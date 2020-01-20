package stub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TestCase struct {
	List []Loan `json:"list"`
}

type Loan struct {
	Path       string      `json:"path"`
	Request    Request     `json:"request"`
	Response   interface{} `json:"response"`
	StatusCode int         `json:"status_code"`
}

type Request struct {
	Loan struct {
		ProductIntent      string  `json:"productIntent"`
		TotalRequestAmount float64 `json:"totalRequestAmount"`
		PaymentFrequency   string  `json:"paymentFrequency"`
		LoanTenor          int     `json:"loanTenor"`
		InstallmentAmount  float64 `json:"installmentAmount"`
		GracePeriod        int     `json:"gracePeriod"`
		DueDay             int     `json:"dueDay"`
	} `json:"loan"`
}

type Response struct {
	Data struct {
		Loan struct {
			TotalRequestAmount float64   `json:"totalRequestAmount"`
			LoanTenor          int       `json:"loanTenor"`
			InstallmentAmount  float64   `json:"installmentAmount"`
			PaymentFrequency   string    `json:"paymentFrequency"`
			InterestRate       int       `json:"interestRate"`
			FirstDueDate       time.Time `json:"firstDueDate"`
			Installment        struct {
				MinAmount float64 `json:"minAmount"`
				MaxAmount float64 `json:"maxAmount"`
			} `json:"installment"`
		} `json:"loan"`
	} `json:"data"`
}

func Route(route *gin.Engine, dataPath string) {
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		logrus.Fatalln("contact ReadFile", err.Error())
	}

	var testCases []Loan
	err = json.Unmarshal(data, &testCases)
	if err != nil {
		logrus.Fatalln("setup error Unmarshal", err.Error())
	}

	route.POST("/v1/loanorigination/simulator/calculate", func(context *gin.Context) {
		var request Request
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "error na ja"})
			return
		}

		for _, mockData := range testCases {
			if request.Loan.ProductIntent == mockData.Request.Loan.ProductIntent {
				context.JSON(http.StatusOK, mockData.Response)
				return
			}
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "ProductIntent ไม่ตรง"})
	})
}
