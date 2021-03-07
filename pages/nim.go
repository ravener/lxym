
package pages

const Nim = "Nim (formerly Nimrod) is a statically typed, imperative programming language\nthat gives the programmer power without compromises on runtime efficiency.\n\nNim is efficient, expressive, and elegant.\n\n```nim\n# Single-line comments start with a #\n\n#[\n  This is a multiline comment.\n  In Nim, multiline comments can be nested, beginning with #[\n  ... and ending with ]#\n]#\n\ndiscard \"\"\"\nThis can also work as a multiline comment.\nOr for unparsable, broken code\n\"\"\"\n\nvar                     # Declare (and assign) variables,\n  letter: char = 'n'    # with or without type annotations\n  lang = \"N\" & \"im\"\n  nLength: int = len(lang)\n  boat: float\n  truth: bool = false\n\nlet            # Use let to declare and bind variables *once*.\n  legs = 400   # legs is immutable.\n  arms = 2_000 # _ are ignored and are useful for long numbers.\n  aboutPi = 3.15\n\nconst            # Constants are computed at compile time. This provides\n  debug = true   # performance and is useful in compile time expressions.\n  compileBadCode = false\n\nwhen compileBadCode:            # `when` is a compile time `if`\n  legs = legs + 1               # This error will never be compiled.\n  const input = readline(stdin) # Const values must be known at compile time.\n\ndiscard 1 > 2 # Note: The compiler will complain if the result of an expression\n              # is unused. `discard` bypasses this.\n\n\n#\n# Data Structures\n#\n\n# Tuples\n\nvar\n  child: tuple[name: string, age: int]   # Tuples have *both* field names\n  today: tuple[sun: string, temp: float] # *and* order.\n\nchild = (name: \"Rudiger\", age: 2) # Assign all at once with literal ()\ntoday.sun = \"Overcast\"            # or individual fields.\ntoday.temp = 70.1\n\n# Sequences\n\nvar\n  drinks: seq[string]\n\ndrinks = @[\"Water\", \"Juice\", \"Chocolate\"] # @[V1,..,Vn] is the sequence literal\n\ndrinks.add(\"Milk\")\n\nif \"Milk\" in drinks:\n  echo \"We have Milk and \", drinks.len - 1, \" other drinks\"\n\nlet myDrink = drinks[2]\n\n#\n# Defining Types\n#\n\n# Defining your own types puts the compiler to work for you. It's what makes\n# static typing powerful and useful.\n\ntype\n  Name = string # A type alias gives you a new type that is interchangeable\n  Age = int     # with the old type but is more descriptive.\n  Person = tuple[name: Name, age: Age] # Define data structures too.\n  AnotherSyntax = tuple\n    fieldOne: string\n    secondField: int\n\nvar\n  john: Person = (name: \"John B.\", age: 17)\n  newage: int = 18 # It would be better to use Age than int\n\njohn.age = newage # But still works because int and Age are synonyms\n\ntype\n  Cash = distinct int    # `distinct` makes a new type incompatible with its\n  Desc = distinct string # base type.\n\nvar\n  money: Cash = 100.Cash # `.Cash` converts the int to our type\n  description: Desc  = \"Interesting\".Desc\n\nwhen compileBadCode:\n  john.age  = money        # Error! age is of type int and money is Cash\n  john.name = description  # Compiler says: \"No way!\"\n\n#\n# More Types and Data Structures\n#\n\n# Enumerations allow a type to have one of a limited number of values\n\ntype\n  Color = enum cRed, cBlue, cGreen\n  Direction = enum # Alternative formatting\n    dNorth\n    dWest\n    dEast\n    dSouth\nvar\n  orient = dNorth # `orient` is of type Direction, with the value `dNorth`\n  pixel = cGreen # `pixel` is of type Color, with the value `cGreen`\n\ndiscard dNorth > dEast # Enums are usually an \"ordinal\" type\n\n# Subranges specify a limited valid range\n\ntype\n  DieFaces = range[1..20] # Only an int from 1 to 20 is a valid value\nvar\n  my_roll: DieFaces = 13\n\nwhen compileBadCode:\n  my_roll = 23 # Error!\n\n# Arrays\n\ntype\n  RollCounter = array[DieFaces, int]  # Array's are fixed length and\n  DirNames = array[Direction, string] # indexed by any ordinal type.\n  Truths = array[42..44, bool]\nvar\n  counter: RollCounter\n  directions: DirNames\n  possible: Truths\n\npossible = [false, false, false] # Literal arrays are created with [V1,..,Vn]\npossible[42] = true\n\ndirections[dNorth] = \"Ahh. The Great White North!\"\ndirections[dWest] = \"No, don't go there.\"\n\nmy_roll = 13\ncounter[my_roll] += 1\ncounter[my_roll] += 1\n\nvar anotherArray = [\"Default index\", \"starts at\", \"0\"]\n\n# More data structures are available, including tables, sets, lists, queues,\n# and crit bit trees.\n# http://nim-lang.org/docs/lib.html#collections-and-algorithms\n\n#\n# IO and Control Flow\n#\n\n# `case`, `readLine()`\n\necho \"Read any good books lately?\"\ncase readLine(stdin)\nof \"no\", \"No\":\n  echo \"Go to your local library.\"\nof \"yes\", \"Yes\":\n  echo \"Carry on, then.\"\nelse:\n  echo \"That's great; I assume.\"\n\n# `while`, `if`, `continue`, `break`\n\nimport strutils as str # http://nim-lang.org/docs/strutils.html\necho \"I'm thinking of a number between 41 and 43. Guess which!\"\nlet number: int = 42\nvar\n  raw_guess: string\n  guess: int\nwhile guess != number:\n  raw_guess = readLine(stdin)\n  if raw_guess == \"\": continue # Skip this iteration\n  guess = str.parseInt(raw_guess)\n  if guess == 1001:\n    echo(\"AAAAAAGGG!\")\n    break\n  elif guess > number:\n    echo(\"Nope. Too high.\")\n  elif guess < number:\n    echo(guess, \" is too low\")\n  else:\n    echo(\"Yeeeeeehaw!\")\n\n#\n# Iteration\n#\n\nfor i, elem in [\"Yes\", \"No\", \"Maybe so\"]: # Or just `for elem in`\n  echo(elem, \" is at index: \", i)\n\nfor k, v in items(@[(person: \"You\", power: 100), (person: \"Me\", power: 9000)]):\n  echo v\n\nlet myString = \"\"\"\nan <example>\n`string` to\nplay with\n\"\"\" # Multiline raw string\n\nfor line in splitLines(myString):\n  echo(line)\n\nfor i, c in myString:       # Index and letter. Or `for j in` for just letter\n  if i mod 2 == 0: continue # Compact `if` form\n  elif c == 'X': break\n  else: echo(c)\n\n#\n# Procedures\n#\n\ntype Answer = enum aYes, aNo\n\nproc ask(question: string): Answer =\n  echo(question, \" (y/n)\")\n  while true:\n    case readLine(stdin)\n    of \"y\", \"Y\", \"yes\", \"Yes\":\n      return Answer.aYes  # Enums can be qualified\n    of \"n\", \"N\", \"no\", \"No\":\n      return Answer.aNo\n    else: echo(\"Please be clear: yes or no\")\n\nproc addSugar(amount: int = 2) = # Default amount is 2, returns nothing\n  assert(amount > 0 and amount < 9000, \"Crazy Sugar\")\n  for a in 1..amount:\n    echo(a, \" sugar...\")\n\ncase ask(\"Would you like sugar in your tea?\")\nof aYes:\n  addSugar(3)\nof aNo:\n  echo \"Oh do take a little!\"\n  addSugar()\n# No need for an `else` here. Only `yes` and `no` are possible.\n\n#\n# FFI\n#\n\n# Because Nim compiles to C, FFI is easy:\n\nproc strcmp(a, b: cstring): cint {.importc: \"strcmp\", nodecl.}\n\nlet cmp = strcmp(\"C?\", \"Easy!\")\n```\n\nAdditionally, Nim separates itself from its peers with metaprogramming,\nperformance, and compile-time features.\n\n## Further Reading\n\n* [Home Page](http://nim-lang.org)\n* [Download](http://nim-lang.org/download.html)\n* [Community](http://nim-lang.org/community.html)\n* [FAQ](http://nim-lang.org/question.html)\n* [Documentation](http://nim-lang.org/documentation.html)\n* [Manual](http://nim-lang.org/docs/manual.html)\n* [Standard Library](http://nim-lang.org/docs/lib.html)\n* [Rosetta Code](http://rosettacode.org/wiki/Category:Nim)"