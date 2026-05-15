# bitmap

1. SETBIT:

   设置某个偏移量的值

   ```
   SETBIT dianzan 0 1 设置0位置处为1
   SETBIT dianzan 1 0 设置1位置处为0
   ```

   

2. GETBIT :

   ```
   GETBIT dianzan 0 获取0偏移量的值
   ```

   

3. SET：

   一次性设置多个偏移量

   ```
   SET dianzan "\xF0" 运用十六进制数设置 八位数钱四位是一，后四位是零
   ```

   

4. BITCOUNTL:

   获取某一个KEY的值里面有多少个bit是1

   ```
   BITCOUNT dianzan 
   ```

   

5. BITPOS :

   获取第一次出现0/1的位置

   ```
   BITPOS dianzan 0 获取第一次0的位置
   BITPOS dianzan 1 获取第一次1的位置
   ```

   