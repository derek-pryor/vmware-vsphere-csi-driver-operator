package assets

import (
	"embed"
)

//go:embed *.yaml rbac/*.yaml webhook/*.yaml webhook/rbac/*.yaml
var f embed.FS

// ReadFile reads and returns the content of the named file.
func ReadFile(name string) ([]byte, error) {
	return f.ReadFile(name)
}
