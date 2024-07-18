
# repo/tut_ATourOfGo

following [A Tour of Go](https://go.dev/tour/welcome/1)

Only loosely following & noting. Skimming sections and stuff since im not starting from absolute 0.

> these notes are hardly more effective than just going through A Tour of Go again. Its just that writing it down helps me remember even if I dont reference it later

---

syntax notes:
- can name return values at the top of the function in the return-type. Declares the vars as in the arg list, & then when returning can just return\n. Use with caution
- typecasts: targetType(variable). must be explicit


misc:
- variables are initialized to a default value by default / 0, false, ""
- int/uint are 32 or 64 bits for respective systems
- division floors result unless an operand is a float
- untyped consts are ?kind of like macro definitions? They act like literals & take the type of context

---

### for loop

only loop keyword in go

has 3 sections: init, condition, post

- init executes before first iteration
- condition evaluated before each iteration
- post executed after each iteration

can exclude the init and/or post.

can exclude all 3 for a "while true" effect


### ifs

ifs can have an 'init' section like a for loop

init-ed vars are within scope of 'else' blocks too


### switches

don't need 'break' statements. its automatic at the end of a case/start of the next

cases can be non-integers or non-constants
	// how does it build the index then? Is it just less efficient in those cases?

switches without a condition have a 'true' value


### defer

a 'deferr'-ed line executes that ~~statement~~ (function) only after the current function finishes executing

last-in-first-out defer execution order / on a stack


### pointers

there is no pointer arithmetic


### structs

can use '.' to access struct-pointer-fields rather than '->' (thank god)

literals are similar to struct literals in C except you dont need parentheses around the struct name

```
type StructName struct {
	...
}
```

when initializing structs with values, use ':' instead of '=' and no '.' necessary. (speaking relative to C)


### arrays

declare array-types as: [...]type
- ... is a special case that specifies array as opposed to fluctuating-size slice

literals outlined similar to struct literals


### slices

- declare like an array but omit size inside brackets []
- can create a slice much like python/javascript with array[i:i2]

> when slicing like this, is a ptr back to original array/slice. It does not initialize values to 0
> 
> Slicing out-of-range with this will throw a runtime error

- slices are ptrs to underlying arrays btw
- can create slice-literals
- get length with len(slice)
- get capacity with cap(slice)
- a slice == nil when len == 0, cap == 0, and there is no underlying array
- can create dynamically sized arrays with 'make' built-in function

> it allocates the array & initializes to 0
>
> ` make(sliceType, length, capacity) // where capacity is optional. returns slice`

- appending to a slice

> ` append(slice, value1, value2 ...) // returns slice.`
>
> it automatically allocates a larger array if necessary & returns slice to new array. (I guess the old slice still exists though somewhere?)


### Range keyword

Range is used with a for loop to iterate over the elements of a slice/map

> ? can range iterate over custom collection types ?

range will give 2 values from a slice: Current index + copy of that index's element


### Maps

> (like a dict)

` map[typeOfKeys]typeOfValues `

can use "make()" function with maps as well

` make(mapType) `

can create map literals

> (accessing/setting values is similar to dicts in python)

check if a key has a value:

```
element, exists = map[key]
//exists will be true or false
```

### Function stuffs

#### Functions as values

` func(argumentType1, argumentType2, ...)(returnType 1, returnType 2, ...) `


#### Function Closures

> ( it seems like a convoluted way to work around pointers. Why would you want to do this? )
>
> Its kind of interesting from a syntax perspective, but it seems like it achieves the same thing as just using some pointers so I don't really understand why you would actually want it

by returning a function from within another function, you create an internal state unique to that returned function which can be altered each time the function is called. (the returned function references data from its creator-function, which persists in Go apparently)

since my description is poor: [link](https://go.dev/tour/moretypes/25)


## Methods and Interfaces

### Methods

Go methods are just functions that can take 1 argument using different syntax

```
func (meMyselfI int) FunctionName() { ... }

var number int
number.FunctionName() // <- works with a COPY of the int, not the int itself unless specified as ptr
```

the type that can call the method (named meMyselfI above) is called the *reciever* type

reciever types MUST BE DECLARED in the same package the method is declared in. To get around, can to a typedef (just "type" in Go)


#### pointer recievers

if the reciever type is specified as a pointer, you dont have to use (&var).method / you can just call the method normally. That is to say, it can be either an actual address-of-type or just type

ofc, if is a pointer reciever, will be working with original copy of data rather than new copy

best practice recommends having all recievers for a type be either pointers or values. Not a mixture of both

the ptr value passed into the reciever can be nil


### Interfaces

interfaces are a group of method signatures. They create a type that any conforming types can convert to. Conforming types must have equivalent methods to the interface-defined ones to cast without an error

Interfaces are pointer-receiver sensitive!

dont need to specify at type declaration that it implements an interface. It just does if it has those methods or doesnt

---

#### nil in interfaces

---

conceptually, think of interfaces as having 2 values: 1) value 2) concrete type
- "concrete type" is the type of the data that was cast to the interface. This is used to call the proper version of the method
- value is the actual data itself. The data which used to be the "concrete type" before it was cast to the interface

