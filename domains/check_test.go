package domains

import (
	"testing"
)

func Test_for_Verify(t *testing.T) {
	err := Verify("cara")

	if err != nil {
		t.Fatalf("expected no error, got %v", err.Error())
	}
}

func Test_for_ToLower(t *testing.T) {
	actual := ToLower("TESTING FoR LowER cASe")
	expect := "testing for lower case"

	if actual != expect {
		t.Fatalf("expected %v, got %v", expect, actual)
	}
}
