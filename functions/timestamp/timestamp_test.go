package timestamp

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

var in = &fnDateTimeToTimeStamp{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	final, err := in.Eval()
	assert.Nil(t, err)
	//assert.Equal(t, int(579), final)
	fmt.Printf("%v\n", final)
}