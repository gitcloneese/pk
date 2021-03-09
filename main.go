package pk

func init() {
	Conn = 10
	Chnn = 5
}

func Get_pk_val() (a, b int) {
	a = Conn
	b = Chnn
	return
}

func Set_pk_val(a, b int) {
	Conn = a
	Chnn = b
}
