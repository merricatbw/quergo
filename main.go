package main

import (
	"fmt"
	"strings"

	"github.com/merricatbw/quergo/pkg/conversion"
)

func main() {
	hello := "Hello, world! 123"
	fmt.Println(strings.Join(conversion.StringToBytes(hello), ""))
}
