package sample

import (
	"context"
	"testing"
)

func TestSample_GetCursor(t *testing.T) {
	var sample = &Sample{}
	result, err := sample.GetCursor(context.Background(), 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
