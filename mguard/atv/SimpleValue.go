package atv

import (
	"fmt"
	"strconv"
	"strings"
)

// SimpleValue represents a simple value in an ATV document.
type SimpleValue struct {
	Value string `@String`
}

// WriteDocumentPart writes a part of the ATV document to the specified writer.
func (value *SimpleValue) WriteDocumentPart(writer *strings.Builder, indent int) error {
	line := fmt.Sprintf("%s\n", strconv.Quote(value.Value))
	_, err := writer.WriteString(line)
	return err
}
