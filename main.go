package main

import (
	"fmt"

	"github.com/yimkh/procon/model"
	"github.com/yimkh/procon/testmethod"
)

func main() {
	env := &model.Environment{}
	env.Mutex = 1
	env.Empty = 5

	fmt.Printf("initial: %v\n", *env)

	testmethod.TestProCon(env)
}
