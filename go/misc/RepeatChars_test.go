package misc

import "testing"

func TestRepeatChars(t *testing.T) {
	if output := RepeatChars("-", 1); output != "-" {
		t.Fail()
	}
	if output := RepeatChars("-", 2); output != "--" {
		t.Fail()
	}
	if output := RepeatChars("-", 3); output != "---" {
		t.Fail()
	}
	if output := RepeatChars("-", 4); output != "----" {
		t.Fail()
	}
	if output := RepeatChars("-", 5); output != "-----" {
		t.Fail()
	}
	if output := RepeatChars("-", 6); output != "------" {
		t.Fail()
	}
	if output := RepeatChars("-", 7); output != "-------" {
		t.Fail()
	}
	if output := RepeatChars("-", 8); output != "--------" {
		t.Fail()
	}
	if output := RepeatChars("-", 0); output != "" {
		t.Fail()
	}
	if output := RepeatChars("-", -1); output != "" {
		t.Fail()
	}
	if output := RepeatChars("-", -2); output != "" {
		t.Fail()
	}
	if output := RepeatChars("-", -3); output != "" {
		t.Fail()
	}
	if output := RepeatChars("-", -4); output != "" {
		t.Fail()
	}
	if output := RepeatChars("-", -5); output != "" {
		t.Fail()
	}

}
