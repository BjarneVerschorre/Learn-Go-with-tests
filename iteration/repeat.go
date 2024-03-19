package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	builder := strings.Builder{}
	for i := 0; i < repeatCount; i++ {
		builder.WriteString(character)
	}
	return builder.String()
}