解析 Golang 依赖注入经典解决方案 uber/fx 理论篇：https://juejin.cn/post/7153582825399124005

解析 Golang 依赖注入经典解决方案 uber/fx 实战篇：https://juejin.cn/post/7153992019193856031

案例源码：https://github.com/ag9920/learnfx

fx gihub地址：https://github.com/uber-go/fx

fx 能为开发者提供的三大优势：

- 代码复用：方便开发者构建松耦合，可复用的组件；
- 消除全局状态：Fx 会帮我们维护好单例，无需借用 init() 函数或者全局变量来做这件事了；
- 经过多年 Uber 内部验证，足够可信。

添加 fx 的依赖需要用下面的命令：

```shell
go get go.uber.org/fx@v1

```

