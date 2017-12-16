package common

import (
	"fmt"
	"time"
	"strings"
	"unicode/utf8"
	"errors"
)

/*
		The source code for errutility.go is located in source
		code repository:

						https://github.com/MikeAustin71/ErrHandlerGo.git

	Dependencies:
	=============

	This file depends on utility routines provided by source code
	file, 'datetimeutility.go'. This source code file is located
	in the source code repository:

						https://github.com/MikeAustin71/datetimeopsgo.git
*/

// SpecErrMsgType - a series of constants used to describe
// Error Message Type
type SpecErrMsgType int

// String - Method used to display the text
// name of an Error Message Type.
func (errmsgtype SpecErrMsgType) String() string {
	return ErrMsgTypeNames[errmsgtype]
}

const (

	// SpecErrTypeNOERRORSALLCLEAR - Describes a state where
	// there are no errors, no warnings and no information
	// messages.
	SpecErrTypeNOERRORSALLCLEAR SpecErrMsgType = iota

	// SpecErrTypeFATAL - Describes a an error which is fatal to
	// program execution. This type of error is equated with
	// 'panic' errors.
	SpecErrTypeFATAL

	// SpecErrTypeERROR - Standard Error
	SpecErrTypeERROR

	// SpecErrTypeWARNING - Not an error. This message types
	// signals a serious warning.
	SpecErrTypeWARNING

	// SpecErrTypeINFO - Not an error. For information purposes
	// only.
	SpecErrTypeINFO

	// SpecErrTypeSUCCESSFULCOMPLETION - Signals that the operation
	// completed successfully with no errors.
	SpecErrTypeSUCCESSFULCOMPLETION

)

// ErrMsgTypeNames - String Array holding Error Message Type names.
var ErrMsgTypeNames = [...]string{"NOERRORSALLCLEAR","FATAL", "ERROR", "WARNING", "INFO","SUCCESS"}

// ErrBaseInfo is intended for use with
// the SpecErr Structure. It sets up base
// error info to be used repeatedly.
type ErrBaseInfo struct {
	SourceFileName   string
	ParentObjectName string
	FuncName         string
	BaseErrorId      int64
}

// New - returns a new, populated ErrBaseInfo Structure
func (b ErrBaseInfo) New(srcFile, parentObjName, funcName string, baseErrID int64) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: srcFile, ParentObjectName: parentObjName, FuncName: funcName, BaseErrorId: baseErrID}
}

// NewFuncName - Returns a New ErrBaseInfo structure with a different Func Name
func (b ErrBaseInfo) NewFunc(funcName string) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: funcName, BaseErrorId: b.BaseErrorId}
}

// NewOpsMsgContextInfo - returns a deep copy of the current
// ErrBaseInfo structure.
func (b ErrBaseInfo) NewBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: b.FuncName, BaseErrorId: b.BaseErrorId}
}

// DeepCopyOpsMsgContextInfo - Same as NewOpsMsgContextInfo()
func (b ErrBaseInfo) DeepCopyBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: b.FuncName, BaseErrorId: b.BaseErrorId}
}

// Equal - Compares two ErrBaseInfo objectes to determine
// if they are equivalent
func (b *ErrBaseInfo) Equal(b2 *ErrBaseInfo) bool {

	if b.SourceFileName != b2.SourceFileName ||
		b.ParentObjectName != b2.ParentObjectName ||
			b.FuncName != b2.FuncName ||
				b.BaseErrorId != b2.BaseErrorId {
					return false
	}

	return true
}

// GetBaseSpecErr - Returns an empty
// SpecErr structure populated with
// Base Error Info
func (b ErrBaseInfo) GetBaseSpecErr() SpecErr {

	return SpecErr{BaseInfo: b.NewBaseInfo()}
}

// GetNewParentInfo - Returns a slice of ErrBaseInfo
// structures with the first element initialized to a
// new ErrBaseInfo structure.
func (b ErrBaseInfo) GetNewParentInfo(srcFile, parentObj, funcName string, baseErrID int64) []ErrBaseInfo {
	var parent []ErrBaseInfo

	bi := b.New(srcFile, parentObj, funcName, baseErrID)

	return append(parent, bi)
}

