package common

import (
	"testing"
	"strings"
	"time"
)

func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewNoErrorsNoMessagesMsg("",msgId)

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewNoErrorsNoMessagesMsg("",msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_03(t *testing.T) {

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(28)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om := OpsMsgDto{}.NewNoErrorsNoMessagesMsg("",msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetNoErrorsNoMessages_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetNoErrorsNoMessages_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om.SetMessageContext(testMsgContext)
	om.SetNoErrorsNoMessages("",msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetNoErrorsNoMessages_03(t *testing.T) {

	om := OpsMsgDto{}

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(28)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om.SetNoErrorsNoMessages("",msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetNoErrorsNoMessages_04(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}

	xMsg := "Xray = 6"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

	om.SetMessageContext(testMsgContext)
	om.SetNoErrorsNoMessages(xMsg,msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgClass_String(t *testing.T) {
	om := testOpsMsgDtoCreateInfoMsg()

	actualMsg := om.String()

	if !strings.Contains(actualMsg,"This is Information Message for test object") {
		t.Errorf("Expected message to contain string 'This is Information Message for test object'. It did NOT! msg= '%v'",actualMsg)
	}
}


func TestOpsMsgDto_GetError_01(t *testing.T) {
	om := testOpsMsgDtoCreateStdErrorMsg()

	err := om.GetError()

	if err == nil {
		t.Error("Expected a valid 'error' type to be returned from standard error om.GetError(). It was 'nil'!")
	}

	actualMsg := err.Error()

	expectedMsg := "This is Standard Error Msg for test object"

	if !strings.Contains(actualMsg, expectedMsg) {
		t.Errorf("Expected error message returned by error type to include text, '%v'. Actual message was '%v'",expectedMsg, actualMsg)
	}

}

func TestOpsMsgDto_GetError_02(t *testing.T) {
	om := testOpsMsgDtoCreateFatalErrorMsg()

	err := om.GetError()

	if err == nil {
		t.Error("Expected a valid 'error' type to be returned from fatal error om.GetError(). It was 'nil'!")
	}

	actualMsg := err.Error()

	expectedMsg := "This is FATAL Error Msg for test object"

	if !strings.Contains(actualMsg, expectedMsg) {
		t.Errorf("Expected error message returned by error type to include text, '%v'. Actual message was '%v'",expectedMsg, actualMsg)
	}

}

func TestOpsMsgDto_GetError_03(t *testing.T) {
	om := testOpsMsgDtoCreateInfoMsg()

	err := om.GetError()

	if err != nil {
		t.Error("Expected 'nil' to be returned from information message om.GetError(). It was non nil 'nil'!")
	}

}

func TestOpsMsgDto_CopyIn_01(t *testing.T) {

	om1 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 := testOpsMsgDtoCreateInfoMsg()

	om1.CopyIn(&om2)

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
	}

	if om1.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om1.MsgClass)
	}

	if om1.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om1.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om1.IsFatalError())
	}

	mId := om1.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om1.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om1.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om1.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om1.MsgTimeUTC == Zero. om1.MsgTimeUTC== '%v'", om1.MsgTimeUTC)
	}

	if om1.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om1.MsgTimeLocal == Zero. om1.MsgTimeLocal== '%v'",om1.MsgTimeLocal)
	}

	if om1.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om1.MsgLocalTimeZone is NOT set to 'Local'. om1.MsgLocalTimeZone== '%v' ", om1.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("om1 should equal om2. It did NOT!")
	}

}

func TestOpsMsgDto_CopyIn_02(t *testing.T) {

	om1 := OpsMsgDto{}

	om2 := testOpsMsgDtoCreateInfoMsg()

	om1.CopyIn(&om2)

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
	}

	if om1.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om1.MsgClass)
	}

	if om1.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om1.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om1.IsFatalError())
	}

	mId := om1.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om1.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om1.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om1.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om1.MsgTimeUTC == Zero. om1.MsgTimeUTC== '%v'", om1.MsgTimeUTC)
	}

	if om1.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om1.MsgTimeLocal == Zero. om1.MsgTimeLocal== '%v'",om1.MsgTimeLocal)
	}

	if om1.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om1.MsgLocalTimeZone is NOT set to 'Local'. om1.MsgLocalTimeZone== '%v' ", om1.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("om1 should equal om2. It did NOT!")
	}

}

