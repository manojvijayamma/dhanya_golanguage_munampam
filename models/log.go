package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	
	"time"
	
)



func init() {
	 orm.RegisterModel(new(LogMaster))
}

type LogMaster struct {
	Id int64
	Date  time.Time
	RowId int64 	
	UserId int
	EndPoint string
	Action string
	Data string	
}


func ActivityLog(A LogMaster) {
	
	t := time.Now()
	t.Format("2006-01-02 15:04:05") // convert time.Time
	A.Date=t

	o := orm.NewOrm()	
	_, err := o.Insert(&A)
	if err == nil{ 
    		
    }
    	

}








