package main

import (
	"MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
)

/*

import (
	"MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"errors"
	"fmt"
)

*/

func main() {

	//common.TestOpsMsgDto_Example006_NewStdErrorMsg()


}


func testOpsMsgCollectionCreateContextInfoObj() common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testOpsMsgCollectionCreateParentHistory() []common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ci.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ci.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]common.OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}


func testOpsMsgCollectionCreateStdErrorMsg_01() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg for test object", 429)
	return om
}

func testOpsMsgCollectionCreateStdErrorMsg_02() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg #2 for test object", 430)
	return om
}

func testOpsMsgCollectionCreateFatalErrorMsg_01() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg for test object", 152)
	return om
}

func testOpsMsgCollectionCreateFatalErrorMsg_02() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg #2 for test object", 153)
	return om
}

func testOpsMsgCollectionCreateInfoMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage("This is Information Message for test object", 19)
	return om
}

func testOpsMsgCollectionCreateWarningMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetWarningMessage("This is Warning Message for test object.", 67)
	return om
}

func testOpsMsgCollectionCreateDebugMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage("This is DEBUG Message for test object.", 238)
	return om
}

func testOpsMsgCollectionCreateSuccessfulCompletionMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetSuccessfulCompletionMessage("", 64)
	return om
}

func testOpsMsgCollectionCreateNoErrorsNoMessagesMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetNoErrorsNoMessages("", 28)
	return om
}

func testOpsMsgCollectionCreateT01Collection() common.OpsMsgCollection {

	om1 := testOpsMsgCollectionCreateFatalErrorMsg_02()

	om2 := testOpsMsgCollectionCreateStdErrorMsg_02()

	om3 := testOpsMsgCollectionCreateInfoMsg()

	om4 := testOpsMsgCollectionCreateDebugMsg()

	om5 := testOpsMsgCollectionCreateWarningMsg()

	om6 := testOpsMsgCollectionCreateStdErrorMsg_01()

	om7 := testOpsMsgCollectionCreateFatalErrorMsg_01()

	opMsgs := common.OpsMsgCollection{}

	opMsgs.AddOpsMsg(om1)
	opMsgs.AddOpsMsg(om2)
	opMsgs.AddOpsMsg(om3)
	opMsgs.AddOpsMsg(om4)
	opMsgs.AddOpsMsg(om5)
	opMsgs.AddOpsMsg(om6)
	opMsgs.AddOpsMsg(om7)

	return opMsgs
}

func testOpsMsgCollectionCreateT02Collection() common.OpsMsgCollection {

	om1 := testOpsMsgCollectionCreateFatalErrorMsg_01()

	om2 := testOpsMsgCollectionCreateStdErrorMsg_01()

	om3 := testOpsMsgCollectionCreateInfoMsg()

	om4 := testOpsMsgCollectionCreateDebugMsg()

	om5 := testOpsMsgCollectionCreateWarningMsg()

	om6 := testOpsMsgCollectionCreateStdErrorMsg_02()

	om7 := testOpsMsgCollectionCreateFatalErrorMsg_02()

	opMsgs := common.OpsMsgCollection{}

	opMsgs.AddOpsMsg(om1)
	opMsgs.AddOpsMsg(om2)
	opMsgs.AddOpsMsg(om3)
	opMsgs.AddOpsMsg(om4)
	opMsgs.AddOpsMsg(om5)
	opMsgs.AddOpsMsg(om6)
	opMsgs.AddOpsMsg(om7)

	return opMsgs
}

/*
func testOpsMsgDtoCreateParentHistory() []common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ci.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ci.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]common.OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}

*/

func testOpsMsgDtoCreateStdErrorMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg for test object", 429)
	return om
}

func testOpsMsgDtoCreateFatalErrorMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg for test object", 152)
	return om
}

func testOpsMsgDtoCreateInfoMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage("This is Information Message for test object", 19)
	return om
}

func testOpsMsgDtoCreateWarningMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetWarningMessage("This is Warning Message for test object.", 67)
	return om
}

func testOpsMsgDtoCreateDebugMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage("This is DEBUG Message for test object.", 238)
	return om
}

func testOpsMsgDtoCreateSuccessfulCompletionMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetSuccessfulCompletionMessage( "", 64)
	return om
}

func testOpsMsgDtoCreateNoErrorsNoMessagesMsg() common.OpsMsgDto {
	om := common.OpsMsgDto{}.InitializeAllContextInfo(testOpsMsgDtoCreateParentHistory(), testOpsMsgDtoCreateContextInfoObj())
	om.SetNoErrorsNoMessages("", 28)
	return om
}


func testOpsMsgDtoCreateContextInfoObj() common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testOpsMsgDtoCreateParentHistory() []common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ci.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ci.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]common.OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}

