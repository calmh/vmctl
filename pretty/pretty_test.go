package pretty

import (
	"testing"
	"time"
)

type testStruct struct {
	Hello   string `pretty:"Hello There"`
	Second  int    `pretty:"seconds" format:"%05d"`
	Second2 int    `pretty:"-" format:"%05d"`
	Struct  struct {
		Foo string
		Bar time.Time
	}
}

func TestPretty(t *testing.T) {
	ts := testStruct{
		Hello:   "foo",
		Second:  42,
		Second2: 33,
		Struct: struct {
			Foo string
			Bar time.Time
		}{
			Foo: "bar",
			Bar: time.Now(),
		},
	}

	Print(ts)
}
