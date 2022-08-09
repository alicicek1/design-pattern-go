package main

// Separate interfaces don't create a god interface

type Document struct {
}
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {
	//ok
}

func (m *MultiFunctionPrinter) Fax(d Document) {
	//ok
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	//ok
}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {
	//ok
}

// Deprecated: Old fashioned printer doesn't support faxing.
func (o *OldFashionedPrinter) Fax(d Document) {
	//TODO implement me
	panic("operation is not supported")
}

// Deprecated: Old fashioned printer doesn't support scanning.
func (o *OldFashionedPrinter) Scan(d Document) {
	//TODO implement me
	panic("operation is not supported")
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type MyPrinter struct {
}

func (m *MyPrinter) Print(d Document) {
	//ok
}

type PhotoCopier struct {
}

func (p *PhotoCopier) Fax(d Document) {
	//ok
}

func (p *PhotoCopier) Scan(d Document) {
	//ok
}

func (p *PhotoCopier) Print(d Document) {
	//ok
}

type MutliFunctionDevice interface {
	Printer
	Scanner
	Faxer
}

//decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
	faxer   Faxer
}

func (m MultiFunctionMachine) Fax(d Document) {
	m.faxer.Fax(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {
}
