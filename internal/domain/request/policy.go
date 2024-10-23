package request

// Policy 保单信息
type Policy struct {
	PlanCode                string `xml:"planCode"`                          // 险种
	ProductCode             string `xml:"productCode"`                       // 产品计划
	Count                   string `xml:"count"`                             // 具体份数根据产品规则定义，从目前产品规则来看都为1份
	PaymentMode             string `xml:"paymentMode"`                       // 支付方式 非见费传0 见费传8
	LastPolicyNo            string `xml:"lastPolicyNo,omitempty"`            // 上年保单号（续保时必传）(否)
	TransApplyDate          string `xml:"transApplyDate"`                    // 申请时间：格式 不传默认当前系统时间 yyyy-mm-dd hh:mm:ss
	TransBeginDate          string `xml:"transBeginDate"`                    // 保险起期：格式 yyyy-mm-dd hh:mm:ss
	TransEndDate            string `xml:"transEndDate"`                      // 保险止期：格式 yyyy-mm-dd hh:mm:ss
	AgentCode               string `xml:"agentCode"`                         // 代理人，代理人节点和业务员节点必须有一个不为空
	SaleAgentCode           string `xml:"saleAgentCode,omitempty"`           // 业务员节点同（agentCode） (否)
	CurrencyCode            string `xml:"currencyCode"`                      // 币种:人民币 （01）、美元（04）、欧元（16）、港币（03）、新加坡元（08）
	RateScale               string `xml:"rateScale,omitempty"`               // 保单手续费，精度（3,2）(否)
	CustomerSource          string `xml:"customerSource,omitempty"`          // 客户来源 (否)
	DisputedSettleMode      string `xml:"disputedSettleMode,omitempty"`      // 争议解决方式 默认 参考字典 (否)
	GroupPersonFlag         string `xml:"groupPersonFlag,omitempty"`         // 意外险控制人数 (否)
	MedicalAmount           string `xml:"medicalAmount,omitempty"`           // 每次事故总赔偿限额 (否)
	UnderwriteScope         string `xml:"underwriteScope"`                   // 承保范围 0: 仅承保境内(不含港澳台) 1:全球范围
	CommissionFactor        string `xml:"commissionFactor"`                  // 佣金比例
	Region                  string `xml:"region,omitempty"`                  // 地区 001:全国（含湖北） 002:全国（除湖北） 参考字典 (否)
	BusSrePcmac             string `xml:"busSrePcmac,omitempty"`             // 终端设备的物理网卡地址 例如 E4-54-E8-A3-CC-BF (否)
	BusSrePcip              string `xml:"busSrePcip,omitempty"`              // 终端设备的ip地址 例如 10.138.91.99 (否)
	BusSrePcname            string `xml:"busSrePcname,omitempty"`            // 终端设备类型 可移动设备/不可移动设备 例如 不可移动设备 (否)
	CheckRealFlag           string `xml:"checkRealFlag,omitempty"`           // 是否调用查验接口标识 1就是调用查验，传0就不调用 (否)
	MerchantBusinessSource  string `xml:"merchantBusinessSource,omitempty"`  // 二级渠道商户号 (否)
	CompanyTravelYearPeople string `xml:"companyTravelYearPeople,omitempty"` // 公司旅行保险年单人数(T80年单计划必填) (否)
	TravelGuaranteePeriod   string `xml:"travelGuaranteePeriod,omitempty"`   // 每次旅行保障期限(T80年单计划必填)详细 (否)
	CompanyTravelPeopleDay  string `xml:"companyTravelPeopleDay,omitempty"`  // 公司总旅行人天数(T80短期计划必填) (否)
	TotalDutyQuota          string `xml:"totalDutyQuota,omitempty"`          // 保单累计赔偿限额(T80必填) (否)
	OneDutyQuota            string `xml:"oneDutyQuota,omitempty"`            // 每次事故赔偿限额(T80必填) (否)
	MininumberDays          string `xml:"mininumberDays,omitempty"`          // 年度最低投保天数(T80) (否)
	MininumberPerson        string `xml:"mininumberPerson,omitempty"`        // 年度最低投保人数(T80) (否)
}
