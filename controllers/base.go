package controllers

import (
    "myapi/models"
	"errors"
	"strings"
	"github.com/astaxie/beego"
    
	jwt "github.com/dgrijalva/jwt-go"   

)

// Operations about Users
type BaseController struct {
	beego.Controller
}

func (c *BaseController) ParseToken() (t *jwt.Token, e error) {
    authString := c.Ctx.Input.Header("Authorization")

    errInputData:=errors.New("User Not Exist")
    errExpired:=errors.New("User Not Exist")
    beego.Debug("AuthString:", authString)

    kv := strings.Split(authString, " ")
    if len(kv) != 2 || kv[0] != "Bearer" {
        beego.Error("AuthString invalid:", authString)
        return nil, errInputData
    }
    tokenString := kv[1]

    // Parse token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("mykey"), nil
    })
    if err != nil {
        beego.Error("Parse token:", err)
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                // That's not even a token
                return nil, errInputData
            } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
                // Token is either expired or not active yet
                return nil, errExpired
            } else {
                // Couldn't handle this token
                return nil, errInputData
            }
        } else {
            // Couldn't handle this token
            return nil, errInputData
        }
    }
    if !token.Valid {
        beego.Error("Token invalid:", tokenString)
        return nil, errInputData
    }
    beego.Debug("Token:", token)
    return token, nil
}

func (c *BaseController) isAuthUser() (t int) {
    
    token, e := c.ParseToken()
    if e != nil {
        
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        
    }
    var userId string = claims["userId"].(string)
   
    beego.Error("db user:", userId)
    if userId!="" {
        result,_:=models.CheckValidUser(userId)        

        if result.Id>0{
            return result.Id

        }
   }

    c.Ctx.Output.SetStatus(401)
    c.Data["json"] = map[string]interface{}{"status": 401, "message": "login success "}    
    c.ServeJSON()

    return 

}





