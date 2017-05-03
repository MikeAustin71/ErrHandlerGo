package common

import (
	"fmt"
)

// SpecError - A data structure used
// to hold custom error information
type SpecError struct {
	IsErr     bool
	IsPanic   bool
	PrefixMsg string
	ErrMsg    string
	FuncName  string
	ErrNo     int64
}

// NewError - Creates new SpecError Type
func (s *SpecError) NewError(prefix string, err error, isPanic bool, funcName string, errNo int64) {

	s.PrefixMsg = prefix
	s.FuncName = funcName
	s.ErrNo = errNo
	s.IsPanic = isPanic
	if err != nil {
		s.ErrMsg = err.Error()
		s.IsErr = true
	} else {
		s.ErrMsg = ""
		s.IsErr = false
	}

}

// Error - Implements Error Interface
func (s *SpecError) Error() string {
	m := s.PrefixMsg
	m += "\n" + s.ErrMsg
	m += "\nFuncName: " + s.FuncName
	m += fmt.Sprintf("\nErrNo: %v", s.ErrNo)
	m += fmt.Sprintf("\nIsErr: %v", s.IsErr)
	m += fmt.Sprintf("\nIsPanic: %v", s.IsPanic)
	return m
}

// CheckErrPanic - Checks for error and then
// executes 'panic'
func CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// SpecCheckErrPanic - Creates a Speck Error
// data type and issues a 'panic' command if
// passed error is valid.
func SpecCheckErrPanic(prefix string, err error) {
	if err == nil {
		return
	}

	e := SpecError{PrefixMsg: prefix, ErrMsg: err.Error()}

	panic(e)
}
