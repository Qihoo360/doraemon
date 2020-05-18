package controllers

//import (
//	"github.com/astaxie/beego"
//)
//
//type LabelController struct {
//	beego.Controller
//}

//func (c *LabelController) URLMapping() {
//	c.Mapping("GetAllLabel", c.GetAllLabel)
//	c.Mapping("AddLabel", c.AddLabel)
//	c.Mapping("DeleteLabel", c.DeleteLabel)
//}
//
//// @router / [get]
//func (c *LabelController) GetAllLabel(){
//	var Label *models.Labels
//	labels:=Label.GetAll()
//	c.Data["json"] = &common.Res{0,"",labels}
//	c.ServeJSON()
//}
//
//// @router / [post]
//func (c *LabelController) AddLabel(){
//	Label:=models.Labels{}
//	var ans common.Res
//	value:= struct {
//		Value string
//	}{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &value)
//	if err != nil {
//		logs.Error("Unmarshal plan error:%v", err)
//		ans.Code = 1
//		ans.Msg = "Unmarshal error"
//	} else {
//		Label.Label=value.Value
//		err=Label.AddLabel()
//		if err!=nil{
//			ans.Code = 1
//			ans.Msg = "插入数据库错误：" + err.Error()
//		}
//	}
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}
//
//// @router /:id [delete]
//func (c *LabelController) DeleteLabel(){
//	labelid := c.Ctx.Input.Param(":id")
//	var Label *models.Labels
//	var ans common.Res
//	err:=Label.DeleteLabel(labelid)
//	if err != nil {
//		ans.Code=1
//		ans.Msg = "数据库删除记录错误：" + err.Error()
//	}
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}
