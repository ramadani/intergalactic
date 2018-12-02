package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ramadani/intergalactic/querier"

	"github.com/ramadani/intergalactic/conversion"
	"github.com/ramadani/intergalactic/credit"

	"github.com/ramadani/intergalactic/converter"
	"github.com/ramadani/intergalactic/numeral/roman"
)

func main() {
	romanNumeralsEngine := roman.NewRoman()
	converter := converter.NewConverter(romanNumeralsEngine)
	credit := credit.NewCredit()
	querier := querier.NewQuerier()
	conversion := conversion.NewConversion(querier, converter, credit)

	reader := bufio.NewReader(os.Stdin)
	isContinue := true

	closeHandler(func() {
		isContinue = false
	})

	fmt.Println(":: Intergalactic numerals conversion ::")

	for isContinue {
		stmt := readLine(reader)
		res, err := conversion.Query(stmt)
		if err != nil {
			fmt.Println(err.Error())
		}

		if res != "" {
			fmt.Println(res)
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func closeHandler(handle func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		handle()
		fmt.Println("\rExit")
		os.Exit(0)
	}()
}
