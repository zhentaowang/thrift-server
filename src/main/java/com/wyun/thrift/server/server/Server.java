package com.wyun.thrift.server.server;

import com.wyun.thrift.server.MyService;
import com.wyun.thrift.server.business.BusinessServiceMap;
import com.wyun.thrift.server.business.TestService;
import com.wyun.thrift.server.processor.WyunServiceImpl;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TThreadedSelectorServer;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TNonblockingServerSocket;

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

    public void setProcessor(MyService.Processor processor) {
        this.processor = processor;
    }

    private MyService.Iface processor;

    public void setServerPort(int serverPort) {
        this.serverPort = serverPort;
    }

    ExecutorService executor = Executors.newSingleThreadExecutor();


    public void startServer() {
        try {
            TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(serverPort);
            // Selector这个类，是不是很熟悉。
//            serverTransport.registerSelector(Selector.open());
            TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
            tArgs.processor(processor);
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
        BusinessServiceMap businessServiceMap=new BusinessServiceMap();
        WyunServiceImpl wyunServiceImpl=new WyunServiceImpl(businessServiceMap);
        server.setProcessor(wyunServiceImpl);
        server.startServer();
    }
}
