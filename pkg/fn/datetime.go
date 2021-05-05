package fn

import (
	ns "github.com/HerringtonDarkholme/go-newspeak/pkg"
)

// https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html

func ADDDATE() ns.Expression { return nil }
func ADDTIME() ns.Expression { return nil }

// func CONVERT_TZ() ns.Expression        { return nil }
// func CURDATE() ns.Expression           { return nil }
func CURRENT_DATE() ns.Expression      { return nil }
func CURRENT_TIME() ns.Expression      { return nil }
func CURRENT_TIMESTAMP() ns.Expression { return nil }

// func CURTIME() ns.Expression           { return nil }
func DATE() ns.Expression           { return nil }
func DATE_ADD() ns.Expression       { return nil }
func DATE_FORMAT() ns.Expression    { return nil }
func DATE_SUB() ns.Expression       { return nil }
func DATEDIFF() ns.Expression       { return nil }
func DAY() ns.Expression            { return nil }
func DAYNAME() ns.Expression        { return nil }
func DAYOFMONTH() ns.Expression     { return nil }
func DAYOFWEEK() ns.Expression      { return nil }
func DAYOFYEAR() ns.Expression      { return nil }
func EXTRACT() ns.Expression        { return nil }
func FROM_DAYS() ns.Expression      { return nil }
func FROM_UNIXTIME() ns.Expression  { return nil }
func GET_FORMAT() ns.Expression     { return nil }
func HOUR() ns.Expression           { return nil }
func LAST_DAY() ns.Expression       { return nil }
func LOCALTIME() ns.Expression      { return nil }
func LOCALTIMESTAMP() ns.Expression { return nil }
func MAKEDATE() ns.Expression       { return nil }
func MAKETIME() ns.Expression       { return nil }
func MICROSECOND() ns.Expression    { return nil }
func MINUTE() ns.Expression         { return nil }
func MONTH() ns.Expression          { return nil }
func MONTHNAME() ns.Expression      { return nil }
func NOW() ns.Expression            { return nil }
func PERIOD_ADD() ns.Expression     { return nil }
func PERIOD_DIFF() ns.Expression    { return nil }
func QUARTER() ns.Expression        { return nil }
func SEC_TO_TIME() ns.Expression    { return nil }
func SECOND() ns.Expression         { return nil }
func STR_TO_DATE() ns.Expression    { return nil }
func SUBDATE() ns.Expression        { return nil }
func SUBTIME() ns.Expression        { return nil }
func SYSDATE() ns.Expression        { return nil }
func TIME() ns.Expression           { return nil }
func TIME_FORMAT() ns.Expression    { return nil }
func TIME_TO_SEC() ns.Expression    { return nil }
func TIMEDIFF() ns.Expression       { return nil }
func TIMESTAMP() ns.Expression      { return nil }
func TIMESTAMPADD() ns.Expression   { return nil }
func TIMESTAMPDIFF() ns.Expression  { return nil }
func TO_DAYS() ns.Expression        { return nil }
func TO_SECONDS() ns.Expression     { return nil }
func UNIX_TIMESTAMP() ns.Expression { return nil }
func UTC_DATE() ns.Expression       { return nil }
func UTC_TIME() ns.Expression       { return nil }
func UTC_TIMESTAMP() ns.Expression  { return nil }
func WEEK() ns.Expression           { return nil }
func WEEKDAY() ns.Expression        { return nil }
func WEEKOFYEAR() ns.Expression     { return nil }
func YEAR() ns.Expression           { return nil }
func YEARWEEK() ns.Expression       { return nil }
