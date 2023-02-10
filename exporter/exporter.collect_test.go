package exporter

import (
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/solacedev/pubsubplus-prometheus-exporter/semp"
	"testing"
)

func TestExporter_Collect(t *testing.T) {
	type fields struct {
		config    *Config
		logger    log.Logger
		lastError error
		semp      *semp.Semp
	}
	type args struct {
		_ chan<- prometheus.Metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Exporter{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				lastError: tt.fields.lastError,
				semp:      tt.fields.semp,
			}
		})
	}
}
