package modules

import (
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// M is map
type M map[string]interface{}

// S is slice
type S []interface{}

// Prom ...
type Prom struct {
	ID  int64
	URL string
}

// Rule ...
type Rule struct {
	ID          int64             `json:"id"`
	PromID      int64             `json:"prom_id"`
	Expr        string            `json:"expr"`
	Op          string            `json:"op"`
	Value       string            `json:"value"`
	For         string            `json:"for"`
	Labels      map[string]string `json:"labels"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
}

// Rules ...
type Rules []Rule

// PromRules ...
type PromRules struct {
	Prom  Prom
	Rules Rules
}

// RulesResp ...
type RulesResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Rules  `json:"data"`
}

// PromsResp ...
type PromsResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Prom `json:"data"`
}

// Content get prom rules
func (r Rules) Content() ([]byte, error) {
	rules := S{}
	for _, i := range r {
		rules = append(rules, M{
			"alert":  strconv.FormatInt(i.ID, 10),
			"expr":   strings.Join([]string{i.Expr, i.Op, i.Value}, " "),
			"for":    i.For,
			"labels": i.Labels,
			"annotations": M{
				"rule_id":     strconv.FormatInt(i.ID, 10),
				"prom_id":     strconv.FormatInt(i.PromID, 10),
				"summary":     i.Summary,
				"description": i.Description,
			},
		})
	}
	result := M{
		"groups": S{
			M{
				"name":  "ruleengine",
				"rules": rules,
			},
		},
	}

	return yaml.Marshal(result)
}

// PromRules cut prom rules
func (r Rules) PromRules() []PromRules {
	tmp := map[int64]Rules{}

	for _, rule := range r {
		if v, ok := tmp[rule.PromID]; ok {
			tmp[rule.PromID] = append(v, rule)
		} else {
			tmp[rule.PromID] = Rules{rule}
		}
	}

	data := []PromRules{}
	for id, rules := range tmp {
		data = append(data, PromRules{
			Prom:  Prom{ID: id},
			Rules: rules,
		})
	}

	return data
}
