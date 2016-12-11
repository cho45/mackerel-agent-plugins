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

	common := []string{
		"status-errors.400",
		"status-errors.403",
		"status-errors.404",
		"status-errors.405",
		"status-errors.416",
		"status-errors.417",
		"status-errors.500",
		"status-errors.502",
		"status-errors.503",
		"http2-errors.protocol",
		"http2-errors.internal",
		"http2-errors.flow-control",
		"http2-errors.settings-timeout",
		"http2-errors.stream-closed",
		"http2-errors.frame-size",
		"http2-errors.refused-stream",
		"http2-errors.cancel",
		"http2-errors.compression",
		"http2-errors.connect",
		"http2-errors.enhance-your-calm",
		"http2-errors.inadequate-security",
		"http2.read-closed",
		"http2.write-closed",
	}
	for _, k := range common {
		v, ok := s[k].(float64)
		if ok {
			stat[k] = v
		}
	}

	durations := []string{
		"connect-time",
		"header-time",
		"body-time",
		"request-total-time",
		"process-time",
		"response-time",
		"duration",
	}
	ranges := []string {
		"0", "25", "50", "75", "99",
	}
	for _, k := range durations {
		for _, r := range ranges {
			v, ok := s[k + "-" + r].(float64)
			if ok {
				stat[k + "." + r] = v
			}
		}
	}

	return stat, nil
}

func (p H2OPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.prefix)

	return map[string]mp.Graphs{
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
		"connect-time": {
			Label: (labelPrefix + " Connect Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "connect-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "connect-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "connect-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "connect-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "connect-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"header-time": {
			Label: (labelPrefix + " Header Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "header-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "header-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "header-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "header-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "header-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"body-time": {
			Label: (labelPrefix + " Body Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "body-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "body-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "body-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "body-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "body-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"request-total-time": {
			Label: (labelPrefix + " Request Time Total Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "request-total-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "request-total-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "request-total-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "request-total-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "request-total-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"process-time": {
			Label: (labelPrefix + " Process Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "process-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "process-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "process-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "process-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "process-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"response-time": {
			Label: (labelPrefix + " Response Time"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "response-time.0", Label: "0", Diff: false, Stacked: true},
				{Name: "response-time.25", Label: "25", Diff: false, Stacked: true},
				{Name: "response-time.50", Label: "50", Diff: false, Stacked: true},
				{Name: "response-time.75", Label: "75", Diff: false, Stacked: true},
				{Name: "response-time.99", Label: "99", Diff: false, Stacked: true},
			},
		},
		"duration": {
			Label: (labelPrefix + " Duration"),
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "duration.0", Label: "0", Diff: false, Stacked: true},
				{Name: "duration.25", Label: "25", Diff: false, Stacked: true},
				{Name: "duration.50", Label: "50", Diff: false, Stacked: true},
				{Name: "duration.75", Label: "75", Diff: false, Stacked: true},
				{Name: "duration.99", Label: "99", Diff: false, Stacked: true},
			},
		},
	}
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
	helper := mp.NewMackerelPlugin(h2o)
	helper.Tempfile = *optTempfile

	helper.Run()
}
