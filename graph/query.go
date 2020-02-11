package graph

import (
	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{
	/*"sms": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "SmsQueryType",
			Fields:      graphql.Fields{},
			IsTypeOf:    nil,
			Description: "短信查询服务",
		}),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "", e
		},
		Description: "短信查询服务",
	},*/
}
