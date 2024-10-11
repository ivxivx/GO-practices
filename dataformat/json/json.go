package json

import (
	"flag"
	"os"
	"testing"

	"github.com/shopspring/decimal"
	"go.uber.org/goleak"
)

type Record struct {
	ID          string          `json:"id"`
	Amount      decimal.Decimal `json:"amount"`
	Currency    string          `json:"currency"`
	Description *string         `json:"description,omitempty"`
	Label       *string         `json:"label"`
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
