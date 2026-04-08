package column

import (
	"gycsv/file"
)

type Column struct {
	Path string
}

func New(path string) *Column {
	return &Column{
		Path: path,
	}
}

// Add add a new column
func (c *Column) Add(field string) error {
	rwf := file.New(c.Path)

	err := rwf.ReadHeader()
	if err != nil {
		return err
	}

	err = rwf.ReadTable()
	if err != nil {
		return err
	}

	rwf.Header = append(rwf.Header, field)

	err = rwf.WriteHeader()
	if err != nil {
		return err
	}

	err = rwf.WriteTable()
	if err != nil {
		return err
	}

	return nil
}


// Alter alter old column name to new column name
func (c *Column) Alter(before string, after string) error {
	rwf := file.New(c.Path)

	err := rwf.ReadHeader()
	if err != nil {
		return err
	}

	err = rwf.ReadTable()
	if err != nil {
		return err
	}

	for i := 0; i < len(rwf.Header); i++ {
		if rwf.Header[i] == before {
			rwf.Header[i] = after
		}
	}

	err = rwf.WriteHeader()
	if err != nil {
		return err
	}

	err = rwf.WriteTable()
	if err != nil {
		return err
	}

	return nil
}

// TODO Delete column
