package com.zhiweicloud.guest.common;

import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.Request;
import com.wyun.thrift.server.Response;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

import java.nio.ByteBuffer;
import java.nio.CharBuffer;
import java.nio.charset.Charset;
import java.util.Map;

/**
 * call service util class for thrift
 * Copyright(C) 2017 杭州风数信息技术有限公司
 *
 * 2017/5/10 13:38
 * @author tiecheng
 */
public class ThriftClientUtils {

    public static String invokeRemoteMethodCallBack(Map<String,Object> params, String serviceName,int port) {
        String result = null;
        try {
            String hostName = "localhost";
            TTransport transport = new TFramedTransport(new TSocket(hostName,
                    port, 1000));
            transport.open();
            TProtocol protocol = new TBinaryProtocol(transport);

            MyService.Client client = new MyService.Client(protocol);
            Request request = new Request();
            request.setServiceName(serviceName);
            String strParams = JSONObject.toJSONString(params);
            request.setParamJSON(byteChangeHelper.getBytes(strParams.toCharArray()));
            Response response = client.send(request);
            result = new String(byteChangeHelper.getChars(response.getResponseJSON()));
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        return result;
    }

    public static String invokeRemoteMethodCallBack(Map<String,Object> params, String serviceName) {
        String result = null;
        try {
            int port = 80;
            String hostName = serviceName;
            TTransport transport = new TFramedTransport(new TSocket(hostName,
                    port, 1000));
            transport.open();
            TProtocol protocol = new TBinaryProtocol(transport);

            MyService.Client client = new MyService.Client(protocol);
            Request request = new Request();
            request.setServiceName(serviceName);
            String strParams = JSONObject.toJSONString(params);
            request.setParamJSON(byteChangeHelper.getBytes(strParams.toCharArray()));
            Response response = client.send(request);
            result = new String(byteChangeHelper.getChars(response.getResponseJSON()));
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        return result;
    }

    static class byteChangeHelper{

        private static byte[] getBytes(char[] chars){
            Charset cs = Charset.forName ("UTF-8");
            CharBuffer cb = CharBuffer.allocate (chars.length);
            cb.put (chars);
            cb.flip ();
            ByteBuffer bb = cs.encode (cb);
            return bb.array();
        }

        private static char[] getChars (byte[] bytes) {
            Charset cs = Charset.forName ("UTF-8");
            ByteBuffer bb = ByteBuffer.allocate (bytes.length);
            bb.put (bytes);
            bb.flip ();
            CharBuffer cb = cs.decode (bb);
            return cb.array();
        }

    }

}
