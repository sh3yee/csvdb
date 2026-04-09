package row

type Row struct {
	Path string
}

func New(path string) *Row {
	return &Row{
		Path: path,
	}
}
