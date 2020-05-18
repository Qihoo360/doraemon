package initial

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type Record struct {
	Id              int64
	RuleId          int64
	Value           float64
	Count           int
	Summary         string
	Description     string
	Hostname        string
	ConfirmedBefore *time.Time
	FiredAt         *time.Time
	Labels          string
}

type RecoverRecord struct {
	Id       int64
	RuleId   int64
	Value    float64
	Count    int
	Summary  string
	Hostname string
}

func UpdateMaintainlist() {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in UpdateMaintainlist:%v\n%s", e, buf)
		}
	}()
	delta, _ := time.ParseDuration("30s")
	datetime := time.Now().Add(delta)
	now := datetime.Format("15:04")
	maintainIds := []struct {
		Id int64
	}{}
	models.Ormer().Raw("SELECT id FROM maintain WHERE valid>=? AND day_start<=? AND day_end>=? AND (flag=true AND (time_start<=? OR time_end>=?) OR flag=false AND time_start<=? AND time_end>=?) AND month&"+strconv.Itoa(int(math.Pow(2, float64(time.Now().Month()))))+">0", datetime.Format("2006-01-02 15:04:05"), datetime.Day(), datetime.Day(), now, now, now, now).QueryRows(&maintainIds)
	//fmt.Println("abc",datetime.Format("2006-01-02 15:04:05"),datetime.Day(),now,maintainids)
	m := map[string]bool{}
	for _, mid := range maintainIds {
		hosts := []struct {
			Hostname string
		}{}
		models.Ormer().Raw("SELECT hostname FROM host WHERE mid=?", mid.Id).QueryRows(&hosts)
		for _, name := range hosts {
			m[name.Hostname] = true
		}
	}
	res, _ := common.HttpGet(beego.AppConfig.String("BrokenUrl"), nil, map[string]string{"Authorization": "Bearer 8gi6UvoPJgIRcunHBWDHel4fCLQVn9"})
	jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(jsonDataFromHttp))
	brokenList := common.BrokenList{}
	json.Unmarshal(jsonDataFromHttp, &brokenList)
	for _, i := range brokenList.Hosts {
		m[i.Hostname] = true
	}
	common.Rw.Lock()
	common.Maintain = m
	common.Rw.Unlock()
}

func SendAll(method string, from string, param map[string]string, content []common.Ready2Send, now string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in SendAll:%v\n%s", e, buf)
		}
	}()
	if method == "SMS" {
		url := beego.AppConfig.String("SmsUrl")
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			msg = append(msg, "[时间] "+now)
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	} else if method == "LANXIN" {
		url := beego.AppConfig.String("LanxinUrl")
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			for _, j := range i.Alerts {
				duration := ""
				if j.Count >= 60 {
					duration += strconv.FormatInt(int64(j.Count/60), 10) + "h" + strconv.FormatInt(int64(j.Count%60), 10) + "m"
				} else {
					duration = strconv.FormatInt(int64(j.Count), 10) + "m"
				}
				id := strconv.FormatInt(j.Id, 10)
				value := strconv.FormatFloat(j.Value, 'f', 2, 64)
				msg = append(msg, "["+duration+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
			}
			msg = append(msg, "[时间] "+now)
			msg = append(msg, "[确认链接] "+beego.AppConfig.String("WebUrl")+"/alerts_confirm/"+strconv.FormatInt(i.RuleId, 10)+"?start="+strconv.FormatInt(i.Start, 10))
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	} else {
		url := beego.AppConfig.String("CallUrl")
		for _, i := range content {
			msg := []string{"故障" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条 " + i.Alerts[0].Summary + " 详细信息请到蓝信查看"}
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	}

}

func Send2Hook(content []common.Ready2Send, now string, t string, url string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in Send2Hook:%v\n%s", e, buf)
		}
	}()
	if t == "recover" {
		for _, i := range content {
			data, _ := json.Marshal(
				struct {
					Type   string               `json:"type"`
					Time   string               `json:"time"`
					RuleId int64                `json:"rule_id"`
					To     []string             `json:"to"`
					Alerts []common.SingleAlert `json:"alerts"`
				}{
					Type:   t,
					RuleId: i.RuleId,
					Time:   now,
					To:     i.User,
					Alerts: i.Alerts,
				})
			common.HttpPost(url, nil, nil, data)
		}
	} else {
		for _, i := range content {
			data, _ := json.Marshal(
				struct {
					Type        string               `json:"type"`
					Time        string               `json:"time"`
					RuleId      int64                `json:"rule_id"`
					To          []string             `json:"to"`
					ConfirmLink string               `json:"confirm_link"`
					Alerts      []common.SingleAlert `json:"alerts"`
				}{
					Type:        t,
					RuleId:      i.RuleId,
					Time:        now,
					ConfirmLink: beego.AppConfig.String("WebUrl") + "/alerts_confirm/" + strconv.FormatInt(i.RuleId, 10) + "?start=" + strconv.FormatInt(i.Start, 10),
					To:          i.User,
					Alerts:      i.Alerts,
				})
			common.HttpPost(url, nil, nil, data)
		}
	}
}

