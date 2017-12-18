package common

import (
	"testing"
	"strings"
	"errors"
)


func TestOpsMsgDto_InitializeContextInfo_01(t *testing.T) {
	parentInfo := testOpsMsgDtoCreateParentHistory()
	contextInfo := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}.InitializeAllContextInfo(parentInfo, contextInfo)

	l := len(om.ParentContextHistory)

	if l != 5 {
		t.Errorf("Expected Parent Context History Length = 5. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om.ParentContextHistory[0].SourceFileName != "TSource01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource01'. Instead, SourceFileName= '%v'",om.ParentContextHistory[0].SourceFileName)
	}

	if om.ParentContextHistory[0].ParentObjectName != "PObj01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj01'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[0].ParentObjectName)
	}

	if om.ParentContextHistory[0].FuncName != "Func001" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History FuncName= 'Func001'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[0].FuncName)
	}

	if om.ParentContextHistory[0].BaseMessageId != int64(1000) {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History BaseMessageId = 1000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[0].BaseMessageId)
	}

	if om.ParentContextHistory[1].SourceFileName != "TSource02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource02'. Instead, SourceFileName= '%v'",om.ParentContextHistory[1].SourceFileName)
	}


	if om.ParentContextHistory[1].ParentObjectName != "PObj02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj02'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[1].ParentObjectName)
	}

	if om.ParentContextHistory[1].FuncName != "Func002" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History FuncName= 'Func002'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[1].FuncName)
	}

	if om.ParentContextHistory[1].BaseMessageId != int64(2000) {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History BaseMessageId = 2000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[1].BaseMessageId)
	}

	if om.ParentContextHistory[2].SourceFileName != "TSource03" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource03'. Instead, SourceFileName= '%v'",om.ParentContextHistory[2].SourceFileName)
	}

	if om.ParentContextHistory[2].ParentObjectName != "PObj03" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj03'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[2].ParentObjectName)
	}

	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

	if om.ParentContextHistory[2].BaseMessageId != int64(3000) {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History BaseMessageId = 3000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[2].BaseMessageId)
	}


	if om.ParentContextHistory[3].SourceFileName != "TSource04" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource04'. Instead, SourceFileName= '%v'",om.ParentContextHistory[3].SourceFileName)
	}

	if om.ParentContextHistory[3].ParentObjectName != "PObj04" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj04'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[3].ParentObjectName)
	}

	if om.ParentContextHistory[3].FuncName != "Func004" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History FuncName= 'Func004'. Instead, FuncName= '%v'",om.ParentContextHistory[3].FuncName)
	}

	if om.ParentContextHistory[3].BaseMessageId != int64(4000) {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History BaseMessageId = 4000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[3].BaseMessageId)
	}

	if om.ParentContextHistory[4].SourceFileName != "TSource05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource05'. Instead, SourceFileName= '%v'",om.ParentContextHistory[4].SourceFileName)
	}

	if om.ParentContextHistory[4].ParentObjectName != "PObj05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History ParentObjectName = 'PObj05'. Instead, ParentObjectName = '%v'", om.ParentContextHistory[4].ParentObjectName)
	}

	if om.ParentContextHistory[4].FuncName != "Func005" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History FuncName= 'Func005'. Instead, FuncName= '%v'",om.ParentContextHistory[4].FuncName)
	}

	if om.ParentContextHistory[4].BaseMessageId != int64(5000) {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History BaseMessageId = 5000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[4].BaseMessageId)
	}

	if om.MsgContext.SourceFileName != "TSource06" {
		t.Errorf("Expected MsgContext.SourceFileName == 'TSource06'. Instead, SourceFileName== '%v'", om.MsgContext.SourceFileName)
	}

	if om.MsgContext.ParentObjectName != "PObj06" {
		t.Errorf("Expected MsgContext.ParentObjectName == 'PObj06'. Instead, ParentObjectName== '%v'", om.MsgContext.ParentObjectName)
	}

	if om.MsgContext.FuncName != "Func006" {
		t.Errorf("Expected MsgContext.FuncName == 'Func006'. Instead, FuncName== '%v'", om.MsgContext.FuncName)
	}

	if om.MsgContext.BaseMessageId != 6000 {
		t.Errorf("Expected MsgContext.BaseMessageId == '6000'. Instead, BaseMessageId== '%v'", om.MsgContext.BaseMessageId)
	}
}



