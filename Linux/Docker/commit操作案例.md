# docker commit案例

## docker commit

提交容器副本使之成为一个新的镜像，可以像一个基本的容器中加上你需要的功能

```
docker commit -m="提交的描述信息" -a="作者" 容器ID 要创建的目标镜像名:[标签名]
```



docker容器上执行两条命令

```
apt-get update

apt-get -y install vim
```

