package main

import (
	"context"
	"fmt"

	api "github.com/magaldima/bizday/api"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewDateCalcClient(conn)

	binaryDateRequest := api.BinaryDateRequest{
		Cal: &api.HolidayCalendar{
			Name: "us-holiday",
		},
		Start: &api.Date{
			Year:  2018,
			Month: api.Month_May,
			Day:   10,
		},
		End: &api.Date{
			Year:  2018,
			Month: api.Month_August,
			Day:   5,
		},
	}
	resp, err := client.HolidaysBetween(context.Background(), &binaryDateRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("number of holidays is %v\n", resp.Days)
}
