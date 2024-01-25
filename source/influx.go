package source

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/metorig/metorig/model"
	"time"
)

type Influx struct {
	writeApi api.WriteAPI
}

func NewInflux(token, url, org, bucket string) *Influx {
	cl := influxdb2.NewClient(url, token)

	return &Influx{
		writeApi: cl.WriteAPI(org, bucket),
	}
}

func (s *Influx) Store(m *model.Metrics) error {
	s.writeValue("metronig.mem.usage", m.UsedMem)
	s.writeValue("metronig.mem.free", m.FreeMem)

	return nil
}

func (s *Influx) writeValue(measure string, value any) error {
	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": value,
	}

	point := write.NewPoint(measure, tags, fields, time.Now())
	s.writeApi.WritePoint(point)

	return nil // Temporary maybe
}
