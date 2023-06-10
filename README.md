# xyu-graphql
基于gqlgen的graphql后端服务器搭建，可直接下载集成到自己项目，根据文档操作即可

### 启动
+ go get github.com/99designs/gqlgen
+ go run main.go

### 操作
#### 新增xxxServer 【可以自己改命名】
+ 1.更新dao/model_gen/下面的 model
+ 2.更新各个module下面的sdl的graphql文件
+ 3.更新schema下面的的graphql文件，定义接口，同时更新dao/model/base.go
+ 4.执行命令,实现接口
#### 更新项目,即更新结构体，增加属性
+ 更新model_gen下面的结构体
+ 更新各个module下面的项目的sdl
+ 更新schema下的base.graphql
+ 命令行执行gqlgen

### 特别注意
+ 1.如果在.grphql文件中，field后面没有加!，则在生成的结构体中是指针类型，保障安全性，但是性能会下降，因此需要酌情使用

![img.png](assets/img.png)
