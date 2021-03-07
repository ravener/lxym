
package pages

const Julia = "Julia is a new homoiconic functional language focused on technical computing.\nWhile having the full power of homoiconic macros, first-class functions,\nand low-level control, Julia is as easy to learn and use as Python.\n\nThis is based on Julia version 1.0.0.\n\n```julia\n# Single line comments start with a hash (pound) symbol.\n#= Multiline comments can be written\n   by putting '#=' before the text  and '=#'\n   after the text. They can also be nested.\n=#\n\n####################################################\n## 1. Primitive Datatypes and Operators\n####################################################\n\n# Everything in Julia is an expression.\n\n# There are several basic types of numbers.\ntypeof(3)       # => Int64\ntypeof(3.2)     # => Float64\ntypeof(2 + 1im) # => Complex{Int64}\ntypeof(2 // 3)  # => Rational{Int64}\n\n# All of the normal infix operators are available.\n1 + 1      # => 2\n8 - 1      # => 7\n10 * 2     # => 20\n35 / 5     # => 7.0\n10 / 2     # => 5.0  # dividing integers always results in a Float64\ndiv(5, 2)  # => 2    # for a truncated result, use div\n5 \\ 35     # => 7.0\n2^2        # => 4    # power, not bitwise xor\n12 % 10    # => 2\n\n# Enforce precedence with parentheses\n(1 + 3) * 2  # => 8\n\n# Julia (unlike Python for instance) has integer under/overflow\n10^19      # => -8446744073709551616\n# use bigint or floating point to avoid this\nbig(10)^19 # => 10000000000000000000\n1e19       # => 1.0e19\n10.0^19    # => 1.0e19\n\n# Bitwise Operators\n~2         # => -3 # bitwise not\n3 & 5      # => 1  # bitwise and\n2 | 4      # => 6  # bitwise or\nxor(2, 4)  # => 6  # bitwise xor\n2 >>> 1    # => 1  # logical shift right\n2 >> 1     # => 1  # arithmetic shift right\n2 << 1     # => 4  # logical/arithmetic shift left\n\n# Use the bitstring function to see the binary representation of a number.\nbitstring(12345)\n# => \"0000000000000000000000000000000000000000000000000011000000111001\"\nbitstring(12345.0)\n# => \"0100000011001000000111001000000000000000000000000000000000000000\"\n\n# Boolean values are primitives\ntrue\nfalse\n\n# Boolean operators\n!true   # => false\n!false  # => true\n1 == 1  # => true\n2 == 1  # => false\n1 != 1  # => false\n2 != 1  # => true\n1 < 10  # => true\n1 > 10  # => false\n2 <= 2  # => true\n2 >= 2  # => true\n# Comparisons can be chained, like in Python but unlike many other languages\n1 < 2 < 3  # => true\n2 < 3 < 2  # => false\n\n# Strings are created with \"\n\"This is a string.\"\n\n# Character literals are written with '\n'a'\n\n# Strings are UTF8 encoded, so strings like \"π\" or \"☃\" are not directly equivalent\n# to an array of single characters.\n# Only if they contain only ASCII characters can they be safely indexed.\nascii(\"This is a string\")[1]    # => 'T'\n# => 'T': ASCII/Unicode U+0054 (category Lu: Letter, uppercase)\n# Beware, Julia indexes everything from 1 (like MATLAB), not 0 (like most languages).\n# Otherwise, iterating over strings is recommended (map, for loops, etc).\n\n# String can be compared lexicographically, in dictionnary order:\n\"good\" > \"bye\"   # => true\n\"good\" == \"good\" # => true\n\"1 + 2 = 3\" == \"1 + 2 = $(1 + 2)\" # => true\n\n# $(..) can be used for string interpolation:\n\"2 + 2 = $(2 + 2)\" # => \"2 + 2 = 4\"\n# You can put any Julia expression inside the parentheses.\n\n# Printing is easy\nprintln(\"I'm Julia. Nice to meet you!\")  # => I'm Julia. Nice to meet you!\n\n# Another way to format strings is the printf macro from the stdlib Printf.\nusing Printf   # this is how you load (or import) a module\n@printf \"%d is less than %f\\n\" 4.5 5.3   # => 5 is less than 5.300000\n\n\n####################################################\n## 2. Variables and Collections\n####################################################\n\n# You don't declare variables before assigning to them.\nsomeVar = 5  # => 5\nsomeVar      # => 5\n\n# Accessing a previously unassigned variable is an error\ntry\n    someOtherVar  # => ERROR: UndefVarError: someOtherVar not defined\ncatch e\n    println(e)\nend\n\n# Variable names start with a letter or underscore.\n# After that, you can use letters, digits, underscores, and exclamation points.\nSomeOtherVar123! = 6  # => 6\n\n# You can also use certain unicode characters\n# here ☃ is a Unicode 'snowman' characters, see http://emojipedia.org/%E2%98%83%EF%B8%8F if it displays wrongly here\n☃ = 8  # => 8\n# These are especially handy for mathematical notation, like the constant π\n2 * π  # => 6.283185307179586\n\n# A note on naming conventions in Julia:\n#\n# * Word separation can be indicated by underscores ('_'), but use of\n#   underscores is discouraged unless the name would be hard to read\n#   otherwise.\n#\n# * Names of Types begin with a capital letter and word separation is shown\n#   with CamelCase instead of underscores.\n#\n# * Names of functions and macros are in lower case, without underscores.\n#\n# * Functions that modify their inputs have names that end in !. These\n#   functions are sometimes called mutating functions or in-place functions.\n\n# Arrays store a sequence of values indexed by integers 1 through n:\na = Int64[] # => 0-element Array{Int64,1}\n\n# 1-dimensional array literals can be written with comma-separated values.\nb = [4, 5, 6] # => 3-element Array{Int64,1}: [4, 5, 6]\nb = [4; 5; 6] # => 3-element Array{Int64,1}: [4, 5, 6]\nb[1]    # => 4\nb[end]  # => 6\n\n# 2-dimensional arrays use space-separated values and semicolon-separated rows.\nmatrix = [1 2; 3 4] # => 2×2 Array{Int64,2}: [1 2; 3 4]\n\n# Arrays of a particular type\nb = Int8[4, 5, 6] # => 3-element Array{Int8,1}: [4, 5, 6]\n\n# Add stuff to the end of a list with push! and append!\n# By convention, the exclamation mark '!' is appended to names of functions\n# that modify their arguments\npush!(a, 1)    # => [1]\npush!(a, 2)    # => [1,2]\npush!(a, 4)    # => [1,2,4]\npush!(a, 3)    # => [1,2,4,3]\nappend!(a, b)  # => [1,2,4,3,4,5,6]\n\n# Remove from the end with pop\npop!(b)  # => 6\nb # => [4,5]\n\n# Let's put it back\npush!(b, 6)  # => [4,5,6]\nb # => [4,5,6]\n\na[1]  # => 1  # remember that Julia indexes from 1, not 0!\n\n# end is a shorthand for the last index. It can be used in any\n# indexing expression\na[end]  # => 6\n\n# we also have popfirst! and pushfirst!\npopfirst!(a)  # => 1 \na # => [2,4,3,4,5,6]\npushfirst!(a, 7)  # => [7,2,4,3,4,5,6]\na # => [7,2,4,3,4,5,6]\n\n# Function names that end in exclamations points indicate that they modify\n# their argument.\narr = [5,4,6]  # => 3-element Array{Int64,1}: [5,4,6]\nsort(arr)      # => [4,5,6]\narr            # => [5,4,6]\nsort!(arr)     # => [4,5,6]\narr            # => [4,5,6]\n\n# Looking out of bounds is a BoundsError\ntry\n    a[0] \n    # => ERROR: BoundsError: attempt to access 7-element Array{Int64,1} at \n    # index [0]\n    # => Stacktrace:\n    # =>  [1] getindex(::Array{Int64,1}, ::Int64) at .\\array.jl:731\n    # =>  [2] top-level scope at none:0\n    # =>  [3] ...\n    # => in expression starting at ...\\LearnJulia.jl:180\n    a[end + 1] \n    # => ERROR: BoundsError: attempt to access 7-element Array{Int64,1} at \n    # index [8]\n    # => Stacktrace:\n    # =>  [1] getindex(::Array{Int64,1}, ::Int64) at .\\array.jl:731\n    # =>  [2] top-level scope at none:0\n    # =>  [3] ...\n    # => in expression starting at ...\\LearnJulia.jl:188\ncatch e\n    println(e)\nend\n\n# Errors list the line and file they came from, even if it's in the standard\n# library. You can look in the folder share/julia inside the julia folder to\n# find these files.\n\n# You can initialize arrays from ranges\na = [1:5;]  # => 5-element Array{Int64,1}: [1,2,3,4,5]\na2 = [1:5]  # => 1-element Array{UnitRange{Int64},1}: [1:5]\n\n# You can look at ranges with slice syntax.\na[1:3]    # => [1, 2, 3]\na[2:end]  # => [2, 3, 4, 5]\n\n# Remove elements from an array by index with splice!\narr = [3,4,5]\nsplice!(arr, 2) # => 4 \narr # => [3,5]\n\n# Concatenate lists with append!\nb = [1,2,3]\nappend!(a, b) # => [1, 2, 3, 4, 5, 1, 2, 3]\na # => [1, 2, 3, 4, 5, 1, 2, 3]\n\n# Check for existence in a list with in\nin(1, a)  # => true\n\n# Examine the length with length\nlength(a)  # => 8\n\n# Tuples are immutable.\ntup = (1, 2, 3)  # => (1,2,3)\ntypeof(tup) # => Tuple{Int64,Int64,Int64}\ntup[1] # => 1\ntry\n    tup[1] = 3  \n    # => ERROR: MethodError: no method matching \n    # setindex!(::Tuple{Int64,Int64,Int64}, ::Int64, ::Int64)\ncatch e\n    println(e)\nend\n\n# Many array functions also work on tuples\nlength(tup) # => 3\ntup[1:2]    # => (1,2)\nin(2, tup)  # => true\n\n# You can unpack tuples into variables\na, b, c = (1, 2, 3)  # => (1,2,3)  \na  # => 1\nb  # => 2\nc  # => 3\n\n# Tuples are created even if you leave out the parentheses\nd, e, f = 4, 5, 6  # => (4,5,6)\nd  # => 4\ne  # => 5\nf  # => 6\n\n# A 1-element tuple is distinct from the value it contains\n(1,) == 1 # => false\n(1) == 1  # => true\n\n# Look how easy it is to swap two values\ne, d = d, e  # => (5,4) \nd  # => 5\ne  # => 4\n\n# Dictionaries store mappings\nemptyDict = Dict()  # => Dict{Any,Any} with 0 entries\n\n# You can create a dictionary using a literal\nfilledDict = Dict(\"one\" => 1, \"two\" => 2, \"three\" => 3)\n# => Dict{String,Int64} with 3 entries:\n# =>  \"two\" => 2, \"one\" => 1, \"three\" => 3\n\n# Look up values with []\nfilledDict[\"one\"]  # => 1\n\n# Get all keys\nkeys(filledDict)\n# => Base.KeySet for a Dict{String,Int64} with 3 entries. Keys:\n# =>  \"two\", \"one\", \"three\"\n# Note - dictionary keys are not sorted or in the order you inserted them.\n\n# Get all values\nvalues(filledDict)\n# => Base.ValueIterator for a Dict{String,Int64} with 3 entries. Values: \n# =>  2, 1, 3\n# Note - Same as above regarding key ordering.\n\n# Check for existence of keys in a dictionary with in, haskey\nin((\"one\" => 1), filledDict)  # => true\nin((\"two\" => 3), filledDict)  # => false\nhaskey(filledDict, \"one\")     # => true\nhaskey(filledDict, 1)         # => false\n\n# Trying to look up a non-existent key will raise an error\ntry\n    filledDict[\"four\"]  # => ERROR: KeyError: key \"four\" not found\ncatch e\n    println(e)\nend\n\n# Use the get method to avoid that error by providing a default value\n# get(dictionary, key, defaultValue)\nget(filledDict, \"one\", 4)   # => 1\nget(filledDict, \"four\", 4)  # => 4\n\n# Use Sets to represent collections of unordered, unique values\nemptySet = Set()  # => Set(Any[])\n# Initialize a set with values\nfilledSet = Set([1, 2, 2, 3, 4])  # => Set([4, 2, 3, 1])\n\n# Add more values to a set\npush!(filledSet, 5)  # => Set([4, 2, 3, 5, 1])\n\n# Check if the values are in the set\nin(2, filledSet)   # => true\nin(10, filledSet)  # => false\n\n# There are functions for set intersection, union, and difference.\notherSet = Set([3, 4, 5, 6])         # => Set([4, 3, 5, 6])\nintersect(filledSet, otherSet)      # => Set([4, 3, 5])\nunion(filledSet, otherSet)          # => Set([4, 2, 3, 5, 6, 1])\nsetdiff(Set([1,2,3,4]), Set([2,3,5])) # => Set([4, 1])\n\n####################################################\n## 3. Control Flow\n####################################################\n\n# Let's make a variable\nsomeVar = 5\n\n# Here is an if statement. Indentation is not meaningful in Julia.\nif someVar > 10\n    println(\"someVar is totally bigger than 10.\")\nelseif someVar < 10    # This elseif clause is optional.\n    println(\"someVar is smaller than 10.\")\nelse                    # The else clause is optional too.\n    println(\"someVar is indeed 10.\")\nend\n# => prints \"some var is smaller than 10\"\n\n# For loops iterate over iterables.\n# Iterable types include Range, Array, Set, Dict, and AbstractString.\nfor animal = [\"dog\", \"cat\", \"mouse\"]\n    println(\"$animal is a mammal\")\n    # You can use $ to interpolate variables or expression into strings.\n    # In this special case, no need for parenthesis: $animal and $(animal) give the same\nend\n# => dog is a mammal\n# => cat is a mammal\n# => mouse is a mammal\n\n# You can use 'in' instead of '='.\nfor animal in [\"dog\", \"cat\", \"mouse\"]\n    println(\"$animal is a mammal\")\nend\n# => dog is a mammal\n# => cat is a mammal\n# => mouse is a mammal\n\nfor pair in Dict(\"dog\" => \"mammal\", \"cat\" => \"mammal\", \"mouse\" => \"mammal\")\n    from, to = pair\n    println(\"$from is a $to\")\nend\n# => mouse is a mammal\n# => cat is a mammal\n# => dog is a mammal\n\nfor (k, v) in Dict(\"dog\" => \"mammal\", \"cat\" => \"mammal\", \"mouse\" => \"mammal\")\n    println(\"$k is a $v\")\nend\n# => mouse is a mammal\n# => cat is a mammal\n# => dog is a mammal\n\n# While loops loop while a condition is true\nlet x = 0\n    while x < 4\n        println(x)\n        x += 1  # Shorthand for in place increment: x = x + 1\n    end\nend\n# => 0\n# => 1\n# => 2\n# => 3\n\n# Handle exceptions with a try/catch block\ntry\n    error(\"help\")\ncatch e\n    println(\"caught it $e\")\nend\n# => caught it ErrorException(\"help\")\n\n####################################################\n## 4. Functions\n####################################################\n\n# The keyword 'function' creates new functions\n# function name(arglist)\n#   body...\n# end\nfunction add(x, y)\n    println(\"x is $x and y is $y\")\n\n    # Functions return the value of their last statement\n    x + y\nend\n\nadd(5, 6)\n# => x is 5 and y is 6\n# => 11\n\n# Compact assignment of functions\nf_add(x, y) = x + y  # => f_add (generic function with 1 method)\nf_add(3, 4)  # => 7\n\n# Function can also return multiple values as tuple\nfn(x, y) = x + y, x - y # => fn (generic function with 1 method)\nfn(3, 4)  # => (7, -1)\n\n# You can define functions that take a variable number of\n# positional arguments\nfunction varargs(args...)\n    return args\n    # use the keyword return to return anywhere in the function\nend\n# => varargs (generic function with 1 method)\n\nvarargs(1, 2, 3)  # => (1,2,3)\n\n# The ... is called a splat.\n# We just used it in a function definition.\n# It can also be used in a function call,\n# where it will splat an Array or Tuple's contents into the argument list.\nadd([5,6]...)  # this is equivalent to add(5,6)\n\nx = (5, 6)  # => (5,6)\nadd(x...)  # this is equivalent to add(5,6)\n\n\n# You can define functions with optional positional arguments\nfunction defaults(a, b, x=5, y=6)\n    return \"$a $b and $x $y\"\nend\n# => defaults (generic function with 3 methods)\n\ndefaults('h', 'g')  # => \"h g and 5 6\"\ndefaults('h', 'g', 'j')  # => \"h g and j 6\"\ndefaults('h', 'g', 'j', 'k')  # => \"h g and j k\"\ntry\n    defaults('h')  # => ERROR: MethodError: no method matching defaults(::Char)\n    defaults()  # => ERROR: MethodError: no method matching defaults()\ncatch e\n    println(e)\nend\n\n# You can define functions that take keyword arguments\nfunction keyword_args(;k1=4, name2=\"hello\")  # note the ;\n    return Dict(\"k1\" => k1, \"name2\" => name2)\nend\n# => keyword_args (generic function with 1 method)\n\nkeyword_args(name2=\"ness\")  # => [\"name2\"=>\"ness\", \"k1\"=>4]\nkeyword_args(k1=\"mine\")     # => [\"name2\"=>\"hello\", \"k1\"=>\"mine\"]\nkeyword_args()              # => [\"name2\"=>\"hello\", \"k1\"=>4]\n\n# You can combine all kinds of arguments in the same function\nfunction all_the_args(normalArg, optionalPositionalArg=2; keywordArg=\"foo\")\n    println(\"normal arg: $normalArg\")\n    println(\"optional arg: $optionalPositionalArg\")\n    println(\"keyword arg: $keywordArg\")\nend\n# => all_the_args (generic function with 2 methods)\n\nall_the_args(1, 3, keywordArg=4)\n# => normal arg: 1\n# => optional arg: 3\n# => keyword arg: 4\n\n# Julia has first class functions\nfunction create_adder(x)\n    adder = function (y)\n        return x + y\n    end\n    return adder\nend\n# => create_adder (generic function with 1 method)\n\n# This is \"stabby lambda syntax\" for creating anonymous functions\n(x -> x > 2)(3)  # => true\n\n# This function is identical to create_adder implementation above.\nfunction create_adder(x)\n    y -> x + y\nend\n# => create_adder (generic function with 1 method)\n\n# You can also name the internal function, if you want\nfunction create_adder(x)\n    function adder(y)\n        x + y\n    end\n    adder\nend\n# => create_adder (generic function with 1 method)\n\nadd_10 = create_adder(10) # => (::getfield(Main, Symbol(\"#adder#11\")){Int64}) \n                          # (generic function with 1 method)\nadd_10(3) # => 13\n\n\n# There are built-in higher order functions\nmap(add_10, [1,2,3])  # => [11, 12, 13]\nfilter(x -> x > 5, [3, 4, 5, 6, 7])  # => [6, 7]\n\n# We can use list comprehensions\n[add_10(i) for i = [1, 2, 3]]   # => [11, 12, 13]\n[add_10(i) for i in [1, 2, 3]]  # => [11, 12, 13]\n[x for x in [3, 4, 5, 6, 7] if x > 5] # => [6, 7]\n\n####################################################\n## 5. Types\n####################################################\n\n# Julia has a type system.\n# Every value has a type; variables do not have types themselves.\n# You can use the `typeof` function to get the type of a value.\ntypeof(5)  # => Int64\n\n# Types are first-class values\ntypeof(Int64)     # => DataType\ntypeof(DataType)  # => DataType\n# DataType is the type that represents types, including itself.\n\n# Types are used for documentation, optimizations, and dispatch.\n# They are not statically checked.\n\n# Users can define types\n# They are like records or structs in other languages.\n# New types are defined using the `struct` keyword.\n\n# struct Name\n#   field::OptionalType\n#   ...\n# end\nstruct Tiger\n    taillength::Float64\n    coatcolor  # not including a type annotation is the same as `::Any`\nend\n\n# The default constructor's arguments are the properties\n# of the type, in the order they are listed in the definition\ntigger = Tiger(3.5, \"orange\")  # => Tiger(3.5,\"orange\")\n\n# The type doubles as the constructor function for values of that type\nsherekhan = typeof(tigger)(5.6, \"fire\")  # => Tiger(5.6,\"fire\")\n\n# These struct-style types are called concrete types\n# They can be instantiated, but cannot have subtypes.\n# The other kind of types is abstract types.\n\n# abstract Name\nabstract type Cat end  # just a name and point in the type hierarchy\n\n# Abstract types cannot be instantiated, but can have subtypes.\n# For example, Number is an abstract type\nsubtypes(Number)  # => 2-element Array{Any,1}:\n                  # =>  Complex\n                  # =>  Real\nsubtypes(Cat)  # => 0-element Array{Any,1}\n\n# AbstractString, as the name implies, is also an abstract type\nsubtypes(AbstractString)  # => 4-element Array{Any,1}:\n                          # =>  String\n                          # =>  SubString\n                          # =>  SubstitutionString\n                          # =>  Test.GenericString\n\n# Every type has a super type; use the `supertype` function to get it.\ntypeof(5) # => Int64\nsupertype(Int64)    # => Signed\nsupertype(Signed)   # => Integer\nsupertype(Integer)  # => Real\nsupertype(Real)     # => Number\nsupertype(Number)   # => Any\nsupertype(supertype(Signed))  # => Real\nsupertype(Any)      # => Any\n# All of these type, except for Int64, are abstract.\ntypeof(\"fire\")      # => String\nsupertype(String)   # => AbstractString\n# Likewise here with String\nsupertype(SubString)  # => AbstractString\n\n# <: is the subtyping operator\nstruct Lion <: Cat  # Lion is a subtype of Cat\n    maneColor\n    roar::AbstractString\nend\n\n# You can define more constructors for your type\n# Just define a function of the same name as the type\n# and call an existing constructor to get a value of the correct type\nLion(roar::AbstractString) = Lion(\"green\", roar)\n# This is an outer constructor because it's outside the type definition\n\nstruct Panther <: Cat  # Panther is also a subtype of Cat\n    eyeColor\n    Panther() = new(\"green\")\n    # Panthers will only have this constructor, and no default constructor.\nend\n# Using inner constructors, like Panther does, gives you control\n# over how values of the type can be created.\n# When possible, you should use outer constructors rather than inner ones.\n\n####################################################\n## 6. Multiple-Dispatch\n####################################################\n\n# In Julia, all named functions are generic functions\n# This means that they are built up from many small methods\n# Each constructor for Lion is a method of the generic function Lion.\n\n# For a non-constructor example, let's make a function meow:\n\n# Definitions for Lion, Panther, Tiger\nfunction meow(animal::Lion)\n    animal.roar  # access type properties using dot notation\nend\n\nfunction meow(animal::Panther)\n    \"grrr\"\nend\n\nfunction meow(animal::Tiger)\n    \"rawwwr\"\nend\n\n# Testing the meow function\nmeow(tigger)  # => \"rawwwr\"\nmeow(Lion(\"brown\", \"ROAAR\"))  # => \"ROAAR\"\nmeow(Panther()) # => \"grrr\"\n\n# Review the local type hierarchy\nTiger   <: Cat  # => false\nLion    <: Cat  # => true\nPanther <: Cat  # => true\n\n# Defining a function that takes Cats\nfunction pet_cat(cat::Cat)\n    println(\"The cat says $(meow(cat))\")\nend\n# => pet_cat (generic function with 1 method)\n\npet_cat(Lion(\"42\")) # => The cat says 42\ntry\n    pet_cat(tigger) # => ERROR: MethodError: no method matching pet_cat(::Tiger)\ncatch e\n    println(e)\nend\n\n# In OO languages, single dispatch is common;\n# this means that the method is picked based on the type of the first argument.\n# In Julia, all of the argument types contribute to selecting the best method.\n\n# Let's define a function with more arguments, so we can see the difference\nfunction fight(t::Tiger, c::Cat)\n    println(\"The $(t.coatcolor) tiger wins!\")\nend\n# => fight (generic function with 1 method)\n\nfight(tigger, Panther())  # => The orange tiger wins!\nfight(tigger, Lion(\"ROAR\")) # => The orange tiger wins!\n\n# Let's change the behavior when the Cat is specifically a Lion\nfight(t::Tiger, l::Lion) = println(\"The $(l.maneColor)-maned lion wins!\")\n# => fight (generic function with 2 methods)\n\nfight(tigger, Panther())  # => The orange tiger wins!\nfight(tigger, Lion(\"ROAR\")) # => The green-maned lion wins!\n\n# We don't need a Tiger in order to fight\nfight(l::Lion, c::Cat) = println(\"The victorious cat says $(meow(c))\")\n# => fight (generic function with 3 methods)\n\nfight(Lion(\"balooga!\"), Panther())  # => The victorious cat says grrr\ntry\n    fight(Panther(), Lion(\"RAWR\"))  \n    # => ERROR: MethodError: no method matching fight(::Panther, ::Lion)\n    # => Closest candidates are:\n    # =>   fight(::Tiger, ::Lion) at ...\n    # =>   fight(::Tiger, ::Cat) at ...\n    # =>   fight(::Lion, ::Cat) at ...\n    # => ...\ncatch e\n    println(e)\nend\n\n# Also let the cat go first\nfight(c::Cat, l::Lion) = println(\"The cat beats the Lion\")\n# => fight (generic function with 4 methods)\n\n# This warning is because it's unclear which fight will be called in:\ntry\n    fight(Lion(\"RAR\"), Lion(\"brown\", \"rarrr\"))\n    # => ERROR: MethodError: fight(::Lion, ::Lion) is ambiguous. Candidates:\n    # =>   fight(c::Cat, l::Lion) in Main at ...\n    # =>   fight(l::Lion, c::Cat) in Main at ...\n    # => Possible fix, define\n    # =>   fight(::Lion, ::Lion)\n    # => ...\ncatch e\n    println(e)\nend\n# The result may be different in other versions of Julia\n\nfight(l::Lion, l2::Lion) = println(\"The lions come to a tie\") \n# => fight (generic function with 5 methods)\nfight(Lion(\"RAR\"), Lion(\"brown\", \"rarrr\"))  # => The lions come to a tie\n\n\n# Under the hood\n# You can take a look at the llvm  and the assembly code generated.\n\nsquare_area(l) = l * l  # square_area (generic function with 1 method)\n\nsquare_area(5)  # => 25\n\n# What happens when we feed square_area an integer?\ncode_native(square_area, (Int32,), syntax = :intel)\n\t#         .text\n\t# ; Function square_area {\n\t# ; Location: REPL[116]:1       # Prologue\n\t#         push    rbp\n\t#         mov     rbp, rsp\n\t# ; Function *; {\n\t# ; Location: int.jl:54\n\t#         imul    ecx, ecx      # Square l and store the result in ECX\n\t# ;}\n\t#         mov     eax, ecx\n\t#         pop     rbp           # Restore old base pointer\n\t#         ret                   # Result will still be in EAX\n\t#         nop     dword ptr [rax + rax]\n\t# ;}\n\ncode_native(square_area, (Float32,), syntax = :intel)\n    #         .text\n    # ; Function square_area {\n    # ; Location: REPL[116]:1\n    #         push    rbp\n    #         mov     rbp, rsp\n    # ; Function *; {\n    # ; Location: float.jl:398\n    #         vmulss  xmm0, xmm0, xmm0  # Scalar single precision multiply (AVX)\n    # ;}\n    #         pop     rbp\n    #         ret\n    #         nop     word ptr [rax + rax]\n    # ;}\n\ncode_native(square_area, (Float64,), syntax = :intel)\n    #         .text\n    # ; Function square_area {\n    # ; Location: REPL[116]:1\n    #         push    rbp\n    #         mov     rbp, rsp\n    # ; Function *; {\n    # ; Location: float.jl:399\n    #         vmulsd  xmm0, xmm0, xmm0  # Scalar double precision multiply (AVX)\n    # ;}\n    #         pop     rbp\n    #         ret\n    #         nop     word ptr [rax + rax]\n    # ;}\n\n# Note that julia will use floating point instructions if any of the\n# arguments are floats.\n# Let's calculate the area of a circle\ncircle_area(r) = pi * r * r     # circle_area (generic function with 1 method)\ncircle_area(5)  # 78.53981633974483\n\ncode_native(circle_area, (Int32,), syntax = :intel)\n    #         .text\n    # ; Function circle_area {\n    # ; Location: REPL[121]:1\n    #         push    rbp\n    #         mov     rbp, rsp\n    # ; Function *; {\n    # ; Location: operators.jl:502\n    # ; Function *; {\n    # ; Location: promotion.jl:314\n    # ; Function promote; {\n    # ; Location: promotion.jl:284\n    # ; Function _promote; {\n    # ; Location: promotion.jl:261\n    # ; Function convert; {\n    # ; Location: number.jl:7\n    # ; Function Type; {\n    # ; Location: float.jl:60\n    #         vcvtsi2sd       xmm0, xmm0, ecx     # Load integer (r) from memory\n    #         movabs  rax, 497710928              # Load pi\n    # ;}}}}}\n    # ; Function *; {\n    # ; Location: float.jl:399\n    #         vmulsd  xmm1, xmm0, qword ptr [rax] # pi * r\n    #         vmulsd  xmm0, xmm1, xmm0            # (pi * r) * r\n    # ;}}\n    #         pop     rbp\n    #         ret\n    #         nop     dword ptr [rax]\n    # ;}\n\ncode_native(circle_area, (Float64,), syntax = :intel)\n    #         .text\n    # ; Function circle_area {\n    # ; Location: REPL[121]:1\n    #         push    rbp\n    #         mov     rbp, rsp\n    #         movabs  rax, 497711048\n    # ; Function *; {\n    # ; Location: operators.jl:502\n    # ; Function *; {\n    # ; Location: promotion.jl:314\n    # ; Function *; {\n    # ; Location: float.jl:399\n    #         vmulsd  xmm1, xmm0, qword ptr [rax]\n    # ;}}}\n    # ; Function *; {\n    # ; Location: float.jl:399\n    #         vmulsd  xmm0, xmm1, xmm0\n    # ;}\n    #         pop     rbp\n    #         ret\n    #         nop     dword ptr [rax + rax]\n    # ;}\n```\n\n## Further Reading\n\nYou can get a lot more detail from the [Julia Documentation](https://docs.julialang.org/)\n\nThe best place to get help with Julia is the (very friendly) [Discourse forum](https://discourse.julialang.org/)."
