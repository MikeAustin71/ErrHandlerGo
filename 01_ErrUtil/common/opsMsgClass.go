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
