package models

type Hosts struct {
	Id       int64  `orm:"auto" json:"id,omitempty"`
	Mid      int64  `json:"mid"`
	Hostname string `orm:"size(255)" json:"hostname"`
}

func (*Hosts) TableName() string {
	return "host"
}

func (u *Hosts) TableUnique() [][]string {
	return [][]string{
		[]string{"Mid", "Hostname"},
	}
}

func (u *Hosts) GetHosts(mid string) []string {
	hosts := []struct {
		Hostname string
	}{}
	Ormer().Raw("SELECT hostname FROM host WHERE mid=?", mid).QueryRows(&hosts)
	res := []string{}
	for _, i := range hosts {
		res = append(res, i.Hostname)
	}
	return res
}
