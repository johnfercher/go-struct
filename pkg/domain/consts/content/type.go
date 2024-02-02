package content

type Type string

const (
	Go         Type = "go"
	GoMod      Type = "gomod"
	Dockerfile Type = "dockerfile"
	Makefile   Type = "makefile"
	Unknown    Type = "unknown"
)
