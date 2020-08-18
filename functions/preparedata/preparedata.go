package string

import (
	
	"fmt"
	"sort"
	"strings"
       //"reflect"
	"net/url"
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
		
		if k=="currencyCode"{
			fmt.Println(md[k])
			fmt.Println(url.PathEscape(fmt.Sprint(md[k])))
			str+=k+"="+fmt.Sprint(md[k])+"&"
        	}else{
			str+=k+"="+url.PathEscape(fmt.Sprint(md[k]))+"&"
        	}
		
		fmt.Println(str)
        }
	fstr:=strings.ReplaceAll(str, "%20", "+")
	//fstr=strings.ReplaceAll(fstr, "%2B", "+")
	fstr=strings.TrimSuffix(fstr, "&")
	fmt.Println("%s", fstr)
	
	
	return fstr, nil


}

