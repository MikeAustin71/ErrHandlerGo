package common

import "testing"

/*  'opsmsgtype_test.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/

func TestOpsMsgTypeText001(t *testing.T) {
	var r OpsMsgType

	r = OpsNOERRORNOMSGTYPE

	var s string

	s = r.String()

	if s != "NOERRORSNOMSGS" {
		t.Errorf("Expected string 'NOERRORSNOMSGS'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText002(t *testing.T) {
	var r OpsMsgType

	r = OpsERRORMSGTYPE

	var s string

	s = r.String()

	if s != "ERROR" {
		t.Errorf("Expected string 'ERROR'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText003(t *testing.T) {
	var r OpsMsgType

	r = OpsINFOMSGTYPE

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText004(t *testing.T) {
	var r OpsMsgType

	r = OpsWARNINGMSGTYPE

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText005(t *testing.T) {
	var r OpsMsgType

	r = OpsDEBUGMSGTYPE

	var s string

	s = r.String()

	if s != "DEBUG" {
		t.Errorf("Expected string 'DEBUG'. Instead, got %v", s)
	}

}


func TestOpsMsgTypeValue001(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsNOERRORNOMSGTYPE

	i = int(r)

	if r != 0 {
		t.Errorf("Expected 'OpsNOERRORNOMSGTYPE' value = 0. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue002(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsERRORMSGTYPE

	i = int(r)

	if r != 1 {
		t.Errorf("Expected 'OpsERRORMSGTYPE' value = 1. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue003(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsINFOMSGTYPE

	i = int(r)

	if r != 2 {
		t.Errorf("Expected 'OpsINFOMSGTYPE' value = 2. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue004(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsWARNINGMSGTYPE

	i = int(r)

	if r != 3 {
		t.Errorf("Expected 'OpsWARNINGMSGTYPE' value = 3. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue005(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsDEBUGMSGTYPE

	i = int(r)

	if r != 4 {
		t.Errorf("Expected 'OpsDEBUGMSGTYPE' value = 4. Instead, got %v", i)
	}

}
