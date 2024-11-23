package conversion_test

import (
	"log/slog"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/conversion"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
)

var _ = Describe("SchemaConverter", func() {
	var (
		converter *conversion.SchemaConverter
		logger    *slog.Logger
	)

	BeforeEach(func() {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		converter = conversion.NewSchemaConverter(logger)
	})

	Context("when converting a file with an enum", func() {
		It("should convert enum correctly", func() {
			fileInfo := &ast.FileInfo{
				Namespace: "FlatData",
				Enums: map[string]*ast.EnumInfo{
					"TestEnum": {
						Name:     "TestEnum",
						BaseType: "int",
						Keys:     []string{"None", "First", "Second"},
						Values:   []string{"0", "1", "2"},
					},
				},
			}

			schema, err := converter.Convert(fileInfo)
			Expect(err).NotTo(HaveOccurred())
			Expect(schema.Namespace).To(Equal("FlatData"))
			Expect(schema.Enums).To(HaveLen(1))
			Expect(schema.Enums[0].Name).To(Equal("TestEnum"))
			Expect(schema.Enums[0].Values).To(HaveLen(3))
		})
	})

	Context("when converting a file with a struct", func() {
		It("should convert struct to table correctly", func() {
			fileInfo := &ast.FileInfo{
				Namespace: "FlatData",
				Structs: map[string]*ast.StructInfo{
					"TestTable": {
						Name: "TestTable",
						Fields: []*ast.FieldInfo{
							{
								Name:      "Id",
								Type:      "int",
								Modifiers: []string{"public"},
							},
							{
								Name:      "Name",
								Type:      "string",
								Modifiers: []string{"public"},
							},
						},
					},
				},
			}

			schema, err := converter.Convert(fileInfo)
			Expect(err).NotTo(HaveOccurred())
			Expect(schema.Tables).To(HaveLen(1))
			Expect(schema.Tables[0].Name).To(Equal("TestTable"))
			Expect(schema.Tables[0].Fields).To(HaveLen(2))
		})
	})
})
