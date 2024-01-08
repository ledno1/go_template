# sms_server



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

      
