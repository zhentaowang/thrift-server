协议
```java
TNonblockingServerSocket serverTransport = new TNonblockingServerSocket(serverPort);    //非阻塞式服务
TThreadedSelectorServer.Args tArgs = new TThreadedSelectorServer.Args(serverTransport);
tArgs.processor(tMultiplexedProcessor);                 //multiple processor，多处理器
tArgs.transportFactory(new TFramedTransport.Factory()); //设置窗口
tArgs.protocolFactory(new TBinaryProtocol.Factory());   //二进制传输
TServer server = new TThreadedSelectorServer(tArgs);
```
客户端与服务端对应
```java
TTransport transport = new TFramedTransport(new TSocket(host, port));   //设置窗口
transport.open();
TProtocol protocol = new TBinaryProtocol(transport);    //二进制传输
//use multiplexed protocol
TMultiplexedProtocol multiplexedProtocol = new TMultiplexedProtocol(protocol, serviceName); //设置处理器名称
MyService.Client client = new MyService.Client(multiplexedProtocol);
```
服务端使用
```java
Server server = new Server(9099);
server.registerService("testService", new TestService());
server.startServer();
```
- spring 配置
```xml
todo
```
客户端使用
- 提供短连接调用方式
```java
ByteBuffer byteBuffer = ByteBuffer.wrap("{\"key\":\"abcde\", \"int\":100}".getBytes());
Response response = ShortConnectionClient.send("localhost", 9099, new Request(byteBuffer, "testService"));
System.out.println(response);
```