本章主要是对于学习Linux和Docker的知识点进行总结，一些常用的命令操作以及容易忘记的命令操作

# Linux

## 文件与目录操作

对于文件与目录的一些相关操作命令，进行使用的说明

### 列出目录内容(ls)

```
ls （列出当前目录）
ls -l （以长格式列出，显示详细信息）
ls -a （列出所有文件，包括隐藏文件，以.开头的文件）
ls /home （列出指定目录的内容）
```



### 切换当前工作目录(cd)

```
cd /home/username/Documents （切换到绝对路径）
cd .. （切换到上级目录）
cd ~ 或 cd （切换到当前用户的家目录）
cd - （切换到上一个所在的目录）
```



### 显示当前所在目录的绝对路径（pwd）

```
$ pwd
/home/username
```



### 创建新目录（mkdir）

```
mkdir new_folder
mkdir -p parent/child/grandchild （递归创建多级目录）
```



### touch

创建新的空文件

更新文件的访问和修改时间戳

```
touch file.txt （如果 file.txt 不存在则创建，存在则更新其时间戳）
```



### cp

复制文件或目录

```
cp file1.txt file2.txt （复制 file1.txt 为 file2.txt）
cp file1.txt /tmp/ （将 file1.txt 复制到 /tmp 目录下）
cp -r dir1 dir2 （递归复制目录及其所有内容）
```



### mv

移动文件或目录

```
mv file.txt /home/username/Documents/ （移动文件）
```



重命名文件或目录

```
mv oldname.txt newname.txt（重命名）
```



### rm 

删除文件或目录

```
rm file.txt （删除文件）
rm -r my_directory （递归删除目录及其所有内容）
rm -f file.txt （强制删除，不提示）【非常危险！】
```

删除空目录

```
rmdir empty_dir
```



### which/find

which命令：可以查看锁使用的一些列的命令的程序文件放在哪里

```
which 要查找的命令
例如
which touch
```



find命令：可以去搜索指定的文件

```
按文件名进行查找
find 起始地址 -name "被查找文件名"

按文件大小进行查找:
find 起始路径 -size +/-n[kMG]
例如：
- 查找小于10KB的文件：find / -size -10kb

- 查找大于100MB的文件：find / -size +100M

- 查找大于1GB的文件：find / -size +1G
```



## 文件内容查看与编辑

为查看文件进行的命令操作

### cat

连接文件并打印到标准输出（通常是屏幕），适合查看短文件

```
cat file.txt(查看文件内容)
cat file1.txt file2.txt > combined.txt(合并两个文件)
```



### less

分页查看文件内容，适合长文件

```
less long_file.log
在 less 中：
- 空格键： 向下翻一页
- b： 向上翻一页
- /keyword： 搜索
- q： 退出
```



### head/tail

分别显示文件的开头和末尾部分，可以用来查看日志（一般默认是十行）

```
显示文件的开头部分
head file.log
head -n 20 file.log （显示前20行）

显示文件的末尾部分
tail file.log
tail -n 20 file.log （显示最后20行）
tail -f file.log （实时追踪文件的新增内容，监控日志神器）
```



### grep

从文件中通过关键字进行过滤文件行

```
grep [-n] 关键字 文件路径 
```

- 选项-n,可选，表示在结果中显示匹配的行的行号
- 参数，关键字，必填，表示过滤的关键字，带有空格或其他特俗符号，建议使用`“”`将关键字包围起来
- 参数，文件路径，必填，表示要过滤内容的文件路径，可作为内容输入端口（可配合管道符进行使用）



### wc 

对文件内容的行数，单词数量等进行统计

```
wc -c -m -l -w 文件路径
实例
wc -cmlw /home/1.txt
统计1.txt文件中的-c:bytes数量 -m:字符数量 -l:行数 -w:单词数量
```



## 进程管理

查看系统中的进程以及内存占用情况

### ps

显示当前进程

```
ps （显示当前用户的进程）
ps aux （显示所有用户的详细进程信息，最常用）
```



### top/htop

动态、实时地显示系统进程和资源占用情况，类似任务管理器

htop需要进行安装

```
top
```



### kill 

用于终止进程

```
kill 1234 （终止 PID 为 1234 的进程，发送默认的 TERM 信号）
kill -9 1234 （强制终止进程，使用 KILL 信号，无法被进程捕获或忽略）
```



### df 

显示磁盘空间使用情况

```\
df -h （以人类易读的格式，如 G, M 显示）
```



