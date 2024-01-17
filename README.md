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