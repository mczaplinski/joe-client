# JOEctl

JOEctl is a tool to send orders to the JOE API.

## Build or install

    go build -o ../bin/joectl
    go install .

## Usage

    joeconvert [options] 

## Options

    -h, --help                       Show this message
    -v, --version                    Show version

    Global options:

    Input options (choose one)
    -d, --data=DATA                  Input xml data (e.g. '<...>')
    -f, --file=FORMAT                Input file path (e.g. assets/example_order.xml)
    -k, --apikey=APIKEY              JOE API key (or use JOE_API_KEY env var)

    Output will be written to stdout unless you specify an output file:
    -o, --output=OUTPUT              Output filepath (e.g. order.xml)

## Commands

    order                        Send order to JOE API
    list                         List orders from JOE API
    get <order_id>               Get order from JOE API

## Examples

    # Set API key to environment
    export JOE_API_KEY=<your_api_key>

    # Read from file and write to file
    joectl --file=assets/example_order.xml order

    # Read from argument and write to stdout
    joectl --data '<?xml version="1.0" encoding="UTF-8"?>
        <ORDER>
        <ORDER_HEADER>
            <ORDER_INFO>
            <ORDER_ID>S1348</ORDER_ID>
            <ORDER_DATE>2023-02-17T18:25:43</ORDER_DATE>
            <LANGUAGE>ger</LANGUAGE>
            <PARTIES>
                <PARTY>
                <PARTY_ID>100001</PARTY_ID>
                <PARTY_ROLE>supplier</PARTY_ROLE>
                <ADDRESS>
                    <NAME>Jacob Elektronik GmbH</NAME>
                    <STREET>Greschbachstraße 2</STREET>
                    <ZIP>76229</ZIP>
                    <CITY>Karlsruhe</CITY>
                    <COUNTRY>Deutschland</COUNTRY>
                    <COUNTRY_CODED>DE</COUNTRY_CODED>
                    <EMAIL>info@jacob.de</EMAIL>
                </ADDRESS>
                </PARTY>
                <PARTY>
                <PARTY_ID>7023320</PARTY_ID>
                <PARTY_ROLE>buyer</PARTY_ROLE>
                <ADDRESS>
                    <NAME>Jacob Elektronik GmbH</NAME>
                    <STREET>An der Rossweid 5</STREET>
                    <ZIP>76229</ZIP>
                    <CITY>Karlsruhe</CITY>
                    <COUNTRY>Deutschland</COUNTRY>
                    <COUNTRY_CODED>DE</COUNTRY_CODED>
                    <EMAIL>info@jacob.de</EMAIL>
                </ADDRESS>
                </PARTY>
                <PARTY>
                <PARTY_ID>7023320</PARTY_ID>
                <PARTY_ROLE>delivery</PARTY_ROLE>
                <ADDRESS>
                    <NAME>Jacob Elektronik GmbH</NAME>
                    <NAME2>z. H. Maximilan Glaeser</NAME2>
                    <STREET>Greschbachstraße 2</STREET>
                    <ZIP>76229</ZIP>
                    <CITY>Karlsruhe</CITY>
                    <COUNTRY>Deutschland</COUNTRY>
                    <COUNTRY_CODED>DE</COUNTRY_CODED>
                    <EMAIL>info@jacob.de</EMAIL>
                </ADDRESS>
                </PARTY>
            </PARTIES>
            <ORDER_PARTIES_REFERENCE>
                <BUYER_IDREF>7023320</BUYER_IDREF>
                <SUPPLIER_IDREF>100001</SUPPLIER_IDREF>
            </ORDER_PARTIES_REFERENCE>
            <CURRENCY>EUR</CURRENCY>
            </ORDER_INFO>
            <CONTROL_INFO>
            <STOP_AUTOMATIC_PROCESSING>1</STOP_AUTOMATIC_PROCESSING>
            <GENERATOR_INFO>data bridged using joeconvert</GENERATOR_INFO>
            <GENERATION_DATE>2023-02-26T22:54:51.113938536Z</GENERATION_DATE>
            </CONTROL_INFO>
        </ORDER_HEADER>
        <ORDER_ITEM_LIST>
            <ORDER_ITEM>
            <LINE_ITEM_ID>1</LINE_ITEM_ID>
            <PRODUCT_ID>
                <SUPPLIER_PID>1301014</SUPPLIER_PID>
                <BUYER_PID>1301014</BUYER_PID>
                <DESCRIPTION_SHORT>Auch ein Hanuta schmeckt besser mit verbauten Halbleitern</DESCRIPTION_SHORT>
            </PRODUCT_ID>
            <QUANTITY>2</QUANTITY>
            <ORDER_UNIT>C62</ORDER_UNIT>
            <PRODUCT_PRICE_FIX>
                <PRICE_AMOUNT>10.000000</PRICE_AMOUNT>
            </PRODUCT_PRICE_FIX>
            <PRICE_LINE_AMOUNT>20.000000</PRICE_LINE_AMOUNT>
            </ORDER_ITEM>
        </ORDER_ITEM_LIST>
        <ORDER_SUMMARY>
            <TOTAL_ITEM_NUM>1</TOTAL_ITEM_NUM>
            <TOTAL_AMOUNT>20.000000</TOTAL_AMOUNT>
        </ORDER_SUMMARY>
        </ORDER>}' order

## Package Usage

Alternatively, you can use the code as a module in your own project:

    import "github.com/mczaplinski/joe-client/joectl/pkg/joe"

    // Place order
    joe.Order(xmlBody, apiKey)

    // List orders
    joe.Orders(apiKey)

    // Get order
    joe.Get(orderId, apiKey)