// SpecErr - A data structure used
// to hold custom error information
type SpecErr struct {
	ParentInfo         	[]ErrBaseInfo
	BaseInfo           	ErrBaseInfo
	ErrorMsgTimeUTC    	time.Time
	ErrorMsgTimeLocal  	time.Time
	ErrorLocalTimeZone 	string
	ErrorMsgType				SpecErrMsgType
	IsErr              	bool
	IsPanic            	bool
	ErrMsg             	string	// Original Error Message passed in by caller
	FmtErrMsg						string	// Formatted Error Message
	ErrId								int64		// Original Error Id Number
	ErrNo              	int64   // Error Id + BaseInfo.BaseErrId
}


// AddParentContextHistory - Adds ParentInfo elements to
// the current SpecErr ParentInfo slice
func (s *SpecErr) AddParentInfo(parent []ErrBaseInfo) {
	if len(parent) == 0 {
		return
	}

	x := s.DeepCopyParentInfo(parent)

	for _, bi := range x {
		s.ParentInfo = append(s.ParentInfo, bi.NewBaseInfo())
	}

	return

}

// AddBaseToParentInfo - Adds the structure's
// ErrBaseInfo data to ParentInfo and returns a
// new ParentInfo Array
func (s *SpecErr) AddBaseToParentInfo() []ErrBaseInfo {

	a := s.DeepCopyParentInfo(s.ParentInfo)
	return append(a, s.BaseInfo.DeepCopyBaseInfo())
}

// CheckErrPanic - Checks for error and then
// executes 'panic'
func (s *SpecErr) CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckIsSpecErr - If error is present,
// returns 'true'.  If NO Error, returns
// 'false'.
func (s *SpecErr) CheckIsSpecErr() bool {

	if s.IsErr {
		return true
	}

	return false

}

// CheckIsSpecErrPanic - Returns 'true' if
// SpecErr object is configured as a panic
// error.
func (s *SpecErr) CheckIsSpecErrPanic() bool {

	return s.IsPanic
}

// ConfigureContext - Receives a 'SpecErr' object and a 'ErrBaseInfo' object
// which are used to populate the current SpecErr fields, 'ParentInfo'
// and 'BaseInfo'.
func (s *SpecErr) ConfigureContext(parentSpecErr SpecErr, newBaseInfo ErrBaseInfo) {
	s.ConfigureParentInfoFromParentSpecErr(parentSpecErr)
	s.ConfigureBaseInfo(newBaseInfo)
}

// ConfigureParentInfoFromParentSpecErr - Receives a SpecErr object
// as input parameter ('parentSpecErr'). The ParentInfo array from
// 'parentSpecErr' object is added to the current or host SpecErr
// object. In addition the 'parentSpecErr' BaseInfo object is also
// added to the current or host ParentInfo array
func (s *SpecErr) ConfigureParentInfoFromParentSpecErr(parentSpecErr SpecErr) {

	s.AddParentInfo(parentSpecErr.ParentInfo)

	s.ParentInfo = append(s.ParentInfo, parentSpecErr.DeepCopyBaseInfo())

}

// ConfigureBaseInfo - Sets 'BaseInfo' for the current or host
// SpecErr object to the input parameter 'newBaseInfo'
func (s *SpecErr) ConfigureBaseInfo(newBaseInfo ErrBaseInfo) {
	s.BaseInfo = newBaseInfo.DeepCopyBaseInfo()
}

