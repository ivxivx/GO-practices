package json

import "fmt"

type UnmarshalJsonTypeError struct {
	Data string `json:"data"`
	Err  error  `json:"error"`
}

func (e UnmarshalJsonTypeError) Error() string {
	return fmt.Sprintf("error parsing Type `%s`: %v", e.Data, e.Err)
}
