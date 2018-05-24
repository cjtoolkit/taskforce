/*
TaskForce, a simple, functional task runner without the plugin nonsense!

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
*/
package taskforce

import (
	"fmt"
	"os"
	"os/exec"
)

type TaskForce struct {
	tasks      map[string]func()
	registerFn func(name string, task func())
	runFn      func(names ...string)
	util       utilI
}

// Create new instance of TaskForce
func InitTaskForce() *TaskForce {
	tf := &TaskForce{
		tasks: map[string]func(){},
		util:  util{},
	}
	tf.registerFn = func(name string, task func()) {
		tf.register(name, task)
	}
	tf.runFn = func(names ...string) {
		tf.firstRun(names...)
	}

	return tf
}

/*
Register Task to TaskForce, has no effect after first run.

Not concurrent safe.
*/
func (tf *TaskForce) Register(name string, task func()) {
	tf.registerFn(name, task)
}

/*
Run a selected task.

Non concurrent safe on first run, but is concurrent safe after first run

Note: will recover from error, on first run.
*/
func (tf *TaskForce) Run(names ...string) {
	tf.runFn(names...)
}

// Execute Terminal Command, will panic if there is error with the command.
func (tf *TaskForce) ExecCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	tf.CheckError(cmd.Run())
}

// Check Error, will panic if there is an error.
func (tf *TaskForce) CheckError(err error) {
	if nil != err {
		tf.util.DoPanic(err)
	}
}

func (tf *TaskForce) firstRun(names ...string) {
	tf.runFn = func(names ...string) {
		tf.nextRun(names...)
	}
	tf.registerFn = func(name string, task func()) {}
	defer tf.recover()
	tf.run(names...)
}

func (tf *TaskForce) nextRun(names ...string) {
	tf.run(names...)
}

func (tf *TaskForce) register(name string, task func()) {
	tf.tasks[name] = task
}

func (tf *TaskForce) recover() {
	if r := recover(); nil != r {
		tf.util.DoRecover(r)
	} else {
		tf.util.DisplaySuccess()
	}
}

func (tf *TaskForce) run(names ...string) {
	for _, name := range names {
		fmt.Println("Running Task:", name)
		tf.tasks[name]()
		fmt.Println()
	}
}
