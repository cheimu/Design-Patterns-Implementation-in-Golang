package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (lp *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return newMsg
}

type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (pa *PrinterAdapter) PrintStored() string {
	var newMsg string
	if pa.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", pa.Msg)
		newMsg = pa.OldPrinter.Print(newMsg)
	} else {
		newMsg = pa.Msg
	}
	fmt.Println(newMsg)
	return newMsg
}
