package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "universe") {
		t.Errorf("Expected to find universe, but it is not there")
	}

	if !strings.Contains(output, "world") {
		t.Errorf("Expected to find world, but it is not there")
	}

	if !strings.Contains(output, "cosmos") {
		t.Errorf("Expected to find cosmos, but it is not there")
	}
}

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("epsilon", &wg)
	wg.Wait()

	if msg != "epsilon" {
		t.Error("The message was not updated correctly")
	}

}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)
	go updateMessage("epsilon", &wg)
	wg.Wait()

	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}
