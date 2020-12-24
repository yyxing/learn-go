package util

import (
	jsoniter "github.com/json-iterator/go"
	imgext "github.com/shamsher31/goimgext"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
	"learn-go/web/enums"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	SystemAccountNo = "6225775552527361"
)

var (
	seed = rand.NewSource(time.Now().UnixNano())
)

var (
	snowflake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Time{},
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func ModelToMap(model interface{}) map[string]interface{} {
	var resultMap map[string]interface{}
	data, err := Marshal(model)
	if err != nil {
		logrus.Errorf("json Marshal failed %s", err)
	}
	if err = Unmarshal(data, &resultMap); err != nil {
		logrus.Errorf("json Unmarshal failed %s", err)
	}
	return resultMap
}

func Marshal(model interface{}) ([]byte, error) {
	json, err := json.Marshal(model)
	return json, err
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

func NextId() string {
	id, err := snowflake.NextID()
	for err != nil {
		id, err = snowflake.NextID()
	}
	return strconv.FormatUint(id, 10)
}

func GenerateAccountNo() string {
	prefix := "622577"
	id, _ := snowflake.NextID()
	suffix := strconv.FormatUint(id, 10)
	return prefix + suffix[len(suffix)-10:]
}

func GenerateRedEnvelopeAmount(remainQuantity int64, remainAmount decimal.Decimal, envelopeType enums.RedEnvelopeType) decimal.Decimal {
	if envelopeType == 1 {
		return remainAmount.Div(decimal.NewFromInt(remainQuantity))
	}
	penny := decimal.NewFromInt(100)
	if remainQuantity == 1 {
		return remainAmount.Div(penny).Mul(penny)
	}
	min := decimal.NewFromFloat(1)
	max := remainAmount.Div(decimal.NewFromInt(remainQuantity)).Mul(decimal.NewFromInt(200))
	random := rand.New(seed)
	return decimal.NewFromInt(random.Int63n(max.IntPart())).Add(min).Div(penny)
}

func GetFileType(bytes []byte) string {
	filetype := http.DetectContentType(bytes[:20])

	ext := imgext.Get()

	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:]) {
			return filetype
		}
	}
	return ""
}
