#!/bin/bash
#cnl = cloudNativeLearn
sudo docker build  -t my/cnl-httpserver:v1.0 -f Dockerfile-httpserver-multiple .

sudo docker run --name httpserver -i -p 7000:7000 -d my/cnl-httpserver:v1.0

#推送镜像
#sudo docker push my/cnl-httpserver:v1.0

#健康检查
curl -s localhost:7000/healthz -w 'http_status:%{http_code}'

#查找容器网络ns
#sudo lsns -t net
#sudo lsns -t net -r | grep httpserver | awk '{print $4}'
pid=$(lsns -t net -r | grep httpserver | cut -d ' ' -f4)

#输出容器进程ns信息
cd /proc/"$pid"/ns && ls -al

#输出容器ip配置
sudo nsenter -t "$pid" -n ip addr