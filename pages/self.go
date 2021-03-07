
package pages

const Self = "Self is a fast prototype based OO language which runs in its own JIT vm. Most development is done through interacting with live objects through a visual development environment called *morphic* with integrated browsers and debugger.\n\nEverything in Self is an object. All computation is done by sending messages to objects. Objects in Self can be understood as sets of key-value slots.\n\n# Constructing objects\n\nThe inbuild Self parser can construct objects, including method objects.\n\n```\n\"This is a comment\"\n\n\"A string:\"\n'This is a string with \\'escaped\\' characters.\\n'\n\n\"A 30 bit integer\"\n23\n\n\"A 30 bit float\"\n3.2\n\n\"-20\"\n-14r16\n\n\"An object which only understands one message, 'x' which returns 20\"\n(|\n  x = 20.\n|)\n\n\"An object which also understands 'x:' which sets the x slot\"\n(|\n  x <- 20.\n|)\n\n\"An object which understands the method 'doubleX' which\ndoubles the value of x and then returns the object\"\n(|\n  x <- 20.\n  doubleX = (x: x * 2. self)\n|)\n\n\"An object which understands all the messages\nthat 'traits point' understands\". The parser\nlooks up 'traits point' by sending the messages\n'traits' then 'point' to a known object called\nthe 'lobby'. It looks up the 'true' object by\nalso sending the message 'true' to the lobby.\"\n(|     parent* = traits point.\n       x = 7.\n       y <- 5.\n       isNice = true.\n|)\n```\n\n# Sending messages to objects\n\nMessages can either be unary, binary or keyword. Precedence is in that order. Unlike Smalltalk, the precedence of binary messages must be specified, and all keywords after the first must start with a capital letter. Messages are separated from their destination by whitespace.\n\n```\n\"unary message, sends 'printLine' to the object '23'\nwhich prints the string '23' to stdout and returns the receiving object (ie 23)\"\n23 printLine\n\n\"sends the message '+' with '7' to '23', then the message '*' with '8' to the result\"\n(23 + 7) * 8\n\n\"sends 'power:' to '2' with '8' returns 256\"\n2 power: 8\n\n\"sends 'keyOf:IfAbsent:' to 'hello' with arguments 'e' and '-1'.\nReturns 1, the index of 'e' in 'hello'.\"\n'hello' keyOf: 'e' IfAbsent: -1\n```\n\n# Blocks\n\nSelf defines flow control like Smalltalk and Ruby by way of blocks. Blocks are delayed computations of the form:\n\n```\n[|:x. localVar| x doSomething with: localVar]\n```\n\nExamples of the use of a block:\n\n```\n\"returns 'HELLO'\"\n'hello' copyMutable mapBy: [|:c| c capitalize]\n\n\"returns 'Nah'\"\n'hello' size > 5 ifTrue: ['Yay'] False: ['Nah']\n\n\"returns 'HaLLO'\"\n'hello' copyMutable mapBy: [|:c|\n   c = 'e' ifTrue: [c capitalize]\n            False: ['a']]\n```\n\nMultiple expressions are separated by a period. ^ returns immediately.\n\n```\n\"returns An 'E'! How icky!\"\n'hello' copyMutable mapBy: [|:c. tmp <- ''|\n   tmp: c capitalize.\n   tmp = 'E' ifTrue: [^ 'An \\'E\\'! How icky!'].\n   c capitalize\n   ]\n```\n\nBlocks are performed by sending them the message 'value' and inherit (delegate to) their contexts:\n```\n\"returns 0\"\n[|x|\n    x: 15.\n    \"Repeatedly sends 'value' to the first block while the result of sending 'value' to the\n     second block is the 'true' object\"\n    [x > 0] whileTrue: [x: x - 1].\n    x\n] value\n```\n\n# Methods\n\nMethods are like blocks but they are not within a context but instead are stored as values of slots. Unlike Smalltalk, methods by default return their final value not 'self'.\n\n```\n\"Here is an object with one assignable slot 'x' and a method 'reduceXTo: y'.\nSending the message 'reduceXTo: 10' to this object will put\nthe object '10' in the 'x' slot and return the original object\"\n(|\n    x <- 50.\n    reduceXTo: y = (\n        [x > y] whileTrue: [x: x - 1].\n        self)\n|)\n.\n```\n\n# Prototypes\n\nSelf has no classes. The way to get an object is to find a prototype and copy it.\n\n```\n| d |\nd: dictionary copy.\nd at: 'hello' Put: 23 + 8.\nd at: 'goodbye' Put: 'No!.\n\"Prints No!\"\n( d at: 'goodbye' IfAbsent: 'Yes! ) printLine.\n\"Prints 31\"\n( d at: 'hello' IfAbsent: -1 ) printLine.\n```\n\n# Further information\n\nThe [Self handbook](http://handbook.selflanguage.org) has much more information, and nothing beats hand-on experience with Self by downloading it from the [homepage](http://www.selflanguage.org)."