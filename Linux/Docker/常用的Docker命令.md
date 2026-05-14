# docker的常用命令

此篇中，只统计了一docker一部分的知识点，主要是关于docker的镜像以及容器的基本操作命令。下面表格中就是相关命令：

| **命令**          | **功能**                         | **示例**                                   |
| :---------------- | :------------------------------- | :----------------------------------------- |
| `docker run`      | 启动一个新的容器并运行命令       | `docker run -d ubuntu`                     |
| `docker ps`       | 列出当前正在运行的容器           | `docker ps`                                |
| `docker ps -a`    | 列出所有容器（包括已停止的容器） | `docker ps -a`                             |
| `docker build`    | 使用 Dockerfile 构建镜像         | `docker build -t my-image .`               |
| `docker images`   | 列出本地存储的所有镜像           | `docker images`                            |
| `docker pull`     | 从 Docker 仓库拉取镜像           | `docker pull ubuntu`                       |
| `docker push`     | 将镜像推送到 Docker 仓库         | `docker push my-image`                     |
| `docker exec`     | 在运行的容器中执行命令           | `docker exec -it container_name bash`      |
| `docker stop`     | 停止一个或多个容器               | `docker stop container_name`               |
| `docker start`    | 启动已停止的容器                 | `docker start container_name`              |
| `docker restart`  | 重启一个容器                     | `docker restart container_name`            |
| `docker rm`       | 删除一个或多个容器               | `docker rm container_name`                 |
| `docker rmi`      | 删除一个或多个镜像               | `docker rmi my-image`                      |
| `docker logs`     | 查看容器的日志                   | `docker logs container_name`               |
| `docker inspect`  | 获取容器或镜像的详细信息         | `docker inspect container_name`            |
| `docker exec -it` | 进入容器的交互式终端             | `docker exec -it container_name /bin/bash` |



部分命令说明以及示例如下：

## 新建并启动一个容器

```
docker run [OPTIONS] IMAGE[COMMAND][ARG...]
例如新建一个Ubuntu容器：
docker run -it --name=qwe1 ubuntu /bin/bash
```

OPTIONS(说明):有些是一个减号，有些是两个减号



--name="容器的新名字"	为容器制定一个名称

-d:后台运行容器并返回容器ID,也即启动守护式容器（后台运行）

-i:已交互模式运行同期，通常于-t同时使用；

也即启动交互式容器（前台有伪终端，等待交互）

-P:随机端口映射，大写P

-p:制定端口映射，小写p



退出终端输入:exit

## 列出当前所有正在运行的容器

```
docker ps [OPTIONS]
```

OPTIONS：

-a；列出当前所有正在运行的容器+历史上运行过的

-l:显示最近创建的容器

-n:显示最近n个创建的容器

-q:静默模式，只显示容器编号



## 退出容器

```
run 进去的容器，用exit退出，容器停止
用ctrl+p+q退出，容器不停止
```



## 启动已停止的容器

```
docker start 容器ID或者容器名
```



## 重启容器

```
docker restart 容器ID或者容器名
```



## 停止容器

```
docker stop 容器ID或者容器名
```



## 强制停止容器

```
docker kill 容器ID或者容器名
```



## 删除已停止的容器

```
docker rm 容器ID
!!!!! 一次性删除多个容器
docker rm -f $(docker ps -a -q)
docker ps -a -q | xargs docker rm
```



## 启动守护式容器（后台运行）

```
docker run -d 容器名
```

Docker容器后台运行的时候，必须要有一个前台的进程，容器运行的命令如果不是那些一起挂起的命令（如top,tail），就会自动退出



```
docker logs 容器ID 查看容器日志
```



## 查看容器内部细节

```
docker inspect 容器ID 
```





## 进入正在运行的容器并以命令行交互

 ```
1. docker exec -it 容器ID bashShell
2. docker attach 容器ID
 ```

- attach直接进入容器启动命令的终端，不会启动新的进程，用exit退出，会导致容器的停止
- exec是在容器中打开新的终端，并且可以启动新的进程，用exit退出，不会导致容器的停止

一般推荐使用exec命令



## 从容器上拷贝文件到主机上

```
docker cp 容器ID:容器内路径 目的的路径
```





## docker images

罗列出本地的镜像：

REPOSITORY    TAG       IMAGE ID       CREATED       SIZE
hello-world   latest    1b44b5a3e06a   5 weeks ago   10.1kB

- REPOSITORY：表示镜像的仓库源
- TAG:镜像的标签版本号
- IMAGE ID:镜像的ID
- CREATED:镜像的创建时间
- SIZE:镜像大小

同一个仓库源可以有多个TAG版本，代表这个仓库源的不同个版本，我们可以使用REPOSITORY :TAG来定义不同的镜像。如果在使用的时候不去指定版本标签，例如你只使用ubuntu,docker将默认使用:ubuntu:latest镜像。

可以在后面加上选项：加-a,是可以列出本地所有的镜像	加-q,是只显示镜像


## docker search

查询某个镜像

```
在使用的时候要加上docker search docker.1ms.run/要查询的东西
内容如下:
[root@localhost ~]# docker search docker.1ms.run/hello-world
NAME                                DESCRIPTION                                     STARS     OFFICIAL
hello-world                         Hello World! (an example of minimal Dockeriz…   2489      [OK]
rancher/hello-world                 This container image is no longer maintained…   6         
okteto/hello-world                                                                  0         
atlassian/hello-world                                                               1         
goharbor/hello-world                                                                0         
tutum/hello-world                   Image to test docker deployments. Has Apache…   91        
dockercloud/hello-world             Hello World!                                
```

-Name镜像名称	-DESCRIPTION镜像说明	STARS点赞数	OFFICIAL是否为官方的	AUTOMATED是否是自动构建的

- 可以限制查询出来的镜像数：docker search --limit 5（限制输出五个） docker.1ms.run/要查询的东西





## 拉取镜像（docker pull）

下载镜像

```
docker pull 镜像名字 [TAG]可以选择下载不同版本
```



## docker system df

```
docker system df 查看镜像/容器/数据卷所占的空间
```

查看docker容器中下载的镜像内存等





## docker rmi 删除镜像

1. 删除单个

```
docker rmi 某个镜像的名字ID
```

加上-f强制删除

2. 删除多个

```
docker rmi -f 镜像名1:TAG 镜像名2:TAG
```

3. 删除全部：

```
docker rmi -f $(docker images -qa)
```

