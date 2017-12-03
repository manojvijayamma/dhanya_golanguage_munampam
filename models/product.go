package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"bytes"
	"math"
	"time"
	
)


type ProductMaster struct {	
	Id       int64 
	Title  string	
	OpenningBalance string
	Status string	
	Priority string	
	CategoryId string
	BranchId string
	OpeningBalDate string
	Type string
	ApplyTo string
}

type OrderProductMaster struct {	
	Id       int64 
	Title  string	
	
}


type ProductCategoryBranch struct {	
	Id       int64 
	Title  string	
	Branch string
	Category string
}	



func init() {
	 orm.RegisterModel(new(ProductMaster))
	 orm.RegisterModel(new(OrderProductMaster))
}


func AddProduct(u ProductMaster) (Id int64) {
	var OrderDetail OrderDetails
	o := orm.NewOrm()

	t := time.Now()
	t.Format("2006-01-02 15:04:05")


	value  := u.OpeningBalDate  
    layout := "02-01-2006"
	p, _ := time.Parse(layout, value) // string to time.Time
	//fmt.Println(p); 
	ps:=p.Format("2006-01-02") //time to string
	//fmt.Println(ps); 
	u.OpeningBalDate=ps 

    if u.Id>0 {
    	_, err := o.Update(&u)
    	orm.NewOrm().Raw("UPDATE order_details SET amount = ?,payment_date=?,type=? WHERE product_id= ? AND is_ob = ?", u.OpenningBalance, ps, u.Type, u.Id,"1").Exec()
    	if err==nil{}
    	return u.Id
    }else{    	
    	u.Status="A"
		Id, _ := o.Insert(&u)

		OrderDetail.ProductId=Id
		OrderDetail.Date=t
		OrderDetail.PaymentDate=ps
		OrderDetail.Amount=u.OpenningBalance
		OrderDetail.Remarks="Opening Balance"
		OrderDetail.Type=u.Type
		OrderDetail.IsOb="1"

		_, err := o.Insert(&OrderDetail)
		if err==nil{}
		return Id
	}
}

func GetProduct(uid string) (u *ProductMaster, err error) {
	
	var Product ProductMaster
	err = orm.NewOrm().QueryTable(new(ProductMaster)).Filter("id", uid).One(&Product)
	return &Product,err
}


