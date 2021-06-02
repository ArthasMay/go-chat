### 完成一个聊天系统

#### 连接层的设计(connect)

Server: 对Socket/WebSocket服务器的抽象

Buckets: 减少锁的竞争，会在connect层分bucket处理

Room & Pair: 对房间和P2PChat的抽象，room中的连接以双向链表的形式存储

Channel: Client 和 Server之间连接的抽象