// CopyIn - Stores a copy of incoming SpecErr type.
// Receives an incoming SpecErr type (s2),
// creates a deep copy and stores the incoming
// data in the current or host SpecErr object.
func (s *SpecErr) CopyIn(s2 SpecErr) {
	s.Empty()
	s.ParentInfo         	= s2.DeepCopyParentInfo(s2.ParentInfo)
	s.BaseInfo           	= s2.DeepCopyBaseInfo()
	s.ErrorMsgTimeUTC    	= s2.ErrorMsgTimeUTC
	s.ErrorMsgTimeLocal  	= s2.ErrorMsgTimeLocal
	s.ErrorLocalTimeZone 	= s2.ErrorLocalTimeZone
	s.ErrorMsgType				= s2.ErrorMsgType
	s.IsErr              	= s2.IsErr
	s.IsPanic            	= s2.IsPanic
	s.ErrMsg             	= s2.ErrMsg
	s.FmtErrMsg           = s2.FmtErrMsg
	s.ErrId								= s2.ErrId
	s.ErrNo              	= s2.ErrNo
}

// CopyOut - Creates a deep copy of the
// current or host SpecErr object and
// returns the copy.
func (s *SpecErr) CopyOut() SpecErr {

	se := SpecErr{}.InitializeBaseInfo(s.ParentInfo, s.BaseInfo)
	se.ErrorMsgTimeUTC    	= s.ErrorMsgTimeUTC
	se.ErrorMsgTimeLocal  	= s.ErrorMsgTimeLocal
	se.ErrorLocalTimeZone 	= s.ErrorLocalTimeZone
	se.ErrorMsgType				= s.ErrorMsgType
	se.IsErr              	= s.IsErr
	se.IsPanic            	= s.IsPanic
	se.ErrMsg             	= s.ErrMsg
	se.FmtErrMsg            = s.FmtErrMsg
	se.ErrId								= s.ErrId
	se.ErrNo              	= s.ErrNo

	return se
}

// DeepCopyOpsMsgContextInfo - Returns a deep copy of the
// current BaseInfo structure.
func (s *SpecErr) DeepCopyBaseInfo() ErrBaseInfo {
	return s.BaseInfo.DeepCopyBaseInfo()
}

// DeepCopyParentContextHistory - Receives an array of slices
// type ErrBaseInfo and appends deep copies
// of those slices to the SpecErr ParentInfo
// field.
func (s *SpecErr) DeepCopyParentInfo(pi []ErrBaseInfo) []ErrBaseInfo {

	if len(pi) == 0 {
		return pi
	}

	a := make([]ErrBaseInfo, 0, len(pi)+10)
	for _, bi := range pi {
		a = append(a, bi.NewBaseInfo())
	}

	return a
}

// Equal - Compares two SpecErr objects to
// determine if they are equivalent.
func (s *SpecErr) Equal( s2 *SpecErr) bool {

	if len(s.ParentInfo) != len(s2.ParentInfo) {
		return false
	}

	for i:= 0; i < len(s.ParentInfo); i++ {
		if !s.ParentInfo[i].Equal(&s2.ParentInfo[i]) {
			return false
		}
	}

	if !s.BaseInfo.Equal(&s2.BaseInfo) {
		return false
	}

	if 	s.ErrorMsgTimeUTC 		!= s2.ErrorMsgTimeUTC 		||
			s.ErrorMsgTimeLocal 	!= s2.ErrorMsgTimeLocal 	||
			s.ErrorLocalTimeZone 	!= s2.ErrorLocalTimeZone 	||
			s.ErrorMsgType 				!= s2.ErrorMsgType       	||
			s.IsErr								!= s2.IsErr              	||
			s.IsPanic							!= s2.IsPanic            	||
			s.ErrMsg 							!= s2.ErrMsg             	||
			s.FmtErrMsg           != s2.FmtErrMsg 					||
			s.ErrId								!= s2.ErrId   						||
			s.ErrNo 							!= s2.ErrNo {

			return false
	}

	return true
}

// Empty - Sets all data fields in the current or host SpecErr
// object to an uninitialized or 'empty' state.
func (s *SpecErr) Empty() {

	s.ParentInfo  = make([]ErrBaseInfo, 0, 20)
	s.BaseInfo 		= ErrBaseInfo{}
	s.EmptyMsgData()
}

