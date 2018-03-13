package main

// structs below imported from JIRA Go package https://github.com/andygrunwald/go-jira

import(
  "time"
)

// User represents a JIRA user.
type User struct {
        Self            string     `json:"self,omitempty" structs:"self,omitempty"`
        Name            string     `json:"name,omitempty" structs:"name,omitempty"`
        Password        string     `json:"-"`
        Key             string     `json:"key,omitempty" structs:"key,omitempty"`
        EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
        AvatarUrls      AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
        DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
        Active          bool       `json:"active,omitempty" structs:"active,omitempty"`
        TimeZone        string     `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
        ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}
// ChangelogItems reflects one single changelog item of a history item
type ChangelogItems struct {
        Field      string      `json:"field" structs:"field"`
        FieldType  string      `json:"fieldtype" structs:"fieldtype"`
        From       interface{} `json:"from" structs:"from"`
        FromString string      `json:"fromString" structs:"fromString"`
        To         interface{} `json:"to" structs:"to"`
        ToString   string      `json:"toString" structs:"toString"`
}
// ChangelogHistory reflects one single changelog history entry
type Changelog struct {
        Id      string           `json:"id" structs:"id"`
        Items   []ChangelogItems `json:"items" structs:"items"`
}

// Issue represents a JIRA issue.
type Issue struct {
        Expand    string       `json:"expand,omitempty" structs:"expand,omitempty"`
        ID        string       `json:"id,omitempty" structs:"id,omitempty"`
        Self      string       `json:"self,omitempty" structs:"self,omitempty"`
        Key       string       `json:"key,omitempty" structs:"key,omitempty"`
        Fields    *IssueFields `json:"fields,omitempty" structs:"fields,omitempty"`
}
// Attachment represents a JIRA attachment
type Attachment struct {
        Self      string `json:"self,omitempty" structs:"self,omitempty"`
        ID        string `json:"id,omitempty" structs:"id,omitempty"`
        Filename  string `json:"filename,omitempty" structs:"filename,omitempty"`
        Author    *User  `json:"author,omitempty" structs:"author,omitempty"`
        Created   string `json:"created,omitempty" structs:"created,omitempty"`
        Size      int    `json:"size,omitempty" structs:"size,omitempty"`
        MimeType  string `json:"mimeType,omitempty" structs:"mimeType,omitempty"`
        Content   string `json:"content,omitempty" structs:"content,omitempty"`
        Thumbnail string `json:"thumbnail,omitempty" structs:"thumbnail,omitempty"`
}
// Epic represents the epic to which an issue is associated
// Not that this struct does not process the returned "color" value
type Epic struct {
        ID      int    `json:"id" structs:"id"`
        Key     string `json:"key" structs:"key"`
        Self    string `json:"self" structs:"self"`
        Name    string `json:"name" structs:"name"`
        Summary string `json:"summary" structs:"summary"`
        Done    bool   `json:"done" structs:"done"`
}
// IssueFields represents single fields of a JIRA issue.
// Every JIRA issue has several fields attached.
type IssueFields struct {
        // TODO Missing fields
        //      * "aggregatetimespent": null,
        //      * "workratio": -1,
        //      * "lastViewed": null,
        //      * "aggregatetimeoriginalestimate": null,
        //      * "aggregatetimeestimate": null,
        //      * "environment": null,
        Expand               string        `json:"expand,omitempty" structs:"expand,omitempty"`
        Type                 IssueType     `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
        Project              Project       `json:"project,omitempty" structs:"project,omitempty"`
        Resolution           *Resolution   `json:"resolution,omitempty" structs:"resolution,omitempty"`
        Priority             *Priority     `json:"priority,omitempty" structs:"priority,omitempty"`
        Resolutiondate       Time          `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
        Created              Time          `json:"created,omitempty" structs:"created,omitempty"`
        Duedate              Date          `json:"duedate,omitempty" structs:"duedate,omitempty"`
        Watches              *Watches      `json:"watches,omitempty" structs:"watches,omitempty"`
        Assignee             *User         `json:"assignee,omitempty" structs:"assignee,omitempty"`
        Updated              Time          `json:"updated,omitempty" structs:"updated,omitempty"`
        Description          string        `json:"description,omitempty" structs:"description,omitempty"`
        Summary              string        `json:"summary,omitempty" structs:"summary,omitempty"`
        Creator              *User         `json:"Creator,omitempty" structs:"Creator,omitempty"`
        Reporter             *User         `json:"reporter,omitempty" structs:"reporter,omitempty"`
        Components           []*Component  `json:"components,omitempty" structs:"components,omitempty"`
        Status               *Status       `json:"status,omitempty" structs:"status,omitempty"`
        Progress             *Progress     `json:"progress,omitempty" structs:"progress,omitempty"`
        AggregateProgress    *Progress     `json:"aggregateprogress,omitempty" structs:"aggregateprogress,omitempty"`
        TimeTracking         *TimeTracking `json:"timetracking,omitempty" structs:"timetracking,omitempty"`
        TimeSpent            int           `json:"timespent,omitempty" structs:"timespent,omitempty"`
        TimeEstimate         int           `json:"timeestimate,omitempty" structs:"timeestimate,omitempty"`
        TimeOriginalEstimate int           `json:"timeoriginalestimate,omitempty" structs:"timeoriginalestimate,omitempty"`
        Worklog              *Worklog      `json:"worklog,omitempty" structs:"worklog,omitempty"`
        IssueLinks           []*IssueLink  `json:"issuelinks,omitempty" structs:"issuelinks,omitempty"`
        Comments             *Comments     `json:"comment,omitempty" structs:"comment,omitempty"`
        FixVersions          []*FixVersion `json:"fixVersions,omitempty" structs:"fixVersions,omitempty"`
        Labels               []string      `json:"labels,omitempty" structs:"labels,omitempty"`
        Subtasks             []*Subtasks   `json:"subtasks,omitempty" structs:"subtasks,omitempty"`
        Attachments          []*Attachment `json:"attachment,omitempty" structs:"attachment,omitempty"`
        Epic                 *Epic         `json:"epic,omitempty" structs:"epic,omitempty"`
        Parent               *Parent       `json:"parent,omitempty" structs:"parent,omitempty"`
//        Unknowns             tcontainer.MarshalMap
}
// AvatarUrls represents different dimensions of avatars / images
type AvatarUrls struct {
        Four8X48  string `json:"48x48,omitempty" structs:"48x48,omitempty"`
        Two4X24   string `json:"24x24,omitempty" structs:"24x24,omitempty"`
        One6X16   string `json:"16x16,omitempty" structs:"16x16,omitempty"`
        Three2X32 string `json:"32x32,omitempty" structs:"32x32,omitempty"`
}
// IssueType represents a type of a JIRA issue.
// Typical types are "Request", "Bug", "Story", ...
type IssueType struct {
        Self        string `json:"self,omitempty" structs:"self,omitempty"`
        ID          string `json:"id,omitempty" structs:"id,omitempty"`
        Description string `json:"description,omitempty" structs:"description,omitempty"`
        IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
        Name        string `json:"name,omitempty" structs:"name,omitempty"`
        Subtask     bool   `json:"subtask,omitempty" structs:"subtask,omitempty"`
        AvatarID    int    `json:"avatarId,omitempty" structs:"avatarId,omitempty"`
}

// Resolution represents a resolution of a JIRA issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
        Self        string `json:"self" structs:"self"`
        ID          string `json:"id" structs:"id"`
        Description string `json:"description" structs:"description"`
        Name        string `json:"name" structs:"name"`
}
// Priority represents a priority of a JIRA issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
        Self    string `json:"self,omitempty" structs:"self,omitempty"`
        IconURL string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
        Name    string `json:"name,omitempty" structs:"name,omitempty"`
        ID      string `json:"id,omitempty" structs:"id,omitempty"`
}
// Watches represents a type of how many and which user are "observing" a JIRA issue to track the status / updates.
type Watches struct {
        Self       string     `json:"self,omitempty" structs:"self,omitempty"`
        WatchCount int        `json:"watchCount,omitempty" structs:"watchCount,omitempty"`
        IsWatching bool       `json:"isWatching,omitempty" structs:"isWatching,omitempty"`
        Watchers   []*Watcher `json:"watchers,omitempty" structs:"watchers,omitempty"`
}
// Watcher represents a simplified user that "observes" the issue
type Watcher struct {
        Self        string `json:"self,omitempty" structs:"self,omitempty"`
        Name        string `json:"name,omitempty" structs:"name,omitempty"`
        DisplayName string `json:"displayName,omitempty" structs:"displayName,omitempty"`
        Active      bool   `json:"active,omitempty" structs:"active,omitempty"`
}
// Component represents a "component" of a JIRA issue.
// Components can be user defined in every JIRA instance.
type Component struct {
        Self string `json:"self,omitempty" structs:"self,omitempty"`
        ID   string `json:"id,omitempty" structs:"id,omitempty"`
        Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// Status represents the current status of a JIRA issue.
// Typical status are "Open", "In Progress", "Closed", ...
// Status can be user defined in every JIRA instance.
type Status struct {
        Self           string         `json:"self" structs:"self"`
        Description    string         `json:"description" structs:"description"`
        IconURL        string         `json:"iconUrl" structs:"iconUrl"`
        Name           string         `json:"name" structs:"name"`
        ID             string         `json:"id" structs:"id"`
        StatusCategory StatusCategory `json:"statusCategory" structs:"statusCategory"`
}
// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every JIRA instance.
type StatusCategory struct {
        Self      string `json:"self" structs:"self"`
        ID        int    `json:"id" structs:"id"`
        Name      string `json:"name" structs:"name"`
        Key       string `json:"key" structs:"key"`
        ColorName string `json:"colorName" structs:"colorName"`
}
// Progress represents the progress of a JIRA issue.
type Progress struct {
        Progress int `json:"progress" structs:"progress"`
        Total    int `json:"total" structs:"total"`
}
// Parent represents the parent of a JIRA issue, to be used with subtask issue types.
type Parent struct {
        ID  string `json:"id,omitempty" structs:"id"`
        Key string `json:"key,omitempty" structs:"key"`
}

// Time represents the Time definition of JIRA as a time.Time of go
type Time time.Time
// Date represents the Date definition of JIRA as a time.Time of go
type Date time.Time

// Wrapper struct for search result
type transitionResult struct {
        Transitions []Transition `json:"transitions" structs:"transitions"`
}
// Transition represents an issue transition in JIRA
type Transition struct {
        ID     string                     `json:"id" structs:"id"`
        Name   string                     `json:"name" structs:"name"`
        To     Status                     `json:"to" structs:"status"`
        Fields map[string]TransitionField `json:"fields" structs:"fields"`
}
// TransitionField represents the value of one Transition
type TransitionField struct {
        Required bool `json:"required" structs:"required"`
}

// Option represents an option value in a SelectList or MultiSelect
// custom issue field
type Option struct {
        Value string `json:"value" structs:"value"`
}
// UnmarshalJSON will transform the JIRA time into a time.Time
// during the transformation of the JIRA JSON response
func (t *Time) UnmarshalJSON(b []byte) error {
        // Ignore null, like in the main JSON package.
        if string(b) == "null" {
                return nil
        }
        ti, err := time.Parse("\"2006-01-02T15:04:05.999-0700\"", string(b))
        if err != nil {
                return err
        }
        *t = Time(ti)
        return nil
}
// UnmarshalJSON will transform the JIRA date into a time.Time
// during the transformation of the JIRA JSON response
func (t *Date) UnmarshalJSON(b []byte) error {
        // Ignore null, like in the main JSON package.
        if string(b) == "null" {
                return nil
        }
        ti, err := time.Parse("\"2006-01-02\"", string(b))
        if err != nil {
                return err
        }
        *t = Date(ti)
        return nil
}
// MarshalJSON will transform the Date object into a short
// date string as JIRA expects during the creation of a
// JIRA request
func (t Date) MarshalJSON() ([]byte, error) {
        time := time.Time(t)
        return []byte(time.Format("\"2006-01-02\"")), nil
}

// Worklog represents the work log of a JIRA issue.
// One Worklog contains zero or n WorklogRecords
// JIRA Wiki: https://confluence.atlassian.com/jira/logging-work-on-an-issue-185729605.html
type Worklog struct {
        StartAt    int             `json:"startAt" structs:"startAt"`
        MaxResults int             `json:"maxResults" structs:"maxResults"`
        Total      int             `json:"total" structs:"total"`
        Worklogs   []WorklogRecord `json:"worklogs" structs:"worklogs"`
}
// WorklogRecord represents one entry of a Worklog
type WorklogRecord struct {
        Self             string `json:"self,omitempty" structs:"self,omitempty"`
        Author           *User  `json:"author,omitempty" structs:"author,omitempty"`
        UpdateAuthor     *User  `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
        Comment          string `json:"comment,omitempty" structs:"comment,omitempty"`
        Created          *Time  `json:"created,omitempty" structs:"created,omitempty"`
        Updated          *Time  `json:"updated,omitempty" structs:"updated,omitempty"`
        Started          *Time  `json:"started,omitempty" structs:"started,omitempty"`
        TimeSpent        string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
        TimeSpentSeconds int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
        ID               string `json:"id,omitempty" structs:"id,omitempty"`
        IssueID          string `json:"issueId,omitempty" structs:"issueId,omitempty"`
}
// TimeTracking represents the timetracking fields of a JIRA issue.
type TimeTracking struct {
        OriginalEstimate         string `json:"originalEstimate,omitempty" structs:"originalEstimate,omitempty"`
        RemainingEstimate        string `json:"remainingEstimate,omitempty" structs:"remainingEstimate,omitempty"`
        TimeSpent                string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
        OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty" structs:"originalEstimateSeconds,omitempty"`
        RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty" structs:"remainingEstimateSeconds,omitempty"`
        TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
}
// Subtasks represents all issues of a parent issue.
type Subtasks struct {
        ID     string      `json:"id" structs:"id"`
        Key    string      `json:"key" structs:"key"`
        Self   string      `json:"self" structs:"self"`
        Fields IssueFields `json:"fields" structs:"fields"`
}
// IssueLink represents a link between two issues in JIRA.
type IssueLink struct {
        ID           string        `json:"id,omitempty" structs:"id,omitempty"`
        Self         string        `json:"self,omitempty" structs:"self,omitempty"`
        Type         IssueLinkType `json:"type" structs:"type"`
        OutwardIssue *Issue        `json:"outwardIssue" structs:"outwardIssue"`
        InwardIssue  *Issue        `json:"inwardIssue" structs:"inwardIssue"`
        Comment      *Comment      `json:"comment,omitempty" structs:"comment,omitempty"`
}
// IssueLinkType represents a type of a link between to issues in JIRA.
// Typical issue link types are "Related to", "Duplicate", "Is blocked by", etc.
type IssueLinkType struct {
        ID      string `json:"id,omitempty" structs:"id,omitempty"`
        Self    string `json:"self,omitempty" structs:"self,omitempty"`
        Name    string `json:"name" structs:"name"`
        Inward  string `json:"inward" structs:"inward"`
        Outward string `json:"outward" structs:"outward"`
}
// Comments represents a list of Comment.
type Comments struct {
        Comments []*Comment `json:"comments,omitempty" structs:"comments,omitempty"`
}
// Comment represents a comment by a person to an issue in JIRA.
type Comment struct {
        ID           string            `json:"id,omitempty" structs:"id,omitempty"`
        Self         string            `json:"self,omitempty" structs:"self,omitempty"`
        Name         string            `json:"name,omitempty" structs:"name,omitempty"`
        Author       User              `json:"author,omitempty" structs:"author,omitempty"`
        Body         string            `json:"body,omitempty" structs:"body,omitempty"`
        UpdateAuthor User              `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
        Updated      string            `json:"updated,omitempty" structs:"updated,omitempty"`
        Created      string            `json:"created,omitempty" structs:"created,omitempty"`
        Visibility   CommentVisibility `json:"visibility,omitempty" structs:"visibility,omitempty"`
}
// FixVersion represents a software release in which an issue is fixed.
type FixVersion struct {
        Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
        ID              string `json:"id,omitempty" structs:"id,omitempty"`
        Name            string `json:"name,omitempty" structs:"name,omitempty"`
        ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"`
        ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
        Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
        Self            string `json:"self,omitempty" structs:"self,omitempty"`
        UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
}
// CommentVisibility represents he visibility of a comment.
// E.g. Type could be "role" and Value "Administrators"
type CommentVisibility struct {
        Type  string `json:"type,omitempty" structs:"type,omitempty"`
        Value string `json:"value,omitempty" structs:"value,omitempty"`
}

