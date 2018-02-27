package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chairraver/go2cdash/cmake"
	"github.com/jstemmer/go-junit-report/parser"
)

type TestEvent struct {
	Time    time.Time // encodes as an RFC3339-format string
	Action  string
	Package string
	Test    string
	Elapsed float64 // seconds
	Output  string
}

var packageName = "testing"

func main() {

	t, err := cmake.NewTestXML()
	if err != nil {
		log.Fatal(err)
	}

	t.SetName("ene mene muh")
	t.SetBuildName("huldi-test-linux")

	_, err = xml.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + string(xmlstring))

	// Read input
	report, err := parser.Parse(os.Stdin, packageName)

	for _, pkg := range report.Packages {

		fmt.Println(pkg.Name)
		for _, tests := range pkg.Tests {
			fmt.Println(tests.Name)
			fmt.Println(tests.Output)
		}
		fmt.Println(pkg.Time)
	}

	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}

}

// Local Variables:
// compile-command: "go run go2cdash.go"
// End:
