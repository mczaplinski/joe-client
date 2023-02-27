# JOE API Example-Client

This is a simple example client for the [JOE API](https://github.com/jacob-elektronik/joe-api-documentation) library.

It consists of two parts:

- A simple command line client to communicate with the JOE API (for now: send order only)
- A command line tool to convert other formats to the JOE API / OpenTrans 2.1 format

Refer to the READMEs in the respective subdirectories for more information:

- [joectl](joectl/README.md)
- [joeconvert](joeconvert/README.md)

## Usage

Run example:

    cd examples && go run main.go

Build binaries and run example:

    make build
    make install

    export JOE_API_KEY=<your-api-key>
    
    // convert JSON to OpenTrans 2.1 XML
    joeconvert -f input.json -o order.xml

    // place order
    joectl -f order.xml order

    // view order
    joectl get <order-id> -o response.xml

    // one-liner
    joectl --data="$(joeconvert -f input.json)" order

Other:

    make test
    make lint