// ProjectCategory represents a single project category
type ProjectCategory struct {
        Self        string `json:"self" structs:"self,omitempty"`
        ID          string `json:"id" structs:"id,omitempty"`
        Name        string `json:"name" structs:"name,omitempty"`
        Description string `json:"description" structs:"description,omitempty"`
}
// Project represents a JIRA Project.
type Project struct {
        Expand       string             `json:"expand,omitempty" structs:"expand,omitempty"`
        Self         string             `json:"self,omitempty" structs:"self,omitempty"`
        ID           string             `json:"id,omitempty" structs:"id,omitempty"`
        Key          string             `json:"key,omitempty" structs:"key,omitempty"`
        Description  string             `json:"description,omitempty" structs:"description,omitempty"`
        Lead         User               `json:"lead,omitempty" structs:"lead,omitempty"`
        Components   []ProjectComponent `json:"components,omitempty" structs:"components,omitempty"`
        IssueTypes   []IssueType        `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
        URL          string             `json:"url,omitempty" structs:"url,omitempty"`
        Email        string             `json:"email,omitempty" structs:"email,omitempty"`
        AssigneeType string             `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
        Versions     []Version          `json:"versions,omitempty" structs:"versions,omitempty"`
        Name         string             `json:"name,omitempty" structs:"name,omitempty"`
        Roles        struct {
                Developers string `json:"Developers,omitempty" structs:"Developers,omitempty"`
        } `json:"roles,omitempty" structs:"roles,omitempty"`
        AvatarUrls      AvatarUrls      `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
        ProjectCategory ProjectCategory `json:"projectCategory,omitempty" structs:"projectCategory,omitempty"`
}
// ProjectComponent represents a single component of a project
type ProjectComponent struct {
        Self                string `json:"self" structs:"self,omitempty"`
        ID                  string `json:"id" structs:"id,omitempty"`
        Name                string `json:"name" structs:"name,omitempty"`
        Description         string `json:"description" structs:"description,omitempty"`
        Lead                User   `json:"lead,omitempty" structs:"lead,omitempty"`
        AssigneeType        string `json:"assigneeType" structs:"assigneeType,omitempty"`
        Assignee            User   `json:"assignee" structs:"assignee,omitempty"`
        RealAssigneeType    string `json:"realAssigneeType" structs:"realAssigneeType,omitempty"`
        RealAssignee        User   `json:"realAssignee" structs:"realAssignee,omitempty"`
        IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid" structs:"isAssigneeTypeValid,omitempty"`
        Project             string `json:"project" structs:"project,omitempty"`
        ProjectID           int    `json:"projectId" structs:"projectId,omitempty"`
}
// Version represents a single release version of a project
type Version struct {
        Self            string `json:"self,omitempty" structs:"self,omitempty"`
        ID              string `json:"id,omitempty" structs:"id,omitempty"`
        Name            string `json:"name,omitempty" structs:"name,omitempty"`
        Description     string `json:"description,omitempty" structs:"name,omitempty"`
        Archived        bool   `json:"archived,omitempty" structs:"archived,omitempty"`
        Released        bool   `json:"released,omitempty" structs:"released,omitempty"`
        ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
        UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
        ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
}
