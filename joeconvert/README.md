# JOEconvert

JOEconvert is a tool to convert between different order formats.
It reads JSON and outputs OpenTrans2.1 XML.

## Build or install

    go build -o ../bin/joeconvert
    go install .

## Usage

    joeconvert [options] 

## Options

    -h, --help                       Show this message
    -v, --version                    Show version

    Input options (choose one)
    -d, --data=DATA                  Input json data (e.g. '{...}')
    -f, --file=FORMAT                Input file path (e.g. assets/example_input.json)

    Output will be written to stdout unless you specify an output file:
    -o, --output=OUTPUT              Output filepath (e.g. order.xml)

## Examples

    # Read from file and write to file
    joeconvert --file assets/example_input.json --output order.xml

    # Read from argument and write to stdout
    joeconvert --data '{
            "ODate": "2023-02-17T18:25:43.511Z",
            "BestellNummer": "S1337",
            "Testbestellung": 1,
            "LieferantNr": 100001,
            "Artikel": [
                {
                "ArtikelNummer": "1301014",
                "Artikel Name ": "Das Elektronische Hanuta",
                "Artikel Beschreibung": "Auch ein Hanuta schmeckt besser mit verbauten Halbleitern",
                "BestellMenge": 2,
                "Preis": 10.00
                }
            ]
        }'

## Package Usage

Alternatively, you can use the code as a module in your own project:

    package main

    import (
        "fmt"
        "log"

        "github.com/mczaplinski/joe-client/joeconvert/pkg/convert"
    )

    func main() {
        data := `{
            "ODate": "2023-02-17T18:25:43.511Z",
            "BestellNummer": "S1337",
            "Testbestellung": 1,
            "LieferantNr": 100001,
            "Artikel": [
                {
                "ArtikelNummer": "1301014",
                "Artikel Name ": "Das Elektronische Hanuta",
                "Artikel Beschreibung": "Auch ein Hanuta schmeckt besser mit verbauten Halbleitern",
                "BestellMenge": 2,
                "Preis": 10.00
                }
            ]
        }`

        order, err := convert.Convert([]byte(data))
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(order))
    }
