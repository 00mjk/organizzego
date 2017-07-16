package main

import (
	"github.com/B0go/organizzego/model"
)

func main() {
	organizzeApi := OrganizzeApi{"username", "api-token", "You (you@yourdomain.com)"}

	organizzeApi.CreateStatement(model.Statement{"blah2", "2017-07-16", 10})
}
