package base

import (
	"os"
	"encoding/csv"
	"strings"
	"regexp"
	"bufio"
	"fmt"
)

type CSVReader struct {
	filename string
	hasHeader bool
}

func (reader *CSVReader) Read() (err error, dataSet DataSet){
	f, err := os.Open(reader.filename)
	if err != nil {
		return
	}
	defer f.Close()
	err, features := reader.ParseFeatures()
	if err != nil {
		return
	}
	dataSet = NewDataGrid()
	for _, f := range features {
		err := dataSet.AddFeature(f)
		if err != nil {
			return
		}
	}
	err = reader.BuildDataSetFromReader(features, dataSet)
	if err != nil {
		return
	}
	return
}

func (reader *CSVReader) BuildDataSetFromReader(features []Feature, dataSet DataSet) error {

}

func (reader *CSVReader) ParseFeatures() (err error, features []Feature) {
	err, features = reader.ParseFeatureType()
	if err != nil {
		return
	}
	err, names := reader.ParseFeatureName()
	if err != nil {
		return
	}
	for i, feature := range features {
		feature.setName(names[i])
	}
	return
}

func (reader *CSVReader) ParseFeatureType() (err error, features []Feature) {
	f, err := os.Open(reader.filename)
	if err != nil {
		return
	}
	defer f.Close()
	freader := csv.NewReader(f)
	if reader.hasHeader {
		_, err = freader.Read()
		if err != nil {
			return
		}
	}
	items, err := freader.Read()
	if err != nil {
		return
	}
	for _, item := range items {
		item = strings.Trim(item, " ")
		matched, err := regexp.MatchString("^[-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?$", item)
		if err != nil {
			return
		}
		if matched {
			features = append(features, NewContinuousFeature(""))
		} else {
			features = append(features, NewDiscreteFeature(""))
		}
	}
	err, maxPrecision := reader.ParseMaxPrecision()
	if err != nil {
		return
	}
	for _, feat := range features {
		if f, ok := feat.(ContinuousFeature); ok {
			f.Precision = maxPrecision
		}
	}
	return
}
func (reader *CSVReader) ParseFeatureName() (err error, names []string) {
	f, err := os.Open(reader.filename)
	if err != nil {
		return
	}
	defer f.Close()
	freader := csv.NewReader(reader.filename)
	names, err = freader.Read()
	if reader.hasHeader {
		for i, n := range names {
			names[i] = strings.TrimSpace(n)
		}
		return
	}
	for i := range names {
		names[i] = fmt.Sprintf("%d",i)
	}
	return names
}
func (reader *CSVReader) ParseMaxPrecision() (err error, precision int) {
	rexp := regexp.MustCompile("[0-9]+(.[0-9]+)?")
	f, err := os.Open(reader.filename)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	maxPrecision := 0
	lineCount := 0
	for scanner.Scan() {
		if lineCount > 10 {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == "$" {
			continue
		}
		if line[0] == "%" {
			continue
		}
		if line[0] == "@" {
			continue
		}
		matches := rexp.FindAllString(line, -1)
		for _, match := range matches {
			splite := strings.Split(match, ".")
			if len(splite) == 2 {
				if p := len(splite[1]); p > maxPrecision {
					maxPrecision = p
				}
			}
		}
		lineCount += 1
	}
	return nil, precision
}
func NewCSVReader(filename string, hasHeader bool) *CSVReader {
	return &CSVReader{
		filename:filename,
		hasHeader:hasHeader,
	}
}
