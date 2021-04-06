
package timestamp

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)


type totimestamp struct {
}

func init() {
	function.Register(&totimestamp{})
}

func (s *totimestamp) Name() string {
	return "totimestamp"
}

func (s *totimestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *totimestamp) Eval(in ...interface{}) (interface{}, error) {
	return time.Now().UnixNano() / int64(time.Millisecond),nil
}
