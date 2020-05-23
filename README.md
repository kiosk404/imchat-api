# imchat-api
基于 Beego + websocket 实现的IM聊天程序。目前还算是十分简陋的，还需要不断的优化

配置在conf中，主要依赖Mysql。
- 启动
``` go
go run main.go
```
- 其他
默认监听8080端口，如果配合有`Nginx`反向代理. 需要将 /ws 路径配置websocket代理。
```cassandraql
        location /ws {
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header X-real-ip $remote_addr;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_pass http://im_backend;
        }

```

#Todo
- [ ] 服务端鉴权
- [ ] 智能心跳机制
- [ ] 消息稳定性, 应用层ACK, 消息序号生成器 
- [ ] 历史消息保存, 一致性Hash   
- [ ] 前端优化
- [ ] Electron开发

