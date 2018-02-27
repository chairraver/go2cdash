/*
Package cmake implements the data structures in the three XML files,
which are typically send to CDash after a ctest execution.
*/
package cmake

import (
	"bytes"
	"encoding/xml"

	"golang.org/x/sys/unix"
)

// TestXML describes a number of fields describing the particular
// hardware to test is executed on.
type TestXML struct {
	XMLName                      xml.Name `xml:"Site"`
	BuildName                    string   `xml:"BuildName,attr"`
	BuildStamp                   string   `xml:"BuildStamp,attr"`
	Name                         string   `xml:"Name,attr"`
	Generator                    string   `xml:"Generator,attr"`
	CompilerName                 string   `xml:"CompilerName,attr"`
	CompilerVersion              string   `xml:"CompilerVersion,attr"`
	OSName                       string   `xml:"OSName,attr"`
	HostName                     string   `xml:"Hostname,attr"`
	OSRelease                    string   `xml:"OSRelease,attr"`
	OSVersion                    string   `xml:"OSVersion,attr"`
	OSPlatform                   string   `xml:"OSPlatform,attr"`
	Is64Bits                     string   `xml:"Is64Bits,attr"`
	VendorString                 string   `xml:"VendorString,attr"`
	VendorID                     string   `xml:"VendorID,attr"`
	FamilyID                     string   `xml:"FamilyID,attr"`
	ModelID                      string   `xml:"ModelID,attr"`
	ProcessorCacheSize           int64    `xml:"ProcessorCacheSize,attr"`
	NumberOfLogicalCPU           int64    `xml:"NumberOfLogicalCPU,attr"`
	NumberOfPhysicalCPU          int64    `xml:"NumberOfPhysicalCPU,attr"`
	TotalVirtualMemory           int64    `xml:"TotalVirtualMemory,attr"`
	TotalPhysicalMemory          int64    `xml:"TotalPhysicalMemory,attr"`
	LogicalProcessorsPerPhysical int64    `xml:"LogicalProcessorsPerPhysical,attr"`
	ProcessorClockFrequency      float64  `xml:"ProcessorClockFrequency,attr"`
	Testing                      Testing
}

// Testing describes the tests, which are performed and some
// time related data (start time, execution time, etc).
type Testing struct {
	StartDateTime  string   `xml:"StartDateTime,omitempty"`
	StartTestTime  int64    `xml:"StartTestTime,omitempty"`
	TestList       []string `xml:"TestList>Test"`
	Test           []Test
	EndDateTime    string  `xml:"EndDateTime,omitempty"`
	EndTestTime    int64   `xml:"EndTestTime,omitempty"`
	ElapsedMinutes float64 `xml:"ElapsedMinutes,omitempty"`
}

// Test describes an individual test with pathnames and the invocation
// parameters.
type Test struct {
	Status          string `xml:"Status,attr"`
	Name            string `xml:"Name"`
	Path            string `xml:"Path"`
	FullName        string `xml:"FullName"`
	FullCommandLine string `xml:"FullCommandLine,omitempty"`
	Results         Results
}

// Results decribe the outcome of each executed test with some
// additional information.
type Results struct {
	NamedMeasurement []NamedMeasurement
	Measurement      string `xml:"Measurement>Value"`
}

// NamedMeasurement describes additional useful information related to
// the individual test, something like test author or links to websites.
type NamedMeasurement struct {
	Type  string `xml:"type,attr"`
	Name  string `xml:"name,attr"`
	Value string `xml:"Value"`
}

// SetName sets a new Name for TestXML.
func (t *TestXML) SetName(n string) {
	t.Name = n
}

// SetBuildName sets a new BuildName for TestXML.
func (t *TestXML) SetBuildName(b string) {
	t.BuildName = b
}

// NewTestXML return a newly allocated TestXML data structure
// with the hardware related  parameters already filled in.
func NewTestXML() (*TestXML, error) {
	cmake := TestXML{}
	cmake.Is64Bits = "0"

	var un unix.Utsname

	err := unix.Uname(&un)
	if err != nil {
		return nil, err
	}

	n := bytes.IndexByte(un.Nodename[:], 0)
	cmake.HostName = string(un.Nodename[:n])
	n = bytes.IndexByte(un.Sysname[:], 0)
	cmake.OSName = string(un.Sysname[:n])
	n = bytes.IndexByte(un.Release[:], 0)
	cmake.OSRelease = string(un.Release[:n])
	n = bytes.IndexByte(un.Machine[:], 0)
	cmake.OSPlatform = string(un.Machine[:n])
	if cmake.OSPlatform == "x86_64" {
		cmake.Is64Bits = "1"
	}
	n = bytes.IndexByte(un.Version[:], 0)
	cmake.OSVersion = string(un.Version[:n])

	return &cmake, nil
}
