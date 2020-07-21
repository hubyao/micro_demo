# demo
## 结构

```
.
├── basic   // 设置与中间件初始化解析的
│   ├── basic.go
│   ├── config
│   │   ├── config.go
│   │   ├── etcd.go
│   │   ├── jwt.go
│   │   ├── mysql.go
│   │   ├── profiles.go
│   │   └── redis.go
│   ├── db
│   │   ├── db.go
│   │   ├── mysql.go
│   │   └── xgorm
│   │       └── db.go
│   ├── go.mod
│   └── redis
│       └── redis.go
├── comm  // 公共目录
│   ├── go.mod
│   └── xhttp
│       ├── errno
│       │   ├── code.go
│       │   └── errno.go
│       ├── req.go
│       └── rsp.go
├── proto // pb文件目录
│   ├── go.mod
│   └── user
│       ├── user.pb.go
│       ├── user.pb.micro.go
│       ├── user.proto
├── README.md
├── user-srv   // 用户srv服务
│   ├── conf
│   │   ├── application-db.yml
│   │   ├── application-etcd.yml
│   │   └── application.yml
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   │   └── user.go  
│   ├── main.go
│   ├── model
│   │   ├── init.go
│   │   └── user  // 操作数据c层
│   │       └── user.go 一个表对应文件
│   ├── plugin.go
│   └── README.md
└── user-web // web服务
    ├── conf
    │   ├── application-etcd.yml
    │   └── application.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── handler
    │   ├── info
    │   └── handler.go // 路由逻辑处理, 建议一个路由一个文件
    ├── main.go  // 主文件
    ├── Makefile
    ├── plugin.go
    ├── README.md
    └── router
        └── router.go // 路由

```

### 第三方框架
- go-micro
- gorm
- gin


## TODO
- [ ] 动态配置文件
- [x] 日志持久化
- [ ] 钉钉机器人报警
- [ ] Swagger 集成
- [ ] 熔断、降级、容错与健康检查
- [ ] 链路追踪

## 开发规范

### 编码规范
变量命名规则
- 表明字符类型的字符应写在变量名之前
- 使用驼峰命名,不要使用下划线
- 状态类型 使用Status与Enable结尾
- 用于判断的bool类型的变量使用变量名为ok
- 复数避免使用[s],可以使用list
- [方法] 
    - update 前缀
    - 删除方法使用 delete 前缀
    - 插入方法使用 save 
    - 获取单个数据方法使用 get 前缀
    - 获取多个数据方法使用 list 前缀
    - 统计方法使用 count 前缀
    - 判断方法使用 is 前缀
- 建议每个方法有相应的注释

### git [TODO]



### 数据库 [TODO]
- 每个表必须要有update_at字段与create_at 字段
- 所有表名为单数
- 所有表需要有表注释与字段注释


字段值 |字段名 |例子| 备注
---|---|---|---
_status | 状态值 | audit_status |注释: 字段注释:1=状态1,2=状态2,3=状态3; 一般以1为起始值,0为全部值; 
_enable |开关值| order_enable |确定是bool值时使用
_list |数组| image_list|
_at| 时间类型|create_at|时间(字段类型timestamp)|
_time|时间类型|create_time|时间戳(字段类型 int)| 
_file|文件类型|....|文件类型


### web 服务
- 每个请求尽量绑定实体接收,方便以后使用 Swagger
- 每个错误类型尽量抛出详细错误,尽量不实现errno里的公共错误码
- web/hander 里,建议每个路由使用一个文件
- 网关服务 如果使用了网关服务:把添加路由的根节点为服务名
```
例如

服务名为:go.micro.api.user
服务名不能由下划线构成
则该服务的根路由为 /user/v1/user/name
```

### 网关服务
当引用网关服务后,所有的web服务将通过网关服务进行转发,在网关服务中,可以添加自定义


### srv 服务
- srv/model 一个表对应一个文件



### proto 

- 所有的 response 都要组合BaseResponse
```
message Error{
    int32 code = 1; // 错误码
    string message = 2; // 错误信息
}

message BaseResponse{
    bool success = 1; // 是否成功 true:成功,false:失败
    Error error = 2; // error 
}

```

- message 的字段后面跟上注释
- 字段使用驼峰命名,公共的使用大驼峰,内部使用小驼峰
- rpc 结构后跟上注释