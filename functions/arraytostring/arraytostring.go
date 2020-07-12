package string

import (
	
	"fmt"
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
	
	arrV:=reflect.ValueOf(items)
	fmt.Println(arrV.Kind()) 
	var len=arrV.Len()
	
	var strArray = make([]string, len)
	
	copy(strArray[:], arrV)
	return strings.Join(strArray, seperator), nil


}
