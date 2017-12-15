package common

import "fmt"

// TestOpsMsgDto_Example_001_NewInfoMsg - Test
// Ops Message New Information Message Display
func TestOpsMsgDto_Example001_NewInfoMsg() {
	msg := OpsMsgDto{}.NewInfoMsg("This is the message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example_002_NewInfoMsg - Test
// Ops Message New Information Message Display
func TestOpsMsgDto_Example002_NewInfoMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewInfoMsg("This is the message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example_003_NewInfoMsg - Test
// Ops Message New Information Message Display
func TestOpsMsgDto_Example003_NewInfoMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewInfoMsg("This is the message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example_004_NewInfoMsg - Test
// Ops Message New Information Message Display
func TestOpsMsgDto_Example004_NewInfoMsg() {
	msg := OpsMsgDto{}.NewInfoMsg("This is the message text.", 0)

	fmt.Printf(msg.GetMessage())
}



// TestOpsMsgDto_Example001_NewStdErrorMsg - Test
// Ops Message Standard Error Display
func TestOpsMsgDto_Example001_NewStdErrorMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewStdErrorMsg("This is Standard Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example002_NewStdErrorMsg - Test
// Ops Message Standard Error Display
func TestOpsMsgDto_Example002_NewStdErrorMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewStdErrorMsg("This is Standard Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewStdErrorMsg - Test
// Ops Message Standard Error Display
func TestOpsMsgDto_Example003_NewStdErrorMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewStdErrorMsg("This is Standard Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewStdErrorMsg - Test
// Ops Message Standard Error Display
func TestOpsMsgDto_Example004_NewStdErrorMsg() {
	msg := OpsMsgDto{}.NewStdErrorMsg("This is Standard Error message text.", 0)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example001_NewFatalErrorMsg - Test
// Ops Message Fatal Error Display
func TestOpsMsgDto_Example001_NewFatalErrorMsg() {
	msg := OpsMsgDto{}.NewFatalErrorMsg("This is FATAL Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example002_NewFatalErrorMsg - Test
// Ops Message Fatal Error Display
func TestOpsMsgDto_Example002_NewFatalErrorMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewFatalErrorMsg("This is FATAL Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewFatalErrorMsg - Test
// Ops Message Fatal Error Display
func TestOpsMsgDto_Example003_NewFatalErrorMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewFatalErrorMsg("This is FATAL Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewFatalErrorMsg - Test
// Ops Message Fatal Error Display
func TestOpsMsgDto_Example004_NewFatalErrorMsg() {
	msg := OpsMsgDto{}.NewFatalErrorMsg("This is FATAL Error message text.", 0)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example001_NewWarningMsg - Test
// Ops Message Warning Message Display
func TestOpsMsgDto_Example001_NewWarningMsg() {
	msg := OpsMsgDto{}.NewWarningMsg("This is Warning Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TTestOpsMsgDto_Example002_NewWarningMsg - Test
// Ops Message Warning Message Display
func TestOpsMsgDto_Example002_NewWarningMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewWarningMsg("This is Warning Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewWarningMsg - Test
// Ops Message Warning Message Display
func TestOpsMsgDto_Example003_NewWarningMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewWarningMsg("This is Warning message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewWarningMsg - Test
// Ops Message Warning Message Display
func TestOpsMsgDto_Example004_NewWarningMsg() {
	msg := OpsMsgDto{}.NewWarningMsg("This is Warning message text.", 0)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example001_NewDEBUGMsg - Test
// Ops Message DEBUG Message Display
func TestOpsMsgDto_Example001_NewDEBUGMsg() {
	msg := OpsMsgDto{}.NewDebugMsg("This is DEBUG Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example002_NewDEBUGMsg - Test
// Ops Message DEBUG Message Display
func TestOpsMsgDto_Example002_NewDEBUGMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewDebugMsg("This is DEBUG Error message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewDEBUGMsg - Test
// Ops Message DEBUG Message Display
func TestOpsMsgDto_Example003_NewDEBUGMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewDebugMsg("This is DEBUG message text.", 974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewDEBUGMsg - Test
// Ops Message DEBUG Message Display
func TestOpsMsgDto_Example004_NewDEBUGMsg() {
	msg := OpsMsgDto{}.NewDebugMsg("This is DEBUG message text.", 0)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example001_NewSuccessfulCompletionMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example001_NewSuccessfulCompletionMsg() {
	msg := OpsMsgDto{}.NewSuccessfulCompletionMsg("", 974)

	fmt.Printf(msg.GetMessage())
}

// TTestOpsMsgDto_Example002_NewSuccessfulCompletionMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example002_NewSuccessfulCompletionMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewSuccessfulCompletionMsg("",974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewSuccessfulCompletionMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example003_NewSuccessfulCompletionMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewSuccessfulCompletionMsg("",974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewSuccessfulCompletionMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example004_NewSuccessfulCompletionMsg() {
	msg := OpsMsgDto{}.NewSuccessfulCompletionMsg("", 0)

	fmt.Printf(msg.GetMessage())
}

// TTestOpsMsgDto_Example002_NewSuccessfulCompletionMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example005_NewSuccessfulCompletionMsg() {
	msgText := "Xray = 6"
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewSuccessfulCompletionMsg(msgText,974)

	fmt.Printf(msg.GetMessage())
}


// TestOpsMsgDto_Example001_NewNoErrorsNoMessagesMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example001_NewNoErrorsNoMessagesMsg() {
	msg := OpsMsgDto{}.NewNoErrorsNoMessagesMsg("", 974)

	fmt.Printf(msg.GetMessage())
}

// TTestOpsMsgDto_Example002_NewNoErrorsNoMessagesMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example002_NewNoErrorsNoMessagesMsg() {
	msg := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(),testExampleOpsMsgDtoCreateContextInfoObj()).NewNoErrorsNoMessagesMsg("",974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example003_NewNoErrorsNoMessagesMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example003_NewNoErrorsNoMessagesMsg() {
	msg := OpsMsgDto{}.InitializeWithMessageContext(testExampleOpsMsgDtoCreateContextInfoObj()).NewNoErrorsNoMessagesMsg("",974)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewNoErrorsNoMessagesMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example004_NewNoErrorsNoMessagesMsg() {
	msg := OpsMsgDto{}.NewNoErrorsNoMessagesMsg("", 0)

	fmt.Printf(msg.GetMessage())
}

// TestOpsMsgDto_Example004_NewNoErrorsNoMessagesMsg - Test
// Ops Message SuccessfulCompletion Message Display
func TestOpsMsgDto_Example005_NewNoErrorsNoMessagesMsg() {

	xMsg := "Xray = 6"

	msg := OpsMsgDto{}.NewNoErrorsNoMessagesMsg(xMsg, 0)

	fmt.Printf(msg.GetMessage())
}



/*
=======================================================================================================
								Private Methods
=======================================================================================================
 */

func testExampleOpsMsgDtoCreateContextInfoObj() OpsMsgContextInfo {
	ci := OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testExampleOpsMsgDtoCreateParentHistory() []OpsMsgContextInfo {
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




func testExampleOpsMsgDtoCreateStdErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetStdErrorMessage("This is Standard Error Msg for test object", 429)
	return om
}

func testExampleOpsMsgDtoCreateFatalErrorMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage("This is FATAL Error Msg for test object", 152)
	return om
}

func testExampleOpsMsgDtoCreateInfoMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage("This is Information Message for test object", 19)
	return om
}

func testExampleOpsMsgDtoCreateWarningMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetWarningMessage("This is Warning Message for test object.", 67)
	return om
}

func testExampleOpsMsgDtoCreateDebugMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage("This is DEBUG Message for test object.", 238)
	return om
}

func testExampleOpsMsgDtoCreateSuccessfulCompletionMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetSuccessfulCompletionMessage("", 64)
	return om
}

func testExampleOpsMsgDtoCreateNoErrorsNoMessagesMsg() OpsMsgDto {
	om := OpsMsgDto{}.InitializeAllContextInfo(testExampleOpsMsgDtoCreateParentHistory(), testExampleOpsMsgDtoCreateContextInfoObj())
	om.SetNoErrorsNoMessages("",28)
	return om
}