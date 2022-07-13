#!/bin/bash
cd daliuren
echo "测试天干寄宫"
go test jigong.go jigong_test.go
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi

echo "测试天盘"
go test tianpan.go tianpan_test.go
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi

echo "测试天将"
go test tianpan.go tianjiang.go tianjiang_test.go
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi

echo "测试四课"
go test tianpan.go sike.go jigong.go sike_test.go
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi

echo "测试三传"
go test tianpan.go sike.go jigong.go sanchuan.go sanchuan_test.go 
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi


echo "测试函数:NewDaliurenShiPan"
# EPHE_PATH=$1  CGO_LDFLAGS="-L$2 -lswe -lm -ldl -static" go test tianpan.go sike.go jigong.go sanchuan.go tianjiang.go daliuren.go daliuren_test.go
LD_LIBRARY_PATH=$2 EPHE_PATH=$1 CGO_LDFLAGS="-L$2 -lswe -lm -ldl" go test tianpan.go sike.go jigong.go sanchuan.go tianjiang.go daliuren.go daliuren_test.go
if [ $? -ne 0 ]
then
    echo "测试失败"
    exit 1
fi
