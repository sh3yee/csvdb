package file

import (
	"encoding/csv"
	"io/fs"
	"os"
)

type Data struct {
	Path   string
	Header []string
	Table  [][]string
}

func New(path string) *Data {
	return &Data{
		Path: path,
	}
}

// Read 读取整个文件（header + table）
func (d *Data) Read() error {
	rf, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer rf.Close()

	reader := csv.NewReader(rf)

	header, err := reader.Read()
	if err != nil {
		return err
	}
	d.Header = header

	table, err := reader.ReadAll()
	if err != nil {
		return err
	}
	d.Table = table

	return nil
}

// Write 写入整个文件（header + table）
func (d *Data) Write() error {
	wf, err := os.OpenFile(d.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		return err
	}
	defer wf.Close()

	writer := csv.NewWriter(wf)

	if err := writer.Write(d.Header); err != nil {
		return err
	}

	for _, row := range d.Table {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	writer.Flush()
	return writer.Error()
}
