CSVPretty
=============================================================================
[CSVPretty]: #CSVPretty

CSVPretty is a small tool that pretty prints CSV files.

There are many tools online that claim to pretty print CSV files when 
really they are just reformatted for display. This tool produces files
which actually *are* CSV files, just with the columns aligned with the
[Elastic Tabstops] algorithm.

The CSV files should produce identical data when parsed provided your 
CSV parser is configured to drop preceeding whitespace in CSV records
(see [Limitations]).

[Elastic Tabstops]: https://godoc.org/text/tabwriter



Example
-----------------------------------------------------------------------------
[Example]: #example

For the input:

[test.csv                                                      ](test.csv)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~      csv
Name, Subspecies, Species, Genus
West African Giraffe, G. c. peralta, G. camelopardalis, Giraffa
Meerkat, , s.Suricata, Suricata
Dog, C. l. familiaris, C. lupus, Canis
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~



CSVPretty produces:

[test_out.csv                                                  ](test_out.csv)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~          csv
Name,                   Subspecies,         Species
West African Giraffe,   G. c. peralta,      G. camelopardalis
Meerkat,                ,                   s.Suricata
Dog,                    C. l. familiaris,   C. lupus
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~



Limitations
-----------------------------------------------------------------------------
[Limitations]: #limitations

The only limitation to CSVPretty is that your CSV parser drops leading
whitespace in records. This is a common option that's tenuously
considered part of the CSV format. In Go, for example you can set the 
TrimLeadingSpace field of [csv.Reader] to `true` for this.

[csv.Reader]: https://godoc.org/encoding/csv



License
-----------------------------------------------------------------------------
[License]: #license


MIT License

Copyright (c) 2019 Zemnmez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

(See [LICENSE])

[LICENSE]: LICENSE