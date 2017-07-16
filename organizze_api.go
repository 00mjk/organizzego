package main

import (
	"github.com/B0go/organizzego/model"
	"github.com/B0go/organizzego/requests"
	"github.com/B0go/organizzego/transformation"
)

type OrganizzeApi struct {
	Username  string
	Password  string
	UserAgent string
}

func (organizzeApi OrganizzeApi) CreateStatement(statement model.Statement) {
	str := transformation.JsonTransformation{}.MarshalToJSON(statement)

	httpRequests := requests.HttpRequests{organizzeApi.Username, organizzeApi.Password, organizzeApi.UserAgent}

	httpRequests.PostToOrganizze("https://api.organizze.com.br/rest/v2/transactions", str)
}
