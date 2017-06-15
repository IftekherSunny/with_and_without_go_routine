package with_and_without_go_routine

import (
	"testing"
	"time"
)

func TestMessageWithGoRoutine(t *testing.T) {
	expect := "Message: Golang"

	found := <-MessageWithGoRoutine("Golang")

	if expect != found {
		assert(t, expect, found)
	}
}

func TestWithGoRoutine(t *testing.T) {
	expect := 2 * time.Second

	found := WithGoRoutine()

	if found > expect {
		assert(t, expect, found)
	}
}
