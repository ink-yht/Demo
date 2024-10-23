package response

type TXTpaic struct {
	TXTpaicResponse TXTpaicResponse `xml:"TXTpaicResponse"`
}

// TXTpaicResponse 响应结构体
type TXTpaicResponse struct {
	Head Head `xml:"head"`
	Body Body `xml:"body"`
}

// Head 响应头
type Head struct {
	TransRefId   string `xml:"transRefId"`        // 交易流水号，唯一性，由请求发起方产生，相同视为同一批次请求
	TransType    string `xml:"transType"`         // 交易类型: 对应操作类型
	TransExeDate string `xml:"transExeDate"`      // 发送请求消息日期： yyyy-mm-dd hh:mm:ss
	Message      string `xml:"message,omitempty"` // 交易信息描述 (否)
	Status       string `xml:"status"`            // 交易状态:0------正常-1----异常
	StatusCode   string `xml:"statusCode"`        // 交易状态代码0000----成功 -999----其他未知错误详细
}

// Body 响应体
type Body struct {
	Policy Policy `xml:"policy"`
	Plans  Plan   `xml:"plans>plan"`
}
