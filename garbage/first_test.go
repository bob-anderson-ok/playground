package garbage

import "testing"

func TestMessage(t *testing.T) {
	var ans = Judy()
	if ans != "This is Judy returning your call" {
		t.Errorf("Judy said: %s", ans)
	}
}
