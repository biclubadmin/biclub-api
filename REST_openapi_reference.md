# 行情类接口


#### 
``` 
访问根路径：https://api.biclub.com/api/market
symbol规则： 基础币种+计价币种。如btc-usdt, bz-usdt, eth-usdt 以此类推
```

### 交易对API

 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v2/market/common/symbols
 ```
 
 #### GET /market/common/symbols 获取最新交易对
 
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 
 data 说明:
 
 ```
    "data": {
         "symbol": 交易对,
         "status": 交易对状态，默认trading,
         "baseAsset": 基础货币,
         "baseAssetPrecision": 基础货币精度,
         "quoteAsset": 计价货币,
         "quoteAssetPrecision": 计价货币精度
   }
 ```
 
  请求响应示例
 
 ```
 /* GET /market/common/symbols */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": [
        {
            "symbol": "btc-usdt",
            "status": "trading",
            "baseAsset": "btc",
            "baseAssetPrecision": 4,
            "quoteAsset": "usdt",
            "quoteAssetPrecision": 2
        },
        ...
    ]
 }
 ```
 
 ### 行情API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/market/tickers?symbol=bz-usdt
 ```
 
 #### GET /market/tickers 获取市场24小时成交量数据
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": {
        "id": 消息ID,
        "minPrice": 24小时最低价,
        "maxPrice": 24小时最高价,
        "averagePrice": 平均价,
        "number": 24小时成交量,
        "count": 24小时累积成交数,
        "vol": 累积成交额：每一笔成交价 * 该笔的成交量,
        "open": 开盘价,
        "last": 新成交价,
        "time": 返回数据时服务器时间,
        "buy": 买一价,
        "sell": 卖一价
   }
 ```
 
 
 请求响应示例
 
 ```
 /* GET /market/tickers?symbol=bz-usdt */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": {
        "id": "6d9ae8ec-e96a-4313-acc8-c4497cd06aee",
        "minPrice": 0,
        "maxPrice": 0,
        "averagePrice": 0,
        "number": 0,
        "count": 0,
        "vol": 0,
        "open": 0,
        "last": 0,
        "time": 1534406100072,
        "buy": 0,
        "sell": 10000
   }
 }
 ```
 
 ### 挂单深度API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/market/depth?symbol=btc-usdt
 ```
 
 #### GET /market/depth 获取交易对市场挂单深度信息
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": {
        "asks": 卖方深度【价格,数量】,
        "bids": 买方深度【价格,数量】
   }
 ```
 
 
 请求响应示例
 
 ```
 /* GET /market/depth?symbol=bz-usdt */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": {
        "bids": [
                    [200.01, 2],
                    [100,500],
                    [2.22, 1.555],
                    [0.21,2]
              ],
        "asks": [
                    [430.01,20.241],
                    [420.01,1.4],
                    [410.01,1.505],
                    [300,48.55]
              ]
   }
 }
 ```
 
 ### 最新成交记录API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/market/trades?symbol=bz-usdt&size=10
 ```
 
 #### GET /market/trades 获取交易对最新成交信息
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | size    | true  | int | 数量   |  20  | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": {
        "timestamp": 交易时间,
        "price": 价格,
        "amount": 数量,
        "side": 方向：买入、卖出
   }
 ```
 
 
 请求响应示例
 
 ```
 /* GET /market/trades?symbol=bz-usdt&size=2 */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": {
        {
            "timestamp": 1534123476275,
            "price": 854.11,
            "amount": 1,
            "side": "buy"
        },
        {
            "timestamp": 1534123082697,
            "price": 854.11,
            "amount": 2,
            "side": "sell"
        }
   }
 }
 ```
 
 ### K线API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/market/history/kline?symbol=btc-usdt&period=30&size=3
 ```
 
 #### GET /market/history 获取交易对K线信息
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | period    | true  | string | 时间颗粒   |    | 1,15,30,60,D,2D,3D,W,3W,M,6M|
 | size    | true  | int | 数量   |  150  | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": {
        "klineList": [
            {
                "timestamp": 交易时间,
                "close": 收盘价,
                "open": 开盘价,
                "high": 最高价格,
                "low": 最低价格,
                "vol": 交易量
            }
   }
 ```
 
 
 请求响应示例
 
 ```
 /* GET /market/history/kline?symbol=btc-usdt&period=30&size=3 */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": {
        "status": "ok",
        "errmsg": null,
        "klineList": [
            {
                "timestamp": "1526117337",
                "close": "132.00",
                "open": "132.00",
                "high": "132.00",
                "low": "132.00",
                "vol": "0.0000"
            },
            {
                "timestamp": "1526203737",
                "close": "124.00",
                "open": "124.00",
                "high": "124.00",
                "low": "124.00",
                "vol": "0.0000"
            },
            {
                "timestamp": "1526462937",
                "close": "72.00",
                "open": "88.00",
                "high": "90.00",
                "low": "70.00",
                "vol": "332.0000"
            }
        ]
   }
 }
 ```
 # 交易类接口
 
 
 #### 
``` 
访问根路径：https://api.biclub.com/api/v1
symbol规则： 基础币种+计价币种。如btc-usdt, bz-usdt, eth-usdt 以此类推
```

 
 ### 委托下单API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/order/orders/place
 ```
 
 #### POST /v1/order/orders/place 下单操作
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | number    | true  | string | 委托数量   |    | |
 | orderType    | true  | string | 委托类型   |    |buy-limit:现价买入,sell-limit:限价卖出，buy-market：市价买入，sell-market：市价卖出 |
 | price    | true  | string | 委托价格   |    | |
 | source    | true  | string | 来源   |    | web，ios，android，api |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 订单ID
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/order/orders/place */
 {
   "number": "1",
   "orderType": "buy-limit",
   "price": "6530",
   "source": "web",
   "symbol": "btc-usdt",
   "timestamp":1535955784075,
   "accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
   "sign":"b210c02f58c193176515020bc7ff332b00efd55285b12df3cdfe06e8253c0490"
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/order/orders/place */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": "87e8c3ed-dd11-4522-a968-aba88ae2733b"
 }
 ```
 
 ### 取消委托订单API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/order/orders/cancel
 ```
 
 #### POST /v1/order/orders/cancel 取消订单操作
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | orderId    | true  | string | 委托订单号   |    | 委托下单返回的订单号 |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 操作结果
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/order/orders/cancel */
 {
   "orderId":"87e8c3ed-dd11-4522-a968-aba88ae2733b",
   "symbol": "btc-usdt",
   "timestamp":1535955784075,
   "accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
   "sign":"b210c02f58c193176515020bc7ff332b00efd55285b12df3cdfe06e8253c0490"
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/order/orders/cancel */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": "取消成功"
 }
 ```
 
  ### 批量取消委托订单API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/order/orders/batchcancel
 ```
 
 #### POST /v1/order/orders/batchcancel 批量取消订单操作
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | orderIdList    | true  | list | 委托订单号   |    | 委托下单返回的订单号最多50条 |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 操作结果
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/order/orders/batchcancel */
 {
   "orderIdList":["87e8c3ed-dd11-4522-a968-aba88ae2733b","e0b4293e-1fc3-4328-b021-e542148824fa"],
   "symbol": "btc-usdt",
   "timestamp":1535955784075,
   "accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
   "sign":"b210c02f58c193176515020bc7ff332b00efd55285b12df3cdfe06e8253c0490"
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/order/orders/batchcancel */
 {
   "success": true,
   "returnCode": "200",
   "returnMsg": "success",
   "data": "取消成功"
 }
 ```
 
 
 ### 查询指定委托单API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/order/orders/consign
 ```
 
 #### POST /v1/order/orders/consign 获取指定委托单
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | orderId    | true  | string | 委托订单号   |    | 委托下单返回的订单号 |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 
     {
       "orderId": 委托订单号,
       "datetime": 委托时间,
       "symbol": 交易对,
       "direct": 交易方向，0：买 1：卖,
       "price": 委托价格,
       "amount": 委托数量,
       "allBalance": 委托总额,
       "doneAmount": 已成交数量,
       "noAmount": 未成交数量,
       "donePrice": 成交均价,
       "status": 状态：0:挂单、1:部分成交、2:部分成交撤销、3:已撤销、4:全部成交,
       "orderType": 委托类型：限价买：buy-limit,限价卖:sell-limit,市价买:buy-market,市价卖:sell-market
     }
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/order/orders/consign */
 {
    "symbol":"btc-usdt",
	"orderId":"51de9b3d-11f0-4c46-9aa2-abd389797659",
	"accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
	"sign":"efcc6861e5b81a9d5624eaac745819d2e4b9b7edc513d3fef9a33bc351a53c6d",
	"timestamp":1535943779122
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/order/orders/consign */
 {
    "success": true,
    "returnCode": "200",
    "returnMsg": "success",
    "data": {
        "orderId": "6326adf2-0389-46f8-bb01-1b4d1663e517",
        "datetime": "2018-08-21 15:21:58.928948",
        "symbol": "btc-usdt",
        "direct": 0,
        "price": 6530,
        "amount": 1,
        "allBalance": null,
        "doneAmount": 0,
        "noAmount": null,
        "donePrice": null,
        "status": "0",
        "orderType": "buy-limit"
    }
 }
 ```
 
 ### 查询委托单列表API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/order/orders/consigns
 ```
 
 #### POST /v1/order/orders/consigns 获取委托单列表
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | symbol    | true  | string | 交易对   |    | btc-usdt, bz-usdt, eth-usdt ...|
 | direct    | true  | string | 方向   |    | 全部、0买入、1卖出 |
 | size    | true  | int | 数量   |    |  |
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 
     {
       "orderId": 委托订单号,
       "datetime": 委托时间,
       "symbol": 交易对,
       "direct": 交易方向，0：买 1：卖,
       "price": 委托价格,
       "amount": 委托数量,
       "allBalance": 委托总额,
       "doneAmount": 已成交数量,
       "noAmount": 未成交数量,
       "donePrice": 成交均价,
       "status": 状态：0:挂单、1:部分成交、2:部分成交撤销、3:已撤销、4:全部成交,
       "orderType": 委托类型：限价买：buy-limit,限价卖:sell-limit,市价买:buy-market,市价卖:sell-market
     }
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/order/orders/consigns */
 {
    "symbol":"btc-usdt",
	"direct":"1",
	"size":1,
    "timestamp":1535955784075,
    "accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
    "sign":"b210c02f58c193176515020bc7ff332b00efd55285b12df3cdfe06e8253c0490"
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/order/orders/consigns */
 {
   "success": true,
    "returnCode": "200",
    "returnMsg": "success",
    "data": {
        "orderId": "6326adf2-0389-46f8-bb01-1b4d1663e517",
        "datetime": "2018-08-21 15:21:58.928948",
        "symbol": "btc-usdt",
        "direct": 0,
        "price": 6530,
        "amount": 1,
        "allBalance": null,
        "doneAmount": 0,
        "noAmount": null,
        "donePrice": null,
        "status": "0",
        "orderType": "buy-limit"
    },
    ...
 }
 ```
 
 ### 查询账户资产列表API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/account/balance/list
 ```
 
 #### POST /v1/account/balance/list 获取账户资产列表
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | coinType    | true  | string | 币种   |    |  |
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 
     {
       "coinType": 币种,
       "totalAmount": 总资产,
       "available": 冻结数量,
       "frozen": 可用资产数量,
       "depositFlag": 是否可充值,
       "withdrawFlag": 是否可提现
     }
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/account/balance/list */
 {
    "accessKey":"ed221a61-67bb-4764-b602-749bf1354b59",
	"coinType":"btc",
	"sign":"c5935ff633d8474c5fdea61d40d2c4fd6970ee19ac323cfa01b63bf1acf481de",
	"timestamp":1535679734680
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/account/balance/list */
 {
   "success": true,
    "returnCode": "200",
    "returnMsg": "success",
    "data": [
        {
            "coinType": "usdt",
            "totalAmount": "1000.00000000",
            "available": "999.00000000",
            "frozen": "1.0000000",
            "depositFlag": "0",
            "withdrawFlag": "0"
        },
        ...
    ]
 }
 ```
 
 ### 查询充值地址API
 
 API接口请求地址
 
 例：
 
 ```
 https://api.biclub.com/api/v1/account/deposit/address
 ```
 
 #### POST /v1/account/deposit/address 获取充值地址
 
  请求参数:
 
 | 参数名称    | 是否必须  | 类型  | 描述    | 默认值   | 取值范围  |
 | ------- | ----- | ------ | ----- | ----- | ----  |
 | coinTypes    | true  | string | 币种列表   |    |btc,usdt,eth,ltc  |
 | accessKey    | true  | string | accessKey   |    | |
 | timestamp    | true  | string | 时间戳   |    | |
 | sign    | true  | string | 数字签名   |    | |
 
  响应数据:
 
 | 参数名称   | 是否必须 | 数据类型   | 描述     | 取值范围     |
 | ---------- | ---- | ------- | ------- | ------  |
 | success    | true | boolean |     | true 或者 false |
 | returnCode | true | string  | 响应返回码 如：200     |      |
 | returnMsg  | true | string  | 响应返回信息     |    |
 | data       | true | object  | 响应数据 |    |
 
 data 说明:
 
 ```
   "data": 
     {
       "coinType": 币种,
       "address": 充币地址,
       "tag": 充币标志（XRP币充值时需要）
     }
 ```
 
 
 请求示例
 
 ```
 /* POST /v1/account/deposit/address */
 {
    "coinTypes":"bch,btc,usdt",
    "timestamp":1535955784075,
    "accessKey":"4f47445f-efc3-4063-ae56-1165b357c747",
    "sign":"b210c02f58c193176515020bc7ff332b00efd55285b12df3cdfe06e8253c0490"
 }
 ```
 
 响应示例
 
 ```
 /* POST /v1/account/deposit/address */
 {
   "success": true,
    "returnCode": "200",
    "returnMsg": "success",
     "data": [
        {
            "coinType": "usdt",
            "address": "16MFa7tExeZ5DmUDSTj2xEfaSKiaFWJtUJ",
            "tag": ""
        },
        {
            "coinType": "btc",
            "address": "34fzJzMCHLvP6GcWTWxK2LjYMPApr7XgLS",
            "tag": ""
        },
        {
            "coinType": "bch",
            "address": "qpanyy4su8djqj8mwxtc3w57gukr65ffqqlgcv7z80",
            "tag": ""
        }
    ]
 }
 ```
