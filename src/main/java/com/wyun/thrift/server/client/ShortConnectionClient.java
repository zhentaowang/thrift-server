package com.wyun.thrift.server.client;

import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.Request;
import com.wyun.thrift.server.Response;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TMultiplexedProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;
import org.apache.thrift.transport.TTransportException;

import java.nio.ByteBuffer;

/**
 * Created by doctor on 17-6-6.
 */
public class ShortConnectionClient {
    public static Response send(String host, int port, Request request) {
        if (host == null || request == null) {
            return null;
        }
        String serviceName = request.getServiceName();
        TTransport transport = new TFramedTransport(new TSocket(host, port));
        try {
            transport.open();
        } catch (TTransportException e) {
            e.printStackTrace();
        }
        TProtocol protocol = new TBinaryProtocol(transport);
        MyService.Client client = new MyService.Client(protocol);
        Response response = null;
        try {
            response = client.send(request);
        } catch (TException e) {
            e.printStackTrace();
        }
        transport.close();
        return response;
    }
    public static void main(String[] args) {
        ByteBuffer byteBuffer = ByteBuffer.wrap("{\"key\":\"abcde\", \"int\":100}".getBytes());
        Response response = ShortConnectionClient.send("localhost", 9002, new Request(byteBuffer, "testService","test"));
        System.out.println(response);
    }

}
