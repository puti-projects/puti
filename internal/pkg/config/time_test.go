package config

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeLoc(t *testing.T) {
	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	tests := []struct {
		name string
		want *time.Location
	}{
		{"shanghai", shanghai},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeLoc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
