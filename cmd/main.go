package main

import (
	"github.com/B0go/organizzego"
	"github.com/B0go/organizzego/model"
)

func main() {
	organizzeApi := organizzego.OrganizzeApi{"username", "api-token", "You (you@yourdomain.com)"}

	installmentsAttributes := model.InstallmentsAttributes{"monthly", 3}

	statement2 := model.Statement{}.New("blah2", "2017-07-19", 10, false)
	statement2.CategoryID = 10
	statement2.AccountID = 10
	statement2.InstallmentsAttributes = &installmentsAttributes

	organizzeApi.CreateStatement(statement2)
}