### du

显示文件或目录的磁盘使用情况

```
du -sh /home/username （以人类易读的格式汇总（-s）显示指定目录的大小）
```



## 权限管理

Linux中每个文件和目录都有所有者、所属组和其他的rwx(读，写，执行)权限

对于文件和目录的权限管理命令，如下

### chmod

改变文件或目录的权限

```
数字模式:
chmod 755 script.sh （所有者：rwx，组：r-x，其他人：r-x）
符号模式：
chmod u+x script.sh （给所有者增加执行权限）
```



### chown

改变文件或目录的所有者和所属组

```
chown username file.txt （改变所有者）
chown username:groupname file.txt （同时改变所有者和所属组）
chown -R username:groupname /path/to/dir （递归改变目录及其内容）
```



### sudo

以超级管理员（root）权限执行命令

```
sudo apt update （以 root 权限更新软件列表）
sudo su - （切换到 root 用户）
```



## 功能性命令操作

### echo

将文本"回显"到标准输出（通常是终端屏幕）

```
echo [选项] [字符串]
```



也可以使用反引号`,打印出其他命令的输出结果

``` 
例如：
echo `pwd` 输出当前的所在地址
```





### 管道符

管道符为“|”

可以将管道符左边命令的结果，作为右边命令的输入

```
命令1 | 命令2
```

利用管道符的核心思想是，类似于流水线作业：

谦虚工序的输出->传输管道->后续工序的输入



### 重定向符

重定向符分为两种，分别为">"和">>"

">"将左侧命令的结果，覆盖写入符号右侧的文件中

">>"将左侧的命令的结果，追加写入到符号右侧指定的文件中

只要是有结果输出的命令都可以结合重定向符对文件内容进行写入



## 压缩与解压

### tar

打包与解压的工具，常与gzip/bzip2结合使用

```
tar -czvf archive.tar.gz /path/to/folder （create zip verbose file：创建 .tar.gz 压缩包）
tar -xzvf archive.tar.gz （extract：解压 .tar.gz 压缩包）
tar -xjvf archive.tar.bz2 （解压 .tar.bz2 压缩包）
```



### zip/unzip

创建和解压.zip格式的压缩包

```
zip -r archive.zip folder/ （递归压缩目录）
unzip archive.zip
```



## vim编辑器

### 主要模式

| 模式         | 进入方式              | 功能                 | 退出方式                     |
| :----------- | :-------------------- | :------------------- | :--------------------------- |
| **普通模式** | 启动默认 / 按 `Esc`   | 移动光标、执行命令   | 按 `i`、`a` 等进入其他模式   |
| **插入模式** | 按 `i`、`a`、`o` 等   | 输入和编辑文本       | 按 `Esc`                     |
| **可视模式** | 按 `v`、`V`、`Ctrl+v` | 选择文本块           | 按 `Esc`                     |
| **命令模式** | 按 `:`                | 执行保存、退出等命令 | 按 `Enter` 执行或 `Esc` 取消 |



### 删除操作

| 命令        | 功能           |
| :---------- | :------------- |
| `dw`        | 删除一个单词   |
| `dd`        | 删除当前行     |
| `ndd`       | 删除 n 行      |
| `d$` 或 `D` | 删除到行尾     |
| `d^`        | 删除到行首     |
| `dG`        | 删除到文件末尾 |
| `dgg`       | 删除到文件开头 |



### 光标移动（普通模式）

| 命令            | 功能                     |
| :-------------- | :----------------------- |
| `h` `j` `k` `l` | 左、下、上、右           |
| `w`             | 移动到下一个单词开头     |
| `b`             | 移动到上一个单词开头     |
| `e`             | 移动到单词结尾           |
| `0`             | 移动到行首               |
| `^`             | 移动到行首第一个非空字符 |
| `$`             | 移动到行尾               |
| `gg`            | 移动到文件开头           |
| `G`             | 移动到文件末尾           |
| `nG` 或 `:n`    | 移动到第 n 行            |
| `Ctrl+f`        | 向下翻页                 |
| `Ctrl+b`        | 向上翻页                 |



### 查找与替换

#### 查找

| 命令       | 功能                     |
| :--------- | :----------------------- |
| `/pattern` | 向前搜索 pattern         |
| `?pattern` | 向后搜索 pattern         |
| `n`        | 下一个匹配项             |
| `N`        | 上一个匹配项             |
| `*`        | 搜索当前光标下的单词     |
| `#`        | 反向搜索当前光标下的单词 |

