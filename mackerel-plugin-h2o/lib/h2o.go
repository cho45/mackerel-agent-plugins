package mph2o

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

type H2OPlugin struct {
	URI string
	prefix      string
	withDurations bool
}

func (p H2OPlugin) FetchMetrics() (map[string]interface{}, error) {
	resp, err := http.Get(p.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return p.parseStats(resp.Body)
}

func (p H2OPlugin) parseStats(body io.Reader) (map[string]interface{}, error) {
	stat := make(map[string]interface{})
	decoder := json.NewDecoder(body)

	var s map[string]interface{}
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}

	for k, raw := range s {
		v, ok := raw.(float64)
		if ok {
			stat[k] = v
		}
	}

	return stat, nil
}

func (p H2OPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.prefix)

	base := map[string]mp.Graphs{
		"status-errors": {
			Label: (labelPrefix + " Status Errors"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "status-errors.400", Label: "400", Diff: false},
				{Name: "status-errors.403", Label: "403", Diff: false},
				{Name: "status-errors.404", Label: "404", Diff: false},
				{Name: "status-errors.405", Label: "405", Diff: false},
				{Name: "status-errors.416", Label: "416", Diff: false},
				{Name: "status-errors.417", Label: "417", Diff: false},
				{Name: "status-errors.500", Label: "500", Diff: false},
				{Name: "status-errors.502", Label: "502", Diff: false},
				{Name: "status-errors.503", Label: "503", Diff: false},
			},
		},
		"http2-errors": {
			Label: (labelPrefix + " HTTP2 Errors"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "http2-errors.protocol", Label: "protocol", Diff: false},
				{Name: "http2-errors.internal", Label: "internal", Diff: false},
				{Name: "http2-errors.flow-control", Label: "flow-control", Diff: false},
				{Name: "http2-errors.settings-timeout", Label: "settings-timeout", Diff: false},
				{Name: "http2-errors.stream-closed", Label: "stream-closed", Diff: false},
				{Name: "http2-errors.frame-size", Label: "frame-size", Diff: false},
				{Name: "http2-errors.refused-stream", Label: "refused-stream", Diff: false},
				{Name: "http2-errors.cancel", Label: "cancel", Diff: false},
				{Name: "http2-errors.compression", Label: "compression", Diff: false},
				{Name: "http2-errors.connect", Label: "connect", Diff: false},
				{Name: "http2-errors.enhance-your-calm", Label: "enhance-your-calm", Diff: false},
				{Name: "http2-errors.inadequate-security", Label: "inadequate-security", Diff: false},
			},
		},
		"http2": {
			Label: (labelPrefix + " HTTP2"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "http2.read-closed", Label: "read-closed", Diff: false},
				{Name: "http2.write-closed", Label: "write-closed", Diff: false},
			},
		},
	}

	durations := map[string]mp.Graphs{
		"duration-0": {
			Label: (labelPrefix + " Duration 0"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time-0", Label: "connect-time", Diff: false, Stacked: false},
				{Name: "header-time-0", Label: "header-time", Diff: false, Stacked: false},
				{Name: "body-time-0", Label: "body-time", Diff: false, Stacked: false},
				{Name: "request-total-time-0", Label: "request-total-time", Diff: false, Stacked: false},
				{Name: "process-time-0", Label: "process-time", Diff: false, Stacked: false},
				{Name: "response-time-0", Label: "response-time", Diff: false, Stacked: false},
				{Name: "duration-0", Label: "duration", Diff: false, Stacked: false},
			},
		},
		"duration-25": {
			Label: (labelPrefix + " Duration 25"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time-25", Label: "connect-time", Diff: false, Stacked: false},
				{Name: "header-time-25", Label: "header-time", Diff: false, Stacked: false},
				{Name: "body-time-25", Label: "body-time", Diff: false, Stacked: false},
				{Name: "request-total-time-25", Label: "request-total-time", Diff: false, Stacked: false},
				{Name: "process-time-25", Label: "process-time", Diff: false, Stacked: false},
				{Name: "response-time-25", Label: "response-time", Diff: false, Stacked: false},
				{Name: "duration-25", Label: "duration", Diff: false, Stacked: false},
			},
		},
		"duration-50": {
			Label: (labelPrefix + " Duration 50"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time-50", Label: "connect-time", Diff: false, Stacked: false},
				{Name: "header-time-50", Label: "header-time", Diff: false, Stacked: false},
				{Name: "body-time-50", Label: "body-time", Diff: false, Stacked: false},
				{Name: "request-total-time-50", Label: "request-total-time", Diff: false, Stacked: false},
				{Name: "process-time-50", Label: "process-time", Diff: false, Stacked: false},
				{Name: "response-time-50", Label: "response-time", Diff: false, Stacked: false},
				{Name: "duration-50", Label: "duration", Diff: false, Stacked: false},
			},
		},
		"duration-75": {
			Label: (labelPrefix + " Duration 75"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time-75", Label: "connect-time", Diff: false, Stacked: false},
				{Name: "header-time-75", Label: "header-time", Diff: false, Stacked: false},
				{Name: "body-time-75", Label: "body-time", Diff: false, Stacked: false},
				{Name: "request-total-time-75", Label: "request-total-time", Diff: false, Stacked: false},
				{Name: "process-time-75", Label: "process-time", Diff: false, Stacked: false},
				{Name: "response-time-75", Label: "response-time", Diff: false, Stacked: false},
				{Name: "duration-75", Label: "duration", Diff: false, Stacked: false},
			},
		},
		"duration-99": {
			Label: (labelPrefix + " Duration 99"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time-99", Label: "connect-time", Diff: false, Stacked: false},
				{Name: "header-time-99", Label: "header-time", Diff: false, Stacked: false},
				{Name: "body-time-99", Label: "body-time", Diff: false, Stacked: false},
				{Name: "request-total-time-99", Label: "request-total-time", Diff: false, Stacked: false},
				{Name: "process-time-99", Label: "process-time", Diff: false, Stacked: false},
				{Name: "response-time-99", Label: "response-time", Diff: false, Stacked: false},
				{Name: "duration-99", Label: "duration", Diff: false, Stacked: false},
			},
		},
	}

	if p.withDurations {
		for k, v := range durations {
			base[k] = v
		}
	}
	return base
}

func (p H2OPlugin) MetricKeyPrefix() string {
	if p.prefix == "" {
		return "h2o"
	} else {
		return p.prefix
	}
}

func Do() {
	optHost := flag.String("host", "localhost", "Hostname")
	optScheme := flag.String("scheme", "http", "Scheme")
	optPort := flag.String("port", "80", "Port")
	optPath := flag.String("path", "/server-status/json", "Path")
	optBasicAuth := flag.String("basic-auth", "", "BasicAuth")
	optWithDurations := flag.Bool("with-durations", false, "With durations (requires duration-stats: ON setting on h2o.conf)")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	optMetricKeyPrefix := flag.String("metric-key-prefix", "h2o", "Metric Key Prefix")
	flag.Parse()

	var h2o H2OPlugin

	if *optBasicAuth != "" {
		h2o.URI = fmt.Sprintf("%s://%s@%s:%s%s", *optScheme, *optBasicAuth, *optHost, *optPort, *optPath)
	} else {
		h2o.URI = fmt.Sprintf("%s://%s:%s%s", *optScheme, *optHost, *optPort, *optPath)
	}

	h2o.prefix = *optMetricKeyPrefix
	h2o.withDurations = *optWithDurations
	helper := mp.NewMackerelPlugin(h2o)
	helper.Tempfile = *optTempfile

	helper.Run()
}
