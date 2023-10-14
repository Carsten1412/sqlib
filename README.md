# sqlib
A VERY SMALL SQLITE PACKAGE

Here a few examples of the functions available in this package.

Creates a new Database:d

myMap, err := sqlib.CreateNewDatabase("TEST.data)
if err != nil {
  fmt.Println(err.Error())
}
  
