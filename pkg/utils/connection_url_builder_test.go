package utils

import (
	"fmt"
	"net/url"
	"qontak_integration/pkg/configs"
	"testing"
)

func TestConnectionURLBuilder(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Success",
			args{n: "postgres"},
			"",
			false,
		},
		{
			"No url",
			args{n: ""},
			"",
			true,
		},
		{
			"Redis",
			args{n: "redis"},
			configs.Config.Redis.Address,
			false,
		},
		{
			"MysSql",
			args{n: "mysql"},
			fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
				configs.Config.Database.Username,
				configs.Config.Database.Password,
				configs.Config.Database.Host,
				configs.Config.Database.Port,
				configs.Config.Database.Schema,
				url.QueryEscape("Asia/Jakarta")),
			false,
		},
		{
			"Fiber",
			args{n: "fiber"},
			fmt.Sprintf(
				"%s:%d",
				"0.0.0.0",
				configs.Config.Apps.HttpPort,
			),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectionURLBuilder(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectionURLBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConnectionURLBuilder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
