package common

import (
	"testing"
	"strings"
	"errors"
	"time"
)

func TestOpsMsgDto_ChangeMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	testMsgCtx := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "This is Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"

	err := errors.New(errMsgText)
	xMsg := "This is New Standard Error Message"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgCtx)
	om.SetStdError(prefixMsgText, err, msgId)
	om.ChangeMsg(xMsg)

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

	if !testMsgCtx.Equal(&om.MsgContext) {
		t.Error("Expected testMsgCtx to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
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

func TestOpsMsgDto_ChangeMsg_02(t *testing.T){
	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is the 'changed' Information Message"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om := testOpsMsgDtoCreateInfoMsg()

	om.ChangeMsg(xMsg)

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

func TestOpsMsgDto_ChangeMsgId_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	testMsgCtx := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "This is Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"

	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgCtx)
	om.SetStdError(prefixMsgText, err, msgId)

	msgId = int64(52)
	msgNo = int64(6052)
	om.ChangeMsgId(msgId)

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

	if !testMsgCtx.Equal(&om.MsgContext) {
		t.Error("Expected testMsgCtx to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
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

func TestOpsMsgDto_ChangeMsgId_02(t *testing.T){
	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om := testOpsMsgDtoCreateInfoMsg()

	msgId = int64(849)
	msgNo = int64(6849)

	om.ChangeMsgId(msgId)

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

func TestOpsMsgDto_CopyIn_01(t *testing.T) {

	om1 := testOpsMsgDtoCreateFatalErrorMsg()

	om2 := testOpsMsgDtoCreateInfoMsg()

	om1.CopyIn(&om2)

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
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

	actMsg := om1.GetFmtMessage()

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

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
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

	actMsg := om1.GetFmtMessage()

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

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
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

	actMsg := om1.GetFmtMessage()

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

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
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

	actMsg := om1.GetFmtMessage()

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
	msgType := OpsMsgTypeFATALERRORMSG


	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
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

	actMsg := om2.GetFmtMessage()

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
	msgType := OpsMsgTypeFATALERRORMSG


	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
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

	actMsg := om2.GetFmtMessage()

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
	msgType := OpsMsgTypeFATALERRORMSG

	om1 := OpsMsgDto{}
	om1.SetFatalErrorMessage(xMsg, msgId)

	om2:= testOpsMsgDtoCreateInfoMsg()

	om2 = om1.CopyOut()

	if om2.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om2.MsgType)
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

	actMsg := om2.GetFmtMessage()

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

	if om1.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om1.MsgType)
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

	actMsg := om1.GetFmtMessage()

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

	om2.fmtMessage = "xxxx"

	if om2.Equal(&om1) {
		t.Error("Expected om2!=om1. om2 DID EQUAL om1 - ERROR!")
	}

}

func TestOpsMsgContextInfo_Equal_06(t *testing.T) {

	om1 := testOpsMsgDtoCreateInfoMsg()

	om2 := testOpsMsgDtoCreateFatalErrorMsg()

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

func testOpsMsgDtoCreateContextInfoObj2() OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}
	return ci.New("T2Source04", "P2Obj04", "F2unc004", 4000)
}

func testOpsMsgDtoCreateParentHistory2() []OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}

	x1 := ci.New("T2Source01", "P2Obj01", "F2unc001", 1000)
	x2 := ci.New("T2Source02", "P2Obj02", "F2unc002", 2000)
	x3 := ci.New("T2Source03", "P2Obj03", "F2unc003", 3000)

	parent := make([]OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)

	return parent
}


func testOpsMsgDtoCreateContextInfoObj3() OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}
	return ci.New("T3Source04", "P3Obj04", "F3unc004", 3000)
}

func testOpsMsgDtoCreateParentHistory3() []OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}

	x1 := ci.New("T3Source01", "P3Obj01", "F3unc001", 1000)
	x2 := ci.New("T3Source02", "P3Obj02", "F3unc002", 2000)

	parent := make([]OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)

	return parent
}


func testOpsMsgDtoCreateStdErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg for test object", 429)
	return om
}

func testOpsMsgDtoCreateStdErrorMsg_02() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory2(), testOpsMsgDtoCreateContextInfoObj2())
	om.SetStdErrorMessage("This is Test Standard Error Msg #2", 229)
	return om
}

func testOpsMsgDtoCreateStdErrorMsg_03() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory3(), testOpsMsgDtoCreateContextInfoObj3())
	om.SetStdErrorMessage("This is Test Standard Error Msg #3", 339)
	return om
}

func testOpsMsgDtoCreateFatalErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg for test object", 152)
	return om
}


func testOpsMsgDtoCreateFatalErrorMsg_02() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory2(), testOpsMsgDtoCreateContextInfoObj2())
	om.SetFatalErrorMessage("This is Test FATAL Error Msg #2", 22152)
	return om
}

func testOpsMsgDtoCreateFatalErrorMsg_03() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory3(), testOpsMsgDtoCreateContextInfoObj3())
	om.SetFatalErrorMessage("This is Test FATAL Error Msg #3", 33152)
	return om
}

func testOpsMsgDtoCreateInfoMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage("This is Information Message for test object", 19)
	return om
}

func testOpsMsgDtoCreateInfoMsg_02() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory2(), testOpsMsgDtoCreateContextInfoObj2())
	om.SetInfoMessage("This is Test Information Message #2", 229)
	return om
}

func testOpsMsgDtoCreateInfoMsg_03() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory3(), testOpsMsgDtoCreateContextInfoObj3())
	om.SetInfoMessage("This is Test Information Message #3", 339)
	return om
}

func testOpsMsgDtoCreateWarningMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetWarningMessage("This is Warning Message for test object.", 67)
	return om
}

func testOpsMsgDtoCreateDebugMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage("This is DEBUG Message for test object.", 238)
	return om
}

func testOpsMsgDtoCreateSuccessfulCompletionMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetSuccessfulCompletionMessage("", 64)
	return om
}

func testOpsMsgDtoCreateNoErrorsNoMessagesMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetNoErrorsNoMessages("", 28)
	return om
}

func testOpsMsgDtoCreateSpecErrParentBaseInfo5Elements() []ErrBaseInfo {
	parentBaseInfo := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}

	a := bi.New("TSource01", "PObj01", "Func001", 1000)
	b := bi.New("TSource02", "PObj02", "Func002", 2000)
	c := bi.New("TSource03", "PObj03", "Func003", 3000)
	d := bi.New("TSource04", "PObj04", "Func004", 4000)
	e := bi.New("TSource05", "PObj05", "Func005", 5000)

	parentBaseInfo = append(parentBaseInfo, a, b, c, d, e)


	return parentBaseInfo
}


func testOpsMsgDtoCreateSpecErrBaseInfoObject() ErrBaseInfo {
	bi := ErrBaseInfo{}

	a := bi.New("TSource06", "PObj06", "Func006", 6000)

	return a
}