package model

type Cmd interface {
	Run(args []string)
	Help()
	ShortHelp()
}
