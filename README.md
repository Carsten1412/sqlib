# sqlib
A VERY SMALL SQLITE PACKAGE

Here a few examples of the functions available in this pachage.

Creates a new Database:

myMap, err := sqlib.GetAllFromTable("TEST", "Person")
  if err != nil {
    fmt.Println(err.Error())
  }
