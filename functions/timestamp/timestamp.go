package timestamp

import (
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
	return []data.Type{}, false
}

func (s *fnDateTimeToTimeStamp) Eval(in ...interface{}) (interface{}, error) {
	return time.Now().UnixNano() / int64(time.Millisecond),nil
}
