# Aula 2 - My First API

Question -> What's **difference** between a function with Uppercase and Lowercase first letter

```
func main() {
    http.HandleFunc()
	http.handleFunc()
}
```

When it starts with **uppercase** is a **public** function, **lowercase** is **private**

#### Variables

Knowing GO has a **STATIC TYPE** -> A variable cannot change your type while running

We have two way to instance a variable

```
var number int // declare
anotherNumber := 1 // declare and set a value
```

#### Functions

```
func multiply(number1 int, number2 int) int {
	return number1 * number2
}
// Void function
func multiplyNoReturn(number1 int, number2 int) {
	fmt.Println(number1 *  number2)
}
```

#### Structs

Look like a class, you can set attributes and methods

```
type Response struct {
	Message string
	Data    any
}
```

Setting a method in a struct

```
func (r Response) SomeMethod() {
	...
}
```

#### Logging

```
logger := log.Default()
logger.Fatal()
logger.Panic()
logger.Print()
```
