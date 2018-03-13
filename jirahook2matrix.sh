#!/bin/sh
#
#
#
# PROVIDE: jirahook2matrix
# REQUIRE: NETWORK
# KEYWORD: shutdown
#

. /etc/rc.subr

#
# to create local user webhook to run this service:
# NAME=webhook ID=8082 && pw groupadd -n $NAME -g $ID && pw useradd -n $NAME -g $NAME -u $ID -s /bin/sh -w no -d /home/$NAME -m -M 750
#

user=webhook
name=jirahook2matrix
logfile=/var/log/${name}.log
rcvar=${name}_enable

procname="/usr/local/bin/${name}"
procargs=" >/var/log/${name}.log 2>&1"
start_cmd="my_start"

load_rc_config $name

: ${jirahook2matrix_enable:="NO"}

my_start()
{
  touch $logfile
  chown ${user} $logfile
  cd "/home/${user}"
  echo "Starting ${name} as user ${user}"
  su ${user} -c "${procname} ${procargs} &"
}

run_rc_command "$1"
