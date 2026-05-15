# Redis

## 什么是Redis

​	Redis（Remote Dictionary Server）是一个开源的内存数据库，遵守 BSD 协议，它提供了一个高性能的键值（key-value）存储系统，常用于缓存、消息队列、会话存储等应用场景。

- **性能极高：**Redis 以其极高的性能而著称，能够支持每秒数十万次的读写操作24。这使得Redis成为处理高并发请求的理想选择，尤其是在需要快速响应的场景中，如缓存、会话管理、排行榜等。
- **丰富的数据类型：**Redis 不仅支持基本的键值存储，还提供了丰富的数据类型，包括字符串、列表、集合、哈希表、有序集合等。这些数据类型为开发者提供了灵活的数据操作能力，使得Redis可以适应各种不同的应用场景。
- **原子性操作：**Redis 的所有操作都是原子性的，这意味着操作要么完全执行，要么完全不执行。这种特性对于确保数据的一致性和完整性至关重要，尤其是在高并发环境下处理事务时。
- **持久化：**Redis 支持数据的持久化，可以将内存中的数据保存到磁盘中，以便在系统重启后恢复数据。这为 Redis 提供了数据安全性，确保数据不会因为系统故障而丢失。
- **支持发布/订阅模式：**Redis 内置了发布/订阅模式（Pub/Sub），允许客户端之间通过消息传递进行通信。这使得 Redis 可以作为消息队列和实时数据传输的平台。
- **单线程模型：**尽管 Redis 是单线程的，但它通过高效的事件驱动模型来处理并发请求，确保了高性能和低延迟。单线程模型也简化了并发控制的复杂性。
- **主从复制：**Redis 支持主从复制，可以通过从节点来备份数据或分担读请求，提高数据的可用性和系统的伸缩性。
- **应用场景广泛：**Redis 被广泛应用于各种场景，包括但不限于缓存系统、会话存储、排行榜、实时分析、地理空间数据索引等。
- **社区支持：**Redis 拥有一个活跃的开发者社区，提供了大量的文档、教程和第三方库，这为开发者提供了强大的支持和丰富的资源。
- **跨平台兼容性：**Redis 可以在多种操作系统上运行，包括 Linux、macOS 和 Windows，这使得它能够在不同的技术栈中灵活部署。

## Redis与其他的key-value存储有什么不同

​	Redis 与其他 key-value 存储系统的主要区别在于其提供了丰富的数据类型、高性能的读写能力、原子性操作、持久化机制、以及丰富的特性集

- **丰富的数据类型：**Redis 不仅仅支持简单的 key-value 类型的数据，还提供了 list、set、zset（有序集合）、hash 等数据结构的存储。这些数据类型可以更好地满足特定的业务需求，使得 Redis 可以用于更广泛的应用场景。
- **高性能的读写能力：**Redis 能读的速度是 110000次/s，写的速度是 81000次/s。这种高性能主要得益于 Redis 将数据存储在内存中，从而显著提高了数据的访问速度。
- **原子性操作：**Redis 的所有操作都是原子性的，这意味着操作要么完全执行，要么完全不执行。这种特性对于确保数据的一致性和完整性非常重要。
- **持久化机制：**Redis 支持数据的持久化，可以将内存中的数据保存在磁盘中，以便在系统重启后能够再次加载使用。这为 Redis 提供了数据安全性，确保数据不会因为系统故障而丢失。
- **丰富的特性集：**Redis 还支持 publish/subscribe（发布/订阅）模式、通知、key 过期等高级特性。这些特性使得 Redis 可以用于消息队列、实时数据分析等复杂的应用场景。
- **主从复制和高可用性：**Redis 支持 master-slave 模式的数据备份，提供了数据的备份和主从复制功能，增强了数据的可用性和容错性。
- **支持 Lua 脚本：**Redis 支持使用 Lua 脚本来编写复杂的操作，这些脚本可以在服务器端执行，提供了更多的灵活性和强大的功能。
- **单线程模型：**尽管 Redis 是单线程的，但它通过高效的事件驱动模型来处理并发请求，确保了高性能和低延迟



