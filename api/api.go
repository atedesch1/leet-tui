package api

import (
	"github.com/machinebox/graphql"
)

var client *graphql.Client

// var ctx *context.Context

func Init() {
	client = graphql.NewClient("https://leetcode.com/graphql")
	// ctx = context.Background()
}

func GetClient() *graphql.Client {
	return client
}
