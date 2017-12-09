package common

import (
	"time"
	"strings"
	"fmt"
)

/*  'opsmsgdto.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/

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
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, FuncName: ci.FuncName, BaseMessageId: ci.BaseMessageId}
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
	Message              []string
	msgId                int64 // The identifying number for this message
	msgNumber            int64 //  Message Number = msgId + MsgContext.BaseMessageId. This is the number displayed in the message
	MsgType              OpsMsgType
	MsgClass             OpsMsgClass
	MsgTimeUTC           time.Time
	MsgTimeLocal         time.Time
	MsgLocalTimeZone     string
	ErrDto               SpecErr
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


// GetMessage - Returns the Operations Message
// stored in this object. Note that the underling
// message is stored as a string array and may
// therefore accommodate multiple messages.
func (opsMsg *OpsMsgDto) GetMessage() string {

	output := ""

	for i:=0; i < len(opsMsg.Message); i++ {
		if i==0 {
			output = opsMsg.Message[i]
		} else {
			output += "\n"
			output += opsMsg.Message[i]
		}

	}

	return output
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
	om.MsgType = OpsMsgTypeDEBUGMSG
	om.MsgClass = OpsMsgClassDEBUG

	if msgNo == 0 {
		om.msgId = 0
	} else {
		om.msgId = msgNo
		om.msgNumber = msgNo + om.MsgContext.BaseMessageId
	}

	om.SetTime("Local")
	om.setDebugMsg(msg)

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
	om.MsgType = OpsMsgTypeINFOMSG
	om.MsgClass = OpsMsgClassINFO

	if msgNo == 0 {
		om.msgId = 0
	} else {
		om.msgId = msgNo
		om.msgNumber = msgNo + om.MsgContext.BaseMessageId
	}

	om.SetTime("Local")
	om.setMsg(msg)

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
	om.MsgType = OpsMsgTypeERRORMSG
	om.MsgClass = OpsMsgClassFATAL

	if errNo == 0 {
		om.msgId = 0
	} else {
		om.msgId = errNo
		om.msgNumber = errNo + om.MsgContext.BaseMessageId
	}

	om.SetTime("Local")
	om.setMsg(errMsg)

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
func (opsMsg OpsMsgDto) NewStdErrorMsg(msg string, errNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.MsgType = OpsMsgTypeERRORMSG
	om.MsgClass = OpsMsgClassOPERROR

	if errNo == 0 {
		om.msgId = 0
		om.msgNumber = 0
	} else {
		om.msgId = errNo
		om.msgNumber = errNo + om.MsgContext.BaseMessageId
	}

	om.SetTime("Local")
	om.setMsg(msg)

	return om

}

// NewMsgFromSpecErrMsg - Create a new Operations Message based on
// the error information contained in a Type SpecErr passed
// into the method. The new message will be designated as
// an error message.
func (opsMsg *OpsMsgDto) NewMsgFromSpecErrMsg(se SpecErr) OpsMsgDto {

	om := OpsMsgDto{}

	if se.ErrorMsgType == SpecErrTypeFATAL {
		om.MsgType = OpsMsgTypeERRORMSG
		om.MsgClass = OpsMsgClassFATAL
		om.ErrDto = se

	} else if se.ErrorMsgType == SpecErrTypeERROR {
		om.MsgType = OpsMsgTypeERRORMSG
		om.MsgClass = OpsMsgClassOPERROR
		om.ErrDto = se

	} else if se.ErrorMsgType == SpecErrTypeWARNING {

		om.MsgType = OpsMsgTypeWARNINGMSG
		om.MsgClass = OpsMsgClassWARNING
		om.ErrDto = se

	} else if se.ErrorMsgType == SpecErrTypeINFO {

		om.MsgType = OpsMsgTypeINFOMSG
		om.MsgClass = OpsMsgClassINFO
		om.ErrDto = se

	} else if se.ErrorMsgType == SpecErrTypeNOERRORSALLCLEAR {
		om.MsgType = OpsMsgTypeNOERRORNOMSG
		om.MsgClass = OpsMsgClassNOERRORSNOMESSAGES
		om.ErrDto = se

	} else {
		om.MsgType = OpsMsgTypeINFOMSG
		om.MsgClass = OpsMsgClassINFO
		om.ErrDto = se
	}

	if se.ErrNo == 0 {
		om.msgId = 0
		om.msgNumber = 0
	} else {
		om.msgId = se.ErrNo
		om.msgNumber = se.ErrNo
	}

	om.MsgTimeUTC = se.ErrorMsgTimeUTC
	om.MsgTimeLocal = se.ErrorMsgTimeLocal
	om.MsgLocalTimeZone = se.ErrorLocalTimeZone
	om.Message = append(opsMsg.Message, se.Error())

	x := se.DeepCopyParentInfo(se.ParentInfo)

	for _, bi := range x {
		ci := OpsMsgContextInfo{SourceFileName:bi.SourceFileName, ParentObjectName: bi.ParentObjectName, FuncName: bi.FuncName, BaseMessageId: bi.BaseErrorID}
		om.ParentContextHistory = append(om.ParentContextHistory, ci)
	}

	y := se.DeepCopyBaseInfo()

	om.MsgContext = OpsMsgContextInfo{SourceFileName:y.SourceFileName, ParentObjectName: y.ParentObjectName, FuncName: y.FuncName, BaseMessageId: y.BaseErrorID}

	return om
}


// NewWarningMsg - Creates a new Warning Message
func (opsMsg OpsMsgDto) NewWarningMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.MsgType = OpsMsgTypeWARNINGMSG
	om.MsgClass = OpsMsgClassWARNING
	if msgNo == 0 {
		om.msgId = 0
		om.msgNumber = 0
	} else {
		om.msgId = msgNo
		om.msgNumber = msgNo + om.MsgContext.BaseMessageId
	}
	om.SetTime("Local")
	om.setMsg(msg)

	return om

}

// getMsgTitle - Returns the Message title, message number and the
// banner line separator based on value of OpsMsgDto.MsgClass
func (opsMsg *OpsMsgDto) getMsgTitle() (string, string, string) {
	var title string
	var banner string
	var msgNo	string

	i := int(opsMsg.MsgClass)

	msgPrefix := "Message Number: "

	switch i {

	case 0:
		// OpsMsgClassDEBUG - 0 Message is a Debug Message
		title = "DEBUG Message"
		banner = strings.Repeat("*", 75)
	case 1:
		// OpsMsgClassOPERROR - 1 Message is an Error Message
		title = "Standard ERROR Message"
		msgPrefix = "Error Number: "
		banner = strings.Repeat("X", 75)
	case 2:
		// OpsMsgClassFATAL - 2 Message is a Fatal Error Message
		title = "FATAL ERROR Message"
		msgPrefix = "Error Number: "
		banner = strings.Repeat("!", 75)
	case 3:
		// OpsMsgClassINFO - 3 Message is an Informational Message
		title = "Information Message"
		banner = strings.Repeat("_", 75)
case 4:
		// OpsMsgClassWARNING - 4 Message is a warning Message
		title = "WARNING Message"
		banner = strings.Repeat("?", 75)
	default:
		title = "Message"
		banner = strings.Repeat("_", 75)
	}

	if opsMsg.msgId != 0 {
		msgNo = msgPrefix + ": " + string(opsMsg.msgNumber)
	} else {
		msgNo = ""
	}


	return title, msgNo, banner
}


func(opsMsg *OpsMsgDto) setDebugMsg(msg string) {
	banner := "\n" + strings.Repeat("+", 75)
	m := "\n\n"
	m += banner
	// FmtDateTimeTzNanoYMD
	dt := DateTimeUtility{}
	timeStamp := "Local Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeLocal)
	timeStamp += "     UTC Time: " + dt.GetDateTimeTzNanoSecText(opsMsg.MsgTimeUTC)
	m += "\nDEBUG Message: "
	m += timeStamp
	m += "\n" + msg
	m += banner

	opsMsg.Message = append(opsMsg.Message, m)
}

func(opsMsg *OpsMsgDto) setMsg(msg string) {

  title, banner, msgNo := opsMsg.getMsgTitle()

	m:= "\n\n"
	m += "\n" + banner
	m += "\n                " + title
	m += "\n" + banner
	m += "\n" + msg
	m += "\n" + banner

	if msgNo != "" {
		m += "\n" + msgNo
	}

	m += "\n Message Type: " + opsMsg.MsgType.String()
	m += "\nMessage Level: " + opsMsg.MsgClass.String()
	m += "\n" + banner
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
	m += "\n" + banner

	opsMsg.Message = append(opsMsg.Message, m)

}

// SetTime - Sets the time stamp for this Operations
// Message. Notice that the input parameter 'localTimeZone'
// is the Iana Time Zone for local time.
//
// Reference Iana Time Zones: https://www.iana.org/time-zones
//
// If the 'localTimeZone' parameter string is empty or an
// invalid time zone, local time zone will default to 'Local'.
// The 'Local' time zone is determined by the host computer.
func(opsMsg *OpsMsgDto)SetTime(localTimeZone string){

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