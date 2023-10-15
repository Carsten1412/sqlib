# sqlib
## A VERY SMALL SQLITE PACKAGE

Here a few examples of the functions available in this package.

### Creates a new database:

```
func CreateNewDatabase(filename string) error
```
Example:
```
err := sqlib.CreateNewDatabase("TEST")
if err != nil {
	fmt.Println(err.Error())
}
```  

### Creates a new table and adds rows

```
func NewRowsList() TableRows
```
```
func (r *TableRows) AddRow(name string, datatype string)
```
```
func CreateTable(database string, table string, rowlist TableRows) error
```
Example:
```

rowlist := sqlib.NewRowsList() 
rowlist.AddRow("name", "TEXT") 
rowlist.AddRow("vorname", "TEXT")
rowlist.AddRow("age", "INTEGER")

err := sqlib.CreateTable("TEST", "Person", rowlist) 
if err != nil {
   fmt.Println(err.Error())
}

```

### Add values to the columns

```
func AddValue(database string, table string, key string, value string)
```
Example:
```

sqlib.AddValue("TEST", "Person", "name", "Mustermann")
sqlib.AddValue("TEST", "Person", "vorname", "Paul")
sqlib.AddValue("TEST", "Person", "age", "58")

```

### Get all data from table

```
func GetAllFromTable(database string, table string) ([]map[string]interface{}, error)
```
Example:
```

myMap, err := sqlib.GetAllFromTable("TEST", "Person")
if err != nil {
fmt.Println(err.Error())
}

```

to be continued ...

