A CLI tool for Concise Message Wrapper.

This tool can:
1. Create CMW collections from a set of individual files, with files containing Evidence CBOR.
2. Append to existing CMW collections, by passing in file(s) as input.

Currently in a prototype stage.
TODO LIST:

- ☑️ Serialise a CMW with  a basic evidence CBOR
- ☑️ Seralise multiple files into the same CMW
- Append CBOR evidence(s) to an existing CMW

CURRENT PROGRESS:

To create a record from a single file:

`go run main.go collection --data file.cbor --output out.cbor`

Flags:
`data` - The file path of the CBOR evidence file to convert into a CMW.
`output` - The file to write the CMW to.

To create a collection:

`go run main.go collection --data file.cbor --label "label" [--data file2.cbor --label "label2" ...] [--cmwc_t "URI/OID"] --output out.cbor`

Flags:
`data` - The file path(s) of the CBOR evidence file(s) to add to the CMW collection.
`label` - The label to prepend to its respective evidence in the CMW.
`cmwc_t` - (Optional) The URI / OID of the collection.
`output` - The file to write the CMW to.

For more information, see: https://www.rfc-editor.org/info/rfc9999/#section-6

