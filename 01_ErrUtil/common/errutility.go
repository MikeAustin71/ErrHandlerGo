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
	SourceFileName 		string
	ParentObjectName	string
	FuncName       		string
	BaseErrorID    		int64
}

// New - returns a new, populated ErrBaseInfo Structure
func (b ErrBaseInfo) New(srcFile, parentObjName, funcName string, baseErrID int64) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: srcFile, ParentObjectName: parentObjName, FuncName: funcName, BaseErrorID: baseErrID}
}

// NewFuncName - Returns a New ErrBaseInfo structure with a different Func Name
func (b ErrBaseInfo) NewFunc(funcName string) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: funcName, BaseErrorID: b.BaseErrorID}
}

// NewOpsMsgContextInfo - returns a deep copy of the current
// ErrBaseInfo structure.
func (b ErrBaseInfo) NewBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: b.FuncName, BaseErrorID: b.BaseErrorID}
}

// DeepCopyOpsMsgContextInfo - Same as NewOpsMsgContextInfo()
func (b ErrBaseInfo) DeepCopyBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, ParentObjectName: b.ParentObjectName, FuncName: b.FuncName, BaseErrorID: b.BaseErrorID}
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
	PrefixMsg          	string
	ErrMsgLabel        	string
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

// CheckIsSpecErrPanic - Returns 'true' if
// SpecErr object is configured as a panic
// error.
func (s *SpecErr) CheckIsSpecErrPanic() bool {

	return s.IsPanic
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

// Error - Implements Error Interface.
// Call this method to produce the error
// message as a string.
func (s SpecErr) Error() string {

	var banner1, banner2, mTitle, m string


	if s.ErrorMsgType == SpecErrTypeNOERRORSALLCLEAR {
		return "No Errors - No Messages"
	} else if	s.ErrorMsgType == SpecErrTypeINFO {

		banner1 = "\n" + strings.Repeat("*", 78)
		banner2 = "\n" + strings.Repeat("=", 78)
		mTitle = "Information Message"
	} else if s.ErrorMsgType == SpecErrTypeERROR ||
		s.ErrorMsgType == SpecErrTypeFATAL {

		banner1 = "\n" + strings.Repeat("!", 78)
		banner2 = "\n" + strings.Repeat("-", 78)
		mTitle = "Error Message"
	} else if s.ErrorMsgType == SpecErrTypeWARNING {
		banner1 = "\n" + strings.Repeat("?", 78)
		banner2 = "\n" + strings.Repeat("_", 78)
		mTitle = "Warning Message"

	} else {
		// Must be successful completion message
		return "Successful Completion!"
	}

	m = "\n" + banner1
	m += "\n" + mTitle
	m += banner2

	if s.ErrNo != 0 {
		m += fmt.Sprintf("\n  ErrNo: %v", s.ErrNo)
	}

	if s.PrefixMsg != "" {
		m += "\n"
		m += s.PrefixMsg
	}

	m+= "\n"

	if s.ErrMsgLabel != "" {
		m+= s.ErrMsgLabel + ": "
	}

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
			if bi.BaseErrorID != 0 {
				m += fmt.Sprintf("  ErrorID: %v", bi.BaseErrorID)
			}
		}

		//m += "\n"
	}

	m += banner2
	m += "\n  Error Time"
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
// prefix string - 	This string will be prefixed and printed before the error
//									message.
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
func (s SpecErr) Initialize(parent []ErrBaseInfo, bi ErrBaseInfo, prefix string, err error, errType SpecErrMsgType, errNo int64) SpecErr {
	return s.InitializeBaseInfo(parent, bi).New(prefix, err, errType, errNo)

}

// New - Creates new SpecErr Type. Uses existing
// Parent and ErrBaseInfo data. The error is based on
// a parameter of type 'error' passed to the method.
//
// Note: If you set errNo == zero, no error number will be displayed in the
// in the error message.
//
// Input Parameters:
// prefix string - 	This string will be prefixed and printed before the error
//									message.
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
func (s SpecErr) New(prefix string, err error, errType SpecErrMsgType, errNo int64) SpecErr {


	x := SpecErr{
		ParentInfo: s.DeepCopyParentInfo(s.ParentInfo),
		BaseInfo:   s.BaseInfo.DeepCopyBaseInfo(),
	}

	x.SetError(prefix, err, errType, errNo)

	return x
}