// EmptyMsgData - Sets all data fields except 'ParentInfo'
// and 'BaseInfo' to an uninitialized or 'empty' state.
func (s *SpecErr) EmptyMsgData() {
	s.ErrorMsgTimeUTC 		= time.Time{}
	s.ErrorMsgTimeLocal 	= time.Time{}
	s.ErrorLocalTimeZone 	= ""
	s.ErrorMsgType				= SpecErrTypeNOERRORSALLCLEAR
	s.IsErr              	= false
	s.IsPanic            	= false
	s.ErrMsg             	= ""
	s.FmtErrMsg						= ""
	s.ErrId								= int64(0)
	s.ErrNo              	= int64(0)
}


// Error - Implements Error Interface.
// Call this method to produce the error
// message as a string.
func (s SpecErr) Error() string {

	return s.FmtErrMsg

}

// InitializeBaseInfoWithSpecErr - Initialize a SpecErr object by passing in a parent SpecErr object and
// the current BaseInfo data.
func (s SpecErr) InitializeBaseInfoWithSpecErr(parentSpeErr SpecErr, currentBaseInfo ErrBaseInfo) SpecErr {

	se := SpecErr{}

	se.ConfigureContext(parentSpeErr, currentBaseInfo)

	return se

}

// InitializeBaseInfo - Initializes a SpecErr Structure
// from a ParentInfo array and a ErrBaseInfo
// structure
func (s SpecErr) InitializeBaseInfo(parent []ErrBaseInfo, currentBaseInfo ErrBaseInfo) SpecErr {

	return SpecErr{
		ParentInfo: s.DeepCopyParentInfo(parent),
		BaseInfo:   currentBaseInfo.DeepCopyBaseInfo()}
}

// InitializeCurrentBaseInfo - Initialize a SpecErr Structure
// wherein only the current BaseInfo object is initialized. The
// ParentInfo remains empty or uninitialized.
func (s SpecErr) InitializeCurrentBaseInfo(currentBaseInfo ErrBaseInfo) SpecErr {

	return SpecErr{
		BaseInfo:   currentBaseInfo.DeepCopyBaseInfo()}

}

// Initialize - Initializes all elements of
// the SpecErr structure
//
// Input Parameters:
// parent [] ErrBaseInfo - 	This represents history data of the function chain
//													which preceded the function in which this error occurred.
//
// bi ErrBaseInfo 			 -	This represents the base information associated with the
//													current function in which the error occurred.
//
// err error		 - 	Type Error containing the error message which will be associated
//									with this SpecErr object.
//
// errType SpecErrMsgType	 -	A constant designating the type
//														of error message to be created.
//
// errNo	int64  - 	An int64 value which specifies the error number associated with this
//									error message. If 'errNo' is set to zero - no error number will be
//									will be displayed in the final error message.
//
func (s SpecErr) Initialize(parent []ErrBaseInfo, bi ErrBaseInfo, err error, errType SpecErrMsgType, errNo int64) SpecErr {
	return s.InitializeBaseInfo(parent, bi).New(err, errType, errNo)

}



// New - Creates new SpecErr Type. Uses existing
// Parent and ErrBaseInfo data. The error is based on
// a parameter of type 'error' passed to the method.
//
// Note: If you set errNo == zero, no error number will be displayed in the
// in the error message.
//
// Input Parameters:
//
// err error		 - 	Type Error containing the error message which will be associated
//									with this SpecErr object.
//
// errType SpecErrMsgType	 -	A constant designating the type
//														of error message to be created.
//
// errNo	int64  - 	An int64 value which specifies the error number associated with this
//									error message. If 'errNo' is set to zero - no error number will be
//									will be displayed in the final error message.
//
func (s SpecErr) New(err error, errType SpecErrMsgType, errNo int64) SpecErr {


	se := SpecErr{
		ParentInfo: s.DeepCopyParentInfo(s.ParentInfo),
		BaseInfo:   s.BaseInfo.DeepCopyBaseInfo(),
	}

	errMsg := err.Error()

	switch errType {
	case SpecErrTypeERROR:
		se.SetStdError(errMsg, errNo)

	case SpecErrTypeFATAL:
		se.SetFatalError(errMsg, errNo)

	case SpecErrTypeINFO:
		se.SetInfoMessage(errMsg, errNo)

	case SpecErrTypeWARNING:
		se.SetWarningMessage(errMsg, errNo)

	case SpecErrTypeSUCCESSFULCOMPLETION:
		se.SetSuccessfulCompletion(errMsg, errNo)

	case SpecErrTypeNOERRORSALLCLEAR:
		se.SetNoErrorsNoMessages(errMsg, errNo)
	}

	return se
}

