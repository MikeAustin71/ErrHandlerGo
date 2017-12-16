package common

import (
	"testing"
	"strings"
)


func TestOpsMsgDto_IsInfoMsg_01(t *testing.T) {
	
	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsInfoMsg := true

	actualIsInfoMsg := om.IsInfoMsg()

	if expectedIsInfoMsg != actualIsInfoMsg {
		t.Errorf("Expected om.IsInfoMsg() = '%v'. Actual om.IsInfoMsg()= '%v'", expectedIsInfoMsg, actualIsInfoMsg)
	}
	
}

func TestOpsMsgDto_IsInfoMsg_02(t *testing.T) {

	om := testOpsMsgDtoCreateStdErrorMsg()

	expectedIsInfoMsg := false

	actualIsInfoMsg := om.IsInfoMsg()

	if expectedIsInfoMsg != actualIsInfoMsg {
		t.Errorf("Expected om.IsInfoMsg() = '%v'. Actual om.IsInfoMsg()= '%v'", expectedIsInfoMsg, actualIsInfoMsg)
	}

}

func TestOpsMsgDto_IsInfoMsg_03(t *testing.T) {

	om := testOpsMsgDtoCreateWarningMsg()

	expectedIsInfoMsg := false

	actualIsInfoMsg := om.IsInfoMsg()

	if expectedIsInfoMsg != actualIsInfoMsg {
		t.Errorf("Expected om.IsInfoMsg() = '%v'. Actual om.IsInfoMsg()= '%v'", expectedIsInfoMsg, actualIsInfoMsg)
	}

}

func TestOpsMsgDto_IsInfoMsg_04(t *testing.T) {

	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	expectedIsInfoMsg := false

	actualIsInfoMsg := om.IsInfoMsg()

	if expectedIsInfoMsg != actualIsInfoMsg {
		t.Errorf("Expected om.IsInfoMsg() = '%v'. Actual om.IsInfoMsg()= '%v'", expectedIsInfoMsg, actualIsInfoMsg)
	}

}

func TestOpsMsgDto_IsNoErrorsNoMessages_01(t *testing.T) {
	
	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	expectedIsNoErrorsNoMessages := true

	actualIsNoErrorsNoMessages := om.IsNoErrorsNoMessages()

	if expectedIsNoErrorsNoMessages != actualIsNoErrorsNoMessages {
		t.Errorf("Expected om.IsNoErrorsNoMessages() = '%v'. Actual om.IsNoErrorsNoMessages()= '%v'", expectedIsNoErrorsNoMessages, actualIsNoErrorsNoMessages)
	}
	
}

func TestOpsMsgDto_IsNoErrorsNoMessages_02(t *testing.T) {

	om := testOpsMsgDtoCreateStdErrorMsg()

	expectedIsNoErrorsNoMessages := false

	actualIsNoErrorsNoMessages := om.IsNoErrorsNoMessages()

	if expectedIsNoErrorsNoMessages != actualIsNoErrorsNoMessages {
		t.Errorf("Expected om.IsNoErrorsNoMessages() = '%v'. Actual om.IsNoErrorsNoMessages()= '%v'", expectedIsNoErrorsNoMessages, actualIsNoErrorsNoMessages)
	}

}

func TestOpsMsgDto_IsNoErrorsNoMessages_03(t *testing.T) {

	om := testOpsMsgDtoCreateSuccessfulCompletionMsg()

	expectedIsNoErrorsNoMessages := false

	actualIsNoErrorsNoMessages := om.IsNoErrorsNoMessages()

	if expectedIsNoErrorsNoMessages != actualIsNoErrorsNoMessages {
		t.Errorf("Expected om.IsNoErrorsNoMessages() = '%v'. Actual om.IsNoErrorsNoMessages()= '%v'", expectedIsNoErrorsNoMessages, actualIsNoErrorsNoMessages)
	}

}

func TestOpsMsgDto_IsNoErrorsNoMessages_04(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsNoErrorsNoMessages := false

	actualIsNoErrorsNoMessages := om.IsNoErrorsNoMessages()

	if expectedIsNoErrorsNoMessages != actualIsNoErrorsNoMessages {
		t.Errorf("Expected om.IsNoErrorsNoMessages() = '%v'. Actual om.IsNoErrorsNoMessages()= '%v'", expectedIsNoErrorsNoMessages, actualIsNoErrorsNoMessages)
	}

}

func TestOpsMsgDto_IsSuccessfulCompletionMsg_01(t *testing.T) {
	
	om := testOpsMsgDtoCreateSuccessfulCompletionMsg()

	expectedIsSuccessfulCompletionMsg := true

	actualIsSuccessfulCompletionMsg := om.IsSuccessfulCompletionMsg()

	if expectedIsSuccessfulCompletionMsg != actualIsSuccessfulCompletionMsg {
		t.Errorf("Expected om.IsSuccessfulCompletionMsg() = '%v'. Actual om.IsSuccessfulCompletionMsg()= '%v'", expectedIsSuccessfulCompletionMsg, actualIsSuccessfulCompletionMsg)
	}
	
}