// NewErrorMsgString - Creates a new error message
// based on an error message string.
//
// Note: If you set errNo == zero, no error number will be displayed in the
// in the error message.
// Input Parameters:
// prefix string - 	This string will be prefixed and printed before the error
//									message.
//
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
func (s SpecErr) NewErrorMsgString(prefix string, errMsg string, errType SpecErrMsgType, errNo int64 ) SpecErr {

	if errType == SpecErrTypeERROR ||
			errType == SpecErrTypeFATAL {

		er := errors.New(errMsg)

		return s.New(prefix, er, errType, errNo)
	}

	se := SpecErr{
		ParentInfo: s.DeepCopyParentInfo(s.ParentInfo),
		BaseInfo:   s.BaseInfo.DeepCopyBaseInfo(),
	}

	se.SetErrorWithMessage(prefix, errMsg, errType, errNo)

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
	se.SetSuccessfulCompletion()
	return se
}


// SetBaseInfo - Sets the SpecErr ErrBaseInfo internal
// structure. This data is used for creating repetitive
// error information.
func (s *SpecErr) SetBaseInfo(bi ErrBaseInfo) {
	s.BaseInfo = bi.NewBaseInfo()
}

// SetError - Sets the error message for the current or host SpecErr object.
func (s *SpecErr) SetError(prefix string, err error, errType SpecErrMsgType, errNo int64) {

	if errType == SpecErrTypeSUCCESSFULCOMPLETION {
		s.SetSuccessfulCompletion()
		return
	}

	s.IsPanic = false

	if errType == SpecErrTypeFATAL {
		s.IsPanic = true
	}

	s.ErrorMsgType = errType
	s.PrefixMsg = prefix

	if errNo != 0 {
		s.ErrNo = errNo + s.BaseInfo.BaseErrorID
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
func (s *SpecErr) SetErrorWithMessage(prefix string, errMsg string, errType SpecErrMsgType, errNo int64 ) {

	if errType == SpecErrTypeERROR || errType== SpecErrTypeFATAL {
		err := errors.New(errMsg)
		s.SetError(prefix, err, errType, errNo)
		return
	}

	if errType == SpecErrTypeSUCCESSFULCOMPLETION {
		s.SetSuccessfulCompletion()
		return
	}

	s.ErrorMsgType = errType
	s.IsErr = false
	s.IsPanic = false
	s.ErrMsg = errMsg
	s.PrefixMsg = prefix

	if errNo == 0 {
		s.ErrNo = 0
	} else {
		s.ErrNo = errNo + s.BaseInfo.BaseErrorID
	}

}

// SetSuccessfulCompletion - Sets values for the current
// or host SpecErr object to reflect successful completion
// of the operation.
func (s *SpecErr) SetSuccessfulCompletion() {
	s.IsErr = false
	s.IsPanic = false
	s.ErrorMsgType = SpecErrTypeSUCCESSFULCOMPLETION
	s.ErrMsg = "Successful Completion!"
	s.ErrNo = 0
	s.PrefixMsg = ""
	s.SetTime("Local")
}

// SetErrorLabel - If an Error Message Label is needed
// the Error message, set the value Error Message Label
// here.  This method merely sets the SpecErr string field,
// SpecErr.ErrMsgLabel. Of course this field can also be
// set directly with the use of this method.
//
// If the SpecErr.ErrMsgLabel is set to "StdOut Err", the
// error message will be formatted as :
// 						"StdOut Err: Your Error Message"
func (s *SpecErr) SetErrorMessageLabel(errorMsgLabel string) {
	s.ErrMsgLabel = errorMsgLabel
}
// SetParentInfo - Sets the ParentInfo Slice for
// the current SpecErr structure
func (s *SpecErr) SetParentInfo(parent []ErrBaseInfo) {
	if len(parent) == 0 {
		return
	}

	s.ParentInfo = s.DeepCopyParentInfo(parent)
}

// SetTime - Sets the time stamp for this Error
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

// String - Returns the string message
// compiled by the Error() method.
func (s SpecErr) String() string {
	return s.Error()
}

var blankErrBaseInfo = ErrBaseInfo{}
var blankParentInfo = make([]ErrBaseInfo, 0, 10)



