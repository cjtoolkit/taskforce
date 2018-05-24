//go:generate mockgen -write_package_comment=false -package=taskforce -source=util.go -destination=util.mock.go
//go:generate debugflag util.mock.go

package taskforce

import (
	"fmt"
	"os"
	"runtime/debug"
)

type utilI interface {
	DoPanic(err error)
	DoRecover(r interface{})
	DisplaySuccess()
}

type util struct{}

func (u util) DoPanic(err error) {
	panic(err)
}

func (u util) DoRecover(r interface{}) {
	fmt.Fprintln(os.Stderr, "Has Errored", r)
	fmt.Fprintln(os.Stderr, string(debug.Stack()))
	fmt.Println("-- Failed --")
	os.Exit(1)
}

func (u util) DisplaySuccess() {
	fmt.Println("-- Success --")
}
