
package pages

const Edn = "Extensible Data Notation (EDN) is a format for serializing data.\n\nEDN is a subset of the syntax used by Clojure. Reading data defined by EDN is\nsafer than that defined by the full Clojure syntax, especially from untrusted\nsources. EDN is restricted to data, no code. It is similar in intent to JSON.\nThough it is more commonly used in Clojure, there are implementations of EDN\nfor many other languages.\n\nThe main benefit of EDN over JSON and YAML is that it is extensible. We\nwill see how it is extended later on.\n\n```clojure\n; Comments start with a semicolon.\n; Anything after the semicolon is ignored.\n\n;;;;;;;;;;;;;;;;;;;\n;;; Basic Types ;;;\n;;;;;;;;;;;;;;;;;;;\n\nnil         ; also known in other languages as null\n\n; Booleans\ntrue\nfalse\n\n; Strings are enclosed in double quotes\n\"hungarian breakfast\"\n\"farmer's cheesy omelette\"\n\n; Characters are preceded by backslashes\n\\g \\r \\a \\c \\e\n\n; Keywords start with a colon. They behave like enums. Kind of\n; like symbols in Ruby.\n:eggs\n:cheese\n:olives\n\n; Symbols are used to represent identifiers. \n; You can namespace symbols by using /. Whatever precedes / is\n; the namespace of the symbol.\nspoon\nkitchen/spoon ; not the same as spoon\nkitchen/fork\ngithub/fork   ; you can't eat with this\n\n; Integers and floats\n42\n3.14159\n\n; Lists are sequences of values\n(:bun :beef-patty 9 \"yum!\")\n\n; Vectors allow random access\n[:gelato 1 2 -2]\n\n; Maps are associative data structures that associate the key with its value\n{:eggs        2\n :lemon-juice 3.5\n :butter      1}\n\n; You're not restricted to using keywords as keys\n{[1 2 3 4] \"tell the people what she wore\",\n [5 6 7 8] \"the more you see the more you hate\"}\n\n; You may use commas for readability. They are treated as whitespace.\n\n; Sets are collections that contain unique elements.\n#{:a :b 88 \"huat\"}\n\n;;;;;;;;;;;;;;;;;;;;;;;\n;;; Tagged Elements ;;;\n;;;;;;;;;;;;;;;;;;;;;;;\n\n; EDN can be extended by tagging elements with # symbols.\n\n#MyYelpClone/MenuItem {:name \"eggs-benedict\" :rating 10}\n\n; Let me explain this with a Clojure example. Suppose I want to transform that\n; piece of EDN into a MenuItem record.\n\n(defrecord MenuItem [name rating])\n\n; defrecord defined, among other things, map->MenuItem which will take a map\n; of field names (as keywords) to values and generate a user.MenuItem record\n\n; To transform EDN to Clojure values, I will need to use the built-in EDN\n; reader, clojure.edn/read-string\n\n(clojure.edn/read-string \"{:eggs 2 :butter 1 :flour 5}\")\n; -> {:eggs 2 :butter 1 :flour 5}\n\n; To transform tagged elements, pass to clojure.edn/read-string an option map\n; with a :readers map that maps tag symbols to data-reader functions, like so\n\n(clojure.edn/read-string\n    {:readers {'MyYelpClone/MenuItem map->MenuItem}}\n    \"#MyYelpClone/MenuItem {:name \\\"eggs-benedict\\\" :rating 10}\")\n; -> #user.MenuItem{:name \"eggs-benedict\", :rating 10}\n\n```\n\n# References\n\n- [EDN spec](https://github.com/edn-format/edn)\n- [Implementations](https://github.com/edn-format/edn/wiki/Implementations)\n- [Tagged Elements](http://www.compoundtheory.com/clojure-edn-walkthrough/)"
