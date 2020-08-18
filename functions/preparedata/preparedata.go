package string

import (
	
	"fmt"
	"sort"
	//"strings"
       //"reflect"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&preparedata{})
}

type preparedata struct {
}

func (s *preparedata) Name() string {
	return "preparedata"
}

func (preparedata) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (preparedata) Eval(params ...interface{}) (interface{}, error) {
	items:=params[0]
	md:= items.(map[string]interface{})
	str:=""
	i:=0
	keys := make([]string, len(md))
	for key, _ := range md{
		keys[i]=key
		i++
        }
	sort.Strings(keys) 
	for _, k := range keys {
       	  str+=k+"="+fmt.Sprint(md[k])+"&"
        }
	fmt.Println("%s", str)
	
	/*mapString := make(map[string]string)
	for key, value := range items {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapString[strKey] = strValue
	}
	fmt.Printf("%T\n", items)
	fmt.Println("%v\n", items)
	fmt.Println("%#v", items)
	fmt.Printf("%#v", mapString)
	*/
	/*keys := make([]string, len(items))
	i := 0
	for k := range items {
	  keys[i] = k
	  fmt.Printf("%s\n", k)	
	  i++
	}*/
	/*var paramSlice []string 
	arrV := reflect.ValueOf(items)
	for i:=0;i<arrV.Len();i++ {
		inter:=arrV.Index(i).Interface()
	
                for _ , v := range inter.(map[string]interface{}){
		   
			paramSlice = append(paramSlice, v.(string) )
		} 
	}
	separator:=params[1].(string)
        var paramSlice []string 
	arrV := reflect.ValueOf(items)
	for i:=0;i<arrV.Len();i++ {
		inter:=arrV.Index(i).Interface()
	
                for _ , v := range inter.(map[string]interface{}){
		   
			paramSlice = append(paramSlice, v.(string) )
		} 
	}
			
	str:=strings.Join(paramSlice, separator)*/
	return str, nil


}

