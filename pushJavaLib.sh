#/bin/bash
## 将java grpc api 文件传送到ftp里面
#PUTFILE=cloud-lazycat-apis-1.0.jar
#ftp -v -n 192.168.1.189<<EOF
#user anonymous lnks
#binary
#cd ./sdk
#lcd ./java_api_workspace/target/
#prompt
#del $PUTFILE
#put $PUTFILE
#bye
##here document
#EOF
#echo "push java-lib complete"

cd java_api_workspace/target
md5sum cloud-lazycat-apis-1.0.jar | cut -c-32 > cloud-lazycat-apis-1.0.jar.md5
scp -O -P 2022  cloud-lazycat-apis-1.0.jar cloud-lazycat-apis-1.0.jar.md5 root@192.168.1.189:/mnt/dl/sdk/
