package conversion

import "fmt"

func asciiToByte(r rune) string {
	return fmt.Sprintf("%08b", r)
}

func StringToBytes(s string) []string {
	var bytes []string
	for _, char := range s {
		bytes = append(bytes, (asciiToByte(char)))
	}
	return bytes
}
