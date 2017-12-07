package common

import (
	"testing"
)


func TestErrMsgType_String001(t *testing.T) {

	var r ErrMsgType

	r = ErrTypeNOERRORSALLCLEAR

	var s string

	s = r.String()

	if s != "NOERRORSALLCLEAR" {
		t.Errorf("Expected string 'NOERRORSALLCLEAR'. Instead got %v", s)
	}

}

func TestErrMsgType_String002(t *testing.T) {

	var r ErrMsgType

	r = ErrTypeFATAL

	var s string

	s = r.String()

	if s != "FATAL" {
		t.Errorf("Expected string 'FATAL'. Instead got %v", s)
	}

}

func TestErrMsgType_String003(t *testing.T) {

	var r ErrMsgType

	r = ErrTypeERROR

	var s string

	s = r.String()

	if s != "ERROR" {
		t.Errorf("Expected string 'ERROR'. Instead got %v", s)
	}

}

func TestErrMsgType_String004(t *testing.T) {

	var r ErrMsgType

	r = ErrTypeWARNING

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead got %v", s)
	}

}

func TestErrMsgType_String005(t *testing.T) {

	var r ErrMsgType

	r = ErrTypeInfo

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead got %v", s)
	}

}


func TestErrMsgType_Value001(t *testing.T) {

	var r ErrMsgType

	var i int

	r = ErrTypeNOERRORSALLCLEAR

	i = int(r)

	if r != 0 {
		t.Errorf("Expected ErrTypeNOERRORSALLCLEAR value = ZERO (0). Instead got %v", i)
	}

}

func TestErrMsgType_Value002(t *testing.T) {

	var r ErrMsgType

	var i int

	r = ErrTypeFATAL

	i = int(r)

	if r != 1 {
		t.Errorf("Expected ErrTypeFATAL value = 1. Instead got %v", i)
	}

}

func TestErrMsgType_Value003(t *testing.T) {

	var r ErrMsgType

	var i int

	r = ErrTypeERROR

	i = int(r)

	if r != 2 {
		t.Errorf("Expected ErrTypeERROR value = 2. Instead got %v", i)
	}

}

func TestErrMsgType_Value004(t *testing.T) {

	var r ErrMsgType

	var i int

	r = ErrTypeWARNING

	i = int(r)

	if r != 3 {
		t.Errorf("Expected ErrTypeWARNING value = 3. Instead got %v", i)
	}

}

func TestErrMsgType_Value005(t *testing.T) {

	var r ErrMsgType

	var i int

	r = ErrTypeInfo

	i = int(r)

	if r != 4 {
		t.Errorf("Expected ErrTypeInfo value = 4. Instead got %v", i)
	}

}