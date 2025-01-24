/*
Copyright © 2022 Telnyx LLC
*/
package main

import (
	"github.com/team-telnyx/telnyx-cli/cmd"
	_ "github.com/team-telnyx/telnyx-cli/cmd/get"
)

func main() {
	cmd.Execute()
}
