// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"myapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/branch",
			beego.NSInclude(
				&controllers.BranchController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/usergroup",
			beego.NSInclude(
				&controllers.UserGroupController{},
			),
		),
		beego.NSNamespace("/links",
			beego.NSInclude(
				&controllers.LinksController{},
			),
		),
		beego.NSNamespace("/activity",
			beego.NSInclude(
				&controllers.ActivityController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/receipt",
			beego.NSInclude(
				&controllers.ReceiptController{},
			),
		),
		
		
	)
	beego.AddNamespace(ns)
}
