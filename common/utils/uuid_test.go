package utils

import (
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
)

func TestUUIDMap(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "uuid-test-1",
			args: args{
				str: "82410302-039e-41b6-98b0-d964084b4170",
			},
			want:    uuid.FromStringOrNil("82410302-039e-41b6-98b0-d964084b4170"),
			wantErr: false,
		},
		{
			name: "uuid-test-2",
			args: args{
				str: "88c502e6-d7eb-4c8e-8259-94cb13d83c77",
			},
			want:    uuid.FromStringOrNil("88c502e6-d7eb-4c8e-8259-94cb13d83c77"),
			wantErr: false,
		},
		{
			name: "uuid-map-1",
			args: args{
				str: "123456",
			},
			want:    uuid.FromStringOrNil("f8598425-92f2-5508-a071-4fc67f9040ac"),
			wantErr: false,
		},
		{
			name: "uuid-map-2",
			args: args{
				str: "a9dk23bz0",
			},
			want:    uuid.FromStringOrNil("c91481b6-fc0f-5d9e-b166-5ddf07b9c3c5"),
			wantErr: false,
		},
		{
			name: "uuid-map-2",
			args: args{
				str: "中文123",
			},
			want:    uuid.FromStringOrNil("145c544c-2229-59e5-8dbb-3f33b7610d26"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UUIDMap(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UUIDMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUIDMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
