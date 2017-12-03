package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"math"
	
)

type BranchMaster struct {

	Id       int64 
	Name  string
	Address1  string
	Address2  string
	Locations  string
	City  string
	District  string
	Pincode  string
	Phone  string
	Email  string
	Fax  string
	Priority  string
	Status string
	Tax string
	CheckoutType string
	GraceTime string
}

func init() {
	 orm.RegisterModel(new(BranchMaster))
}


func AddBranch(branch BranchMaster) (Id int64) {
	o := orm.NewOrm() 

    if branch.Id>0 {
    	_, err := o.Update(&branch)
    	if err == nil{ 
    		return branch.Id
    	}
    	return branch.Id	
    }else{
    	/*
    	id, err := o.Insert(&user) // return last inserted id
    	*/
    	branch.Status="A"
		Id, _ := o.Insert(&branch)
		return Id
	}	
    

}

func GetBranch(uid string) (u *BranchMaster, err error) {
	
	var branch BranchMaster
	err = orm.NewOrm().QueryTable("BranchMaster").Filter("id", uid).One(&branch)
	return &branch,err
}

func GetAllBranches(params map[string]interface{}, sortby string, order string, page int64, limit int64) ([] *BranchMaster, map[string]interface{}) {
	
	var BranchList []*BranchMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(BranchMaster))	

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
	num, _:=qs.All(&BranchList)

	

	totalPages:= math.Ceil(float64(Total)/float64(limit))

	pager := map[string]interface{}{
		"totalPages": totalPages,
		"totalRecords":  Total,
		"startSI": offset+1,
		"endSI": offset+num,
		"page": page,
	}

	return BranchList,pager
}


func DeleteBranch(did string) (count int64) {
	//delete(UserList, uid)
	var branchId1 int64
	var branchId2 int64
	orm.NewOrm().Raw("SELECT count(id) as total from product_master WHERE branch_id=? ", did).QueryRow(&branchId1)
	

	orm.NewOrm().Raw("SELECT count(id) as total from order_master WHERE branch_id=? ", did).QueryRow(&branchId2)
	

	if branchId1==0 && branchId2==0{
		num, err:= orm.NewOrm().QueryTable("BranchMaster").Filter("Id", did).Delete()
		//fmt.Printf("Affected Num: %s, %s", num, err)
		if err == nil{ 
			return num
		}
	}	

	return 0

}

func UpdateBranchStatus(uid string){	
	var branch BranchMaster
	
	err := orm.NewOrm().QueryTable("BranchMaster").Filter("Id", uid).One(&branch)
	
	if err == nil{
		if branch.Status=="A" {
			branch.Status="D"
		}else{
			branch.Status="A"
		}

		_, err := orm.NewOrm().Update(&branch)
		fmt.Println(err); 
	}	
	
}

func UpdateBranchPriority(id int64, priority string){	
	orm.NewOrm().Raw("UPDATE branch_master SET priority = ? WHERE id = ?", priority, id).Exec()
	
}

func GetBranchOptions() [] *BranchMaster {
	var BranchList []*BranchMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(BranchMaster))
	qs = qs.OrderBy("Name")
	num, _:= qs.All(&BranchList)	
	fmt.Println(num)
	
	return BranchList
	
}
