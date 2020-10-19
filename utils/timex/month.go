package timex

import "time"

//Month 当前月份
func (t *Timex) Month() int {
	return int(t.T.Month())
}

//BeginOfMonth begin of month
func (t *Timex) BeginOfMonth() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), 1, 0, 0, 0, 0, t.T.Location()),
	}
}

//EndOfMonth end of month
func (t *Timex) EndOfMonth() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month()+1, 1, 0, 0, 0, 0, t.T.Location()),
	}
}
