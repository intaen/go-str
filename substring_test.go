package substr

import "testing"

func TestSubStr(t *testing.T) {
	data := "Test Substring"
	strlen := 5
	cutlen := 6
	expected := "Sub"
	res := SubStr(data, strlen, cutlen)
	if res != expected {
		t.Errorf("This function doesn't work, if '%s' got cut for '%d', the rest of string must be '%s' not '%s'", data, cutlen, expected, res)
	}
}
