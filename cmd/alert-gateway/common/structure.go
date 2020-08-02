package common

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var ErrHttpRequest = errors.New("create HTTP request failed")
var Maintain map[string]bool
var RuleCount map[[2]int64]int64
var Recover2Send = map[string]map[[2]int64]*Ready2Send{
	AlertMethodLanxin: {},
	//"HOOK":   map[[2]int64]*Ready2Send{},
}

var Lock sync.Mutex
var Rw sync.RWMutex

func UpdateRecovery2Send(ug UserGroup, alert Alert, users []string, alertId int64, alertCount int, hostname string) {

	ruleId, _ := strconv.ParseInt(alert.Annotations.RuleId, 10, 64)

	Lock.Lock()
	defer Lock.Unlock()
	if _, ok := Recover2Send[ug.Method]; !ok {
		Recover2Send[ug.Method] = map[[2]int64]*Ready2Send{{ruleId, ug.Id}: {
			RuleId: ruleId,
			Start:  ug.Id,
			User:   users,
			Alerts: []SingleAlert{{
				Id:       alertId,
				Count:    alertCount,
				Value:    alert.Value,
				Summary:  alert.Annotations.Summary,
				Hostname: hostname,
			}},
		}}
	} else {
		if _, ok := Recover2Send[ug.Method][[2]int64{ruleId, ug.Id}]; !ok {
			Recover2Send[ug.Method][[2]int64{ruleId, ug.Id}] = &Ready2Send{
				RuleId: ruleId,
				Start:  ug.Id,
				User:   users,
				Alerts: []SingleAlert{{
					Id:       alertId,
					Count:    alertCount,
					Value:    alert.Value,
					Summary:  alert.Annotations.Summary,
					Hostname: hostname,
					Labels:   alert.Labels,
				}},
			}
		} else {
			Recover2Send[ug.Method][[2]int64{ruleId, ug.Id}].Alerts = append(Recover2Send[ug.Method][[2]int64{ruleId, ug.Id}].Alerts, SingleAlert{
				Id:       alertId,
				Count:    alertCount,
				Value:    alert.Value,
				Summary:  alert.Annotations.Summary,
				Hostname: hostname,
			})
		}
	}
}

// AuthModel holds information used to authenticate.
type AuthModel struct {
	Username string
	Password string
}

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type BrokenList struct {
	Hosts []struct {
		Hostname string `json:"hostname"`
	} `json:"hosts"`
	Error interface{} `json:"error"`
}

type Msg struct {
	Content string   `json:"content"`
	From    string   `json:"from"`
	Title   string   `json:"title"`
	To      []string `json:"to"`
}

type SingleAlert struct {
	Id       int64             `json:"id"`
	Count    int               `json:"count"`
	Value    float64           `json:"value"`
	Summary  string            `json:"summary"`
	Hostname string            `json:"hostname"`
	Labels   map[string]string `json:"labels"`
}

type Ready2Send struct {
	RuleId int64
	Start  int64
	User   []string
	Alerts []SingleAlert
}

type UserGroup struct {
	Id                    int64
	StartTime             string
	EndTime               string
	Start                 int
	Period                int
	ReversePolishNotation string
	User                  string
	Group                 string
	DutyGroup             string
	Method                string
}

/*
 Check if UserGroup is valid.
*/
func (u UserGroup) IsValid() bool {
	return u.User != "" || u.DutyGroup != "" || u.Group != ""
}

/*
 IsOnDuty return if current UserGroup is on duty or not by StartTime & EndTime.
 If the UserGroup is not on duty, alerts should not be sent to them.
*/
func (u UserGroup) IsOnDuty() bool {
	now := time.Now().Format("15:04")

	return (u.StartTime <= u.EndTime && u.StartTime <= now && u.EndTime >= now) || // 不跨 00:00
		(u.StartTime > u.EndTime && (u.StartTime <= now || now <= u.EndTime)) // // 跨 00:00
}

type Alerts []Alert

type Alert struct {
	ActiveAt    time.Time `json:"active_at"`
	Annotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
		RuleId      string `json:"rule_id"`
	} `json:"annotations"`
	FiredAt    time.Time         `json:"fired_at"`
	Labels     map[string]string `json:"labels"`
	LastSentAt time.Time         `json:"last_sent_at"`
	ResolvedAt time.Time         `json:"resolved_at"`
	State      int               `json:"state"`
	ValidUntil time.Time         `json:"valid_until"`
	Value      float64           `json:"value"`
}

type AlertForShow struct {
	Id              int64             `json:"id,omitempty"`
	RuleId          int64             `json:"rule_id"`
	Labels          map[string]string `json:"labels"`
	Value           float64           `json:"value"`
	Count           int               `json:"count"`
	Status          int8              `json:"status"`
	Summary         string            `json:"summary"`
	Description     string            `json:"description"`
	ConfirmedBy     string            `json:"confirmed_by"`
	FiredAt         *time.Time        `json:"fired_at"`
	ConfirmedAt     *time.Time        `json:"confirmed_at"`
	ConfirmedBefore *time.Time        `json:"confirmed_before"`
	ResolvedAt      *time.Time        `json:"resolved_at"`
}

type Confirm struct {
	Duration int
	User     string
	Ids      []int
}

type ValidUserGroup struct {
	User      string
	Group     string
	DutyGroup string
}

func GenerateJsonHeader() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func HttpPost(url string, params map[string]string, headers map[string]string, body []byte) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, ErrHttpRequest
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{Timeout: 5 * time.Second} //Add the timeout,the reason is that the default client has no timeout set; if the remote server is unresponsive, you're going to have a bad day.
	return client.Do(req)
}

func HttpGet(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, ErrHttpRequest
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{Timeout: 5 * time.Second} //Add the timeout,the reason is that the default client has no timeout set; if the remote server is unresponsive, you're going to have a bad day.
	return client.Do(req)
}
