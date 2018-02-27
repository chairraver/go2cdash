/*
Package cmake implements the data structures in the three XML files,
which are typically send to CDash after a ctest execution.
*/
package cmake

import "encoding/xml"

type ConfigureXML struct {
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
	Configure                    Configure
}

type Configure struct {
	StartDateTime      string  `xml:"StartDateTime"`
	StartConfigureTime int64   `xml:"StartConfigureTime"`
	ConfigureCommand   string  `xml:"ConfigureCommand"`
	Log                string  `xml:"Log"`
	ConfigureStatus    int64   `xml:"ConfigureStatus"`
	EndDateTime        string  `xml:"EndDateTime"`
	EndConfigureTime   int64   `xml:"EndConfigureTime"`
	ElapsedMinutes     float64 `xml:"ElapsedMinutes"`
}

// SetName sets a new Name for ConfigureXML.
func (c *ConfigureXML) SetName(n string) {
	c.Name = n
}

// SetBuildName sets a new BuildName for ConfigureXML.
func (c *ConfigureXML) SetBuildName(b string) {
	c.BuildName = b
}

// NewConfigureXML return a newly allocated ConfigureXML data structure
// with the hardware related  parameters already filled in.
func NewConfigureXML() *ConfigureXML {
	cmake := ConfigureXML{}

	return &cmake
}
