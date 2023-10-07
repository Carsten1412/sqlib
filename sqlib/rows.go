package sqlib

type TableRows struct {
	rowslist map[string]string
}

func (r *TableRows) addRow(name string, datatype string) {
	r.rowslist["name"] = datatype
}

func (r *TableRows) toString() string {
	var line string

	for k, v := range r.rowslist {
		line += k + " "
		line += v + ","
	}

	if len(line) > 0 {
		line = line[:len(line)-1]
	}

	return line
}
