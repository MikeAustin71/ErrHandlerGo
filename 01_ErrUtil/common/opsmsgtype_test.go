package common

import "testing"

/*  'opsmsgtype_test.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/

func TestLogMsgTypeText_01(t *testing.T) {
	var r OpsMsgType

	r = OpsERRORMSGTYPE

	var s string

	s = r.String()

	if s != "ERROR" {
		t.Errorf("Expected string 'ERROR'. Instead, got %v", s)
	}

}

func TestLogMsgTypeText_02(t *testing.T) {
	var r OpsMsgType

	r = OpsINFOMSGTYPE

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead, got %v", s)
	}

}

func TestLogMsgTypeText_03(t *testing.T) {
	var r OpsMsgType

	r = OpsWARNINGMSGTYPE

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead, got %v", s)
	}

}


func TestLogMsgTypeValue_01(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsERRORMSGTYPE

	i = int(r)

	if r != 0 {
		t.Error("Expected 'OpsERRORMSGTYPE' value = 0. Instead, got %v", i)
	}

}

func TestLogMsgTypeValue_02(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsINFOMSGTYPE

	i = int(r)

	if r != 1 {
		t.Error("Expected 'OpsINFOMSGTYPE' value = 1. Instead, got %v", i)
	}

}

func TestLogMsgTypeValue_03(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsWARNINGMSGTYPE

	i = int(r)

	if r != 2 {
		t.Error("Expected 'OpsWARNINGMSGTYPE' value = 2. Instead, got %v", i)
	}

}
