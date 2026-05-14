# Dockerfile

## 什么是Dockerfile

Dockerfile 是一个文本文件，包含了构建 Docker 镜像的所有指令。

Dockerfile 是一个用来构建镜像的文本文件，文本内容包含了一条条构建镜像所需的指令和说明。

通过定义一系列命令和参数，Dockerfile 指导 Docker 构建一个自定义的镜像。



## 构建Dockerfile步骤

三步走：

- 编写Dockerfile文件
- docker build命令构建镜像
- docker run 依照新编镜像运行容器实例



## 基础知识

1. 每条指令都必须为大写字母且后面至少跟随一个参数
2. 指令从上到下，顺序执行
3. #表示注释
4. 每条指令都会创建一个新的镜像层并对镜像进行提交

Dockerfile进行编写，创建Docker镜像，运行Docker容器实例



## 使用Dockerfile定制镜像

**1、下面以定制一个 nginx 镜像（构建好的镜像内会有一个 /usr/share/nginx/html/index.html 文件）**

在一个空目录下，新建一个名为 Dockerfile 文件，并在文件内添加以下内容：

```
FROM nginx
RUN echo '这是一个本地构建的nginx镜像' > /usr/share/nginx/html/index.html
```





## 构建镜像

在 Dockerfile 文件的存放目录下，执行构建动作。

以下示例，通过目录下的 Dockerfile 构建一个 nginx:v3（镜像名称:镜像标签）

```
docker build -t nginx:v3 .
```



### 上下文路径

上一节中，有提到指令最后一个 **.** 是上下文路径，那么什么是上下文路径呢？

$ docker build -t nginx:v3 .

上下文路径，是指 docker 在构建镜像，有时候想要使用到本机的文件（比如复制），docker build 命令得知这个路径后，会将路径下的所有内容打包。

**解析**：由于 docker 的运行模式是 C/S。我们本机是 C，docker 引擎是 S。实际的构建过程是在 docker 引擎下完成的，所以这个时候无法用到我们本机的文件。这就需要把我们本机的指定目录下的文件一起打包提供给 docker 引擎使用。

如果未说明最后一个参数，那么默认上下文路径就是 Dockerfile 所在的位置。

**注意**：上下文路径下不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。



## Dockerfile常用的保留字指令

### FROM

基础镜像，当前新镜像是基于那个镜像的，指定一个已经存在的镜像作为模版，第一条一般为FROM



### MAINTAINER

**镜像维护者的姓名和邮箱地址**



### RUN

**容器构建的时候需要执行的命令**

一般为两种格式：shell格式，exec格式

RUN这个命令是在docker build是运行

```
RUN <命令行命令>
#<命令行命令>等同于，在终端操作的shell命令

RUN ["可执行文件","参数1"，"参数2"]
#例如
#RUN ["./test.php","dev","offline"]等价于 RUN ./test.php dev offline
```



### EXPOSE

当前容器对外暴露出的端口



### WORKDIR

指定在创建容器后，终端默认登录的进来工作目录，一个落脚点



### USER

指定该镜像以什么样的用户去执行，如果不指定，默认是root



### ENV

用来在构建镜像过程中设置环境变量

```
ENV MY_PATH /usr/mytest
该环境变量可以在后续的任何RUN指令中使用，这就如同在命令前面制定了环境变量前缀一样，也可以 在其他指令中直接使用这些环境变量
比如：WORKDIR $MY_PATH
```



### ADD

将宿主机目录下的文件拷贝进镜像且会自动处理URL和解压tar压缩包



### COPY

类似于ADD，拷贝文件和目录到镜像中。将从构建上下文目录中<源路径>的文件/目录复制道心的一层的镜像内的<目的路径>位置



### VOLUME

容器数据卷，用于数据保存和持久化工作



### CMD

制定容器启动后要干的事情



**CMD容器启动命令**

CMD指令的格式和RUN相似，也是两种格式：

```
shell格式:CMD<命令>
exec格式:CMD ["可执行文件","参数1","参数2"...]
参数列表格式:CMD ["参数1","参数2"...],在指定ENTRYPOINT指令后，用CMD指定具体的参数
```

Dockerfile中可以有多个CMD指令，但只有最后一个生效，CMD会被docker run之后的参数替换

！！！

CMD和RUN命令的区别：

CMD是在docker run时运行

RUN实在docker build时运行



### ENTRYPOINT

也是用来指定一个容器启动时要运行的命令

类似于CMD指令，但是ENTRYPOINT不会被docker run后面的命令覆盖，而且这些命令行参数会被当做参数送给ENTRYPOINT指令指定的程序

当二者一起使用的时候，CMD不在直接运行而是将内容作为参数传递给ENTRYPOINT指令，如下组合：

```
<ENTRYPOINT> "<CMD>"
```





## 指令详解

| Dockerfile 指令 | 说明                                                         |
| :-------------- | :----------------------------------------------------------- |
| FROM            | 指定基础镜像，用于后续的指令构建。                           |
| MAINTAINER      | 指定Dockerfile的作者/维护者。（已弃用，推荐使用LABEL指令）   |
| LABEL           | 添加镜像的元数据，使用键值对的形式。                         |
| RUN             | 在构建过程中在镜像中执行命令。                               |
| CMD             | 指定容器创建时的默认命令。（可以被覆盖）                     |
| ENTRYPOINT      | 设置容器创建时的主要命令。（不可被覆盖）                     |
| EXPOSE          | 声明容器运行时监听的特定网络端口。                           |
| ENV             | 在容器内部设置环境变量。                                     |
| ADD             | 将文件、目录或远程URL复制到镜像中。                          |
| COPY            | 将文件或目录复制到镜像中。                                   |
| VOLUME          | 为容器创建挂载点或声明卷。                                   |
| WORKDIR         | 设置后续指令的工作目录。                                     |
| USER            | 指定后续指令的用户上下文。                                   |
| ONBUILD         | 当该镜像被用作另一个构建过程的基础时，添加触发器。           |
| SHELL           | 覆盖Docker中默认的shell，用于RUN、CMD和ENTRYPOINT指令。      |
