package src

type OrderStatus string

// 订单状态常量
const (
	START     OrderStatus = "100"
	WAIT_PAY  OrderStatus = "110"
	PAIED     OrderStatus = "120"
	WAIT_SEND OrderStatus = "130"
	SENDED    OrderStatus = "140"
	SIGNED    OrderStatus = "150"
	CANCEL    OrderStatus = "160"
	FINISHED  OrderStatus = "170"
)

// 订单事件常量

type OrderEvent string

const (
	CREATE      OrderEvent = "creat"
	USER_PAIED  OrderEvent = "user_paied"
	SELLER_SEND OrderEvent = "seller_send"
	USER_SIGNED OrderEvent = "user_signed"
	USER_CANCEL OrderEvent = "user_cancel"
)

func (e OrderEvent) String() string {
	return string(e)
}

func (e OrderStatus) String() string {
	return string(e)
}
