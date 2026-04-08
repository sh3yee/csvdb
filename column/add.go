package column

import "gycsv/file"

// Add 添加新列
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
