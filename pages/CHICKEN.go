
package pages

const Chicken = "CHICKEN is an implementation of Scheme programming language that can\ncompile Scheme programs to C code as well as interpret them. CHICKEN\nsupports R5RS and R7RS (work in progress) standards and many extensions.\n\n\n```scheme\n;; #!/usr/bin/env csi -s\n\n;; Run the CHICKEN REPL in the commandline as follows :\n;; $ csi\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 0. Syntax\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Single line comments start with a semicolon\n\n#| Block comments\n   can span multiple lines and...\n   #| can be nested\n   |#\n|#\n\n;; S-expression comments are used to comment out expressions\n#; (display \"nothing\")    ; discard this expression \n\n;; CHICKEN has two fundamental pieces of syntax: Atoms and S-expressions\n;; an atom is something that evaluates to itself\n;; all builtin data types viz. numbers, chars, booleans, strings etc. are atoms\n;; Furthermore an atom can be a symbol, an identifier, a keyword, a procedure\n;; or the empty list (also called null)\n'athing              ;; => athing \n'+                   ;; => + \n+                    ;; => <procedure C_plus>\n\n;; S-expressions (short for symbolic expressions) consists of one or more atoms\n(quote +)            ;; => + ; another way of writing '+\n(+ 1 2 3)            ;; => 6 ; this S-expression evaluates to a function call\n'(+ 1 2 3)           ;; => (+ 1 2 3) ; evaluates to a list \n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 1. Primitive Datatypes and Operators \n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Numbers\n99999999999999999999 ;; integers\n#b1010               ;; binary ; => 10\n#o10                 ;; octal  ; => 8\n#x8ded               ;; hexadecimal ; => 36333\n3.14                 ;; real\n6.02e+23\n3/4                  ;; rational\n\n;;Characters and Strings\n#\\A                  ;; A char\n\"Hello, World!\"      ;; strings are fixed-length arrays of characters\n\n;; Booleans\n#t                  ;; true\n#f                  ;; false\n\n;; Function call is written as (f x y z ...)\n;; where f is a function and x,y,z, ... are arguments\n(print \"Hello, World!\")    ;; => Hello, World!\n;; formatted output\n(printf \"Hello, ~a.\\n\" \"World\")  ;; => Hello, World.\n\n;; print commandline arguments\n(map print (command-line-arguments)) \n\n(list 'foo 'bar 'baz)          ;; => (foo bar baz)\n(string-append \"pine\" \"apple\") ;; => \"pineapple\"\n(string-ref \"tapioca\" 3)       ;; => #\\i;; character 'i' is at index 3\n(string->list \"CHICKEN\")       ;; => (#\\C #\\H #\\I #\\C #\\K #\\E #\\N)\n(string-intersperse '(\"1\" \"2\") \":\") ;; => \"1:2\"\n(string-split \"1:2:3\" \":\")     ;; => (\"1\" \"2\" \"3\")\n\n\n;; Predicates are special functions that return boolean values\n(atom? #t)                ;; => #t\n\n(symbol? #t)              ;; => #f\n\n(symbol? '+)              ;; => #t\n\n(procedure? +)            ;; => #t\n\n(pair? '(1 2))            ;; => #t\n\n(pair? '(1 2 . 3))        ;; => #t\n\n(pair? '())               ;; => #f\n\n(list? '())               ;; => #t\n\n\n;; Some arithmetic operations\n\n(+ 1 1)                   ;; => 2\n(- 8 1)                   ;; => 7\n(* 10 2)                  ;; => 20\n(expt 2 3)                ;; => 8\n(remainder 5 2)           ;; => 1\n(/ 35 5)                  ;; => 7\n(/ 1 3)                   ;; => 0.333333333333333\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 2. Variables\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; You can create variables with define\n;; A variable name can use any character except: ()[]{}\",'`;#\\\n(define myvar 5)\nmyvar        ;; => 5\n\n;; Alias to a procedure\n(define ** expt)\n(** 2 3)     ;; => 8\n\n;; Accessing an undefined variable raises an exception\ns            ;; => Error: unbound variable: s\n\n;; Local binding\n(let ((me \"Bob\"))\n  (print me))     ;; => Bob\n\n(print me)        ;; => Error: unbound variable: me\n\n;; Assign a new value to previously defined variable\n(set! myvar 10) \nmyvar             ;; => 10\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 3. Collections\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Pairs\n;; 'cons' constructs pairs, \n;; 'car' extracts the first element, 'cdr' extracts the rest of the elements\n(cons 'subject 'verb)       ;; => '(subject . verb)\n(car (cons 'subject 'verb)) ;; => subject\n(cdr (cons 'subject 'verb)) ;; => verb\n\n;; Lists\n;; cons creates a new list if the second item is a list\n(cons 0 '())         ;; => (0)\n(cons 1 (cons 2  (cons 3 '())))    ;; => (1 2 3)\n;; 'list' is a convenience variadic constructor for lists\n(list 1 2 3)    ;; => (1 2 3)\n\n\n;; Use 'append' to append lists together\n(append '(1 2) '(3 4)) ;; => (1 2 3 4)\n\n;; Some basic operations on lists\n(map add1 '(1 2 3))    ;; => (2 3 4)\n(reverse '(1 3 4 7))   ;; => (7 4 3 1)\n(sort '(11 22 33 44) >)   ;; => (44 33 22 11)\n\n(define days '(SUN MON FRI))\n(list-ref days 1)      ;; => MON\n(set! (list-ref days 1) 'TUE)\ndays                   ;; => (SUN TUE FRI)\n\n;; Vectors\n;; Vectors are heterogeneous structures whose elements are indexed by integers\n;; A Vector typically occupies less space than a list of the same length\n;; Random access of an element in a vector is faster than in a list\n#(1 2 3)                     ;; => #(1 2 3) ;; literal syntax\n(vector 'a 'b 'c)            ;; => #(a b c) \n(vector? #(1 2 3))           ;; => #t\n(vector-length #(1 (2) \"a\")) ;; => 3\n(vector-ref #(1 (2) (3 3)) 2);; => (3 3)\n\n(define vec #(1 2 3))\n(vector-set! vec 2 4)\nvec                         ;; => #(1 2 4)\n\n;; Vectors can be created from lists and vice-verca\n(vector->list #(1 2 4))     ;; => '(1 2 4)\n(list->vector '(a b c))     ;; => #(a b c)\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 4. Functions\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Use 'lambda' to create functions.\n;; A function always returns the value of its last expression\n(lambda () \"Hello World\")   ;; => #<procedure (?)> \n\n;; Use extra parens around function definition to execute \n((lambda () \"Hello World\")) ;; => Hello World ;; argument list is empty\n\n;; A function with an argument\n((lambda (x) (* x x)) 3)           ;; => 9\n;; A function with two arguments\n((lambda (x y) (* x y)) 2 3)       ;; => 6\n\n;; assign a function to a variable\n(define sqr (lambda (x) (* x x)))\nsqr                        ;; => #<procedure (sqr x)>\n(sqr 3)                    ;; => 9\n\n;; We can shorten this using the function definition syntactic sugar\n(define (sqr x) (* x x))\n(sqr 3)                    ;; => 9\n\n;; We can redefine existing procedures\n(foldl cons '() '(1 2 3 4 5)) ;; => (((((() . 1) . 2) . 3) . 4) . 5)\n(define (foldl func accu alist)\n  (if (null? alist)\n    accu\n    (foldl func (func (car alist) accu) (cdr alist))))\n\n(foldl cons '() '(1 2 3 4 5))   ;; => (5 4 3 2 1)\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 5. Equality\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; For numbers use '='\n(= 3 3.0)                  ;; => #t\n(= 2 1)                    ;; => #f\n\n;; 'eq?' returns #t if two arguments refer to the same object in memory\n;; In other words, it's a simple pointer comparison.\n(eq? '() '())              ;; => #t ;; there's only one empty list in memory\n(eq? (list 3) (list 3))    ;; => #f ;; not the same object\n(eq? 'yes 'yes)            ;; => #t\n(eq? 3 3)                  ;; => #t ;; don't do this even if it works in this case\n(eq? 3 3.0)                ;; => #f ;; it's better to use '=' for number comparisons\n(eq? \"Hello\" \"Hello\")      ;; => #f\n\n;; 'eqv?' is same as 'eq?' all datatypes except numbers and characters\n(eqv? 3 3.0)               ;; => #f\n(eqv? (expt 2 3) (expt 2 3)) ;; => #t\n(eqv? 'yes 'yes)           ;; => #t\n\n;; 'equal?' recursively compares the contents of pairs, vectors, and strings,\n;; applying eqv? on other objects such as numbers and symbols. \n;; A rule of thumb is that objects are generally equal? if they print the same.\n\n(equal? '(1 2 3) '(1 2 3)) ;; => #t\n(equal? #(a b c) #(a b c)) ;; => #t\n(equal? 'a 'a)             ;; => #t\n(equal? \"abc\" \"abc\")       ;; => #t\n\n;; In Summary:\n;; eq? tests if objects are identical\n;; eqv? tests if objects are operationally equivalent\n;; equal? tests if objects have same structure and contents\n\n;; Comparing strings for equality\n(string=? \"Hello\" \"Hello\") ;; => #t\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 6. Control Flow\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Conditionals\n(if #t                     ;; test expression\n  \"True\"                   ;; then expression\n  \"False\")                 ;; else expression\n                           ;; => \"True\"\n\n(if (> 3 2)\n  \"yes\"\n  \"no\")                    ;; => \"yes\"\n\n;; In conditionals, all values that are not '#f' are treated as true.\n;; 0, '(), #() \"\" , are all true values\n(if 0\n  \"0 is not false\"\n  \"0 is false\")            ;; => \"0 is not false\"\n\n;; 'cond' chains a series of tests and returns as soon as it encounters a true condition\n;; 'cond' can be used to simulate 'if/elseif/else' statements\n(cond ((> 2 2) \"not true so don't return this\")\n      ((< 2 5) \"true, so return this\")\n      (else \"returning default\"))    ;; => \"true, so return this\"\n\n\n;; A case expression is evaluated as follows:\n;; The key is evaluated and compared with each datum in sense of 'eqv?',\n;; The corresponding clause in the matching datum is evaluated and returned as result\n(case (* 2 3)              ;; the key is 6\n  ((2 3 5 7) 'prime)       ;; datum 1\n  ((1 4 6 8) 'composite))  ;; datum 2; matched!\n                           ;; => composite\n\n;; case with else clause\n(case (car '(c d))\n  ((a e i o u) 'vowel)\n  ((w y) 'semivowel)\n  (else 'consonant))       ;; =>  consonant\n\n;; Boolean expressions\n;; 'and' returns the first expression that evaluates to #f\n;; otherwise, it returns the result of the last expression\n(and #t #f (= 2 2.0))                ;; => #f\n(and (< 2 5) (> 2 0) \"0 < 2 < 5\")    ;; => \"0 < 2 < 5\"\n\n;; 'or' returns the first expression that evaluates to #t \n;; otherwise the result of the last expression is returned\n(or #f #t #f)                        ;; => #t\n(or #f #f #f)                        ;; => #f\n\n;; 'when' is like 'if' without the else expression\n(when (positive? 5) \"I'm positive\")  ;; => \"I'm positive\"\n\n;; 'unless' is equivalent to (when (not <test>) <expr>)\n(unless (null? '(1 2 3)) \"not null\") ;; => \"not null\"\n\n\n;; Loops\n;; loops can be created with the help of tail-recursions\n(define (loop count)\n  (unless (= count 0)\n    (print \"hello\") \n    (loop (sub1 count))))\n(loop 4)                             ;; => hello, hello ...\n\n;; Or with a named let\n(let loop ((i 0) (limit 5))\n  (when (< i limit)\n    (printf \"i = ~a\\n\" i)\n    (loop (add1 i) limit)))          ;; => i = 0, i = 1....\n\n;; 'do' is another iteration construct\n;; It initializes a set of variables and updates them in each iteration\n;; A final expression is evaluated after the exit condition is met\n(do ((x 0 (add1 x )))            ;; initialize x = 0 and add 1 in each iteration\n  ((= x 10) (print \"done\"))      ;; exit condition and final expression\n  (print x))                     ;; command to execute in each step\n                                 ;; => 0,1,2,3....9,done\n\n;; Iteration over lists \n(for-each (lambda (a) (print (* a a)))\n          '(3 5 7))                  ;; => 9, 25, 49\n\n;; 'map' is like for-each but returns a list\n(map add1 '(11 22 33))               ;; => (12 23 34)\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 7. Extensions\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; The CHICKEN core is very minimal, but additional features are provided by library extensions known as Eggs.\n;; You can install Eggs with 'chicken-install <eggname>' command.\n\n;; 'numbers' egg provides support for full numeric tower.\n(require-extension numbers)\n;; complex numbers\n3+4i                               ;; => 3+2i\n;; Supports fractions without falling back to inexact flonums\n1/3                                ;; => 1/3\n;; provides support for large integers through bignums\n(expt 9 20)                        ;; => 12157665459056928801 \n;; And other 'extended' functions\n(log 10 (exp 1))                   ;; => 2.30258509299405\n(numerator 2/3)                    ;; => 2\n\n;; 'utf8' provides unicode support\n(require-extension utf8)\n\"\\u03BBx:(\\u03BC\\u0251.\\u0251\\u2192\\u0251).xx\" ;; => \"λx:(μɑ.ɑ→ɑ).xx\"\n\n;; 'posix' provides file I/O and lots of other services for unix-like operating systems\n;; Some of the functions are not available in Windows system,\n;; See http://wiki.call-cc.org/man/4/Unit%20posix for more details\n\n;; Open a file to append, open \"write only\" and create file if it does not exist\n(define outfn (file-open \"chicken-hen.txt\" (+ open/append open/wronly open/creat)))\n;; write some text to the file\n(file-write outfn \"Did chicken came before hen?\") \n;; close the file\n(file-close outfn)\n;; Open the file \"read only\"\n(define infn (file-open \"chicken-hen.txt\" open/rdonly))\n;; read some text from the file\n(file-read infn 30)         ;; => (\"Did chicken came before hen?  \", 28)\n(file-close infn)\n\n;; CHICKEN also supports SRFI (Scheme Requests For Implementation) extensions\n;; See 'http://srfi.schemers.org/srfi-implementers.html\" to see srfi's supported by CHICKEN\n(require-extension srfi-1)         ;; list library\n(filter odd? '(1 2 3 4 5 6 7))     ;; => (1 3 5 7)\n(count even? '(1 2 3 4 5))         ;; => 2\n(take '(12 24 36 48 60) 3)         ;; => (12 24 36)\n(drop '(12 24 36 48 60) 2)         ;; => (36 48 60)\n(circular-list 'z 'q)              ;; => z q z q ...\n\n(require-extension srfi-13)        ;; string library\n(string-reverse \"pan\")             ;; => \"nap\"\n(string-index \"Turkey\" #\\k)        ;; => 3\n(string-every char-upper-case? \"CHICKEN\") ;; => #t\n(string-join '(\"foo\" \"bar\" \"baz\") \":\")    ;; => \"foo:bar:baz\"\n\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 8. Macros\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; A 'for .. in ..' iteration like python, for lists\n(define-syntax for\n  (syntax-rules (in)\n                ((for elem in alist body ...)\n                 (for-each (lambda (elem) body ...) alist))))\n\n(for x in '(2 4 8 16)\n     (print x))          ;; => 2, 4, 8, 16\n\n(for chr in (string->list \"PENCHANT\")\n     (print chr))        ;; => P, E, N, C, H, A, N, T\n\n;; While loop\n(define-syntax while\n  (syntax-rules ()\n                ((while cond body ...)\n                 (let loop ()\n                   (when cond\n                     body ...\n                     (loop))))))\n\n(let ((str \"PENCHANT\") (i 0))\n  (while (< i (string-length str))     ;; while (condition)\n         (print (string-ref str i))    ;; body \n         (set! i (add1 i))))           \n                                       ;; => P, E, N, C, H, A, N, T\n\n;; Advanced Syntax-Rules Primer -> http://petrofsky.org/src/primer.txt\n;; Macro system in chicken -> http://lists.gnu.org/archive/html/chicken-users/2008-04/msg00013.html\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n; 9. Modules\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n;; Also See http://wiki.call-cc.org/man/4/Modules\n\n;; The 'test' module exports a value named 'hello' and a macro named 'greet'\n(module test (hello greet)\n  (import scheme)\n\n  (define-syntax greet\n    (syntax-rules ()\n      ((_ whom) \n       (begin\n         (display \"Hello, \")\n         (display whom)\n         (display \" !\\n\") ) ) ) )\n\n  (define (hello)\n    (greet \"world\") )  )\n\n;; we can define our modules in a separate file (say test.scm) and load them to the interpreter with\n;;         (load \"test.scm\")\n\n;; import the module\n(import test)\n(hello)                ;; => Hello, world !\n(greet \"schemers\")     ;; => Hello, schemers !\n\n;; We can compile the module files in to shared libraries by using following command,\n;;         csc -s test.scm\n;;         (load \"test.so\")\n\n;; Functors\n;; Functors are high level modules that can be parameterized by other modules\n;; Following functor requires another module named 'M' that provides a function called 'multiply'\n;; The functor itself exports a generic function 'square'\n(functor (squaring-functor (M (multiply))) (square)\n         (import scheme M) \n         (define (square x) (multiply x x)))\n\n;; Module 'nums' can be passed as a parameter to 'squaring-functor'\n(module nums (multiply) \n        (import scheme)     ;; predefined modules\n        (define (multiply x y) (* x y))) \n;; the final module can be imported and used in our program\n(module number-squarer = (squaring-functor nums)) \n\n(import number-squarer)\n(square 3)              ;; => 9\n\n;; We can instantiate the functor for other inputs\n;; Here's another example module that can be passed to squaring-functor\n(module stars (multiply)\n        (import chicken scheme)  ;; chicken module for the 'use' keyword\n        (use srfi-1)             ;; we can use external libraries in our module\n        (define (multiply x y)\n          (list-tabulate x (lambda _ (list-tabulate y (lambda _ '*))))))\n(module star-squarer = (squaring-functor stars))\n\n(import star-squarer)\n(square 3)              ;; => ((* * *)(* * *)(* * *))\n\n```\n## Further Reading\n* [CHICKEN User's Manual](http://wiki.call-cc.org/man/4/The%20User%27s%20Manual).\n* [R5RS standards](http://www.schemers.org/Documents/Standards/R5RS)\n\n\n## Extra Info\n\n* [For programmers of other languages](http://wiki.call-cc.org/chicken-for-programmers-of-other-languages)\n* [Compare CHICKEN syntax with other languages](http://plr.sourceforge.net/cgi-bin/plr/launch.py)"