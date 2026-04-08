package file

import (
	"bufio"
	"encoding/csv"
	"io"
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

func (d *Data) ReadHeader() error {

	rf, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer rf.Close()

  // at begin of file, offset is 0
	_, err = rf.Seek(0, 0)
	if err != nil {
		return err
	}

	reader := csv.NewReader(rf)
	column, err := reader.Read()
	if err != nil {
		return err
	}

	d.Header = column
	return nil
}

func (d *Data) ReadTable() error {

	rf, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer rf.Close()

	br := bufio.NewReader(rf)
	line, _, err := br.ReadLine()
	if err != nil {
		return err
	}

	_, err = rf.Seek(int64(len(line))+1, 0)
	if err != nil {
		return err
	}

	cr := csv.NewReader(rf)
	table, err := cr.ReadAll()
	if err != nil {
		return err
	}

	d.Table = table
	return nil
}

func (d *Data) WriteHeader() error {
	wf, err := os.OpenFile(d.Path, os.O_RDWR|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		return err
	}
	defer wf.Close()

	// 偏移到第一行
	_, err = wf.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// 写入列名
	writer := csv.NewWriter(wf)
	defer writer.Flush()

	err = writer.Write(d.Header)
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) WriteTable() error {

	wf, err := os.OpenFile(d.Path, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(wf)
	line, _, err := reader.ReadLine()
	if err != nil {
		return err
	}

	// offset one line
	_, err = wf.Seek(int64(len(line))+1, io.SeekStart)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(wf)
	defer writer.Flush()

	for _, t := range d.Table {
		err = writer.Write(t)
		if err != nil {
			return err
		}
	}

	return nil
}
