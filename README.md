CSVPretty
=======================================================================
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
-----------------------------------------------------------------------
[Example]: #example

For the input:

[test.csv]
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~    csv
Name, Subspecies, Species, Genus
West African Giraffe, G. c. peralta, G. camelopardalis, Giraffa
Meerkat, , s.Suricata, Suricata
Dog, C. l. familiaris, C. lupus, Canis
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

CSVPretty produces:

[test_out.csv]
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~    csv
Name,                   Subspecies,         Species
West African Giraffe,   G. c. peralta,      G. camelopardalis
Meerkat,                ,                   s.Suricata
Dog,                    C. l. familiaris,   C. lupus
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

[test.csv]: test.csv
[test_out.csv]: test_out.csv



Limitations
-----------------------------------------------------------------------
[Limitations]: #limitations

The only limitation to CSVPretty is that your CSV parser drops leading
whitespace in records. This is a common option that's tenuously
considered part of the CSV format. In Go, for example you can set the 
TrimLeadingSpace field of [csv.Reader] to `true` for this.

[csv.Reader]: https://godoc.org/encoding/csv