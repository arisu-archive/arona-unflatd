# arona-unflatd

AronaUnflatd is a specialized tool designed to reconstruct FlatBuffer schema files from decompiled C# code. It analyzes the structure and attributes of decompiled FlatBuffer-generated C# classes and automatically generates the corresponding `.fbs` schema files, making it easier to recover original schemas when they're not available.

## Prerequisites

To use this tool, you need to have the following installed on your system:

|  Tool  | Version |
| :----: | :-----: |
|   Go   | ≥ 1.22  |

## Key Features

- Converts decompiled C# FlatBuffer classes back to `.fbs` schema
- Preserves table structures, fields, and attributes
- Reconstructs enum definitions
- Maintains field ordering and types
- Supports nested tables and unions

## Installation

```sh
go install github.com/arisu-archive/arona-unflatd@latest
```

## Usage

```sh
arona-unflatd [flags] <input-directory>
```

### Flags

| Flag | Description |
| :--: | :---------: |
| `-i` | Input directory for decompiled C# code |
| `-o` | Output directory for generated schema files |
| `-v` | Enable verbose logging |

### Example

```sh
arona-unflatd -i ./src/FlatData -o ./schema
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
