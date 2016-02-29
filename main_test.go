package main_test

import (
	"flag"
	"os"
	"testing"
)

func Test(t *testing.T) {
	expected := 0
	actual := 0
	if expected != actual {
		t.Errorf("%v is not %v :'(", actual, expected)
	}
}

func SetupTest() {
	// ...
}

func TearDownTest() {
	// ...
}

func TestMain(m *testing.M) {
	flag.Parse()
	SetupTest()
	res := m.Run()
	TearDownTest()
	os.Exit(res)
}

func BenchmarkMain(b *testing.B) {
	// ...
}
