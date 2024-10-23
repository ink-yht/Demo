package request

// Plans 计划信息
type Plans struct {
	Plan Plan `xml:"plan"`
}

type Plan struct {
	PlanNo         string      `xml:"planNo,omitempty"` // 计划号，传纯数字序列 1 2 3... (否)
	PlanName       string      `xml:"planName"`         // 计划名称
	OccupationType string      `xml:"occupationType"`   // 职业类别（P21产品传） 参照职业码表
	Coverages      []Coverage  `xml:"coverages>coverage"`
	Insurances     []Insurance `xml:"insurances>insurance"`
}

// Coverage 责任信息
type Coverage struct {
	CoverageCode       string `xml:"coverageCode"`                 // 责任编码
	InsuredAmount      string `xml:"insuredAmount"`                // 保险金额（16,2）
	TotalActualPremium string `xml:"totalActualPremium,omitempty"` // 实交保费合计(16,2) (否)
	PaymentRatio       string `xml:"paymentRatio,omitempty"`       // 赔付比例(6,2) (否)
	DeductionAmount    string `xml:"deductionAmount,omitempty"`    // 免赔额（16,2） (否)
	DeductionDay       string `xml:"deductionDay,omitempty"`       // 免赔天数 (否)
	PerDayAmount       string `xml:"perDayAmount,omitempty"`       // 每次事故免赔额（16,2） (否)
	MaxDay             string `xml:"maxDay,omitempty"`             // 最高赔偿天数 (否)
	EachDay            string `xml:"eachDay,omitempty"`            // 每次赔偿天数（每次住院天数） (否)
	SelfPaymentRatio   string `xml:"selfPaymentRatio,omitempty"`   // 自费部分赔付比例（6,2） (否)
}

// Insurance 被保险人信息
type Insurance struct {
	InsuranceName      string        `xml:"insuranceName"`                // 被保险人姓名
	BirthDate          string        `xml:"birthDate"`                    // 被保险人生日。格式：yyyy-MM-dd
	CertiNo            string        `xml:"certiNo"`                      // 证件号
	CertiType          string        `xml:"certiType"`                    // 证件号类型
	RelApplicant       string        `xml:"relApplicant"`                 // 与投保人关系
	Gender             string        `xml:"gender"`                       // 被保险人性别，F：女 M:男
	OccupationCode     string        `xml:"occupationCode"`               // 职业代码（工种代码） 参考字典
	OccupationType     string        `xml:"occupationType"`               // 职业类型（P95,T80必传）
	NationalityType    string        `xml:"nationalityType"`              // 国籍
	Address            string        `xml:"address,omitempty"`            // 通信地址 (否)
	TargetFlag         string        `xml:"targetFlag,omitempty"`         // 是否不记名投保 (否)
	PersPhone          string        `xml:"persPhone,omitempty"`          // 手机号 (否)
	IsHaveSocial       string        `xml:"isHaveSocial,omitempty"`       // 是否有社保 (否)
	EmployeePost       string        `xml:"employeePost,omitempty"`       // 员工岗位 参考字典 (否)
	InsuranceStartTime string        `xml:"insuranceStartTime,omitempty"` // 标的保险起期 格式 yyyy-mm-dd hh:mm:ss(T80) (否)
	InsuranceEndTime   string        `xml:"insuranceEndTime,omitempty"`   // 标的保险止期 格式 yyyy-mm-dd hh:mm:ss(T80) (否)
	BeneficiaryType    string        `xml:"beneficiaryType,omitempty"`    // 受益人类型，1、法定，2、顺序，3、比例 (否)
	Beneficiaries      []Beneficiary `xml:"beneficiaries>beneficiary"`
	CertStartDate      string        `xml:"certStartDate,omitempty"`  // 证件起期 格式：2021-01-01 (否)
	CertEndDate        string        `xml:"certEndDate,omitempty"`    // 证件止期 格式：2021-01-01 (否)
	CertLongFlag       string        `xml:"certLongFlag,omitempty"`   // 是否长期有效 1：是；0：否 (否)
	InsureProvince     string        `xml:"insureProvince,omitempty"` // 所在省份 (否)
	InsureCity         string        `xml:"insureCity,omitempty"`     // 所在城市 (否)
	InsureDistrict     string        `xml:"insureDistrict,omitempty"` // 所在区 (否)
	FurthestCity       string        `xml:"furthestCity,omitempty"`   // 目的地详细 (否)
}

// Beneficiary 受益人信息
type Beneficiary struct {
	BenefitName        string `xml:"benefitName,omitempty"`        // 受益人姓名 (否)
	BenefitCertiType   string `xml:"benefitCertiType,omitempty"`   // 受益人证件类型，同被保人证件类型 (否)
	BenefitCertiNo     string `xml:"benefitCertiNo,omitempty"`     // 受益人证件号 (否)
	BenefitRelation    string `xml:"benefitRelation,omitempty"`    // 受益人与被保人关系 0-配偶，1-子女 2-父母，3-其他 (否)
	BenefitScale       string `xml:"benefitScale,omitempty"`       // 受益比例，受益人类型为比例是传入例：0.75 (否)
	BenefitOrderNo     string `xml:"benefitOrderNo,omitempty"`     // 受益顺序，受益人类型为顺位时传入 例（整数）：1、2、3 (否)
	CertStartDate      string `xml:"certStartDate,omitempty"`      // 证件起期 格式：2021-01-01 (否)
	CertEndDate        string `xml:"certEndDate,omitempty"`        // 证件止期 格式：2021-01-01 (否)
	CertLongFlag       string `xml:"certLongFlag,omitempty"`       // 是否长期有效 1：是；0：否 (否)
	BeneficiaryAddress string `xml:"beneficiaryAddress,omitempty"` // 受益人联系地址 (否)
	BeneficiaryPhone   string `xml:"beneficiaryPhone,omitempty"`   // 受益人联系电话 (否)
}
