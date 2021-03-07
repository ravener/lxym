
package pages

const Citron = "```ruby\n# Comments start with a '#'\n# All comments encompass a single line\n\n###########################################\n## 1. Primitive Data types and Operators\n###########################################\n\n# You have numbers\n3. # 3\n\n# Numbers are all doubles in interpreted mode\n\n# Mathematical operator precedence is not respected.\n# binary 'operators' are evaluated in ltr order\n1 + 1. # 2\n8 - 4. # 4\n10 + 2 * 3. # 36\n\n# Division is always floating division\n35 / 2 # 17.5.\n\n# Integer division is non-trivial, you may use floor\n(35 / 2) floor # 17.\n\n# Booleans are primitives\nTrue.\nFalse.\n\n# Boolean messages\nTrue not. # False\nFalse not. # True\n1 = 1. # True\n1 !=: 1. # False\n1 < 10. # True\n\n# Here, `not` is a unary message to the object `Boolean`\n# Messages are comparable to instance method calls\n# And they have three different forms:\n#   1. Unary messages: Length > 1, and they take no arguments:\n        False not.\n#   2. Binary Messages: Length = 1, and they take a single argument:\n        False & True.\n#   3. Keyword messages: must have at least one ':', they take as many arguments\n#      as they have `:` s\n        False either: 1 or: 2. # 2\n\n# Strings\n'This is a string'.\n'There are no character types exposed to the user'.\n# \"You cannot use double quotes for strings\" <- Error\n\n# Strins can be summed\n'Hello, ' + 'World!'. # 'Hello, World!'\n\n# Strings allow access to their characters\n'This is a beautiful string' at: 0. # 'T'\n\n###########################################\n## intermission: Basic Assignment\n###########################################\n\n# You may assign values to the current scope:\nvar name is value. # assignes `value` into `name`\n\n# You may also assign values into the current object's namespace\nmy name is value. # assigns `value` into the current object's `name` property\n\n# Please note that these names are checked at compile (read parse if in interpreted mode) time\n# but you may treat them as dynamic assignments anyway\n\n###########################################\n## 2. Lists(Arrays?) and Tuples\n###########################################\n\n# Arrays are allowed to have multiple types\nArray new < 1 ; 2 ; 'string' ; Nil. # Array new < 1 ; 2 ; 'string' ; Nil\n\n# Tuples act like arrays, but are immutable.\n# Any shenanigans degrade them to arrays, however\n[1, 2, 'string']. # [1, 2, 'string']\n\n# They can interoperate with arrays\n[1, 'string'] + (Array new < 'wat'). # Array new < 1 ; 'string' ; 'wat'\n\n# Indexing into them\n[1, 2, 3] at: 1. # 2\n\n# Some array operations\nvar arr is Array new < 1 ; 2 ; 3.\n\narr head. # 1\narr tail. # Array new < 2 ; 3.\narr init. # Array new < 1 ; 2.\narr last. # 3\narr push: 4. # Array new < 1 ; 2 ; 3 ; 4.\narr pop. # 4\narr pop: 1. # 2, `arr` is rebound to Array new < 1 ; 3.\n\n# List comprehensions\n[x * 2 + y,, arr, arr + [4, 5],, x > 1]. # Array ← 7 ; 9 ; 10 ; 11\n# fresh variable names are bound as they are encountered,\n# so `x` is bound to the values in `arr`\n# and `y` is bound to the values in `arr + [4, 5]`\n#\n# The general format is: [expr,, bindings*,, predicates*]\n\n\n####################################\n## 3. Functions\n####################################\n\n# A simple function that takes two variables\nvar add is {:a:b ^a + b.}.\n\n# this function will resolve all its names except the formal arguments\n# in the context it is called in.\n\n# Using the function\nadd applyTo: 3 and: 5. # 8\nadd applyAll: [3, 5]. # 8\n\n# Also a (customizable -- more on this later) pseudo-operator allows for a shorthand\n# of function calls\n# By default it is REF[args]\n\nadd[3, 5]. # 8\n\n# To customize this behaviour, you may simply use a compiler pragma:\n#:callShorthand ()\n\n# And then you may use the specified operator.\n# Note that the allowed 'operator' can only be made of any of these: []{}()\n# And you may mix-and-match (why would anyone do that?)\n\nadd(3, 5). # 8\n\n# You may also use functions as operators in the following way:\n\n3 `add` 5. # 8\n# This call binds as such: add[(3), 5]\n# because the default fixity is left, and the default precedance is 1\n\n# You may change the precedence/fixity of this operator with a pragma\n#:declare infixr 1 add\n\n3 `add` 5. # 8\n# now this binds as such: add[3, (5)].\n\n# There is another form of functions too\n# So far, the functions were resolved in a dynamic fashion\n# But a lexically scoped block is also possible\nvar sillyAdd is {\\:x:y add[x,y].}.\n\n# In these blocks, you are not allowed to declare new variables\n# Except with the use of Object::'letEqual:in:`\n# And the last expression is implicitly returned.\n\n# You may also use a shorthand for lambda expressions\nvar mul is \\:x:y x * y.\n\n# These capture the named bindings that are not present in their\n# formal parameters, and retain them. (by ref)\n\n###########################################\n## 5. Control Flow\n###########################################\n\n# inline conditional-expressions\nvar citron is 1 = 1 either: 'awesome' or: 'awful'. # citron is 'awesome'\n\n# multiple lines is fine too\nvar citron is 1 = 1\n    either: 'awesome'\n    or:     'awful'.\n\n# looping\n10 times: {:x\n    Pen writeln: x.\n}. # 10. -- side effect: 10 lines in stdout, with numbers 0 through 9 in them\n\n# Citron properly supports tail-call recursion in lexically scoped blocks\n# So use those to your heart's desire\n\n# mapping most data structures is as simple as `fmap:`\n[1, 2, 3, 4] fmap: \\:x x + 1. # [2, 3, 4, 5]\n\n# You can use `foldl:accumulator:` to fold a list/tuple\n[1, 2, 3, 4] foldl: (\\:acc:x acc * 2 + x) accumulator: 4. # 90\n\n# That expression is the same as\n(2 * (2 * (2 * (2 * 4 + 1) + 2) + 3) + 4)\n\n###################################\n## 6. IO\n###################################\n\n# IO is quite simple\n# With `Pen` being used for console output\n# and Program::'input' and Program::'waitForInput' being used for console input\n\nPen writeln: 'Hello, ocean!' # prints 'Hello, ocean!\\n' to the terminal\n\nPen writeln: Program waitForInput. # reads a line and prints it back\n```"
