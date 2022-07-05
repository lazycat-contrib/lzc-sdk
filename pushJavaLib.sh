#/bin/bash
## 将java grpc api 文件传送到ftp里面
PUTFILE=cloud-lazycat-apis-1.0.jar
ftp -v -n 192.168.1.193<<EOF
user ftp账号 ftp密码
binary
cd ./sdk
lcd ./java_api_workspace/target/
prompt
del $PUTFILE
put $PUTFILE
bye
#here document
EOF
echo "push java-lib complete"
