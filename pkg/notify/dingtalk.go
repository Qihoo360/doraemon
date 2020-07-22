package notify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/astaxie/beego"
	uri "net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// DingTalkMessage the dingTalk request body model
type DingTalkMsg struct {
	Type     string              `json:"msgtype"`
	Markdown DingTalkMarkdownMsg `json:"markdown"`
}

// DingTalkMarkdownMsg the dingTalk markdown request body model
type DingTalkMarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Send2DingTalk send alert or recover to DingTalk IM
func Send2DingTalk(content []common.Ready2Send, isRecover bool, sendTime string, url string, secret string) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in Send2DingTalk: %v\n%s", e, buf)
		}
	}()
	for _, c := range content {
		mk := dingTalkMsg2String(c, sendTime, isRecover)
		var title string
		if isRecover {
			title = fmt.Sprintf("有 %d 条故障恢复", len(c.Alerts))
		} else {
			title = fmt.Sprintf("有 %d 条故障告警", len(c.Alerts))
		}
		post2DingTalk(url, secret, &DingTalkMsg{
			Type: "markdown",
			Markdown: DingTalkMarkdownMsg{
				Title: title,
				Text:  mk,
			},
		})

	}
}

// post2DingTalk
func post2DingTalk(url string, secret string, msg *DingTalkMsg) {
	t := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	addr, err := uri.Parse(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Parse the url: %s err: %v\n", url, err))
		return
	}
	qs := addr.Query()
	qs.Set("timestamp", t)
	qs.Set("sign", signRequestData(t, secret))
	addr.RawQuery = qs.Encode()

	message, err := json.Marshal(msg)
	if err != nil {
		logs.Error(fmt.Sprintf("Marshal the dingTalk msg: %v err: %v\n", msg, err))
		return
	}
	// ignore DingTalk response for now
	_, _ = common.HttpPost(addr.String(), nil, common.GenerateJsonHeader(), message)
}

// dingTalkMsg2String convert Ready2Send to dingTalk markdown string
func dingTalkMsg2String(content common.Ready2Send, sendTime string, isRecover bool) string {
	var msg []string
	if !isRecover {
		msg = []string{fmt.Sprintf("<font color=#e50303 size=4 face=\"楷体\">告警触发: %s</font>", content.Alerts[0].Summary)}
	} else {
		msg = []string{fmt.Sprintf("<font color=#01c101 size=4 face=\"楷体\">告警恢复: %s</font>", content.Alerts[0].Summary)}
	}
	msg = append(msg, "**基本信息**")
	msg = append(msg, fmt.Sprintf("> - 告警时间: %s", sendTime))

	if len(content.User) > 0 {
		msg = append(msg, "**告警人员**")
		for _, u := range content.User {
			msg = append(msg, fmt.Sprintf("> - %s", u))
		}
	}

	msg = append(msg, "**详细指标**")
	for _, i := range content.Alerts {
		msg = append(msg, fmt.Sprintf("> - 告警ID: %d", i.Id))
		msg = append(msg, fmt.Sprintf("> - 异常分钟: %s 分钟", strconv.FormatInt(int64(i.Count), 10)))
		msg = append(msg, fmt.Sprintf("> - 当前数值: %v", i.Value))
		if len(i.Hostname) > 0 {
			msg = append(msg, fmt.Sprintf("> - 故障主机: %s", i.Hostname))

		}
		msg = append(msg, "---")
	}
	/*if len(content.Alerts[0].Description) > 0 {
		msg = append(msg, "**具体描述**")
		msg = append(msg, fmt.Sprintf("> **消息**: %s", content.Alerts[0].Description))
		msg = append(msg, "---")
	}*/
	if !isRecover {
		msg = append(msg, fmt.Sprintf("[点击确认告警](%s/alerts_confirm/%d?start=%d)", beego.AppConfig.String("WebUrl"), content.RuleId, content.Start))
	}
	return strings.Join(msg, "\n\n")

}

// signRequestData DingTalk bot request params sign
func signRequestData(t string, secret string) string {
	string2sign := fmt.Sprintf("%s\n%s", t, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(string2sign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
