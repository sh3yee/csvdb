package column

import "gycsv/internal/file"

// Alter 修改列名
func (c *Column) Alter(before string, after string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	for i := 0; i < len(rwf.Header); i++ {
		if rwf.Header[i] == before {
			rwf.Header[i] = after
		}
	}

	return rwf.Write()
}
