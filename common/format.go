package common

import (
	"bytes"
	"strings"
	"fmt"
)

const (
	LEFT                     = "left"    // The alignment string aligning the cells of a column to the left.
	CENTER                   = "center"  // The alignment string aligning the cells of a column to its center.
	RIGHT                    = "right"   // The alignment string aligning the cells of a column to the right.
	DECIMAL                  = "decimal" // The alignment string aligning the cells of a column to the decimal point.
	DEFAULT_MIN_COLUMN_WIDTH = 1         // The default minimum number of characters a column may have; currently 1.
	DEFAULT_COLUMN_SEPARATOR = " "       // The default string separating any two columns from another; currently " ".
	DEFAULT_ROW_SEPARATOR    = "\n"      // The default string separating any two rows from another; currently "\n".
	DEFAULT_SLICE_SEPARATOR  = "\n\n"    // The default string separating any two slices from another; currently "\n\n".
)

const maxInt = int(^uint(0)>>1)
const minInt = -(maxInt - 1)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var blanksCache []string = setupBlanksCache()

// for efficient String manipulations

// Pre-fabricate 40 static strings with 0,1,2,..,39 blanks, for usage
// within method blanks(length).
// Now, we don't need to construct and fill them on demand, and garbage
// collect them again.
// All 40 strings share the identical char[] array, only with different
// offset and length --> somewhat smaller static memory footprint
func setupBlanksCache() []string {
	size := 40
	blanks := make([]string, size)
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteString(" ")
	}
	str := buf.String()
	for i := 0; i < size; i++ {
		blanks[i] = str[:i]
		//fmt.Println(i + "-" + blanksCache[i] + "-")
	}
	return blanks
}

// Base type for flexible, well human readable matrix print
// formatting. Value type independent. A single cell is formatted via a format
// string. Columns can be aligned left, centered, right and by decimal point.
//
// A column can be broader than specified by the parameter
// minColumnWidth (because a cell may not fit into that width) but a
// column is never smaller than minColumnWidth. Normally one does not
// need to specify minColumnWidth. Cells in a row are separated by a
// separator string, similar separators can be set for rows and slices.
type FormatterBase struct {
	Alignment       string // The default format string for formatting a single cell value; currently "%G".
	Format          string // The default format string for formatting a single cell value; currently "%G".
	MinColumnWidth  int    // The default minimum number of characters a column may have; currently 1.
	ColumnSeparator string // The default string separating any two columns from another; currently " ".
	RowSeparator    string // The default string separating any two rows from another; currently "\n".
	SliceSeparator  string // The default string separating any two slices from another; currently "\n\n".
	PrintShape      bool   // Tells whether String representations are to be preceded with summary of the shape; currently "true".
}

func NewFormatter() *FormatterBase {
	return &FormatterBase{
		LEFT,
		"%G",
		DEFAULT_MIN_COLUMN_WIDTH,
		DEFAULT_COLUMN_SEPARATOR,
		DEFAULT_ROW_SEPARATOR,
		DEFAULT_SLICE_SEPARATOR,
		true,
	}
}

// Returns the index of the decimal point.
func (f *FormatterBase) IndexOfDecimalPoint(s string) int {
	i := strings.LastIndex(s, ".")
	if i < 0 {
		i = strings.LastIndex(s, "e")
	}
	if i < 0 {
		i = strings.LastIndex(s, "E")
	}
	if i < 0 {
		i = len(s)
	}
	return i
}

// Returns the number of characters before the decimal point.
func (f *FormatterBase) Lead(s string) int {
	if f.Alignment == DECIMAL {
		return f.IndexOfDecimalPoint(s)
	}
	return len(s)
}

// Modifies the strings in a column of the string matrix to be aligned
// (left, centered, right, decimal).
func (f *FormatterBase) Align(strings [][]string) {
	rows := len(strings)
	columns := 0
	if rows > 0 {
		columns = len(strings[0])
	}
	maxColWidth := make([]int, columns)
	var maxColLead []int = nil
	isDecimal := f.Alignment == DECIMAL
	if isDecimal {
		maxColLead = make([]int, columns)
	}
	// maxColTrail = make([]int, columns)

	// for each column, determine alignment parameters
	for column := 0; column < columns; column++ {
		maxWidth := f.MinColumnWidth
		maxLead := minInt
		// maxTrail := minInt
		for row := 0; row < rows; row++ {
			s := strings[row][column]
			maxWidth = max(maxWidth, len(s))
			if isDecimal {
				maxLead = max(maxLead, f.Lead(s))
			}
			// maxTrail := math.Max(maxTrail, f.Trail(s))
		}
		maxColWidth[column] = maxWidth
		if isDecimal {
			maxColLead[column] = maxLead
		}
		// maxColTrail[column] = maxTrail
	}

	// format each row according to alignment parameters
	for row := 0; row < rows; row++ {
		f.AlignRow(strings[row], maxColWidth, maxColLead)
	}
}

// Modifies the strings the string matrix to be aligned
// (left, centered, right, decimal).
func (f *FormatterBase) AlignRow(row []string, maxColWidth, maxColLead []int) {
	var s bytes.Buffer

	columns := len(row)
	for column := 0; column < columns; column++ {
		s.Reset()
		c := row[column]
		if f.Alignment == RIGHT {
			s.WriteString(f.Blanks(maxColWidth[column] - s.Len()))
			s.WriteString(c)
		} else if f.Alignment == DECIMAL {
			s.WriteString(f.Blanks(maxColLead[column] - f.Lead(c)))
			s.WriteString(c)
			s.WriteString(f.Blanks(maxColWidth[column] - s.Len()))
		} else if f.Alignment == CENTER {
			s.WriteString(f.Blanks((maxColWidth[column] - len(c))/2))
			s.WriteString(c)
			s.WriteString(f.Blanks(maxColWidth[column] - s.Len()))
		} else if f.Alignment == LEFT {
			s.WriteString(c)
			s.WriteString(f.Blanks(maxColWidth[column] - s.Len()))
		} else {
			fmt.Errorf("invalid alignment:%s", f.Alignment)
			return
		}
		row[column] = s.String()
	}
}

// Returns a string with length blanks.
func (f *FormatterBase) Blanks(length int) string {
	if length < 0 {
		length = 0
	}
	if length < len(blanksCache) {
		return blanksCache[length]
	}

	var buf bytes.Buffer
	for k := 0; k < length; k++ {
		buf.WriteString(" ")
	}
	return buf.String()
}

// Returns a single string representation of the given string arrays.
func (f *FormatterBase) ArrayToString(strings [][]string) string {
	rows := len(strings)
	columns := 0
	if len(strings) > 0 {
		columns = len(strings[0])
	}

	var total bytes.Buffer
	var s bytes.Buffer
	for row := 0; row < rows; row++ {
		s.Reset()
		for column := 0; column < columns; column++ {
			s.WriteString(strings[row][column])
			if column < columns - 1 {
				s.WriteString(f.ColumnSeparator)
			}
		}
		total.WriteString(s.String())
		if row < rows - 1 {
			total.WriteString(f.RowSeparator)
		}
	}
	return total.String()
}
