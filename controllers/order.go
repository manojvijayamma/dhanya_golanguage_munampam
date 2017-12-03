package controllers

import (
	"myapi/models"	
	"encoding/json"
	"strconv"
	"time"
	
)

// Operations about Users
type OrderController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.OrderMaster	true		"body for user content"
// @Success 200 {int} models.OrderMaster.Id
// @Failure 403 body is empty
// @router / [post]
func (this *OrderController) Post() {
   
	var userId=this.isAuthUser()
	var LogMaster models.LogMaster
	
	var Order models.OrderMaster
	json.Unmarshal(this.Ctx.Input.RequestBody, &Order)
	Order.Type="payment"
	LogMaster.RowId=models.AddOrder(Order)

	LogMaster.UserId=userId
	LogMaster.EndPoint="order"
	LogMaster.Action="create"
	if Order.Id >0 {
		LogMaster.Action="update"
	}	
	LogMaster.Data=string(this.Ctx.Input.RequestBody)	
	models.ActivityLog(LogMaster)

	this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	
}



// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *OrderController) GetAll() {
	var userId=this.isAuthUser()
	if userId>0 {}
	var limit,_  = this.GetInt64("rowsize")
	var offset,_  = this.GetInt64("page")
	var sortby  = this.GetString("sortField")
	var order  = this.GetString("sortOrder")

	var paymentDate=this.GetString("PaymentDate")
	if paymentDate!=""{
		value  := paymentDate 
    	layout := "02-01-2006"
		p, _ := time.Parse(layout, value) // string to time.Time
		//fmt.Println(p); 
		ps:=p.Format("2006-01-02") //time to string
		paymentDate=ps
	}
	

	qryparams := map[string]interface{}{
		"BranchId": this.GetString("BranchId"),
		"ProductId":  this.GetString("ProductId"),
		"PaymentDate": paymentDate,	
		
	}

	result,pager := models.GetAllOrders(qryparams,sortby,order,offset,limit)
	//this.Data["json"] = users
	this.Data["json"] = &map[string]interface{}{"result": result, "pager":pager}
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *OrderController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetOrder(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = &map[string]interface{}{"result": user}
			//u.Data["json"] = user
		}
	}
	u.ServeJSON()
}



// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /status/:uid [get]
func (u *OrderController) Status() {
	uid := u.GetString(":uid")
	models.UpdateOrderStatus(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}


// @Title Priority
// @Description priority the user
// @Param	id		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /priority [get]
func (this *OrderController) Priority() {
	id,_ := this.GetInt64("id")
	priority := this.GetString("priority")
	
	models.UpdateOrderPriority(id,priority)
	this.Data["json"] = &map[string]interface{}{"status": "success"}
	this.ServeJSON()
}


// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete/:uid [get]
func (u *OrderController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteOrder(uid)

	var userId=u.isAuthUser()
	var LogMaster models.LogMaster
	LogMaster.UserId=userId
	LogMaster.EndPoint="voucher"
	LogMaster.Action="delete"
	if j, err := strconv.ParseInt(uid, 10, 64); err == nil {
			LogMaster.RowId=j
		}
	models.ActivityLog(LogMaster)

	u.Data["json"]=&map[string]interface{}{"status": "success","message":"Successfully Deleted. "}
	u.ServeJSON()
}