func GetAllProduct(params map[string]interface{}, sortby string, order string, page int64, limit int64) ([] *ProductCategoryBranch, map[string]interface{}) {
	
	/*
	var ProductList []*ProductMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(ProductMaster))	

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
	qs=qs.Limit(limit, offset).RelatedSel()
	num, _:=qs.All(&ProductList)

	

	totalPages:= math.Ceil(float64(Total)/float64(limit))

	pager := map[string]interface{}{
		"totalPages": totalPages,
		"totalRecords":  Total,
		"startSI": offset+1,
		"endSI": offset+num,
		"page": page,
	}

	return ProductList,pager
	*/



	var queryBuffer bytes.Buffer
	var selectBuffer bytes.Buffer

	var ProductList []*ProductCategoryBranch
	o := orm.NewOrm()	


	offset:=(page-1)*limit
	
	switch sortby{
		case "title":
			sortby="P.title"
		break;
		case "branch":
			sortby="B.name"
		break;
		case "category":
			sortby="C.title"
		break;
		
	}

	//join query preparation
	_, err:= queryBuffer.WriteString(" FROM product_master AS P INNER JOIN category_master AS C ON P.category_id = C.id  INNER JOIN branch_master AS B ON B.id = P.branch_id WHERE 1")
	if err != nil {
		return nil, nil
	}

	//query conditions
	sCondition := make([]interface{}, 0, 9)	
	if params["BranchId"] != nil {
		_, err := queryBuffer.WriteString(" AND P.branch_id=?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["BranchId"])
	}
	if params["CategoryId"] != nil {
		_, err := queryBuffer.WriteString(" AND P.category_id=?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["CategoryId"])
	}

	if len(params["Title"].(string)) > 0 {
		_, err := queryBuffer.WriteString(" AND P.title LIKE ?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, "%"+params["Title"].(string)+"%")
	}

	


	//count query
	_, err = selectBuffer.WriteString("SELECT COUNT(P.id) as total")
	if err != nil {
		return nil, nil
	}
	_, err = selectBuffer.WriteString(queryBuffer.String())
	if err != nil {
		return nil, nil
	}
	var count int64
	countQuery := selectBuffer.String()
	o.Raw(countQuery, sCondition...).QueryRow(&count)
	selectBuffer.Reset()
	

	//query data
	_, err = selectBuffer.WriteString("SELECT P.id,P.title,B.Name as branch,C.title as category ")
	if err != nil {
		return nil, nil
	}
	_, err = selectBuffer.WriteString(queryBuffer.String())
	if err != nil {
		return nil, nil
	}
	_, err = selectBuffer.WriteString(" ORDER BY "+sortby+" "+order+" LIMIT ?,?")
	if err != nil {
		return nil, nil
	}

	sCondition = append(sCondition, offset, limit)
	ProductListSql := selectBuffer.String()


	num, err := o.Raw(ProductListSql,sCondition...).QueryRows(&ProductList)
	if err == nil {
	    fmt.Println("user nums: ", num)
	}

	

	totalPages:= math.Ceil(float64(count)/float64(limit))
	pager := map[string]interface{}{
		"totalPages": totalPages,
		"totalRecords":  count,
		"startSI": offset+1,
		"endSI": offset+num,
		"page": page,
	}
	return ProductList,pager
}






func DeleteProduct(uid string) (count int64){
	//delete(UserList, uid)	

	var ProductId int64
	orm.NewOrm().Raw("SELECT count(id) as total from order_details WHERE (product_id=? OR second_product_id=?) AND is_ob!='1'", uid,uid).QueryRow(&ProductId)

	if ProductId==0 {
		num, err := orm.NewOrm().QueryTable("ProductMaster").Filter("Id", uid).Delete()	

		num, err = orm.NewOrm().QueryTable("OrderDetails").Filter("ProductId", uid).Filter("is_ob", "1").Delete()
		
		if err == nil{ 
			return num
		}
	}

	return 0
	
	// DELETE FROM user WHERE name = "slene"
}


func UpdateProductStatus(uid string){	
	var Product ProductMaster	
	err := orm.NewOrm().QueryTable("ProductMaster").Filter("Id", uid).One(&Product)	
	if err == nil{
		if Product.Status=="A" {
			Product.Status="D"
		}else{
			Product.Status="A"
		}

		_, err := orm.NewOrm().Update(&Product)
		fmt.Println(err); 
	}	
	
}

func UpdateProductPriority(id int64, priority string){	
	orm.NewOrm().Raw("UPDATE product_master SET priority = ? WHERE id = ?", priority, id).Exec()
	
}

/*
func GetProductWithCategory() [] *ProductCategory {
	var ProductList []*ProductCategory
	o := orm.NewOrm()	
	num, err := o.Raw("SELECT P.id,P.title as product_name,C.title as category_name FROM product_master AS P LEFT JOIN category_master AS C ON P.parent_id=C.id WHERE P.status=? ", "A").QueryRows(&ProductList)
	if err == nil {
	    fmt.Println("user nums: ", num)
	}
	
	return ProductList
	
}*/

func GetProductOptions(id int64) [] *ProductMaster {
	var ProductList []*ProductMaster
	o := orm.NewOrm()	
	qs:=o.QueryTable(new(ProductMaster)).Filter("BranchId", id)
	qs = qs.OrderBy("Title")
	num, _:= qs.All(&ProductList)	
	fmt.Println(num)
	
	return ProductList
	
}