func TestOpsMsgDto_InitializeContextWithParentOpsMsg_01(t *testing.T) {
	parentInfo := testOpsMsgDtoCreateParentHistory()
	contextInfo := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}.InitializeAllContextInfo(parentInfo, contextInfo)
	newMsg := "Information Message Text 2"
	om.SetInfoMessage("Information Text 1", 122)


	ci := OpsMsgContextInfo{SourceFileName:"TSource07", ParentObjectName:"PObj07", FuncName: "Func007", BaseMessageId: 7000}

	om2 := OpsMsgDto{}.InitializeContextWithOpsMsgDto(om, ci)
	om2.SetInfoMessage(newMsg, 122 )


	l := len(om2.ParentContextHistory)

	if l != 6 {
		t.Errorf("Expected Parent Context History Length = 6. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om2.ParentContextHistory[0].SourceFileName != "TSource01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource01'. Instead, SourceFileName= '%v'",om2.ParentContextHistory[0].SourceFileName)
	}

	if om2.ParentContextHistory[0].ParentObjectName != "PObj01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj01'. Instead, ParentObjectName= '%v'",om2.ParentContextHistory[0].ParentObjectName)
	}


	if om2.ParentContextHistory[1].ParentObjectName != "PObj02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj02'. Instead, ParentObjectName= '%v'",om2.ParentContextHistory[1].ParentObjectName)
	}


	if om2.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om2.ParentContextHistory[2].FuncName)
	}

	if om2.ParentContextHistory[3].BaseMessageId != int64(4000) {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History BaseMessageId = 4000. Instead, BaseMessageId = '%v'", om2.ParentContextHistory[3].BaseMessageId)
	}

	if om2.ParentContextHistory[4].ParentObjectName != "PObj05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History ParentObjectName = 'PObj05'. Instead, ParentObjectName = '%v'", om2.ParentContextHistory[4].ParentObjectName)
	}

	if om2.ParentContextHistory[5].ParentObjectName != "PObj06" {
		t.Errorf("Expected 6th OpsMsgContextInfo object in Parent Context History ParentObjectName = 'PObj06'. Instead, ParentObjectName = '%v'", om2.ParentContextHistory[5].ParentObjectName)
	}

	if om2.MsgContext.SourceFileName != "TSource07" {
		t.Errorf("Expected MsgContext.SourceFileName == 'TSource07'. Instead, SourceFileName== '%v'", om2.MsgContext.SourceFileName)
	}

	if om2.MsgContext.ParentObjectName != "PObj07" {
		t.Errorf("Expected MsgContext.ParentObjectName == 'PObj07'. Instead, ParentObjectName== '%v'", om2.MsgContext.ParentObjectName)
	}

	if om2.MsgContext.FuncName != "Func007" {
		t.Errorf("Expected MsgContext.FuncName == 'Func007'. Instead, FuncName== '%v'", om2.MsgContext.FuncName)
	}

	if om2.MsgContext.BaseMessageId != 7000 {
		t.Errorf("Expected MsgContext.BaseMessageId == '7000'. Instead, BaseMessageId== '%v'", om2.MsgContext.BaseMessageId)
	}

	actualMsg := om2.GetFmtMessage()

	if !strings.Contains(actualMsg, newMsg) {
		t.Errorf("Expected actual message to contain string '%v'. It did NOT! actualMsg='%v'", newMsg, actualMsg)
	}

	msgNo := om2.GetMessageNumber()

	if msgNo != 7122 {
		t.Errorf("Expected om2.msgNumber == '7122'. Instead, om2.msgNumber == '%v'", msgNo)
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

func TestOpsMsgDto_NewDebugMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(6238)
	msgType := OpsMsgTypeDEBUGMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewDebugMsg(xMsg, msgId)

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

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewDebugMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(6238)
	msgType := OpsMsgTypeDEBUGMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewDebugMsg(xMsg, msgId)


	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewDebugMsg_03(t *testing.T) {

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(238)
	msgType := OpsMsgTypeDEBUGMSG

	om := OpsMsgDto{}.NewDebugMsg(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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


func TestOpsMsgDto_NewError_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is FATAL Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

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

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewError_02(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

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

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Standard Error Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewError_03(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Standard Error Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewError_04(t *testing.T) {

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.NewError(prefixMsgText, err, msgType , msgId)


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Standard Error Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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
