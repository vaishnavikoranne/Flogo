package encryptdummy
import(
 "fmt"
 
 "testing"
 
)
var testfn=&AESEncrypt{}
var inputString1,in2,in3 string
func init(){
	
	inputString1="ball"
	in2="47cef24b-2b82-4ac4-a27c-fb0aca32baea"
	in3="ac103458-fcb6-41d3-94r0-43d25b4f4ff4"
}

func TestMyFirstFlogoFunctionFn_Eval_1(t *testing.T){
	outputString, _:=testfn.Eval(inputString1,in2,in3)
	
	fmt.Println("Output ",outputString)

}
