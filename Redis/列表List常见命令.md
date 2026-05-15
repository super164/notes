1. LPUSH letter a:

   添加元素到队列头部中去

   LPUSH letter a b c：

   一次性添加多个元素

2. LRANGE letter 0 -1

   查看列表中的元素

3. RPUSH letter f:

   添加元素到列表的尾部

4. RPOP letter 【删除元素个数】:

   删除列表尾部中的最后一个元素

5. LPOP letter 【删除元素个数】:
   删除列表头部的第一个元素

6. LLEN letter:

   查看列表中的元素个数

7. LTRIM letter start stop:

   删除start到stop范围以外的元素，即只保留范围内的元素