package main

import (
	"strings"
	"testing"
)

func TestTemplateSupportsOptionalErrorDataWithoutChangingIError(t *testing.T) {
	svc := &service{Name: "DemoService"}
	got := svc.execute()

	if !strings.Contains(got, "type iErrorData interface") {
		t.Fatalf("generated template should declare optional iErrorData interface")
	}
	if !strings.Contains(got, "GetData() interface{}") {
		t.Fatalf("generated template should read GetData from optional iErrorData interface")
	}
	if !strings.Contains(got, "if d, ok := e.(iErrorData); ok") {
		t.Fatalf("generated template should read GetData from the matched public iError only")
	}
	if !strings.Contains(got, "data = d.GetData()") {
		t.Fatalf("generated template should use provider data as response data")
	}
	if !strings.Contains(got, "resp.response(ctx, status, code, msg, data)") {
		t.Fatalf("generated template should pass optional error data to the response envelope")
	}
	iErrorDataStart := strings.Index(got, "type iErrorData interface")
	iErrorMatchStart := strings.Index(got, "if errors.As(err, &e) {")
	if iErrorDataStart < 0 || iErrorMatchStart < 0 || iErrorDataStart < iErrorMatchStart {
		t.Fatalf("generated template should only read optional error data after matching public iError")
	}
	if strings.Contains(got, "GetAct") || strings.Contains(got, `body["act"]`) {
		t.Fatalf("generated template must not include login-specific or top-level act support")
	}

	iErrorStart := strings.Index(got, "type iError interface")
	if iErrorStart < 0 {
		t.Fatalf("generated template should still declare iError interface")
	}
	iErrorEnd := strings.Index(got[iErrorStart:], "}")
	if iErrorEnd < 0 {
		t.Fatalf("generated iError interface should have a closing brace")
	}
	iErrorBlock := got[iErrorStart : iErrorStart+iErrorEnd]
	if strings.Contains(iErrorBlock, "GetData") {
		t.Fatalf("generated iError interface must not require GetData")
	}
}
