package com.wyun.thrift.server.business;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.RESCODE;
import com.wyun.thrift.server.Response;

import java.nio.ByteBuffer;

/**
 * Created by doctor on 17-6-6.
 */
public class TestService implements IBusinessService {
    @Override
    public Response handle(String operation, JSONObject request) {
        String key = request.getString("key");
        int i = request.getInteger("int");
        JSONObject result= JSONObject.parseObject("{\"key\":\"" + key + "\",\"value\":" + i + "}");
        String resultString = JSON.toJSONString(result);
        byte[] resultBytes = resultString.getBytes();
        ByteBuffer returnByteBuffer = ByteBuffer.allocate(resultBytes.length);
        returnByteBuffer.put(resultBytes);
        returnByteBuffer.flip();
        return new Response(RESCODE._200, returnByteBuffer);
    }
}
