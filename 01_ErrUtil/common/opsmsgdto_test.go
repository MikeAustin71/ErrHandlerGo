package common

import (
	"testing"
	"strings"
	"time"
)

func TestOpsMsgDto_ParentHistory_01(t *testing.T) {

	ci := OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj02", "Func003", 3000)

	om := OpsMsgDto{}
	om.ParentContextHistory = append(om.ParentContextHistory, x1)
	om.ParentContextHistory = append(om.ParentContextHistory, x2)
	om.ParentContextHistory = append(om.ParentContextHistory, x3)

	l := len(om.ParentContextHistory)

	if l != 3 {
		t.Errorf("Expected Parent Context History Length = 3. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

}

func TestOpsMsgDto_ParentHistory_02(t *testing.T) {

	parent := testOpsMsgDtoCreateParentHistory()

	om := OpsMsgDto{}
	om.AddParentContextHistory(parent)

	l := len(om.ParentContextHistory)

	if l != 5 {
		t.Errorf("Expected Parent Context History Length = 5. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

	if om.ParentContextHistory[3].BaseMessageId != int64(4000) {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History BaseMessageId = 4000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[3].BaseMessageId)
	}

}

func TestOpsMsgDto_InitializeContextInfo_01(t *testing.T) {
	parentInfo := testOpsMsgDtoCreateParentHistory()
	contextInfo := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}.InitializeContextInfo(parentInfo, contextInfo)

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


	if om.ParentContextHistory[1].ParentObjectName != "PObj02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj02'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[1].ParentObjectName)
	}


	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

	if om.ParentContextHistory[3].BaseMessageId != int64(4000) {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History BaseMessageId = 4000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[3].BaseMessageId)
	}

	if om.ParentContextHistory[4].ParentObjectName != "PObj05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History ParentObjectName = 'PObj05'. Instead, ParentObjectName = '%v'", om.ParentContextHistory[4].ParentObjectName)
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

	om := OpsMsgDto{}.InitializeContextInfo(parentInfo, contextInfo)
	newMsg := "Information Message Text 2"
	om.SetInfoMessage("Information Text 1", 122)


	ci := OpsMsgContextInfo{SourceFileName:"TSource07", ParentObjectName:"PObj07", FuncName: "Func007", BaseMessageId: 7000}
	
	om2 := OpsMsgDto{}.InitializeContextWithParentOpsMsg(om, ci)
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

	actualMsg := om2.GetMessage()

	if !strings.Contains(actualMsg, newMsg) {
		t.Errorf("Expected actual message to contain string '%v'. It did NOT! actualMsg='%v'", newMsg, actualMsg)
	}

	msgNo := om2.GetMessageNumber()

	if msgNo != 7122 {
		t.Errorf("Expected om2.msgNumber == '7122'. Instead, om2.msgNumber == '%v'", msgNo)
	}

}

func TestOpsMsgDto_SetStdErrorMessage_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	om := testOpsMsgDtoCreateStdErrorMsg()

	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassOPERROR

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



	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}


	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
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

func TestOpsMsgDto_SetStdErrorMessage_02(t *testing.T) {
	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassOPERROR
	baseMsgId := int64(6000)
	srcFile := "TSource06"

	om := OpsMsgDto{}
	mCtx := testOpsMsgDtoCreateContextInfoObj()
	om.SetMessageContext(mCtx)
	om.SetStdErrorMessage(xMsg, msgId)

	if om.MsgContext.BaseMessageId != baseMsgId {
		t.Errorf("Expected om.MsgContext.BaseMessageId == '%v'. Instead, om.MsgContext.BaseMessageId == '%v'", baseMsgId, om.MsgContext.BaseMessageId )
	}

	if om.MsgContext.SourceFileName != srcFile{
		t.Errorf("Expected om.MsgContext.SourceFileName == '%v'. Instead, om.MsgContext.BaseMessageId == '%v'", srcFile, om.MsgContext.SourceFileName )
	}

	if !om.MsgContext.Equal(&mCtx) {
		t.Error("Expected om.MsgContext==mCtx. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}


	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
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

func TestOpsMsgDto_SetStdErrorMessage_03(t *testing.T) {
	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(429)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassOPERROR

	om := OpsMsgDto{}
	om.SetStdErrorMessage(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}


	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
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

func TestOpsMsgDto_SetFatalErrorMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG
	msgClass := OpsMsgClassFATAL


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
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

func TestOpsMsgDto_SetInfoMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateInfoMsg()

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG
	msgClass := OpsMsgClassINFO

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetWarningMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateWarningMsg()

	xMsg := "This is Warning Message for test object."
	msgId := int64(67)
	msgNo := int64(6067)
	msgType := OpsMsgTypeWARNINGMSG
	msgClass := OpsMsgClassWARNING

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
	}

	if om.IsError() != false {
		t.Error("Expected Warning Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Warning to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetDebugMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateDebugMsg()

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(6238)
	msgType := OpsMsgTypeDEBUGMSG
	msgClass := OpsMsgClassDEBUG

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.MsgClass != msgClass {
		t.Errorf("Expected Messgage Class == '%v'. Instead, Message Class == '%v'.", msgClass, om.MsgClass)
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

func TestOpsMsgDto_SetSuccessfulCompletionMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateSuccessfulCompletionMsg()

	xMsg := "Successful Completion"
	msgId := int64(64)
	msgNo := int64(6064)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION
	msgClass := OpsMsgClassSUCCESSFULCOMPLETION

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

func TestOpsMsgDto_SetNoErrorsNoMessages(t *testing.T) {

	om := testOpsMsgDtoCreateNoErrorsNoMessagesMsg()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSG
	msgClass := OpsMsgClassNOERRORSNOMESSAGES

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

/*
=======================================================================================================
								Private Methods
=======================================================================================================
 */

func testOpsMsgDtoCreateContextInfoObj() OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testOpsMsgDtoCreateParentHistory() []OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ci.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ci.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}

func testOpsMsgDtoCreateStdErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg for test object", 429)
	return om
}

func testOpsMsgDtoCreateFatalErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg for test object", 152)
	return om
}

func testOpsMsgDtoCreateInfoMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage("This is Information Message for test object", 19)
	return om
}

func testOpsMsgDtoCreateWarningMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetWarningMessage("This is Warning Message for test object.", 67)
	return om
}

func testOpsMsgDtoCreateDebugMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage("This is DEBUG Message for test object.", 238)
	return om
}

func testOpsMsgDtoCreateSuccessfulCompletionMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetSuccessfulCompletionMessage( 64)
	return om
}

func testOpsMsgDtoCreateNoErrorsNoMessagesMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetNoErrorsNoMessages(28)
	return om
}