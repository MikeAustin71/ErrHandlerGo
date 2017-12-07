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



// OpsMsgDto - Data Transfer Object
// containing information about an
// operations Message
type OpsMsgDto struct {
	Message          	[]string
	MsgNumber					int64
	MsgType          	OpsMsgType
	MsgClass         	OpsMsgClass
	MsgTimeUTC       	time.Time
	MsgTimeLocal     	time.Time
	MsgLocalTimeZone 	string
	ErrDto           	SpecErr
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
	om.MsgType = OpsDEBUGMSGTYPE
	om.MsgClass = MsgClassDEBUG

	if msgNo == 0 {
		om.MsgNumber = 0
	} else {
		om.MsgNumber = msgNo
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
	om.MsgType = OpsINFOMSGTYPE
	om.MsgClass = MsgClassINFO

	if msgNo == 0 {
		om.MsgNumber = 0
	} else {
		om.MsgNumber = msgNo
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
	om.MsgType = OpsERRORMSGTYPE
	om.MsgClass = MsgClassFATAL

	if errNo == 0 {
		om.MsgNumber = 0
	} else {
		om.MsgNumber = errNo
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
	om.MsgType = OpsERRORMSGTYPE
	om.MsgClass = MsgClassOPERROR

	if errNo == 0 {
		om.MsgNumber = 0
	} else {
		om.MsgNumber = errNo
	}

	om.SetTime("Local")
	om.setMsg(msg)

	return om

}

// NewSpecErrMsg - Create a new Operations Message based on
// the error information contained in a Type SpecErr passed
// into the method. The new message will be designated as
// an error message.
func (opsMsg *OpsMsgDto) NewSpecErrMsg(se SpecErr) OpsMsgDto {

	om := OpsMsgDto{}

	if se.IsPanic {
		om.MsgType = OpsERRORMSGTYPE
		om.MsgClass = MsgClassFATAL
		opsMsg.ErrDto = se

	} else if se.IsErr {
		om.MsgType = OpsERRORMSGTYPE
		om.MsgClass = MsgClassOPERROR
		opsMsg.ErrDto = se

	} else {
		om.MsgType = OpsINFOMSGTYPE
		om.MsgClass = MsgClassINFO
		opsMsg.ErrDto = se
	}

	opsMsg.MsgNumber = se.ErrNo
	opsMsg.MsgTimeUTC = se.ErrorMsgTimeUTC
	opsMsg.MsgTimeLocal = se.ErrorMsgTimeLocal
	opsMsg.MsgLocalTimeZone = se.ErrorLocalTimeZone
	opsMsg.Message = append(opsMsg.Message, se.Error())

	return om
}


// NewWarningMsg - Creates a new Warning Message
func (opsMsg OpsMsgDto) NewWarningMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.MsgType = OpsWARNINGMSGTYPE
	om.MsgClass = MsgClassWARNING
	om.MsgNumber = msgNo
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
		// MsgClassDEBUG - 0 Message is a Debug Message
		title = "DEBUG Message"
		banner = strings.Repeat("*", 75)
	case 1:
		// MsgClassOPERROR - 1 Message is an Error Message
		title = "Standard ERROR Message"
		msgPrefix = "Error Number: "
		banner = strings.Repeat("X", 75)
	case 2:
		// MsgClassFATAL - 2 Message is a Fatal Error Message
		title = "FATAL ERROR Message"
		msgPrefix = "Error Number: "
		banner = strings.Repeat("!", 75)
	case 3:
		// MsgClassINFO - 3 Message is an Informational Message
		title = "Information Message"
		banner = strings.Repeat("_", 75)
case 4:
		// MsgClassWARNING - 4 Message is a warning Message
		title = "WARNING Message"
		banner = strings.Repeat("?", 75)
	default:
		title = "Message"
		banner = strings.Repeat("_", 75)
	}

	if opsMsg.MsgNumber != 0 {
		msgNo = msgPrefix + ": " + string(opsMsg.MsgNumber)
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