# sms_server

修改 项目名 

修改 token密钥

修改 hack 文件夹的config

打包项目名 打包名



# 新建接口与服务

1. 创建接口文件

   ```
   server/api
   ```

   在此文件夹内创建大类 小类接口名

   如 api接口类型 支付功能

   ```
   server/api/api/pay
   ```



此文件用于文档请求解析

2. 控制层创建

   ```
   server/internal/controller
   ```

   在此文件夹内创建大类 小类接口名

   如 api接口类型 支付功能

   ```
   server/internal/controller/api/pay
   ```

   服务器是否有该接口在于控制层是否创建

3. 路由层

   如非独立类型 则只需在对应的 前缀添加控制类

   ![image-20231217221202255](/Users/Apple/Library/Application Support/typora-user-images/image-20231217221202255.png)

4. 服务逻辑层定义接口与注册

    1. 服务器层添加 服务定义添加注册方法

       ```go
       // server/internal/service
       
       package service
       
       import (
           "context"
           "hotgo/internal/model/input/testin"
       )
       
       type (
           ITestCheck interface {
               Add(ctx context.Context, in *testin.CheckAddInp) (err error)
           }
       )
       
       var (
           localTestCheck ITestCheck
       )
       
       func TestCheck() ITestCheck {
           if localTestCheck == nil {
               panic("implement not found for interface ISysProvinces, forgot register?")
           }
           return localTestCheck
       }
       
       func RegisterTestCheck(i ITestCheck) {
           localTestCheck = i
       }
       ```



2. 逻辑层实现方法

   ```go
   //server/internal/logic
   package test
   
   import (
       "context"
       "github.com/gogf/gf/v2/frame/g"
       "hotgo/internal/model/input/testin"
       "hotgo/internal/service"
   )
   
   type sTestCheck struct{}
   
   func NewTestCheck() *sTestCheck {
       return &sTestCheck{}
   }
   
   func init() {
       service.RegisterTestCheck(NewTestCheck())
   }
   
   func (s *sTestCheck) Add(ctx context.Context, in *testin.CheckAddInp) (err error) {
       g.Log().Debug(ctx, "服务逻辑层", in)
   
       return
   }
   
   ```

   并且在服务层文件载入

   server/internal/logic/logic.go

3. 期间要在模型层定义好输入数据的类型

   ```
   server/internal/model/input 输入流
   server/internal/model/entity 数据库模型
   server/internal/model/do 执行模型
   ```

      

# 添加自定义插件组件

https://github.com/bufanyun/hotgo/blob/v2.0/docs/guide-zh-CN/addon-flow.md

1 在开发工具插件管理添加 自动创建插件文件夹 到 指定项目位置

2 开发流程与主服务相同 访问路径不同

3 添加插件接口同主服务相同 数据类型建议继承主服务的input

**创建流程**

1 使用系统配置创建新插件 

2 自己建表 添加模式为插件模式添加entity

3 添加配置文件组件到 web的插件view路径

4  

### 访问路径 必需跟插件名称相同 大小写敏感

#### 后台插件访问路径

```
// IP+端口或域名/admin/插件名称/API路径
如：127.0.0.1:8000/admin/hgexample/index/test
```



- 对应控制器路径：`server/addons/hgexample/controller/admin/sys/index.go`

#### 前端API插件访问路径

```
// IP+端口或域名/api/插件名称/API路径
如：127.0.0.1:8000/api/hgexample/index/test
```



- 对应控制器路径：`server/addons/hgexample/controller/api/index.go`

#### 前台页面插件访问路径

```
// IP+端口或域名/home/插件名称/API路径
如：127.0.0.1:8000/home/hgexample/index/test
```



- 对应控制器路径：`server/addons/hgexample/controller/home/index.go`
- 对应模板路径：`server/addons/hgexample/resource/public/template`

#### 静态资源插件访问路径

```
// IP+端口或域名/home/插件名称/API路径
如：127.0.0.1:8000/addons/hgexample/default
```



- 对应资源路径：`server/addons/hgexample/resource/public`

#### Websocket插件访问路径

```
// IP+端口或域名/socket/插件名称/API路径
如：127.0.0.1:8000/socket/hgexample/index/test
```



- 对应控制器路径：`server/addons/hgexample/controller/socket/index.go`

