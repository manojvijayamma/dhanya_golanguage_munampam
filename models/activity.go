package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	
)

var (
	ActivityList []*ActivityMaster
)

func init() {
	 orm.RegisterModel(new(ActivityMaster))
}

type ActivityMaster struct {

	Id       int `orm:"auto"`
	Title  string `orm:"size(100)"`
	ActivityId string
	
}





func GetActivityForAllocation() [] *ActivityMaster {
	/*
	qs, _ := o.QueryTable(new(Ticket)).Filter("EventId", 2).All(&tickets)
	*/
	o := orm.NewOrm()	
	

	num, err := o.Raw("SELECT AM.id, AM.title,LA.activity_id FROM activity_master as AM left join link_activities  as LA on AM.id=LA.activity_id WHERE AM.status = ?", "A").QueryRows(&ActivityList)
	if err == nil {
	    fmt.Println("user nums: ", num)
	}
	
	return ActivityList
	
}



