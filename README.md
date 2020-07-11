# golang版  唯品会开放平台SDK

具体使用方式参考测试代码

## [唯品会开放平台API列表](http://vop.vip.com/home#/api/service/list)
### [市场和渠道服务](http://vop.vip.com/home#/api/service/list/2)
#### 联盟商品开放服务
- [获取指定商品id集合的商品信息](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.0.0/getByGoodsIds)
- [获取联盟在推商品列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.0.0/goodsList)
- [根据关键词查询商品列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.0.0/query)
- [相似推荐商品列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.0.0/similarRecommend)
- [猜你喜欢商品列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.0.0/userRecommend)
#### 联盟订单开放服务
- [服务健康度检查](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderService-1.0.0/healthCheck)
- [获取订单信息列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderService-1.0.0/orderList)
- [获取维权订单列表](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderService-1.0.0/refundOrderList)
####  联盟PID开放服务
- [创建推广位PID](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidService-1.0.0/genPid)
- [查询推广位PID](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidService-1.0.0/queryPid)
#### 联盟链接开放服务
- [根据商品id生成联盟链接](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.0.0/genByGoodsId)
- [根据唯品会链接生成联盟链接](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.0.0/genByVIPUrl)
- [进行cps链接的解析](http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.0.0/vipLinkCheck)
