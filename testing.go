package example

import "testing"

func assert(t *testing.T, expect interface{}, found interface{}) {
	t.Errorf("\nExpect: \n\t%s \nFound \n\t%s", expect, found)
}
