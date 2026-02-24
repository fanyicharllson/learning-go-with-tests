package mocking

import (
	"bytes"
	"fmt"
	// "time"
)

const limit = 3
const finalword = "Go!"

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out *bytes.Buffer, sleeper Sleeper) {
	for i := limit; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalword)
}
