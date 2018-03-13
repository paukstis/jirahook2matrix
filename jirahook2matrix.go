package main

import(
  "os"
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "bytes"
  "io/ioutil"
  "encoding/json"
)

type Configuration struct {
  DEBUG            string  `json:"DEBUG"`
  ChangelogTrigger string  `json:"ChangelogTrigger"`
  ListenPort       string  `json:"ListenPort"`
  ListenPath       string  `json:"ListenPath"`
  MatrixURL        string  `json:"MatrixURL"`
  Room             string  `json:"Room"`
  Token            string  `json:"Token"`
  Type             string  `json:"Type"`
  AllowedIP        string  `json:"AllowedIP"`
  IssueUrlPrefix   string  `json:"IssueUrlPrefix"`
  ReporterPrefix   string  `json:"ReporterPrefix"`
  SummaryPrefix    string  `json:"SummaryPrefix"`
}
// Event represents an incoming JIRA webhook event
type Event struct {
  WebhookEvent string       `json:"webhookEvent"`
  Timestamp    int64        `json:"timestamp"`
  User         User         `json:"user"`
  Issue        Issue        `json:"issue"`
  Changelog    Changelog    `json:"changelog"`
}
type MatrixPostBody struct {
  Msgtype string  `json:"msgtype"`
  Body    string  `json:"body"`
}
var conf Configuration
var version string


// =======================================================
func main() {
  file, _ := os.Open("conf.json")
  decoder := json.NewDecoder(file)
  err := decoder.Decode(&conf)
  if err != nil {
    fmt.Println("configuration file error")
    panic(err)
  }

  mux := http.NewServeMux()
  mux.HandleFunc(conf.ListenPath, hookHandler)
  fmt.Printf("jirahook2matrix version %s starting, listen on port %s\n", version, conf.ListenPort)
  http.ListenAndServe(conf.ListenPort, mux)
  fmt.Printf("jirahook2matrix version %s stop", version)
}

// =======================================================
func hookHandler(w http.ResponseWriter, r *http.Request) {
  var whe Event

  s := strings.Split(r.RemoteAddr, ":")
  ip, _ := s[0], s[1]    

  if ip != conf.AllowedIP {
    fmt.Println("IP not allowed")
    w.WriteHeader(http.StatusForbidden)
    return
  }
  
  err := json.NewDecoder(r.Body).Decode(&whe)
  if err != nil {
    panic(err)
  }

  if whe.Issue.Self == "" {
    fmt.Println("Bad request, no Issue.Self")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // Prepare URL for JIRA issue browse
  u, _ := url.Parse(whe.Issue.Self)
  issueurl := fmt.Sprintf("%s://%s/browse/%s", u.Scheme, u.Host, whe.Issue.Key)

  //msgtype := conf.Type
  summary, reporter := "", ""
  if whe.Issue.Fields == nil {
    fmt.Println("Bad request, no Issue.Fields")
    w.WriteHeader(http.StatusBadRequest)
    return
  } else {
    summary = whe.Issue.Fields.Summary
    reporter = whe.Issue.Fields.Reporter.DisplayName
  }

  // additional trigger on changelog field if defined in configuration
  endnow := false
  if conf.ChangelogTrigger != "" {
    endnow = true
    if whe.Changelog.Items != nil {
      for index := range whe.Changelog.Items {
        if whe.Changelog.Items[index].Field == conf.ChangelogTrigger {
          endnow = false
        }
      }
    }
  }
  if endnow {
    w.WriteHeader(http.StatusOK)
    return
  }
  
  // Prepare notification message and URL for POST
  postURL := fmt.Sprintf("%s/%s/send/m.room.message?access_token=%s", conf.MatrixURL, conf.Room, conf.Token)
  bodyval := fmt.Sprintf("%s%s%s%s%s%s", conf.IssueUrlPrefix, issueurl, conf.ReporterPrefix, reporter, conf.SummaryPrefix, summary)

  // Prepare Matrix JSON
  mbo := MatrixPostBody { Msgtype: conf.Type, Body: bodyval }
  mbojson, err := json.Marshal(mbo)
  if err != nil {
    panic(err)
  }
  
  // POST to Matrix or just printf if DEBUG YES
  if conf.DEBUG == "YES" {
    fmt.Printf("POST_URL: %s\nPOST_BODY: %+v\n", postURL, string(mbojson))
    w.WriteHeader(http.StatusOK)
  } else {
    resp, err := http.Post(postURL, "application/x-www-form-urlencoded", bytes.NewReader(mbojson))
    if err != nil {
      fmt.Println("Bad gateway")
      w.WriteHeader(http.StatusBadGateway)
      return
    }

    if resp.StatusCode/100 == 2 {
      w.WriteHeader(http.StatusOK)
      return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("POST response body=%s\n",string(body))

    w.WriteHeader(http.StatusServiceUnavailable)
  }
}