func TestOpsMsgDto_CopyIn_03(t *testing.T) {

	om2 := OpsMsgDto{}

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(19)

	om2.SetInfoMessage(xMsg, msgId)

	om1 := OpsMsgDto{}

	om1.CopyIn(&om2)


	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
	}

	if om1.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om1.MsgClass)
	}

	if om1.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om1.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om1.IsFatalError())
	}

	mId := om1.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om1.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om1.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om1.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om1.MsgTimeUTC == Zero. om1.MsgTimeUTC== '%v'", om1.MsgTimeUTC)
	}

	if om1.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om1.MsgTimeLocal == Zero. om1.MsgTimeLocal== '%v'",om1.MsgTimeLocal)
	}

	if om1.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om1.MsgLocalTimeZone is NOT set to 'Local'. om1.MsgLocalTimeZone== '%v' ", om1.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("om1 should equal om2. It did NOT!")
	}

}

func TestOpsMsgDto_CopyIn_04(t *testing.T) {

	om2 := OpsMsgDto{}
	ci := OpsMsgContextInfo{}
	om2.MsgContext = ci.New("TSource06", "PObj06", "Func006", 6000)



	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)

	om2.SetInfoMessage(xMsg, msgId)

	om1 := OpsMsgDto{}

	om1.CopyIn(&om2)


	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
	}

	if om1.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om1.MsgClass)
	}

	if om1.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om1.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om1.IsFatalError())
	}

	mId := om1.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om1.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om1.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om1.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om1.MsgTimeUTC == Zero. om1.MsgTimeUTC== '%v'", om1.MsgTimeUTC)
	}

	if om1.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om1.MsgTimeLocal == Zero. om1.MsgTimeLocal== '%v'",om1.MsgTimeLocal)
	}

	if om1.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om1.MsgLocalTimeZone is NOT set to 'Local'. om1.MsgLocalTimeZone== '%v' ", om1.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("om1 should equal om2. It did NOT!")
	}

}

func TestOpsMsgDto_CopyOut_01(t *testing.T) {
	om1 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 := om1.CopyOut()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassFATAL


	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
	}

	if om2.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om2.MsgClass)
	}

	if om2.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om2.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om2.IsFatalError())
	}

	mId := om2.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om2.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om2.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om2.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om2.MsgTimeUTC == Zero. om2.MsgTimeUTC== '%v'", om2.MsgTimeUTC)
	}

	if om2.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om2.MsgTimeLocal == Zero. om2.MsgTimeLocal== '%v'",om2.MsgTimeLocal)
	}

	if om2.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om2.MsgLocalTimeZone is NOT set to 'Local'. om2.MsgLocalTimeZone== '%v' ", om2.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("Expected om1==om2.  It did NOT!")
	}
}

func TestOpsMsgDto_CopyOut_02(t *testing.T) {
	om1 := testOpsMsgDtoCreateFatalErrorMsg()

	om2:= testOpsMsgDtoCreateInfoMsg()

	om2 = om1.CopyOut()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassFATAL


	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
	}

	if om2.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om2.MsgClass)
	}

	if om2.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om2.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om2.IsFatalError())
	}

	mId := om2.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om2.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om2.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om2.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om2.MsgTimeUTC == Zero. om2.MsgTimeUTC== '%v'", om2.MsgTimeUTC)
	}

	if om2.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om2.MsgTimeLocal == Zero. om2.MsgTimeLocal== '%v'",om2.MsgTimeLocal)
	}

	if om2.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om2.MsgLocalTimeZone is NOT set to 'Local'. om2.MsgLocalTimeZone== '%v' ", om2.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("Expected om1==om2.  It did NOT!")
	}
}

func TestOpsMsgDto_CopyOut_03(t *testing.T) {
	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(152)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassFATAL

	om1 := OpsMsgDto{}
	om1.SetFatalErrorMessage(xMsg, msgId)

	om2:= testOpsMsgDtoCreateInfoMsg()

	om2 = om1.CopyOut()



	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
	}

	if om2.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om2.MsgClass)
	}

	if om2.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om2.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om2.IsFatalError())
	}

	mId := om2.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om2.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om2.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om2.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om2.MsgTimeUTC == Zero. om2.MsgTimeUTC== '%v'", om2.MsgTimeUTC)
	}

	if om2.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om2.MsgTimeLocal == Zero. om2.MsgTimeLocal== '%v'",om2.MsgTimeLocal)
	}

	if om2.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om2.MsgLocalTimeZone is NOT set to 'Local'. om2.MsgLocalTimeZone== '%v' ", om2.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("Expected om1==om2.  It did NOT!")
	}
}

