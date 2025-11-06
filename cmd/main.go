package main

import (
	"IFJudger/cmd/worker"
	"fmt"
)

func main() {
	worker, err := worker.NewWorker()
	if err != nil {
		panic(err)
	}

	worker.SetupPython(64)
	stdout, stderr, err := worker.Execute("var1 = input()\nprint(var1)\nvar2 = input()\nprint(var2)", "teste\ncu", 5)
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
	fmt.Println(stderr)
}
