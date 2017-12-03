package models

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	//"strings"
	"math"
	"strconv"
	"bytes"
	
)




type OrderMaster struct {	
	Id       int64 
	CreatedDate  time.Time 	
	BranchId string
	FromId string
	ToId string
	Amount string 
	PaymentDate string 
	Remarks string
	Status string
	Type string
	RefNumber string
}


type OrderList struct {	
	Id       int64	
	Branch string
	Product1 string
	Product2 string
	Amount string 
	PaymentDate string 	
	Type string
}


type OrderDetails struct {	
	Id       int64 
	OrderId  int64 	
	ProductId int64
	Date time.Time 
	Amount string 
	PaymentDate string 
	Remarks string
	Type string
	SubNarration string
	IsOb string
	RefNumber string
	DayBook string
	SecondProductId int64
}



func init() {
	 orm.RegisterModel(new(OrderMaster))
	 orm.RegisterModel(new(OrderDetails))
}


func AddOrder(Order OrderMaster) (Id int64) {
	var OrderDetail OrderDetails
	var OrderDetail2 OrderDetails
	var ProductTo ProductMaster
	var FromId int64
	var ToId int64
	var ProductFrom ProductMaster

	t := time.Now()
	t.Format("2006-01-02 15:04:05") // convert time.Time
	Order.CreatedDate=t

	value  := Order.PaymentDate  
    layout := "02-01-2006"
	p, _ := time.Parse(layout, value) // string to time.Time
	//fmt.Println(p); 
	ps:=p.Format("2006-01-02") //time to string
	//fmt.Println(ps); 
	Order.PaymentDate=ps 

	o := orm.NewOrm() 


	if i, err := strconv.ParseInt(Order.FromId, 10, 64); err == nil {
		FromId=i			
		err = orm.NewOrm().QueryTable("ProductMaster").Filter("id", i).One(&ProductTo)
	}

	if j, err := strconv.ParseInt(Order.ToId, 10, 64); err == nil {
		ToId=j			
		err = orm.NewOrm().QueryTable("ProductMaster").Filter("id", j).One(&ProductFrom)
	}

	if Order.Type=="receipt" {
		OrderDetail.DayBook="1"
	}else{
		OrderDetail2.DayBook="1"
	}


    if Order.Id>0 {
    	_, err := o.Update(&Order)
    	orm.NewOrm().Raw("UPDATE order_details SET product_id=?, payment_date=?, amount = ?, type=?, Remarks=?, sub_narration=?, ref_number=?, day_book=?, second_product_id=? WHERE order_id=? AND type=?",FromId, ps, Order.Amount, "Dr", Order.Remarks, "To: ", Order.RefNumber, OrderDetail.DayBook, ToId, Order.Id, "Dr" ).Exec()
    	orm.NewOrm().Raw("UPDATE order_details SET product_id=?, payment_date=?, amount = ?, type=?, Remarks=?, sub_narration=?, ref_number=?, day_book=?, second_product_id=? WHERE order_id=? AND type=?",ToId, ps, Order.Amount, "Cr", Order.Remarks, "By: ", Order.RefNumber, OrderDetail2.DayBook, FromId, Order.Id, "Cr" ).Exec()
    	
    	if err != nil {
			return Order.Id
		}
    	return Order.Id
    }else{
    	/*
    	id, err := o.Insert(&user) // return last inserted id
    	*/
    	Order.Status="A"    	
		Id, _ := o.Insert(&Order)

        if FromId >0 {
			OrderDetail.OrderId=Id
			OrderDetail.ProductId=FromId
			OrderDetail.Date=t
			OrderDetail.PaymentDate=ps
			OrderDetail.Amount=Order.Amount
			OrderDetail.Remarks=Order.Remarks
			OrderDetail.Type="Dr"
			OrderDetail.SubNarration="To: "
			OrderDetail.SecondProductId=ToId
			OrderDetail.IsOb="0"
			OrderDetail.RefNumber=Order.RefNumber

			

			_, err1 := o.Insert(&OrderDetail)
			if err1 != nil {
				return Id
			}	
		}

		
		
		if ToId >0 {	
			OrderDetail2.OrderId=Id
			OrderDetail2.ProductId=ToId
			OrderDetail2.Date=t
			OrderDetail2.PaymentDate=ps
			OrderDetail2.Amount=Order.Amount
			OrderDetail2.Remarks=Order.Remarks
			OrderDetail2.Type="Cr"
			OrderDetail2.SubNarration="By: "
			OrderDetail2.SecondProductId=FromId
			OrderDetail2.IsOb="0"
			OrderDetail2.RefNumber=Order.RefNumber

			
			_, err2 := o.Insert(&OrderDetail2)

			if err2 != nil {
				return Id
			}
		}
		
		return Id	
	}	
    

}



