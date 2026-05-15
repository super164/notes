# Redis Stream

Redis Stream 主要用于消息队列（MQ，Message Queue），Redis 本身是有一个 Redis 发布订阅 (pub/sub) 来实现消息队列的功能，但它有个缺点就是消息无法持久化，如果出现网络断开、Redis 宕机等，消息就会被丢弃。

简单来说发布订阅 (pub/sub) 可以分发消息，但无法记录历史消息。

而 Redis Stream 提供了消息的持久化和主备复制功能，可以让任何客户端访问任何时刻的数据，并且能记住每一个客户端的访问位置，还能保证消息不丢失。

**消息队列相关命令：**

- **XADD** - 添加消息到末尾
- **XTRIM** - 对流进行修剪，限制长度
- **XDEL** - 删除消息
- **XLEN** - 获取流包含的元素数量，即消息长度
- **XRANGE** - 获取消息列表，会自动过滤已经删除的消息
- **XREVRANGE** - 反向获取消息列表，ID 从大到小
- **XREAD** - 以阻塞或非阻塞方式获取消息列表

**消费者组相关命令：**

- **XGROUP CREATE** - 创建消费者组
- **XREADGROUP GROUP** - 读取消费者组中的消息
- **XACK** - 将消息标记为"已处理"
- **XGROUP SETID** - 为消费者组设置新的最后递送消息ID
- **XGROUP DELCONSUMER** - 删除消费者
- **XGROUP DESTROY** - 删除消费者组
- **XPENDING** - 显示待处理消息的相关信息
- **XCLAIM** - 转移消息的归属权
- **XINFO** - 查看流和消费者组的相关信息；
- **XINFO GROUPS** - 打印消费者组的信息；
- **XINFO STREAM** - 打印流信息

### XADD

使用 XADD 向队列添加消息，如果指定的队列不存在，则创建一个队列，XADD 语法格式：

```
XADD key ID field value [field value ...]
XADD key 1-0 course field value:手动输入id，前一个数字代表时间戳，后一个整数表示一个序列号
```

- **key** ：队列名称，如果不存在就创建
- **ID** ：消息 id，我们使用 * 表示由 redis 生成，可以自定义，但是要自己保证递增性。
- **field value** ： 记录

### XRANGE

获取消息列表，会自动过滤已经删除的消息



### XTRIM 

删除消息：XTRIM futao MAXLEN 0 表示删除所有消息



### XREAD

```
XREAD COUNT 读取数量 BLOCK 1000 STREAMS 队列名 起始位置
XREAD COUNT 读取数量 BLOCK 1000 STREAMS 队列名 ￥ ：获取从现在开始以最新的消息
```

以阻塞或非阻塞方式获取消息列表



### XGROUP

创建消费者组命令

```
XGROUP CREATE 消息名称 组的名称 ID
```



### XINFO GROUPS

查看消费者组的信息

```
XINFO GROUPS 消息队列名称
```



### XGROUP CREATECONSUMER 

添加消费者
```
XGROUP CREATECONSUMER 消息名称 组的名称 消费者的名字
```



### **XREADGROUP GROUP**

- 读取消费者组中的消息

  ```
  XREADGROUP GROUP 消费组的名称 消费者 COUNT 读取消息数量 BLOCK 阻塞时间
  STREAMS mystream >
  ```

  