If the value of the data is nil, such as a nil ptr, the method will be called with a nil value rather than throwin an exception. Methods are generally expected to handle nil values gracefully


above I was speaking about the value/data being nil, but sometimes the "concrete type" can be nil as well. For example, if an interface is created & nothing is ever assigned to it.

In this case, a runtime exception WILL be thrown, since the interface doesnt even know which type method to call

(interface == nil)
> This is only true when the "concrete type" is nil, not necessarily when the value/data is nil


#### the Empty Interface

> as in, an interface which specifies/requires no methods

every type fits this criteria, so can be cast to the empty interface

Empty interfaces are used by code that handles/passes-around values of arbitrary/unknown type, such as fmt.Printf

` i interface{} `
> since there is only 1 empty interface, can just specify it literally each time like this
>
> ? so its kind of like passing around a void* ? If you cast it back to a given type, is it manipulating the same original data, or a copy?


#### Type Assertions

> by checking the type of an empty interface

` value, isType := emptyInterface.(typeToCheck) `

- if the interface's underlying type is typeToCheck, isType will be true & value will be the data as that type
- if isType is omitted as a return & !isType (what would be), will have a runtime panic. If types match, will just give the value as expected


#### Type Switches

> ( just using an empty interface + a switch statement checking the type to execute code conditionally on the type of the data )


#### "Stringer" interface

> a specific interface in fmt pkg

```
type Stringer interface {
	String() string
} 
```

> pretty self explanatory imo


### Errors

> Go uses errors as values, often returned alongside other return values. Sparse notes here probably

```
type error interface {
	Error() string
}
```


### Readers

io.Reader is an interface. One of the methods defines is:

` func (T) Read(b []byte) (n int, err error) `
	
- returns io.EOF as error when stream ends
- reads as many bytes as the byte slice has room for

Go has many different implementations of the Reader interface for different purposes

shows using "strings.NewReader(stringGoesHere)" to read from a string


### Images

[Image](https://pkg.go.dev/image#Image) interface in "image" package

```
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle 		// aka Image.Rectangle
	At(x, y int) color.Color
}
```

- image.Rectangle is not {x, y, w, h}. Its: {Min point, Max point} where point = {x,y}


## Generics

### Type parameters

> generics are implemented in Go through "type parameters"

` func FunctionName[T constraintName](argumentList) { ... } `

> T can be used in place of anywhere the generic-type is allowed

specifying a constraint allows you to use the generic type in that way.

> the example showed the "comparable" constraint. It allows you to use "==" or "!=" with the type

### Generic Types

> generic types in non-functions. Like structs

```
type StructName[T constraintName] struct {
	variableName T
	...
}
```


## Concurrency
