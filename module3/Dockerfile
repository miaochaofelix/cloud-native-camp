# Step 1: 我的虚拟机是64位ubuntu系统且运行在M1 PRO机器上，所以先用以下命令编译程序
# GOOS=linux GOARCH=arm64 go build

# Step 2: 测试将编译好的程序在虚拟机上运行，可以正常访问 http://10.211.55.3:8080/

# Step 3: 构建docker镜像
# sudo docker build -t httpserver:v1 .

# Step 4: 启动容器
# sudo docker run -p 80:8080 -it httpserver:v1

# Step 5: 测试可以正常访问 http://10.211.55.3

# Step 6: 将镜像推送到官方镜像仓库 https://hub.docker.com/repository/docker/miaochaofelix/httpserver
# sudo docker tag httpserver:v1 miaochaofelix/httpserver:v1
# sudo docker push miaochaofelix/httpserver:v1

# Step 7: 进入容器查看IP配置。通过以下命令可以查看到容器的IP是172.17.0.2
# sudo docker inspect c0455ad2d06b | grep -i pid
# sudo nsenter -t 5719 -n ip a

FROM ubuntu

ADD ./httpserver httpserver
ENTRYPOINT ["/httpserver"]

EXPOSE 8080