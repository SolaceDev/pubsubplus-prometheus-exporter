package semp

import (
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestNewSemp(t *testing.T) {
	type args struct {
		logger             log.Logger
		brokerURI          string
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		exporterVersion    float64
	}
	tests := []struct {
		name string
		args args
		want *Semp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSemp(tt.args.logger, tt.args.brokerURI, tt.args.httpClient, tt.args.httpRequestVisitor, tt.args.exporterVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSemp_GetBridgeSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetBridgeSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBridgeSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetBridgeSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetBridgeStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetBridgeStatsSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBridgeStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetBridgeStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetClientMessageSpoolStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetClientMessageSpoolStatsSemp1(tt.args.ch, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientMessageSpoolStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetClientMessageSpoolStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetClientSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetClientSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetClientSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetClientSlowSubscriberSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetClientSlowSubscriberSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientSlowSubscriberSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetClientSlowSubscriberSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetClientStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetClientStatsSemp1(tt.args.ch, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetClientStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetClusterLinksSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch            chan<- prometheus.Metric
		clusterFilter string
		linkFilter    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetClusterLinksSemp1(tt.args.ch, tt.args.clusterFilter, tt.args.linkFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClusterLinksSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetClusterLinksSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetConfigSyncRouterSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetConfigSyncRouterSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigSyncRouterSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetConfigSyncRouterSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetConfigSyncVpnSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch        chan<- prometheus.Metric
		vpnFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetConfigSyncVpnSemp1(tt.args.ch, tt.args.vpnFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigSyncVpnSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetConfigSyncVpnSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetDiskSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetDiskSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiskSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetDiskSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetGlobalStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetGlobalStatsSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGlobalStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetGlobalStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetHealthSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetHealthSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHealthSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetHealthSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetInterfaceSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch              chan<- prometheus.Metric
		interfaceFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetInterfaceSemp1(tt.args.ch, tt.args.interfaceFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInterfaceSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetInterfaceSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetMemorySemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetMemorySemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemorySemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetMemorySemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetQueueDetailsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetQueueDetailsSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueueDetailsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetQueueDetailsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetQueueRatesSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetQueueRatesSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueueRatesSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetQueueRatesSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetQueueStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetQueueStatsSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueueStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetQueueStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetRedundancySemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetRedundancySemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRedundancySemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetRedundancySemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetReplicationStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetReplicationStatsSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReplicationStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetReplicationStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetSpoolSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetSpoolSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpoolSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetSpoolSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetStorageElementSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch                   chan<- prometheus.Metric
		storageElementFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetStorageElementSemp1(tt.args.ch, tt.args.storageElementFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStorageElementSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetStorageElementSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetTopicEndpointDetailsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetTopicEndpointDetailsSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopicEndpointDetailsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetTopicEndpointDetailsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetTopicEndpointRatesSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetTopicEndpointRatesSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopicEndpointRatesSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetTopicEndpointRatesSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetTopicEndpointStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch         chan<- prometheus.Metric
		vpnFilter  string
		itemFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetTopicEndpointStatsSemp1(tt.args.ch, tt.args.vpnFilter, tt.args.itemFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopicEndpointStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetTopicEndpointStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetVersionSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetVersionSemp1(tt.args.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersionSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetVersionSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetVpnReplicationSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch        chan<- prometheus.Metric
		vpnFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetVpnReplicationSemp1(tt.args.ch, tt.args.vpnFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVpnReplicationSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetVpnReplicationSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetVpnSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch        chan<- prometheus.Metric
		vpnFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetVpnSemp1(tt.args.ch, tt.args.vpnFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVpnSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetVpnSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetVpnSpoolSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch        chan<- prometheus.Metric
		vpnFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetVpnSpoolSemp1(tt.args.ch, tt.args.vpnFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVpnSpoolSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetVpnSpoolSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_GetVpnStatsSemp1(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		ch        chan<- prometheus.Metric
		vpnFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOk  float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			gotOk, err := e.GetVpnStatsSemp1(tt.args.ch, tt.args.vpnFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVpnStatsSemp1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetVpnStatsSemp1() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSemp_postHTTP(t *testing.T) {
	type fields struct {
		logger             log.Logger
		httpClient         http.Client
		httpRequestVisitor func(*http.Request)
		brokerURI          string
		exporterVersion    float64
	}
	type args struct {
		uri  string
		in1  string
		body string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Semp{
				logger:             tt.fields.logger,
				httpClient:         tt.fields.httpClient,
				httpRequestVisitor: tt.fields.httpRequestVisitor,
				brokerURI:          tt.fields.brokerURI,
				exporterVersion:    tt.fields.exporterVersion,
			}
			got, err := s.postHTTP(tt.args.uri, tt.args.in1, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("postHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postHTTP() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeMetricBool(t *testing.T) {
	type args struct {
		item bool
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeMetricBool(tt.args.item); got != tt.want {
				t.Errorf("encodeMetricBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeMetricMulti(t *testing.T) {
	type args struct {
		item     string
		refItems []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeMetricMulti(tt.args.item, tt.args.refItems); got != tt.want {
				t.Errorf("encodeMetricMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
