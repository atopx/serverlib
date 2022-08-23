package trans

import "time"

const (
	DATE_YMDHMS_STORT = "20060102"
	DATE_YMD          = "2006-01-02"
	DATE_YMDHMS       = "2006-01-02 15:04:05"
	DATE_YMDHMS_DOT   = "15:04:05.000"
)

func TimeFormatOnNow(layout string) string {
	return time.Now().Local().Format(layout)
}

func TimeFormatOnUnix(unix int64, layout string) string {
	return time.Unix(unix, 0).Local().Format(layout)
}
