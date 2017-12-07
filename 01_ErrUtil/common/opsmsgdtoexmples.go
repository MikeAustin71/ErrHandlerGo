package common

import "fmt"

// TestOpsMsgDto_Example_001_NewInfoMsg - Test
// Ops Message Display
func TestOpsMsgDtoExample001NewInfoMsg() {
	msg := OpsMsgDto{}.NewInfoMsg("This is the message text.", 974)

	fmt.Printf(msg.GetMessage())
}
