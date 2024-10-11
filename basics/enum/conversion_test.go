package enum

import (
	"slices"
	"testing"
)

func Test_Enum_Convert(t *testing.T) {
	t.Parallel()

	expected := "invalid"
	actual := Name("invalid")

	if actual != "invalid" {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_Enum_Contain(t *testing.T) {
	t.Parallel()

	allNames := []Name{ZhangSan, LiShi}

	invalid := Name("invalid")
	if slices.Contains(allNames, invalid) {
		t.Fatalf("expected %v does not contain %s", allNames, invalid)
	}

	valid := Name("ZhangSan")
	if !slices.Contains(allNames, valid) {
		t.Fatalf("expected %v contains %s", allNames, valid)
	}

	if !slices.Contains(allNames, "LiShi") {
		t.Fatalf("expected %v contains %s", allNames, "LiShi")
	}
}
