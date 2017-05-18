package com.zhiweicloud.guest.processor;

import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.RESCODE;
import com.wyun.thrift.server.Request;
import com.wyun.thrift.server.Response;
import com.zhiweicloud.guest.service.IBusinessService;
import org.apache.thrift.TException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.nio.ByteBuffer;

@Service
public class AirportServiceImpl implements MyService.Iface {
    private final
    IBusinessService iBusinessService;

    @Autowired
    public AirportServiceImpl(IBusinessService iBusinessService) {
        this.iBusinessService = iBusinessService;
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
        } catch (Exception ignored) {
        }

        String result;
        result = iBusinessService.handle(paramJSON);
        byte[] resultBytes = result.getBytes();
        ByteBuffer returnByteBuffer = ByteBuffer.allocate(resultBytes.length);
        returnByteBuffer.put(resultBytes);
        returnByteBuffer.flip();
        return new Response(RESCODE._200, returnByteBuffer);
    }
}
