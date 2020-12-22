package envelopes

import (
	"fmt"
	"github.com/shopspring/decimal"
	"learn-go/web/core/boot"
	"learn-go/web/enums"
	"learn-go/web/util"
	"testing"
)

func TestMain(m *testing.M) {
	_ = boot.TestEnv(
		"C:\\Users\\bodenai\\go\\src\\learn-go\\web\\resource\\application.yml")
	m.Run()
}

func TestRedEnvelopeAmount(t *testing.T) {
	var mean = make([]decimal.Decimal, 10)
	var temp = make([]decimal.Decimal, 10)
	count := 20000
	for i := 0; i < 10; i++ {
		mean[i] = decimal.NewFromInt(0)
	}
	for j := 0; j < count; j++ {
		var totalAmount = decimal.NewFromFloat(88.88)
		for i := 10; i > 0; i-- {
			amount := util.GenerateRedEnvelopeAmount(int64(i), totalAmount, enums.TryLuck)
			totalAmount = totalAmount.Sub(amount)
			temp[i-1] = amount
		}
		for k := 0; k < 10; k++ {
			mean[k] = mean[k].Add(temp[k].Div(decimal.NewFromInt(int64(count))))
		}
	}
	var total = decimal.NewFromFloat(0)
	for i := 0; i < 10; i++ {
		total = total.Add(mean[i])
		fmt.Print(mean[i].String() + " ")
	}
	fmt.Println(total.String())
}
