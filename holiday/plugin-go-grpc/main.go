package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/holiday"
)

// US is a real implementation of Holiday
type US struct {
	dates *treeset.Set
}

func (us US) IsHoliday(d time.Time) (bool, error) {
	return us.dates.Contains(d), nil
}

func (us US) Delta(d1, d2 time.Time) (int32, error) {
	return int32(us.dates.Select(func(i int, v interface{}) bool {
		d := v.(time.Time)
		return d.After(d1) && d.Before(d2)
	}).Size()), nil
}

func main() {
	us := &US{dates: treeset.NewWith(byTime)}
	_, err := us.parseDatesFromCSV("/Users/mmagaldi/go/src/github.com/magaldima/bizday/holiday/plugin-go-grpc/us-holidays.csv")
	if err != nil {
		panic(err)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: holiday.Handshake,
		Plugins: map[string]plugin.Plugin{
			"US": &holiday.HolidayPlugin{Impl: us},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

func (us US) parseDatesFromCSV(filename string) ([]time.Time, error) {
	const dateFormat = "2006-01-02"
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var dates []time.Time
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		t, err := time.Parse(dateFormat, line[0])
		if err != nil {
			return dates, err
		}
		us.dates.Add(t)
		dates = append(dates, t)
	}
	return dates, nil
}

func byTime(a, b interface{}) int {
	c1 := a.(time.Time)
	c2 := b.(time.Time)

	switch {
	case c1.After(c2):
		return 1
	case c1.Before(c2):
		return -1
	default:
		return 0
	}
}
