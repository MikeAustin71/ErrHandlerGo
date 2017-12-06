package common

/*  'opsMsgClass.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

 */

// OpsMsgClass - Holds the Message level indicating the relative importance of a specific log Message.
type OpsMsgClass int

// String - Returns the name of the OpsMsgClass element
// formatted as a string.
func (opsmsgclass OpsMsgClass) String() string {
	return OpsMsgClassNames[opsmsgclass]
}

const (
	// MsgClassDEBUG - 0 Message is a Debug Message
	MsgClassDEBUG OpsMsgClass = iota
	// MsgClassOPERROR - 1 Message is an Error Message
	MsgClassOPERROR
	// MsgClassFATAL - 2 Message is a Fatal Error Message
	MsgClassFATAL
	// MsgClassINFO - 3 Message is an Informational Message
	MsgClassINFO
	// MsgClassWARNING - 4 Message is a warning Message
	MsgClassWARNING
)

// OpsMsgClassNames - string array containing names of Log Levels
var OpsMsgClassNames = [...]string{"DEBUG", "OPERROR", "FATAL", "INFO", "WARNING"}
