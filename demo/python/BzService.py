#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author  kingler  2018/9/26
# reference https://github.com/biclubadmin/biclub-api/wiki/API-Reference
from BzUtil import *


'''
Palace orders

:URL https://api.biclub.com/api/trade/order/orders/place
:method post

:param number string
:param orderType string
:param price  string
:param source string
:param symbol string
'''
def post_place_order(number, order_type, price, source, symbol):
    req_path = "/api/trade/order/orders/place"
    params = {"number":number,
              "orderType":order_type,
              "price":price,
              "source":source,
              "symbol":symbol}
    return do_https_post(req_path, params)

'''
Query the recharge address

:URL https://api.biclub.com/api/trade/account/deposit/address
:method post

:param coinTypes
'''
def post_deposit_address(coin_types):
    req_path = "/api/trade/account/deposit/address"
    params = {"coinTypes":coin_types}
    return do_https_post(req_path, params)

'''
Query the account balance list API

:URL https://api.biclub.com/api/trade/account/balance/list
:method post

:param coinTypes string
'''
def post_account_balance_list(coin_types):
    req_path = "/api/trade/account/balance/list"
    params = {"coinType":coin_types}
    return do_https_post(req_path, params)

'''
Query the consignation orders list api

:URL https://api.biclub.com/api/trade/order/orders/consigns
:method post

:param symbol string
:param direct string
:param status string
:param size  int
'''
def post_orders_consigns(symbol, direct, status, size):
    req_path = "/api/trade/order/orders/consigns"
    params = {"symbol":symbol,
              "direct":direct,
              "status":status,
              "size":size}
    return do_https_post(req_path, params)

'''
Query the specified consignation order api

:URL https://api.biclub.com/api/trade/order/orders/consign
:method post

:param orderId string
:param symbol string
'''
def post_order_consign(order_id, symbol):
    req_path = "/api/trade/order/orders/consign"
    params = {"orderId":order_id,
              "symbol":symbol}
    return do_https_post(req_path, params)

'''
Batch cancel the delegate orders

:URL https://api.biclub.com/api/trade/order/orders/batchcancel
:method post

:param orderIdList list  eg ["bebced7c-b97c-4493-88d5-XXXX","1ff53193-76ab-4eb9-89b8-XXXX"]
:param symbol string
'''
def post_orders_batch_cancel(order_id_list, symbol):
    req_path = "/api/trade/order/orders/batchcancel"
    params = {"orderIdList":order_id_list,
              "symbol":symbol}

    return do_https_post(req_path, params)

'''
Cancel the delegate order by order id and symbol one by one

:URL https://api.biclub.com/api/trade/order/orders/cancel
:method post

:param orderId string  eg: "bebced7c-b97c-4493-88d5-XXXX"
:param symbol string
'''
def post_delegated_order_cancel(order_id, symbol):
    req_path = "/api/trade/order/orders/cancel"
    params = {"orderId":order_id,
              "symbol":symbol}

    return do_https_post(req_path, params)


'''
Query latest transaction record API

:URL eg: https://api.biclub.com/api/market/trades?symbol=bz-usdt&size=20
:method get

:param  symbol string  eg: bch-usdt, bz-usdt
:param  size  int  1-200, default 20
'''
def get_trade2_deal_history(symbol, size):
    req_path = "/api/market/trades"
    params = {"symbol":symbol,
              "size":size}
    return do_https_get(req_path,params)

'''
Query k line transaction pair info

:URL eg: https://api.biclub.com/api/market/history/kline?symbol=btc-usdt&period=30&size=3
:method get

:param  symbol string  eg: bch-usdt, bz-usdt
:param  period string  eg: 1,15,30,60,D,2D,3D,W,3W,M,6M
:param  size   int     default is 150
'''
def get_kline_transpair_info(symbol, period, size):
    req_path = "/api/market/history/kline"
    params = {"symbol":symbol,
              "period":period,
              "size":size}
    return do_https_get(req_path,params)

'''
Query latest transaction record api

:URL eg: https://api.biclub.com/api/market/trades?symbol=bz-usdt&size=10
:method get

:param  symbol string  eg: bch-usdt, bz-usdt
:param  size   int     default is 20
'''
def get_latest_tran_record(symbol, size):
    req_path = "/api/market/trades"
    params = {"symbol":symbol,
              "size":size}
    return do_https_get(req_path,params)

'''
Query order depth

:URL eg: https://api.biclub.com/api/market/depth?symbol=btc-usdt
:method get

:param  symbol string  eg: bch-usdt, bz-usdt
'''
def get_order_depth(symbol):
    req_path = "/api/market/depth"
    params = {"symbol":symbol}
    return do_https_get(req_path,params)

'''
Query market tickers

:URL eg: https://api.biclub.com/api/market/tickers?symbol=bz-usdt
:method get

:param  symbol string  eg: bch-usdt, bz-usdt
'''
def get_market_tickers(symbol):
    req_path = "/api/market/tickers"
    params = {"symbol":symbol}
    return do_https_get(req_path,params)

'''
Query latest transaction pair

:URL eg: https://api.biclub.com/api/market/common/symbols
:method get

'''
def get_transaction_pair():
    req_path = "/api/market/common/symbols"
    return do_https_get(req_path)

if __name__ == '__main__':

    #print post_deposit_address("bz")
    #print post_account_balance_list("bz")
    #print post_orders_consigns("bz-usdt", 1, 0, 20)
    #print post_order_consign("ceecca34-7941-479d-a467-895042f1e606","bz-usdt")
    #aa = ["87068d96-7130-40ed-a4ae-7ad476d4d357", "68206740-7b8f-4232-85d6-e309154aab52","68206740-7b8f-4232-85d6-e309154aab52"]
    #print post_orders_batch_cancel(aa, "bz-usdt")
    #print post_delegated_order_cancel("87068d96-7130-40ed-a4ae-7ad476d4d357","bz-usdt")
    #print post_place_order(22, "sell-limit","7", "api","bz-usdt")
    #print get_kline_transpair_info("bz-usdt", "3D", 10)
    #print get_latest_tran_record("bz-usdt", 200)
    #print get_order_depth("bz-usdt")
    #print get_market_tickers("bz-usdt")
    #print get_transaction_pair()
    pass
