package templates

type GoModData struct {
	ModuleName string
}

const GoModTemplate = `module {{.ModuleName}}

go 1.21
`
