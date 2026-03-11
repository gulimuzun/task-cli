package errorPrint

import (
	"fmt"
	"os"
)
//ErrPrint prints error to stderr and exits the program with status code 1 
func ErrPrint(err error){
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr,"failed:",err)
	os.Exit(1)
}	
