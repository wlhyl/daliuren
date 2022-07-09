# golang 语言编写的六壬排盘server

## 单元测试

* 下载瑞士星历表，并编译
```bash
mkdir /tmp/swe
cd /tmp/swe
wget https://www.astro.com/ftp/swisseph/swe_unix_src_2.10.02.tar.gz
tar xvzf swe_unix_src_2.10.02.tar.gz 
cd swe
make libswe.so
```
* 下载瑞士星历表文件
```bash
wget https://www.astro.com/ftp/swisseph/ephe/semo_18.se1
wget https://www.astro.com/ftp/swisseph/ephe/semom48.se1
wget https://www.astro.com/ftp/swisseph/ephe/sepl_18.se1
wget https://www.astro.com/ftp/swisseph/ephe/seplm48.se1
```

* 单元测试
```bash
./test.sh `pwd`  /tmp/swe/src
```

## 接口文档
[/swagger/index.html#/daliuren/post_api_daliuren](/swagger/index.html)


## 请求api
```bash
data='{"year":2022, "month":7,"day":8,"yue_jiang":"", "divination_time":"寅","diurnal":true, "year_of_birth":2000, "masculine":true}'

curl -X POST  -D  - --data $data  -H "Content-Type: application/json" 127.0.0.1/api/daliuren
```

## 构建镜像
```bash
VUE_APP_API_URI=\'http://your_host/api\' > ui/.env.production.local
make images
```