package timing

import (
	"strings"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mathHelpers"
)

type Stopwatch struct {
	domain   string
	activity string
	notes    string
	start    time.Time
	end      time.Time
	duration time.Duration
}

func Start(domain, activity, notes string) Stopwatch {
	return Stopwatch{domain: domain, activity: activity, start: time.Now(), notes: notes, end: time.Time{}, duration: time.Duration(0)}
}

func (w *Stopwatch) Stop(count int) {
	w.end = time.Now()
	w.duration = w.end.Sub(w.start)
	logHandler.Timing.Printf("Domain=[%v] Activity=[%v] Notes=[%v] Count=[%v] Duration=[%v]", w.domain, strings.ToUpper(w.activity), w.notes, count, w.duration)
}

// SnoozeFor snoozes the application for a given amount of time
// The function SnoozeFor takes in a polling interval and calls the snooze function with that interval.
func SnoozeFor(noSeconds int) {
	snooze(noSeconds)
}

// Snooze snoozes for a random period
// The Snooze function generates a random number between 0 and 10 and then calls the snooze function
func Snooze() {
	snooze(mathHelpers.RandomInt(10))
}

func snooze(noSeconds int) {
	//pollingInterval, _ := strconv.Atoi(inPollingInterval)
	logHandler.Info.Printf("Snooze... Zzzzzz.... snoozing for %d seconds...", noSeconds)
	time.Sleep(time.Duration(noSeconds) * time.Second)
}
