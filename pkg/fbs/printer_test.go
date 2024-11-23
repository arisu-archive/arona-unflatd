package fbs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

var _ = Describe("Printer", func() {
	var p *fbs.Printer

	BeforeEach(func() {
		p = fbs.NewPrinter()
	})

	Context("when printing with indentation", func() {
		It("should handle basic printing", func() {
			p.Print("hello")
			Expect(p.Flush()).To(Equal("hello"))
		})

		It("should handle println", func() {
			p.Println("hello")
			Expect(p.Flush()).To(Equal("hello\n"))
		})

		It("should handle format strings", func() {
			p.Println("hello %s", "world")
			Expect(p.Flush()).To(Equal("hello world\n"))
		})

		It("should handle indentation", func() {
			p.Println("level 0")
			p.Indent()
			p.Println("level 1")
			p.Indent()
			p.Println("level 2")
			p.Unindent()
			p.Println("back to level 1")
			p.Unindent()
			p.Println("back to level 0")

			expected := `level 0
    level 1
        level 2
    back to level 1
back to level 0
`
			Expect(p.Flush()).To(Equal(expected))
		})
	})

	Context("when handling complex structures", func() {
		It("should format nested blocks correctly", func() {
			p.Println("table Test {")
			p.Indent()
			p.Println("id:int;")
			p.Println("name:string;")
			p.Println("items:[Item];")
			p.Unindent()
			p.Println("}")

			expected := `table Test {
    id:int;
    name:string;
    items:[Item];
}
`
			Expect(p.Flush()).To(Equal(expected))
		})

		It("should handle multiple flushes", func() {
			p.Println("first")
			Expect(p.Flush()).To(Equal("first\n"))

			p.Println("second")
			Expect(p.Flush()).To(Equal("second\n"))
		})
	})

	Context("when handling edge cases", func() {
		It("should handle empty strings", func() {
			p.Print("")
			Expect(p.Flush()).To(Equal(""))
		})

		It("should handle multiple newlines", func() {
			p.Println("")
			p.Println("")
			p.Println("text")
			Expect(p.Flush()).To(Equal("\n\ntext\n"))
		})

		It("should handle negative indentation gracefully", func() {
			p.Unindent() // Try to unindent at level 0
			p.Println("text")
			Expect(p.Flush()).To(Equal("text\n"))
		})
	})
})
