package timex

import (
	"fmt"
	"time"
)

// WeekOfBegin 获取本周开始
func (t *Timex) BeginOfWeek() *Timex {
	return t.OneDayOfWeek(time.Monday)
}

// WeekOfEnd 获取本周结束的时间戳
func (t *Timex) EndOfWeek() *Timex {
	return &Timex{
		T: t.BeginOfWeek().T.AddDate(0, 0, 7),
	}
}

// ISOWeek 获取本周是今年的第几周
func (t *Timex) ISOWeek() string {
	year, index := t.T.ISOWeek()
	return fmt.Sprintf("%04d%02d", year, index)
}

// LastISOWeek 获取上周是今年的第几周
func (t *Timex) LastISOWeek() string {
	lastWeek := t.T.Add(-7 * 24 * time.Hour)
	year, index := lastWeek.ISOWeek()
	return fmt.Sprintf("%04d%02d", year, index)
}

// OneDayOfWeek 一周中某天的时间戳
func (t *Timex) OneDayOfWeek(weekday time.Weekday) *Timex {
	var offset int
	if t.T.Weekday() == time.Sunday {
		offset = int(time.Weekday(7) - weekday)
	} else {
		offset = int(t.T.Weekday() - weekday)
	}
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), t.T.Day(), 0, 0, 0, 0, t.T.Location()).AddDate(0, 0, -offset),
	}
}