## 持久化机制

​	Redis的持久化主要是由两个内容进行实现，RDB和AOF,通过这两个机制实现Redis持久化，两个可以说是各有千秋。

### 1. RDB(Redis Database)

- **快照方式**：定期将内存数据全量保存为二进制文件（dump.rdb）。

- 优点：恢复快、文件紧凑、适合备份。

- 缺点：可能丢失最后一次快照后的数据。

- 配置示例：

  ```redis
  save 900 1     # 900秒内至少1个key变化则触发
  save 300 10
  save 60 10000
  ```

### **2. AOF（Append Only File）**

- **日志方式**：记录每个写命令，重启时重放命令恢复数据。

- 优点：数据更安全（可配置 fsync 策略）。

- 缺点：文件大、恢复慢。

- 配置：

  ```redis
  appendonly yes
  appendfsync everysec  # 可选 always / everysec / no
  ```



## 常用的Redis CLI工具

常用命令：

```
redis-cli
redis-cli --raw #可以输出中文
redis-cli -h 127.0.0.1 -p 6379
redis-cli -a yourpassword
redis-cli ping          # 测试连接
redis-cli info memory   # 查看内存信息
redis-cli --scan        # 批量获取 key（替代 KEYS *）
```



## 常用的数据类型以及操作命令

### Key键

1. DEL key :该命令用于在key存在时删除key
2. DUMP key:序列化给定key，并返回被序列化的值。
3. EXISTS key:检查给定key是否存在
4. EXPIRE key seconds:给定key设置过期时间，以秒。
5. PEXPIRE key milliseconds: 设置key的过期时间以毫秒
6. KEYS * :查询redis中所有的键
7. PERSIST key：移除key的过期时间，key持久保存
8. （p）TLL key:查看相应的key的过期时间(P毫米)
9. RENAME key newkey:修改key的名称
10. TYPE key:返回key所存储的值的类型

### String字符串

1. SET key value:设置指定key的值
2. GET key :获取指定key的值
3. GETRANGE key start end:返回key中字符串的子字符
4. MGET key1 [key2]:获取所有给定的key的值，查询多个key
5. SETNX key value:只有在key不存在时设置key的值
6. STRLEN key:返回key所存储的字符串值的长度
7. INCR key :将key中的存储数字加一
8. INCRBY key increment:将key存储的值加上指定的增量值（increment）
9. DECR key:将key中存储的数字减一
10. APPEND key value:如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。

### Hash 哈希

1. HSET key(表示名字)  键 值：

   ```reids
   HSET person name laoyang
   ```

2. HGET person name:获取键值对的值

3. HGETALL person：获取所有的键值对，键值对成对出现

4. HDEL Hash表 键：删除相应的键值对

5. HEXISTS 表 键：判断某个键值对是否存在（1存在 0不存在）

6. HKEYS Hash：获取哈希中的所有键

7. HLEN Hash：获取Hash中所有键值对的数量

### List列表

1. LPUSH letter a:添加元素到队列头部中去

   LPUSH letter a b c：一次性添加多个元素

2. LRANGE letter 0 -1查看列表中的元素

3. RPUSH letter f:添加元素到列表的尾部

4. RPOP letter 【删除元素个数】:删除列表尾部中的最后一个元素

5. LPOP letter 【删除元素个数】:删除列表头部的第一个元素

6. LLEN letter:查看列表中的元素个数

7. LTRIM letter start stop:删除start到stop范围以外的元素，即只保留范围内的元素

8. LREM key count value:移除列表元素

9. LSET key index value:通过索引设置列表元素的值

10. RPUSHX key value:为已存在的列表添加值

### Set集合

1. SADD key member1:向集合中添加元素

2. SCARD key:回去集合的元素数

3. SMEMBERS key:查看集合中的元素

4. 不能添加重复的元素