// NewErrorMsgString - Creates a new error message
// based on an error message string.
//
// Note: If you set errNo == zero, no error number will be displayed in the
// in the error message.
// Input Parameters:
// errMsg string - 	This strings contains the error message which will be associated
//									with this SpecErr object.
//
// errType SpecErrMsgType	 -	A constant designating the type
//														of error message to be created.
//
// errNo	int64  - 	An int64 value which specifies the error number associated with this
//									error message. If 'errNo' is set to zero - no error number will be
//									will be displayed in the final error message.
//
func (s SpecErr) NewErrorMsgString(errMsg string, errType SpecErrMsgType, errNo int64 ) SpecErr {

	se := SpecErr{
		ParentInfo: s.DeepCopyParentInfo(s.ParentInfo),
		BaseInfo:   s.BaseInfo.DeepCopyBaseInfo(),
	}

	switch errType {
	case SpecErrTypeERROR:
		se.SetStdError(errMsg, errNo)

	case SpecErrTypeFATAL:
		se.SetFatalError(errMsg, errNo)

	case SpecErrTypeINFO:
		se.SetInfoMessage(errMsg, errNo)

	case SpecErrTypeWARNING:
		se.SetWarningMessage(errMsg, errNo)

	case SpecErrTypeSUCCESSFULCOMPLETION:
		se.SetSuccessfulCompletion(errMsg, errNo)

	case SpecErrTypeNOERRORSALLCLEAR:
		se.SetNoErrorsNoMessages(errMsg, errNo)
	}

	return se
}

// Panic - Executes 'panic' command
// if IsPanic == 'true'
func (s *SpecErr) Panic() {
	if s.IsPanic {
		panic(s)
	}
}

// PanicOnSpecErr - Issues a 'panic'
// command if SpecErr IsPanic flag is set
func (s *SpecErr) PanicOnSpecErr(eSpec SpecErr) {

	if s.IsPanic {
		panic(s)
	}
}

// SignalNoErrors - Returns a SpecErr
// structure with IsErr set to false.
// Returned error type is SpecErrTypeNOERRORSALLCLEAR
func (s SpecErr) SignalNoErrors() SpecErr {
	//return SpecErr{IsErr: false, IsPanic: false}
	se := SpecErr{}

	se.SetNoErrorsNoMessages("No Errors - No Messages", 0)
	return se
}

// SignalSuccessfulCompletion - Returns a SpecErr structure with
// IsErr set to false. The returned error type is SpecErrTypeSUCCESSFULCOMPLETION
func (s SpecErr) SignalSuccessfulCompletion() SpecErr {

	se := SpecErr{}
	se.SetSuccessfulCompletion("", 0)
	return se
}

// SetBaseInfo - Sets the SpecErr ErrBaseInfo internal
// structure. This data is used for creating repetitive
// error information.
func (s *SpecErr) SetBaseInfo(bi ErrBaseInfo) {
	s.BaseInfo = bi.NewBaseInfo()
}

// SetError - Sets the error message for the current or host SpecErr object.
func (s *SpecErr) SetError(err error, errType SpecErrMsgType, errId int64) {

	if err==nil {
		s.SetNoErrorsNoMessages("", errId)
		return
	}

	switch errType {

	case SpecErrTypeNOERRORSALLCLEAR:
		s.SetNoErrorsNoMessages(err.Error(), errId)

	case SpecErrTypeERROR:
		s.SetStdError(err.Error(), errId)

	case SpecErrTypeFATAL:
		s.SetFatalError(err.Error(), errId)

	case SpecErrTypeINFO:
		s.SetInfoMessage(err.Error(), errId)

	case SpecErrTypeWARNING:
		s.SetWarningMessage(err.Error(), errId)

	case SpecErrTypeSUCCESSFULCOMPLETION:
		s.SetSuccessfulCompletion(err.Error(), errId)

	default:
		panic("SpecErr.SetError() - INVALID SpecErrType!")
	}


	return
}

