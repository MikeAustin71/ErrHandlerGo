package common

import (
	"fmt"
	"errors"
	"time"
	"strings"
)

/*
		The source code for errutility.go is located in source
		code repository:

						https://github.com/MikeAustin71/ErrHandlerGo.git
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
	ErrMsg             	string
	ErrNo              	int64
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

	if 	s.ErrorMsgTimeUTC 		!= s2.ErrorMsgTimeUTC ||
			s.ErrorMsgTimeLocal 	!= s2.ErrorMsgTimeLocal ||
			s.ErrorLocalTimeZone 	!= s2.ErrorLocalTimeZone ||
			s.ErrorMsgType 				!= s2.ErrorMsgType ||
			s.IsErr								!= s2.IsErr ||
			s.IsPanic							!= s2.IsPanic ||
			s.ErrMsg 							!= s2.ErrMsg ||
			s.ErrNo 							!= s2.ErrNo {

				return false
	}

	return true
}

// Empty - Sets all data fields in the current or host SpecErr
// object to an uninitialized or 'empty' state.
func (s *SpecErr) Empty() {

	s.ParentInfo  = make([]ErrBaseInfo, 0, 20)
	s.BaseInfo 						= ErrBaseInfo{}
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
	s.ErrNo              	= int64(0)
}


// Error - Implements Error Interface.
// Call this method to produce the error
// message as a string.
func (s SpecErr) Error() string {


	var banner1, banner2, mTitle, noTitle, m string


	if s.ErrorMsgType == SpecErrTypeNOERRORSALLCLEAR {
		return "No Errors - No Messages"

	} else if	s.ErrorMsgType == SpecErrTypeINFO {

		banner1 = "\n" + strings.Repeat("*", 78)
		banner2 = "\n" + strings.Repeat("=", 78)
		mTitle = "Information Message"
		noTitle = "InfoMsgNo"
	} else if s.ErrorMsgType == SpecErrTypeERROR ||
		s.ErrorMsgType == SpecErrTypeFATAL {

		banner1 = "\n" + strings.Repeat("!", 78)
		banner2 = "\n" + strings.Repeat("-", 78)
		mTitle = "Error Message"
		noTitle = "ErrNo"
	} else if s.ErrorMsgType == SpecErrTypeWARNING {
		banner1 = "\n" + strings.Repeat("?", 78)
		banner2 = "\n" + strings.Repeat("_", 78)
		mTitle = "Warning Message"
		noTitle = "WarningMsgNo"

	} else {
		// Must be successful completion message

		if s.ErrNo != 0 {
			m = "\nMsgNo: " + string(s.ErrNo) + " "
		} else {
			m = "\n"
		}

		m +=  "Successful Completion!"

		return m

	}

	m = "\n" + banner1
	m += "\n" + mTitle
	m += banner2

	if s.ErrNo != 0 {
		m += fmt.Sprintf("\n  %v: %v", noTitle, s.ErrNo)
	}


	m+= "\n"


	m += s.ErrMsg
	m += banner2

	if s.BaseInfo.SourceFileName != "" ||
		s.BaseInfo.ParentObjectName != "" ||
			s.BaseInfo.FuncName != ""  {
				m+= "\nCurrent Base Context Info"
				m+=  banner2

	}

	if s.BaseInfo.SourceFileName != "" {
		m += "\n  SrcFile: " + s.BaseInfo.SourceFileName
	}

	if s.BaseInfo.ParentObjectName != "" {
		m += "\nParentObj: " + s.BaseInfo.ParentObjectName
	}

	if s.BaseInfo.FuncName != "" {
		m += "\n FuncName: " + s.BaseInfo.FuncName
	}


	m += fmt.Sprintf("\n  IsErr: %v", s.IsErr)
	m += fmt.Sprintf("\nIsPanic: %v", s.IsPanic)

	// If parent Function Info Exists
	// Print it out.
	if len(s.ParentInfo) > 0 {
		m += banner2
		m += "\n  Parent Context Info"
		m += banner2

		for _, bi := range s.ParentInfo {
			m += "\n SrcFile: " + bi.SourceFileName +"  ParentObj: " + bi.ParentObjectName + "  FuncName: " + bi.FuncName
			if bi.BaseErrorId != 0 {
				m += fmt.Sprintf("  ErrorID: %v", bi.BaseErrorId)
			}
		}
	}

	m += banner2
	m += "\n  Time Stamp"
	m += banner2
	dt := DateTimeUtility{}
	dtfmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
	m += fmt.Sprintf("\n  Error Time UTC: %v", dt.GetDateTimeCustomFmt(s.ErrorMsgTimeUTC, dtfmt))
	m += fmt.Sprintf("\nError Time Local: %v", dt.GetDateTimeCustomFmt(s.ErrorMsgTimeLocal, dtfmt))
	localTz := s.ErrorLocalTimeZone

	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}

	m += fmt.Sprintf("\nLocal Time Zone : %v", localTz)
	m +=  banner1
	m += "\n"

	return m
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
		se.SetSuccessfulCompletion(errNo)

	case SpecErrTypeNOERRORSALLCLEAR:
		se.EmptyMsgData()
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
		se.SetSuccessfulCompletion(errNo)

	case SpecErrTypeNOERRORSALLCLEAR:
		se.EmptyMsgData()
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
func (s SpecErr) SignalNoErrors() SpecErr {
	//return SpecErr{IsErr: false, IsPanic: false}
	se := SpecErr{}
	se.SetSuccessfulCompletion(0)
	return se
}


// SetBaseInfo - Sets the SpecErr ErrBaseInfo internal
// structure. This data is used for creating repetitive
// error information.
func (s *SpecErr) SetBaseInfo(bi ErrBaseInfo) {
	s.BaseInfo = bi.NewBaseInfo()
}

// SetError - Sets the error message for the current or host SpecErr object.
func (s *SpecErr) SetError(err error, errType SpecErrMsgType, errNo int64) {

	if errType == SpecErrTypeSUCCESSFULCOMPLETION {
		s.SetSuccessfulCompletion(errNo)
		return
	}

	s.IsPanic = false

	if errType == SpecErrTypeFATAL {
		s.IsPanic = true
	}

	s.ErrorMsgType = errType


	if errNo != 0 {
		s.ErrNo = errNo + s.BaseInfo.BaseErrorId
	} else {
		s.ErrNo = 0
	}

	if err != nil {

		s.ErrMsg = err.Error()

		if errType == SpecErrTypeFATAL ||
			errType == SpecErrTypeERROR {
			s.IsErr = true
		} else {
			s.IsErr = false
		}

	} else {
		s.ErrMsg = ""
		s.IsErr = false
		s.IsPanic = false
	}

	s.SetTime("Local")
}

// SetErrorWithMessage - Configures the current or host SpecErr object according to
// input parameters and an error message string.
func (s *SpecErr) SetErrorWithMessage(errMsg string, errType SpecErrMsgType, errNo int64 ) {

	s.EmptyMsgData()

	if errType == SpecErrTypeERROR || errType== SpecErrTypeFATAL {
		err := errors.New(errMsg)
		s.SetError(err, errType, errNo)
		return
	}

	if errType == SpecErrTypeSUCCESSFULCOMPLETION {
		s.SetSuccessfulCompletion(errNo)
		return
	}

	s.ErrorMsgType = errType
	s.IsErr = false
	s.IsPanic = false
	s.ErrMsg = errMsg


	if errNo == 0 {
		s.ErrNo = 0
	} else {
		s.ErrNo = errNo + s.BaseInfo.BaseErrorId
	}

	s.SetTime("Local")
}

// SetFatalError - Sets the value of the current or host SpecErr object
// to a FATAL error.  Both IsPanic IsErr are set to 'true'.
func (s *SpecErr) SetFatalError(errMsg string, errNo int64) {

	s.EmptyMsgData()
	newErr := errors.New(errMsg)
	s.SetError(newErr, SpecErrTypeFATAL, errNo)

}

// SetInfoMessage - Sets the value of the current or host SpecErr object
// to an 'Information' message.  IsPanic and IsErr are both set to 'false'.
func (s *SpecErr) SetInfoMessage(warningMsg string, msgNo int64) {
	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeINFO
	s.IsErr  = false
	s.IsPanic = false
	s.ErrMsg = warningMsg

	if msgNo == 0 {
		s.ErrNo = 0
	} else {
		s.ErrNo = msgNo + s.BaseInfo.BaseErrorId
	}

	s.SetTime("Local")

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
func (s *SpecErr) SetStdError(errMsg string, errNo int64) {

	s.EmptyMsgData()
	newErr := errors.New(errMsg)
	s.SetError(newErr, SpecErrTypeERROR, errNo)

}

// SetSuccessfulCompletion - Sets values for the current
// or host SpecErr object to reflect successful completion
// of the operation.
func (s *SpecErr) SetSuccessfulCompletion(msgNo int64) {
	s.IsErr = false
	s.IsPanic = false
	s.ErrorMsgType = SpecErrTypeSUCCESSFULCOMPLETION
	s.ErrMsg = "Successful Completion!"

	if msgNo == 0 {
		s.ErrNo = 0
	} else {
		s.ErrNo = msgNo + s.BaseInfo.BaseErrorId
	}

	s.SetTime("Local")
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
func(s *SpecErr)SetTime(localTimeZone string){

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

// SetWarningMessage - Sets the value of the current SpecErr object to a
// 'Warning' Message. Both IsPanic and IsErr are set to 'false'
func (s *SpecErr) SetWarningMessage(warningMsg string, msgNo int64) {
	s.EmptyMsgData()
	s.ErrorMsgType	= SpecErrTypeWARNING
	s.IsErr  = false
	s.IsPanic = false
	s.ErrMsg = warningMsg
	s.ErrNo  = msgNo + s.BaseInfo.BaseErrorId
	s.SetTime("Local")
}

// String - Returns the string message
// compiled by the Error() method.
func (s SpecErr) String() string {
	return s.Error()
}

var blankErrBaseInfo = ErrBaseInfo{}
var blankParentInfo = make([]ErrBaseInfo, 0, 10)



