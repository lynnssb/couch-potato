/**
 * @author:       wangxuebing
 * @fileName:     timeutil.go
 * @date:         2022/11/5 22:13
 * @description:
 */

package timeutil

import (
	"couch-potato/common/utils/characterutil"
	"fmt"
	"time"
)

func TimeMilliFormat(n int64) string {
	var timeFormat string
	if n == 0 {
		timeFormat = "-"
	} else {
		tm := time.Unix(n/1e3, 0)
		timeFormat = tm.Format("2006-01-02 15:04:05")
	}
	return timeFormat
}

func TimeFormat(n *time.Time) string {
	var timeFormat string
	if n == nil {
		timeFormat = "——"
	} else {
		timeFormat = n.Format("2006-01-02 15:04:05")
	}
	return timeFormat
}

func TimeFormatYMD(n *time.Time) string {
	var timeFormat string
	if n == nil {
		timeFormat = "—-"
	} else {
		timeFormat = n.Format("2006-01-02")
	}
	return timeFormat
}

func TimeCode(num int64) string {
	code := characterutil.StitchingBufStr(time.Now().Format("20060102"), fmt.Sprintf("%04d", num))
	return code
}