// SetErrorWithMessage - Configures the current or host SpecErr object according to
// input parameters and an error message string.
func (s *SpecErr) SetErrorWithMessage(errMsg string, errType SpecErrMsgType, errId int64 ) {

	s.EmptyMsgData()

	switch errType {

	case SpecErrTypeNOERRORSALLCLEAR:
		s.SetNoErrorsNoMessages(errMsg, errId)

	case SpecErrTypeERROR:
		s.SetStdError(errMsg, errId)

	case SpecErrTypeFATAL:
		s.SetFatalError(errMsg, errId)

	case SpecErrTypeINFO:
		s.SetInfoMessage(errMsg, errId)

	case SpecErrTypeWARNING:
		s.SetWarningMessage(errMsg, errId)

	case SpecErrTypeSUCCESSFULCOMPLETION:
		s.SetSuccessfulCompletion(errMsg, errId)

	default:
		panic("SpecErr.SetErrorWithMessage() - INVALID SpecErrType!")
	}


}

// SetFatalError - Sets the value of the current or host SpecErr object
// to a FATAL error.  Both IsPanic IsErr are set to 'true'.
func (s *SpecErr) SetFatalError(errMsg string, errNo int64) {

	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeFATAL
	s.IsErr  = true
	s.IsPanic = false
	s.setFormatMessage(errMsg, errNo)


}

// SetInfoMessage - Sets the value of the current or host SpecErr object
// to an 'Information' message.  IsPanic and IsErr are both set to 'false'.
func (s *SpecErr) SetInfoMessage(infoMsg string, msgId int64) {
	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeINFO
	s.IsErr  = false
	s.IsPanic = false
	s.setFormatMessage(infoMsg, msgId)
}

// SetNoErrorsNoMessages - Sets or default
// 'empty' message state.
// SpecErrType= SpecErrTypeNOERRORSALLCLEAR
func (s *SpecErr) SetNoErrorsNoMessages(msg string, msgNo int64) {
	s.EmptyMsgData()
	s.setFormatMessage(msg, msgNo)
}

// SetParentInfo - Sets the ParentInfo Slice for
// the current SpecErr structure
func (s *SpecErr) SetParentInfo(parent []ErrBaseInfo) {
	if len(parent) == 0 {
		return
	}

	s.ParentInfo = s.DeepCopyParentInfo(parent)
}

// SetStdError - Sets the value of the current or host SpecErr object
// to a Standard or non-fatal error. IsPanic is set to 'false' IsErr is
// set to 'true'.
func (s *SpecErr) SetStdError(errMsg string, errId int64) {

	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeERROR
	s.IsErr  = true
	s.IsPanic = false
	s.setFormatMessage(errMsg, errId)
}

// SetSuccessfulCompletion - Sets values for the current
// or host SpecErr object to reflect successful completion
// of the operation.
func (s *SpecErr) SetSuccessfulCompletion(msg string, msgId int64) {
	s.IsErr = false
	s.IsPanic = false
	s.ErrorMsgType = SpecErrTypeSUCCESSFULCOMPLETION

	s.setFormatMessage(msg, msgId)
}

// SetWarningMessage - Sets the value of the current SpecErr object to a
// 'Warning' Message. Both IsPanic and IsErr are set to 'false'
func (s *SpecErr) SetWarningMessage(warningMsg string, msgId int64) {
	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeWARNING
	s.IsErr  = false
	s.IsPanic = false
	s.setFormatMessage(warningMsg, msgId)
}

// String - Returns the string message
// compiled by the Error() method.
func (s SpecErr) String() string {
	return s.Error()
}

/*
*********************************************
				Private Methods
				***************
*********************************************
*/

