package waceplugin

import (
	"fmt"

	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/operators"
	"github.com/corazawaf/coraza/v2/types/variables"
)

type wace struct {
}

func (o *wace) Init(data string) error {
	return nil
}

func (o *wace) Evaluate(tx *coraza.Transaction, value string) bool {
	//res, capture := isSQLi(value)
	a := tx.GetCollection(variables.FullRequest).Data()
	for key, value := range a {
		fmt.Print(key)
		fmt.Print(":")
		fmt.Println(value)
	}
	fmt.Println(a)
	return false
}

func init() {
	operators.RegisterPlugin("wace", func() coraza.RuleOperator { return new(wace) })
}

var (
	_ coraza.RuleOperator = &wace{}
)
