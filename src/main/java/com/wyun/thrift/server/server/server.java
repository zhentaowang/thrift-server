package com.wyun.thrift.server.server;

import com.wyun.thrift.server.MyService;
import com.wyun.utils.SpringBeanUtil;
import org.apache.thrift.TProcessor;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TThreadedSelectorServer;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TNonblockingServerSocket;

import javax.annotation.PostConstruct;
import java.nio.channels.Selector;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * User:lpf
 * Date:2017/2/27
 * Time:18:56
 */
public class Server {

    private int serverPort;

    private String processorName;

    public void setProcessorName(String processorName) {
        this.processorName = processorName;
    }

    public void setServerPort(int serverPort) {
        this.serverPort = serverPort;
    }

    ExecutorService executor = Executors.newSingleThreadExecutor();

    public void startServer() {
        try {
            MyService.Iface wyunServiceImpl= SpringBeanUtil.getBean(processorName);
            TProcessor tprocessor = new MyService.Processor<MyService.Iface>(wyunServiceImpl);
            // 非阻塞异步通讯模型（服务器端）
            TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(serverPort);
            // Selector这个类，是不是很熟悉。
            serverTransport.registerSelector(Selector.open());
            //多线程半同步半异步
            TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
            tArgs.processor(tprocessor);
            tArgs.transportFactory(new TFramedTransport.Factory());
            //二进制协议
            tArgs.protocolFactory(new TBinaryProtocol.Factory());
            TServer server = new TThreadedSelectorServer(tArgs);
            System.out.println("fund Server start....");
            server.serve();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @PostConstruct
    public void init() {
        executor.execute(new Runnable() {
            @Override
            public void run() {
                startServer();
            }
        });
    }
}
