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

	// OpsNOERRORNOMSGTYPE
	OpsNOERRORNOMSGTYPE OpsMsgType = iota

	// OpsERRORMSGTYPE - Message Type
	OpsERRORMSGTYPE

	// OpsINFOMSGTYPE - Information Message Type
	OpsINFOMSGTYPE

	// OpsWARNINGMSGTYPE - Warning Message Type
	OpsWARNINGMSGTYPE

	// OpsDEBUGMSGTYPE - Debug Message
	OpsDEBUGMSGTYPE

)

// OpsMsgTypeNames - String Array holding Message Type names.
var OpsMsgTypeNames = [...]string{"NOERRORSNOMSGS","ERROR", "INFO", "WARNING", "DEBUG"}
