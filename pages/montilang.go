
package pages

const Montilang = "MontiLang is a Stack-Oriented concatenative imperative programming language. Its syntax\nis roughly based off of forth with similar style for doing arithmetic in [reverse polish notation.](https://en.wikipedia.org/wiki/Reverse_Polish_notation)\n\nA good way to start with MontiLang is to read the documentation and examples at [montilang.ml](http://montilang.ml),\nthen download MontiLang or build from source code with the instructions provided.\n\n```\n/# Monti Reference sheet #/\n/#\nComments are multiline\nNested comments are not supported \n#/\n/# Whitespace is all arbitrary, indentation is optional #/\n/# All programming in Monti is done by manipulating the parameter stack \narithmetic and stack operations in MontiLang are similar to FORTH\nhttps://en.wikipedia.org/wiki/Forth_(programming_language)\n#/\n\n/# in Monti, everything is either a string or a number. Operations treat all numbers\nsimilarly to floats, but anything without a remainder is treated as type int #/\n\n/# numbers and strings are added to the stack from left to right #/\n\n/# Arithmetic works by manipulating data on the stack #/\n\n5 3 + PRINT . /# 8 #/\n\n/#  5 and 3 are pushed onto the stack\n    '+' replaces top 2 items on stack with sum of top 2 items\n    'PRINT' prints out the top item on the stack\n    '.' pops the top item from the stack. \n    #/\n\n6 7 * PRINT . /# 42 #/\n1360 23 - PRINT . /# 1337 #/\n12 12 / PRINT . /# 1 #/\n13 2 % PRINT . /# 1 #/\n\n37 NEG PRINT . /# -37 #/\n-12 ABS PRINT . /# 12 #/\n52 23 MAX PRINT . /# 52 #/\n52 23 MIN PRINT . /# 23 #/\n\n/# 'PSTACK' command prints the entire stack, 'CLEAR' clears the entire stack #/\n\n3 6 8 PSTACK CLEAR /# [3, 6, 8] #/\n\n/# Monti comes with some tools for stack manipulation #/\n\n2 DUP PSTACK CLEAR /# [2, 2] - Duplicate the top item on the stack#/\n2 6 SWAP PSTACK CLEAR /# [6, 2] - Swap top 2 items on stack #/\n1 2 3 ROT PSTACK CLEAR /# [2, 3, 1] - Rotate top 3 items on stack #/\n2 3 NIP PSTACK CLEAR /# [3] - delete second item from the top of the stack #/\n4 5 6 TRIM PSTACK CLEAR /# [5, 6] - Deletes first item on stack #/\n/# variables are assigned with the syntax 'VAR [name]'#/\n/# When assigned, the variable will take the value of the top item of the stack #/\n\n6 VAR six . /# assigns var 'six' to be equal to 6 #/\n3 6 + VAR a . /# assigns var 'a' to be equal to 9 #/\n\n/# the length of the stack can be calculated with the statement 'STKLEN' #/\n1 2 3 4 STKLEN PRINT CLEAR /# 4 #/\n\n/# strings are defined with | | #/\n\n|Hello World!| VAR world . /# sets variable 'world' equal to string 'Hello world! #/ \n\n/# variables can be called by typing its name. when called, the value of the variable is pushed\nto the top of the stack #/\nworld PRINT .\n\n/# with the OUT statement, the top item on the stack can be printed without a newline #/\n\n|world!| |Hello, | OUT SWAP PRINT CLEAR\n\n/# Data types can be converted between strings and integers with the commands 'TOINT' and 'TOSTR'#/\n|5| TOINT PSTACK . /# [5] #/\n45 TOSTR PSTACK . /# ['45'] #/\n\n/# User input is taken with INPUT and pushed to the stack. If the top item of the stack is a string, \nthe string is used as an input prompt #/\n\n|What is your name? | INPUT NIP \n|Hello, | OUT SWAP PRINT CLEAR\n\n\n/# FOR loops have the syntax 'FOR [condition] [commands] ENDFOR' At the moment, [condition] can\nonly have the value of an integer. Either by using an integer, or a variable call to an integer.\n[commands] will be interpereted the amount of time specified in [condition] #/\n/# E.G: this prints out 1 to 10 #/\n\n1 VAR a .\nFOR 10\n    a PRINT 1 + VAR a\nENDFOR\n\n/# the syntax for while loops are similar. A number is evaluated as true if it is larger than\n0. a string is true if its length > 0. Infinite loops can be used by using literals.\n#/\n10 var loop .\nWHILE loop\n    loop print \n    1 - var loop\nENDWHILE\n/#\nthis loop would count down from 10.\n\nIF statements are pretty much the same, but only are executed once.\n#/\nIF loop\n loop PRINT .\nENDIF\n\n/# This would only print 'loop' if it is larger than 0 #/\n\n/# If you would want to use the top item on the stack as loop parameters, this can be done with the ':' character #/\n\n/# eg, if you wanted to print 'hello' 7 times, instead of using #/\n\nFOR 7\n    |hello| PRINT .\nENDFOR\n\n/# this could be used #/\n7\nFOR : \n    |hello| PRINT .\nENDFOR\n\n/# Equality and inequality statements use the top 2 items on the stack as parameters, and replace the top two items with the output #/\n/# If it is true, the top 2 items are replaced with '1'. If false, with '0'. #/\n\n7 3 > PRINT . /# 1 #/\n2 10 > PRINT . /# 0 #/\n5 9 <= PRINT . /# 1 #/\n5 5 == PRINT . /# 1 #/\n5 7 == PRINT . /# 0 #/\n3 8 != PRINT . /# 1 #/\n\n/# User defined commands have the syntax of 'DEF [name] [commands] ENDDEF'. #/\n/# eg, if you wanted to define a function with the name of 'printseven' to print '7' 10 times, this could be used #/\n\nDEF printseven\n    FOR 10\n       7 PRINT .\n    ENDFOR\nENDDEF\n\n/# to run the defined statement, simply type it and it will be run by the interpereter #/\n\nprintseven\n\n/# Montilang supports AND, OR and NOT statements #/\n\n1 0 AND PRINT . /# 0 #/\n1 1 AND PRINT . /# 1 #/\n1 0 OR PRINT . /# 1 #/\n0 0 OR PRINT . /# 0 #/\n1 NOT PRINT . /# 0 #/\n0 NOT PRINT . /# 1 #/\n\n/# Preprocessor statements are made inbetween '&' characters #/\n/# currently, preprocessor statements can be used to make c++-style constants #/\n\n&DEFINE LOOPSTR 20&\n/# must have & on either side with no spaces, 'DEFINE' is case sensative. #/\n/# All statements are scanned and replaced before the program is run, regardless of where the statements are placed #/\n\nFOR LOOPSTR 7 PRINT . ENDFOR /# Prints '7' 20 times. At run, 'LOOPSTR' in source code is replaced with '20' #/ \n\n/# Multiple files can be used with the &INCLUDE <filename>& Command that operates similar to c++, where the file specified is tokenized,\n   and the &INCLUDE statement is replaced with the file #/\n   \n/# E.G, you can have a program be run through several files. If you had the file 'name.mt' with the following data:\n\n[name.mt]\n|Hello, | OUT . name PRINT .\n\na program that asks for your name and then prints it out can be defined as such: #/\n\n|What is your name? | INPUT VAR name . &INCLUDE name.mt&\n\n/# ARRAYS: #/\n\n/# arrays are defined with the statement 'ARR'\nWhen called, everything currently in the stack is put into one\narray and all items on the stack are replaced with the new array. #/\n\n2 3 4 ARR PSTACK . /# [[2, 3, 4]] #/\n\n/# the statement 'LEN' adds the length of the last item on the stack to the stack.\nThis can be used on arrays, as well as strings. #/\n\n3 4 5 ARR LEN PRINT . /# 3 #/\n\n/# values can be appended to an array with the statement 'APPEND' #/\n\n1 2 3 ARR 5 APPEND . PRINT . /# [1, 2, 3, 5] #/\n\n/# an array at the top of the stack can be wiped with the statement 'WIPE' #/\n3 4 5 ARR WIPE PRINT . /# [] #/\n\n/# The last item of an array can be removed with the statement 'DROP' #/\n\n3 4 5 ARR DROP PRINT . /# [3, 4]\n/# arrays, like other datatypes can be stored in variables #/\n5 6 7 ARR VAR list .\nlist PRINT . /# [5, 6, 7] #/\n\n/# Values at specific indexes can be changed with the statement 'INSERT <index>' #/\n4 5 6 ARR \n97 INSERT 1 . PRINT /# 4, 97, 6 #/\n\n/# Values at specific indexes can be deleted with the statement 'DEL <index>' #/\n1 2 3 ARR\nDEL 1 PRINT . /# [1, 3] #/\n\n/# items at certain indexes of an array can be gotten with the statement 'GET <index>' #/\n\n1 2 3 ARR GET 2 PSTACK /# [[1, 2, 3], 3] #/\n```\n\n## Extra information\n\n- [MontiLang.ml](http://montilang.ml/)\n- [Github Page](https://github.com/lduck11007/MontiLang)"