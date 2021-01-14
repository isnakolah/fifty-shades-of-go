package main

import (
	"fmt"
	"gotchas/beginner"
)

func main() {
	var data beginner.Info
	var err error

	data.Result, err = beginner.Work()
	fmt.Println(data, err)
}
