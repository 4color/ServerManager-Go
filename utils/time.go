package utils

import (
	"strings"
	"time"
)

func init() {
	//调整时区
	timelocal := time.FixedZone("CST", 3600*8)
	time.Local = timelocal
}

func Timeformat(oldtime string) string {

	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", oldtime, time.Local)
	return tt.Format("2006-01-02 15:04:05")

}

func Timenow() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")

}

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

//序列化
func (t Time) MarshalJSON() ([]byte, error) {

	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	if t.Year() > 1000 {
		b := time.Time(t).AppendFormat(b, timeFormart)
		b = append(b, '"')
		return b, nil
	} else {
		b = append(b, '"')
		return b, nil
	}
}

//反序列化
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	newstring := string(data)
	var err error
	if strings.Index(newstring, ".") != -1 {
		newstring = newstring[0:strings.Index(newstring, ".")]
	}

	if strings.Index(newstring, "T") != -1 {
		newstring = strings.Replace(newstring, "T", " ", -1)
	}
	newstring = strings.Replace(newstring, "\"", "", -1)
	if newstring != "" {
		tt, err2 := time.ParseInLocation(timeFormart, newstring, time.Local)
		if err2 != nil {
			return err2
		} else {
			*t = Time(tt)
		}
	}
	return err
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t Time) Date() string {
	return time.Time(t).Format("2006-01-02")
}

func (t Time) Year() int {
	inty := time.Time(t).Year()
	return inty
}

func (t Time) Now() Time {
	tt := Time(time.Now())
	return tt
}
