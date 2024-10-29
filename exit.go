package gotools

import (
	"fmt"
	"os"
)

func FailIfError(err error) {
	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		os.Exit(1)
	}
}
