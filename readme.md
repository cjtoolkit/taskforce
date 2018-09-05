[![Build Status](https://travis-ci.com/cjtoolkit/taskforce.svg?branch=master)](https://travis-ci.com/cjtoolkit/taskforce)

# TaskForce

a simple, functional task runner without the plugin nonsense!

``` go
package main

import (
	"fmt"
	"os"

	"github.com/cjtoolkit/taskforce"
)

func task() *taskforce.TaskForce {
	tf := taskforce.InitTaskForce()

	tf.Register("hello", func() {
		fmt.Println("Hello,")
	})

	tf.Register("world", func() {
		fmt.Println("World.")
	})

	tf.Register("echo-world", func() {
		tf.ExecCmd("echo", "world")
	})

	tf.Register("both", func() {
		tf.Run("hello", "world")
	})

	return tf
}

func main() {
	task().Run(os.Args[1:]...)
}
```