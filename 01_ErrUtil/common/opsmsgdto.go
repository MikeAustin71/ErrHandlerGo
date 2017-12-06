package common

import "time"

/*  'opsmsgdto.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/



// OpsMsgDto - Data Transfer Object
// containing information about an
// operations Message
type OpsMsgDto struct {
	Message       []string
	MsgType       OpsMsgType
	MsgLevel      OpsMsgClass
	MsgTimeUTC    time.Time
	MsgTimeLocal  time.Time
	LocalTimeZone string
	ErrDto        SpecErr
}

// NewSpecErr - Create a new Operations Message based on
// the error information contained in a Type SpecErr passed
// into the method. The new message will be designated as
// an error message.
func (opsMsg *OpsMsgDto) NewSpecErr(se SpecErr) OpsMsgDto {

	om := OpsMsgDto{}

	if se.IsPanic {
		om.MsgType = OpsERRORMSGTYPE
		om.MsgLevel = MsgClassFATAL
		opsMsg.ErrDto = se

	} else if se.IsErr {
		om.MsgType = OpsERRORMSGTYPE
		om.MsgLevel = MsgClassDEBUG
		opsMsg.ErrDto = se

	} else {
		om.MsgType = OpsINFOMSGTYPE
		om.MsgLevel = MsgClassINFO
		opsMsg.ErrDto = se
	}

	opsMsg.Message = append(opsMsg.Message, se.Error())

	return om
}

// NewInfoMsg - Create a new Operations Message which is
// an Informational Message.
func(opsMsg *OpsMsgDto) NewInfoMsg(msg string) OpsMsgDto {

	om := OpsMsgDto{}
	om.MsgType = OpsINFOMSGTYPE
	om.MsgLevel = MsgClassINFO

	om.Message = append(om.Message, msg)

	return om
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
	opsMsg.LocalTimeZone = localTimeZone

	tzLocal, _ := tz.ConvertTz(opsMsg.MsgTimeUTC, opsMsg.LocalTimeZone)
	opsMsg.MsgTimeLocal = tzLocal.TimeOut

}