# Executing:
```sh
go run something.go
```

# Tips:


# Comments and indentations:
1. Comment using `//` 
2. Avoid using walrus operator `:=` for defining variables 
3. Maintain Naming Convenions
4. Organize properly from the start (use Moduleraization)
5. Split and package if needed
   

# Naming Conventions:

1. Files are named using snake case ie: `data_structures.go`, or `magical_script_praser.go`
2. Variables/identifiers are named in mixedCaps also known as camelCase ie `age`, `sex`, `dateOfBrith`
3. If the indentifier is needed to be exported and be used on outside package use CamelCase starting with a capital letter ie `Name`, `CountryOfOrigin`
4. Functions/Interfaces are named using Capital starting CamelCase ending with "er" suffix ie: `Writer`, `Reader`, & `HtmlCleanParser`
5. Short variable names (it needs to be short but descriptive) 
6. Unique names with capital letters should not be lowercased ie `productAPI`, `HTTPScontent` etc.


# DataStructures in Go:

There are 4 fundamential types which have their own subtypes:

1. String
2. Boolean (true/false)
3. Numeric (int/float)
4. Derived (Pointer, Array, Map, Structure, & Interface)

4.1. Pointer:


4.2.


# DS tips:
1. You can definitely use (:=) walrus operator for everything. But, it’s difficult and confusing to trace the error. So, rather you should have a practice of defining the data type of the variable.