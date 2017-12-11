package common

import (
	"time"
	"strings"
	"fmt"
	"errors"
)

/*  'opsmsgdto.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/

// OpsMsgType - Designates type of Message being logged
type OpsMsgType int

// String - Method used to display the text
// name of an Operations Message Type.
func (opstype OpsMsgType) String() string {
	return OpsMsgTypeNames[opstype]
}

const (

	// OpsMsgTypeNOERRORNOMSG - 0 Uninitialized -
	// no errors and no messages
	OpsMsgTypeNOERRORNOMSG OpsMsgType = iota

	// OpsMsgTypeERRORMSG - 1 Error Message
	OpsMsgTypeERRORMSG

	// OpsMsgTypeINFOMSG - 2 Information Message Type
	OpsMsgTypeINFOMSG

	// OpsMsgTypeWARNINGMSG - 3 Warning Message Type
	OpsMsgTypeWARNINGMSG

	// OpsMsgTypeDEBUGMSG - 4 Debug Message
	OpsMsgTypeDEBUGMSG

	// OpsMsgTypeSUCCESSFULCOMPLETION - 5 Message signalling
	// successful completion of the operation.
	OpsMsgTypeSUCCESSFULCOMPLETION

)

// OpsMsgTypeNames - String Array holding Message Type names.
var OpsMsgTypeNames = [...]string{"NOERRORSNOMSGS","ERROR", "INFO", "WARNING", "DEBUG", "SUCCESS"}


// OpsMsgClass - Holds the Message level indicating the relative importance of a specific log Message.
type OpsMsgClass int

// String - Returns the name of the OpsMsgClass element
// formatted as a string.
func (opsmsgclass OpsMsgClass) String() string {
	return OpsMsgClassNames[opsmsgclass]
}


const (
	// OpsMsgClassNOERRORSNOMESSAGES - 0 Signals uninitialized message
	// with no errors and no messages
	OpsMsgClassNOERRORSNOMESSAGES OpsMsgClass = iota

	// OpsMsgClassOPERROR - 1 Message is an Error Message
	OpsMsgClassOPERROR

	// OpsMsgClassFATAL - 2 Message is a Fatal Error Message
	OpsMsgClassFATAL

	// OpsMsgClassINFO - 3 Message is an Informational Message
	OpsMsgClassINFO

	// OpsMsgClassWARNING - 4 Message is a warning Message
	OpsMsgClassWARNING

	// OpsMsgClassDEBUG - 5 Message is a Debug Message
	OpsMsgClassDEBUG

	// OpsMsgClassSUCCESSFULCOMPLETION - 6 Message signalling successful
	// completion of the operation
	OpsMsgClassSUCCESSFULCOMPLETION
)

// OpsMsgClassNames - string array containing names of Log Levels
var OpsMsgClassNames = [...]string{"NOERRORSNOMESSAGES", "OPERROR", "FATAL", "INFO", "WARNING", "DEBUG", "SUCCESS"}


// OpsMsgContextInfo - Contains context information describing
// the current environment in which the message was generated.
type OpsMsgContextInfo struct {
	SourceFileName   string
	ParentObjectName string
	FuncName         string
	BaseMessageId    int64
}

// New - returns a new, populated OpsMsgContextInfo Structure
func (ci OpsMsgContextInfo) New(srcFile, parentObject, funcName string, baseMsgID int64) OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: srcFile, ParentObjectName:parentObject, FuncName: funcName, BaseMessageId: baseMsgID}
}

// NewFuncName - Returns a New OpsMsgContextInfo structure with a different Func Name
func (ci OpsMsgContextInfo) NewFuncName(funcName string) OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, FuncName: funcName, BaseMessageId: ci.BaseMessageId}
}

// NewOpsMsgContextInfo - returns a deep copy of the current
// OpsMsgContextInfo structure.
func (ci OpsMsgContextInfo) NewOpsMsgContextInfo() OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, ParentObjectName: ci.ParentObjectName, FuncName: ci.FuncName, BaseMessageId: ci.BaseMessageId}
}

// DeepCopyOpsMsgContextInfo - Same as NewOpsMsgContextInfo()
func (ci OpsMsgContextInfo) DeepCopyOpsMsgContextInfo() OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, ParentObjectName: ci.ParentObjectName, FuncName: ci.FuncName, BaseMessageId: ci.BaseMessageId}
}

// GetBaseOpsMsgDto - Returns an empty
// OpsMsgDto structure populated with
// Base Message Context Information
func (ci OpsMsgContextInfo) GetBaseOpsMsgDto() OpsMsgDto {

	return OpsMsgDto{MsgContext: ci.NewOpsMsgContextInfo()}
}

// GetNewParentInfo - Returns a slice of OpsMsgContextInfo
// structures with the first element initialized to a
// new OpsMsgContextInfo structure.
func (ci OpsMsgContextInfo) GetNewParentInfo(srcFile, parentObject, funcName string, baseMsgID int64) []OpsMsgContextInfo {
	var parent []OpsMsgContextInfo

	newCi := ci.New(srcFile, parentObject, funcName, baseMsgID)

	return append(parent, newCi)
}


// OpsMsgDto - Data Transfer Object
// containing information about an
// operations Message
type OpsMsgDto struct {
	ParentContextHistory [] OpsMsgContextInfo // Function tree showing the execution path leading to this method
	MsgContext           OpsMsgContextInfo
	Message              string // The original message sent to OpsMsgDto
	FmtMessage					 string // The formatted message
	msgId                int64 // The identifying number for this message
	msgNumber            int64 //  Message Number = msgId + MsgContext.BaseMessageId. This is the number displayed in the message
	MsgType              OpsMsgType
	MsgClass             OpsMsgClass
	MsgTimeUTC           time.Time
	MsgTimeLocal         time.Time
	MsgLocalTimeZone     string

}


// AddParentContextHistory - Adds ParentInfo elements to
// the current SpecErr ParentInfo slice
func (opsMsg *OpsMsgDto) AddParentContextHistory(parent []OpsMsgContextInfo) {

	if len(parent) == 0 {
		return
	}

	x := opsMsg.DeepCopyParentContextHistory(parent)

	for _, bi := range x {
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, bi.NewOpsMsgContextInfo())
	}

	return
}

// AddOpsMsgContextInfoToParentHistory - Adds an OpsMsgContextInfo object to the Parent
// Context History maintained by the current or host OpsMsgDto object.
func (opsMsg *OpsMsgDto) AddOpsMsgContextInfoToParentHistory(newContextInfo OpsMsgContextInfo) {
	ci := newContextInfo.DeepCopyOpsMsgContextInfo()
	opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, ci)
}

// ConfigureContextHistoryFromParentOpsMsgDto - Receives an OpsMsgDto object as
// an input parameter ('parentOpsMsgDto'). The parent context history from this
// second OpsMsgDto object ('parentOpsMsgDto') is added to the parent history
// of the current or host OpsMsgDto Object. In addition, the MsgContext object
// from 'parentOpsMsgDto' is also added to the parent history of the current
// or host OpsMsgDto Object.
func (opsMsg *OpsMsgDto) ConfigureContextHistoryFromParentOpsMsgDto(parentOpsMsgDto OpsMsgDto) {
	opsMsg.AddParentContextHistory(parentOpsMsgDto.ParentContextHistory)
	opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, parentOpsMsgDto.DeepCopyMsgContext())

}

// ConfigureMessageContext - Sets MsgContext for the current or host OpsMsgDto object
// to the input parameter 'newMsgContext'
func (opsMsg *OpsMsgDto) ConfigureMessageContext(newMsgContext OpsMsgContextInfo) {
	opsMsg.MsgContext = newMsgContext.DeepCopyOpsMsgContextInfo()
}

// DeepCopyOpsMsgContextInfo - Returns a deep copy of the
// current MsgContext (OpsMsgContextInfo structure).
func (opsMsg *OpsMsgDto) DeepCopyMsgContext() OpsMsgContextInfo {
	return opsMsg.MsgContext.DeepCopyOpsMsgContextInfo()
}


// DeepCopyParentContextHistory - Receives an array of slices
// type OpsMsgContextInfo and appends deep copies
// of those slices to the OpsMsgDto ParentContextHistory
// field.
func (opsMsg *OpsMsgDto) DeepCopyParentContextHistory(pi []OpsMsgContextInfo) []OpsMsgContextInfo {

	if len(pi) == 0 {
		return pi
	}

	newHistory := make([]OpsMsgContextInfo, 0, len(pi)+10)
	for _, ci := range pi {
		newHistory = append(newHistory, ci.NewOpsMsgContextInfo())
	}

	return newHistory
}

// Empty - Resets the current OpsMsgDto object to
// an uninitialized or 'empty' state.
func (opsMsg *OpsMsgDto) Empty() {
	opsMsg.ParentContextHistory = make([] OpsMsgContextInfo, 0, 30)
	opsMsg.MsgContext = OpsMsgContextInfo{}
	opsMsg.EmptyMsgData()
}

// EmptyMsgData - Resets all OpsMsgDto fields, with
// the exception of ParentContextHistory and MsgContext,
// to an uninitialized or 'empty' state.
func (opsMsg *OpsMsgDto) EmptyMsgData() {
	opsMsg.Message 					= ""
	opsMsg.FmtMessage   		= ""
	opsMsg.msgId          	= int64(0) // The identifying number for this message
	opsMsg.msgNumber      	= int64(0) //  Message Number = msgId + MsgContext.BaseMessageId. This is the number displayed in the message
	opsMsg.MsgType        	= OpsMsgTypeNOERRORNOMSG
	opsMsg.MsgClass       	= OpsMsgClassNOERRORSNOMESSAGES
	opsMsg.MsgTimeUTC     	= time.Time{}
	opsMsg.MsgTimeLocal   	= time.Time{}
	opsMsg.MsgLocalTimeZone	= ""
}

// GetError - If the current OpsMsgDto is
// configured as either a Standard Error or
// Fatal Error, this method will return
// an 'error' type containing the error
// message. If OpsMsgDto is configured as
// a non-error type message, this method
// will return 'nil'.
func (opsMsg *OpsMsgDto) GetError() error {

	if opsMsg.IsError() {
		return errors.New(opsMsg.GetMessage())
	}

	return nil

}

// GetMessage - Returns the Operations Message
// stored in this object. Note that the underling
// message is stored as a string array and may
// therefore accommodate multiple messages.
func (opsMsg *OpsMsgDto) GetMessage() string {

	return opsMsg.FmtMessage
}

// GetMessageId - returns data field 'msgId' for
// the current OpsMsgDto object.
func (opsMsg *OpsMsgDto) GetMessageId() int64 {
	return opsMsg.msgId
}

// GetMessageNumber - returns the data field 'msgNumber'
// for the current OpsMsgDto object. The 'msgNumber' is
// equal to 'msgId' plus 'opsMsg.MsgContext.BaseMessageId'
func (opsMsg *OpsMsgDto) GetMessageNumber() int64 {

	return opsMsg.msgNumber
}

// InitializeContextInfo - Initializes Parent Context History and Message Context Info for a new
// OpsMsgDto object.
//
// Input Parameters:
// =================
//
// parentHistory []OpsMsgContextInfo - An array of OpsMsgContextInfo objects
// 											documenting execution path that led to the generation
//											of this method.
//
// msgContext OpsMsgContextInfo - This object records the current context in
//											which the new OpsMsgDto returned by this method will
//											will be operating.
//
// 											It allows the newly created OpsMsgDto to return data
// 											on the execution path which	led to the generation of
// 											the Operations Message.
//
// Example Usage:
// ==============
//
// oMsg := OpsMsgDto{}.InitializeContextInfo(parentHistory, msgContext)
//
// Parent Context History and current Message Context serve as an important
// purpose. It allows one to maintain a record of the function execution tree
// that led to the generation of this message.
//
func(opsMsg OpsMsgDto) InitializeContextInfo(parentHistory []OpsMsgContextInfo, msgContext OpsMsgContextInfo) OpsMsgDto {
	om := OpsMsgDto{}
	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentHistory)
	om.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()

	return om
}

// InitializeContextWithParentOpsMsg - Initialize a new OpsMsgDto
// object with Parent History data extracted from another OpsMsgDto
// object.
//
// Input Parameters:
// =================
//
//	parentOpsMsg OpsMsgDto 	- The Parent History context from the incoming
//														OpsMsgDto will be added to the new OpsMsgDto
//														object being created by this method. In addition,
//														the Parent OpsMsgDto MsgContext will be added to
//														current OpsMsgDto ParentContextHistory.
//
//	newMsgContext OpsMsgContextInfo - This new OpsMsgContextInfo object will
//										be configured as the 'MsgContext' field in the new
// 										OpsMsgDto object created by this method.
//
//										It allows the newly created OpsMsgDto to return data
// 										on the execution path which	led to the generation of
// 										the Operations Message.
//
//	Example Usage:
//  ==============
//
//	parentOpsMsgDto // OpsMsgDto object created in the parent function
//	currentMsgContext = OpsMsgContextInfo{SourceFileName:"xray.go",
// 											ParentObjectName: "stringutil", FuncName:"DoSomeWork",
//											BaseMessgeId:int64(8000)
//
// Parent Context History and current Message Context serve as an important
// purpose. It allows one to maintain a record of the function execution tree
// that led to the generation of this message.
//
func(opsMsg OpsMsgDto) InitializeContextWithParentOpsMsg(parentOpsMsg OpsMsgDto, newMsgContext OpsMsgContextInfo) OpsMsgDto {

	om := OpsMsgDto{}

	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentOpsMsg.ParentContextHistory)
	om.AddOpsMsgContextInfoToParentHistory(parentOpsMsg.MsgContext)
	om.MsgContext = newMsgContext.DeepCopyOpsMsgContextInfo()

	return om
}

// IsFatalError - If the current OpsMsgDto object is configured
// as a fatal error, this method will return true.
func (opsMsg *OpsMsgDto) IsFatalError() bool {

	if opsMsg.MsgClass == OpsMsgClassFATAL {
		return true
	}

	return false

}

// IsError - Returns a boolean value signaling
// whether the current OpsMsgDto is an 'error'
// message.
//
// If the return value is true, the OpsMsgDto
// will be configured as either a Standard Error
// or a Fatal Error. (See Method IsFatalError())
func (opsMsg *OpsMsgDto) IsError() bool {

	if opsMsg.MsgType == OpsMsgTypeERRORMSG {
		return true
	}

	return false
}

// NewDebugMsg - Create a new Debug Message
//
// Input Parameters
// ****************
//
//	msg string 		- The text of the Debug Message
//
//	msgNo	int64		- The message number to be associated with
//									this message. If 'msgNo' is equal to zero,
//									no message number will be displayed in the
//									final message
func(opsMsg OpsMsgDto) NewDebugMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetDebugMessage(msg, msgNo)

	return om
}


// NewInfoMsg - Create a new Operations Message which is
// an Informational Message.
//
// Input Parameters
// ****************
//
//	msg string 		- The text of the Information Message
//
//	msgNo	int64		- The message number to be associated with
//									this message. If 'msgNo' is equal to zero,
//									no message number will be displayed in the
//									final message
func(opsMsg OpsMsgDto) NewInfoMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetInfoMessage(msg, msgNo)

	return om
}

// NewFatalErrorMsg - Creates a New FATAL Error Message
//
// Input Parameters
// ****************
//
//	errMsg string	- The text of the Error Message
//
//	errNo	int64		- The error number to be associated with
//									this message. If 'errNo' is equal to zero,
//									no error number will be displayed in the
//									final error message
func (opsMsg OpsMsgDto) NewFatalErrorMsg(errMsg string, errNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetFatalErrorMessage(errMsg, errNo)
	return om

}

// NewStdErrorMsg - Creates a new non-fatal error message
//
// Input Parameters
// ****************
//
//	errMsg string	- The text of the Error Message
//
//	errNo	int64		- The error number to be associated with
//									this message. If 'errNo' is equal to zero,
//									no error number will be displayed in the
//									final error message
func (opsMsg OpsMsgDto) NewStdErrorMsg(errMsg string, errNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetStdErrorMessage(errMsg, errNo)
	return om

}

// NewMsgFromSpecErrMsg - Create a new Operations Message based on
// the error information contained in a Type SpecErr passed
// into the method. The new message will be designated as
// an error message.
func (opsMsg *OpsMsgDto) NewMsgFromSpecErrMsg(se SpecErr) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetFromSpecErrMessage(se)

	return om
}


// NewWarningMsg - Creates a new Warning Message
// and returns it as a new OpsMsgDto object.
func (opsMsg OpsMsgDto) NewWarningMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}

	om.SetWarningMessage(msg, msgNo)

	return om

}

// SetDebugMessage - Configures the current or host
// OpsMsgDto object as a 'DEBUG' message.
func (opsMsg *OpsMsgDto) SetDebugMessage(msg string, msgId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeDEBUGMSG
	opsMsg.MsgClass = OpsMsgClassDEBUG

	opsMsg.setMsgText(msg, msgId)

}

// SetFatalErrorMessage - Configures the current or host
// OpsMsgDto object as an information message.
func (opsMsg *OpsMsgDto) SetFatalErrorMessage(errMsg string, errId int64) {

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeERRORMSG
	opsMsg.MsgClass = OpsMsgClassFATAL

	opsMsg.setMsgText(errMsg, errId)

}

// SetFromSpecErrMessage - Sets an instance of OpsMsgDto type based
// on a SpecErr object passed as an input parameter.
func (opsMsg *OpsMsgDto) SetFromSpecErrMessage(se SpecErr) {

	opsMsg.Empty()
	
	x := se.DeepCopyParentInfo(se.ParentInfo)

	for _, bi := range x {
		ci := OpsMsgContextInfo{SourceFileName:bi.SourceFileName, ParentObjectName: bi.ParentObjectName, FuncName: bi.FuncName, BaseMessageId: bi.BaseErrorId}
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, ci)
	}

	y := se.DeepCopyBaseInfo()

	opsMsg.MsgContext = OpsMsgContextInfo{SourceFileName:y.SourceFileName, ParentObjectName: y.ParentObjectName, FuncName: y.FuncName, BaseMessageId: y.BaseErrorId}


	switch se.ErrorMsgType {

	case SpecErrTypeNOERRORSALLCLEAR:
		opsMsg.SetNoErrorsNoMessages(se.ErrId)

	case SpecErrTypeERROR:
		opsMsg.SetStdErrorMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeFATAL:
		opsMsg.SetFatalErrorMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeINFO:
		opsMsg.SetInfoMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeWARNING:
		opsMsg.SetWarningMessage(se.ErrMsg, se.ErrId )

	case SpecErrTypeSUCCESSFULCOMPLETION:
		opsMsg.SetSuccessfulCompletionMessage(se.ErrId)

	default:
		panic("OpsMsgDto.SetFromSpecErrMessage() - INVALID SpecErrType Code")
	}

}

// SetInfoMessage - Configures the current or host
// OpsMsgDto object as an information message.
func (opsMsg *OpsMsgDto) SetInfoMessage(msg string, msgId int64) {
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeINFOMSG
	opsMsg.MsgClass = OpsMsgClassINFO

	opsMsg.setMsgText(msg, msgId)
}

// SetStdErrorMessage - Configures the current or host
// OpsMsgDto object as a standard error message.
func (opsMsg *OpsMsgDto) SetStdErrorMessage(errMsg string, errId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeERRORMSG
	opsMsg.MsgClass = OpsMsgClassOPERROR

	opsMsg.setMsgText(errMsg, errId)

}

// SetNoErrorsNoMessages - Configures the current or host
// OpsMsgDto object for the default message type,
// 'No Errors and No Messages'.
func (opsMsg *OpsMsgDto) SetNoErrorsNoMessages(msgId int64) {

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeNOERRORNOMSG
	opsMsg.MsgClass = OpsMsgClassNOERRORSNOMESSAGES

	opsMsg.setMsgText("No Errors - No Messages", msgId)

}

// SetSuccessfulCompletionMessage - Configures the current or host
// OpsMsgDto object as a Successful Completion Message.
func (opsMsg *OpsMsgDto) SetSuccessfulCompletionMessage(msgId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeSUCCESSFULCOMPLETION
	opsMsg.MsgClass = OpsMsgClassSUCCESSFULCOMPLETION

	opsMsg.setMsgText("Successful Completion", msgId)

}

// SetWarningMessage - Configures the current or host
// OpsMsgDto object as a Warning Message.
func (opsMsg *OpsMsgDto) SetWarningMessage(msg string, msgId int64) {
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeWARNINGMSG
	opsMsg.MsgClass = OpsMsgClassWARNING

	opsMsg.setMsgText(msg, msgId)

}

// String - returns the operations message as a
// string.
func (opsMsg *OpsMsgDto) String() string {
	return opsMsg.GetMessage()
}

// ***********************************************
// private methods
// ***********************************************

// getMsgTitle - Returns the Message title, message number and the
// banner line separator based on value of OpsMsgDto.MsgClass
func (opsMsg *OpsMsgDto) getMsgTitle() (banner1, banner2, title, numTitle string, ) {

	switch opsMsg.MsgClass {

	case OpsMsgClassNOERRORSNOMESSAGES:
		// OpsMsgClassNOERRORSNOMESSAGES - 0 Signals uninitialized message
		// with no errors and no messages
		title = "Empty Message - No Errors and No Messages"
		numTitle = "Msg Number"
		banner1 = strings.Repeat("%", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassOPERROR:
		// OpsMsgClassOPERROR - 1 Message is an Error Message
		title = "Standard ERROR Message"
		numTitle = "Error Number: "
		banner1 = strings.Repeat("#", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassFATAL:
		// OpsMsgClassFATAL - 2 Message is a Fatal Error Message
		title = "FATAL ERROR Message"
		numTitle = "Error Number: "
		banner1 = strings.Repeat("!", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassINFO:
		// OpsMsgClassINFO - 3 Message is an Informational Message
		title = "Information Message"
		numTitle = "Info Msg Number"
		banner1 = strings.Repeat("*", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassWARNING:
		// OpsMsgClassWARNING - 4 Message is a warning Message
		title = "WARNING Message"
		numTitle = "Warning Msg Number"
		banner1 = strings.Repeat("?", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassDEBUG:
		// OpsMsgClassDEBUG - 5 Message is a Debug Message
		title = "DEBUG Message"
		numTitle = "DEBUG - Message Number"
		banner1 = strings.Repeat("@", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassSUCCESSFULCOMPLETION:
		// OpsMsgClassSUCCESSFULCOMPLETION - 6 Message signalling successful
		// completion of the operation
		title = "Successful Completion"
		numTitle = "Successful Completion Msg Number"
		banner1 = strings.Repeat("&", 78)
		banner2 = strings.Repeat("-", 78)

	default:
		// This should never happen
		panic("OpsMsgDto.getMsgTitle() - Invalid opsMsg.MsgClass")
	}

	return banner1, banner2, title, numTitle
}


func(opsMsg *OpsMsgDto) setDebugMsgText(banner1, banner2, title, numTitle string) {

	m := "\n\n"
	m += "\n" + banner1
	m += "\n" + title
	// FmtDateTimeTzNanoYMD
	dt := DateTimeUtility{}
	timeStamp := "Local Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeLocal)
	timeStamp += "    UTC Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeUTC)

	m += "\n"
	if opsMsg.msgId != 0 {
		m+= numTitle + ": " + string(opsMsg.msgNumber)
		m+= "    Time Stamp: " + timeStamp
	} else {
		m+= "Time Stamp: " + timeStamp
	}

	m += "\nMessage: " + opsMsg.Message
	m += "\n" + banner1


	opsMsg.FmtMessage =  m
}

func(opsMsg *OpsMsgDto) setMsgText(msg string, msgId int64) {

	opsMsg.setMsgIdAndMsgNumber(msgId)

	opsMsg.setTime("Local")

	opsMsg.Message = msg

	var m string
	banner1, banner2, title, numTitle := opsMsg.getMsgTitle()

	if opsMsg.MsgClass == OpsMsgClassSUCCESSFULCOMPLETION {
		opsMsg.setSuccessfulCompletionMsgText(banner1, banner2, title, numTitle)
		return
	}

	if opsMsg.MsgClass == OpsMsgClassNOERRORSNOMESSAGES {
		opsMsg.setEmptyMessageText(banner1, banner2, title, numTitle)
		return
	}


	if opsMsg.MsgClass == OpsMsgClassDEBUG {
		opsMsg.setDebugMsgText(banner1, banner2, title, numTitle)
		return
	}

	m= "\n\n"
	m += "\n" + banner1
	m += "\n     " + title
	m += "\n" + banner1
	if opsMsg.msgNumber != 0 {
		m+= "\n"  + numTitle + ": " + string(opsMsg.msgNumber)
	}
	m += "\n" + msg
	m += "\n" + banner2


	m += "\n Message Type: " + opsMsg.MsgType.String()
	m += "\nMessage Class: " + opsMsg.MsgClass.String()
	m += "\n" + banner2
	m += "\n" + "Time Stamp:"
	m += "\n" + banner2
	dt := DateTimeUtility{}
	dtFmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
	m += fmt.Sprintf("\n  Message Time UTC: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeUTC, dtFmt))
	m += fmt.Sprintf("\nMessage Time Local: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeLocal, dtFmt))
	m += "\n   Local Time Zone:"
	localTz := opsMsg.MsgLocalTimeZone

	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}

	m += localTz
	m += "\n" + banner1

	opsMsg.FmtMessage =  m
}

func (opsMsg *OpsMsgDto) setEmptyMessageText(banner1, banner2, title, numTitle string) {
	m := "\n\n"
	m += "\n" + banner1
	m += "\n     " + title
	m += "\n" + banner1
	if opsMsg.msgNumber != 0 {
		m+= "\n"  + numTitle + ": " + string(opsMsg.msgNumber)
	}
	m += "\n" + banner2
	m += "\n" + "Time Stamp:"
	m += "\n" + banner2
	dt := DateTimeUtility{}
	dtFmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
	m += fmt.Sprintf("\n  Message Time UTC: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeUTC, dtFmt))
	m += fmt.Sprintf("\nMessage Time Local: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeLocal, dtFmt))
	m += "\n   Local Time Zone:"
	localTz := opsMsg.MsgLocalTimeZone

	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}

	m += localTz
	m += "\n" + banner1

	opsMsg.FmtMessage =  m

}

func (opsMsg *OpsMsgDto) setMsgIdAndMsgNumber(msgId int64) {
	
	if msgId == 0 {
		opsMsg.msgId = 0
		opsMsg.msgNumber = 0
	} else {
		opsMsg.msgId = msgId
		opsMsg.msgNumber = msgId + opsMsg.MsgContext.BaseMessageId
	}
	
	
}

func (opsMsg *OpsMsgDto) setSuccessfulCompletionMsgText(banner1, banner2, title, numTitle string) {
	m := "\n\n"
	m += "\n" + banner1
	m += "\n     " + title
	m += "\n" + banner1
	if opsMsg.msgNumber != 0 {
		m+= "\n"  + numTitle + ": " + string(opsMsg.msgNumber)
	}
	m += "\n" + banner2
	m += "\n" + "Time Stamp:"
	m += "\n" + banner2
	dt := DateTimeUtility{}
	dtFmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
	m += fmt.Sprintf("\n  Message Time UTC: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeUTC, dtFmt))
	m += fmt.Sprintf("\nMessage Time Local: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeLocal, dtFmt))
	m += "\n   Local Time Zone:"
	localTz := opsMsg.MsgLocalTimeZone

	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}

	m += localTz
	m += "\n" + banner1

	opsMsg.FmtMessage =  m

}


// setTime - Sets the time stamp for this Operations
// Message. Notice that the input parameter 'localTimeZone'
// is the Iana Time Zone for local time.
//
// Reference Iana Time Zones: https://www.iana.org/time-zones
//
// If the 'localTimeZone' parameter string is empty or an
// invalid time zone, local time zone will default to 'Local'.
// The 'Local' time zone is determined by the host computer.
func(opsMsg *OpsMsgDto) setTime(localTimeZone string){

	tz := TimeZoneUtility{}

	isValid, _, _ := tz.IsValidTimeZone(localTimeZone)

	if !isValid {
		localTimeZone = "Local"
	}

	opsMsg.MsgTimeUTC = time.Now().UTC()
	opsMsg.MsgLocalTimeZone = localTimeZone

	tzLocal, _ := tz.ConvertTz(opsMsg.MsgTimeUTC, opsMsg.MsgLocalTimeZone)
	opsMsg.MsgTimeLocal = tzLocal.TimeOut

}

