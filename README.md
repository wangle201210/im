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
需要参数:{
	name		字符串		姓名,
	password	字符串		密码,
	room		数字			房间号,
}
返回结果:{
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
需要参数:{
	name		字符串		姓名,
	password	字符串		密码,
	room		数字			房间号,
}
返回结果:{
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
接收到的content:{"Type":2,"User":"用户名","room":1,"Timestamp":1548658898216,"content":"这里是内容"} 

注解:
    token       登陆后从c端获取到的用于验证身份
    room        房间号
    content     发送(接受)的内容

    content.Type        c端不用
    content.User        用户名
    content.room        房间号
    content.Timestamp   时间戳
    content.content     消息内容
```

### 登陆
```
功能:超级管理员登陆(用于上传和修改资料,主要是获取token)
方法:post
接口:/api/login
需要参数:{
    "name":"admin",
    "password":"password"
}
返回结果:{
    "code": 200,
    "msg": "登陆成功！",
    "data": {
        "user": {
            "id": 1,
            "name": "admin",
            "password": "password",
            "role": "admin",
            "token": "",
            "chat": null,
            "created_at": "2019-01-27T22:38:23+08:00",
            "updated_at": "2019-01-27T22:38:29+08:00"
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDk5OTI4MDEsImlhdCI6MTU0OTk4OTIwMSwiaW5mbyI6eyJVaWQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJQYXNzd29yZCI6InBhc3N3b3JkIiwiUm9sZSI6ImFkbWluIn0sIm5iZiI6MTU0OTk4OTIwMX0.cw5JAKnQddcNxZ-J16DTwFs0eyKmAL4bh3aqo0wUytY"
    }
}

普通登陆在界面完成即可
```

### 视频
```
功能:c端新增视频
方法:post
接口:/api/admin/videos
需要参数:{
        "url": "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
        "content": "描述(备注)",
        "room": 3,
        "c_id": 1
}
返回结果:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 4,
        "c_id": 0,
        "url": "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
        "content": "描述(备注)",
        "room": 3,
        "created_at": "2019-02-11T23:44:55.592745+08:00",
        "updated_at": "2019-02-11T23:44:55.592747+08:00"
    }
}

功能:c端删除视频,c_id为c端id
方法:delete
接口:/api/admin/video/{c_id}
需要参数:{
    
}
返回结果:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改视频,c_id为c端id
方法:put
接口:/api/admin/video/{c_id}
需要参数:{
        "url": "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
        "content": "描述(备注)",
        "room": 1,
        "c_id": 1
}
返回结果:{
    "code": 205,
    "msg": "更新成功",
    "data": {
        "id": 1,
        "c_id": 1,
        "url": "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
        "content": "描述(备注)",
        "room": 1,
        "created_at": "2019-01-26T14:06:41+08:00",
        "updated_at": "2019-02-11T23:48:21.115248+08:00"
    }
}

功能:查询所有视频
方法:get
接口:/api/videos
需要参数:{

}
返回结果:{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "id": 14,
            "c_id": 1,
            "title": "这里是标题",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-12T15:33:45+08:00",
            "updated_at": "2019-02-12T15:33:45+08:00"
        },
        {
            "id": 13,
            "c_id": 1,
            "title": "这里是标题2",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-12T15:33:32+08:00",
            "updated_at": "2019-02-12T15:33:32+08:00"
        }
    ]
}
```

### 菜单
```
功能:c端新增图片
方法:post
接口:/api/admin/pics
需要参数:{
        "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
        "order": 3,
        "room": 1,
        "c_id": 1
}
返回结果:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 16,
        "c_id": 1,
        "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
        "order": 3,
        "room": 1,
        "created_at": "2019-02-11T23:50:52.505634+08:00",
        "updated_at": "2019-02-11T23:50:52.505635+08:00"
    }
}

功能:c端删除图片同步更新用户界面,c_id为c端id
方法:delete
接口:/api/admin/pic/{c_id}
需要参数:{
    
}
返回结果:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改图片,c_id为c端id
方法:put
接口:/api/admin/pic/{c_id}
需要参数:{
    "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
    "order": 999,
    "room": 1
}
返回结果:{
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

功能:查询所有菜单
方法:get
接口:/api/pics
需要参数:{

}
返回结果:{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "id": 4,
            "c_id": 4,
            "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
            "order": 999,
            "room": 1,
            "created_at": "2019-01-27T15:45:04+08:00",
            "updated_at": "2019-01-28T05:33:35+08:00"
        },
        {
            "id": 5,
            "c_id": 5,
            "url": "http://bpic.588ku.com//back_water_img/18/06/13/5a31c43a3c4df7a20f3cb7cdc873bd47.jpg",
            "order": 999,
            "room": 1,
            "created_at": "2019-01-27T15:45:04+08:00",
            "updated_at": "2019-01-28T05:36:31+08:00"
        }
    ]
}

