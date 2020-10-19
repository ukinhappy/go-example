package timex

import (
	"testing"
	"time"
)

func TestTimex_UnixSecond(t *testing.T) {
	t.Log(Now().UnixSecond())
}

func TestTimex_UnixMilliSecond(t *testing.T) {
	t.Log(Now().UnixMilliSecond())
}

func TestTimex_UniMicrosecond(t *testing.T) {
	t.Log(Now().UniMicrosecond())
}

func TestTimex_BeginOfDate(t *testing.T) {
	t.Log(Now().BeginOfDate().String())
}

func TestTimex_BeginOfHour(t *testing.T) {
	t.Log(Now().BeginOfHour().String())
}

func TestTimex_BeginOfMonth(t *testing.T) {
	t.Log(Now().BeginOfMonth().String())
}

func TestTimex_BeginOfYear(t *testing.T) {
	t.Log(Now().BeginOfYear().String())
}

func TestTimex_Day(t *testing.T) {
	t.Log(Now().Day())
}

func TestTimex_EndOfDate(t *testing.T) {
	t.Log(Now().EndOfDate().String())
}

func TestTimex_UnixNano(t *testing.T) {
	t.Log(Now().UnixNano())
}

func TestTimex_EndOfHour(t *testing.T) {
	t.Log(Now().EndOfHour().String())
}

func TestTimex_EndOfMonth(t *testing.T) {
	t.Log(Now().EndOfMonth().String())
}

func TestTimex_EndOfYear(t *testing.T) {
	t.Log(Now().EndOfYear().String())
}
func TestTimex_Hour(t *testing.T) {
	t.Log(Now().Hour())
}

func TestTimex_ISOWeek(t *testing.T) {
	t.Log(Now().ISOWeek())
}

func TestTimex_LastISOWeek(t *testing.T) {
	t.Log(Now().LastISOWeek())
}

func TestTimex_Month(t *testing.T) {
	t.Log(Now().Month())
}

func TestTimex_RemainingOfDate(t *testing.T) {
	t.Log(Now().RemainingOfDate())
}

func TestTimex_BeginOfWeek(t *testing.T) {
	t.Log(Timer(time.Date(Now().T.Year(), Now().T.Month(), 7, 0, 0, 0, 0, Now().T.Location())).BeginOfWeek())
}
