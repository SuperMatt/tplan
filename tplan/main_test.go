package main

import (
	"fmt"
	"testing"
)

func TestLoadPlan(t *testing.T) {
	_, err := LoadPlan("../testdata/badplan/tf_plan")
	if err == nil {
		t.Error("Loaded bad plan. This shouldn't happen")
	}

	_, err = LoadPlan("../testdata/change/tf_plan")
	if err != nil {
		t.Error(err)
	}
}

func TestDiff(t *testing.T) {
	p, err := Diff("../testdata/change/tf_plan")
	if err != nil {
		t.Error(err)
	}
	if p == false {
		t.Error("Diff was false, when it should be true")
	}

	p, err = Diff("../testdata/nochange/tf_plan")
	if err != nil {
		t.Error(err)
	}
	if p {
		t.Error("Diff was true, when it should be false")
	}
}

func TestGet(t *testing.T) {
	p, err := Get("../testdata/change/tf_plan")
	if err != nil {
		t.Error(err)
	}
	if p.Diff.Empty() {
		t.Error("Plan is empty when it should contain a change")
	}

	p, err = Get("../testdata/nochange/tf_plan")
	if err != nil {
		t.Error(err)
	}
	if !p.Diff.Empty() {
		t.Error("Plan contains change when it should be empty")
	}
}

func TestRun(t *testing.T) {
	filename := "../testdata/change/tf_plan"

	exitcode := Run(filename, true)
	if exitcode != 1 {
		err := fmt.Sprintf("Exit code is %d when it should be 1", exitcode)
		t.Error(err)
	}

	filename = "../testdata/nochange/tf_plan"

	exitcode = Run(filename, false)
	if exitcode != 0 {
		err := fmt.Sprintf("Exitcode is %d when it should be 0", exitcode)
		t.Error(err)
	}

	filename = "../testdata/badplan/tf_plan"

	exitcode = Run(filename, false)
	if exitcode != 2 {
		err := fmt.Sprintf("Exitcode is %d when it should be 2", exitcode)
		t.Error(err)
	}
}
