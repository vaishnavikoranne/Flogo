package string

import (
	
	
	//"strings"
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
	separator:=params[1].(string)
    var paramSlice string 
	
	
	arrV := reflect.ValueOf(items)
 	if arrV.Kind() == reflect.Slice {
  		for _, v := range items.([]interface{}) {
   		paramSlice =paramSlice+ v.(string)+separator
  		}
  		
 	}
 	paramSlice=paramSlice[:len(paramSlice)-1]
			
	//str:=strings.Join(paramSlice, separator)
	return paramSlice, nil


}
