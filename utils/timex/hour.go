package timex

import "time"

//Hour 当前小时
func (t *Timex) Hour() int {
	return int(t.T.Hour())
}

//BeginOfHour begin of hour
func (t *Timex) BeginOfHour() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), t.T.Day(), t.T.Hour(), 0, 0, 0, t.T.Location()),
	}

}

//EndOfHour
func (t *Timex) EndOfHour() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), t.T.Month(), t.T.Day()+1, t.T.Hour()+1, 0, 0, 0, t.T.Location()),
	}
}
