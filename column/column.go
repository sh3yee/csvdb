package column

type Column struct {
	Path string
}

func New(path string) *Column {
	return &Column{
		Path: path,
	}
}
