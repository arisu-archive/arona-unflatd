package fbs

import (
	"fmt"
	"strings"
)

type Printer struct {
	w   *strings.Builder
	ind int // indentation level
}

func NewPrinter() *Printer {
	return &Printer{w: &strings.Builder{}}
}

func (p *Printer) Indent()   { p.ind++ }
func (p *Printer) Unindent() { p.ind-- }
func (p *Printer) Newline()  { p.w.WriteByte('\n') }

func (p *Printer) Print(format string, args ...any) {
	if p.ind > 0 {
		p.w.WriteString(strings.Repeat("    ", p.ind))
	}
	fmt.Fprintf(p.w, format, args...)
}

func (p *Printer) Println(format string, args ...any) {
	p.Print(format, args...)
	p.Newline()
}

func (p *Printer) Flush() string {
	result := p.w.String()
	p.w.Reset()
	return result
}
