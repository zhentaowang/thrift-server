package com.wyun.thrift.server.server;

import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.business.IBusinessService;
import com.wyun.thrift.server.business.TestService;
import com.wyun.thrift.server.processor.MyServiceImpl;
import org.apache.thrift.TMultiplexedProcessor;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TCompactProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TThreadedSelectorServer;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TNonblockingServerSocket;

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

    public Server(int serverPort) {
        this.serverPort = serverPort;
    }

    private TMultiplexedProcessor tMultiplexedProcessor = new TMultiplexedProcessor();

    public void setServerPort(int serverPort) {
        this.serverPort = serverPort;
    }

    ExecutorService executor = Executors.newSingleThreadExecutor();

    public void registerService(String name, IBusinessService instance) {
        MyServiceImpl myService = new MyServiceImpl(instance);
        tMultiplexedProcessor.registerProcessor(name, new MyService.Processor<>(myService));
    }

    public void startServer() {
        try {
            TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(serverPort);
            // Selector这个类，是不是很熟悉。
//            serverTransport.registerSelector(Selector.open());
            TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
            tArgs.processor(tMultiplexedProcessor);
            tArgs.transportFactory(new TFramedTransport.Factory());
//            tArgs.executorService(Executors.newFixedThreadPool(20));
            tArgs.protocolFactory(new TBinaryProtocol.Factory());
            TServer server = new TThreadedSelectorServer(tArgs);
            System.out.println("Thrift Server start....");
            executor.execute(new Runnable() {
                @Override
                public void run() {
                    server.serve();
                }
            });
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        Server server = new Server(9099);
        server.registerService("testService", new TestService());
        server.startServer();
    }
}
