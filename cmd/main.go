package main

import accountmanager "github.com/StefanUA/AccountManager/pkg/accountManager.go"

func main() {
	accountManager := accountmanager.NewCommand()
	accountManager.Execute()
}
