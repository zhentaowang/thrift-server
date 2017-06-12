package com.wyun.thrift.server.server;

import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.business.BusinessServiceMap;
import com.wyun.thrift.server.business.TestService;
import com.wyun.thrift.server.processor.WyunServiceImpl;
import org.apache.thrift.TProcessor;
import org.apache.thrift.protocol.TBinaryProtocol;
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

    public void setWyunServiceImpl(MyService.Iface wyunServiceImpl) {
        this.wyunServiceImpl = wyunServiceImpl;
    }

    private MyService.Iface wyunServiceImpl;

    public void setServerPort(int serverPort) {
        this.serverPort = serverPort;
    }

    ExecutorService executor = Executors.newSingleThreadExecutor();


    public void startServer() {
        try {
            TProcessor tprocessor = new MyService.Processor<MyService.Iface>(wyunServiceImpl);
            TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(serverPort);
            serverTransport.registerSelector(Selector.open());
            TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
            tArgs.processor(tprocessor);
            tArgs.transportFactory(new TFramedTransport.Factory());
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
        BusinessServiceMap businessServiceMap = new BusinessServiceMap();
        businessServiceMap.registerService("testService",new TestService());
        WyunServiceImpl wyunServiceImpl = new WyunServiceImpl(businessServiceMap);
        server.setWyunServiceImpl(wyunServiceImpl);
        server.startServer();
    }
}
