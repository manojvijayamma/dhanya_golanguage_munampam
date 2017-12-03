package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"math"
	
)


type CategoryMaster struct {	
	Id       int64 `orm:"pk"`
	Title  string 	
	Status string
	ParentId string 
	Priority string
}

func init() {
	 orm.RegisterModel(new(CategoryMaster))
}


func AddCategory(u CategoryMaster) (Id int64) {
	o := orm.NewOrm()
    if u.Id>0 {
    	_, err := o.Update(&u)
    	if err == nil{ 
    		return u.Id
    	}
    	return u.Id	
    }else{    	
    	u.Status="A"
		Id,_:= o.Insert(&u)
		return Id
	}
}

func GetCategory(uid string) (u *CategoryMaster, err error) {
	
	var category CategoryMaster
	err = orm.NewOrm().QueryTable("CategoryMaster").Filter("id", uid).One(&category)
	return &category,err
}


func GetAllCategory(params map[string]interface{}, sortby string, order string, page int64, limit int64) ([] *CategoryMaster, map[string]interface{}) {
	
	var CategoryList []*CategoryMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(CategoryMaster))	

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
	num, _:=qs.All(&CategoryList)

	

	totalPages:= math.Ceil(float64(Total)/float64(limit))

	pager := map[string]interface{}{
		"totalPages": totalPages,
		"totalRecords":  Total,
		"startSI": offset+1,
		"endSI": offset+num,
		"page": page,
	}

	return CategoryList,pager
}






func DeleteCategory(uid string) (count int64){
	//delete(UserList, uid)	
	var CategoryId int64
	orm.NewOrm().Raw("SELECT count(id) as total from product_master WHERE category_id=? ", uid).QueryRow(&CategoryId)
	
	if CategoryId==0 {
		num, err := orm.NewOrm().QueryTable("CategoryMaster").Filter("Id", uid).Delete()
		if err == nil{ 
			return num
		}
	}
	return 0
}


func UpdateCategoryStatus(uid string){	
	var category CategoryMaster	
	err := orm.NewOrm().QueryTable("CategoryMaster").Filter("Id", uid).One(&category)	
	if err == nil{
		if category.Status=="A" {
			category.Status="D"
		}else{
			category.Status="A"
		}

		_, err := orm.NewOrm().Update(&category)
		fmt.Println(err); 
	}	
	
}

func UpdateCategoryPriority(id int64, priority string){	
	orm.NewOrm().Raw("UPDATE category_master SET priority = ? WHERE id = ?", priority, id).Exec()
	
}

func GetCategoryOptions() [] *CategoryMaster {
	var CategoryList []*CategoryMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(CategoryMaster))
	qs = qs.OrderBy("Title")
	num, _:= qs.All(&CategoryList)	
	fmt.Println(num)
	
	return CategoryList
	
}



