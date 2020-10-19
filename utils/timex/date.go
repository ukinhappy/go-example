package timex

import (
	"time"
)

// Day 当前天
func (t *Timex) Day() int {
	return t.T.Day()
}

// BeginOfDate begin of date
func (t *Timex) BeginOfDate() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), t.T.Day(), 0, 0, 0, 0, t.T.Location()),
	}
}

//EndOfDate end of date
func (t *Timex) EndOfDate() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), t.T.Day()+1, 0, 0, 0, 0, t.T.Location()),
	}
}

//RemainingOfDate remaining of date
func (t *Timex) RemainingOfDate() int64 {
	return t.EndOfDate().UnixSecond() - t.UnixSecond()
}
