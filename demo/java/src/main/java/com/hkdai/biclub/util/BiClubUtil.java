package com.hkdai.biclub.util;

import java.util.HashMap;
import java.util.Map;

/**
 * TODO 描述功能
 *
 * @author: daihk
 * @date: 2018/8/14 14:59
 * @version:1.0
 * @since:1.0
 */
public class BiClubUtil {

    static final String ORDER_URL = "https://api.biclub.com/api/trade/order/orders/place";

    static final String accessKey = "";

    static final String secretKey = "";

    public static void main(String[] args) {
        try {
            for (int i = 0; i < 50; i++) {
                placeOrder();
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static void placeOrder() {
        Map<String, String> paramMap = new HashMap<>();
        paramMap.put("number", "10");
        paramMap.put("orderType", "buy-limit");
        paramMap.put("price", "0.001");
        paramMap.put("source", "api");
        paramMap.put("symbol", "bz-usdt");
        paramMap.put("timestamp", System.currentTimeMillis() + "");
        paramMap.put("accessKey", accessKey);

        paramMap = MapSort.sortMapByKey(paramMap); // 按Key进行排序，根据首字母hashcode
        paramMap.put("sign", SignatureUtil.getSignByMap(paramMap, secretKey));

        String result = HttpUtil.doPost(ORDER_URL, paramMap);
        System.out.println(result);
    }

}
