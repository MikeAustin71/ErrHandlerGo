package main

import (
	common "MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"errors"
	"fmt"
)

func main() {

	err := errors.New("Some Error")
	var se common.SpecErr
	x := se.New("errprefix", err, true, "main.go", "main", 334)

	fmt.Println(x.Error())

}
