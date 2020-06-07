package helper

import (
	"fmt"
)

//ComposeSource compose source command to output
func ComposeSource(file string) string {
	return "source " + ConvertToUnixPath(file)

}

//ComposeFPath compose fpath command to output
func ComposeFPath(path string) string {
	return fmt.Sprintf("fpath+=( %s )", ConvertToUnixPath(path))
}

//ComposeEnvPath compose env path command to output
func ComposeEnvPath(path string) string {
	return "export PATH=\"" + ConvertToUnixPath(path) + ":$PATH\""
}