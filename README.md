# intergalactic

Intergalactic numerals conversion

## System Design Solution

## Requirements

* Go >= 1.10
* A built-in testing command `go test`

## Installation

Just clone this repository and make sure `go` was installed on your machine.

## Testing

Run the tests with:

```bash
go test ./...
```

Or see the log all tests and the coverage the code with:

```bash
go test ./... -v -cover
```

## Running

To start using this application, you can use:

```bash
$ go run main.go
```

or you can use the `intergalactic` binary based on your OS from the build directory. Example:

```bash
$ build/macos/64/intergalactic
```

## Usage

This application is a CLI (Command Line Interface) tool, after you run this application, you can start using it.

```bash
:: Intergalactic numerals conversion ::
glob is I
prok is V
pish is X
tegj is L
```

```bash
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
```

```bash
how much is pish tegj glob glob ?
pish tegj glob glob is 42
```

```bash
how many Credits is glob prok Silver ?
glob prok Silver is 68 Credits
```

```bash
how many Credits is glob prok Gold ?
glob prok Gold is 57800 Credits
```

```bash
how many Credits is glob prok Iron ?
glob prok Iron is 782 Credits
```

```bash
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?
I have no idea what you are talking about
```

To exit this application, just press `ctrl+c`.
