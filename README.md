# im
## 基于beego WebIM开发的在线聊天工具
 ```
 1.clone 项目 --git clone https://github.com/wangle201210/im.git
 2.本地运行项目(默认位8081端口) --bee run 
 3.访问localhost:8081
 ```

## 线上运行
```
1.打包linux版本项目 --CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
2.给可执行文件权限  --chmod +x [可执行文件名字]
3.运行 --nohup ./[可执行文件名字]
4.退出 
  先使用--ps -aux | grep [可执行文件名字] 获取PID
  然后 kill [PID]
  ```
## 功能
```
基础聊天功能
多频道
管理员未登陆或下线后不允许用户登陆
手机端兼容
视屏播放
同账号可多处登陆
```