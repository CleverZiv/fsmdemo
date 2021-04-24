package src

import (
	"github.com/looplab/fsm"
)

// 定义状态转换
func NewOrderFsm(initState string) *fsm.FSM {
	return fsm.NewFSM(
		initState,
		fsm.Events{
			{
				Name: CREATE.String(),
				Src:  []string{START.String()},
				Dst:  WAIT_PAY.String(),
			},
			{
				Name: USER_PAIED.String(),
				Src:  []string{WAIT_PAY.String()},
				Dst:  PAIED.String(),
			},
			{
				Name: SELLER_SEND.String(),
				Src:  []string{WAIT_SEND.String(), PAIED.String()},
				Dst:  SENDED.String(),
			},
			{
				Name: USER_SIGNED.String(),
				Src:  []string{SENDED.String()},
				Dst:  FINISHED.String(),
			},
			{
				Name: USER_CANCEL.String(),
				Src:  []string{WAIT_PAY.String(), PAIED.String(), WAIT_SEND.String(), SENDED.String()},
				Dst:  FINISHED.String(),
			},
		},
		fsm.Callbacks{
			"enter_state":               enterState,
			CREATE.afterEventKey():      afterCreate,
			USER_PAIED.afterEventKey():  afterPaied,
			SELLER_SEND.afterEventKey(): afterSend,
			USER_SIGNED.afterEventKey(): afterSigned,
			USER_CANCEL.afterEventKey(): afterCancel,
		},
	)

}

func (e OrderEvent) afterEventKey() string {
	return "after_" + string(e)
}