func Sender(SendClass map[string][]common.Ready2Send, now string) {
	for k, v := range SendClass {
		switch k {
		case "SMS":
			go SendAll("SMS", "mis", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		case "LANXIN":
			go SendAll("LANXIN", "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
			//logs.Alertloger.Info("[%s]%v:", now, v)
		case "CALL":
			go SendAll("CALL", "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		default:
			go Send2Hook(v, now, "alert", k[5:])
		}
	}
}

func RecoverSender(SendClass map[string]map[[2]int64]*common.Ready2Send, now string) {
	lanxin := []common.Ready2Send{}
	for _, v := range SendClass["LANXIN"] {
		lanxin = append(lanxin, *v)
	}
	go SendRecover(beego.AppConfig.String("LanxinUrl"), "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, lanxin, now)
	//logs.Panic.Info("send[%s]:%v", now, lanxin)
	delete(SendClass, "LANXIN")
	for k := range SendClass {
		hook := []common.Ready2Send{}
		for _, u := range SendClass[k] {
			hook = append(hook, *u)
		}
		go Send2Hook(hook, now, "recover", k[5:])
	}
}

func SendRecover(url string, from string, param map[string]string, content []common.Ready2Send, now string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in SendRecover:%v\n%s", e, buf)
		}
	}()
	for _, i := range content {
		msg := []string{"[故障恢复:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
		for _, j := range i.Alerts {
			duration := ""
			if j.Count >= 60 {
				duration += strconv.FormatInt(int64(j.Count/60), 10) + "h" + strconv.FormatInt(int64(j.Count%60), 10) + "m"
			} else {
				duration = strconv.FormatInt(int64(j.Count), 10) + "m"
			}
			id := strconv.FormatInt(j.Id, 10)
			value := strconv.FormatFloat(j.Value, 'f', 2, 64)
			msg = append(msg, "["+duration+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
		}
		msg = append(msg, "[时间] "+now)
		data, _ := json.Marshal(common.Msg{
			Content: strings.Join(msg, "\n"),
			From:    from,
			Title:   "Alerts",
			To:      i.User})
		common.HttpPost(url, param, nil, data)
	}
}

func Filter(alerts map[int64][]Record, maxCount map[int64]int) map[string][]common.Ready2Send {
	SendClass := map[string][]common.Ready2Send{
		"SMS":    []common.Ready2Send{},
		"LANXIN": []common.Ready2Send{},
		"CALL":   []common.Ready2Send{},
		//"HOOK":   []common.Ready2Send{},
	}
	Cache := map[int64][]common.UserGroup{}
	NewRuleCount := map[[2]int64]int64{}
	datetime := time.Now()
	now := datetime.Format("15:04")
	for key := range alerts {
		var usergroupList []common.UserGroup
		var planId struct {
			PlanId  int64
			Summary string
		}
		AlertsMap := map[int][]common.SingleAlert{}
		models.Ormer().Raw("SELECT plan_id,summary FROM rule WHERE id=?", key).QueryRow(&planId)
		if _, ok := Cache[planId.PlanId]; !ok {
			models.Ormer().Raw("SELECT id,start_time,end_time,start,period,reverse_polish_notation,user,`group`,duty_group,method FROM plan_receiver WHERE plan_id=?", planId.PlanId).QueryRows(&usergroupList)
			Cache[planId.PlanId] = usergroupList
		}
		for _, element := range Cache[planId.PlanId] {
			if element.User != "" || element.DutyGroup != "" || element.Group != "" {
				if (element.StartTime <= element.EndTime && element.StartTime <= now && element.EndTime >= now) || (element.StartTime > element.EndTime && (element.StartTime <= now || now <= element.EndTime)) {
					if maxCount[key] >= element.Start {
						if _, ok := common.RuleCount[[2]int64{key, int64(element.Start)}]; !ok {
							NewRuleCount[[2]int64{key, int64(element.Start)}] = -1
						} else {
							NewRuleCount[[2]int64{key, int64(element.Start)}] = common.RuleCount[[2]int64{key, int64(element.Start)}]
						}
						NewRuleCount[[2]int64{key, int64(element.Start)}] += 1

						if NewRuleCount[[2]int64{key, int64(element.Start)}]%int64(element.Period) == 0 {
							if _, ok := AlertsMap[element.Start]; !ok {
								AlertsMap[element.Start] = []common.SingleAlert{}
							} else {
								if len(AlertsMap[element.Start]) > 0 {
									if element.ReversePolishNotation == "" {
										SendClass[element.Method] = append(SendClass[element.Method], common.Ready2Send{
											RuleId: key,
											Start:  element.Id,
											User: models.SendAlertsFor(&common.ValidUserGroup{
												User:      element.User,
												Group:     element.Group,
												DutyGroup: element.DutyGroup,
											}),
											Alerts: AlertsMap[element.Start],
										})
									} else {
										filteredAlerts := []common.SingleAlert{}
										for _, alert := range AlertsMap[element.Start] {
											if common.CalculateReversePolishNotation(alert.Labels, element.ReversePolishNotation) {
												filteredAlerts = append(filteredAlerts, alert)
											}
										}
										if len(filteredAlerts) > 0 {
											SendClass[element.Method] = append(SendClass[element.Method], common.Ready2Send{
												RuleId: key,
												Start:  element.Id,
												User: models.SendAlertsFor(&common.ValidUserGroup{
													User:      element.User,
													Group:     element.Group,
													DutyGroup: element.DutyGroup,
												}),
												Alerts: filteredAlerts,
											})
										}
									}
								}
								continue
							}
							for _, alert := range alerts[key] {
								if alert.Count >= element.Start {
									if _, ok := common.Maintain[alert.Hostname]; !ok {
										labelMap := map[string]string{}
										if alert.Labels != "" {
											for _, j := range strings.Split(alert.Labels, "\v") {
												kv := strings.Split(j, "\a")
												labelMap[kv[0]] = kv[1]
											}
										}
										AlertsMap[element.Start] = append(AlertsMap[element.Start], common.SingleAlert{
											Id:       alert.Id,
											Count:    alert.Count,
											Value:    alert.Value,
											Summary:  planId.Summary,
											Hostname: alert.Hostname,
											Labels:   labelMap,
										})
									}
								}
							}
							if len(AlertsMap[element.Start]) > 0 {
								if element.ReversePolishNotation == "" {
									SendClass[element.Method] = append(SendClass[element.Method], common.Ready2Send{
										RuleId: key,
										Start:  element.Id,
										User: models.SendAlertsFor(&common.ValidUserGroup{
											User:      element.User,
											Group:     element.Group,
											DutyGroup: element.DutyGroup,
										}),
										Alerts: AlertsMap[element.Start],
									})
								} else {
									filteredAlerts := []common.SingleAlert{}
									for _, alert := range AlertsMap[element.Start] {
										if common.CalculateReversePolishNotation(alert.Labels, element.ReversePolishNotation) {
											filteredAlerts = append(filteredAlerts, alert)
										}
									}
									if len(filteredAlerts) > 0 {
										SendClass[element.Method] = append(SendClass[element.Method], common.Ready2Send{
											RuleId: key,
											Start:  element.Id,
											User: models.SendAlertsFor(&common.ValidUserGroup{
												User:      element.User,
												Group:     element.Group,
												DutyGroup: element.DutyGroup,
											}),
											Alerts: filteredAlerts,
										})
									}
								}
							}
						}
					}
				}
			}
		}
	}
	common.RuleCount = NewRuleCount
	//logs.Alertloger.Debug("RuleCount: %v", common.RuleCount)
	return SendClass
}

func init() {
	go func() {
		for {
			current := time.Now()
			time.Sleep(time.Duration(90-current.Second()) * time.Second)
			UpdateMaintainlist()
		}
	}()
	go func() {
		for {
			//time.Sleep(time.Second)
			current := time.Now()
			time.Sleep(time.Duration(60-current.Second()) * time.Second)
			now := time.Now().Format("2006-01-02 15:04:05")
			go func() {
				defer func() {
					if e := recover(); e != nil {
						buf := make([]byte, 16384)
						buf = buf[:runtime.Stack(buf, false)]
						logs.Panic.Error("Panic in timer:%v\n%s", e, buf)
					}
				}()
				var info []Record
				//_, err := o.QueryTable(models.Alerts{}).Limit(-1).Filter("status", 2).Update(orm.Params{"count": orm.ColValue(orm.ColAdd, 1)})
				models.Ormer().Raw("UPDATE alert SET status=2 WHERE status=1 AND confirmed_before<?", now).Exec()
				o := orm.NewOrm()
				o.Begin()
				o.Raw("UPDATE alert SET count=count+1 WHERE status!=0").Exec()
				o.Raw("SELECT id,rule_id,value,count,summary,description,hostname,confirmed_before,fired_at,labels FROM alert WHERE status = ?", 2).QueryRows(&info)
				//filter alerts...
				aggregation := map[int64][]Record{}
				maxCount := map[int64]int{}
				for _, i := range info {
					aggregation[i.RuleId] = append(aggregation[i.RuleId], i)
					if _, ok := maxCount[i.RuleId]; !ok {
						maxCount[i.RuleId] = i.Count
					} else {
						if i.Count > maxCount[i.RuleId] {
							maxCount[i.RuleId] = i.Count
						}
					}
				}
				common.Rw.RLock()
				ready2send := Filter(aggregation, maxCount)
				common.Rw.RUnlock()
				o.Commit()
				logs.Alertloger.Info("Alerts to send:%v", ready2send)
				Sender(ready2send, now)
				common.Lock.Lock()
				recover2send := common.Recover2Send
				common.Recover2Send = map[string]map[[2]int64]*common.Ready2Send{
					"LANXIN": map[[2]int64]*common.Ready2Send{},
					//"HOOK":   map[[2]int64]*common.Ready2Send{},
				}
				common.Lock.Unlock()
				logs.Alertloger.Info("Recoveries to send:%v", recover2send)
				RecoverSender(recover2send, now)
			}()
		}
	}()
}
