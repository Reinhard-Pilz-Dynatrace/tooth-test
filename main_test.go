package main_test

import "testing"

func TestMain(t *testing.T) {
	t.Log("done")
	t.Error("this test has failed")
	t.FailNow()
}
