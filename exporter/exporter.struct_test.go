package exporter

import (
	"github.com/go-kit/kit/log"
	"reflect"
	"testing"
)

func TestNewExporter(t *testing.T) {
	type args struct {
		logger     log.Logger
		conf       *Config
		dataSource *[]DataSource
		version    float64
	}
	tests := []struct {
		name string
		args args
		want *Exporter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExporter(tt.args.logger, tt.args.conf, tt.args.dataSource, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
