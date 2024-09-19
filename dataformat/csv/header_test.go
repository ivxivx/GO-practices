package csv

import (
	"bytes"
	"encoding/csv"
	"testing"

	"github.com/jszwec/csvutil"
)

func Test_WriteHeader_WithEncoder_Auto(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte{})

	writer := csv.NewWriter(buffer)

	encoder := csvutil.NewEncoder(writer)

	encoder.SetHeader([]string{"amount", "currency"})

	// First call to Encode will write a header
	if err := encoder.Encode(records); err != nil {
		t.Fatal(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	expected := "amount,currency\n100,USD\n200,EUR\n"
	actual := buffer.String()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_WriteHeader_WithEncoder_Manual1(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte{})

	writer := csv.NewWriter(buffer)

	encoder := csvutil.NewEncoder(writer)

	encoder.SetHeader([]string{"amount", "currency"})
	encoder.AutoHeader = false

	// manually write header
	encoder.EncodeHeader(&Record{})

	if err := encoder.Encode(records); err != nil {
		t.Fatal(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	expected := "amount,currency\n100,USD\n200,EUR\n"
	actual := buffer.String()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_WriteHeader_WithEncoder_Manual2(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte{})

	writer := csv.NewWriter(buffer)

	encoder := csvutil.NewEncoder(writer)

	encoder.SetHeader([]string{"amount", "currency"})
	encoder.AutoHeader = false

	if err := encoder.Encode(records); err != nil {
		t.Fatal(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	// no header
	expected := "100,USD\n200,EUR\n"
	actual := buffer.String()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_WriteHeader_WithoutEncoder(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte{})

	writer := csv.NewWriter(buffer)

	// manually write header
	writer.Write([]string{"amount", "currency"})

	for _, record := range records {
		if err := writer.Write([]string{record.Amount.String(), record.Currency}); err != nil {
			t.Fatal(err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	expected := "amount,currency\n100,USD\n200,EUR\n"
	actual := buffer.String()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func Test_WriteHeader_DifferentOrder(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte{})

	writer := csv.NewWriter(buffer)

	encoder := csvutil.NewEncoder(writer)

	encoder.SetHeader([]string{"currency", "amount"})

	if err := encoder.Encode(records); err != nil {
		t.Fatal(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	expected := "currency,amount\nUSD,100\nEUR,200\n"
	actual := buffer.String()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}
