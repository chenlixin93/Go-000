#### Week03 作业题目：
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

代码来自 @rushuinet 同学，做了一些修改和注释

windows7测试过程：
```
go build main.go // 生成可执行文件

main.exe // 执行exe文件
```

测试结果1：
```
D:\GoCode\Go-000\Week03>main.exe
开始捕捉err
httpServer start
协程1：访问/close路由，触发信号关闭1
协程2：来自信号关闭1
httpServer结束1
======= 协程2：【通知】http服务关闭
```
测试结果2：
```
D:\GoCode\Go-000\Week03>main.exe
开始捕捉err
httpServer start
协程2：监听系统信号量，触发信号关闭2
协程1：来自信号关闭2
======= 协程2：信号关闭2
```
