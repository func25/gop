package gopnet

import (
	"fmt"
	"testing"

	"github.com/google/go-querystring/query"
)

type TestQuery struct {
	Int    int      `url:"int"`
	String string   `url:"string"`
	Array  []string `url:"arr" del:","`
}

func TestToQuery(t *testing.T) {
	v, err := query.Values(nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v.Encode())
}
