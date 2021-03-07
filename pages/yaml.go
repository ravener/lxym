
package pages

const Yaml = "YAML is a data serialisation language designed to be directly writable and\nreadable by humans.\n\nIt's a strict superset of JSON, with the addition of syntactically\nsignificant newlines and indentation, like Python. Unlike Python, however,\nYAML doesn't allow literal tab characters for indentation.\n\n```yaml\n---  # document start\n\n# Comments in YAML look like this.\n\n################\n# SCALAR TYPES #\n################\n\n# Our root object (which continues for the entire document) will be a map,\n# which is equivalent to a dictionary, hash or object in other languages.\nkey: value\nanother_key: Another value goes here.\na_number_value: 100\nscientific_notation: 1e+12\n# The number 1 will be interpreted as a number, not a boolean. if you want\n# it to be interpreted as a boolean, use true\nboolean: true\nnull_value: null\nkey with spaces: value\n# Notice that strings don't need to be quoted. However, they can be.\nhowever: 'A string, enclosed in quotes.'\n'Keys can be quoted too.': \"Useful if you want to put a ':' in your key.\"\nsingle quotes: 'have ''one'' escape pattern'\ndouble quotes: \"have many: \\\", \\0, \\t, \\u263A, \\x0d\\x0a == \\r\\n, and more.\"\n# UTF-8/16/32 characters need to be encoded\nSuperscript two: \\u00B2\n\n# Multiple-line strings can be written either as a 'literal block' (using |),\n# or a 'folded block' (using '>').\nliteral_block: |\n    This entire block of text will be the value of the 'literal_block' key,\n    with line breaks being preserved.\n\n    The literal continues until de-dented, and the leading indentation is\n    stripped.\n\n        Any lines that are 'more-indented' keep the rest of their indentation -\n        these lines will be indented by 4 spaces.\nfolded_style: >\n    This entire block of text will be the value of 'folded_style', but this\n    time, all newlines will be replaced with a single space.\n\n    Blank lines, like above, are converted to a newline character.\n\n        'More-indented' lines keep their newlines, too -\n        this text will appear over two lines.\n\n####################\n# COLLECTION TYPES #\n####################\n\n# Nesting uses indentation. 2 space indent is preferred (but not required).\na_nested_map:\n  key: value\n  another_key: Another Value\n  another_nested_map:\n    hello: hello\n\n# Maps don't have to have string keys.\n0.25: a float key\n\n# Keys can also be complex, like multi-line objects\n# We use ? followed by a space to indicate the start of a complex key.\n? |\n  This is a key\n  that has multiple lines\n: and this is its value\n\n# YAML also allows mapping between sequences with the complex key syntax\n# Some language parsers might complain\n# An example\n? - Manchester United\n  - Real Madrid\n: [2001-01-01, 2002-02-02]\n\n# Sequences (equivalent to lists or arrays) look like this\n# (note that the '-' counts as indentation):\na_sequence:\n  - Item 1\n  - Item 2\n  - 0.5  # sequences can contain disparate types.\n  - Item 4\n  - key: value\n    another_key: another_value\n  -\n    - This is a sequence\n    - inside another sequence\n  - - - Nested sequence indicators\n      - can be collapsed\n\n# Since YAML is a superset of JSON, you can also write JSON-style maps and\n# sequences:\njson_map: {\"key\": \"value\"}\njson_seq: [3, 2, 1, \"takeoff\"]\nand quotes are optional: {key: [3, 2, 1, takeoff]}\n\n#######################\n# EXTRA YAML FEATURES #\n#######################\n\n# YAML also has a handy feature called 'anchors', which let you easily duplicate\n# content across your document. Both of these keys will have the same value:\nanchored_content: &anchor_name This string will appear as the value of two keys.\nother_anchor: *anchor_name\n\n# Anchors can be used to duplicate/inherit properties\nbase: &base\n  name: Everyone has same name\n\n# The regexp << is called Merge Key Language-Independent Type. It is used to\n# indicate that all the keys of one or more specified maps should be inserted\n# into the current map.\n\nfoo:\n  <<: *base\n  age: 10\n\nbar:\n  <<: *base\n  age: 20\n\n# foo and bar would also have name: Everyone has same name\n\n# YAML also has tags, which you can use to explicitly declare types.\nexplicit_string: !!str 0.5\n# Some parsers implement language specific tags, like this one for Python's\n# complex number type.\npython_complex_number: !!python/complex 1+2j\n\n# We can also use yaml complex keys with language specific tags\n? !!python/tuple [5, 7]\n: Fifty Seven\n# Would be {(5, 7): 'Fifty Seven'} in Python\n\n####################\n# EXTRA YAML TYPES #\n####################\n\n# Strings and numbers aren't the only scalars that YAML can understand.\n# ISO-formatted date and datetime literals are also parsed.\ndatetime: 2001-12-15T02:59:43.1Z\ndatetime_with_spaces: 2001-12-14 21:59:43.10 -5\ndate: 2002-12-14\n\n# The !!binary tag indicates that a string is actually a base64-encoded\n# representation of a binary blob.\ngif_file: !!binary |\n  R0lGODlhDAAMAIQAAP//9/X17unp5WZmZgAAAOfn515eXvPz7Y6OjuDg4J+fn5\n  OTk6enp56enmlpaWNjY6Ojo4SEhP/++f/++f/++f/++f/++f/++f/++f/++f/+\n  +f/++f/++f/++f/++f/++SH+Dk1hZGUgd2l0aCBHSU1QACwAAAAADAAMAAAFLC\n  AgjoEwnuNAFOhpEMTRiggcz4BNJHrv/zCFcLiwMWYNG84BwwEeECcgggoBADs=\n\n# YAML also has a set type, which looks like this:\nset:\n  ? item1\n  ? item2\n  ? item3\nor: {item1, item2, item3}\n\n# Sets are just maps with null values; the above is equivalent to:\nset2:\n  item1: null\n  item2: null\n  item3: null\n\n...  # document end\n```\n\n### More Resources\n\n+ [YAML official website](https://yaml.org/)\n+ [Online YAML Validator](http://www.yamllint.com/)"
