package taskforce

import (
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
)

type TaskForce struct {
	tasks      map[string]func()
	registerFn func(name string, task func())
	runFn      func(names ...string)
}

func InitTaskForce() *TaskForce {
	tf := &TaskForce{
		tasks: map[string]func(){},
	}
	tf.registerFn = func(name string, task func()) {
		tf.register(name, task)
	}
	tf.runFn = func(names ...string) {
		tf.firstRun(names...)
	}

	return tf
}

func (tf *TaskForce) Register(name string, task func()) {
	tf.registerFn(name, task)
}

func (tf *TaskForce) Run(names ...string) {
	tf.runFn(names...)
}

func (tf *TaskForce) ExecCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	tf.CheckError(cmd.Run())
}

func (tf *TaskForce) CheckError(err error) {
	if nil != err {
		panic(err)
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
		fmt.Fprintln(os.Stderr, "Has Errored", r)
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		fmt.Println("-- Failed --")
		os.Exit(1)
	} else {
		fmt.Println("-- Success --")
	}
}

func (tf *TaskForce) run(names ...string) {
	for _, name := range names {
		fmt.Println("Running Task:", name)
		tf.tasks[name]()
		fmt.Println()
	}
}
