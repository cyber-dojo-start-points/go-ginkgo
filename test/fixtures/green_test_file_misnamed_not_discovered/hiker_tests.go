package hiker

import "testing"

// go test collects test functions from files whose names end _test.go. This
// file ends _tests.go, so it is compiled as ordinary source and this function
// is never called. It asserts three digits, which would fail if it ran.
//
// A ginkgo spec is different: Describe registers at package init, so a spec in
// this file would run from here. See red_spec_in_misnamed_file_still_runs.
func Test_answer_is_three_digits(t *testing.T) {
	if have := answer(); have < 100 || have > 999 {
		t.Errorf("answer: have %d; want: three digits", have)
	}
}
