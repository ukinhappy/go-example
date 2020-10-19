/*
timex 对time系统包的扩展
*/
package timex

import "time"

//Timex time 扩展
type Timex struct {
	T time.Time
}

//Now 当前时间
func Now() *Timex {
	return &Timex{
		T: time.Now(),
	}
}

//Timer t的时间
func Timer(t time.Time) *Timex {
	return &Timex{
		T: t,
	}
}

//String 2019-0909 10:10:10
func String(t string) *Timex {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return &Timex{
		T: ti,
	}
}

//Unix 111111
func Unix(t int64) *Timex {
	return &Timex{
		T: time.Unix(t, 0),
	}
}

// ----------------------------------------------------------------------
func (t *Timex) UnixSecond() int64 {
	return t.T.Unix()
}

//UnixMilliSecond 毫秒
func (t *Timex) UnixMilliSecond() int64 {
	return t.T.UnixNano() / 1e6
}

//UniMicrosecond 微妙
func (t *Timex) UniMicrosecond() int64 {
	return t.T.UnixNano() / 1e3
}

//UnixNano 纳秒
func (t *Timex) UnixNano() int64 {
	return t.T.UnixNano()
}

// String 2019-09-09 10:00:00
func (t *Timex) String() string {
	return t.T.Format("2006-01-02 15:04:05")
}
