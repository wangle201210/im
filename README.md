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
后台图片更新订阅者界面自动更新
```
## 接口
```
c端的接口名字可以有c端修改,但是我们传入的数据格式和我们所需要的数据必须保持一致
时间均取当前服务器时间,不单独发送
带***的功能为需要c端实现的
c端都通过http接口向b端发送信息,b端接收到后通过ws向订阅者界面推送
```
### 登陆
```
功能:c端管理员登陆系统登陆时会向c端发送验证请求
方法:post
接口:/login
req:{
	name		字符串		姓名,
	password	字符串		密码,
	room		数字			房间号,
}
resp:{
    "code": 200,
    "msg": "登陆成功！",
    "data": {
        "user": {
            c端返回的用户信息
        },
        "token": b端颁发用户的token	用于验证用户	c端必须保留token用于之后c端发送信息到b端
    }
}

功能:发送登陆信息到c端并验证     ***
方法:post
接口:/c/user/login
req:{
	name		字符串		姓名,
	password	字符串		密码,
	room		数字			房间号,
}
resp:{
	name		字符串		姓名,
	room		数字			房间号,
	role		字符串		权限(admin|user),
	info		json		用户其他信息,
}
```
### 聊天记录
```
功能:c端发送聊天到b端某房间及c端接收b端信息    ***
方法:ws
接口:ws://域名/ws/join?token=xxxxx&room=xxx
发送消息调用:websock.send(content)
接受接受调用:websock.onmessage(content)
接收到的content:{"Type":2,"User":"用户名","Room":1,"Timestamp":1548658898216,"Content":"这里是内容"} 

注解:
    token       登陆后从c端获取到的用于验证身份
    room        房间号
    content     发送(接受)的内容

    content.Type        c端不用
    content.User        用户名
    content.Room        房间号
    content.Timestamp   时间戳
    content.Content     消息内容
```
### 图片
```
功能:c端新增图片同步推到用户界面
方法:post
接口:/pics
req:{
    c_id		数字		c端id
    url			字符串	图片网址
    order		数字		排序(越大显示越靠前)
    room		数字		房间号
}
resp:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 14,
        "c_id": 0,
        "url": "http://pic17.nipic.com/20111021/8633866_210108284151_2.jpg",
        "order": 2,
        "room": 1,
        "created_at": "2019-01-28T11:16:06.589705+08:00",
        "updated_at": "2019-01-28T11:16:06.589708+08:00"
    }
}

功能:c端删除图片同步更新用户界面,c_id为c端id
方法:delete
接口:/pic/c_id
req:{
    
}
resp:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改图片,c_id为c端id
方法:put
接口:/pic/c_id
req:{
    "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
    "order": 999,
    "room": 1
}
resp:{
    "code": 205,
    "msg": "更新成功",
    "data": {
        "id": 5,
        "c_id": 5,
        "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
        "order": 999,
        "room": 1,
        "created_at": "2019-01-27T23:45:04+08:00",
        "updated_at": "2019-01-28T13:36:31.667765+08:00"
    }
}
```











