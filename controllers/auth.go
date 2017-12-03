package controllers

import (
	"myapi/models"
	"encoding/json"	
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

// Operations about Users
type AuthController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AuthController) Post() {
	var user models.User

	u.CheckTialPeriod()

	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//fmt.Println(user)
	result,_:=models.DoLogin(user)
	//u.Data["json"] = map[string]string{"uid": uid}
	fmt.Println("user nums: ", result)
	if result.Id==0 {
		//u.Ctx.Output.SetStatus(401)
		u.Data["json"] = map[string]interface{}{"status": 200, "message": "Invalid Credentials"}	
		u.ServeJSON()
	}

	claims := make(jwt.MapClaims)
	claims["userId"] = result.Email
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("mykey"))

	if err != nil {              
    
    }
    u.Data["json"] = map[string]interface{}{"status": 200, "message": "login success ", "token": tokenString,"name":result.Name}     
	
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *AuthController) GetAll() {
	//users := models.GetAllUsers()
	u.Data["json"] = ""
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *AuthController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}



// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *AuthController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}



// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *AuthController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}



func (c *AuthController) CheckTialPeriod(){
   /* t := time.Now().UTC()
    t.Format("2006-01-02 15:04:05")

    value  := "01-10-2017" 
    layout := "02-01-2006"
    e, _ := time.Parse(layout, value)

    diff := e.Sub(t)
    days := int(diff.Hours() / 24)
    //fmt.Println(days); 
    if days < 0 {
         c.Ctx.Output.SetStatus(200)
         c.Data["json"] = map[string]interface{}{"status": 200, "message": "Your trial period for this app has expired."}    
         c.ServeJSON()
    }
    */
}


