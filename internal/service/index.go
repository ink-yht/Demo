package service

import (
	"Demo/internal/domain/response"
	"Demo/internal/repository"
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strings"
)

type Service struct {
	repo *repository.Repository
}

func NewServe(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// SendRequest 发送请求
func (svc *Service) SendRequest() (*response.TXTpaic, error) {

	// 获取请求体
	request, err := svc.repo.CreateRequest()
	if err != nil {
		log.Fatalf("获取请求体失败: %v", err)
	}

	// 将请求体编码为XML
	xmlBytes, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		log.Fatalf("编码失败: %v", err)
	}

	// 创建HTTP客户端和请求
	client, req := CreateClient(err, xmlBytes)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 创建一个新的xml.Decoder并设置CharsetReader
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charsetReader

	// 解析XML数据
	var responseStruct response.TXTpaic
	err = decoder.Decode(&responseStruct)
	if err != nil {
		log.Fatalf("解析 xml 数据失败: %v", err)
	}

	// 下面代码是将响应写到xml文件中，如不使用，则写到结构体中

	//xmlresponse, err := xml.MarshalIndent(responseStruct, "", "  ")
	//if err != nil {
	//	log.Fatalf("编码失败: %v", err)
	//}
	//
	//// 检查response.xml文件是否存在
	//if _, err := os.Stat("response.xml"); os.IsNotExist(err) {
	//	// 文件不存在，创建文件
	//	_, err := os.Create("response.xml")
	//	if err != nil {
	//		log.Fatalf("Failed to create file: %v", err)
	//	}
	//}
	//
	//// 将XML数据写入文件
	//err = ioutil.WriteFile("response.xml", xmlresponse, 0644)
	//if err != nil {
	//	log.Fatalf("Failed to write XML to file: %v", err)
	//}

	log.Printf("解析成功: %+v", responseStruct)
	return &responseStruct, nil
}

func CreateClient(err error, xmlBytes []byte) (*http.Client, *http.Request) {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建POST请求
	req, err := http.NewRequest("POST", "https://uats.axa.cn/webservice/groupApply.do", bytes.NewBuffer(xmlBytes))
	if err != nil {
		log.Fatalf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "text/xml;charset=GBK")

	// 设置基本认证
	req.SetBasicAuth("avaiTest", "password")
	return client, req
}

// 自定义的CharsetReader函数
func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, "gbk") || strings.EqualFold(charset, "gb2312") {
		reader := transform.NewReader(input, simplifiedchinese.GBK.NewDecoder())
		return reader, nil
	}
	return nil, fmt.Errorf("unsupported charset: %s", charset)
}
