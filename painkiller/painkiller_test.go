package painkiller

import (
	"fmt"
	"testing"
)

func TestPill_String(t *testing.T) {
	if fmt.Sprintf("%s", Aspirin) != "Aspirin" {
		t.Error("oops")
	}
}