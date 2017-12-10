package common

import (
	"testing"
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

	parent := testCreateOpsMsgDtoParentHistory()

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
	parentInfo := testCreateOpsMsgDtoParentHistory()
	contextInfo := testCreateOpsMsgContextInfoObj()

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

// TODO - fix this broken test
/*
func TestOpsMsgDto_InitializeContextWithParentOpsMsg_01(t *testing.T) {
	parentInfo := testCreateOpsMsgDtoParentHistory()
	contextInfo := testCreateOpsMsgContextInfoObj()

	om := OpsMsgDto{}.InitializeContextInfo(parentInfo, contextInfo)
	newMsg := "Information Message Text"
	om.SetInfoMessage(newMsg, 122)


	ci := OpsMsgContextInfo{SourceFileName:"TSource07", ParentObjectName:"PObj07", FuncName: "Func007", BaseMessageId: 7000}
	
	om2 := OpsMsgDto{}.InitializeContextWithParentOpsMsg(om, ci)
	
	l := len(om.ParentContextHistory)

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
*/

func testCreateOpsMsgContextInfoObj() OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testCreateOpsMsgDtoParentHistory() []OpsMsgContextInfo {
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