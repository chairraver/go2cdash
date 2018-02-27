/*
Package cmake implements the data structures in the three XML files,
which are typically send to CDash after a ctest execution.
*/
package cmake

import "encoding/xml"

// BuildXML is the structure where CMake/CTest records the
// build process of a project. This structure is currently unused.
type BuildXML struct {
	XMLName                      xml.Name `xml:"Site"`
	BuildName                    string   `xml:"BuildName,attr"`
	BuilStamp                    string   `xml:"BuildStamp,attr"`
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
	Build                        Build
}

// Build describes the actual build command, that controlls the build
// process of the whole project.
type Build struct {
	StartDateTime  string `xml:"StartDateTime"`
	StartBuildTime int64  `xml:"StartBuildTime"`
	BuildCommand   string `xml:"BuildCommand"`
	Warning        []Warning
}

// Warning describes the individual compilation warnings during
// the compilation process.
type Warning struct {
	BuildLogLine     int64  `xml:"BuildLogLine"`
	Text             string `xml:"Text"`
	SourceFile       string `xml:"SourceFile"`
	SourceLineNumber int64  `xml:"SourceLineNumber"`
	PreContext       string `xml:"PreContext"`
	PostContext      string `xml:"PostContext"`
	RepeatCount      int64  `xml:"RepeatCount"`
}

// SetName sets a new Name for BuildXML.
func (b *BuildXML) SetName(n string) {
	b.Name = n
}

// SetBuildName sets a new BuildName for BuildXML.
func (b *BuildXML) SetBuildName(bn string) {
	b.BuildName = bn
}

// NewBuildXML return a newly allocated BuildXML data structure
// with the hardware related  parameters already filled in.
func NewBuildXML() *BuildXML {
	cmake := BuildXML{}

	return &cmake
}
