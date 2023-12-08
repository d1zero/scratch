package templates

type GoModData struct {
	ModuleName string
}

const GoModTemplate = `module {{.ModuleName}}

go 1.21

require (
	github.com/d1zero/scratch v0.0.15
)
`
