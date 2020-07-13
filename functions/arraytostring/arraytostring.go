package string

import (
	
	
	"strings"
        "reflect"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&arraytostring{})
}

type arraytostring struct {
}

func (s *arraytostring) Name() string {
	return "arraytostring"
}

func (arraytostring) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

func (arraytostring) Eval(params ...interface{}) (interface{}, error) {
	items:=params[0]
	seperator:=params[1].(string)
        var paramSlice []string 
	arrV := reflect.ValueOf(items)
	for i:=0;i<arrV.Len();i++ {
		paramSlice = append(paramSlice, arrV.Index(i).(string))	
	}
      
	 
	str:=strings.Join(paramSlice, seperator)
	return str, nil


}
