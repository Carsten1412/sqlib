# sqlib
## A VERY SMALL SQLITE PACKAGE

Here a few examples of the functions available in this package.

### Creates a new database:

```
err := sqlib.CreateNewDatabase("TEST")
if err != nil {
	fmt.Println(err.Error())
}
```  
