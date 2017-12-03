package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["myapi/controllers:ActivityController"] = append(beego.GlobalControllerRouter["myapi/controllers:ActivityController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["myapi/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["myapi/controllers:AuthController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["myapi/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["myapi/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["myapi/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:BranchController"] = append(beego.GlobalControllerRouter["myapi/controllers:BranchController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/options`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["myapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/options`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:LinksController"] = append(beego.GlobalControllerRouter["myapi/controllers:LinksController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:LinksController"] = append(beego.GlobalControllerRouter["myapi/controllers:LinksController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:LinksController"] = append(beego.GlobalControllerRouter["myapi/controllers:LinksController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:LinksController"] = append(beego.GlobalControllerRouter["myapi/controllers:LinksController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:OrderController"] = append(beego.GlobalControllerRouter["myapi/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ProductController"] = append(beego.GlobalControllerRouter["myapi/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/options`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ReceiptController"] = append(beego.GlobalControllerRouter["myapi/controllers:ReceiptController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Priority",
			Router: `/priority`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Changepassword",
			Router: `/changepassword`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserGroupController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserGroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserGroupController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserGroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserGroupController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserGroupController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserGroupController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserGroupController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserGroupController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserGroupController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
