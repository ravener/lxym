
package pages

const Pcre = "A regular expression (regex or regexp for short) is a special text string for describing a search pattern. e.g. to extract domain name from a string we can say `/^[a-z]+:/` and it will match `http:` from `http://github.com/`.\n\nPCRE (Perl Compatible Regular Expressions) is a C library implementing regex. It was written in 1997 when Perl was the de-facto choice for complex text processing tasks. The syntax for patterns used in PCRE closely resembles Perl. PCRE syntax is being used in many big projects including PHP, Apache, R to name a few.\n\n\nThere are two different sets of metacharacters:\n\n* Those that are recognized anywhere in the pattern except within square brackets\n\n```\n  \\      general escape character with several uses\n  ^      assert start of string (or line, in multiline mode)\n  $      assert end of string (or line, in multiline mode)\n  .      match any character except newline (by default)\n  [      start character class definition\n  |      start of alternative branch\n  (      start subpattern\n  )      end subpattern\n  ?      extends the meaning of (\n         also 0 or 1 quantifier\n         also quantifier minimizer\n  *      0 or more quantifier\n  +      1 or more quantifier\n         also \"possessive quantifier\"\n  {      start min/max quantifier\n```\n\n* Those that are recognized within square brackets. Outside square brackets. They are also called as character classes.\n\n```\n\n  \\      general escape character\n  ^      negate the class, but only if the first character\n  -      indicates character range\n  [      POSIX character class (only if followed by POSIX syntax)\n  ]      terminates the character class\n\n```\n\nPCRE provides some generic character types, also called as character classes.\n\n```\n  \\d     any decimal digit\n  \\D     any character that is not a decimal digit\n  \\h     any horizontal white space character\n  \\H     any character that is not a horizontal white space character\n  \\s     any white space character\n  \\S     any character that is not a white space character\n  \\v     any vertical white space character\n  \\V     any character that is not a vertical white space character\n  \\w     any \"word\" character\n  \\W     any \"non-word\" character\n```\n\n## Examples\n\nWe will test our examples on the following string:\n\n```\n66.249.64.13 - - [18/Sep/2004:11:07:48 +1000] \"GET /robots.txt HTTP/1.0\" 200 468 \"-\" \"Googlebot/2.1\"\n```\n\n It is a standard Apache access log.\n\n| Regex | Result          | Comment |\n| :---- | :-------------- | :------ |\n| `GET`   | GET | GET matches the characters GET literally (case sensitive) |\n| `\\d+.\\d+.\\d+.\\d+` | 66.249.64.13 | `\\d+` match a digit [0-9] one or more times defined by `+` quantifier, `\\.` matches `.` literally |\n| `(\\d+\\.){3}\\d+` | 66.249.64.13 | `(\\d+\\.){3}` is trying to match group (`\\d+\\.`) exactly three times. |\n| `\\[.+\\]` | [18/Sep/2004:11:07:48 +1000] | `.+` matches any character (except newline), `.` is any character |\n| `^\\S+` | 66.249.64.13 | `^` means start of the line, `\\S+` matches any number of non-space characters |\n| `\\+[0-9]+` | +1000 | `\\+` matches the character `+` literally. `[0-9]` character class means single number. Same can be achieved using `\\+\\d+` |\n\n## Further Reading\n[Regex101](https://regex101.com/) - Regular Expression tester and debugger"
