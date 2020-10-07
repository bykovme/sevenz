package sevenz

import "testing"

func TestCheck7zAvailability(t *testing.T) {
	err := Check7zAvailability()
	if err != nil {
		t.Error(err)
	}
}