#### 替换（命令模式）

| 命令                | 功能                 |
| :------------------ | :------------------- |
| `:s/old/new`        | 替换当前行第一个 old |
| `:s/old/new/g`      | 替换当前行所有 old   |
| `:%s/old/new/g`     | 替换全文所有 old     |
| `:%s/old/new/gc`    | 替换全文，每次确认   |
| `:n1,n2s/old/new/g` | 在 n1 到 n2 行间替换 |



# Docker

对于Docker有四大核心概念，镜像、容器、仓库、Dockerfile

| 概念                | 说明                          |
| :------------------ | :---------------------------- |
| **镜像(Image)**     | 只读模板，用于创建容器        |
| **容器(Container)** | 镜像的运行实例                |
| **仓库(Registry)**  | 存放镜像的地方，如 Docker Hub |
| **Dockerfile**      | 构建镜像的脚本文件            |



## 镜像管理命令

### 镜像的搜索与下载

```
# 搜索镜像
docker search nginx
# 拉取镜像
docker pull nginx (直接拉取最新版本)
docker pull nginx:20.04 (拉取的是指定的版本)

# 拉取私有仓库镜像
docker pull myregistry.com/image:tag
```



### 镜像列表与检查

```
# 查看本地镜像
docker images
docker image ls

# 查看镜像详情
docker image inspect nginx:latest

# 查看镜像历史
docker history nginx
```



### 镜像清理

```docker
# 删除镜像
docker rmi nginx:latest
docker image rm nginx

# 删除悬空镜像（无标签）
docker image prune

# 强制删除所有未使用镜像
docker image prune -a
```



### 镜像导出导入

```docker
# 导出镜像
docker save -o nginx.tar nginx:latest

# 导入镜像
docker load -i nginx.tar

# 导出容器为镜像
docker export -o container.tar container_id
```



## 容器管理命令

### 容器生命周期

```docker
# 创建并启动容器
docker run -d --name mynginx -p 80:80 nginx
docker run -it ubuntu:20.04 /bin/bash

# 启动/停止/重启容器
docker start container_name
docker stop container_name
docker restart container_name

# 暂停/恢复容器
docker pause container_name
docker unpause container_name
```



### 运行选项详解

```docekr
# 常用运行选项
docker run -d \                    # 后台运行
  --name myapp \                   # 容器名称
  -p 8080:80 \                     # 端口映射
  -v /host/path:/container/path \  # 数据卷挂载
  -e ENV_VAR=value \               # 环境变量
  --network mynetwork \            # 指定网络
  --restart=always \               # 自动重启
  nginx:latest
```



### 容器的监控

```docker
# 查看运行中的容器
docker ps

# 查看所有容器（包括停止的）
docker ps -a

# 查看容器日志
docker logs container_name
docker logs -f container_name      # 实时日志
docker logs --tail 100 container_name  # 最后100行

# 进入容器
docker exec -it container_name /bin/bash
docker exec -it container_name sh

# 查看容器资源使用
docker stats container_name
# 所有容器资源使用
docker stats                       

# 查看容器进程
docker top container_name
```



### 容器清理

```docker
# 删除容器
docker rm container_name

# 强制删除运行中的容器
docker rm -f container_name       

# 清理所有停止的容器
docker container prune
```



## 数据卷管理

### 数据卷命令

```docker
q# 创建数据卷
docker volume create myvolume

# 查看数据卷
docker volume ls

# 查看数据卷详情
docker volume inspect myvolume

# 删除数据卷
docker volume rm myvolume
docker volume prune                # 清理未使用数据卷
```



### 挂载方式

```docker
# 挂载数据卷
docker run -v myvolume:/app/data nginx

# 绑定挂载（主机目录）
docker run -v /host/path:/container/path nginx

# 只读挂载
docker run -v /host/path:/container/path:ro nginx
```



## Docker网络管理

针对于Docker网络的管理，需要知道四种网络管理模式：

**bridge:**

为每一个容器分配、设置IP等，并将容器连接到一个docker0，虚拟网桥默认为该模式

使用network bridge制定，默认使用docker0

**host** :

容器将不会虚拟出自己的网卡，配置自己的ip等，而是使用宿主机的ip和端口

**none :**

容器有独立的Network,namespace,DNA并没有对其进行任何网络设置，如分配veth pair和网络桥连接，IP等

**container:**

