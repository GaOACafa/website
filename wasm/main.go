package main

import(
    "fmt"
    "syscall/js"
	"go/token"
    "go/types"
)

func eval(expr string) (types.TypeAndValue, error) {
    return types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
}

func mainFunc(this js.Value, args []js.Value) interface{} {
	formula := js.Global().Get("document").Call("getElementById", "formula").Get("value").String()
	tv, err := eval(formula)
    ans := ""
	if err != nil {
        fmt.Println(err)
		ans = err.Error()
    } else {
		fmt.Println(tv.Value)
		ans = tv.Value.String()
	}
	js.Global().Get("document").Call("getElementById", "ans").Set("innerHTML", ans)
	return nil
}

func main() {
	ch := make(chan struct{})
	js.Global().Set("calcWasm", js.FuncOf(mainFunc))
	<-ch
}