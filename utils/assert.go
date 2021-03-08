package utils

import (
	"os"

	log "github.com/optim-kazuhiro-seida/loglog"
)

func EAssert(err error) Assert {
	instance := Assert{
		Err: err,
	}
	return instance
}
func (self Assert) Log() Assert {
	if self.Err != nil {
		log.Error(self.Err.Error())
	}
	return self
}

func (self Assert) OnErr(errFn func()) Assert {
	if self.Err != nil {
		errFn()
	}
	return self
}
func (self Assert) NoneErr(elseFun func()) Assert {
	if self.Err == nil {
		elseFun()
	}
	return self
}
func (self Assert) NoneErrAssert(e error) Assert {
	if self.Err == nil {
		self.Err = e
	}
	return self
}
func (self Assert) NoneErrPrintln(str ...interface{}) Assert {
	if self.Err == nil {
		Println(str...)
	}
	return self
}

func (self Assert) NoneErrPrint(str ...interface{}) Assert {
	if self.Err == nil {
		Print(str...)
	}
	return self
}

func (self Assert) ExitWith(number int) Assert {
	if self.Err != nil {
		os.Exit(number)
	}
	return self
}

func (self Assert) Exit() Assert {
	if self.Err != nil {
		os.Exit(1)
	}
	return self
}

func (self Assert) Return() error {
	return self.Err
}
func (self Assert) ErrNotNil() bool {
	return self.Err != nil
}

func (self Assert) ErrNil() bool {
	return self.Err == nil
}
