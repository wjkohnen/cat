package cat

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	cat := NewCat(context.Background(), os.Stdout)
	n, err := fmt.Fprintln(cat, "Hello, World!")
	if err != nil {
		t.Errorf("Error writing message to cat: %v", err)
	}
	t.Logf("Wrote %d bytes into cat.", n)
	err = cat.Close()
	if err != nil {
		t.Errorf("Error closing cat: %v", err)
	}
}

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cat := NewCat(ctx, os.Stdout)
	n, err := fmt.Fprintln(cat, "Hello, World!")
	if err != nil {
		t.Errorf("Error writing message to cat: %v", err)
	}
	t.Logf("Wrote %d bytes into cat.", n)
	cancel()
	err = cat.Close()
	t.Logf("Error of closing a canceled subprocess was: %v", err)
}
