package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitConfig(t *testing.T) {
	pwd, _ := os.Getwd()

	type args struct {
		cfg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"testCfg",
			args{cfg: filepath.Join(filepath.Dir(pwd), "../../configs/config.yaml")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
