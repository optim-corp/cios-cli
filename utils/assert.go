package utils

import (
	"fmt"
	"os"

	"github.com/fcfcqloow/go-advance/log"
)

func EAssert(err error) *Assert {
	instance := &Assert{
		Err: err,
	}
	return instance
}
func (self *Assert) ErrNotNil() bool {
	return self.Err != nil
}

func (self *Assert) ErrNil() bool {
	return self.Err == nil
}

func (self *Assert) OnErr(errFn func()) *Assert {
	if self.ErrNotNil() {
		errFn()
	}
	return self
}
func (self *Assert) NoneErr(elseFun func()) *Assert {
	if self.ErrNil() {
		elseFun()
	}
	return self
}
func (self *Assert) Log() *Assert {
	return self.OnErr(func() { log.Error(self.Err.Error()) })
}
func (self *Assert) NoneErrAssert(e error) *Assert {
	return self.NoneErr(func() { self.Err = e })
}
func (self *Assert) NoneErrAssertFn(eFn func() error) *Assert {
	return self.NoneErrAssert(eFn())
}
func (self *Assert) NoneErrPrintln(args ...interface{}) *Assert {
	return self.NoneErr(func() { fmt.Println(args...) })
}

func (self *Assert) NoneErrPrint(args ...interface{}) *Assert {
	return self.NoneErr(func() { fmt.Print(args...) })
}

func (self *Assert) ExitWith(number int) *Assert {
	return self.OnErr(func() { os.Exit(number) })
}

func (self *Assert) Exit() *Assert {
	return self.ExitWith(1)
}

func (self *Assert) Return() error {
	return self.Err
}
