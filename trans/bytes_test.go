package trans

import "testing"

const test_string = `goroutine 43 [running]:`

var test_bytes = []byte(test_string)

func TestUnsafeBytesToString(t *testing.T) {
	target := UnsafeBytesToString(test_bytes)
	if target != test_string {
		t.Errorf("ERROR: %s", target)
	}
}

func TestReflectStringToBytes(t *testing.T) {
	target := string(UnsafeStringToBytes(test_string))
	if target != test_string {
		t.Errorf("ERROR: %s", target)
	}
}