5. SISMEMBER key member:查看相应的元素是否存在集合中

6. SREM key member1:删除集合中的元素

7. SPOP key: 移除并返回集合中的一个随机元素

8. 集合中的交并差集运算：SINTER,SUNION,SDIFF

### Sorted set 有序集合

有序集合的相关命令都是以Z开头的

1. ZADD 集合名 分数 成员

   ```redis
   添加一个高考录取分数的有序集合
   ZADD result 680 清华 660 北大 650 复旦 640 浙大
   ```

2. ZRANGE 集合名 起始位置 结束位置：

   查看集合中的元素

3. ZRANGE 集合名 起始位置 结束位置 WITHSCORES:

   在输出集合元素的时候同时输出分数

4. ZSCORE 集合名 元素名：

   输出集合中指定元素的分数

5. ZRANK 集合名 元素名：（从小到大）

   该元素在此有序集合中的排序的位置

6. ZREVRANK  集合名 元素名：（从大到小）

7. ZREM 集合名 元素：

   删除集合中某位元素

### HyperLogLog

​	Redis HyperLogLog 是用来做基数统计的算法，HyperLogLog 的优点是，在输入元素的数量或者体积非常非常大时，计算基数所需的空间总是固定 的、并且是很小的。

1. PFADD key element:添加指定元素到 HyperLogLog 中。
2. PFCOUNT key:返回给定 HyperLogLog 的基数估算值。
3. PERMERGE destkey sourcekey:将多个HyperLogLog和并为一个HyperLogLog



## 事务和原子性操作

**事务（MULTI / EXEC）**

- 不支持回滚（即使出错也继续执行）
- 保证命令**顺序执行、不被其他客户端打断**

1. MULTI 开启事务
2. EXEC 提交事务
3. DISCARD： 取消事务，放弃执行事务块内的所有命令

**原子操作（天然原子）**

- 所有单条 Redis 命令都是原子的（如 `INCR`, `HSET`, `LPUSH`）



​	单个 Redis 命令的执行是原子性的，但 Redis 没有在事务上增加任何维持原子性的机制，所以 Redis 事务的执行并不是原子性的。

​	事务可以理解为一个打包的批量执行脚本，但批量指令并非原子化的操作，中间某条指令的失败不会导致前面已做指令的回滚，也不会造成后续的指令不做。



## Redis 发布订阅

Redis 发布订阅 (pub/sub) 是一种消息通信模式：发送者 (pub) 发送消息，订阅者 (sub) 接收消息。

Redis 客户端可以订阅任意数量的频道。

1. SUBSCRIBE channel ...：订阅一个或者多个频道信息
2. UNSUNCRIBE channel ...:退订给定的频道
3. PUBLISH channel message:将信息发送给指定的频道

## Redis GEO

Redis GEO 主要用于存储地理位置信息，并对存储的信息进行操作，该功能在 Redis 3.2 版本新增。

Redis GEO 操作方法有：

- geoadd：添加地理位置的坐标。
- geopos：获取地理位置的坐标。
- geodist：计算两个位置之间的距离。
- georadius：根据用户给定的经纬度坐标来获取指定范围内的地理位置集合。
- georadiusbymember：根据储存在位置集合里面的某个地点获取指定范围内的地理位置集合。
- geohash：返回一个或多个位置对象的 geohash 值。



## Redis Stream

​	Redis Stream 是 Redis 5.0 版本新增加的数据结构。Redis Stream 主要用于消息队列（MQ，Message Queue），Redis 本身是有一个 Redis 发布订阅 (pub/sub) 来实现消息队列的功能，但它有个缺点就是消息无法持久化，如果出现网络断开、Redis 宕机等，消息就会被丢弃。

​	发布订阅 (pub/sub) 可以分发消息，但无法记录历史消息。而 Redis Stream 提供了消息的持久化和主备复制功能，可以让任何客户端访问任何时刻的数据，并且能记住每一个客户端的访问位置，还能保证消息不丢失。



**常用的命令：**

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

  
