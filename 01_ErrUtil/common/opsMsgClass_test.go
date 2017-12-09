package common

import "testing"

func TestOpsMsgClass_String_01(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassNOERRORSNOMESSAGES

	var s string

	s = r.String()

	if s != "NOERRORSNOMESSAGES" {
		t.Errorf("Expected string 'NOERRORSNOMESSAGES'. Instead got %v", s)
	}

}

func TestOpsMsgClass_String_02(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassOPERROR

	var s string

	s = r.String()

	if s != "OPERROR" {
		t.Errorf("Expected string 'OPERROR'. Instead got  %v", s)
	}

}


func TestOpsMsgClass_String_03(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassFATAL

	var s string

	s = r.String()

	if s != "FATAL" {
		t.Errorf("Expected string 'FATAL'. Instead got  %v", s)
	}

}

func TestOpsMsgClass_String_04(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassINFO

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead got  %v", s)
	}

}

func TestOpsMsgClass_String_05(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassWARNING

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead got  %v", s)
	}

}

func TestOpsMsgClass_String_06(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassDEBUG

	var s string

	s = r.String()

	if s != "DEBUG" {
		t.Errorf("Expected string 'DEBUG'. Instead got %v", s)
	}

}

func TestOpsMsgClass_String_07(t *testing.T) {
	var r OpsMsgClass

	r = OpsMsgClassSUCCESSFULCOMPLETION

	var s string

	s = r.String()

	if s != "SUCCESS" {
		t.Errorf("Expected string 'SUCCESS'. Instead got %v", s)
	}

}


func TestOpsMsClass_Value_01(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassNOERRORSNOMESSAGES

	i = int(r)

	if r != 0 {
		t.Errorf("Expected OpsMsgClassNOERRORSNOMESSAGES value = ZERO (0). Instead got %v", i)
	}
}




func TestOpsMsClass_Value_02(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassOPERROR

	i = int(r)

	if r != 1 {
		t.Errorf("Expected OPERROR value = 1. Instead got %v", i)
	}
}


func TestOpsMsClass_Value_03(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassFATAL

	i = int(r)

	if r != 2 {
		t.Errorf("Expected FATAL value = 2. Instead, got %v", i)
	}
}

func TestOpsMsClass_Value_04(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassINFO

	i = int(r)

	if r != 3 {
		t.Errorf("Expected INFO value = 3. Instead, got %v", i)
	}
}

func TestOpsMsClass_Value_05(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassWARNING

	i = int(r)

	if r != 4 {
		t.Errorf("Expected WARNING value = 4. Instead, got %v", i)
	}
}

func TestOpsMsClass_Value_06(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassDEBUG

	i = int(r)

	if r != 5 {
		t.Errorf("Expected DEBUG value = 5. Instead got %v", i)
	}
}

func TestOpsMsClass_Value_07(t *testing.T) {
	var r OpsMsgClass

	var i int

	r = OpsMsgClassSUCCESSFULCOMPLETION

	i = int(r)

	if r != 6 {
		t.Errorf("Expected SUCCESSFUL COMPLETION value = 6. Instead got %v", i)
	}
}
