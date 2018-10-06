#!/usr/bin/env python
# -*- coding: utf-8 -*-
# reference http://docs.python-requests.org/zh_CN/latest/user/quickstart.html
import hashlib
import json
import time
import requests

REQUEST_URL = "https://api.biclub.com"
ACCESS_kEY = ""
SECRET_KEY = ""
HEADERS = {"Accept": "application/json, text/plain, */*",
           "Content-Type": "application/json;charset=utf-8",
           "User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0"}
TIMEOUT = 5

#create sign
def create_sign(params, secret_key):
    # Sorted by params key ASCII from small to large
    signed_str = ""
    sorted_params = sorted(params.items(), key=lambda d: d[0], reverse=False)
    for key, value in sorted_params:
        signed_str += str(key) + str(value)
    #add secret key
    signed_str += secret_key
    #replace single quotes with double quotes, and remove Spaces to prevent signing incorrectly when value is a list type
    signed_str = signed_str.replace("'", "\"")
    signed_str = ''.join(signed_str.split())
    #create sign by sha256
    sha256 = hashlib.sha256()
    sha256.update(signed_str)
    sign = sha256.hexdigest()
    return sign

def do_https_post(req_path, params, other_headers = None):
    url = REQUEST_URL + req_path
    #13-bit millisecond eg: 1536225580880
    timestamp = int(round(time.time() * 1000))
    params.update({
        "accessKey":ACCESS_kEY,
        "timestamp":timestamp
        })
    sign = create_sign(params, SECRET_KEY)
    #add sign
    params.update({
        "sign":sign
    })
    params = json.dumps(params)
    if other_headers:
        HEADERS.update(other_headers)
    try:
        response = requests.post(url, data = params, headers = HEADERS, timeout = TIMEOUT)
        if 200 == response.status_code:
            return response.json()
        else:
            return response.json()
    except Exception as e:
        print("oh, post exception is %s"%e)
        return {"status":"failed", "exception":e}
'''
parameter signature is not required for all get requests
eg:
@ reqPath /api/market/trades
@ params {"symbol":"bch-usdt", "size":"5"}
URL is  https://api.biclub.com/api/market/trades?symbol=bch-usdt&size=5
'''
def do_https_get(req_path, params = None, other_headers = None):
    url = REQUEST_URL + req_path
    if other_headers:
        HEADERS.update(other_headers)
    try:
        response = requests.get(url, params = params, headers = HEADERS, timeout = TIMEOUT)
        if 200 == response.status_code:
            return response.json()
        else:
            return {"getStatus":"failed"}
    except Exception as e:
        print("oh, get exception is %s" %e)
        return {"status":"failed", "exception": e}


