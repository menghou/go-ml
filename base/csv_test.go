package base

import (
	"fmt"
	"testing"
)

func TestCSVReader_ParseRowCount(t *testing.T) {
	reader := NewCSVReader("../examples/datasets/iris.csv", false)
	err, rowCount := reader.ParseRowCount()
	if err != nil {
		t.Error(err)
	}
	if rowCount != 150 {
		t.Errorf("wrong rowCount :%d", rowCount)
	}
}

func TestCSVReader_String(t *testing.T) {
	reader := NewCSVReader("../examples/datasets/iris.csv", false)
	err, instance := reader.Read()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(instance)
}
func TestCSVReader_ParseMaxPrecision(t *testing.T) {
	reader := NewCSVReader("../examples/datasets/iris.csv", false)
	err, maxPrecision := reader.ParseMaxPrecision()
	if err != nil {
		t.Error(err)
	}
	if maxPrecision != 1 {
		t.Errorf("max precision received %d", maxPrecision)
	}
}
