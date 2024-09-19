package csv

import (
	"flag"
	"os"
	"testing"

	"github.com/shopspring/decimal"
	"go.uber.org/goleak"
)

type Record struct {
	ID       string          `csv:"id"`
	Amount   decimal.Decimal `csv:"amount"`
	Currency string          `csv:"currency"`
}

var records = []Record{
	{
		ID:       "id1",
		Amount:   decimal.NewFromInt(100),
		Currency: "USD",
	},
	{
		ID:       "id2",
		Amount:   decimal.NewFromInt(200),
		Currency: "EUR",
	},
}

func TestMain(m *testing.M) {
	leak := flag.Bool("leak", false, "use leak detector")
	flag.Parse()

	if *leak {
		goleak.VerifyTestMain(m,
			goleak.IgnoreTopFunction("net/http.(*persistConn).writeLoop"),
			goleak.IgnoreTopFunction("internal/poll.runtime_pollWait"),
			goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*recursiveTree).dispatch"),
		)

		return
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}
