package com.wyun.thrift.server.processor;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.*;
import com.wyun.thrift.server.business.IBusinessService;
import com.wyun.utils.SpringBeanUtil;
import org.apache.thrift.TException;
import org.springframework.beans.factory.annotation.Autowired;

import java.nio.ByteBuffer;

/**
 * User:lpf
 * Date:2017/2/27
 * Time:18:56
 */
public class WyunServiceImpl implements MyService.Iface{
    @Autowired
    SpringBeanUtil springBeanUtil;

    @Override
    public Response send(Request request) throws ServiceException, TException {
        String serviceName = request.getServiceName();
        JSONObject paramJSON = null;
        try {
            byte [] paramJSON_bytes = request.getParamJSON();
            if(paramJSON_bytes != null && paramJSON_bytes.length > 0) {
                String paramJSON_string = new String(paramJSON_bytes);
                paramJSON = JSONObject.parseObject(paramJSON_string);
            }
        } catch(Exception e) {
        }
        IBusinessService businessService=null;
        try {
            businessService=  springBeanUtil.getBean(serviceName);
        } catch (Exception  e) {
            e.printStackTrace();
        }
        JSONObject result;
        result = businessService.handle(paramJSON);
        String resultString= JSON.toJSONString(result);
        byte[] resultBytes = resultString.getBytes();
        ByteBuffer returnByteBuffer = ByteBuffer.allocate(resultBytes.length);
        returnByteBuffer.put(resultBytes);
        returnByteBuffer.flip();
        Response response = new Response(RESCODE._200, returnByteBuffer);
        return response;
    }
}
