
package pages

const Crystal = "```crystal\n\n# This is a comment\n\n# Everything is an object\nnil.class  #=> Nil\n100.class  #=> Int32\ntrue.class #=> Bool\n\n# Falsey values are: nil, false and null pointers\n!nil   #=> true  : Bool\n!false #=> true  : Bool\n!0     #=> false : Bool\n\n# Integers\n\n1.class #=> Int32\n\n# Five signed integer types\n1_i8.class   #=> Int8\n1_i16.class  #=> Int16\n1_i32.class  #=> Int32\n1_i64.class  #=> Int64\n1_i128.class #=> Int128\n\n# Five unsigned integer types\n1_u8.class   #=> UInt8\n1_u16.class  #=> UInt16\n1_u32.class  #=> UInt32\n1_u64.class  #=> UInt64\n1_u128.class #=> UInt128\n\n2147483648.class          #=> Int64\n9223372036854775808.class #=> UInt64\n\n# Binary numbers\n0b1101 #=> 13 : Int32\n\n# Octal numbers\n0o123 #=> 83 : Int32\n\n# Hexadecimal numbers\n0xFE012D #=> 16646445 : Int32\n0xfe012d #=> 16646445 : Int32\n\n# Floats\n\n1.0.class #=> Float64\n\n# There are two floating point types\n1.0_f32.class #=> Float32\n1_f32.class   #=> Float32\n\n1e10.class    #=> Float64\n1.5e10.class  #=> Float64\n1.5e-7.class  #=> Float64\n\n# Chars\n\n'a'.class #=> Char\n\n# Octal codepoint\n'\\101' #=> 'A' : Char\n\n# Unicode codepoint\n'\\u0041' #=> 'A' : Char\n\n# Strings\n\n\"s\".class #=> String\n\n# Strings are immutable\ns = \"hello, \"  #=> \"hello, \"        : String\ns.object_id    #=> 134667712        : UInt64\ns += \"Crystal\" #=> \"hello, Crystal\" : String\ns.object_id    #=> 142528472        : UInt64\n\n# Supports interpolation\n\"sum = #{1 + 2}\" #=> \"sum = 3\" : String\n\n# Multiline string\n\"This is\n   multiline string\"\n\n# String with double quotes\n%(hello \"world\") #=> \"hello \\\"world\\\"\"\n\n# Symbols\n# Immutable, reusable constants represented internally as Int32 integer value.\n# They're often used instead of strings to efficiently convey specific,\n# meaningful values\n\n:symbol.class #=> Symbol\n\nsentence = :question?     # :\"question?\" : Symbol\n\nsentence == :question?    #=> true  : Bool\nsentence == :exclamation! #=> false : Bool\nsentence == \"question?\"   #=> false : Bool\n\n# Arrays\n\n[1, 2, 3].class         #=> Array(Int32)\n[1, \"hello\", 'x'].class #=> Array(Int32 | String | Char)\n\n# Empty arrays should specify a type\n[]               # Syntax error: for empty arrays use '[] of ElementType'\n[] of Int32      #=> [] : Array(Int32)\nArray(Int32).new #=> [] : Array(Int32)\n\n# Arrays can be indexed\narray = [1, 2, 3, 4, 5] #=> [1, 2, 3, 4, 5] : Array(Int32)\narray[0]                #=> 1               : Int32\narray[10]               # raises IndexError\narray[-6]               # raises IndexError\narray[10]?              #=> nil             : (Int32 | Nil)\narray[-6]?              #=> nil             : (Int32 | Nil)\n\n# From the end\narray[-1] #=> 5\n\n# With a start index and size\narray[2, 3] #=> [3, 4, 5]\n\n# Or with range\narray[1..3] #=> [2, 3, 4]\n\n# Add to an array\narray << 6  #=> [1, 2, 3, 4, 5, 6]\n\n# Remove from the end of the array\narray.pop #=> 6\narray     #=> [1, 2, 3, 4, 5]\n\n# Remove from the beginning of the array\narray.shift #=> 1\narray       #=> [2, 3, 4, 5]\n\n# Check if an item exists in an array\narray.includes? 3 #=> true\n\n# Special syntax for an array of string and an array of symbols\n%w(one two three) #=> [\"one\", \"two\", \"three\"] : Array(String)\n%i(one two three) #=> [:one, :two, :three]    : Array(Symbol)\n\n# There is a special array syntax with other types too, as long as\n# they define a .new and a #<< method\nset = Set{1, 2, 3} #=> [1, 2, 3]\nset.class          #=> Set(Int32)\n\n# The above is equivalent to\nset = Set(typeof(1, 2, 3)).new\nset << 1\nset << 2\nset << 3\n\n# Hashes\n\n{1 => 2, 3 => 4}.class   #=> Hash(Int32, Int32)\n{1 => 2, 'a' => 3}.class #=> Hash(Int32 | Char, Int32)\n\n# Empty hashes should specify a type\n{}                     # Syntax error\n{} of Int32 => Int32   # {}\nHash(Int32, Int32).new # {}\n\n# Hashes can be quickly looked up by key\nhash = {\"color\" => \"green\", \"number\" => 5}\nhash[\"color\"]        #=> \"green\"\nhash[\"no_such_key\"]  #=> Missing hash key: \"no_such_key\" (KeyError)\nhash[\"no_such_key\"]? #=> nil\n\n# Check existence of keys hash\nhash.has_key? \"color\" #=> true\n\n# Special notation for symbol and string keys\n{key1: 'a', key2: 'b'}     # {:key1 => 'a', :key2 => 'b'}\n{\"key1\": 'a', \"key2\": 'b'} # {\"key1\" => 'a', \"key2\" => 'b'}\n\n# Special hash literal syntax with other types too, as long as\n# they define a .new and a #[]= methods\nclass MyType\n  def []=(key, value)\n    puts \"do stuff\"\n  end\nend\n\nMyType{\"foo\" => \"bar\"}\n\n# The above is equivalent to\ntmp = MyType.new\ntmp[\"foo\"] = \"bar\"\ntmp\n\n# Ranges\n\n1..10                  #=> Range(Int32, Int32)\nRange.new(1, 10).class #=> Range(Int32, Int32)\n\n# Can be inclusive or exclusive\n(3..5).to_a  #=> [3, 4, 5]\n(3...5).to_a #=> [3, 4]\n\n# Check whether range includes the given value or not\n(1..8).includes? 2 #=> true\n\n# Tuples are a fixed-size, immutable, stack-allocated sequence of values of\n# possibly different types.\n{1, \"hello\", 'x'}.class #=> Tuple(Int32, String, Char)\n\n# Access tuple's value by its index\ntuple = {:key1, :key2}\ntuple[1] #=> :key2\ntuple[2] #=> syntax error : Index out of bound\n\n# Can be expanded into multiple variables\na, b, c = {:a, 'b', \"c\"}\na #=> :a\nb #=> 'b'\nc #=> \"c\"\n\n# Procs represent a function pointer with an optional context (the closure data)\n# It is typically created with a proc literal\nproc = ->(x : Int32) { x.to_s }\nproc.class # Proc(Int32, String)\n# Or using the new method\nProc(Int32, String).new { |x| x.to_s }\n\n# Invoke proc with call method\nproc.call 10 #=> \"10\"\n\n# Control statements\n\nif true\n  \"if statement\"\nelsif false\n  \"else-if, optional\"\nelse\n  \"else, also optional\"\nend\n\nputs \"if as a suffix\" if true\n\n# If as an expression\na = if 2 > 1\n      3\n    else\n      4\n    end\n\na #=> 3\n\n# Ternary if\na = 1 > 2 ? 3 : 4 #=> 4\n\n# Case statement\ncmd = \"move\"\n\naction = case cmd\n  when \"create\"\n    \"Creating...\"\n  when \"copy\"\n    \"Copying...\"\n  when \"move\"\n    \"Moving...\"\n  when \"delete\"\n    \"Deleting...\"\nend\n\naction #=> \"Moving...\"\n\n# Loops\nindex = 0\nwhile index <= 3\n  puts \"Index: #{index}\"\n  index += 1\nend\n# Index: 0\n# Index: 1\n# Index: 2\n# Index: 3\n\nindex = 0\nuntil index > 3\n  puts \"Index: #{index}\"\n  index += 1\nend\n# Index: 0\n# Index: 1\n# Index: 2\n# Index: 3\n\n# But the preferable way is to use each\n(1..3).each do |index|\n  puts \"Index: #{index}\"\nend\n# Index: 1\n# Index: 2\n# Index: 3\n\n# Variable's type depends on the type of the expression\n# in control statements\nif a < 3\n  a = \"hello\"\nelse\n  a = true\nend\ntypeof a #=> (Bool | String)\n\nif a && b\n  # here both a and b are guaranteed not to be Nil\nend\n\nif a.is_a? String\n  a.class #=> String\nend\n\n# Functions\n\ndef double(x)\n  x * 2\nend\n\n# Functions (and all blocks) implicitly return the value of the last statement\ndouble(2) #=> 4\n\n# Parentheses are optional where the call is unambiguous\ndouble 3 #=> 6\n\ndouble double 3 #=> 12\n\ndef sum(x, y)\n  x + y\nend\n\n# Method arguments are separated by a comma\nsum 3, 4 #=> 7\n\nsum sum(3, 4), 5 #=> 12\n\n# yield\n# All methods have an implicit, optional block parameter\n# it can be called with the 'yield' keyword\n\ndef surround\n  puts '{'\n  yield\n  puts '}'\nend\n\nsurround { puts \"hello world\" }\n\n# {\n# hello world\n# }\n\n\n# You can pass a block to a function\n# \"&\" marks a reference to a passed block\ndef guests(&block)\n  block.call \"some_argument\"\nend\n\n# You can pass a list of arguments, which will be converted into an array\n# That's what splat operator (\"*\") is for\ndef guests(*array)\n  array.each { |guest| puts guest }\nend\n\n# If a method returns an array, you can use destructuring assignment\ndef foods\n    [\"pancake\", \"sandwich\", \"quesadilla\"]\nend\nbreakfast, lunch, dinner = foods\nbreakfast #=> \"pancake\"\ndinner    #=> \"quesadilla\"\n\n# By convention, all methods that return booleans end with a question mark\n5.even? # false\n5.odd?  # true\n\n# And if a method ends with an exclamation mark, it does something destructive\n# like mutate the receiver. Some methods have a ! version to make a change, and\n# a non-! version to just return a new changed version\ncompany_name = \"Dunder Mifflin\"\ncompany_name.gsub \"Dunder\", \"Donald\"  #=> \"Donald Mifflin\"\ncompany_name  #=> \"Dunder Mifflin\"\ncompany_name.gsub! \"Dunder\", \"Donald\"\ncompany_name  #=> \"Donald Mifflin\"\n\n\n# Define a class with the class keyword\nclass Human\n\n  # A class variable. It is shared by all instances of this class.\n  @@species = \"H. sapiens\"\n\n  # type of name is String\n  @name : String\n\n  # Basic initializer\n  # Assign the argument to the \"name\" instance variable for the instance\n  # If no age given, we will fall back to the default in the arguments list.\n  def initialize(@name, @age = 0)\n  end\n\n  # Basic setter method\n  def name=(name)\n    @name = name\n  end\n\n  # Basic getter method\n  def name\n    @name\n  end\n\n  # The above functionality can be encapsulated using the propery method as follows\n  property :name\n\n  # Getter/setter methods can also be created individually like this\n  getter :name\n  setter :name\n\n  # A class method uses self to distinguish from instance methods.\n  # It can only be called on the class, not an instance.\n  def self.say(msg)\n    puts msg\n  end\n\n  def species\n    @@species\n  end\nend\n\n\n# Instantiate a class\njim = Human.new(\"Jim Halpert\")\n\ndwight = Human.new(\"Dwight K. Schrute\")\n\n# Let's call a couple of methods\njim.species #=> \"H. sapiens\"\njim.name #=> \"Jim Halpert\"\njim.name = \"Jim Halpert II\" #=> \"Jim Halpert II\"\njim.name #=> \"Jim Halpert II\"\ndwight.species #=> \"H. sapiens\"\ndwight.name #=> \"Dwight K. Schrute\"\n\n# Call the class method\nHuman.say(\"Hi\") #=> print Hi and returns nil\n\n# Variables that start with @ have instance scope\nclass TestClass\n  @var = \"I'm an instance var\"\nend\n\n# Variables that start with @@ have class scope\nclass TestClass\n  @@var = \"I'm a class var\"\nend\n# Variables that start with a capital letter are constants\nVar = \"I'm a constant\"\nVar = \"can't be updated\" # Already initialized constant Var\n\n# Class is also an object in crystal. So class can have instance variables.\n# Class variable is shared among the class and all of its descendants.\n\n# base class\nclass Human\n  @@foo = 0\n\n  def self.foo\n    @@foo\n  end\n\n  def self.foo=(value)\n    @@foo = value\n  end\nend\n\n# derived class\nclass Worker < Human\nend\n\nHuman.foo   #=> 0\nWorker.foo  #=> 0\n\nHuman.foo = 2 #=> 2\nWorker.foo    #=> 0\n\nWorker.foo = 3 #=> 3\nHuman.foo   #=> 2\nWorker.foo  #=> 3\n\nmodule ModuleExample\n  def foo\n    \"foo\"\n  end\nend\n\n# Including modules binds their methods to the class instances\n# Extending modules binds their methods to the class itself\n\nclass Person\n  include ModuleExample\nend\n\nclass Book\n  extend ModuleExample\nend\n\nPerson.foo     # => undefined method 'foo' for Person:Class\nPerson.new.foo # => 'foo'\nBook.foo       # => 'foo'\nBook.new.foo   # => undefined method 'foo' for Book\n\n\n# Exception handling\n\n# Define new exception\nclass MyException < Exception\nend\n\n# Define another exception\nclass MyAnotherException < Exception; end\n\nex = begin\n   raise MyException.new\nrescue ex1 : IndexError\n  \"ex1\"\nrescue ex2 : MyException | MyAnotherException\n  \"ex2\"\nrescue ex3 : Exception\n  \"ex3\"\nrescue ex4 # catch any kind of exception\n  \"ex4\"\nend\n\nex #=> \"ex2\"\n\n```\n\n## Additional resources\n\n- [Official Documentation](https://crystal-lang.org/)"