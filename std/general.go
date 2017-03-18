package std

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Hr line
func Hr() {
	fmt.Println(aurora.Magenta(strings.Repeat("=", 120)))
}

// Border for message
func Border() {
	fmt.Println(aurora.Magenta(strings.Repeat("-", 120)))
}

// Nl prints new line
func Nl() {
	fmt.Println(getPrefix())
}

// Set output prefix
func getPrefix() string {
	return aurora.Magenta("# ").String()
}

// GetWithPrefix string with prefix
func GetWithPrefix(prefix string, format string, a ...interface{}) string {
	raw := fmt.Sprintf(format, a...)
	msg := fmt.Sprintf("%s%s", prefix, raw)
	return msg
}

// Msg is line of multiline message
func Msg(format string, a ...interface{}) {
	msg := GetWithPrefix("  ", format, a...)
	fmt.Println(msg)
}

// Body is line within the body surrounded with Hr's
func Body(format string, a ...interface{}) (n int, err error) {
	raw := fmt.Sprintf(format, a...)
	border := getPrefix()
	message := []string{border, " ", raw}
	return fmt.Println(strings.Join(message, ""))
}
