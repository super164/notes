# Dockerfile案例

## Centos7镜像具备vim+ifconfig+jdk8

1. 准备编写Dockerfile文件
2. 构建新的镜像

```
docker build -t 新镜像名字:TAG.
```

Dockerfile文件内容配置:

```
FROM centos:7
MAINTAINER zzyy<zzyy@126.com>
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo && \
    curl -o /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo && \
    sed -i 's/$releasever/7/g' /etc/yum.repos.d/CentOS-Base.repo && \
    yum makecache
ENV MYPATH /usr/local
WORKDIR $MYPATH

#安装vim编辑器
RUN yum -y install vim
#安装ifconfig命令查看网络IP
RUN yum -y install net-tools
#安装java8及lib库
RUN yum -y install glibc.i686
RUN mkdir /usr/local/java
#ADD 是相对路径jar,把jdk 8u171-linux-x64.tar.gz添加到容器中，安装包必须要和Dockerfile文件在同一位置
ADD jdk-8u461-linux-i586.tar.gz /usr/local/java/
#配置java环境变量
ENV JAVA_HOME /usr/local/java/jdk1.8.0_461
ENV JRE_HOME $JAVA_HOME/jre
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib:$CLASSPATH
ENV PATH $JAVA_HOME/bin:$PATH

EXPOSE 80

CMD echo $MYPATH
CMD echo "success-...... ok"
CMD /bin/bash

```



## Dockerfile配置虚悬镜像

配置一个Dockerfile文件

```
FROM ubuntu
CMD -echo 'sahdfjkashfk'
```

使用下面命令可以进行查询出所有虚悬镜像

```
docker image ls -f dangling=true
```



删除虚悬镜像命令:

```
docker image prune
```

