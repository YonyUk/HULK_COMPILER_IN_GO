package fylesystempackage

func IFileReaderCode() string {
	return `package filesystem

// File reader interface
type IFileReader interface {
	ReadFile() (string, error)
}
`
}
