### 注意事项
考虑到stargo是一个工具集合、在为服务的场景下，结合go work 。

本示例使用了 [gostar](https://github.com/starfork/gostar) 生成的代码，因此遵循一下目录结构

```
.....
├── service # your services
│   ├── examples 当前项目所在目录
├── go.work  
```

### 关于 makefile使用

因为各自习惯。makefile中 仅能使用

`make proto` #需要安装protoc相关

`make test`

其他的命令，需要根据自己的情况调整


### 测试

examples 目录下  运行  
```
make test
```
运行结果见到
```
starting: gRPC Listener 51100
```

表示成功


[使用命令行调试工具 grpc-client-cli ](https://github.com/starfork/stargo?tab=readme-ov-file#grpc-client-cli-%E5%91%BD%E4%BB%A4%E8%A1%8C%E8%B0%83%E8%AF%95%E5%B7%A5%E5%85%B7)
```
? Choose a method: FetchEcho

Message json (type ? to see defaults): {}

```
运行FetchEcho可以得到如下内容
```
{
  "count": "2",
  "data": [
    {
      "id": 1,
      "name": "name1"
    },
    {
      "id": 2,
      "name": "name2"
    }
  ]
}

```