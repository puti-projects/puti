package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	pwd, _ := os.Getwd()

	type args struct {
		cfg string
	}
	tests := []struct {
		name string
		args args
	}{
		{"cfg", args{cfg: filepath.Join(filepath.Dir(pwd), "../../configs/config.yaml")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.cfg)
		})
	}
}
