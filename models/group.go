package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	
)

var (
	GroupList []*Group
)

func init() {
	 orm.RegisterModel(new(Group))
}

type Group struct {
	//Id       int `orm:"auto"`
	//Username string `orm:"size(100)"`
	//Password string `orm:"size(100)"`
	Id       int `orm:"auto"`
	Title  string `orm:"size(100)"`
	
}


type GroupActivities struct {	
	Id       int `orm:"auto"`
	Title  string `orm:"size(100)"`
	
}


func AddGroup(u Group) (error) {
	o := orm.NewOrm()
   // o.Using("default") // Using default, you can use other database
	fmt.Println(u)

    if u.Id>0 {
    	_, err := o.Update(&u)
    	return err
    }else{
    	/*
    	id, err := o.Insert(&user) // return last inserted id
    	*/
    	
		_, err := o.Insert(&u)
		return err
	}	
    

}


func GetAllGroups() [] *Group {
	/*
	qs, _ := o.QueryTable(new(Ticket)).Filter("EventId", 2).All(&tickets)
	*/
	o := orm.NewOrm()	
	qs,_:=o.QueryTable(new(Group)).All(&GroupList)	
	fmt.Println(qs)
	
	return GroupList
	
}