func TestOpsMsgDto_IsSuccessfulCompletionMsg_02(t *testing.T) {

	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	expectedIsSuccessfulCompletionMsg := false

	actualIsSuccessfulCompletionMsg := om.IsSuccessfulCompletionMsg()

	if expectedIsSuccessfulCompletionMsg != actualIsSuccessfulCompletionMsg {
		t.Errorf("Expected om.IsSuccessfulCompletionMsg() = '%v'. Actual om.IsSuccessfulCompletionMsg()= '%v'", expectedIsSuccessfulCompletionMsg, actualIsSuccessfulCompletionMsg)
	}

}

func TestOpsMsgDto_IsSuccessfulCompletionMsg_03(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	expectedIsSuccessfulCompletionMsg := false

	actualIsSuccessfulCompletionMsg := om.IsSuccessfulCompletionMsg()

	if expectedIsSuccessfulCompletionMsg != actualIsSuccessfulCompletionMsg {
		t.Errorf("Expected om.IsSuccessfulCompletionMsg() = '%v'. Actual om.IsSuccessfulCompletionMsg()= '%v'", expectedIsSuccessfulCompletionMsg, actualIsSuccessfulCompletionMsg)
	}

}

func TestOpsMsgDto_IsSuccessfulCompletionMsg_04(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsSuccessfulCompletionMsg := false

	actualIsSuccessfulCompletionMsg := om.IsSuccessfulCompletionMsg()

	if expectedIsSuccessfulCompletionMsg != actualIsSuccessfulCompletionMsg {
		t.Errorf("Expected om.IsSuccessfulCompletionMsg() = '%v'. Actual om.IsSuccessfulCompletionMsg()= '%v'", expectedIsSuccessfulCompletionMsg, actualIsSuccessfulCompletionMsg)
	}

}

func TestOpsMsgDto_IsWarningMsg_01(t *testing.T) {
	
	om := testOpsMsgDtoCreateWarningMsg()

	expectedIsWarningMsg := true

	actualIsWarningMsg := om.IsWarningMsg()

	if expectedIsWarningMsg != actualIsWarningMsg {
		t.Errorf("Expected om.IsWarningMsg() = '%v'. Actual om.IsWarningMsg()= '%v'", expectedIsWarningMsg, actualIsWarningMsg)
	}
	
}

func TestOpsMsgDto_IsWarningMsg_02(t *testing.T) {

	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	expectedIsWarningMsg := false

	actualIsWarningMsg := om.IsWarningMsg()

	if expectedIsWarningMsg != actualIsWarningMsg {
		t.Errorf("Expected om.IsWarningMsg() = '%v'. Actual om.IsWarningMsg()= '%v'", expectedIsWarningMsg, actualIsWarningMsg)
	}

}

func TestOpsMsgDto_IsWarningMsg_03(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	expectedIsWarningMsg := false

	actualIsWarningMsg := om.IsWarningMsg()

	if expectedIsWarningMsg != actualIsWarningMsg {
		t.Errorf("Expected om.IsWarningMsg() = '%v'. Actual om.IsWarningMsg()= '%v'", expectedIsWarningMsg, actualIsWarningMsg)
	}

}

func TestOpsMsgDto_IsWarningMsg_04(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsWarningMsg := false

	actualIsWarningMsg := om.IsWarningMsg()

	if expectedIsWarningMsg != actualIsWarningMsg {
		t.Errorf("Expected om.IsWarningMsg() = '%v'. Actual om.IsWarningMsg()= '%v'", expectedIsWarningMsg, actualIsWarningMsg)
	}

}

func TestOpsMsgDto_IsWarningMsg_05(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	expectedIsWarningMsg := false

	actualIsWarningMsg := om.IsWarningMsg()

	if expectedIsWarningMsg != actualIsWarningMsg {
		t.Errorf("Expected om.IsWarningMsg() = '%v'. Actual om.IsWarningMsg()= '%v'", expectedIsWarningMsg, actualIsWarningMsg)
	}

}

func TestOpsMsgDto_SignalSuccessfulCompletion_01(t *testing.T) {

	omx := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	om := omx.SignalSuccessfulCompletion(64)

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "Successful Completion"
	msgId := int64(64)
	msgNo := int64(6064)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION
	msgClass := OpsMsgClassSUCCESSFULCOMPLETION

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
		t.Error("Expected Successful Completion Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Successful Completion Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

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
func TestOpsMsgDto_SignalSuccessfulCompletion_02(t *testing.T) {

	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj()).SignalSuccessfulCompletion(0)

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "Successful Completion"
	msgId := int64(0)
	msgNo := int64(0)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION
	msgClass := OpsMsgClassSUCCESSFULCOMPLETION

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
		t.Error("Expected Successful Completion Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Successful Completion Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

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
