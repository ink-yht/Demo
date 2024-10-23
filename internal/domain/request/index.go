package request

type TXTpaic struct {
	TXTpaicRequest TXTpaicRequest `xml:"TXTpaicRequest"`
}

// TXTpaicRequest 请求结构体
type TXTpaicRequest struct {
	Head Head `xml:"head"`
	Body Body `xml:"body"`
}

// Head 请求头
type Head struct {
	TransRefId   string `xml:"transRefId"`   // 交易流水号，唯一性，由请求发起方产生，相同视为同一批次请求
	TransType    string `xml:"transType"`    // 交易类型:001详细
	TransExeDate string `xml:"transExeDate"` // 发送请求消息日期： yyyy-mm-dd hh:mm:ss
}

// Body 请求体
type Body struct {
	Policy    Policy    `xml:"policy"`
	Applicant Applicant `xml:"applicant"`
	Plans     Plans     `xml:"plans"`
}
