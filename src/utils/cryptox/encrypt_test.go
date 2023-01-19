package cryptox

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		request  interface{}
		MySecret string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test encrypt",
			args: args{
				request: map[string]string{
					"test": "test",
				},
				MySecret: "nkjnj32nof2hee98hneni2oi0923ujcdkjw09f032jocj0w9j9020j2jf",
			},
			want: "UhGrn6ont1bpjaAhOHsH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.request, tt.args.MySecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
