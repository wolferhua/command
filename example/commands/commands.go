package commands

import "github.com/wolferhua/command"

func init() {
	command.Registry(
		&Test{},
	)
}
