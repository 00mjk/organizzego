package main

import (
	"fmt"

	"github.com/B0go/organizzego"
	"github.com/B0go/organizzego/model"
)

func main() {
	organizzeApi := organizzego.OrganizzeApi{"you@yourdomain.com", "api-key", "You (you@yourdomain.com)"}

	installmentsAttributes := model.InstallmentsAttributes{"monthly", 3}

	statement2 := model.Statement{}.New("blah2", "2017-07-19", 10, false)
	statement2.CategoryID = 10
	statement2.AccountID = 10
	statement2.InstallmentsAttributes = &installmentsAttributes

	organizzeApi.CreateStatement(statement2)

	accounts := organizzeApi.FindAccounts()

	fmt.Printf("\n %s \n", accounts[0].Name)

	creditcards := organizzeApi.FindCreditCards()

	fmt.Printf("\n %s \n", creditcards[0].Name)
}
