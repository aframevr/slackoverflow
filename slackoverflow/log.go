package slackoverflow

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
	"github.com/logrusorgru/aurora"
)

// Emergency message
func Emergency(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" emergency %s ", aurora.Red("\u2718").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
	stackoverflow.Close(1)
}

// Alert message
func Alert(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" alert     %s ", aurora.Red("\u2718").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
}

// Critical message
func Critical(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" critical  %s ", aurora.Red("\u2718").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
}

// Error message
func Error(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" error     %s ", aurora.Red("\u2717").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
}

// Warning message
func Warning(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" warning   %s ", aurora.Brown("\u26A0").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
}

// Ok message
func Ok(format string, a ...interface{}) {
	prefix := fmt.Sprintf(" ok        %s ", aurora.Green("\u2714").Bold())
	msg := std.GetWithPrefix(prefix, format, a...)
	fmt.Println(msg)
}

// Notice message
func Notice(format string, a ...interface{}) {
	if stackoverflow.logLevel < 300 {
		prefix := fmt.Sprintf(" notice    %s ", aurora.Cyan("\u26A0").Bold())
		msg := std.GetWithPrefix(prefix, format, a...)
		fmt.Println(msg)
	}
}

// Info message
func Info(format string, a ...interface{}) {
	if stackoverflow.logLevel < 250 {
		prefix := fmt.Sprintf(" info      %s ", aurora.Blue("!").Bold())
		msg := std.GetWithPrefix(prefix, format, a...)
		fmt.Println(msg)
	}
}

// Debug message
func Debug(format string, a ...interface{}) {
	if stackoverflow.logLevel == 100 {
		prefix := fmt.Sprintf(" debug     %s ", aurora.Gray("\u2699").Bold())
		msg := std.GetWithPrefix(prefix, format, a...)
		fmt.Println(msg)
	}
}

// DebugArray message
func DebugArray(dbg map[string]interface{}) {
	if stackoverflow.logLevel == 100 {
		fmt.Println(fmt.Sprintf(" debug     %s: %+v", aurora.Gray("\u2699").Bold(), dbg))
	}
}
