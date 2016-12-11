package mph2o

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestParse(t *testing.T) {
	var h2o H2OPlugin
	json := `
{
 
 "server-version": "2.1.0-DEV",
 "openssl-version": "LibreSSL 2.4.4",
 "current-time": "11/Dec/2016:17:12:07 +0900",
 "restart-time": "11/Dec/2016:14:29:06 +0900",
 "uptime": 9781,
 "generation": 11,
 "connections": 2,
 "max-connections": 1024,
 "listeners": 2,
 "worker-threads": 3,
 "num-sessions": 1446,
 "requests": [
 ],
 "status-errors.400": 0,
 "status-errors.403": 0,
 "status-errors.404": 6,
 "status-errors.405": 0,
 "status-errors.416": 0,
 "status-errors.417": 0,
 "status-errors.500": 0,
 "status-errors.502": 0,
 "status-errors.503": 0,
 "http2-errors.protocol": 0, 
 "http2-errors.internal": 0, 
 "http2-errors.flow-control": 0, 
 "http2-errors.settings-timeout": 0, 
 "http2-errors.stream-closed": 0, 
 "http2-errors.frame-size": 0, 
 "http2-errors.refused-stream": 0, 
 "http2-errors.cancel": 0, 
 "http2-errors.compression": 0, 
 "http2-errors.connect": 0, 
 "http2-errors.enhance-your-calm": 0, 
 "http2-errors.inadequate-security": 0, 
 "http2.read-closed": 9, 
 "http2.write-closed": 0
,
 "connect-time-0": 32772,
 "connect-time-25": 415118,
 "connect-time-50": 2609807,
 "connect-time-75": 6262158,
 "connect-time-99": 10042828
, "header-time-0": 0,
 "header-time-25": 0,
 "header-time-50": 0,
 "header-time-75": 0,
 "header-time-99": 0
, "body-time-0": 0,
 "body-time-25": 0,
 "body-time-50": 0,
 "body-time-75": 0,
 "body-time-99": 0
, "request-total-time-0": 0,
 "request-total-time-25": 0,
 "request-total-time-50": 0,
 "request-total-time-75": 0,
 "request-total-time-99": 0
, "process-time-0": 0,
 "process-time-25": 0,
 "process-time-50": 0,
 "process-time-75": 0,
 "process-time-99": 151195
, "response-time-0": 0,
 "response-time-25": 0,
 "response-time-50": 0,
 "response-time-75": 0,
 "response-time-99": 704009
, "duration-0": 0,
 "duration-25": 0,
 "duration-50": 0,
 "duration-75": 0,
 "duration-99": 753367
}
	`

	stat, err := h2o.parseStats(bytes.NewBufferString(json))
	fmt.Println(stat)
	assert.Nil(t, err)
	assert.EqualValues(t, reflect.TypeOf(stat["status-errors.400"]).String(), "float64")
	assert.EqualValues(t, reflect.TypeOf(stat["duration-99"]).String(), "float64")
	assert.EqualValues(t, stat["duration-99"], 753367)
}
