

ent cli工具安装

官方文档：https://entgo.io/zh/docs/getting-started/

```shell
go install entgo.io/ent/cmd/ent@latest
```

### 创建schema
`--target`：指定创建的schema所在目录
`User`: 创建的schema

```shell
ent new --target ./schema User
```

### 生成代码
`./schema`: schema 所在目录
`--target`：生成代码所在目录
```shell
go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --feature sql/lock --target ./ent
```
一般是将生成命令单独写在一个文件中如：
```go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --feature sql/lock --target ./ent
```
然后执行命令`go generate ./...` 来生成