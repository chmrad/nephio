package main
/*ttt*/
import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	fnr "github.com/nephio-project/nephio/krm-functions/nad-fn/fn"
	"os"
)

r := &fnr.FnR{
	Client: nil,
}


func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(fnr.Run)); err != nil {
		os.Exit(1)
	}
}
