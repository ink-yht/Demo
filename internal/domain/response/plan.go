package response

// Plan 计划信息响应
type Plan struct {
	PlanNo     string      `xml:"planNo"`   // 计划号
	PlanName   string      `xml:"planName"` // 计划名称
	Insurances []Insurance `xml:"insurances>insurance"`
}

// Insurance 被保险人信息响应
type Insurance struct {
	InsuranceName    string `xml:"insuranceName"`              // 被保险人姓名
	PersonCode       string `xml:"personCode"`                 // 被保险人id，用来批改时使用
	CertiNo          string `xml:"certiNo"`                    // 证件号
	CertiType        string `xml:"certiType"`                  // 证件号类型
	InsurancePremium string `xml:"insurancePremium,omitempty"` // 标的保费变化 (否)
}
