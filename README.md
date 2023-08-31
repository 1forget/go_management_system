## Setup

### 1. Clone项目并添加依赖
```shell
git clone https://github.com/Godzizizilla/Management-System.git
```

### 2. 配置`config.yml`
配置文件在./config/config.yml


### 3. 启动服务
```shell
go run server/main.go 
```
## API
### /user/Login      用于学生端或管理员登录（public） 
### /user/CreateUser 用于学生端注册（public）
### /user/UpdateUser 用于学生端修改信息 （private）
### /user/DeleteUser 用于学生端删除个人信息 （private）
### /user/GetUser    用于学生端获取个人信息

### /admin/Update    用于管理员修改学生信息（private）
### /admin/getAllUser    用于管理员获取所有学生信息（private）

## 特点
### 采用redis缓存过期token解决token无状态的特性
### 在PostgreSQL中用timestamptz 类型定义createdAt PostgreSQL中比Mysql多了许多字段类型
### 比较清晰的分包结构吧
### 个人觉得token是项目重中之中  实现了统一的认证和鉴权 相当于一个小网关
### 获取 X-Forwarded-For 头部 将token绑定IP 可能可以避免一些攻击
### 写出了比较完善的swagger文档 并完全全部测试

## 不足之处
1. 将配置文件连带着加进代码仓库中 可以在部署环境中用一份yaml代替
2. 一些变量或者包的命名不太规范 （go中可以通过包名调用方法 可能UserDao或者AdminDao需要在方法名字上多加区分）
<<<<<<< HEAD
3. 因为有事耽搁了 到最近两天才写这个练手项目 拖的有点久
4. 项目比较简单 没有service层
5. 或许可以把路由信息全部抽取出来 而不是分散在多个controller中？
6. 还是不怎么会在go中做单元测试 属于是遇到问题再做测试 没有像在springboot中集成测试
7. git中commit界面卡住 git提交信息单一

### 未来改进
1. 尝试将swagger调用结果保存
2. 尝试改进并保存现有框架 作为日后开发的基础（便于导包什么的...）
=======
3. 在使用go-swagger中 只生成了一个登录用的接口 个人感觉这种手撸swagger的方式有的麻烦而且前端也不好改 之前用习惯了apifox定义接口文档并直接测试 感觉会比这个更直接 所以就没手撸那么多文档 想看看有没有自动生成swagger的工具
4. 感觉在定义swagger文档中 可以用一个结构体对应一个请求 这样可能更方便些
5. 因为有事耽搁了 到最近两天才写这个练手项目 拖的有点久
6. 学生端的token保存的信息太多 但是管理员有一些数据不用到 比如学号 不能做到用一样的方法检验token 
7. 或许可以把路由信息全部抽取出来 而不是分散在多个controller中？
8. 还是不怎么会在go中做单元测试 属于是遇到问题再做测试 没有像在springboot中集成测试
9. 登录就直接获得token 可能需要前端配合着做一下或者说可能一个用户可以获得多个token
>>>>>>> f8dcf1926eb6f23ffb430374c796762e5bb09f4e
