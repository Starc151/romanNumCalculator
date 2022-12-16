package quit

import (
	"fmt"
	"os"
)

func Quit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
