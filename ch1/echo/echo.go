package echo

import (
	"fmt"
	"os"
	"strings"
)

func Execute() {
	result := strings.Join(os.Args[1:], " ")
	fmt.Println(result)
}
