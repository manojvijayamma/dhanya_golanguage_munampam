package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"math"
	
)

type User struct {	
	Id       int 
	Name  string `orm:"size(100)"`
	Email string 
	Mobile string	
	Address string
	Password string
	Username string
	Status string
	Priority string
	GroupId string `orm:"size(100)"` // in db it should be group_id
}

type Changepassword struct {	
	NewPassword string
	Password string	
	UserId int
}



func init() {
	 orm.RegisterModel(new(User))
}


func AddUser(u User) (error) {
	o := orm.NewOrm() 

    if u.Id>0 {
    	_, err := o.Update(&u)
    	return err
    }else{
    	/*
    	id, err := o.Insert(&user) // return last inserted id
    	*/
    	u.Status="A"
		_, err := o.Insert(&u)
		return err
	}	
    

}

func GetUser(uid string) (u *User, err error) {
	
	var user User
	err = orm.NewOrm().QueryTable("user").Filter("id", uid).One(&user)
	return &user,err
}

func GetAllUsers(params map[string]interface{}, sortby string, order string, page int64, limit int64) ([] *User, map[string]interface{}) {
	
	var UserList []*User
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(User))	

	//where
	for k, v := range params {
		// rewrite dot-notation to Object__Attribute
		if v!=""{
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}	
	}

	//get total
	Total, _:=qs.Count()

	//sorting
	if order=="DESC"{
		sortby="-"+sortby
	}
	qs = qs.OrderBy(sortby)
	offset:=(page-1)*limit
	qs=qs.Limit(limit, offset)
	num, _:=qs.All(&UserList)

	

	totalPages:= math.Ceil(float64(Total)/float64(limit))

	pager := map[string]interface{}{
		"totalPages": totalPages,
		"totalRecords":  Total,
		"startSI": offset+1,
		"endSI": offset+num,
		"page": page,
	}

	return UserList,pager
}


func DeleteUser(did string) {
	//delete(UserList, uid)
	
	num, err := orm.NewOrm().QueryTable("User").Filter("Id", did).Delete()
	fmt.Printf("Affected Num: %s, %s", num, err)
	// DELETE FROM user WHERE name = "slene"

}

func UpdateUserStatus(uid string){	
	var user User
	
	err := orm.NewOrm().QueryTable("User").Filter("Id", uid).One(&user)
	
	if err == nil{
		if user.Status=="A" {
			user.Status="D"
		}else{
			user.Status="A"
		}

		_, err := orm.NewOrm().Update(&user)
		fmt.Println(err); 
	}	
	
}

func UpdateUserPriority(id int64, priority string){	
	orm.NewOrm().Raw("UPDATE User SET priority = ? WHERE id = ?", priority, id).Exec()
	
}


func DoLogin(u User) (a *User, err error) {
	//o := orm.NewOrm()
	//user := User{Username: u.Username} 
	//fmt.Println(u); 
	//err= o.Read(&user)
	var user User
	err = orm.NewOrm().QueryTable("user").Filter("username", u.Username).Filter("password",u.Password).One(&user)
	

	//o.Raw("select *  from user where username = ? and password=?", u.Username,u.Password).QueryRow(&u)



	/*
	if err1 == orm.ErrNoRows {
    	fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.Name)
	}
	*/

    return &user,err

}


func CheckValidUser(userId string) (a *User, err error) {
	//o := orm.NewOrm()
	//user := User{Username: u.Username} 
	//fmt.Println(u); 
	//err= o.Read(&user)
	var user User
	err = orm.NewOrm().QueryTable("user").Filter("Email", userId).One(&user)
	

	//o.Raw("select *  from user where username = ? and password=?", u.Username,u.Password).QueryRow(&u)



	/*
	if err1 == orm.ErrNoRows {
    	fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.Name)
	}
	*/

    return &user,err

}

func ChangePassword(u Changepassword) (bool) {
	var user User
	_ = orm.NewOrm().QueryTable("user").Filter("Password", u.Password).Filter("Id", u.UserId).One(&user)
	if user.Id>0 {
		orm.NewOrm().Raw("UPDATE User SET password = ? WHERE id = ?", u.NewPassword, user.Id).Exec()
		return true
	}
	return false
}
