package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"

	AMUtilErrors "../../../internal/apimachinery/pkg/util/errors" //TODO: create package import
)

const (
	DefaultErrorExitCode	= 1
	ErrExitText 			= "exit"
)

// ErrExit may be passed to CheckError to instruct it to output nothing but exit with
// status code 1.
var ErrExit = fmt.Errorf(ErrExitText)

// The fatal error handler
var fatalErrHandler = fatal


// checkErr formats a given error as a string and calls the passed handleErr
// func with that string and an msgraph exit code.
func checkErr(err error, handleErr func(string, int)) {
	// unwrap aggregates of 1
	if agg, ok := err.(AMUtilErrors.Aggregate); ok && len(agg.Errors()) == 1 {
		err = agg.Errors()[0]
	}

	if err == nil {
		return
	}

	//TODO: implement better/more error handling cases
	switch {
	case err == ErrExit:
		handleErr("", DefaultErrorExitCode)
	default:
		//TODO: implement better/more error handling cases
		//Example: https://github.com/kubernetes/kubernetes/blob/f077d6736b968b320a4fcc9fe57a342eae86acc5/pkg/kubectl/cmd/util/helpers.go
		handleErr("Unhandled Error", DefaultErrorExitCode)
	}

}

// fatal prints the message (if provided) and then exits. If V(2) or greater,
// glog.Fatal is invoked for extended information.
func fatal(msg string, code int) {
	/*if glog.V(2) {//TODO: Setup logrus structured logging
		glog.FatalDepth(2, msg)
	}*/
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code. Unrecognized errors will be printed with an "error: " prefix.
//
// This method is generic to the command in use and may be used by msgraph
// commands.
func CheckErr(err error) {
	checkErr(err, fatalErrHandler)
}

// DefaultSubCommandRun prints a command's help string to the specified output if no
// arguments (sub-commands) are provided, or a usage error otherwise.
func DefaultSubCommandRun(out io.Writer) func(c *cobra.Command, args []string) {
	return func(c *cobra.Command, args []string) {
		c.SetOutput(out)
		RequireNoArguments(c, args)
		c.Help()
		CheckErr(ErrExit)
	}
}

// RequireNoArguments exits with a usage error if extra arguments are provided.
func RequireNoArguments(c *cobra.Command, args []string) {
	if len(args) > 0 {
		CheckErr(UsageErrorf(c, "unknown command %q", strings.Join(args, " ")))
	}
}

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '%s -h' for help and examples.", msg, cmd.CommandPath())
}