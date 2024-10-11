package json

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/ivxivx/go-practices/util"
)

func Test_Omitempty_Nil(t *testing.T) {
	t.Parallel()

	record := Record{
		ID:          "id1",
		Description: nil,
	}

	marshaled, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"id":"id1","amount":"0","currency":""}`
	actual := string(marshaled)

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_Omitempty_Empty(t *testing.T) {
	t.Parallel()

	record := Record{
		ID:          "id1",
		Description: util.ToPointer(""),
	}

	marshaled, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"id":"id1","amount":"0","currency":"","description":""}`
	actual := string(marshaled)

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_NoOmitempty_Nil(t *testing.T) {
	t.Parallel()

	record := Record{
		ID:    "id1",
		Label: nil,
	}

	marshaled, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"id":"id1","amount":"0","currency":"","label":null}`
	actual := string(marshaled)

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_NoOmitempty_Empty(t *testing.T) {
	t.Parallel()

	record := Record{
		ID:    "id1",
		Label: util.ToPointer(""),
	}

	marshaled, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"id":"id1","amount":"0","currency":"","label":""}`
	actual := string(marshaled)

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}
