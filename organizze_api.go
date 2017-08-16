package organizzego

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

func (organizzeApi OrganizzeApi) FindAccounts() []model.BankAccount {

	accounts := []model.BankAccount{}

	organizzeApi.find("https://api.organizze.com.br/rest/v2/accounts", &accounts)

	return accounts
}

func (organizzeApi OrganizzeApi) FindCreditCards() []model.CreditCard {

	creditCards := []model.CreditCard{}

	organizzeApi.find("https://api.organizze.com.br/rest/v2/credit_cards", &creditCards)

	return creditCards
}

func (organizzeApi OrganizzeApi) find(url string, data interface{}) {
	httpRequests := requests.HttpRequests{organizzeApi.Username, organizzeApi.Password, organizzeApi.UserAgent}

	jsonBytes := httpRequests.GetFromOrganizze(url)

	transformation.JsonTransformation{}.UnmarshalFromJSON(jsonBytes, data)
}