// setMessageText - Sets the original message and formats
// the message for display
func(s *SpecErr) setFormatMessage(msg string, msgNo int64){

	s.setMsgIdMsgNo(msgNo)

	s.setTime("Local")

	s.ErrMsg = msg

	banner1, banner2, mTitle, numTitle := s.setMsgParms()

	var m string
	dt := DateTimeUtility{}
	dtfmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
	lineWidth := len(banner1)


	// Common Message Formatting
	m = "\n\n"
	m += "\n" + banner1

	s1 := (lineWidth / 3) * 2
	s2 := lineWidth - s1

	if s.ErrNo != 0 {
		sNo:= fmt.Sprintf("%v: %v", numTitle, s.ErrNo)
		str1, _ := s.strCenterInStr(mTitle, s1)
		str2, _ := s.strRightJustify(sNo, s2)
		m+= "\n" + str1 + str2
	} else {
		str1, _ := s.strCenterInStr(mTitle, s1)
		m+= "\n" + str1
	}

	nextBanner := banner1

	if s.ErrorMsgType == SpecErrTypeERROR || s.ErrorMsgType == SpecErrTypeFATAL {

		m += "\n" + nextBanner
		nextBanner = banner2
		str1 := fmt.Sprintf(" IsError: %v     Is Panic/Fatal Error: %v", s.IsErr, s.IsPanic)
		m += "\n" + str1

	}

	if s.ErrMsg != "" {
		m += "\n" + nextBanner
		nextBanner = banner2

		m += "\n Message: "

		if len(s.ErrMsg) > 67 {
			m += "\n  "
		}

		m += s.ErrMsg
	} else {
		m += "\n" + nextBanner
		nextBanner = banner2
		m += "\n Message: NO MESSAGE TEXT PROVIDED!!"

	}


	// If parent Function Info Exists
	// Print it out.
	if len(s.ParentInfo) > 0 {
		m += "\n" + nextBanner
		nextBanner = banner2

		m += "\n Parent Context Info:"

		for _, bi := range s.ParentInfo {
			m += "\n  SrcFile: " + bi.SourceFileName
			m += " -ParentObj: " + bi.ParentObjectName
			m += " -FuncName: " + bi.FuncName
			m += " -BaseMsgId: " + fmt.Sprintf("%v", bi.BaseErrorId)
			}
	}


	if s.BaseInfo.SourceFileName != "" ||
		s.BaseInfo.ParentObjectName != "" ||
		s.BaseInfo.FuncName != "" {
		m += "\n" + nextBanner
		nextBanner = banner2

		m += "\n Current Base Context Info:"

		if s.BaseInfo.SourceFileName != "" {
			m += "\n  SrcFile: " + s.BaseInfo.SourceFileName
		}

		if s.BaseInfo.ParentObjectName != "" {
			m += " -ParentObj: " + s.BaseInfo.ParentObjectName
		}

		if s.BaseInfo.FuncName != "" {
			m += " -FuncName: " + s.BaseInfo.FuncName
		}

		if s.BaseInfo.BaseErrorId != 0 {
			m += " -BaseMsgId: " + fmt.Sprintf("%v", s.BaseInfo.BaseErrorId)
		}


	}

	m += "\n" + nextBanner
	nextBanner = banner2
	m += fmt.Sprintf("\n   UTC Time: %v", dt.GetDateTimeCustomFmt(s.ErrorMsgTimeUTC, dtfmt))
	m += fmt.Sprintf("\n Local Time: %v", dt.GetDateTimeCustomFmt(s.ErrorMsgTimeLocal, dtfmt))
	m += "\n" + banner1

	s.FmtErrMsg = m
	return

}

