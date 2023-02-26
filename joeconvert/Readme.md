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
    -d, --data=DATA                  Input json data (e.g. '{...}')
    -f, --file=FORMAT                Input file path (e.g. /tmp/orders.json)
    -o, --output=OUTPUT              Output filepath (e.g. /tmp/orders.xml)
    -v, --version                    Show version
    -h, --help                       Show this message

## Examples

    joeconvert --file pkg/data/example.json --output order.xml