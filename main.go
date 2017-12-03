package main

import (
	_ "myapi/routers"
		"os/exec"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    //"github.com/astaxie/beego/context"
    "github.com/astaxie/beego/plugins/cors"
    "time"
)

func init() {	
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //orm.RegisterDataBase("default", "mysql", "root:root@/angulargo_db?charset=utf8")  
   // orm.RegisterDataBase("default", "mysql", "root:@/mini_tally_db?charset=utf8")  
   orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/mini_tally_db")  

}


func main() {
	
	orm.Debug = true
	/*
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {	    
	        ctx.Output.Header("Access-Control-Allow-Origin", "*")
	        ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	        ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	        ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept,  Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,Access-Control-Allow-Origin")
	        ctx.Output.Header("Access-Control-Expose-Headers", "X-Pagination-Current-Page")
	        //ctx.Abort(200, "Hello")	    
	})*/

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin","Authorization","X-Requested-With"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	if beego.BConfig.RunMode == "dev" {		
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		
	}
	
 


	beego.Run()


		go func() {
             <- time.After(100 * time.Millisecond)
             exec.Command("xdg-open", "http://example.com/").Run()
        }()
       
}
