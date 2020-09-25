package main

import (
	"context"
	"fmt"
	"github.com/djsxianglei/iris-demo/app/api/controllers"
	"github.com/djsxianglei/iris-demo/common"
	"github.com/djsxianglei/iris-demo/repositories"
	"github.com/djsxianglei/iris-demo/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//1.创建iris实例
	app := iris.New()
	//2.设置错误模式在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./api/views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板目录
	app.HandleDir("/assets", "./api/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println("数据库连接", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//5.注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))


	//6.启动服务
	app.Run(
		iris.Addr("localhost:30080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
