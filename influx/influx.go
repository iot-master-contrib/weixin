package influx

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

var Writer api.WriteAPI
var Reader api.QueryAPI

type Point struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}

var bucket string
var client influxdb2.Client

func Open() {
	client = influxdb2.NewClient(options.Url, options.Token)
	Writer = client.WriteAPI(options.Org, options.Bucket)
	Reader = client.QueryAPI(options.Org)
	bucket = options.Bucket
}

func Close() {
	client.Close()
}

func Insert(measurement, id string, fields map[string]interface{}, ts time.Time) {
	Writer.WritePoint(write.NewPoint(measurement, map[string]string{"id": id}, fields, ts))
}

func Query(measurement, id, field, start, end, win, fn string) ([]Point, error) {
	flux := "from(bucket: \"" + bucket + "\")\n"
	flux += "|> range(start: " + start + ", stop: " + end + ")\n"
	flux += "|> filter(fn: (r) => r[\"_measurement\"] == \"" + measurement + "\")\n"
	flux += "|> filter(fn: (r) => r[\"id\"] == \"" + id + "\")\n"
	flux += "|> filter(fn: (r) => r[\"_field\"] == \"" + field + "\")"
	flux += "|> aggregateWindow(every: " + win + ", fn: " + fn + ", createEmpty: false)\n"
	flux += "|> yield(name: \"" + fn + "\")"

	result, err := Reader.Query(context.Background(), flux)
	if err != nil {
		return nil, err
	}

	records := make([]Point, 0)
	for result.Next() {
		//result.TableChanged()
		records = append(records, Point{
			Value: result.Record().Value(),
			Time:  result.Record().Time(),
		})
	}
	return records, result.Err()
}

func Rate(measurement, id, field, start, end, win string) ([]Point, error) {
	flux := "import \"experimental/aggregate\""
	flux += "from(bucket: \"" + bucket + "\")\n"
	flux += "|> range(start: " + start + ", stop: " + end + ")\n"
	flux += "|> filter(fn: (r) => r[\"_measurement\"] == \"" + measurement + "\")\n"
	//如果是空ID，则全场统计，但是只能做一种模型的统计
	if id != "" {
		flux += "|> filter(fn: (r) => r[\"id\"] == \"" + id + "\")\n"
	}
	flux += "|> filter(fn: (r) => r[\"_field\"] == \"" + field + "\")\n"
	flux += "|> aggregate.rate(every: " + win + ", unit: " + win + ")"

	//log.Println(flux)

	result, err := Reader.Query(context.Background(), flux)
	if err != nil {
		return nil, err
	}

	records := make([]Point, 0)
	for result.Next() {
		//result.TableChanged()
		records = append(records, Point{
			Value: result.Record().Value(),
			Time:  result.Record().Time(),
		})
	}
	return records, result.Err()
}
