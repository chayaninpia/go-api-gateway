package cryptox

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	type args struct {
		text     string
		MySecret string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test decrypt",
			args: args{
				text:     "UhGrn6ont1bpjaAhOHsH",
				MySecret: "nkjnj32nof2hee98hneni2oi0923ujcdkjw09f032jocj0w9j9020j2jf",
			},
			want: []byte(`{"test":"test"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.text, tt.args.MySecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
