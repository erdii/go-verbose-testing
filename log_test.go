package goverbosetesting

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestLogHiddenNoFail(t *testing.T) {
	t.Log("TestLogHiddenNoFail: hidden")
}

func TestLogVisibleFail(t *testing.T) {
	t.Log("TestLogVisibleFail: visible")
	t.Fail()
}

func TestFmtPrintlnHiddenNoFail(t *testing.T) {
	fmt.Println("TestFmtPrintlnHiddenNoFail: hidden")
}

func TestFmtPrintlnVisibleFail(t *testing.T) {
	fmt.Println("TestFmtPrintlnVisibleFail: visible")
	t.Fail()
}

func TestSlogHiddenNoFail(t *testing.T) {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	log.Info("TestSlogHiddenNoFail: hidden")
}

func TestSlogVisibleFail(t *testing.T) {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	log.Info("TestSlogVisibleFail: visible")
	t.Fail()
}
