package models

import (
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
)

type Alerts struct {
	Id              int64      `orm:"column(id);auto" json:"id,omitempty"`
	Rule            *Rules     `orm:"rel(fk)" json:"rule_id"`
	Labels          string     `orm:"column(labels);size(4095)" json:"labels"`
	Value           float64    `orm:"column(value)" json:"value"`
	Count           int        `json:"count"`
	Status          int8       `orm:"index" json:"status"`
	Summary         string     `orm:"column(summary);size(1023)" json:"summary"`
	Description     string     `orm:"column(description);size(1023)" json:"description"`
	Hostname        string     `orm:"column(hostname);size(255)" json:"hostname"`
	ConfirmedBy     string     `orm:"column(confirmed_by);size(1023)" json:"confirmed_by"`
	FiredAt         *time.Time `orm:"type(datetime)" json:"fired_at"`
	ConfirmedAt     *time.Time `orm:"null" json:"confirmed_at"`
	ConfirmedBefore *time.Time `orm:"null" json:"confirmed_before"`
	ResolvedAt      *time.Time `orm:"null" json:"resolved_at"`
}

type OneAlert struct {
	ID              int64      `json:"id"`
	RuleID          int64      `json:"rule_id"`
	Value           float64    `json:"value"`
	Status          int8       `json:"status"`
	Count           int        `json:"count"`
	Summary         string     `json:"summary"`
	Description     string     `json:"description"`
	ConfirmedBy     string     `json:"confirmed_by"`
	FiredAt         *time.Time `json:"fired_at"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	ConfirmedBefore *time.Time `json:"confirmed_before"`
	ResolvedAt      *time.Time `json:"resolved_at"`
}

type ShowAlerts struct {
	Alerts []common.AlertForShow `json:"alerts"`
	Total  int64                 `json:"total"`
}

func (*Alerts) TableName() string {
	return "alert"
}

//func (u *Alerts) TableUnique() [][]string {
//	return [][]string{
//		[]string{"Rule", "Labels", "FiredAt"},
//	}
//}

type record struct {
	Id              int64
	RuleId          int64
	Labels          string
	Value           float64
	Count           int
	Status          int8
	Summary         string
	Description     string
	ConfirmedBy     string
	FiredAt         *time.Time
	ConfirmedAt     *time.Time
	ConfirmedBefore *time.Time
	ResolvedAt      *time.Time
}

func (r record) toOneAlert() OneAlert {
	return OneAlert{
		ID:              r.Id,
		RuleID:          r.RuleId,
		Value:           r.Value,
		Status:          r.Status,
		Count:           r.Count,
		Summary:         r.Summary,
		Description:     r.Description,
		ConfirmedBy:     r.ConfirmedBy,
		FiredAt:         r.FiredAt,
		ConfirmedAt:     r.ConfirmedAt,
		ConfirmedBefore: r.ConfirmedBefore,
		ResolvedAt:      r.ResolvedAt,
	}
}

func (r record) getLabelMap() map[string]string {
	label := map[string]string{}
	if r.Labels != "" {
		for _, e := range strings.Split(r.Labels, "\v") {
			kv := strings.Split(e, "\a")
			label[kv[0]] = kv[1]
		}
	}
	return label
}

func (r record) toAlertForShow() common.AlertForShow {

	return common.AlertForShow{
		Id:              r.Id,
		RuleId:          r.RuleId,
		Labels:          r.getLabelMap(),
		Value:           r.Value,
		Count:           r.Count,
		Status:          r.Status,
		Summary:         r.Summary,
		Description:     r.Description,
		ConfirmedBy:     r.ConfirmedBy,
		FiredAt:         r.FiredAt,
		ConfirmedAt:     r.ConfirmedAt,
		ConfirmedBefore: r.ConfirmedBefore,
		ResolvedAt:      r.ResolvedAt,
	}
}

func (u *Alerts) ClassifyAlerts() map[string]map[string][]OneAlert {
	var records []record

	Ormer().
		Raw("SELECT id,rule_id,labels,value,status,count,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE status=2 AND count!=-1").
		QueryRows(&records)
	res := map[string]map[string][]OneAlert{}
	for _, i := range records {
		if i.Labels != "" {
			for _, j := range strings.Split(i.Labels, "\v") {
				kv := strings.Split(j, "\a")
				if _, ok := res[kv[0]]; ok {
					res[kv[0]][kv[1]] = append(res[kv[0]][kv[1]], i.toOneAlert())
				} else {
					res[kv[0]] = map[string][]OneAlert{}
					res[kv[0]][kv[1]] = append(res[kv[0]][kv[1]], i.toOneAlert())
				}
			}
		} else {
			if _, ok := res["no label"]; ok {
				res["no label"]["no label"] = append(res["no label"]["no label"], i.toOneAlert())
			} else {
				res["no label"] = map[string][]OneAlert{}
				res["no label"]["no label"] = append(res["no label"]["no label"], i.toOneAlert())
			}
		}
	}
	return res
}

func (u *Alerts) GetAlerts(pageNo int64, pageSize int64, timeStart string, timeEnd string, status string, summary string) ShowAlerts {
	var showAlerts ShowAlerts
	showAlerts.Alerts = []common.AlertForShow{}
	var records []record

	if summary != "" {
		if status != "" {
			if timeStart != "" {
				if timeEnd != "" {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND fired_at<=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND fired_at<=? AND status=? AND summary LIKE ?", timeStart, timeEnd, status, "%"+summary+"%").QueryRow(&showAlerts.Total)
				} else {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND status=? AND summary LIKE ?", timeStart, status, "%"+summary+"%").QueryRow(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at<=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeEnd, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at<=? AND status=? AND summary LIKE ?", timeEnd, status, "%"+summary+"%").QueryRow(&showAlerts.Total)
			} else {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE status=? AND summary LIKE ?", status, "%"+summary+"%").QueryRow(&showAlerts.Total)
			}
		} else {
			if timeStart != "" {
				if timeEnd != "" {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND fired_at<=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND fired_at<=? AND summary LIKE ?", timeStart, timeEnd, "%"+summary+"%").QueryRow(&showAlerts.Total)
				} else {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND summary LIKE ?", timeStart, "%"+summary+"%").QueryRow(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at<=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeEnd, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at<=? AND summary LIKE ?", timeEnd, "%"+summary+"%").QueryRow(&showAlerts.Total)
			} else {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE summary LIKE ? ORDER BY id DESC LIMIT ?,?", "%"+summary+"%", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE summary LIKE ?", "%"+summary+"%").QueryRow(&showAlerts.Total)
			}
		}
	} else {
		if status != "" {
			if timeStart != "" {
				if timeEnd != "" {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND fired_at<=? AND status=? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, status, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND fired_at<=? AND status=?", timeStart, timeEnd, status).QueryRow(&showAlerts.Total)
				} else {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND status=? ORDER BY id DESC LIMIT ?,?", timeStart, status, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND status=?", timeStart, status).QueryRow(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at<=? AND status=? ORDER BY id DESC LIMIT ?,?", timeEnd, status, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at<=? AND status=?", timeEnd, status).QueryRow(&showAlerts.Total)
			} else {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE status=? ORDER BY id DESC LIMIT ?,?", status, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE status=?", status).QueryRow(&showAlerts.Total)
			}
		} else {
			if timeStart != "" {
				if timeEnd != "" {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? AND fired_at<=? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=? AND fired_at<=?", timeStart, timeEnd).QueryRow(&showAlerts.Total)
				} else {
					_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at>=? ORDER BY id DESC LIMIT ?,?", timeStart, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
					_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at>=?", timeStart).QueryRow(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE fired_at<=? ORDER BY id DESC LIMIT ?,?", timeEnd, (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert WHERE fired_at<=?", timeEnd).QueryRow(&showAlerts.Total)
			} else {
				_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert ORDER BY id DESC LIMIT ?,?", (pageNo-1)*pageSize, pageSize).QueryRows(&records)
				_ = Ormer().Raw("SELECT count(*) FROM alert").QueryRow(&showAlerts.Total)
			}
		}
	}

	for _, i := range records {
		showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
	}
	//showalerts.Total, _ = Ormer().QueryTable(Alerts{}).Limit(-1).Count()
	return showAlerts
}

func (u *Alerts) ShowAlerts(ruleId string, start string, pageNo int64, pageSize int64) ShowAlerts {
	var showAlerts ShowAlerts
	showAlerts.Alerts = []common.AlertForShow{}
	var records []record
	strategy := struct {
		ReversePolishNotation string
		Start                 int
	}{}
	if start != "" {
		_ = Ormer().Raw("SELECT start,reverse_polish_notation FROM plan_receiver WHERE id=?", start).QueryRow(&strategy)
	}
	_, _ = Ormer().Raw("SELECT id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at FROM alert WHERE count>=? AND rule_id=? AND status!=0 ORDER BY status DESC,id DESC", strategy.Start, ruleId).QueryRows(&records)
	for _, i := range records {
		label := i.getLabelMap()
		if strategy.ReversePolishNotation != "" {
			if common.CalculateReversePolishNotation(label, strategy.ReversePolishNotation) {
				showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
			}
		} else {
			showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
		}
	}
	showAlerts.Total = int64(len(showAlerts.Alerts))
	if showAlerts.Total == 0 {
		return showAlerts
	} else if showAlerts.Total < pageNo*pageSize {
		showAlerts.Alerts = showAlerts.Alerts[(pageNo-1)*pageSize:]
		return showAlerts
	} else {
		showAlerts.Alerts = showAlerts.Alerts[(pageNo-1)*pageSize : pageNo*pageSize]
		return showAlerts
	}
}

func (u *Alerts) ConfirmAll(confirmList *common.Confirm) error {
	now := time.Now()
	var err error
	for _, id := range confirmList.Ids {
		var rs struct {
			Status uint8
		}
		o := orm.NewOrm()
		o.Begin()
		err = o.Raw("SELECT status,rule_id FROM alert WHERE id=? LOCK IN SHARE MODE", id).QueryRow(&rs)
		if err != nil {
			o.Rollback()
			return errors.Wrap(err, "database query error")
		} else {
			const AlertStatusOn = 2
			if rs.Status == AlertStatusOn {
				_, err = o.Raw("UPDATE alert SET status=1,confirmed_at=?,confirmed_by=?,confirmed_before=? WHERE id=?", now.Format("2006-01-02 15:04:05"), confirmList.User, now.Add(time.Duration(confirmList.Duration)*time.Minute).Format("2006-01-02 15:04:05"), id).Exec()
				if err != nil {
					o.Rollback()
					return errors.Wrap(err, "database update error")
				}
			}
		}
		o.Commit()
	}
	return errors.Wrap(err, "database update error")
}

type alertForQuery struct {
	*common.Alert
	label    string
	hostname string
	ruleId   int64
	firedAt  time.Time
}

/*
 set value for fields in alertForQuery
*/
func (a *alertForQuery) setFields() {
	var orderKey []string
	var labels []string

	// set ruleId
	a.ruleId, _ = strconv.ParseInt(a.Annotations.RuleId, 10, 64)
	for key := range a.Labels {
		orderKey = append(orderKey, key)
	}
	sort.Strings(orderKey)
	for _, i := range orderKey {
		labels = append(labels, i+"\a"+a.Labels[i])
	}
	// set label
	a.label = strings.Join(labels, "\v")
	// set firedAt
	a.firedAt = a.FiredAt.Truncate(time.Second)
	// set hostname
	a.setHostname()
}

/*
 set hostname by instance label on data
*/
func (a *alertForQuery) setHostname() {
	h := ""
	if _, ok := a.Labels["instance"]; ok {
		h = a.Labels["instance"]
		boundary := strings.LastIndex(h, ":")
		if boundary != -1 {
			h = h[:boundary]
		}
	}
	a.hostname = h
}

func (u *Alerts) AlertsHandler(alert *common.Alerts) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in AlertsHandler:%v\n%s", e, buf)
		}
	}()
	//rlist := []int64{}
	Cache := map[int64][]common.UserGroup{}
	now := time.Now().Format("15:04:05")
	todayZero, _ := time.ParseInLocation("2006-01-02", "2019-01-01 15:22:22", time.Local)
	for _, elemt := range *alert {

		var queryres []struct {
			Id     int64
			Status uint8
		}

		a := &alertForQuery{Alert: &elemt}
		a.setFields()

		_, err := Ormer().Raw("SELECT id,status FROM alert WHERE rule_id =? AND labels=? AND fired_at=?", a.ruleId, a.label, a.firedAt).QueryRows(&queryres)
		if err == nil {
			if len(queryres) > 0 {
				if queryres[0].Status != 0 {
					const AlertStatusOff = 0
					if elemt.State == AlertStatusOff {
						//rlist = append(rlist, queryres[0].Id)
						recoverInfo := struct {
							Id       int64
							Count    int
							Hostname string
						}{}
						o := orm.NewOrm()
						o.Begin()
						err = o.Raw("SELECT id,count,hostname From alert WHERE rule_id =? AND labels=? AND fired_at=? FOR UPDATE", a.ruleId, a.label, a.firedAt).QueryRow(&recoverInfo)
						if err == nil {
							if recoverInfo.Id != 0 {
								_, err = o.Raw("UPDATE alert SET status=?,summary=?,description=?,value=?,resolved_at=? WHERE id=?", elemt.State, elemt.Annotations.Summary, elemt.Annotations.Description, elemt.Value, elemt.ResolvedAt, recoverInfo.Id).Exec()
								if err == nil {
									//logs.Alertloger.Info("AlertRecovered:%s", elemt)
									common.Rw.RLock()
									if _, ok := common.Maintain[a.hostname]; !ok {
										var userGroupList []common.UserGroup
										var planId struct {
											PlanId  int64
											Summary string
										}
										Ormer().Raw("SELECT plan_id,summary FROM rule WHERE id=?", a.ruleId).QueryRow(&planId)
										if _, ok := Cache[planId.PlanId]; !ok {
											Ormer().Raw("SELECT id,start_time,end_time,start,period,reverse_polish_notation,user,`group`,duty_group,method FROM plan_receiver WHERE plan_id=? AND (method='LANXIN' OR method LIKE 'HOOK %')", planId.PlanId).QueryRows(&userGroupList)
											Cache[planId.PlanId] = userGroupList
										}
										for _, element := range Cache[planId.PlanId] {
											if element.IsValid() && element.IsOnDuty() {
												if recoverInfo.Count >= element.Start {
													sendFlag := false
													if recoverInfo.Count-element.Start >= element.Period {
														sendFlag = true
													} else {
														if _, ok := common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]; ok {
															logs.Panic.Debug("[%s] id:%d,rulecount:%d,count:%d,start:%d,period:%d", now, recoverInfo.Id, common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}], recoverInfo.Count, element.Start, element.Period)
															if common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}] >= int64(recoverInfo.Count-element.Start) {
																logs.Panic.Debug("[%s] id:%d %d,%s", now, recoverInfo.Id, (common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-int64(recoverInfo.Count)+int64(element.Start))%int64(element.Period), common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-((common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-int64(recoverInfo.Count)+int64(element.Start))/int64(element.Period))*int64(element.Period) >= int64(element.Period))
																if (common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-int64(recoverInfo.Count)+int64(element.Start))%int64(element.Period) == 0 || common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-((common.RuleCount[[2]int64{a.ruleId, int64(element.Start)}]-int64(recoverInfo.Count)+int64(element.Start))/int64(element.Period))*int64(element.Period) >= int64(element.Period) {
																	sendFlag = true
																}
															}
														}
													}
													if sendFlag {
														if element.ReversePolishNotation == "" || common.CalculateReversePolishNotation(elemt.Labels, element.ReversePolishNotation) {
															users := SendAlertsFor(&common.ValidUserGroup{
																User:      element.User,
																Group:     element.Group,
																DutyGroup: element.DutyGroup,
															})
															common.UpdateRecovery2Send(element, elemt, users, recoverInfo.Id, recoverInfo.Count, recoverInfo.Hostname)
														}
													}
												}
											}
										}
									}
									common.Rw.RUnlock()
									o.Commit()
								} else {
									o.Rollback()
									//logs.Alertloger.Error("models.AlertsHandler alertsrecover sql error:%s", err.Error())
								}
							}
							o.Commit()
						} else {
							o.Rollback()
							Ormer().Raw("UPDATE alert SET status=?,summary=?,description=?,value=?,resolved_at=? WHERE id=?", elemt.State, elemt.Annotations.Summary, elemt.Annotations.Description, elemt.Value, elemt.ResolvedAt, recoverInfo.Id).Exec() //if exceed the max waiting time for getting the lock
						}
						//send the recover message
					} else {
						Ormer().Raw("UPDATE alert SET summary=?,description=?,value=? WHERE rule_id =? AND labels=? AND fired_at=?", elemt.Annotations.Summary, elemt.Annotations.Description, elemt.Value, a.ruleId, a.label, a.firedAt).Exec()
					}
				} else {
					continue
				}
			} else {
				var alert Alerts
				alert.Id = 0 //reset the "Id" to 0,which is very important:after a record is inserted,the value of "Id" will not be 0,but the auto primary key of the record
				alert.Rule = &Rules{Id: a.ruleId}
				alert.Labels = a.label
				alert.FiredAt = &a.firedAt
				alert.Description = elemt.Annotations.Description
				alert.Summary = elemt.Annotations.Summary
				alert.Count = -1
				alert.Value = elemt.Value
				alert.Status = int8(elemt.State)
				alert.Hostname = a.hostname
				alert.ConfirmedAt = &todayZero
				alert.ConfirmedBefore = &todayZero
				alert.ResolvedAt = &todayZero
				_, err := Ormer().Insert(&alert)
				if err != nil {
					logs.Error("Insert alter failed:%s", err)
				}
			}
		}
	}
	//logs.Panic.Debug("[%s] recoverid: %v", now, rlist)
}
