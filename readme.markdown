# FSM 的 demo 设计实现
官方文档：https://pkg.go.dev/github.com/looplab/fsm#readme-usage-as-a-struct-field

原项目地址：https://github.com/looplab/fsm

## 对象--商品订单
以电商场景下的订单实体为研究对象，可以归纳出以下几个状态
1. 创建
2. 待支付
3. 已支付
4. 待发货
5. 已发货
6. 已签收   
7. 取消
8. 完成
