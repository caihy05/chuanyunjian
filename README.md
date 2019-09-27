#开发环境
语言golang 12.7
系统win10
#使用说明
项目名称穿云箭
场景一：
cyj -listen 28888
cyj -reflect 10.100.0.100:28888
相当于
nc -lp 28888
nc 10.100.0.100 28888

场景二：内网穿透(代理)
cyj -tran 192.168.0.100:80,202.200.100.100:30000
cyj -monitor 30000,40000

#编译
##Windows环境：
cmd中执行build.bat
##Linux环境：
chmod +x build.sh && ./build.sh
其他环境不支持
编译二进制文件在bin目录下

#其他
限制端口为20000-60000
后续准备添加：端口扫描

#借鉴
https://github.com/cw1997/NATBypass.git

