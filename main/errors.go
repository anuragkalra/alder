package main

import "fmt"

func httpStatusError(what string, val int) error {
	return fmt.Errorf("%s %d", what, val)
}
