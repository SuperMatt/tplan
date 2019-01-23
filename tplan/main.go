package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform/terraform"
)

//LoadPlan loads the plan from file and returns the plan object and/or an error
func LoadPlan(fname string) (*terraform.Plan, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(b)
	return terraform.ReadPlan(r)
}

//Diff loads the plan and returns true if there's a diff
func Diff(fname string) (bool, error) {
	plan, err := LoadPlan(fname)
	if err != nil {
		return false, err
	}

	if plan.Diff.Empty() {
		return false, nil
	}

	return true, nil
}

//Get loads the plan and shows its contents
func Get(fname string) (*terraform.Plan, error) {
	plan, err := LoadPlan(fname)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

//Run runs the main logic
func Run(filename string, show bool) int {
	if show {
		p, err := Get(filename)
		if err != nil {
			fmt.Println(err)
			return 2
		}

		if p.Diff.Empty() {
			fmt.Println("No changes")
			return 0
		}
		fmt.Println(p.Diff.String())
		return 1
	}
	diff, err := Diff(filename)
	if err != nil {
		return 2
	}

	if diff {
		return 1
	}
	return 0
}

func main() {
	fset := flag.NewFlagSet("fset", flag.ExitOnError)

	filename := fset.String("filename", "", "The filename of the plan")
	show := fset.Bool("show", false, "Show the contents of the file, rather than return the error code")
	version := fset.Bool("version", false, "Terraform version number")
	fset.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("")
		fmt.Println("Return codes:")
		fmt.Println("  0 - No planned changes")
		fmt.Println("  1 - Changes to be applied")
		fmt.Println("  2 - Error reading plan file")
		fmt.Println("")

		fset.PrintDefaults()
	}

	fset.Parse(os.Args[1:])

	if *version {
		fmt.Println(terraform.Version)
		os.Exit(0)
	}

	os.Exit(Run(*filename, *show))
}
