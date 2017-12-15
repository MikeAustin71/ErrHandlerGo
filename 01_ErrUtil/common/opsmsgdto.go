package common

import (
	"time"
	"strings"
	"fmt"
	"errors"
	"unicode/utf8"
)

/*  'opsmsgdto.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

	Dependencies:
	=============

	This file depends on utility routines provided by source code
	file, 'datetimeutility.go'. This source code file is located
	in the source code repository:

						https://github.com/MikeAustin71/datetimeopsgo.git
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

// Equal - Compares two OpsMsgContextInfo objects
// to determine if they are equivalent.
func (ci *OpsMsgContextInfo) Equal(ci2 *OpsMsgContextInfo) bool {
	if ci.SourceFileName 		!= ci2.SourceFileName 	||
			ci.ParentObjectName != ci2.ParentObjectName ||
			ci.FuncName 				!= ci2.FuncName					||
			ci.BaseMessageId    != ci2.BaseMessageId		{

				return false
	}

	return true
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

// CopyIn - Receives an OpsMsgDto object as input.
// Then a deep copy is created and used to populate
// the current OpsMsgDto object.
func (opsMsg *OpsMsgDto) CopyIn(opsMsg2 *OpsMsgDto) {
	opsMsg.Empty()
	opsMsg.ParentContextHistory = opsMsg2.DeepCopyParentContextHistory(opsMsg2.ParentContextHistory)
	opsMsg.MsgContext = opsMsg2.MsgContext.DeepCopyOpsMsgContextInfo()

	opsMsg.Message       		= opsMsg2.Message
	opsMsg.FmtMessage				= opsMsg2.FmtMessage
	opsMsg.msgId            = opsMsg2.GetMessageId()
	opsMsg.msgNumber        = opsMsg2.GetMessageNumber()
	opsMsg.MsgType          = opsMsg2.MsgType
	opsMsg.MsgClass         = opsMsg2.MsgClass
	opsMsg.MsgTimeUTC       = opsMsg2.MsgTimeUTC
	opsMsg.MsgTimeLocal     = opsMsg2.MsgTimeLocal
	opsMsg.MsgLocalTimeZone = opsMsg2.MsgLocalTimeZone

}

// CopyOut - Makes a deep copy of the current OpsMsgDto
// and returns it as a new OpsMsgDto with equivalent
// field values.
func (opsMsg *OpsMsgDto) CopyOut() OpsMsgDto {
	
	opsMsg2 := OpsMsgDto{}

	opsMsg2.ParentContextHistory = opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory)
	opsMsg2.MsgContext = opsMsg.MsgContext.DeepCopyOpsMsgContextInfo()

	opsMsg2.Message       	= opsMsg.Message
	opsMsg2.FmtMessage			= opsMsg.FmtMessage
	opsMsg2.msgId            = opsMsg.GetMessageId()
	opsMsg2.msgNumber        = opsMsg.GetMessageNumber()
	opsMsg2.MsgType          = opsMsg.MsgType
	opsMsg2.MsgClass         = opsMsg.MsgClass
	opsMsg2.MsgTimeUTC       = opsMsg.MsgTimeUTC
	opsMsg2.MsgTimeLocal     = opsMsg.MsgTimeLocal
	opsMsg2.MsgLocalTimeZone = opsMsg.MsgLocalTimeZone
	
	return opsMsg2
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

	opsMsg.EmptyParentHistory()
	opsMsg.EmptyMessageContext()
	opsMsg.EmptyMsgData()
}

// EmptyParentHistory - Deletes the current ParentHistory ([] OpsMsgContextInfo)
// and resets it to an 'empty' or uninitialized state (zero length array).
func (opsMsg *OpsMsgDto) EmptyParentHistory() {
	opsMsg.ParentContextHistory = make([] OpsMsgContextInfo, 0, 30)
}

// EmptyMessageContext - Deletes the current message context
// (OpsMsgDto.MsgContext) and resets it to an uninitialized state.
func (opsMsg *OpsMsgDto) EmptyMessageContext() {
	opsMsg.MsgContext = OpsMsgContextInfo{}
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

// Equal - Compares an incoming OpsMsgDto object
// with the current OpsMsgDto object to determine
// if they are equivalent.
func (opsMsg *OpsMsgDto) Equal(opsMsg2 *OpsMsgDto) bool {

	l1 := len(opsMsg.ParentContextHistory)
	l2 := len(opsMsg2.ParentContextHistory)

	if l1 != l2 {
		return false
	}

	for i:= 0; i < l1; i++ {
		if !opsMsg.ParentContextHistory[i].Equal(&opsMsg2.ParentContextHistory[i]) {
			return false
		}
	}

	if !opsMsg.MsgContext.Equal(&opsMsg2.MsgContext) {
		return false
	}

	if  opsMsg.Message     			!= opsMsg2.Message 						||
			opsMsg.FmtMessage				!= opsMsg2.FmtMessage					||
			opsMsg.msgId            != opsMsg2.GetMessageId()			||
			opsMsg.msgNumber        != opsMsg2.GetMessageNumber()	||
			opsMsg.MsgType          != opsMsg2.MsgType						||
			opsMsg.MsgClass         != opsMsg2.MsgClass						||
			opsMsg.MsgTimeUTC       != opsMsg2.MsgTimeUTC					||
			opsMsg.MsgTimeLocal     != opsMsg2.MsgTimeLocal				||
			opsMsg.MsgLocalTimeZone != opsMsg2.MsgLocalTimeZone {

		return false
	}

	return true

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

// InitializeAllContextInfo - Initializes Parent Context History and Message Context Info for a new
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
// oMsg := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgContext)
//
// Parent Context History and current Message Context serve as an important
// purpose. It allows one to maintain a record of the function execution tree
// that led to the generation of this message.
//
func(opsMsg OpsMsgDto) InitializeAllContextInfo(parentHistory []OpsMsgContextInfo, msgContext OpsMsgContextInfo) OpsMsgDto {
	om := OpsMsgDto{}
	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentHistory)
	om.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()

	return om
}


// InitializeWithMessageContext - Initialize a new OpsMsgDto object
// with only a Message Context - No ParentHistory.
func(opsMsg OpsMsgDto) InitializeWithMessageContext(msgContext OpsMsgContextInfo) OpsMsgDto {
	om := OpsMsgDto{}
	om.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()
	return om
}

// InitializeContextWithParentHistoryPlusMsgContext - Initialize a new OpsMsgDto
// object with Parent History plus the OpsMsgContextInfo object passed as an input
// parameter, 'newMsgContext'.
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
func(opsMsg OpsMsgDto) InitializeContextWithParentHistoryPlusMsgContext(parentOpsMsg OpsMsgDto, newMsgContext OpsMsgContextInfo) OpsMsgDto {

	om := OpsMsgDto{}

	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentOpsMsg.ParentContextHistory)
	om.AddOpsMsgContextInfoToParentHistory(parentOpsMsg.MsgContext)
	om.MsgContext = newMsgContext.DeepCopyOpsMsgContextInfo()

	return om
}

// IsDebugMsg - Returns a boolean value indicating
// whether or not this message is a 'Debug'
// type message.
func(opsMsg OpsMsgDto) IsDebugMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeDEBUGMSG {
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

// IsFatalError - If the current OpsMsgDto object is configured
// as a fatal error, this method will return true.
func (opsMsg *OpsMsgDto) IsFatalError() bool {

	if opsMsg.MsgClass == OpsMsgClassFATAL {
		return true
	}

	return false

}

// IsInfoMsg - Returns a boolean value indicating
// whether or not this message is an 'Information'
// type message.
func (opsMsg *OpsMsgDto) IsInfoMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeINFOMSG {
		return true
	}

	return false


}

// IsNoErrorsNoMessages  - Returns a boolean value indicating
// whether or not this message is a 'No Errors No Messages'
// type message.
//
// 'No Errors No Messages' is the type of message assigned to
// uninitialized OpsMsgDto objects.
func (opsMsg *OpsMsgDto) IsNoErrorsNoMessages() bool {

	if opsMsg.MsgType == OpsMsgTypeNOERRORNOMSG {
		return true
	}

	return false

}

// IsSuccessfulCompletion - Returns a boolean value indicating
// whether or not this message is a 'Successful Completion' type
// message.
func (opsMsg *OpsMsgDto) IsSuccessfulCompletionMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeSUCCESSFULCOMPLETION {
		return true
	}

	return false

}

// IsWarningMsg - Returns a boolean value indicating
// whether or not this message is a 'Warning' type
// message.
func (opsMsg *OpsMsgDto) IsWarningMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeWARNINGMSG {
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
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
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
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
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
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetFatalErrorMessage(errMsg, errNo)
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
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetStdErrorMessage(errMsg, errNo)
	return om
}

// NewSuccessfulCompletionMsg - Creates a new Successful Completion
// Message and returns it as a new OpsMsgDto object.
func (opsMsg OpsMsgDto) NewSuccessfulCompletionMsg(msg string, msgId int64) OpsMsgDto {
	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetSuccessfulCompletionMessage(msg, msgId)
	return om
}

// NewWarningMsg - Creates a new Warning Message
// and returns it as a new OpsMsgDto object.
func (opsMsg OpsMsgDto) NewWarningMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)

	om.SetWarningMessage(msg, msgNo)

	return om

}

// NewNoErrorsNoMessagesMsg - Creates a new No Errors and No
// Messages Message and returns it as a new OpsMsgDto object.
func (opsMsg OpsMsgDto) NewNoErrorsNoMessagesMsg(msg string,msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)

	om.SetNoErrorsNoMessages(msg, msgNo)

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
		opsMsg.SetNoErrorsNoMessages(se.ErrMsg, se.ErrId)

	case SpecErrTypeERROR:
		opsMsg.SetStdErrorMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeFATAL:
		opsMsg.SetFatalErrorMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeINFO:
		opsMsg.SetInfoMessage(se.ErrMsg, se.ErrId)

	case SpecErrTypeWARNING:
		opsMsg.SetWarningMessage(se.ErrMsg, se.ErrId )

	case SpecErrTypeSUCCESSFULCOMPLETION:
		opsMsg.SetSuccessfulCompletionMessage(se.ErrMsg, se.ErrId)

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

// SetMsgContext - Receives an OpsMsgContextInfo object as
// an input parameter and configures the current OpsMsgDto
// MessageContext.
func (opsMsg *OpsMsgDto) SetMessageContext(msgContext OpsMsgContextInfo) {
	opsMsg.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()
}

// SetParentMessageContextHistory - Deletes the current opsMsg.ParentContextHistory
// and replaeces it with the input parameter, 'parentHistory',
func (opsMsg *OpsMsgDto) SetParentMessageContextHistory( parentHistory []OpsMsgContextInfo) {
	opsMsg.ParentContextHistory = make([] OpsMsgContextInfo, 0, 30)
	l1 := len(parentHistory)

	for i:= 0; i < l1; i++ {
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, parentHistory[i])
	}

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
func (opsMsg *OpsMsgDto) SetNoErrorsNoMessages(msg string, msgId int64) {

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeNOERRORNOMSG
	opsMsg.MsgClass = OpsMsgClassNOERRORSNOMESSAGES

	opsMsg.setMsgText(msg, msgId)

}

// SetSuccessfulCompletionMessage - Configures the current or host
// OpsMsgDto object as a Successful Completion Message.
func (opsMsg *OpsMsgDto) SetSuccessfulCompletionMessage(msg string, msgId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeSUCCESSFULCOMPLETION
	opsMsg.MsgClass = OpsMsgClassSUCCESSFULCOMPLETION

	opsMsg.setMsgText( msg, msgId)

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
		title = "No Errors and No Messages"
		numTitle = "Msg Number"
		banner1 = strings.Repeat("&", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassOPERROR:
		// OpsMsgClassOPERROR - 1 Message is an Error Message
		title = "Standard ERROR Message"
		numTitle = "Error No"
		banner1 = strings.Repeat("#", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassFATAL:
		// OpsMsgClassFATAL - 2 Message is a Fatal Error Message
		title = "FATAL ERROR Message"
		numTitle = "Error No"
		banner1 = strings.Repeat("!", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassINFO:
		// OpsMsgClassINFO - 3 Message is an Informational Message
		title = "Information Message"
		numTitle = "Msg No"
		banner1 = strings.Repeat("*", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassWARNING:
		// OpsMsgClassWARNING - 4 Message is a warning Message
		title = "WARNING Message"
		numTitle = "Msg No"
		banner1 = strings.Repeat("?", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassDEBUG:
		// OpsMsgClassDEBUG - 5 Message is a Debug Message
		title = "DEBUG Message"
		numTitle = " Number"
		banner1 = strings.Repeat("@", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgClassSUCCESSFULCOMPLETION:
		// OpsMsgClassSUCCESSFULCOMPLETION - 6 Message signalling successful
		// completion of the operation
		title = "Successful Completion"
		numTitle = "Msg No"
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

	if opsMsg.msgNumber != 0 {
		m += fmt.Sprintf("\n %v %v: %v", title, numTitle, opsMsg.msgNumber)
	} else {
		m += fmt.Sprintf("\n %v -", title)
	}

	m += "\n  " + opsMsg.Message

	l1 := len(opsMsg.ParentContextHistory)
	if l1 > 0 {
		m += "\n" + banner2
		m += "\n Parent Context History:"
		for i:=0; i < l1; i++ {
			m+= "\n  Src File: " + opsMsg.ParentContextHistory[i].SourceFileName
			m+= "   Parent Obj: " + opsMsg.ParentContextHistory[i].ParentObjectName
			m+= "   Func Name: " + opsMsg.ParentContextHistory[i].FuncName
		}

	}

	if opsMsg.MsgContext.SourceFileName != "" ||
		opsMsg.MsgContext.ParentObjectName != "" ||
		opsMsg.MsgContext.FuncName != "" {
		m += "\n" + banner2
		m += "\n Current Message Context:"
		if opsMsg.MsgContext.SourceFileName != "" {
			m+= "\n  Src File: " + opsMsg.MsgContext.SourceFileName
		}

		if opsMsg.MsgContext.ParentObjectName != "" {
			m+= "   Parent Obj: " + opsMsg.MsgContext.ParentObjectName
		}

		if opsMsg.MsgContext.FuncName != "" {
			m+= "   Func Name: " + opsMsg.MsgContext.FuncName
		}
	}

	// FmtDateTimeTzNanoYMD
	dt := DateTimeUtility{}
	localTz := opsMsg.MsgLocalTimeZone

	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}
	m += "\n" + banner2
	m += "\n   UTC Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeUTC)
	m += "\n Local Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeLocal) + "   Time Zone: " + localTz

	m += "\n" + banner1


	opsMsg.FmtMessage =  m
}

func(opsMsg *OpsMsgDto) setMsgText(msg string, msgId int64) {

	opsMsg.setMsgIdAndMsgNumber(msgId)

	opsMsg.setTime("Local")

	opsMsg.Message = msg

	var m string
	banner1, banner2, title, numTitle := opsMsg.getMsgTitle()

	if opsMsg.MsgClass == OpsMsgClassDEBUG {
		opsMsg.setDebugMsgText(banner1, banner2, title, numTitle)
		return
	}
	lineWidth := len(banner1)
	dt := DateTimeUtility{}
	dtFmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"

	m= "\n\n"
	m += "\n" + banner1
	nextBanner := banner1
	s1 := (lineWidth / 3) * 2
	s2 := lineWidth - s1

	if opsMsg.msgNumber != 0 {
		sNo:= fmt.Sprintf("%v: %v", numTitle, opsMsg.msgNumber)
		str1, _ := opsMsg.strCenterInStr(title, s1)
		str2, _ := opsMsg.strRightJustify(sNo, s2)
		m+= "\n" + str1 + str2
	} else {
		str1, _ := opsMsg.strCenterInStr(title, s1)
		m+= "\n" + str1
	}

	if opsMsg.Message != "" {
		m += "\n" + nextBanner

		m += "\n Message: "

		if len(opsMsg.Message) > 67 {
			m += "\n  "
		}

		m += opsMsg.Message

		nextBanner = banner2
	}

	l1 := len(opsMsg.ParentContextHistory)
	if l1 > 0 {
		m += "\n" + nextBanner
		m += "\n Parent Context History:"
		for i:=0; i < l1; i++ {
			m+= "\n  Src File: " + opsMsg.ParentContextHistory[i].SourceFileName
			m+= "   Parent Obj: " + opsMsg.ParentContextHistory[i].ParentObjectName
			m+= "   Func Name: " + opsMsg.ParentContextHistory[i].FuncName
		}

		nextBanner = banner2
	}

	if opsMsg.MsgContext.SourceFileName != "" ||
			opsMsg.MsgContext.ParentObjectName != "" ||
				opsMsg.MsgContext.FuncName != "" {
		m += "\n" + nextBanner
		nextBanner = banner2
		m += "\n Current Message Context:"
		if opsMsg.MsgContext.SourceFileName != "" {
			m+= "\n  Src File: " + opsMsg.MsgContext.SourceFileName
		}

		if opsMsg.MsgContext.ParentObjectName != "" {
			m+= "   Parent Obj: " + opsMsg.MsgContext.ParentObjectName
		}

		if opsMsg.MsgContext.FuncName != "" {
			m+= "   Func Name: " + opsMsg.MsgContext.FuncName
		}
	}

	m += "\n" + nextBanner
	m += fmt.Sprintf("\n   Message Time UTC: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeUTC, dtFmt))
	m += fmt.Sprintf("\n Message Time Local: %v ", dt.GetDateTimeCustomFmt(opsMsg.MsgTimeLocal, dtFmt))
	m += "\n" + banner1

	opsMsg.FmtMessage =  m
}

// setMsgIdAndMsgNumber - This method is called internally
// to set the OpsMsgDto.msgId and OpsMsgDto.msgNumber fields.
func (opsMsg *OpsMsgDto) setMsgIdAndMsgNumber(msgId int64) {
	
	if msgId == 0 {
		opsMsg.msgId = 0
		opsMsg.msgNumber = 0
	} else {
		opsMsg.msgId = msgId
		opsMsg.msgNumber = msgId + opsMsg.MsgContext.BaseMessageId
	}
	
	
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


/*

Private String Management Methods

*/

// strCenterInStr - returns a string which includes
// a left pad blank string plus the original string.
// The complete string will effectively center the
// original string is a field of specified length.
func (opsMsg *OpsMsgDto) strCenterInStr(strToCenter string, fieldLen int) (string, error) {

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
func (opsMsg *OpsMsgDto) strRightJustify(strToJustify string, fieldLen int) (string, error) {

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
func (opsMsg *OpsMsgDto) strPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

	sLen := opsMsg.strGetRuneCnt(strToCenter)

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
func (opsMsg *OpsMsgDto) strGetRuneCnt(targetStr string) int {
	return utf8.RuneCountInString(targetStr)
}

// strGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (opsMsg *OpsMsgDto) strGetCharCnt(targetStr string) int {
	return len([]rune(targetStr))
}


