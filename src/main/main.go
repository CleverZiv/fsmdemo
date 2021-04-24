package main

import (
	"code.leng.org/fsmdemo/src"
	"fmt"
)

func main() {
	myFsm := src.NewOrderFsm(src.START.String())
	fmt.Println(myFsm.Current())
	myFsm.Event(src.CREATE.String())

	myFsm.Event(src.USER_PAIED.String())

	myFsm.Event(src.SELLER_SEND.String())

	myFsm.Event(src.USER_CANCEL.String())

	myFsm.Event(src.USER_SIGNED.String())

}
