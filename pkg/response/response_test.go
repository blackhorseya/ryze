package response

import (
	"reflect"
	"testing"
)

func TestResponse_WithData(t *testing.T) {
	type fields struct {
		Code    int
		Message string
		Data    interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		{
			name: "with data",
			fields: fields{
				Code:    200,
				Message: "ok",
				Data:    nil,
			},
			args: args{
				data: "data",
			},
			want: &Response{
				Code:    200,
				Message: "ok",
				Data:    "data",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Data:    tt.fields.Data,
			}
			if got := resp.WithData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithData() = %v, want %v", got, tt.want)
			}
		})
	}
}