func GetOrder(uid string) (u *OrderMaster, err error) {
	
	var Order OrderMaster
	err = orm.NewOrm().QueryTable("OrderMaster").Filter("id", uid).One(&Order)
	return &Order,err
}

func GetAllOrders(params map[string]interface{}, sortby string, order string, page int64, limit int64) ([] *OrderList, map[string]interface{}) {
	
	var queryBuffer bytes.Buffer
	var selectBuffer bytes.Buffer

	var OrderLists []*OrderList
	o := orm.NewOrm()	


	offset:=(page-1)*limit
	
	switch sortby{
		case "PaymentDate":
			sortby="payment_date"
		break;
		case "Amount":
			sortby="amount"
		break;
		case "Source":
			sortby="P1.title"
		break;
		case "Towords":
			sortby="P2.title"
		break;
	case "Branch":
			sortby="B.name"
		break;
	}

	//join query preparation
	_, err:= queryBuffer.WriteString(" FROM order_master AS O INNER JOIN product_master AS P1 ON P1.id = O.from_id INNER JOIN product_master AS P2 ON P2.id = O.to_id INNER JOIN branch_master AS B ON B.id = O.branch_id WHERE 1")
	if err != nil {
		return nil, nil
	}

	//query conditions
	sCondition := make([]interface{}, 0, 9)	
	if params["BranchId"] != "" {
		_, err := queryBuffer.WriteString(" AND O.branch_id=?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["BranchId"])
	}
	if params["ProductId"] != "" {
		_, err := queryBuffer.WriteString(" AND (O.from_id=?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["ProductId"])

		_, err = queryBuffer.WriteString(" OR O.to_id=?)")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["ProductId"])
	}

	if params["PaymentDate"] != "" {		
		_, err := queryBuffer.WriteString(" AND O.payment_date=?")
		if err != nil {
			return nil, nil
		}
		sCondition = append(sCondition, params["PaymentDate"])
	}



	//count query
	_, err = selectBuffer.WriteString("SELECT COUNT(O.id) as total")
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
	_, err = selectBuffer.WriteString("SELECT O.id,O.amount,O.type,O.payment_date,P1.title as product1,P2.title as product2,B.Name as branch ")
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
	OrderListSql := selectBuffer.String()


	num, err := o.Raw(OrderListSql,sCondition...).QueryRows(&OrderLists)
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
	return OrderLists,pager
}




func UpdateOrderStatus(uid string){	
	var Order OrderMaster
	
	err := orm.NewOrm().QueryTable("OrderMaster").Filter("Id", uid).One(&Order)
	
	if err == nil{
		if Order.Status=="A" {
			Order.Status="D"
		}else{
			Order.Status="A"
		}

		_, err := orm.NewOrm().Update(&Order)
		fmt.Println(err); 
	}	
	
}

func UpdateOrderPriority(id int64, priority string){	
	orm.NewOrm().Raw("UPDATE order_master SET priority = ? WHERE id = ?", priority, id).Exec()
	
}


func DeleteOrder(uid string) {
	//delete(UserList, uid)	
	num, err := orm.NewOrm().QueryTable("OrderMaster").Filter("Id", uid).Delete()
	fmt.Printf("Affected Num: %s, %s", num, err)

	num, err = orm.NewOrm().QueryTable("OrderDetails").Filter("OrderId", uid).Delete()
	fmt.Printf("Affected Num: %s, %s", num, err)

	
	// DELETE FROM user WHERE name = "slene"
}








