###################################################################
# 使用此脚本前必须设置 Docker 用户名称和密码对应的环境变量。它们分别为:
# - DOCKER_NAME
# - DOCKER_PASSWORD
#
# 使用示例:
# chmod 755 ci.sh
# ./ci.sh
###################################################################
#!/bin/bash

# 镜像类型 test 测试类型、release 生产类型
IMAGE_TYPE=$1
# Docker 仓库地址
DOCKER_REGISTRY_PROXY="registry.com"
# 组
DOCKER_GROUP="group"
# 镜像名
DOCKER_IMAGE_NAME="test"
# 版本
DOCKER_TAG=$(cat version)"-$IMAGE_TYPE"
# 镜像名称
DOCKER_IMAGE_NAME_TAG=$DOCKER_REGISTRY_PROXY"/"$DOCKER_GROUP"/"$DOCKER_IMAGE_NAME":"$DOCKER_TAG

echo "开始构建 Docker 镜像: "$DOCKER_IMAGE_NAME_TAG
# 登录远程仓库
docker login $DOCKER_REGISTRY_PROXY -u $DOCKER_NAME -p $DOCKER_PASSWORD
# 执行构建
docker build --no-cache=false -t $DOCKER_IMAGE_NAME_TAG -f Dockerfile .
# 推送镜像到仓库
docker push $DOCKER_IMAGE_NAME_TAG
# 执行完成
echo "任务执行完成"
