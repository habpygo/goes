package main

import (
	"TheGoProgrammingLanguage/websnippets/bankapp/web"
	"TheGoProgrammingLanguage/websnippets/bankapp/web/bankaccounts"
	"TheGoProgrammingLanguage/websnippets/bankapp/web/controllers"
)

func main() {

	myAccount := &bankaccounts.BankAccount{}

	app := &controllers.Application{
		Bank: myAccount,
	}
	web.Serve(app)
}
