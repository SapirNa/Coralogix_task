package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type ReadCSV struct {
	file_name string
	Reader    chan []string
	output    [][]string
}

func Read(file_name string) *ReadCSV {
	/*
		contractor for ReadCSV.
		get the file name path and create new object of CSV reader.
	*/
	r := &ReadCSV{file_name: file_name, Reader: Reader(file_name)}
	return r
}

func (r ReadCSV) GetOutput() [][]string {
	// return the final output of operations
	if r.output != nil {
		return r.output
	} else {
		return [][]string{}
	}
}

func Reader(file_name string) chan []string {
	/*
		exectue code for the CSV reader.
		read from the CSV line and move the data to the next operation.
	*/
	ch := make(chan []string)
	go func() {
		defer close(ch)
		inputFile, err := os.Open(file_name)
		if err != nil {
			log.Fatal(err)
		}
		csvReader := csv.NewReader(inputFile)
		for record, err := csvReader.Read(); err == nil; record, err = csvReader.Read() {
			if err != nil {
				if err == io.EOF {
					log.Print("End of file")
				} else {
					log.Fatal(err)
				}
				break
			}
			ch <- record
		}
	}()
	return ch
}

func Write(table [][]string, path string) {
	// create new CSV with the path and data from Read class
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)

	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()

	for _, row := range table {
		csvWriter.Write(row)
	}
}

func (r ReadCSV) FilterRows(record_num int, condition string) ReadCSV {
	// operation that filter row by column number (record num) and the condition
	temp := [][]string{}
	var con = condition
	for line := range r.Reader {
		if line[record_num] == con {
			temp = append(temp, line)
		}
	}
	r.output = temp
	return r
}

func (r ReadCSV) GetColumn(column_num int) ReadCSV {
	// operation that return one column from the Reader output
	if r.output == nil {
		temp := make([][]string, 0)
		for line := range r.Reader {
			temp = append(temp, line)
		}
		r.output = temp
	}
	column := make([]string, 0)
	for _, line := range r.output {
		column = append(column, line[column_num])
	}
	r.output = make([][]string, 0)
	for i := range column {
		t := make([]string, 0)
		t = append(t, column[i])
		r.output = append(r.output, t)
	}
	return r
}

func (r ReadCSV) Sum_column() ReadCSV {
	// operation that sum the column
	avg_column := make([]string, 0)
	sum_column := 0
	for _, line := range r.output {
		for idx := range len(r.output) - 1 {
			num, err := strconv.Atoi(line[idx])
			if err != nil {
				log.Fatal(err)
			}
			sum_column += num
		}
	}
	avg_column = append(avg_column, strconv.Itoa(sum_column))
	r.output = make([][]string, 0)
	r.output = append(r.output[:0], avg_column)
	return r
}
