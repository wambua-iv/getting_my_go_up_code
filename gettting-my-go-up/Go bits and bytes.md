
Go does not have a concept of absence of value
representing a null value may necessitate creating a nil pointer

## Pre-declared Identifiers 

In Go, **predeclared identifiers** are special names that are built into the language. They serve specific purposes and are not reserved words, allowing the set to expand without invalidating existing programs. Here are some notable predeclared identifiers:

1. **`true`** and **`false`**: Untyped boolean values.
2. **`iota`**: Represents the untyped integer ordinal number of the current const specification in a const declaration.
3. **`nil`**: The zero value for pointers, channels, functions, interfaces, maps, or slices.
4. **Types**: **`bool`**, **`byte`**, **`complex64`**, **`complex128`**, **`error`**, **`float32`**, **`float64`**, **`int`**, **`int8`**, **`int16`**, **`int32`**, **`int64`**, **`rune`**, **`string`**, **`uint`**, **`uint8`**, **`uint16`**, **`uint32`**, **`uint64`**, and **`uintptr`**.
5. **Functions**: <span class="purple"> **`append`**</span>, **`cap`**, **`copy`**, **`delete`**, **`imag`**, **`len`**, **`make`**, **`new`**, **`panic`**, **`print`**, **`println`**, **`real`**, and **`recover`**.

⇒ _You must be very careful to **never redefine any identifiers in the universe block**_


> [!NOTE] Shadowing
>  `:=` reuses only variables that are declared in the current block. When using `:=`, make sure that you don’t have any variables from an outer scope on the lefthand side unless you intend to shadow them.

`var` keyword initializes a variable to its zero value


