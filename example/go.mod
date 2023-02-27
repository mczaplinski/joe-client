module github.com/mczaplinski/joe-client/example

go 1.20

replace github.com/mczaplinski/joe-client/joeconvert => ../joeconvert

replace github.com/mczaplinski/joe-client/joectl => ../joectl

replace github.com/mczaplinski/joe-client/utils => ../utils

require (
	github.com/mczaplinski/joe-client/joeconvert v0.0.0-00010101000000-000000000000
	github.com/mczaplinski/joe-client/joectl v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mczaplinski/joe-client/utils v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
)
