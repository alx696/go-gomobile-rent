#!/bin/sh -e

# 设置目标文件夹
TARGET_DIR="/home/m/dev/android-rent/gomobile/"

echo "打包aar: ${PACKAGE}"

# 复制依赖
go mod vendor; cp -r vendor/* $GOPATH/src/; rm -rf $GOPATH/src/pkg $GOPATH/src/modules.txt ; rm -rf vendor

# 复制源码
TEMPPATH="$GOPATH/src/lilu.red/temp"
mkdir -p $TEMPPATH
# dns包
cp -r dns $TEMPPATH
# safe包
cp -r safe $TEMPPATH

# 打包
GO111MODULE="off"
gomobile bind -v -o "${TARGET_DIR}gomobile.aar" -target=android "${TEMPPATH}/dns" "${TEMPPATH}/safe"
rm -rf TEMPPATH

echo "打包完成"