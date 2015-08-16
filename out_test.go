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

func TestprintoutMarker(t *testing.T) {
	printout(os.Stdout, "hello, world", "green")
	// Output:  >>> hello, world
}

func TestprintoutArgs(t *testing.T) {
	printout(os.Stdout, "hello, %s", "red", "world")
	// Output:  >>> hello, world
}

func TestprintoutColorsSubstitution(t *testing.T) {
	printout(os.Stdout, "hello, [white]world", "green")
	// Output:  >>> hello, world
}

func runCmd(t *testing.T, args ...string) (string, string) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("go", append([]string{"run", "samples/main.go"}, args...)...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		t.Error("errored: ", err)
	}
	return out.String(), stderr.String()
}

func colored(message, color string) string {
	return "\033[" + colorstring.DefaultColors[color] + "m" + message
}

func TestAll(t *testing.T) {
	message := "hello, world"

	stdout, stderr := runCmd(t, "Success", message)
	assert.Contains(t, stdout, colored(message, "green"))
	assert.Equal(t, "", stderr)

	stdout, stderr = runCmd(t, "Error", message)
	assert.Equal(t, "", stdout)
	assert.Contains(t, stderr, colored(message, "red"))

	stdout, stderr = runCmd(t, "Info", message)
	assert.Contains(t, stdout, colored(message, "blue"))
	assert.Equal(t, "", stderr)

	stdout, stderr = runCmd(t, "Warn", message)
	assert.Equal(t, "", stdout)
	assert.Contains(t, stderr, colored(message, "yellow"))

	os.Setenv("DEBUG", "truthy")
	stdout, stderr = runCmd(t, "Debug", message)
	assert.Equal(t, "", stdout)
	assert.Contains(t, stderr, colored(message, "cyan"))

	// without DEBUG
	os.Setenv("DEBUG", "")
	stdout, stderr = runCmd(t, "Debug", message)
	assert.Equal(t, "", stdout)
	assert.Equal(t, "", stderr)
}
