package main

import (
	"fmt"
	gpowsh "win-log/components/gopowershell"
)

func main() {
	po := gpowsh.PowerShell{LocalExecPath: "powershell.exe"}

	posh, err := po.New()

	stdout, err := posh.Execute("Get-EventLog -List")
	fmt.Println(stdout)
	if err != nil {
		fmt.Println(err)
	}

}