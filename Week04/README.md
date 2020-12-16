#### Week04 作业题目：

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

代码来自 [kratos](https://github.com/go-kratos/kratos) 生成的 demo

**如何运行**
```
cd Week04/cmd
go build
./cmd -conf ../configs
```
打开浏览器访问 http://localhost:8000/Week04/start

你会看到输出了
`Terry I Love You!`

#### 参考 go-kratos 的文档

[kratos 快速上手](https://github.com/go-kratos/kratos/blob/master/docs/quickstart.md)

[blademaster](https://github.com/go-kratos/kratos/blob/master/docs/blademaster.md) (HTTP框架)

[warden](https://github.com/go-kratos/kratos/blob/master/docs/warden.md) (改良的gRPC框架)

不得不说，b站起名字真够中二的！哈哈哈！

