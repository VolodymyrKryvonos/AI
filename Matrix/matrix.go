package Matrix

type Matrix struct {
	columns int
	rows    int
	m       [][]float64
}

func New(r int, c int) *Matrix {
	m := make([][]float64, r)
	for i := range m {
		m[i] = make([]float64, c)
	}
	return &Matrix{
		columns: c,
		rows:    r,
		m:       m,
	}
}

func (m *Matrix) Set(i int, j int, value float64) {
	if i < m.rows && j < m.columns {
		m.m[i][j] = value
	} else {
		panic("index out of range")
	}
}

func (m Matrix) At(i, j int) float64 {
	if i < m.rows && j < m.columns {
		return m.m[i][j]
	} else {
		panic("index out of range")
	}
}

func Product(m1, m2 Matrix) Matrix {
	if m1.columns == m2.rows {
		matrix := make([][]float64, m1.rows)
		for i := 0; i < m1.rows; i++ {
			matrix[i] = make([]float64, m2.columns)
			for j := 0; j < m1.rows; j++ {
				for k := 0; k < m1.rows; k++ {
					matrix[i][j] += m1.m[i][k] * m2.m[k][j]
				}
			}
		}
		return Matrix{
			columns: m2.columns,
			rows:    m1.rows,
			m:       matrix,
		}
	} else {
		panic("invalid shape")
	}
}

//func (m Matrix) ActivateLayer(activationFunction func(f float64) float64) Matrix {
//	if m.columns==1{
//		matrix := make([][]float64,m.rows)
//		for i:=0;i<m.rows;i++{
//			matrix[i]=activationFunction(m.m[i][0])
//		}
//		return Matrix{
//			columns: 1,
//			rows:    m.rows,
//			m:       matrix,
//		}
//	} else {
//		panic("used only for one layer")
//	}
//}
