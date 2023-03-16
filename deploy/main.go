package main

import (
	"git-codecommit.eu-west-1.amazonaws.com/v1/repos/BigUpData.Transactions.Gui.git/integrate"
)

func main() {
	integration.ServerlessListener(integrate.Router, ":80")
}