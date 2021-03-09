package pk

func init() {
	conn = 10
	chnn = 5
}

func get_pk_val() (a, b int) {
	a = conn
	b = chnn
	return
}

func set_pk_val(a, b int) {
	conn = a
	chnn = b
}