新创建的容器不会创建自己的网卡和配置自己的IP,而是和一个指定的容器共享IP、端口范围等

使用network container:NAME或者容器ID指定



### 网络基础命令

```
# 查看网络
docker network ls

# 创建网络
docker network create mynetwork
docker network create --driver bridge mybridge

# 查看网络详情
docker network inspect mynetwork

# 连接容器到网络
docker network connect mynetwork container_name
docker network disconnect mynetwork container_name

# 删除网络
docker network rm mynetwork
```



### 网络类型

```
# 使用不同网络类型
docker run --network=bridge ...    # 默认桥接
docker run --network=host ...      # 主机网络
docker run --network=none ...      # 无网络
docker run --network=container:name ... # 共享其他容器网络
```



## Dockerfile

Dockerfile是用来构建镜像的一个文件，如何标准的去编写一个正确的Dockerfile文件时关键的，下面我将列出一个模版关于Dockerfile文件的编写

```
# 基础镜像
FROM ubuntu:20.04

# 维护者信息
LABEL maintainer="your-email@example.com"

# 设置工作目录
WORKDIR /app

# 复制文件
COPY . .
COPY requirements.txt .

# 添加文件（支持URL和解压）
ADD application.tar.gz /app/

# 运行命令
RUN apt-get update && apt-get install -y python3
RUN pip install -r requirements.txt

# 暴露端口
EXPOSE 80

# 环境变量
ENV ENVIRONMENT=production
ENV DATABASE_URL=postgresql://...

# 启动命令
CMD ["python", "app.py"]
ENTRYPOINT ["/app/start.sh"]
```



### 构建与优化

编写完Dockerfile之后，进行镜像的构建和运行优化

```
# 构建镜像
docker build -t myapp:latest .
docker build -t myapp:1.0 -f Dockerfile.prod .

# 多阶段构建示例
dockerfile
FROM node:14 as builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
```



## Docker compose容器编排

Docker容器编排是对于多个容器的一个统一管理，例如多个容器的启动和停止，启动的先后顺序

### 核心概念

什么是 Docker Compose？

- **定义**：用于定义和运行多容器 Docker 应用程序的工具
- **作用**：通过 YAML 文件配置应用服务，实现一键启动多个关联容器
- **优势**：简化复杂应用的部署和管理，解决容器间的依赖关系

### 核心组件

| 组件                   | 说明                               |
| :--------------------- | :--------------------------------- |
| **docker-compose.yml** | 定义服务、网络、数据卷的配置文件   |
| **Services**           | 容器服务定义，每个服务对应一个容器 |
| **Networks**           | 自定义网络配置                     |
| **Volumes**            | 数据卷配置                         |



### Docker Compose文件结构详解

### 基本语法结构

```
version: '3.8'          # Compose 文件版本

services:               # 服务定义部分
  web:                  # 服务名称
    image: nginx:latest # 服务配置

networks:               # 网络定义（可选）
  frontend:

volumes:                # 数据卷定义（可选）
  db_data:
```



### Compose命令



#### 基本生命周期命令

```
# 启动所有服务（构建镜像）
docker-compose up
docker-compose up -d              # 后台运行
docker-compose up --build         # 重新构建镜像后启动

# 停止服务
docker-compose down
docker-compose down -v            # 同时删除数据卷
docker-compose down --rmi all     # 同时删除镜像

# 重启服务
docker-compose restart
docker-compose restart web        # 重启指定服务
```



#### 服务管理命令

```
# 查看服务状态
docker-compose ps
docker-compose ps web             # 查看指定服务

# 查看服务日志
docker-compose logs
docker-compose logs -f            # 实时日志
docker-compose logs web db        # 查看指定服务日志
docker-compose logs --tail=100    # 查看最后100行

# 服务启停
docker-compose start
docker-compose stop
docker-compose pause
docker-compose unpause
```



#### 执行命令与调试

```
# 在服务中执行命令
docker-compose exec web /bin/bash
docker-compose exec db psql -U postgres

# 运行一次性命令（自动清理）
docker-compose run --rm web python manage.py migrate
docker-compose run --rm db /bin/bash

# 查看服务配置
docker-compose config             # 验证和查看配置
docker-compose config --services  # 列出所有服务
```



#### 镜像管理

```
# 构建镜像
docker-compose build
docker-compose build --no-cache   # 不使用缓存构建

# 拉取镜像
docker-compose pull

# 推送镜像
docker-compose push
```