func TestOpsMsgDto_CopyOut_04(t *testing.T) {

	om2 := OpsMsgDto{}
	ci := OpsMsgContextInfo{}
	om2.MsgContext = ci.New("TSource06", "PObj06", "Func006", 6000)



	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)

	om2.SetInfoMessage(xMsg, msgId)

	om1 := testOpsMsgDtoCreateFatalErrorMsg()

	om1 = om2.CopyOut()


	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
	}

	if om1.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om1.MsgClass)
	}

	if om1.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om1.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om1.IsFatalError())
	}

	mId := om1.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om1.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om1.GetMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om1.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om1.MsgTimeUTC == Zero. om1.MsgTimeUTC== '%v'", om1.MsgTimeUTC)
	}

	if om1.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om1.MsgTimeLocal == Zero. om1.MsgTimeLocal== '%v'",om1.MsgTimeLocal)
	}

	if om1.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om1.MsgLocalTimeZone is NOT set to 'Local'. om1.MsgLocalTimeZone== '%v' ", om1.MsgLocalTimeZone)
	}

	if !om2.Equal(&om1) {
		t.Error("om1 should equal om2. It did NOT!")
	}

}

func TestOpsMsgContextInfo_Equal_01(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()


	om2 = om1.CopyOut()

	if !om2.Equal(&om1) {
		t.Error("Expected om2==om1. It did NOT!")
	}

}

func TestOpsMsgContextInfo_Equal_02(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()


	om2 = om1.CopyOut()

	om2.MsgContext.FuncName = "..."

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR! ")
	}

}

func TestOpsMsgContextInfo_Equal_03(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()


	om2 = om1.CopyOut()

	om2.ParentContextHistory[4].BaseMessageId = 99

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR! ")
	}

}

func TestOpsMsgContextInfo_Equal_04(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 = om1.CopyOut()

	om2.MsgTimeLocal = time.Now()

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgContextInfo_Equal_05(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 = om1.CopyOut()

	om2.FmtMessage = "xxxx"

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgContextInfo_Equal_06(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 = om1.CopyOut()

	om2.MsgClass = OpsMsgClassNOERRORSNOMESSAGES

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgContextInfo_Equal_07(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 = om1.CopyOut()

	om2.MsgType = OpsMsgTypeDEBUGMSG

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgContextInfo_Equal_08(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 = om1.CopyOut()

	om2.Message = "..."

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgDto_IsDebugMsg_01(t *testing.T) {

	om := testOpsMsgDtoCreateDebugMsg()

	expectedIsDebug := true

	actualIsDebug := om.IsDebugMsg()

	if expectedIsDebug != actualIsDebug {
		t.Errorf("Expected om.IsError() = '%v'. Actual om.IsError()= '%v'", expectedIsDebug, actualIsDebug)
	}

}

func TestOpsMsgDto_IsDebugMsg_02(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	expectedIsDebug := false

	actualIsDebug := om.IsDebugMsg()

	if expectedIsDebug != actualIsDebug {
		t.Errorf("Expected om.IsDebugMsg() = '%v'. Actual om.IsDebugMsg()= '%v'", expectedIsDebug, actualIsDebug)
	}

}

func TestOpsMsgDto_IsError_01(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	expectedIsError := true

	actualIsError := om.IsError()

	if expectedIsError != actualIsError {
		t.Errorf("Expected om.IsError() = '%v'. Actual om.IsError()= '%v'", expectedIsError, actualIsError)
	}

}

func TestOpsMsgDto_IsError_02(t *testing.T) {

	om := testOpsMsgDtoCreateStdErrorMsg()

	expectedIsError := true

	actualIsError := om.IsError()

	if expectedIsError != actualIsError {
		t.Errorf("Expected om.IsError() = '%v'. Actual om.IsError()= '%v'", expectedIsError, actualIsError)
	}

}


func TestOpsMsgDto_IsError_03(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsError := false

	actualIsError := om.IsError()

	if expectedIsError != actualIsError {
		t.Errorf("Expected om.IsError() = '%v'. Actual om.IsError()= '%v'", expectedIsError, actualIsError)
	}

}

func TestOpsMsgDto_IsFatalError_01(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	expectedIsFatalError := true

	actualIsFatalError := om.IsFatalError()

	if expectedIsFatalError != actualIsFatalError {
		t.Errorf("Expected om.IsFatalError() = '%v'. Actual om.IsFatalError()= '%v'", expectedIsFatalError, actualIsFatalError)
	}

}

func TestOpsMsgDto_IsFatalError_02(t *testing.T) {

	om := testOpsMsgDtoCreateStdErrorMsg()

	expectedIsFatalError := false

	actualIsFatalError := om.IsFatalError()

	if expectedIsFatalError != actualIsFatalError {
		t.Errorf("Expected om.IsFatalError() = '%v'. Actual om.IsFatalError()= '%v'", expectedIsFatalError, actualIsFatalError)
	}

}


func TestOpsMsgDto_IsFatalError_03(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsFatalError := false

	actualIsFatalError := om.IsFatalError()

	if expectedIsFatalError != actualIsFatalError {
		t.Errorf("Expected om.IsFatalError() = '%v'. Actual om.IsFatalError()= '%v'", expectedIsFatalError, actualIsFatalError)
	}

}