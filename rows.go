package sqlib

type TableRows struct {
	rowslist map[string]string
}

func (r *TableRows) addRow(name string, datatype string) {
	r.rowslist["name"] = datatype
}

func (r *TableRows) toString() string {
	var line string

	for _, v := range r.rowslist {

	}

	return line
}
