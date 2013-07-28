
package tfloat64

type Matrix struct {
	MatrixData
}

// Returns the number of cells which is Rows()*Columns().
func (m *Matrix) Size() int {
	return m.Rows() * m.Columns()
}
