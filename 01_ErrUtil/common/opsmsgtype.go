package common

/*  'opsmsgtype.go' is located in source code
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
