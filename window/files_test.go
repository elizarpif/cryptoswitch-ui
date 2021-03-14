package window

import "testing"

func Test_fileNameWithoutExtension(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check",
			args: args{
				fileName: "lalala.png.enc",
			},
			want: "lalala.png",
		},
		{
			name: "check",
			args: args{
				fileName: "lalala.png.enc.png.enc",
			},
			want: "lalala.png.enc.png",
		},
		{
			name: "check",
			args: args{
				fileName: "lalala.png.enc.enc",
			},
			want: "lalala.png.enc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileNameWithoutExtension(tt.args.fileName); got != tt.want {
				t.Errorf("fileNameWithoutExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
