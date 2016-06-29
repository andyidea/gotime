// timer.go

package gotime

import (
	"time"
)

// 每日24点（次日0点）的timer
func Start24Timer(f func()) {
	go func() {
		for {
			now := time.Now()
			next := GetToday24time()
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}
