package com.hkdai.biclub.util;

import com.alibaba.fastjson.JSON;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.lang.reflect.Field;
import java.math.BigInteger;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;

/**
 * 签名工具
 *
 * @author zhaoc
 */
public class SignatureUtil {

    private static final Logger LOGGER = LoggerFactory.getLogger(SignatureUtil.class);

    // 加密key
    public static String screatKey;



    public static String getSignByMap(Map<String, String> map, String key) {
        LOGGER.info("getSignByMap map:{}", JSON.toJSONString(map));

        ArrayList<String> list = new ArrayList<String>();
        for (Map.Entry<String, String> entry : map.entrySet()) {
            if (entry.getValue() != "") {
                list.add(entry.getKey() +  entry.getValue() );
            }
        }

        int size = list.size();
        String[] arrayToSort = list.toArray(new String[size]);
        Arrays.sort(arrayToSort, String.CASE_INSENSITIVE_ORDER);
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < size; i++) {
            sb.append(arrayToSort[i]);
        }

        String result = sb.toString();
        result += key;
        LOGGER.info("Sign Before MD5:{}", result);
        result = SHAUtils.String2SHA256(result);
        LOGGER.info("Sign Result:{}", result);

        return result;
    }



    public void setScreatKey(String screatKey) {
        SignatureUtil.screatKey = screatKey;
    }
}
