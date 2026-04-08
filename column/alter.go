package column

import "gycsv/file"

// Alter 修改列名
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
