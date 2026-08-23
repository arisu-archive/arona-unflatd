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

		It("should return enum conversion errors with enum context", func() {
			fileInfo := &ast.FileInfo{
				Enums: map[string]*ast.EnumInfo{
					"Unsupported": {
						Name:     "Unsupported",
						BaseType: "float",
						Keys:     []string{"None"},
						Values:   []string{"0"},
					},
				},
			}

			schema, err := converter.Convert(fileInfo)

			Expect(err).To(MatchError(ContainSubstring(`convert enum "Unsupported"`)))
			Expect(schema).To(BeNil())
		})
	})

	Context("when declarations come from maps", func() {
		It("should order tables and enums by name", func() {
			fileInfo := &ast.FileInfo{
				Structs: map[string]*ast.StructInfo{
					"Zulu":  flatbufferStruct("Zulu"),
					"Alpha": flatbufferStruct("Alpha"),
					"Mike":  flatbufferStruct("Mike"),
					"Bravo": flatbufferStruct("Bravo"),
				},
				Enums: map[string]*ast.EnumInfo{
					"ZuluEnum":  enumInfo("ZuluEnum"),
					"AlphaEnum": enumInfo("AlphaEnum"),
					"MikeEnum":  enumInfo("MikeEnum"),
					"BravoEnum": enumInfo("BravoEnum"),
				},
			}

			schema, err := converter.Convert(fileInfo)
			Expect(err).NotTo(HaveOccurred())

			tableNames := make([]string, 0, len(schema.Tables))
			for _, table := range schema.Tables {
				tableNames = append(tableNames, table.Name)
			}
			Expect(tableNames).To(Equal([]string{"Alpha", "Bravo", "Mike", "Zulu"}))

			enumNames := make([]string, 0, len(schema.Enums))
			for _, enum := range schema.Enums {
				enumNames = append(enumNames, enum.Name)
			}
			Expect(enumNames).To(Equal([]string{"AlphaEnum", "BravoEnum", "MikeEnum", "ZuluEnum"}))
		})
	})

	Context("when converting a file with a struct", func() {
		Context("when the struct is a flatbuffer type", func() {
			It("should convert struct to table correctly", func() {
				fileInfo := &ast.FileInfo{
					Namespace: "FlatData",
					Structs: map[string]*ast.StructInfo{
						"TestTable": {
							Name:     "TestTable",
							BaseList: []string{"IFlatbufferObject"},
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
							Methods: map[string]*ast.MethodInfo{
								"AddId": {
									Name:           "AddId",
									Modifiers:      []string{"public", "static"},
									ReturnType:     "void",
									ParameterNames: []string{"builder", "id"},
									ParameterTypes: []string{"FlatBufferBuilder", "int"},
								},
								"AddName": {
									Name:           "AddName",
									Modifiers:      []string{"public", "static"},
									ReturnType:     "void",
									ParameterNames: []string{"builder", "nameOffset"},
									ParameterTypes: []string{"FlatBufferBuilder", "StringOffset"},
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

			Context("when the struct is not a flatbuffer type", func() {
				It("should not convert the struct", func() {
					fileInfo := &ast.FileInfo{
						Namespace: "FlatData",
						Structs: map[string]*ast.StructInfo{
							"TestTable": {
								Name:     "TestTable",
								BaseList: []string{"NotFlatBufferType"},
							},
						},
					}

					schema, err := converter.Convert(fileInfo)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("no tables or enums found"))
					Expect(schema).To(BeNil())
				})
			})
		})
	})
})

func flatbufferStruct(name string) *ast.StructInfo {
	return &ast.StructInfo{
		Name:     name,
		BaseList: []string{"IFlatbufferObject"},
		Fields: []*ast.FieldInfo{
			{Name: "ID", Type: "int", Modifiers: []string{"public"}},
		},
		Methods: map[string]*ast.MethodInfo{
			"AddID": {Name: "AddID"},
		},
	}
}

func enumInfo(name string) *ast.EnumInfo {
	return &ast.EnumInfo{
		Name:     name,
		BaseType: "int",
		Keys:     []string{"None"},
		Values:   []string{"0"},
	}
}
