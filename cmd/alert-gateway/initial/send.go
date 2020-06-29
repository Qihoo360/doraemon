package initial

import (
	"encoding/json"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/astaxie/beego"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
 send alert if rule is triggered.
*/
func Sender(SendClass map[string][]common.Ready2Send, now string) {
	for k, v := range SendClass {
		switch k {
		case common.AlertMethodSms:
			go SendAll(k, "mis", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		case common.AlertMethodLanxin:
			go SendAll(k, "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
			//logs.Alertloger.Info("[%s]%v:", now, v)
		case common.AlertMethodCall:
			go SendAll(k, "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		default:
			go Send2Hook(v, now, "alert", k[5:])
		}
	}
}

/*
 send recovery message if alert recovered.
*/
func RecoverSender(SendClass map[string]map[[2]int64]*common.Ready2Send, now string) {
	lanxin := []common.Ready2Send{}
	for _, v := range SendClass[common.AlertMethodLanxin] {
		lanxin = append(lanxin, *v)
	}
	go SendRecover(beego.AppConfig.String("LanxinUrl"), "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, lanxin, now)
	//logs.Panic.Info("send[%s]:%v", now, lanxin)
	delete(SendClass, common.AlertMethodLanxin)
	for k := range SendClass {
		hook := []common.Ready2Send{}
		for _, u := range SendClass[k] {
			hook = append(hook, *u)
		}
		go Send2Hook(hook, now, "recover", k[5:])
	}
}

func SendAll(method string, from string, param map[string]string, content []common.Ready2Send, sendTime string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in SendAll:%v\n%s", e, buf)
		}
	}()
	if method == common.AlertMethodSms {
		url := beego.AppConfig.String("SmsUrl")
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			msg = append(msg, "[时间] "+sendTime)
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	} else if method == common.AlertMethodLanxin {
		url := beego.AppConfig.String("LanxinUrl")
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			for _, j := range i.Alerts {
				duration, _ := time.ParseDuration(strconv.FormatInt(int64(j.Count), 10) + "m")

				id := strconv.FormatInt(j.Id, 10)
				value := strconv.FormatFloat(j.Value, 'f', 2, 64)
				msg = append(msg, "["+duration.String()+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
			}
			msg = append(msg, "[时间] "+sendTime)
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

type hookRequest struct {
	Type        string               `json:"type"`
	Time        string               `json:"time"`
	RuleId      int64                `json:"rule_id"`
	To          []string             `json:"to"`
	Alerts      []common.SingleAlert `json:"alerts"`
	ConfirmLink string               `json:"confirm_link,omitempty"`
}

func Send2Hook(content []common.Ready2Send, sendTime string, t string, url string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in Send2Hook:%v\n%s", e, buf)
		}
	}()
	if t == "recover" {
		for _, i := range content {
			data, _ := json.Marshal(hookRequest{
				Type:   t,
				RuleId: i.RuleId,
				Time:   sendTime,
				To:     i.User,
				Alerts: i.Alerts,
			})
			common.HttpPost(url, nil, common.GenerateJsonHeader(), data)
		}
	} else {
		for _, i := range content {
			data, _ := json.Marshal(hookRequest{
				Type:        t,
				RuleId:      i.RuleId,
				Time:        sendTime,
				ConfirmLink: beego.AppConfig.String("WebUrl") + "/alerts_confirm/" + strconv.FormatInt(i.RuleId, 10) + "?start=" + strconv.FormatInt(i.Start, 10),
				To:          i.User,
				Alerts:      i.Alerts,
			})
			common.HttpPost(url, nil, common.GenerateJsonHeader(), data)
		}
	}
}

func SendRecover(url string, from string, param map[string]string, content []common.Ready2Send, sendTime string) {
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
			duration, _ := time.ParseDuration(strconv.FormatInt(int64(j.Count), 10) + "m")

			id := strconv.FormatInt(j.Id, 10)
			value := strconv.FormatFloat(j.Value, 'f', 2, 64)
			msg = append(msg, "["+duration.String()+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
		}
		msg = append(msg, "[时间] "+sendTime)
		data, _ := json.Marshal(common.Msg{
			Content: strings.Join(msg, "\n"),
			From:    from,
			Title:   "Alerts",
			To:      i.User})
		common.HttpPost(url, param, nil, data)
	}
}
