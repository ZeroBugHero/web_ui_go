package reports

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
)

type Validator struct {
	Id              int             `json:"id"`
	Type            string          `json:"type"`             // 断言类型
	Locator         string          `json:"locator"`          // 定位器
	Expect          interface{}     `json:"expect"`           // 期望值
	Actual          interface{}     `json:"actual"`           // 实际值
	Message         string          `json:"msg,omitempty"`    // 错误信息
	ValidatorStatic ValidatorStatic `json:"validator_static"` // 验证器静态信息
}

type ValidatorStatic struct {
	Screenshot string `json:"screenshot"`
	Video      string `json:"video"`
	Trace      string `json:"trace"`
}

// validationResults 存储了一系列的 Validator 结果
var validationResults []Validator

// AddResult 将一个新的 Validator 添加到结果切片中
func (val *Validator) AddResult() {
	log.Info().Msgf("添加验证结果：%v", val)
	validationResults = append(validationResults, *val)
	log.Info().Msgf("验证结果：%v", validationResults)
	SaveResultsToJSONFile(validationResults)
}

// SaveResults 将所有验证结果保存到 JSON 文件
func SaveResults(validationResults []Validator) error {
	return SaveResultsToJSONFile(validationResults)
}

// SaveResultsToJSONFile 将验证结果写入 JSON 文件
func SaveResultsToJSONFile(validationResults []Validator) error {
	file, err := os.OpenFile("result.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(validationResults)
}
