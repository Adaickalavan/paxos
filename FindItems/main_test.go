package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	//Save old settings before rewriting settings
	oldArgs := os.Args
	oldOutStream := outStream
	defer func() {
		os.Args = oldArgs
		outStream = oldOutStream
	}()

	//Setup redirection for outputs
	var gotBuf bytes.Buffer
	outStream = &gotBuf

	tests := []struct {
		name    string
		args    []string
		wantBuf []byte
	}{
		{name: "Maximum price matches achievable price",
			args:    []string{"cmd", "prices.txt", "2500", "2"},
			wantBuf: string2Bytes("Candy Bar 500, Earmuffs 2000"),
		},
		{name: "Maximum price higher than achievable price",
			args:    []string{"cmd", "prices.txt", "2300", "2"},
			wantBuf: string2Bytes("Paperback Book 700, Headphones 1400"),
		},
		{name: "Maximum price higher than highest priced item",
			args:    []string{"cmd", "prices.txt", "10000","2"},
			wantBuf: string2Bytes("Earmuffs 2000, Bluetooth Stereo 6000"),
		},
		{name: "Maximum price lower than achievable price",
			args:    []string{"cmd", "prices.txt", "1100", "2"},
			wantBuf: string2Bytes("Not possible"),
		},

		{name: "Maximum price matches achievable price. Select 3 items.",
			args:    []string{"cmd", "prices.txt", "3700", "3"},
			wantBuf: string2Bytes("Paperback Book 700, Detergent 1000, Earmuffs 2000"),
		},
		{name: "Maximum price higher than achievable price. Select 3 items.",
			args:    []string{"cmd", "prices.txt", "4600", "3"},
			wantBuf: string2Bytes("Detergent 1000, Headphones 1400, Earmuffs 2000"),
		},
		{name: "Maximum price higher than highest priced item. Select 3 items.",
			args:    []string{"cmd", "prices.txt", "10000","3"},
			wantBuf: string2Bytes("Headphones 1400, Earmuffs 2000, Bluetooth Stereo 6000"),
		},
		{name: "Maximum price lower than achievable price. Select 3 items.",
			args:    []string{"cmd", "prices.txt", "2100", "3"},
			wantBuf: string2Bytes("Not possible"),
		},

		{name: "Maximum price <= 0.",
		args:    []string{"cmd", "prices.txt", "0","3"},
		wantBuf: string2Bytes("Not possible"),
		},
		{name: "Select <= 0 items.",
			args:    []string{"cmd", "prices.txt", "2100", "0"},
			wantBuf: string2Bytes("Not possible"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			main()
			if !bytes.Equal(gotBuf.Bytes(), tt.wantBuf) {
				t.Errorf("main() = %v, want = %v", gotBuf.String(), string(tt.wantBuf))
			}
		})
		gotBuf.Reset()
	}
}

func string2Bytes(wantBuf string) []byte {
	return bytes.NewBufferString(wantBuf).Bytes()
}
