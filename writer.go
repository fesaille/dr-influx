package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api/write"
)

var token string

type measurement struct {
	temperature float64
	disk_free   float64
	disk_total  float64
	mem_total   float64
	mem_free    uint64
}

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	// token := flag.String("token", "", "Token provided by Influx")
	bucket := flag.String("bucket", "MYBUCKET", "Bucket name")
	org := flag.String("org", "MYORG", "Org name")

	// client := influxdb2.NewClient("http://localhost:8086", token)
	client := influxdb2.NewClientWithOptions("http://localhost:8086", token,
		influxdb2.DefaultOptions().SetBatchSize(50))
	// always close client at the end
	defer client.Close()

	// get non-blocking write client
	writeAPI := client.WriteAPI(*org, *bucket)

	for i := 0; i < 10000; i++ {
		// create point
		// write asynchronously
		p := newPoint(i)
		writeAPI.WritePoint(p)
	}
	// Force all unwritten data to be sent
	writeAPI.Flush()

	// var isinterface bool
	// meas := measurement{1, 1, 1, 1, 1}

	// err := types.IsInterface(meas)
	// if err != nill {
	// 	log.Fatal(err)
	// }

}

func newPoint(i int) *write.Point {

	p := influxdb2.NewPoint(
		"system",
		map[string]string{
			"id":       fmt.Sprintf("rack_%v", i%10),
			"vendor":   "AWS",
			"hostname": fmt.Sprintf("host_%v", i%100),
		},
		map[string]interface{}{
			"temperature": rand.Float64() * 80.0,
			"disk_free":   rand.Float64() * 1000.0,
			"disk_total":  (i/10 + 1) * 1000000,
			"mem_total":   (i/100 + 1) * 10000000,
			"mem_free":    rand.Uint64(),
		},
		time.Now())

	return p

}
