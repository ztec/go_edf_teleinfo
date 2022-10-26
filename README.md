EDF TELEINFO
============

This library offer the tools to read and parse data from French EDF energy meter.
This data is called "Téléinfo" and is named Teleinfo in the library.

# ALPHA STATE

I use this library on my own installation. This installation does not
cover the whole specification of EDF Téléinfo, and some fields are not parsed.
The full specification can be found [in this document](https://www.enedis.fr/sites/default/files/Enedis-NOI-CPT_02E.pdf)

The library provide the raw packet for anyone to use, but you are welcome to contribute

# Standard & Historical formats

This library only works with the historical format. This is the format of EDF meter for decades now.
The linky, the green meter, arrived with a new format called "standard" (Good luck for the next format name ...).
This format must be enabled. Personally it is not enabled and I did not bother to call Enedis to do it. Only them can do
it.
If you do, or you meter is already "standard" enabled, then
maybe [j-vizcaino/goteleinfo](https://github.com/j-vizcaino/goteleinfo)
is a better choice of library for you. I did not test it, but it seems to do the job.

# How-to use the library

import the package via

```shell
go get git2.riper.fr/ztec/go_edf_teleinfo
```

or

```shell
go get github.com/ztec/go_edf_teleinfo
```

in your program you can now use it with `git2.riper.fr/ztec/go_edf_teleinfo` or `github.com/ztec/go_edf_teleinfo`
in your imports.

```go
package main

import (
	"bufio"
	"fmt"
	"git2.riper.fr/ztec/go_edf_teleinfo"
)

func main() {

	fi, err := os.Open("/dev/ttyAMA0") // Open the interface. It must be already configured with correct parameters
	if err != nil {
		fmt.Printf("ERROR %s. \n", err)
		return
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)                // Creating a scanner reading incomming data from interface
	scanner.Split(go_edf_teleinfo.ScannerSplitter) // Adding a "content splitter" to identify each teleinfo messages

	for {
		for scanner.Scan() {
			teleinfo, err := go_edf_teleinfo.PayloadToTeleinfo(scanner.Bytes()) // Reading the latest packet  
			if err != nil {
				fmt.Printf("ERROR %s. %#v\n", err, teleinfo)
				continue
			}
			fmt.Printf("EDF TELEINFO PAYLOAD %#v\n", teleinfo) // You can now use this data as you wish
		}
	}
}
```

# Reference

The library provide the following

### Data structure

Teleinfo data is conveniently stored in the main data structure from [teleinfo.go](./teleinfo.go)

Some data you can expect:

- **HCHC**   Index heures creuses in KWh or KVArh
- **HCHP**   Index heures pleines in KWh or KVArh
- **IINST**  Intensitee instantanee in A rounded to the closer integer
- **PAPP**   Puissance apparente in VA rounded to the tenth
- **PTEC**   Période tarifaire en cours (Heure creuse ou pleine ou bleu ou rouge, ...)

There is also a **RAW** field that hold the original Bytes.

For the full list, just review [teleinfo.go](./teleinfo.go)

### Payload parsing

[PayloadToTeleinfo](./payloadToTeleinfo.go) is a function that will analyse the
given []bytes to generate a teleinfo object. It will parse each lines, and check the checksums
to ensure the data is valid. Any error is raised via the Error returned.

### Scanner splitter

If you decide to read from the serial interface directly in go, you can use the provided
[ScannerSplitter](./scannerSplitter.go) function to slice data received in coherent chunks
containing all data the meter sent.

### Versioning

This library is in alpha state however

- I'll tag any new version
- I won't break compatibility without increasing the second digit (from 0.1 to 0.2 contains breaking changes)
- I don't consider adding new fields to `Teleinfo` (and the code to extract it) a breaking change
- I consider removing a field from `Teleinfo` as a breaking change
- I consider changing the name or the intention of the existing  `PayloadToTeleinfo` and `ScannerSplitter` function as a
  breaking change

# Alternative

[j-vizcaino/goteleinfo](https://github.com/j-vizcaino/goteleinfo) seems to be doing the same job, and support
the new format called "standard" of the Linky. If this library does not work for you, give it a try.
