# Docker三件套

镜像（image）、容器（container）、仓库（repository）

# Docker容器

Docker 容器是一个轻量级、可移植、自给自足的软件环境，用于运行应用程序。

Docker 容器将应用程序及其所有依赖项（包括库、配置文件、系统工具等）封装在一个标准化的包中，使得应用能够在任何地方一致地运行。

## 镜像与容器的关系

- **镜像（Image）**：容器的静态模板，包含了应用程序运行所需的所有依赖和文件。镜像是不可变的。
- **容器（Container）**：镜像的一个运行实例，具有自己的文件系统、进程、网络等，且是动态的。容器从镜像启动，并在运行时保持可变。



# Docker客户端

Docker 客户端是与 Docker 守护进程（Docker Daemon）交互的命令行工具。

docker 客户端非常简单，我们可以直接输入 docker 命令来查看到 Docker 客户端的所有命令选项



## 常用的Docker客户端命令：

| **命令**              | **功能**                                         | **示例**                                   |
| :-------------------- | :----------------------------------------------- | :----------------------------------------- |
| `docker run`          | 启动一个新的容器并运行命令                       | `docker run -d ubuntu`                     |
| `docker ps`           | 列出当前正在运行的容器                           | `docker ps`                                |
| `docker ps -a`        | 列出所有容器（包括已停止的容器）                 | `docker ps -a`                             |
| `docker build`        | 使用 Dockerfile 构建镜像                         | `docker build -t my-image .`               |
| `docker images`       | 列出本地存储的所有镜像                           | `docker images`                            |
| `docker pull`         | 从 Docker 仓库拉取镜像                           | `docker pull ubuntu`                       |
| `docker push`         | 将镜像推送到 Docker 仓库                         | `docker push my-image`                     |
| `docker exec`         | 在运行的容器中执行命令                           | `docker exec -it container_name bash`      |
| `docker stop`         | 停止一个或多个容器                               | `docker stop container_name`               |
| `docker start`        | 启动已停止的容器                                 | `docker start container_name`              |
| `docker restart`      | 重启一个容器                                     | `docker restart container_name`            |
| `docker rm`           | 删除一个或多个容器                               | `docker rm container_name`                 |
| `docker rmi`          | 删除一个或多个镜像                               | `docker rmi my-image`                      |
| `docker logs`         | 查看容器的日志                                   | `docker logs container_name`               |
| `docker inspect`      | 获取容器或镜像的详细信息                         | `docker inspect container_name`            |
| `docker exec -it`     | 进入容器的交互式终端                             | `docker exec -it container_name /bin/bash` |
| `docker network ls`   | 列出所有 Docker 网络                             | `docker network ls`                        |
| `docker volume ls`    | 列出所有 Docker 卷                               | `docker volume ls`                         |
| `docker-compose up`   | 启动多容器应用（从 `docker-compose.yml` 文件）   | `docker-compose up`                        |
| `docker-compose down` | 停止并删除由 `docker-compose` 启动的容器、网络等 | `docker-compose down`                      |
| `docker info`         | 显示 Docker 系统的详细信息                       | `docker info`                              |
| `docker version`      | 显示 Docker 客户端和守护进程的版本信息           | `docker version`                           |
| `docker stats`        | 显示容器的实时资源使用情况                       | `docker stats`                             |
| `docker login`        | 登录 Docker 仓库                                 | `docker login`                             |
| `docker logout`       | 登出 Docker 仓库                                 | `docker logout`                            |

**常用选项说明:**

- **`-d`**：后台运行容器，例如 `docker run -d ubuntu`。
- **`-it`**：以交互式终端运行容器，例如 `docker exec -it container_name bash`。
- **`-t`**：为镜像指定标签，例如 `docker build -t my-image .`



### 常用的命令：

| 功能描述           | 系统命令                   | 说明                           |
| :----------------- | :------------------------- | :----------------------------- |
| **启动Docker服务** | `systemctl start docker`   | 启动Docker守护进程             |
| **停止Docker服务** | `systemctl stop docker`    | 停止Docker守护进程             |
| **重启Docker服务** | `systemctl restart docker` | 重新启动Docker守护进程         |
| **查看服务状态**   | `systemctl status docker`  | 检查Docker运行状态和日志       |
| **设置开机自启**   | `systemctl enable docker`  | 配置Docker在系统启动时自动运行 |
| **查看概要信息**   | `docker info`              | 显示Docker系统范围的详细信息   |
| **查看总体帮助**   | `docker --help`            | 列出所有顶级的Docker命令       |
| **查看命令帮助**   | `docker <命令> --help`     | 查看特定Docker命令的详细用法   |



# 容器使用

## 获取镜像

如果我们本地没有 ubuntu 镜像，我们可以使用 docker pull 命令来载入 ubuntu 镜像：

```
$ docker pull ubuntu
```



## 启动容器

以下命令使用 ubuntu 镜像启动一个容器，参数为以命令行模式进入该容器：

```
$ docker run -it ubuntu /bin/bash
```