package controllers

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
)

type MaintainController struct {
	beego.Controller
}

func (c *MaintainController) URLMapping() {
	c.Mapping("GetAllProms", c.GetAllMaintains)
	c.Mapping("AddProm", c.AddMaintain)
	c.Mapping("UpdateProm", c.UpdateMaintain)
	c.Mapping("DeleteProm", c.DeleteMaintain)
}

// @router / [get]
func (c *MaintainController) GetAllMaintains() {
	var Maintain *models.Maintains
	maintains := Maintain.GetAllMaintains()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: maintains,
	}
	c.ServeJSON()
}

// @router /:id/hosts [get]
func (c *MaintainController) GetHosts() {
	var Host *models.Hosts
	mid := c.Ctx.Input.Param(":id")
	hosts := Host.GetHosts(mid)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: hosts,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *MaintainController) AddMaintain() {
	var data struct {
		Flag      bool   `json:"flag"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
		Hosts     string `json:"hosts"`
	}
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		var maintain models.Maintains
		maintain.TimeStart = data.TimeStart
		maintain.TimeEnd = data.TimeEnd
		maintain.DayStart = data.DayStart
		maintain.DayEnd = data.DayEnd
		if data.TimeStart > data.TimeEnd {
			maintain.Flag = true
		} else {
			maintain.Flag = false
		}
		validTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.Valid, time.Local)
		maintain.Valid = &validTime
		monthList := strings.Split(data.Month, ",")
		for _, m := range monthList {
			e, _ := strconv.ParseFloat(m, 64)
			maintain.Month = maintain.Month | int(math.Pow(2, e))
		}
		err = maintain.AddMaintains(data.Hosts)
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, data)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [put]
func (c *MaintainController) UpdateMaintain() {
	var data struct {
		Flag      bool   `json:"flag"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
		Hosts     string `json:"hosts"`
	}
	var ans common.Res
	mid := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(mid, 10, 64)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err == nil {
		var maintain models.Maintains
		maintain.Id = id
		maintain.TimeStart = data.TimeStart
		maintain.TimeEnd = data.TimeEnd
		maintain.DayStart = data.DayStart
		maintain.DayEnd = data.DayEnd
		if data.TimeStart > data.TimeEnd {
			maintain.Flag = true
		} else {
			maintain.Flag = false
		}
		validTime, _ := time.ParseInLocation("2006-01-02 15:04:05", data.Valid, time.Local)
		maintain.Valid = &validTime
		monthList := strings.Split(data.Month, ",")
		for _, m := range monthList {
			e, _ := strconv.ParseFloat(m, 64)
			maintain.Month = maintain.Month | int(math.Pow(2, e))
		}
		//fmt.Println([]byte(data.Hosts))
		err = maintain.UpdateMaintains(data.Hosts)
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, data)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *MaintainController) DeleteMaintain() {
	id := c.Ctx.Input.Param(":id")
	var maintain *models.Maintains
	var ans common.Res
	err := maintain.DeleteMaintains(id)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, id)
	c.Data["json"] = &ans
	c.ServeJSON()
}
