package common

import "testing"

func TestOpsMsgClass_String_01(t *testing.T) {
	var r OpsMsgClass

	r = MsgClassDEBUG

	var s string

	s = r.String()

	if s != "DEBUG" {
		t.Errorf("Expected string 'DEBUG'. Instead got %v", s)
	}

}

func TestOpsMsgClass_String_02(t *testing.T) {
	var r OpsMsgClass

	r = MsgClassOPERROR

	var s string

	s = r.String()

	if s != "OPERROR" {
		t.Errorf("Expected string 'OPERROR'. Instead got  %v", s)
	}

}


func TestOpsMsgClass_String_03(t *testing.T) {
	var r OpsMsgClass

	r = MsgClassFATAL

	var s string

	s = r.String()

	if s != "FATAL" {
		t.Errorf("Expected string 'FATAL'. Instead got  %v", s)
	}

}

func TestOpsMsgClass_String_04(t *testing.T) {
	var r OpsMsgClass

	r = MsgClassINFO

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead got  %v", s)
	}

}

func TestOpsMsgClass_String_05(t *testing.T) {
	var r OpsMsgClass

	r = MsgClassWARNING

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead got  %v", s)
	}

}

func TestOpsMsClass_Value_01(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = MsgClassDEBUG

	i = int(r)

	if r != 0 {
		t.Errorf("Expected DEBUG value = ZERO (0). Instead got %v", i)
	}
}



func TestOpsMsClass_Value_02(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = MsgClassOPERROR

	i = int(r)

	if r != 1 {
		t.Errorf("Expected OPERROR value = 1. Instead got %v", i)
	}
}


func TestOpsMsClass_Value_03(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = MsgClassFATAL

	i = int(r)

	if r != 2 {
		t.Errorf("Expected FATAL value = 2. Instead, got %v", i)
	}
}

func TestOpsMsClass_Value_04(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = MsgClassINFO

	i = int(r)

	if r != 3 {
		t.Errorf("Expected INFO value = 3. Instead, got %v", i)
	}
}

func TestOpsMsClass_Value_05(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = MsgClassWARNING

	i = int(r)

	if r != 4 {
		t.Errorf("Expected WARNING value = 4. Instead, got %v", i)
	}
}
