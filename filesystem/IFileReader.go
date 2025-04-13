package filesystem

// File reader interface
type IFileReader interface {
	ReadFile() (string, error)
}
