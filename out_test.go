/*
These test suites have been adopted from
https://github.com/mitchellh/colorstring/blob/master/colorstring_test.go
as required.
*/
package out

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/mitchellh/colorstring"
	"github.com/stretchr/testify/assert"
)

func TestPrintMarker(t *testing.T) {
	print("hello, world", "green")
	// Output:  >>> hello, world
}

func TestPrintArgs(t *testing.T) {
	print("hello, %s", "red", "world")
	// Output:  >>> hello, world
}

func TestPrintColorsSubstitution(t *testing.T) {
	print("hello, [white]world", "green")
	// Output:  >>> hello, world
}

func runCmd(t *testing.T, args ...string) string {
	var out bytes.Buffer
	cmd := exec.Command("go", append([]string{"run", "samples/main.go"}, args...)...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Error("errored: ", err)
	}
	return out.String()
}

func colored(message, color string) string {
	return "\033[" + colorstring.DefaultColors[color] + "m" + message
}

func TestAll(t *testing.T) {
	message := "hello, world"

	out := runCmd(t, "Success", message)
	assert.Contains(t, out, colored(message, "green"))

	out = runCmd(t, "Error", message)
	assert.Contains(t, out, colored(message, "red"))

	out = runCmd(t, "Info", message)
	assert.Contains(t, out, colored(message, "blue"))

	out = runCmd(t, "Warn", message)
	assert.Contains(t, out, colored(message, "yellow"))

	os.Setenv("DEBUG", "truthy")
	out = runCmd(t, "Debug", message)
	assert.Contains(t, out, colored(message, "cyan"))

	// without DEBUG
	os.Setenv("DEBUG", "")
	out = runCmd(t, "Debug", message)
	assert.NotContains(t, out, colored(message, "cyan"))
}
