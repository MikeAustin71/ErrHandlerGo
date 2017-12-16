package common

import (
	"testing"
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


