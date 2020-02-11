package graph

import (
	"encoding/json"
	"log"

	g "github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

// 非空字符串类型参数
func ArgString(description string) *g.ArgumentConfig {
	return &g.ArgumentConfig{Type: g.String, Description: description}
}
func ArgNonNullString(description string) *g.ArgumentConfig {
	return &g.ArgumentConfig{Type: g.NewNonNull(g.String), Description: description}
}

var _mutation = g.Fields{
	"sms": &g.Field{
		Name: "",
		Type: g.NewObject(g.ObjectConfig{
			Name:       "SmsMutationType",
			Interfaces: nil,
			Fields: g.Fields{
				"sendCaptcha": &g.Field{
					Description: "发送短信验证码",
					Type:        ResponseCodeType,
					Args: g.FieldConfigArgument{
						"orgId":        ArgNonNullString("机构编号"),
						"mobile":       ArgNonNullString("手机号码"),
						"templateId":   ArgNonNullString("短信模板编号"),
						"captchaId":    ArgNonNullString("图片验证码编号，通过图片验证码服务获取"),
						"captchaValue": ArgNonNullString("图片验证码用户输入值"),
					},
					Resolve: func(p g.ResolveParams) (i interface{}, e error) {
						bytes, e := json.Marshal(p.Args)
						log.Print("发送短信" + string(bytes))
						return ERR_SUCCESS, e
					},
				},
				"verifyCaptcha": &g.Field{
					Description: "验证短信验证码，服务端调用",
					Type:        ResponseCodeType,
					Args: g.FieldConfigArgument{
						"mobile": ArgNonNullString("手机号码"),
						"code":   ArgNonNullString("短信验证码"),
					},
					Resolve: func(p g.ResolveParams) (i interface{}, e error) {
						bytes, e := json.Marshal(p.Args)
						log.Print("验证短信" + string(bytes))
						return ERR_SUCCESS, e
					},
				},
			},
			IsTypeOf:    nil,
			Description: "短信服务更新类型",
		}),
		Resolve: func(p g.ResolveParams) (i interface{}, err error) {
			return "", err
		},
		Description: "短信服务",
	},
}
