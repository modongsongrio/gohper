package time2

import (
	"testing"
	"time"

	"github.com/cosiner/gohper/testing2"
)

func TestFunc(t *testing.T) {
	tt := testing2.Wrap(t)
	type Test struct {
		Time  int64
		Human string
	}

	tests := []Test{
		Test{0, "0ns"},
		Test{999, "999ns"},

		Test{1000, "1us"},
		Test{1499, "1us"},
		Test{1500, "2us"},

		Test{1000 * 1000, "1ms"},
		Test{1000 * 1499, "1ms"},
		Test{1000 * 1500, "2ms"},

		Test{1000 * 1000 * 1000, "1s"},
		Test{1000 * 1000 * 1499, "1s"},
		Test{1000 * 1000 * 1500, "2s"},

		Test{1000 * 1000 * 1000 * 10000, "10000s"},
	}

	for _, test := range tests {
		tt.Eq(test.Human, ToHuman(test.Time))
	}
}

func TestTiming(t *testing.T) {
	tt := testing2.Wrap(t)

	tt.True(Timing(func() {
		time.Sleep(time.Millisecond)
	})/int64(time.Millisecond) == 1)
}
