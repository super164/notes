# Redis持久化

Redis中数据的持久化有两种方式RDB(Redis Database)和AOF(Append Only File)



# RDB

快照模式，适合用来做备份

在执行完命令之后，执行save后，即保存快照



bgsave会单独创建一个子进程，负责将内存中的数据写入到硬盘中



xxd.dump.rdb

xxd是一个可以查看2进制或16进制文件内容的Linux命令，可以看到rdb文件中的内容。

# AOF

​	在执行写命令的时候不仅会将命令写入到内存中，还会将命令写入到一个追加的文件中，AOF文件，以日志的形式在记录每一个写操作，Redis重启的时候，会重新执行AOF文件中的命令，加载数据库的内容到内存中

在AOF文件中，将appendonly no的‘’no‘‘改成"yes"