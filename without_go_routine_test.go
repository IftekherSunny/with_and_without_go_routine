package with_and_without_go_routine

import (
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	expect := "Message: Golang"

	found := Message("Golang")

	if expect != found {
		assert(t, expect, found)
	}
}

func TestWithoutGoRoutine(t *testing.T) {
	expect := 4 * time.Second

	found := WithoutGoRoutine()

	if found > expect {
		assert(t, expect, found)
	}
}
