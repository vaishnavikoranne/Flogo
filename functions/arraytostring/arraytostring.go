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
	items:=reflect.ValueOf(params[0])
	seperator:=reflect.ValueOf(params[1])
	
	return strings.Join(items, seperator), nil

}
