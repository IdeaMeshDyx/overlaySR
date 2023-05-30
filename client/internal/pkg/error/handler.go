package srError

import (
	"fmt"
	"time"
)

// handle it after main thread by using defer
func funcErrHandler(originErr error, funcname string) error {
	defer func() {
		// sleep 600 seconds
		for i := 0; i < 600; i++ {
			time.Sleep(time.Second)
		}
	}()

	// combine the current error and funcname
	err := newFuncReturnError(originErr.Error(), funcname)
   fmt.Printf("\033[1;31m %s \033[0m \033[32m Program will exit in 10 mins \033[0m \n", err.Error())
	return err
}
