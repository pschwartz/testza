package testza

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestCaptureHelper_Stdout(t *testing.T) {
	type args struct {
		capture func(w io.Writer) error
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{capture: func(w io.Writer) error { fmt.Print("Hello, World!"); return nil }}, want: "Hello, World!"},
		{args: args{capture: func(w io.Writer) error { fmt.Print(" Hello, World! "); return nil }}, want: " Hello, World! "},
		{args: args{capture: func(w io.Writer) error { fmt.Print("H\ne\nl\nl\nl\no\n\n\n\nWorld!"); return nil }}, want: "H\ne\nl\nl\nl\no\n\n\n\nWorld!"},
		{args: args{capture: func(w io.Writer) error { fmt.Println("Hello, World!"); return nil }}, want: "Hello, World!\n"},
		{args: args{capture: func(w io.Writer) error { fmt.Println("Hello, \nWorld!"); return nil }}, want: "Hello, \nWorld!\n"},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Use.Capture.Stdout(tt.args.capture)
			Use.Assert.Nil(t, err)
			Use.Assert.Equal(t, got, tt.want)
		})
	}
}

func TestCaptureHelper_Stderr(t *testing.T) {
	type args struct {
		capture func(w io.Writer) error
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{capture: func(w io.Writer) error { fmt.Fprint(os.Stderr, "Hello, World!"); return nil }}, want: "Hello, World!"},
		{args: args{capture: func(w io.Writer) error { fmt.Fprint(os.Stderr, " Hello, World! "); return nil }}, want: " Hello, World! "},
		{args: args{capture: func(w io.Writer) error { fmt.Fprint(os.Stderr, "H\ne\nl\nl\nl\no\n\n\n\nWorld!"); return nil }}, want: "H\ne\nl\nl\nl\no\n\n\n\nWorld!"},
		{args: args{capture: func(w io.Writer) error { fmt.Fprintln(os.Stderr, "Hello, World!"); return nil }}, want: "Hello, World!\n"},
		{args: args{capture: func(w io.Writer) error { fmt.Fprintln(os.Stderr, "Hello, \nWorld!"); return nil }}, want: "Hello, \nWorld!\n"},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Use.Capture.Stderr(tt.args.capture)
			Use.Assert.Nil(t, err)
			Use.Assert.Equal(t, got, tt.want)
		})
	}
}
