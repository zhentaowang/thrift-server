package com.wyun.thrift.server.processor;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.*;
import com.wyun.thrift.server.business.BusinessServiceMap;
import com.wyun.thrift.server.business.IBusinessService;
import org.apache.thrift.TException;

import java.nio.ByteBuffer;

/**
 * User:lpf
 * Date:2017/2/27
 * Time:18:56
 */
public class WyunServiceImpl implements MyService.Iface {
    private BusinessServiceMap businessServiceMap;

    public WyunServiceImpl(BusinessServiceMap businessServiceMap) {
        this.businessServiceMap = businessServiceMap;
    }

    @Override
    public Response send(Request request) throws TException {
        JSONObject paramJSON = null;
        try {
            byte[] paramJSON_bytes = request.getParamJSON();
            if (paramJSON_bytes != null && paramJSON_bytes.length > 0) {
                String paramJSON_string = new String(paramJSON_bytes);
                paramJSON = JSONObject.parseObject(paramJSON_string);
            }
        } catch (Exception e) {
        }
        JSONObject result = businessServiceMap.handle(request.serviceName, "", paramJSON);
        String resultString = JSON.toJSONString(result);
        byte[] resultBytes = resultString.getBytes();
        ByteBuffer returnByteBuffer = ByteBuffer.allocate(resultBytes.length);
        returnByteBuffer.put(resultBytes);
        returnByteBuffer.flip();
        return new Response(RESCODE._200, returnByteBuffer);
    }
}