```


### 教学
```
功能:c端新增教学
方法:post
接口:/api/admin/courses
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "c_id": 1
}
返回结果:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 3,
        "c_id": 1,
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-12T00:05:52.781669+08:00",
        "updated_at": "2019-02-12T00:05:52.781672+08:00"
    }
}

功能:c端删除教学,c_id为c端id
方法:delete
接口:/api/admin/course/{c_id}
需要参数:{
    
}
返回结果:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改教学,c_id为c端id
方法:put
接口:/api/admin/course/{c_id}
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "c_id": 1
}
返回结果:{
    "code": 205,
    "msg": "更新成功",
    "data": {
        "id": 2,
        "c_id": 11,
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-11T16:04:59+08:00",
        "updated_at": "2019-02-12T00:13:41.509851+08:00"
    }
}

功能:查询所有教学
方法:get
接口:/api/courses
需要参数:{

}
返回结果:{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "id": 4,
            "c_id": 1,
            "content": "2这里是内容,可以是富文本",
            "room": 2,
            "created_at": "2019-02-11T16:12:29+08:00",
            "updated_at": "2019-02-11T16:12:29+08:00"
        },
        {
            "id": 2,
            "c_id": 11,
            "content": "1这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-11T08:04:59+08:00",
            "updated_at": "2019-02-11T16:13:41+08:00"
        }
    ]
}
```


### 历史
```
功能:c端新增历史
方法:post
接口:/api/admin/histories
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "title":"这里是标题",
        "c_id": 1
}
返回结果:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 8,
        "c_id": 1,
        "title": "这里是标题",
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-12T00:19:32.748528+08:00",
        "updated_at": "2019-02-12T00:19:32.74853+08:00"
    }
}

功能:c端删除历史,c_id为c端id
方法:delete
接口:/api/admin/history/{c_id}
需要参数:{
    
}
返回结果:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改历史,c_id为c端id
方法:put
接口:/api/admin/history/{c_id}
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "title":"这里是标题2",
        "c_id": 1
}
返回结果:{
    "code": 205,
    "msg": "更新成功",
    "data": {
        "id": 8,
        "c_id": 1,
        "title": "这里是标题2",
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-11T16:19:32+08:00",
        "updated_at": "2019-02-12T00:22:04.562769+08:00"
    }
}

功能:查询所有历史
方法:get
接口:/api/histories
需要参数:{

}
返回结果:{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "id": 8,
            "c_id": 1,
            "title": "这里是标题2",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-11T08:19:32+08:00",
            "updated_at": "2019-02-11T16:22:04+08:00"
        },
        {
            "id": 7,
            "c_id": 111,
            "title": "这里是标题2",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-01-30T02:44:41+08:00",
            "updated_at": "2019-02-12T15:33:14+08:00"
        }
    ]
}
```


### 记录
```
功能:c端新增记录
方法:post
接口:/api/admin/records
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "title":"这里是标题",
        "c_id": 1
}
返回结果:{
    "code": 201,
    "msg": "添加成功",
    "data": {
        "id": 14,
        "c_id": 1,
        "title": "这里是标题",
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-12T23:33:45.356802+08:00",
        "updated_at": "2019-02-12T23:33:45.356803+08:00"
    }
}

功能:c端删除记录,c_id为c端id
方法:delete
接口:/api/admin/record/{c_id}
需要参数:{
    
}
返回结果:{
    "code": 200,
    "msg": "删除成功",
    "data": ""
}

功能:c端修改记录,c_id为c端id
方法:put
接口:/api/admin/record/{c_id}
需要参数:{
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "title":"这里是标题2",
        "c_id": 1
}
返回结果:{
    "code": 205,
    "msg": "更新成功",
    "data": {
        "id": 8,
        "c_id": 1,
        "title": "这里是标题2",
        "content": "这里是内容,可以是富文本",
        "room": 1,
        "created_at": "2019-02-11T16:19:32+08:00",
        "updated_at": "2019-02-12T00:22:04.562769+08:00"
    }
}

功能:查询所有记录
方法:get
接口:/api/records
需要参数:{

}
返回结果:{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "id": 14,
            "c_id": 1,
            "title": "这里是标题",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-12T15:33:45+08:00",
            "updated_at": "2019-02-12T15:33:45+08:00"
        },
        {
            "id": 13,
            "c_id": 1,
            "title": "这里是标题2",
            "content": "这里是内容,可以是富文本",
            "room": 1,
            "created_at": "2019-02-12T15:33:32+08:00",
            "updated_at": "2019-02-12T15:33:32+08:00"
        }
    ]
}
```
{
    Content: "大家好！",
    Room: 1,
    Timestamp: 1552566903938,
    Type: 0,
    User: "u",  
}


## 备注
以上所有接口都需要在请求头里加入token值,token来源于登陆结果
如下:
Authorization : Bearer token