// setMsgParms - Set Message Parameters.
// Called by SpecErr.setMessageText()
func(s *SpecErr) setMsgParms() (banner1, banner2, title, numTitle string) {

	switch s.ErrorMsgType {

	case SpecErrTypeNOERRORSALLCLEAR:
		title = "No Errors - No Messages"
		banner1 =  strings.Repeat("&", 78)
		banner2 =  strings.Repeat("-", 78)
		numTitle = "MsgNo"

	case SpecErrTypeERROR:
		banner1 =  strings.Repeat("#", 78)
		banner2 =  strings.Repeat("-", 78)
		title = "Standard Error Message"
		numTitle = "ErrNo"

	case SpecErrTypeFATAL:
		banner1 =  strings.Repeat("!", 78)
		banner2 =  strings.Repeat("-", 78)
		title = "Fatal Error Message"
		numTitle = "ErrNo"

	case SpecErrTypeINFO:

		banner1 =  strings.Repeat("*", 78)
		banner2 =  strings.Repeat("-", 78)
		title = "Information Message"
		numTitle = "InfoMsgNo"

	case SpecErrTypeWARNING:
		banner1 =  strings.Repeat("?", 78)
		banner2 =  strings.Repeat("-", 78)
		title = "Warning Message"
		numTitle = "WarningMsgNo"

	case SpecErrTypeSUCCESSFULCOMPLETION:
		banner1 =  strings.Repeat("$", 78)
		banner2 =  strings.Repeat("-", 78)
		title = "Successful Completion"
		numTitle = "MsgNo"

	default:
		panic("SpecErr.setMsgParms() - INVALID SpecErrType")
	}

	return banner1, banner2, title, numTitle
}

func(s *SpecErr) setMsgIdMsgNo(msgId int64){

	if msgId == 0 {
		s.ErrId = 0
		s.ErrNo = 0
	} else {
		s.ErrId = msgId
		s.ErrNo = s.ErrId + s.BaseInfo.BaseErrorId
	}
}

// setTime - Sets the time stamp for this Error
// Message. Notice that the input parameter 'localTimeZone'
// is the Iana Time Zone for local time.
//
// Reference Iana Time Zones: https://www.iana.org/time-zones
//
// If the 'localTimeZone' parameter string is empty or an
// invalid time zone, local time zone will default to 'Local'.
//
// By default the 'localTimeZone' is set to "Local" signaling
// that the local time zone for the host computer will be used.
func(s *SpecErr) setTime(localTimeZone string){

	tz := TimeZoneUtility{}

	isValid, _, _ := tz.IsValidTimeZone(localTimeZone)

	if !isValid {
		localTimeZone = "Local"
	}

	s.ErrorMsgTimeUTC = time.Now().UTC()
	s.ErrorLocalTimeZone = localTimeZone

	tzLocal, _ := tz.ConvertTz(s.ErrorMsgTimeUTC, s.ErrorLocalTimeZone)
	s.ErrorMsgTimeLocal = tzLocal.TimeOut

}

/*

Private String Management Methods

*/

// strCenterInStr - returns a string which includes
// a left pad blank string plus the original string.
// The complete string will effectively center the
// original string is a field of specified length.
func (s *SpecErr) strCenterInStr(strToCenter string, fieldLen int) (string, error) {

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter,  fmt.Errorf("'fieldLen' = '%v' strToCenter Length= '%v'. 'fieldLen is shorter than strToCenter Length!", fieldLen, sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen-sLen)/2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}


	return leftPadStr + strToCenter	+ rightPadStr, nil

}

// strRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
func (s *SpecErr) strRightJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify, fmt.Errorf("Length of string to right justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)



	return leftPadStr + strToJustify, nil
}

// strPadLeftToCenter - Returns a blank string
// which allows centering of the target string
// in a fixed length field.
func (s *SpecErr) strPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

	sLen := s.strGetRuneCnt(strToCenter)

	if sLen > fieldLen {
		return "", errors.New("StringUtility:StrPadLeftToCenter() - String To Center is longer than Field Length")
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// strGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (s *SpecErr) strGetRuneCnt(targetStr string) int {
	return utf8.RuneCountInString(targetStr)
}

// strGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (s *SpecErr) strGetCharCnt(targetStr string) int {
	return len([]rune(targetStr))
}



var blankErrBaseInfo = ErrBaseInfo{}
var blankParentInfo = make([]ErrBaseInfo, 0, 10)



