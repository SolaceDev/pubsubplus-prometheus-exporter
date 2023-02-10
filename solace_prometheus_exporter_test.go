package main

import (
	"github.com/go-kit/kit/log"
	"net/http"
	"github.com/solacedev/pubsubplus-prometheus-exporter/exporter"
	"testing"
)

func Test_doHandle(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		dataSource []exporter.DataSource
		conf       *exporter.Config
		logger     log.Logger
	}
	tests := []struct {
		name           string
		args           args
		wantResultCode string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResultCode := doHandle(tt.args.w, tt.args.r, tt.args.dataSource, tt.args.conf, tt.args.logger); gotResultCode != tt.wantResultCode {
				t.Errorf("doHandle() = %v, want %v", gotResultCode, tt.wantResultCode)
			}
		})
	}
}

func Test_logDataSource(t *testing.T) {
	type args struct {
		dataSources []exporter.DataSource
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logDataSource(tt.args.dataSources); got != tt.want {
				t.Errorf("logDataSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
