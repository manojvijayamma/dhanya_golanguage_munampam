package controllers

import (
	"myapi/models"	
	"fmt"
	
	


)

// Operations about Users
type ActivityController struct {
	BaseController
}





// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *ActivityController) GetAll() {
	userId:=this.isAuthUser()
	fmt.Println(userId)
	groups := models.GetActivityForAllocation()
	this.Data["json"] = &map[string]interface{}{"result": groups}
	this.ServeJSON()
}









