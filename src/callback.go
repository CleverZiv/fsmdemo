package src

import (
	"fmt"
	"github.com/looplab/fsm"
)

func enterState(e *fsm.Event) {
	fmt.Printf("初始状态：%s，事件：%s，目标状态：%s\n", e.Src, e.Event, e.Dst)
}

func afterCreate(e *fsm.Event) {
	fmt.Println("afterCreate 订单已经创建.......")
}

func afterPaied(e *fsm.Event) {
	fmt.Println("afterPaied 订单已经被支付.......")
}

func afterSend(e *fsm.Event) {
	fmt.Println("afterSend 商品已经发货.......")
}

func afterSigned(e *fsm.Event) {
	fmt.Println("afterSigned 商品已经签收.......")
}

func afterCancel(e *fsm.Event) {
	fmt.Println("afterCancel 订单被取消.......")
}
