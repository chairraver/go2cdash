# go2cdash

A small utility to convert the `go test` output to the XML format,
that is used by CMake/CTest to report the results of the build/test
process to the [CDash](https://www.cdash.org/) dashboard.

It uses the `go test` output parser from
https://github.com/jstemmer/go-junit-report, which also sparked the
idea to actually do this.

Well, that was the idea before Golang 1.10 came along. Go 1.10 now
supports an options to produce JSON out for `go test`, which I will
use moving forware.
