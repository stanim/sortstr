# sortstr

[![Documentation and Examples](https://godoc.org/github.com/stanim/sortstr?status.svg)](https://godoc.org/github.com/stanim/sortstr)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/stanim/sortstr/blob/master/LICENSE)

Sort slices of strings based on multiple indices or column
headers.
Reverse sort is possible by providing negative index
values or prefixing column header titles with `-`.
Note that indices are not zero based, but start with 1.
If an index is out of range, an empty string will be
used for comparing, rather than throwing a runtime
panic.

Normally you only need to use this function:

    By(rows, indices)

or these two together:

    NewHeaders(titles)
    ByHeaders(headers, rows, titles)

Only if you need to sort the same rows multiple times
in different orders, you'll need to use the `Multi` type
directly.

This was developed with the
[xlsx](https://github.com/tealeg/xlsx)
package in mind, but can be used independently.

See [godoc](https://godoc.org/github.com/stanim/sortstr) for more documentation and examples.

Released under the [MIT License](https://github.com/stanim/sortstr/blob/master/LICENSE).
