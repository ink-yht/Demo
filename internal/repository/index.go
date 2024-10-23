package repository

import (
	"Demo/internal/domain/request"
	"time"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

// CreateRequest 构造请求体
func (repo *Repository) CreateRequest() (*request.TXTpaic, error) {
	head, policy, applicant, plans := ConstructData()
	requestBody := request.TXTpaicRequest{
		Head: head,
		Body: request.Body{
			Policy:    policy,
			Applicant: applicant,
			Plans:     plans,
		},
	}

	textBody := request.TXTpaic{
		TXTpaicRequest: requestBody,
	}

	return &textBody, nil
}

// ConstructData 构造请求数据
func ConstructData() (request.Head, request.Policy, request.Applicant, request.Plans) {
	// 当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 构造请求头
	head := request.Head{
		TransRefId:   "1234567890",
		TransType:    "001",
		TransExeDate: currentTime,
	}

	// 构造保单信息
	policy := request.Policy{
		PlanCode:         "P95",
		ProductCode:      "P95002",
		Count:            "1",
		PaymentMode:      "0",
		TransApplyDate:   currentTime,
		TransBeginDate:   "2023-10-01 00:00:00",
		TransEndDate:     "2024-10-01 00:00:00",
		AgentCode:        "agent81867",
		CurrencyCode:     "01",
		UnderwriteScope:  "0",
		CommissionFactor: "0.10",
	}

	// 构造投保人信息
	applicant := request.Applicant{
		ApplicationPersonnelType: "0",
		Enterprise:               "ABC Company",
		EnterpriseCreditCode:     "123456789012345678",
		StateProperty:            "1",
		LinkManName:              "John Doe",
		LinkManMobileTelephone:   "12345678901",
		CertType:                 "1",
		CertificateNo:            "123456789012345678",
		CertificateType:          "1",
		PersonTypeOverseas:       "0",
		PersonName:               "John Doe",
		PersonCertiCode:          "123456789012345678",
		PersonCertiType:          "1",
		PostCode:                 "100000",
		AppProvinceCode:          "110000",
		AppCityCode:              "110100",
		AppDistrictCode:          "110101",
	}

	plans := request.Plans{
		Plan: request.Plan{
			PlanNo:         "1",
			PlanName:       "Plan A",
			OccupationType: "123",
			Coverages: []request.Coverage{
				{
					CoverageCode:  "COV123",
					InsuredAmount: "100000.00",
				},
			},
			Insurances: []request.Insurance{
				{
					InsuranceName:   "Alice Smith",
					BirthDate:       "1990-01-01",
					CertiNo:         "123456789012345678",
					CertiType:       "1",
					RelApplicant:    "5",
					Gender:          "M",
					OccupationCode:  "123",
					OccupationType:  "1",
					NationalityType: "01",
				},
			},
		},
	}
	return head, policy, applicant, plans
}
