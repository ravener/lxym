
package pages

const Paren = "[Paren](https://bitbucket.org/ktg/paren) is a dialect of Lisp. It is designed to be an embedded language.\n\nSome examples are from <http://learnxinyminutes.com/docs/racket/>.\n\n```scheme\n;;; Comments\n# comments\n\n;; Single line comments start with a semicolon or a sharp sign\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 1. Primitive Datatypes and Operators\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;;; Numbers\n123 ; int\n3.14 ; double\n6.02e+23 ; double\n(int 3.14) ; => 3 : int\n(double 123) ; => 123 : double\n\n;; Function application is written (f x y z ...)\n;; where f is a function and x, y, z, ... are operands\n;; If you want to create a literal list of data, use (quote) to stop it from\n;; being evaluated\n(quote (+ 1 2)) ; => (+ 1 2)\n;; Now, some arithmetic operations\n(+ 1 1)  ; => 2\n(- 8 1)  ; => 7\n(* 10 2) ; => 20\n(^ 2 3) ; => 8\n(/ 5 2) ; => 2\n(% 5 2) ; => 1\n(/ 5.0 2) ; => 2.5\n\n;;; Booleans\ntrue ; for true\nfalse ; for false\n(! true) ; => false\n(&& true false (prn \"doesn't get here\")) ; => false\n(|| false true (prn \"doesn't get here\")) ; => true\n\n;;; Characters are ints.\n(char-at \"A\" 0) ; => 65\n(chr 65) ; => \"A\"\n\n;;; Strings are fixed-length array of characters.\n\"Hello, world!\"\n\"Benjamin \\\"Bugsy\\\" Siegel\"   ; backslash is an escaping character\n\"Foo\\tbar\\r\\n\" ; includes C escapes: \\t \\r \\n\n\n;; Strings can be added too!\n(strcat \"Hello \" \"world!\") ; => \"Hello world!\"\n\n;; A string can be treated like a list of characters\n(char-at \"Apple\" 0) ; => 65\n\n;; Printing is pretty easy\n(pr \"I'm\" \"Paren. \") (prn \"Nice to meet you!\")\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 2. Variables\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; You can create or set a variable using (set)\n;; a variable name can use any character except: ();#\"\n(set some-var 5) ; => 5\nsome-var ; => 5\n\n;; Accessing a previously unassigned variable is an exception\n; x ; => Unknown variable: x : nil\n\n;; Local binding: Use lambda calculus! 'a' and 'b' are bound to '1' and '2' only within the (fn ...)\n((fn (a b) (+ a b)) 1 2) ; => 3\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 3. Collections\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;;; Lists\n\n;; Lists are vector-like data structures. (Random access is O(1).)\n(cons 1 (cons 2 (cons 3 (list)))) ; => (1 2 3)\n;; 'list' is a convenience variadic constructor for lists\n(list 1 2 3) ; => (1 2 3)\n;; and a quote can also be used for a literal list value\n(quote (+ 1 2)) ; => (+ 1 2)\n\n;; Can still use 'cons' to add an item to the beginning of a list\n(cons 0 (list 1 2 3)) ; => (0 1 2 3)\n\n;; Lists are a very basic type, so there is a *lot* of functionality for\n;; them, a few examples:\n(map inc (list 1 2 3))          ; => (2 3 4)\n(filter (fn (x) (== 0 (% x 2))) (list 1 2 3 4))    ; => (2 4)\n(length (list 1 2 3 4))     ; => 4\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 3. Functions\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Use 'fn' to create functions.\n;; A function always returns the value of its last expression\n(fn () \"Hello World\") ; => (fn () Hello World) : fn\n\n;; Use parentheses to call all functions, including a lambda expression\n((fn () \"Hello World\")) ; => \"Hello World\"\n\n;; Assign a function to a var\n(set hello-world (fn () \"Hello World\"))\n(hello-world) ; => \"Hello World\"\n\n;; You can shorten this using the function definition syntactic sugar:\n(defn hello-world2 () \"Hello World\")\n\n;; The () in the above is the list of arguments for the function\n(set hello\n  (fn (name)\n    (strcat \"Hello \" name)))\n(hello \"Steve\") ; => \"Hello Steve\"\n\n;; ... or equivalently, using a sugared definition:\n(defn hello2 (name)\n  (strcat \"Hello \" name))\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 4. Equality\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; for numbers use '=='\n(== 3 3.0) ; => true\n(== 2 1) ; => false\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 5. Control Flow\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;;; Conditionals\n\n(if true               ; test expression\n    \"this is true\"   ; then expression\n    \"this is false\") ; else expression\n; => \"this is true\"\n\n;;; Loops\n\n;; for loop is for number\n;; (for SYMBOL START END STEP EXPR ..)\n(for i 0 10 2 (pr i \"\")) ; => prints 0 2 4 6 8 10\n(for i 0.0 10 2.5 (pr i \"\")) ; => prints 0 2.5 5 7.5 10\n\n;; while loop\n((fn (i)\n  (while (< i 10)\n    (pr i)\n    (++ i))) 0) ; => prints 0123456789\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 6. Mutation\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Use 'set' to assign a new value to a variable or a place\n(set n 5) ; => 5\n(set n (inc n)) ; => 6\nn ; => 6\n(set a (list 1 2)) ; => (1 2)\n(set (nth 0 a) 3) ; => 3\na ; => (3 2)\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;; 7. Macros\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Macros let you extend the syntax of the language.\n;; Paren macros are easy.\n;; In fact, (defn) is a macro.\n(defmacro setfn (name ...) (set name (fn ...)))\n(defmacro defn (name ...) (def name (fn ...)))\n\n;; Let's add an infix notation\n(defmacro infix (a op ...) (op a ...))\n(infix 1 + 2 (infix 3 * 4)) ; => 15\n\n;; Macros are not hygienic, you can clobber existing variables!\n;; They are code transformations.\n```"