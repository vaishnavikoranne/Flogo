package customdatetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)


type fnDateTimeToTimeStamp struct {
}

func init() {
	function.Register(&fnDateTimeToTimeStamp{})
}

func (s *fnDateTimeToTimeStamp) Name() string {
	return "totimestamp"
}

func (s *fnDateTimeToTimeStamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (s *fnDateTimeToTimeStamp) Eval(in ...interface{}) (interface{}, error) {
	plainText := params[0].(string)
	return plainText,nil
}