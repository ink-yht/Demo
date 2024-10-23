package response

// Policy 保单信息响应
type Policy struct {
	PlanCode           string `xml:"planCode"`           // 险种
	ProductCode        string `xml:"productCode"`        // 产品计划
	QuotationNo        string `xml:"quotationNo"`        // 报价成功，返回报价单号
	CurrencyRate       string `xml:"currencyRate"`       // 币种汇率（16,6）
	TotalNetPremium    string `xml:"totalNetPremium"`    // 净额保费(16,2)
	AddedValueTax      string `xml:"addedValueTax"`      // 增值税(16,2)
	TotalActualPremium string `xml:"totalActualPremium"` // 复核保费(16,2)
}
