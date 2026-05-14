# Docker Hello World

在Docker中运行“Hello World”,使用docker run命令来实现，但需要在容器内进行运行

示例代码如下：

```
docker run ubuntu:15.10 /bin/echo "Hello world"
```

- **ubuntu:**15.10:指定要运行的镜像，Docker 首先从本地主机上查找镜像是否存在，如果不存在，Docker 就会从镜像仓库 Docker Hub 下载公共镜像。
- **/bin/echo "Hello world":** 在启动的容器里执行的命令





# 运行交互式的容器

通过 docker 的两个参数 -i -t，让 docker 运行的容器实现**"对话"**的能力：

```
docker run -i -t ubuntu:15.10 /bin/bash
```

- **-t:** 在新容器内指定一个伪终端或终端。
- **-i:** 允许你对容器内的标准输入 (STDIN) 进行交互。



出现“:/#"时候说明已经进入了一个ubuntu15.10系统的容器

可以通过exit命令或者CTRL+D来退出容器



# 启动容器（后台模式）

使用一下命令创建一个一进程方式运行的容器：

```
docker run -d ubuntu:15.10 /bin/sh -c
"while true; do echo hello world; sleep 1; done"
```

此时会输出一段长字符：“90d5c2de0a68d5aab0d634c62dbf2511cc23dd6bff43f6edeafaad4c5e2c6f28”

这段字符为容器ID,可以通过容器ID来查看对应的容器发生了什么



## docker ps

该命令可以查看容器运行状态

```
[root@localhost yum.repos.d]# docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS         PORTS     NAMES
035f341c3fb6   ubuntu:15.10   "/bin/sh -c 'while t…"   3 minutes ago   Up 3 minutes           serene_germain

```

**CONTAINER ID:** 容器 ID。

**IMAGE:** 使用的镜像。

**COMMAND:** 启动容器时运行的命令。

**CREATED:** 容器的创建时间。

**STATUS:** 容器状态

created（已创建）	restarting（重启中）	running 或 Up（运行中）	removing（迁移中）	paused（暂停）	exited（停止）	dead（死亡）

**PORTS:** 容器的端口信息和使用的连接类型（tcp\udp）。

**NAMES:** 自动分配的容器名称。



## docker logs

```
docker logs+CONTAINER ID 或者docker logs +NAMES
```

这两种命令都可以，对容器中的标准进行输出



# 停止容器

## docker stop

通过该命令docker stop命令来停止容器：

```
docker stop CONTAINER ID或docker logs NAMES
```



通过docker ps查看容器工作状态

输出没有内容的时候就说明容器已经停止工作了



