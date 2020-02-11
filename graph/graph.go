package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hgraph"
)

var (
	GraphqlHttpHandler = hgraph.GraphqlHttpHandler
	QueryFields        = hgraph.MergeQueryFields
	MutationFields     = hgraph.MergeMutationFields

	String           = graphql.String
	NewNonNullString = graphql.NewNonNull(graphql.String)
)

// 响应状态码
var ResponseCodeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ResponseCode",
	Fields: graphql.Fields{
		"code": &graphql.Field{Type: graphql.String, Description: "响应码"},
		"msg":  &graphql.Field{Type: graphql.String, Description: "响应描述"},
	},
	Description: "响应状态码，‘00000’成功，其它错误码见接口描述",
})
