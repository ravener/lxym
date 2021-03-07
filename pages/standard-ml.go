
package pages

const Standard_Ml = "Standard ML is a functional programming language with type inference and some\nside-effects.  Some of the hard parts of learning Standard ML are: Recursion,\npattern matching, type inference (guessing the right types but never allowing\nimplicit type conversion). Standard ML is distinguished from Haskell by including\nreferences, allowing variables to be updated.\n\n```ocaml\n(* Comments in Standard ML begin with (* and end with *).  Comments can be\n   nested which means that all (* tags must end with a *) tag.  This comment,\n   for example, contains two nested comments. *)\n\n(* A Standard ML program consists of declarations, e.g. value declarations: *)\nval rent = 1200\nval phone_no = 5551337\nval pi = 3.14159\nval negative_number = ~15  (* Yeah, unary minus uses the 'tilde' symbol *)\n\n(* Optionally, you can explicitly declare types. This is not necessary as\n   ML will automatically figure out the types of your values. *)\nval diameter = 7926 : int\nval e = 2.718 : real\nval name = \"Bobby\" : string\n\n(* And just as importantly, functions: *)\nfun is_large(x : int) = if x > 37 then true else false\n\n(* Floating-point numbers are called \"reals\". *)\nval tau = 2.0 * pi         (* You can multiply two reals *)\nval twice_rent = 2 * rent  (* You can multiply two ints *)\n(* val meh = 1.25 * 10 *)  (* But you can't multiply an int and a real *)\nval yeh = 1.25 * (Real.fromInt 10) (* ...unless you explicitly convert\n                                      one or the other *)\n\n(* +, - and * are overloaded so they work for both int and real. *)\n(* The same cannot be said for division which has separate operators: *)\nval real_division = 14.0 / 4.0  (* gives 3.5 *)\nval int_division  = 14 div 4    (* gives 3, rounding down *)\nval int_remainder = 14 mod 4    (* gives 2, since 3*4 = 12 *)\n\n(* ~ is actually sometimes a function (e.g. when put in front of variables) *)\nval negative_rent = ~(rent)  (* Would also have worked if rent were a \"real\" *)\n\n(* There are also booleans and boolean operators *)\nval got_milk = true\nval got_bread = false\nval has_breakfast = got_milk andalso got_bread  (* 'andalso' is the operator *)\nval has_something = got_milk orelse got_bread   (* 'orelse' is the operator *)\nval is_sad = not(has_something)                 (* not is a function *)\n\n(* Many values can be compared using equality operators: = and <> *)\nval pays_same_rent = (rent = 1300)  (* false *)\nval is_wrong_phone_no = (phone_no <> 5551337)  (* false *)\n\n(* The operator <> is what most other languages call !=. *)\n(* 'andalso' and 'orelse' are called && and || in many other languages. *)\n\n(* Actually, most of the parentheses above are unnecessary.  Here are some\n   different ways to say some of the things mentioned above: *)\nfun is_large x = x > 37  (* The parens above were necessary because of ': int' *)\nval is_sad = not has_something\nval pays_same_rent = rent = 1300  (* Looks confusing, but works *)\nval is_wrong_phone_no = phone_no <> 5551337\nval negative_rent = ~rent  (* ~ rent (notice the space) would also work *)\n\n(* Parentheses are mostly necessary when grouping things: *)\nval some_answer = is_large (5 + 5)      (* Without parens, this would break! *)\n(* val some_answer = is_large 5 + 5 *)  (* Read as: (is_large 5) + 5. Bad! *)\n\n\n(* Besides booleans, ints and reals, Standard ML also has chars and strings: *)\nval foo = \"Hello, World!\\n\"  (* The \\n is the escape sequence for linebreaks *)\nval one_letter = #\"a\"        (* That funky syntax is just one character, a *)\n\nval combined = \"Hello \" ^ \"there, \" ^ \"fellow!\\n\"  (* Concatenate strings *)\n\nval _ = print foo       (* You can print things. We are not interested in the *)\nval _ = print combined  (* result of this computation, so we throw it away. *)\n(* val _ = print one_letter *)  (* Only strings can be printed this way *)\n\n\nval bar = [ #\"H\", #\"e\", #\"l\", #\"l\", #\"o\" ]  (* SML also has lists! *)\n(* val _ = print bar *)  (* Lists are unfortunately not the same as strings *)\n\n(* Fortunately they can be converted.  String is a library and implode and size\n   are functions available in that library that take strings as argument. *)\nval bob = String.implode bar          (* gives \"Hello\" *)\nval bob_char_count = String.size bob  (* gives 5 *)\nval _ = print (bob ^ \"\\n\")            (* For good measure, add a linebreak *)\n\n(* You can have lists of any kind *)\nval numbers = [1, 3, 3, 7, 229, 230, 248]  (* : int list *)\nval names = [ \"Fred\", \"Jane\", \"Alice\" ]    (* : string list *)\n\n(* Even lists of lists of things *)\nval groups = [ [ \"Alice\", \"Bob\" ],\n               [ \"Huey\", \"Dewey\", \"Louie\" ],\n               [ \"Bonnie\", \"Clyde\" ] ]     (* : string list list *)\n\nval number_count = List.length numbers     (* gives 7 *)\n\n(* You can put single values in front of lists of the same kind using\n   the :: operator, called \"the cons operator\" (known from Lisp). *)\nval more_numbers = 13 :: numbers  (* gives [13, 1, 3, 3, 7, ...] *)\nval more_groups  = [\"Batman\",\"Superman\"] :: groups\n\n(* Lists of the same kind can be appended using the @ (\"append\") operator *)\nval guest_list = [ \"Mom\", \"Dad\" ] @ [ \"Aunt\", \"Uncle\" ]\n\n(* This could have been done with the \"cons\" operator.  It is tricky because the\n   left-hand-side must be an element whereas the right-hand-side must be a list\n   of those elements. *)\nval guest_list = \"Mom\" :: \"Dad\" :: [ \"Aunt\", \"Uncle\" ]\nval guest_list = \"Mom\" :: (\"Dad\" :: (\"Aunt\" :: (\"Uncle\" :: [])))\n\n(* If you have many lists of the same kind, you can concatenate them all *)\nval everyone = List.concat groups  (* [ \"Alice\", \"Bob\", \"Huey\", ... ] *)\n\n(* A list can contain any (finite) number of values *)\nval lots = [ 5, 5, 5, 6, 4, 5, 6, 5, 4, 5, 7, 3 ]  (* still just an int list *)\n\n(* Lists can only contain one kind of thing... *)\n(* val bad_list = [ 1, \"Hello\", 3.14159 ] : ??? list *)\n\n\n(* Tuples, on the other hand, can contain a fixed number of different things *)\nval person1 = (\"Simon\", 28, 3.14159)  (* : string * int * real *)\n\n(* You can even have tuples inside lists and lists inside tuples *)\nval likes = [ (\"Alice\", \"ice cream\"),\n              (\"Bob\",   \"hot dogs\"),\n              (\"Bob\",   \"Alice\") ]     (* : (string * string) list *)\n\nval mixup = [ (\"Alice\", 39),\n              (\"Bob\",   37),\n              (\"Eve\",   41) ]  (* : (string * int) list *)\n\nval good_bad_stuff =\n  ([\"ice cream\", \"hot dogs\", \"chocolate\"],\n   [\"liver\", \"paying the rent\" ])           (* : string list * string list *)\n\n\n(* Records are tuples with named slots *)\n\nval rgb = { r=0.23, g=0.56, b=0.91 } (* : {b:real, g:real, r:real} *)\n\n(* You don't need to declare their slots ahead of time. Records with\n   different slot names are considered different types, even if their\n   slot value types match up. For instance... *)\n\nval Hsl = { H=310.3, s=0.51, l=0.23 } (* : {H:real, l:real, s:real} *)\nval Hsv = { H=310.3, s=0.51, v=0.23 } (* : {H:real, s:real, v:real} *)\n\n(* ...trying to evaluate `Hsv = Hsl` or `rgb = Hsl` would give a type\n   error. While they're all three-slot records composed only of `real`s,\n   they each have different names for at least some slots. *)\n\n(* You can use hash notation to get values out of tuples. *)\n\nval H = #H Hsv (* : real *)\nval s = #s Hsl (* : real *)\n\n(* Functions! *)\nfun add_them (a, b) = a + b    (* A simple function that adds two numbers *)\nval test_it = add_them (3, 4)  (* gives 7 *)\n\n(* Larger functions are usually broken into several lines for readability *)\nfun thermometer temp =\n    if temp < 37\n    then \"Cold\"\n    else if temp > 37\n         then \"Warm\"\n         else \"Normal\"\n\nval test_thermo = thermometer 40  (* gives \"Warm\" *)\n\n(* if-sentences are actually expressions and not statements/declarations.\n   A function body can only contain one expression.  There are some tricks\n   for making a function do more than just one thing, though. *)\n\n(* A function can call itself as part of its result (recursion!) *)\nfun fibonacci n =\n    if n = 0 then 0 else                   (* Base case *)\n    if n = 1 then 1 else                   (* Base case *)\n    fibonacci (n - 1) + fibonacci (n - 2)  (* Recursive case *)\n\n(* Sometimes recursion is best understood by evaluating a function by hand:\n\n fibonacci 4\n   ~> fibonacci (4 - 1) + fibonacci (4 - 2)\n   ~> fibonacci 3 + fibonacci 2\n   ~> (fibonacci (3 - 1) + fibonacci (3 - 2)) + fibonacci 2\n   ~> (fibonacci 2 + fibonacci 1) + fibonacci 2\n   ~> ((fibonacci (2 - 1) + fibonacci (2 - 2)) + fibonacci 1) + fibonacci 2\n   ~> ((fibonacci 1 + fibonacci 0) + fibonacci 1) + fibonacci 2\n   ~> ((1 + fibonacci 0) + fibonacci 1) + fibonacci 2\n   ~> ((1 + 0) + fibonacci 1) + fibonacci 2\n   ~> (1 + fibonacci 1) + fibonacci 2\n   ~> (1 + 1) + fibonacci 2\n   ~> 2 + fibonacci 2\n   ~> 2 + (fibonacci (2 - 1) + fibonacci (2 - 2))\n   ~> 2 + (fibonacci (2 - 1) + fibonacci (2 - 2))\n   ~> 2 + (fibonacci 1 + fibonacci 0)\n   ~> 2 + (1 + fibonacci 0)\n   ~> 2 + (1 + 0)\n   ~> 2 + 1\n   ~> 3  which is the 4th Fibonacci number, according to this definition\n\n *)\n\n(* A function cannot change the variables it can refer to.  It can only\n   temporarily shadow them with new variables that have the same names.  In this\n   sense, variables are really constants and only behave like variables when\n   dealing with recursion.  For this reason, variables are also called value\n   bindings. An example of this: *)\n\nval x = 42\nfun answer(question) =\n    if question = \"What is the meaning of life, the universe and everything?\"\n    then x\n    else raise Fail \"I'm an exception. Also, I don't know what the answer is.\"\nval x = 43\nval hmm = answer \"What is the meaning of life, the universe and everything?\"\n(* Now, hmm has the value 42.  This is because the function answer refers to\n   the copy of x that was visible before its own function definition. *)\n\n\n(* Functions can take several arguments by taking one tuples as argument: *)\nfun solve2 (a : real, b : real, c : real) =\n    ((~b + Math.sqrt(b * b - 4.0 * a * c)) / (2.0 * a),\n     (~b - Math.sqrt(b * b - 4.0 * a * c)) / (2.0 * a))\n\n(* Sometimes, the same computation is carried out several times. It makes sense\n   to save and re-use the result the first time. We can use \"let-bindings\": *)\nfun solve2 (a : real, b : real, c : real) =\n    let val discr  = b * b - 4.0 * a * c\n        val sqr = Math.sqrt discr\n        val denom = 2.0 * a\n    in ((~b + sqr) / denom,\n        (~b - sqr) / denom)\n    end\n\n\n(* Pattern matching is a funky part of functional programming.  It is an\n   alternative to if-sentences.  The fibonacci function can be rewritten: *)\nfun fibonacci 0 = 0  (* Base case *)\n  | fibonacci 1 = 1  (* Base case *)\n  | fibonacci n = fibonacci (n - 1) + fibonacci (n - 2)  (* Recursive case *)\n\n(* Pattern matching is also possible on composite types like tuples, lists and\n   records. Writing \"fun solve2 (a, b, c) = ...\" is in fact a pattern match on\n   the one three-tuple solve2 takes as argument. Similarly, but less intuitively,\n   you can match on a list consisting of elements in it (from the beginning of\n   the list only). *)\nfun first_elem (x::xs) = x\nfun second_elem (x::y::xs) = y\nfun evenly_positioned_elems (odd::even::xs) = even::evenly_positioned_elems xs\n  | evenly_positioned_elems [odd] = []  (* Base case: throw away *)\n  | evenly_positioned_elems []    = []  (* Base case *)\n  \n(* The case expression can also be used to pattern match and return a value *)\ndatatype temp =\n      C of real\n    | F of real\n    \n(*  Declaring a new C temp value...\n    val t: temp = C 45.0  *)\n\nfun temp_to_f t =\n    case t of\n      C x => x * (9.0 / 5.0) + 32.0\n    | F x => x\n\n(* When matching on records, you must use their slot names, and you must bind\n   every slot in a record. The order of the slots doesn't matter though. *)\n\nfun rgbToTup {r, g, b} = (r, g, b)    (* fn : {b:'a, g:'b, r:'c} -> 'c * 'b * 'a *)\nfun mixRgbToTup {g, b, r} = (r, g, b) (* fn : {b:'a, g:'b, r:'c} -> 'c * 'b * 'a *)\n\n(* If called with {r=0.1, g=0.2, b=0.3}, either of the above functions\n   would return (0.1, 0.2, 0.3). But it would be a type error to call them\n   with {r=0.1, g=0.2, b=0.3, a=0.4} *)\n\n(* Higher order functions: Functions can take other functions as arguments.\n   Functions are just other kinds of values, and functions don't need names\n   to exist.  Functions without names are called \"anonymous functions\" or\n   lambda expressions or closures (since they also have a lexical scope). *)\nval is_large = (fn x => x > 37)\nval add_them = fn (a,b) => a + b\nval thermometer =\n    fn temp => if temp < 37\n               then \"Cold\"\n               else if temp > 37\n                    then \"Warm\"\n                    else \"Normal\"\n\n(* The following uses an anonymous function directly and gives \"ColdWarm\" *)\nval some_result = (fn x => thermometer (x - 5) ^ thermometer (x + 5)) 37\n\n(* Here is a higher-order function that works on lists (a list combinator) *)\n(* map f l\n       applies f to each element of l from left to right, \n       returning the list of results. *)\nval readings = [ 34, 39, 37, 38, 35, 36, 37, 37, 37 ]  (* first an int list *)\nval opinions = List.map thermometer readings (* gives [ \"Cold\", \"Warm\", ... ] *)\n\n(* And here is another one for filtering lists *)\nval warm_readings = List.filter is_large readings  (* gives [39, 38] *)\n\n(* You can create your own higher-order functions, too.  Functions can also take\n   several arguments by \"currying\" them. Syntax-wise this means adding spaces\n   between function arguments instead of commas and surrounding parentheses. *)\nfun map f [] = []\n  | map f (x::xs) = f(x) :: map f xs\n\n(* map has type ('a -> 'b) -> 'a list -> 'b list and is called polymorphic. *)\n(* 'a is called a type variable. *)\n\n\n(* We can declare functions as infix *)\nval plus = add_them   (* plus is now equal to the same function as add_them *)\ninfix plus            (* plus is now an infix operator *)\nval seven = 2 plus 5  (* seven is now bound to 7 *)\n\n(* Functions can also be made infix before they are declared *)\ninfix minus\nfun x minus y = x - y (* It becomes a little hard to see what's the argument *)\nval four = 8 minus 4  (* four is now bound to 4 *)\n\n(* An infix function/operator can be made prefix with 'op' *)\nval n = op + (5, 5)   (* n is now 10 *)\n\n(* 'op' is useful when combined with high order functions because they expect\n   functions and not operators as arguments. Most operators are really just\n   infix functions. *)\n(* foldl f init [x1, x2, ..., xn]\n       returns\n       f(xn, ...f(x2, f(x1, init))...)\n       or init if the list is empty. *)\nval sum_of_numbers = foldl op+ 0 [1, 2, 3, 4, 5]\n\n\n(* Datatypes are useful for creating both simple and complex structures *)\ndatatype color = Red | Green | Blue\n\n(* Here is a function that takes one of these as argument *)\nfun say(col) =\n    if col = Red then \"You are red!\" else\n    if col = Green then \"You are green!\" else\n    if col = Blue then \"You are blue!\" else\n    raise Fail \"Unknown color\"\n\nval _ = print (say(Red) ^ \"\\n\")\n\n(* Datatypes are very often used in combination with pattern matching *)\nfun say Red   = \"You are red!\"\n  | say Green = \"You are green!\"\n  | say Blue  = \"You are blue!\"\n\n(* We did not include the match arm `say _ = raise Fail \"Unknown color\"`\nbecause after specifying all three colors, the pattern is exhaustive\nand redundancy is not permitted in pattern matching *)\n\n\n(* Here is a binary tree datatype *)\ndatatype 'a btree = Leaf of 'a\n                  | Node of 'a btree * 'a * 'a btree (* three-arg constructor *)\n\n(* Here is a binary tree *)\nval myTree = Node (Leaf 9, 8, Node (Leaf 3, 5, Leaf 7))\n\n(* Drawing it, it might look something like...\n\n           8\n          / \\\n leaf -> 9   5\n            / \\\n   leaf -> 3   7 <- leaf\n *)\n\n(* This function counts the sum of all the elements in a tree *)\nfun count (Leaf n) = n\n  | count (Node (leftTree, n, rightTree)) = count leftTree + n + count rightTree\n\nval myTreeCount = count myTree  (* myTreeCount is now bound to 32 *)\n\n\n(* Exceptions! *)\n(* Exceptions can be raised/thrown using the reserved word 'raise' *)\nfun calculate_interest(n) = if n < 0.0\n                            then raise Domain\n                            else n * 1.04\n\n(* Exceptions can be caught using \"handle\" *)\nval balance = calculate_interest ~180.0\n              handle Domain => ~180.0    (* balance now has the value ~180.0 *)\n\n(* Some exceptions carry extra information with them *)\n(* Here are some examples of built-in exceptions *)\nfun failing_function []    = raise Empty  (* used for empty lists *)\n  | failing_function [x]   = raise Fail \"This list is too short!\"\n  | failing_function [x,y] = raise Overflow  (* used for arithmetic *)\n  | failing_function xs    = raise Fail \"This list is too long!\"\n\n(* We can pattern match in 'handle' to make sure\n   a specific exception was raised, or grab the message *)\nval err_msg = failing_function [1,2] handle Fail _ => \"Fail was raised\"\n                                          | Domain => \"Domain was raised\"\n                                          | Empty  => \"Empty was raised\"\n                                          | _      => \"Unknown exception\"\n\n(* err_msg now has the value \"Unknown exception\" because Overflow isn't\n   listed as one of the patterns -- thus, the catch-all pattern _ is used. *)\n\n(* We can define our own exceptions like this *)\nexception MyException\nexception MyExceptionWithMessage of string\nexception SyntaxError of string * (int * int)\n\n(* File I/O! *)\n(* Write a nice poem to a file *)\nfun writePoem(filename) =\n    let val file = TextIO.openOut(filename)\n        val _ = TextIO.output(file, \"Roses are red,\\nViolets are blue.\\n\")\n        val _ = TextIO.output(file, \"I have a gun.\\nGet in the van.\\n\")\n    in TextIO.closeOut(file)\n    end\n\n(* Read a nice poem from a file into a list of strings *)\nfun readPoem(filename) =\n    let val file = TextIO.openIn filename\n        val poem = TextIO.inputAll file\n        val _ = TextIO.closeIn file\n    in String.tokens (fn c => c = #\"\\n\") poem\n    end\n\nval _ = writePoem \"roses.txt\"\nval test_poem = readPoem \"roses.txt\"  (* gives [ \"Roses are red,\",\n                                                 \"Violets are blue.\",\n                                                 \"I have a gun.\",\n                                                 \"Get in the van.\" ] *)\n\n(* We can create references to data which can be updated *)\nval counter = ref 0 (* Produce a reference with the ref function *)\n\n(* Assign to a reference with the assignment operator *)\nfun set_five reference = reference := 5\n\n(* Read a reference with the dereference operator *)\nfun equals_five reference = !reference = 5\n\n(* We can use while loops for when recursion is messy *)\nfun decrement_to_zero r = if !r < 0\n                          then r := 0\n                          else while !r >= 0 do r := !r - 1\n\n(* This returns the unit value (in practical terms, nothing, a 0-tuple) *)\n\n(* To allow returning a value, we can use the semicolon to sequence evaluations *)\nfun decrement_ret x y = (x := !x - 1; y)\n```\n\n## Further learning\n\n* Install an interactive compiler (REPL), for example\n  [Poly/ML](http://www.polyml.org/),\n  [Moscow ML](http://mosml.org),\n  [SML/NJ](http://smlnj.org/).\n* Follow the Coursera course [Programming Languages](https://www.coursera.org/course/proglang).\n* Read *[ML for the Working Programmer](https://www.cl.cam.ac.uk/~lp15/MLbook/pub-details.html)* by Larry C. Paulson.\n* Use [StackOverflow's sml tag](http://stackoverflow.com/questions/tagged/sml).\n* Solve exercises on [Exercism.io's Standard ML track](https://exercism.io/tracks/sml)."
