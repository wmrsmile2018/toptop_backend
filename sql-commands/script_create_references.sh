red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
purple=`tput setaf 5`
reset=`tput sgr0`

echo "${yellow}create references${purple}"
PGGSSENCMODE=disable psql -h localhost -d toptop -U 19545303 -p 5432 -a -q -f /Users/19545303/project/golang/top-top/sql-commands/sql_create_references.sql
echo "${reset}"
