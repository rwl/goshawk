package tfloat64

import (
	"fmt"
	common "github.com/rwl/goshawk"
)

// Flexible, well human readable matrix print formatting; By default decimal
// point aligned. Currently works on 1-d, 2-d and 3-d matrices. Note that in
// most cases you will not need to get familiar with this type; just call
// matrix.String() and be happy with the default formatting. This type is for
// advanced requirements.
type Formatter struct {
	common.FormatterBase
}

// Constructs and returns a matrix formatter with format "%G".
func NewFormatter() *Formatter {
	return NewFormatterFormat("%G")
}

// Constructs and returns a matrix formatter with the given format used to
// convert a single cell value.
func NewFormatterFormat(format string) *Formatter {
	return &Formatter{
		common.FormatterBase{
			Format: format,
			Alignment: common.DECIMAL,
		},
	}
}

// Converts a given cell to a String; no alignment considered.
func (f *Formatter) Form(vector Vec, index int) string {
	if index < 0 || index >= vector.Size() {
		return "index error"
	}
	return fmt.Sprintf(f.Format, vector.GetQuick(index))
}

//  Returns a string representations of all cells; no alignment considered.
func (f *Formatter) FormatMatrix(matrix Mat) [][]string {
	strings := make([][]string,matrix.Rows(), matrix.Columns())
	/*for row := 0; row < matrix.Rows(); row++ { TODO: implement ViewRow
		strings[row] = f.FormatRow(matrix.ViewRow(row))
	}*/
	return strings
}

//  Returns a string representations of all cells; no alignment considered.
func (f *Formatter) FormatRow(vector Vec) []string {
	s := vector.Size()
	strings := make([]string, s)
	for i := 0; i < s; i++ {
		strings[i] = f.Form(vector, i)
	}
	return strings
}

// Returns a short string representation describing the shape of the vector.
func (f *Formatter) VectorShape(vector Vec) string {
	// return "Matrix1D of size="+matrix.Size
	// return matrix.Size+" element matrix"
	// return "matrix("+matrix.Size+")"
	return fmt.Sprintf("%d vector", vector.Size())
}

// Returns a short string representation describing the shape of the matrix.
func (f *Formatter) MatrixShape(matrix Mat) string {
	return fmt.Sprintf("%d x %d matrix", matrix.Rows(), matrix.Columns())
}

// Returns a short string representation describing the shape of the cube.
func (f *Formatter) CubeShape(matrix Cub) string {
	return fmt.Sprintf("%d x %d x %d matrix", matrix.Slices(), matrix.Rows(), matrix.Columns())
}

// Returns a string representation of the given vector.
func (f *Formatter) VectorToString(v Vec) string {
	//	easy := NewMatrix(1, v.Size())
	//	easy.ViewRow(0).AssignVector(v)
	return ""//f.MatrixToString(easy)
}

// Returns a string representation of the given matrix.
func (f *Formatter) MatrixToString(matrix Mat) string {
	strings := f.FormatMatrix(matrix)
	f.Align(strings)
	total := f.ArrayToString(strings)
	if f.PrintShape {
		total = f.MatrixShape(matrix) + "\n" + total
	}
	return total
}

/*
// Returns a string representation of the given matrix.
func (f *Formatter) CubeToString(matrix *Cube) string {
	var buf bytes.Buffer
	oldPrintShape := f.PrintShape
	f.PrintShape = false
	for slice := 0; slice < matrix.Slices; slice++ {
		if slice != 0 {
			buf.WriteString(f.SliceSeparator)
		}
		buf.WriteString(f.VectorToString(matrix.ViewSlice(slice)))
	}
	f.PrintShape = oldPrintShape
	if printShape {
		return f.CubeShape(matrix) + "\n" + buf.String()
	}
	return buf.String()
}
*/
/*
func (f *Formatter) VectorToSourceCode(matrix Vector) string {
	var copy Formatter = f.Clone()
	copy.PrintShape = false
	copy.ColumnSeparator = ", "
	lead := "{"
	trail := "};"
	return lead + copy.VectorToString(matrix) + trail
}

func (f *Formatter) MatrixToSourceCode(matrix Matrix) string {
	var copy Formatter = f.Clone()
	b3 := blanks(3)
	copy.PrintShape = false
	copy.ColumnSeparator = ", "
	copy.tRowSeparator = "},\n" + b3 + "{"
	lead := "{\n" + b3 + "{"
	trail := "}\n};"
	return lead + copy.MatrixToString(matrix) + trail
}

func (f *Formatter) CubeToSourceCode(matrix Cube) string {
	var copy Formatter := f.Clone()
	b3 := f.Blanks(3)
	b6 := f.Blanks(6)
	copy.PrintShape = false
	copy.ColumnSeparator = ", "
	copy.RowSeparator = "},\n" + b6 + "{"
	copy.SliceSeparator = "}\n" + b3 + "},\n" + b3 + "{\n" + b6 + "{"
	lead := "{\n" + b3 + "{\n" + b6 + "{"
	trail := "}\n" + b3 + "}\n}"
	return lead + copy.String(matrix) + trail
}
*/
