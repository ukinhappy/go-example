package lockx

import "testing"

func TestNewMutex(t *testing.T) {
	InitCluster(WithAddrs([]string{"172.16.1.112:7000"}))

	l := NewMutex("lockx")
	if l.Lock(3) {
		defer l.UnLock()
		t.Log("locked")
		if !l.Lock(3) {
			t.Log("lock success")
		}
	} else {
		t.Errorf("no lockx")
	}

}
