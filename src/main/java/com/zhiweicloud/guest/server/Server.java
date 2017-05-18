package com.zhiweicloud.guest.server;

import com.wyun.thrift.server.MyService;
import org.apache.thrift.TProcessor;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TThreadedSelectorServer;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TNonblockingServerSocket;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.nio.channels.Selector;

/**
 * Created by wzt on 05/05/2017.
 */
@Service
public class Server {
    public static int SERVER_PORT;

    private static MyService.Iface serverServiceImpl;

    @Autowired
    public Server(MyService.Iface servServiceImpl) {
        this.serverServiceImpl = servServiceImpl;
    }

    public static void startServer() {
        try {
            TProcessor tprocessor = new MyService.Processor<MyService.Iface>(serverServiceImpl);
            // 非阻塞异步通讯模型（服务器端）
            TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(SERVER_PORT);
            // Selector这个类，是不是很熟悉。
            serverTransport.registerSelector(Selector.open());
            //多线程半同步半异步
            TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
            tArgs.processor(tprocessor);
            tArgs.transportFactory(new TFramedTransport.Factory());
            //二进制协议
            tArgs.protocolFactory(new TBinaryProtocol.Factory());
            TServer server = new TThreadedSelectorServer(tArgs);
            server.serve();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
