/*
Copyright © 2025 NAME HERE ed.crookshanks@gmail.com
*/
package main

import (
	"fmt"
	"runtime"

	"github.com/ecrookshanks/hist-man/cmd"
)

func main() {
	os := runtime.GOOS

	fmt.Println("Operating system:", os)

	cmd.Execute()
}
