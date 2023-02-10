package exporter

import "testing"

func TestDataSource_String(t *testing.T) {
	type fields struct {
		Name       string
		VpnFilter  string
		ItemFilter string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataSource := DataSource{
				Name:       tt.fields.Name,
				VpnFilter:  tt.fields.VpnFilter,
				ItemFilter: tt.fields.ItemFilter,
			}
			if got := dataSource.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
