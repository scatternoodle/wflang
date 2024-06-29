package jrpc2

import (
	"reflect"
	"testing"
)

func TestEncodeMessage(t *testing.T) {
	input := NewNotification("put")

	msg, err := EncodeMessage(input)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	want := []byte("Content-Length: 32\r\n\r\n{\"jsonrpc\":\"2.0\",\"method\":\"put\"}")
	if !reflect.DeepEqual(msg, want) {
		t.Errorf("msg = %s, want %s", msg, want)
	}
}

func TestDecodeMessage(t *testing.T) {
	type wantT struct {
		method  string
		content string
		err     bool
	}

	tests := []struct {
		name string
		msg  []byte
		want wantT
	}{
		{
			name: "good",
			msg:  []byte("Content-Length: 31\r\n\r\n{\"method\": \"put\", \"test\": \"hi\"}"),
			want: wantT{
				method:  "put",
				content: `{"method": "put", "test": "hi"}`,
				err:     false,
			},
		},
		{
			name: "bad incorrect content length",
			msg:  []byte("Content-Length: 30\r\n\r\n{\"method\": \"put\", \"test\": \"hi\"}"),
			want: wantT{
				method:  "",
				content: "",
				err:     true,
			},
		},
		{
			name: "bad missing header",
			msg:  []byte("Content-Length: 31\r\n\r\n{\"AAAHH!\": \"put\", \"test\": \"hi\"}"),
			want: wantT{
				method:  "",
				content: "",
				err:     true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, content, err := DecodeMessage(tt.msg)

			if err != nil && !tt.want.err {
				t.Fatalf("unexpected error: %+v", err)
			}
			if err != nil && tt.want.err {
				return // no need to test outputs
			}
			if err == nil && tt.want.err {
				t.Fatal("err = nil, want err")
			}
			if method != tt.want.method {
				t.Errorf("method = %s, want %s", method, tt.want.method)
			}
			if string(content) != tt.want.content {
				t.Errorf("content = %s, want %s", string(content), tt.want.content)
			}
		})
	}
}
