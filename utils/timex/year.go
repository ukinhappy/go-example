package timex

import "time"

//BeginOfYear begin of year
func (t *Timex) BeginOfYear() *Timex {
	return &Timex{
		T: time.Date(t.T.Year(), 1, 1, 0, 0, 0, 0, t.T.Location()),
	}
}

//EndOfYear end of year
func (t *Timex) EndOfYear() *Timex {
	return &Timex{
		T: time.Date(t.T.Year()+1, 1, 1, 0, 0, 0, 0, t.T.Location()),
	}
}
