1. HSET key(表示名字)  键 值：

   ```reids
   HSET person name laoyang
   ```

2. HGET person name:

   获取键值对的值

3. HGETALL person：

   获取所有的键值对，键值对成对出现

4. HDEL Hash表 键：

   删除相应的键值对

5. HEXISTS 表 键：

   判断某个键值对是否存在（1存在 0不存在）

6. HKEYS Hash：

   获取哈希中的所有键

7. HLEN Hash：

   获取Hash中所有键值对的数量

