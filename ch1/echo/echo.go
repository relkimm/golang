package echo

import (
	"fmt"
	"os"
)

func Execute() {
	result := ""
	seperator := " "

	for i := 1; i < len(os.Args); i++ {
		result = result + seperator + os.Args[i]
	}
	fmt.Println(result)
}
