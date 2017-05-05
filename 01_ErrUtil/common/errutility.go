package common

import (
	"fmt"
)

// ErrBaseInfo is intended for use with
// the SpecErr Structure. It sets up base
// error info to be used repeatedly.
type ErrBaseInfo struct {
	SourceFileName string
	FuncName       string
	BaseErrorID    int64
}

// New - returns a new, populated ErrBaseInfo Structure
func (b ErrBaseInfo) New(srcFile, funcName string, baseErrID int64) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: srcFile, FuncName: funcName, BaseErrorID: baseErrID}
}

// NewFunc - Returns a New ErrBaseInfo structure with a different Func Name
func (b ErrBaseInfo) NewFunc(funcName string) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, FuncName: funcName, BaseErrorID: b.BaseErrorID}
}

// NewBaseInfo - returns a deep copy of the current
// ErrBaseInfo structure.
func (b ErrBaseInfo) NewBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, FuncName: b.FuncName, BaseErrorID: b.BaseErrorID}
}

// GetBaseSpecErr - Returns an empty
// SpecErr structure populated with
// Base Error Info
func (b ErrBaseInfo) GetBaseSpecErr() SpecErr {

	return SpecErr{BaseInfo: b.NewBaseInfo(), IsBaseInfoSet: true}
}

// SpecErr - A data structure used
// to hold custom error information
type SpecErr struct {
	ParentInfo     []ErrBaseInfo
	BaseInfo       ErrBaseInfo
	IsBaseInfoSet  bool
	IsErr          bool
	IsPanic        bool
	PrefixMsg      string
	ErrMsg         string
	SrcFile        string
	ParentFuncName string
	FuncName       string
	ErrNo          int64
}

// New - Creates new SpecErr Type
func (s SpecErr) New(prefix string, err error, isPanic bool, srcFile string, funcName string, errNo int64) SpecErr {

	x := SpecErr{PrefixMsg: prefix, IsPanic: isPanic, SrcFile: srcFile, FuncName: funcName, ErrNo: errNo}

	if s.IsBaseInfoSet {
		x.BaseInfo = s.BaseInfo.NewBaseInfo()
	}

	if srcFile == "" {
		x.SrcFile = s.BaseInfo.SourceFileName
	}

	if funcName == "" {
		x.FuncName = s.BaseInfo.FuncName
	}

	if err != nil {
		x.ErrMsg = err.Error()
		x.IsErr = true
	} else {
		x.ErrMsg = ""
		x.IsErr = false
		x.IsPanic = false
	}

	return x
}

// NewDetailErr - returns a new, populated SpecErr structure.
// It is designed to work with SpecErr.BaseInfo to replicate pre-fabricated
func (s SpecErr) NewDetailErr(prefix string, err error, isPanic bool, errDetailID int64) SpecErr {

	if s.IsBaseInfoSet {
		id := s.BaseInfo.BaseErrorID + errDetailID
		return s.New(prefix, err, isPanic, s.BaseInfo.SourceFileName, s.BaseInfo.FuncName, id)
	}

	return s.New(prefix, err, isPanic, "", "", errDetailID)

}

// SetNoError - Returns a SpecErr
// structure with IsErr set to false.
func (s SpecErr) SetNoError() SpecErr {
	return SpecErr{IsErr: false, IsPanic: false}
}

// SetBaseInfo - Sets the SpecErr ErrBaseInfo internal
// structure. This data is used for creating repetitive
// error information.
func (s SpecErr) SetBaseInfo(bi ErrBaseInfo) {
	s.BaseInfo = bi.NewBaseInfo()
	s.IsBaseInfoSet = true
}

// SetParentInfo - Receives an array of slices
// type ErrBaseInfo and appends deep copies
// of those slices to the SpecErr ParentInfo
// field.
func (s SpecErr) SetParentInfo(pi []ErrBaseInfo) []ErrBaseInfo {
	a := make([]ErrBaseInfo,0, len(pi)+10)
	for _, bi := range pi {
		a = append(a, bi.NewBaseInfo())
	}

	return a
}

// Panic - Executes 'panic' command
// if IsPanic == 'true'
func (s SpecErr) Panic() {
	if s.IsPanic {
		panic(s)
	}
}

// Error - Implements Error Interface
func (s SpecErr) Error() string {
	m := s.PrefixMsg
	m += "\n" + s.ErrMsg

	if s.SrcFile != "" {
		m += "\nSourceFile: " + s.SrcFile
	}

	if s.ParentFuncName != "" {
		m += "\nParentFuncName: " + s.ParentFuncName
	}

	if s.FuncName != "" {
		m += "\nFuncName: " + s.FuncName
	}

	if s.ErrNo != 0 {
		m += fmt.Sprintf("\nErrNo: %v", s.ErrNo)
	}

	m += fmt.Sprintf("\nIsErr: %v", s.IsErr)
	m += fmt.Sprintf("\nIsPanic: %v", s.IsPanic)

	// If parent Function Info Exists
	// Print it out.
	if len(s.ParentInfo) > 0 {
		m += "---------------------"
		m += "  Parent Func Info"
		m += "---------------------"

		for _, bi := range s.ParentInfo {
			m += ("\n" + bi.SourceFileName + "-" + bi.FuncName)
			if bi.BaseErrorID != 0 {
				m += fmt.Sprintf(" ErrorID: %v", bi.BaseErrorID)
			}
		}
	}

	return m
}

// CheckErrPanic - Checks for error and then
// executes 'panic'
func CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckIsSpecErr - If error is present,
// returns 'true'.  If NO Error, returns
// 'false'.
func CheckIsSpecErr(eSpec SpecErr) bool {

	if eSpec.IsErr {
		return true
	}

	return false

}

// PanicOnSpecErr - Issues a 'panic'
// command if SpecErr IsPanic flag is set
func PanicOnSpecErr(eSpec SpecErr) {

	if eSpec.IsPanic {
		panic(eSpec)
	}
}
