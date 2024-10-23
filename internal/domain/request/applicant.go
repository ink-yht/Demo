package request

// Applicant 投保人信息
type Applicant struct {
	ApplicationPersonnelType string `xml:"applicationPersonnelType"`         // 投保人类型:0公司，1 个人
	Enterprise               string `xml:"enterprise"`                       // 企业名称(团体)
	EnterpriseCreditCode     string `xml:"enterpriseCreditCode"`             // 企业信用代码证号(团体)
	IndustryAttribute        string `xml:"industryAttribute,omitempty"`      // 企业行业属性(团体) (否)
	Address                  string `xml:"address,omitempty"`                // 公司地址 (否)
	StateProperty            string `xml:"stateProperty"`                    // 单位性质（团体）详细
	Email                    string `xml:"email,omitempty"`                  // 邮箱（团体） (否)
	LinkManName              string `xml:"linkManName"`                      // 联系人名称(团体)
	LinkManSexCode           string `xml:"linkManSexCode,omitempty"`         // 联系人性别(团体) (否)
	LinkManMobileTelephone   string `xml:"linkManMobileTelephone"`           // 联系人手机号码(团体) (否)
	LinkManOfficeAreaCode    string `xml:"linkManOfficeAreaCode,omitempty"`  // 联系人办公室电话区号(团体) (否)
	LinkManOfficeTelephone   string `xml:"linkManOfficeTelephone,omitempty"` // 联系人办公室电话号码(团体) (否)
	LinkManDepartment        string `xml:"linkManDepartment,omitempty"`      // 联系人所属部门(团体) (否)
	LinkManPosition          string `xml:"linkManPosition,omitempty"`        // 联系人职位(团体) (否)
	CertType                 string `xml:"certType"`                         // 纳税人类型 1:小规模纳税人,2:增值税一般纳税人,3:非增值税纳税人
	OrgCode                  string `xml:"orgCode,omitempty"`                // 工商注册号(团体) (否)
	TaxRegisterNo            string `xml:"taxRegisterNo,omitempty"`          // 税务登记证(团体) (否)
	BusinessRegisterId       string `xml:"businessRegisterId,omitempty"`     // 组织机构代码证(团体) (否)
	CertificateNo            string `xml:"certificateNo"`                    // 证件号码(团体)
	CertificateType          string `xml:"certificateType"`                  //证件类型参考字典
	PersonTypeOverseas       string `xml:"personTypeOverseas"`               // 0：境内 1 境外 默认境内
	PersonName               string `xml:"personName"`                       // 投保人姓名（个人）
	PersonCertiCode          string `xml:"personCertiCode"`                  // 投保人证件号（个人）
	PersonCertiType          string `xml:"personCertiType"`                  // 投保人证件类型（个人）
	PersonGender             string `xml:"personGender,omitempty"`           // 投保人性别（个人）M男F女 (否)
	PersonBirthDate          string `xml:"personBirthDate,omitempty"`        // 投保人出生日期（个人） 格式：yyyy-mm-dd (否)
	PersonEmail              string `xml:"personEmail,omitempty"`            // 投保人邮箱（个人） (否)
	PersonPhone              string `xml:"personPhone,omitempty"`            // 投保人手机号（个人） (否)
	PersonAdderss            string `xml:"personAdderss,omitempty"`          // 投保人地址（个人） (否)
	PostCode                 string `xml:"postCode"`                         // 邮编
	Nationality              string `xml:"nationality,omitempty"`            // 国籍 (否)
	CertStartDate            string `xml:"certStartDate,omitempty"`          // 证件起期（个人） 格式：2021-01-01 (否)
	CertEndDate              string `xml:"certEndDate,omitempty"`            // 证件止期（个人） 格式：2021-01-01 (否)
	CertLongFlag             string `xml:"certLongFlag,omitempty"`           // 是否长期有效（个人） 1：是；0：否 (否)
	IndustryLevel1           string `xml:"industryLevel1,omitempty"`         // （P21使用）行业等级1 (否)
	IndustryLevel2           string `xml:"industryLevel2,omitempty"`         // （P21使用）行业等级2 (否)
	IndustryLevel3           string `xml:"industryLevel3,omitempty"`         // （P21使用）行业等级3 (否)
	IndustryLevel4           string `xml:"industryLevel4,omitempty"`         // （P21使用）行业等级4 (否)
	InsureProvince           string `xml:"insureProvince,omitempty"`         // 所在省份 (否)
	InsureCity               string `xml:"insureCity,omitempty"`             // 所在城市 (否)
	InsureDistrict           string `xml:"insureDistrict,omitempty"`         // 所在区 (否)
	AppProvinceCode          string `xml:"appProvinceCode"`                  // 投保人省
	AppCityCode              string `xml:"appCityCode"`                      // 投保人市
	AppDistrictCode          string `xml:"appDistrictCode"`                  // 投保人区
}
