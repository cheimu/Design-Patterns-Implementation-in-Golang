package bridge

import (
	"errors"
	"fmt"
	"io"
)

//------------------------Implementor------------------------------------

type PrinterAPI interface {
	PrintMessage(string) error
}

//------------------------Concrete Implementation------------------------------------

type PrinterAPI1 struct{}

func (d *PrinterAPI1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterAPI2 struct {
	Writer io.Writer
}

func (d *PrinterAPI2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("should pass an io.Writer to PrinterAPI2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

//------------------------Abstraction------------------------------------

type PrinterAbstraction interface {
	Print() error
}

//------------------------Refined Abstraction------------------------------------

type PrinterA struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PrinterA) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PrinterB struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PrinterB) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from PrinterB: %s", c.Msg))
	return nil
}
