package string

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

var in = &arraytostring{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	var  balance =[]string{{"1000"}, {"2"}, {"3"}, {"17"}, {"50"}}
	
	final, err := in.Eval(balance,",")
	assert.Nil(t, err)
	//assert.Equal(t, int(579), final)
	fmt.Printf("%v\n", final)
}