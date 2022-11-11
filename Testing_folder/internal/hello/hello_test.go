package hello_test

import (
    "bytes"
    "testing"
    "randomlkjlskj/internal/hello"
)

func TestSample(t *testing.T) {
    expected_string:="Hello, Xebia"
    faketerminal:=&bytes.Buffer{}
    printer:=hello.Printer{Output: faketerminal}
    printer.Print("Hello, Xebia")
    if faketerminal.String()!=expected_string{
        t.Errorf("Expected String(%s) is not same as"+
            " actual string (%s)", expected_string,faketerminal.String())
    }
}