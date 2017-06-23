package com.wyun.thrift.server.business;

import com.alibaba.fastjson.JSONObject;

/**
 * Created by doctor on 17-6-6.
 */
public class TestService implements IBusinessService {
    @Override
    public JSONObject handle(String operation, JSONObject request) {
        String key = request.getString("key");
        int i = request.getInteger("int");
        return JSONObject.parseObject("{\"key\":\"" + key + "\",\"value\":" + i + "}");
    }
}
