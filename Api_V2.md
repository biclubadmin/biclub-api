# API Reference

#### 

``` 
访问根路径：https://api.biclub.com/api/v2
symbol 规则： 基础币种+计价币种。如btc-usdt, bz-usdt, eth-usdt 以此类推
```


### API

API接口请求地址

例：

```
https://api.biclub.com/api/v2/market/tickers?symbol=bz-usdt
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
    "id": 消息id,
    "time": 统计时间,
    "number": 24小时成交量,
    "open": 前推24小时成交价,
    "last": 当前成交价,
    "maxPrice": 近24小时最高价,
    "minPrice": 近24小时最低价,
    "averagePrice": 近24小时平均价,
    "count": 近24小时累积成交数,
    "vol": 近24小时累积成交额, 即 sum(每一笔成交价 * 该笔的成交量)
  }
```


请求响应例子

```
/* GET /market/tickers?symbol=bz-usdt */
{
  "success": true,
  "returnCode": "200",
  "returnMsg": "success",
  "data": {
    "id": "6147e0ba-cf7f-43b5-9ff9-99f7c86d02b9",
    "minPrice": 0.010391,
    "maxPrice": 0.01983,
    "averagePrice": 0.014123666666666666,
    "number": 353.2,
    "count": 27,
    "vol": 4.9814457,
    "open": 0.012089,
    "last": 0.012089,
    "time": 1530150825104
  }
}

/* GET /market/tickers?symbol=not-exists */
{
  "success": false,
  "returnCode": "30024",
  "returnMsg": "交易对不存在",
  "data": null
}
```
