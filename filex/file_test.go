package filex

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	file, err := os.Open("eqps")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := ReadLines(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", lines)
}

func TestMakeDir1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test make dir1",
			args: args{
				name: "dir1",
			},
			wantErr: false,
		},
		{
			name: "test make dir1 again",
			args: args{
				name: "dir1",
			},
			wantErr: false,
		},
		{
			name: "test make dir2/dir22",
			args: args{
				name: "dir2/dir22",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeDir(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("MakeDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClearDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test clear dir1",
			args: args{
				name: "dir1",
			},
			wantErr: false,
		},
		{
			name: "test clear dir1 again",
			args: args{
				name: "dir1",
			},
			wantErr: false,
		},
		{
			name: "test clear dir2/dir22",
			args: args{
				name: "dir2/dir22",
			},
			wantErr: false,
		},
		{
			name: "test clear dir2/dir22 again",
			args: args{
				name: "dir2/dir22",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ClearDir(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ClearDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{
			name: "test create file1",
			args: args{
				name: "file1",
			},
			wantErr: false,
		},
		{
			name: "test create file1 again",
			args: args{
				name: "file1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateFile(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateFile() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestOpenFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{
			name: "test open file1",
			args: args{
				name: "file1",
			},
			wantErr: false,
		},
		{
			name: "test open file2",
			args: args{
				name: "file2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenFile(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("OpenFile() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
