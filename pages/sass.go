
package pages

const Sass = "Sass is a CSS extension language that adds features such as variables, nesting, mixins and more.\nSass (and other preprocessors, such as [Less](http://lesscss.org/)) help developers write maintainable and DRY (Don't Repeat Yourself) code.\n\nSass has two different syntax options to choose from. SCSS, which has the same syntax as CSS but with the added features of Sass. Or Sass (the original syntax), which uses indentation rather than curly braces and semicolons.\nThis tutorial is written using SCSS.\n\nIf you're already familiar with CSS3, you'll be able to pick up Sass relatively quickly. It does not provide any new styling properties but rather the tools to write your CSS more efficiently and make maintenance much easier.\n\n```scss\n\n\n//Single line comments are removed when Sass is compiled to CSS.\n\n/* Multi line comments are preserved. */\n\n\n\n/* Variables\n============================== */\n\n\n\n/* You can store a CSS value (such as a color) in a variable.\nUse the '$' symbol to create a variable. */\n\n$primary-color: #A3A4FF;\n$secondary-color: #51527F;\n$body-font: 'Roboto', sans-serif;\n\n/* You can use the variables throughout your stylesheet.\nNow if you want to change a color, you only have to make the change once. */\n\nbody {\n\tbackground-color: $primary-color;\n\tcolor: $secondary-color;\n\tfont-family: $body-font;\n}\n\n/* This would compile to: */\nbody {\n\tbackground-color: #A3A4FF;\n\tcolor: #51527F;\n\tfont-family: 'Roboto', sans-serif;\n}\n\n/* This is much more maintainable than having to change the color\neach time it appears throughout your stylesheet. */\n\n\n\n/* Control Directives\n============================== */\n\n/* Sass lets you use @if, @else, @for, @while, and @each to control the\n   compilation of your code to CSS. */\n\n/* @if/@else blocks behave exactly as you might expect */\n\n$debug: true !default;\n\n@mixin debugmode {\n\t@if $debug {\n\t\t@debug \"Debug mode enabled\";\n\n\t\tdisplay: inline-block;\n\t}\n\t@else {\n\t\tdisplay: none;\n\t}\n}\n\n.info {\n\t@include debugmode;\n}\n\n/* If $debug is set to true, .info is displayed; if it's set to false then\n.info is not displayed.\n\nNote: @debug will output debugging information to the command line.\nUseful for checking variables while debugging your SCSS. */\n\n.info {\n\tdisplay: inline-block;\n}\n\n/* @for is a control loop that iterates through a range of values.\nParticularly useful for setting styles on a collection of items.\nThere are two forms, \"through\" and \"to\". The former includes the last value,\nthe latter stops at the last value. */\n\n@for $c from 1 to 4 {\n\tdiv:nth-of-type(#{$c}) {\n\t\tleft: ($c - 1) * 900 / 3;\n\t}\n}\n\n@for $c from 1 through 3 {\n\t.myclass-#{$c} {\n\t\tcolor: rgb($c * 255 / 3, $c * 255 / 3, $c * 255 / 3);\n\t}\n}\n\n/* Will compile to: */\n\ndiv:nth-of-type(1) {\n\tleft: 0;\n}\n\ndiv:nth-of-type(2) {\n\tleft: 300;\n}\n\ndiv:nth-of-type(3) {\n\tleft: 600;\n}\n\n.myclass-1 {\n\tcolor: #555555;\n}\n\n.myclass-2 {\n\tcolor: #aaaaaa;\n}\n\n.myclass-3 {\n\tcolor: white;\n// SASS automatically converts #FFFFFF to white\n}\n\n/* @while is very straightforward: */\n\n$columns: 4;\n$column-width: 80px;\n\n@while $columns > 0 {\n\t.col-#{$columns} {\n\t\twidth: $column-width;\n\t\tleft: $column-width * ($columns - 1);\n\t}\n\n\t$columns: $columns - 1;\n}\n\n/* Will output the following CSS: */\n\n.col-4 {\n\twidth: 80px;\n\tleft: 240px;\n}\n\n.col-3 {\n\twidth: 80px;\n\tleft: 160px;\n}\n\n.col-2 {\n\twidth: 80px;\n\tleft: 80px;\n}\n\n.col-1 {\n\twidth: 80px;\n\tleft: 0px;\n}\n\n/* @each functions like @for, except using a list instead of ordinal values\nNote: you specify lists just like other variables, with spaces as\ndelimiters. */\n\n$social-links: facebook twitter linkedin reddit;\n\n.social-links {\n\t@each $sm in $social-links {\n\t\t.icon-#{$sm} {\n\t\t\tbackground-image: url(\"images/#{$sm}.png\");\n\t\t}\n\t}\n}\n\n/* Which will output: */\n\n.social-links .icon-facebook {\n\tbackground-image: url(\"images/facebook.png\");\n}\n\n.social-links .icon-twitter {\n\tbackground-image: url(\"images/twitter.png\");\n}\n\n.social-links .icon-linkedin {\n\tbackground-image: url(\"images/linkedin.png\");\n}\n\n.social-links .icon-reddit {\n\tbackground-image: url(\"images/reddit.png\");\n}\n\n\n/* Mixins\n==============================*/\n\n/* If you find you are writing the same code for more than one\nelement, you might want to store that code in a mixin.\n\nUse the '@mixin' directive, plus a name for your mixin. */\n\n@mixin center {\n\tdisplay: block;\n\tmargin-left: auto;\n\tmargin-right: auto;\n\tleft: 0;\n\tright: 0;\n}\n\n/* You can use the mixin with '@include' and the mixin name. */\n\ndiv {\n\t@include center;\n\tbackground-color: $primary-color;\n}\n\n/* Which would compile to: */\ndiv {\n\tdisplay: block;\n\tmargin-left: auto;\n\tmargin-right: auto;\n\tleft: 0;\n\tright: 0;\n\tbackground-color: #A3A4FF;\n}\n\n/* You can use mixins to create a shorthand property. */\n\n@mixin size($width, $height) {\n\twidth: $width;\n\theight: $height;\n}\n\n/* Which you can invoke by passing width and height arguments. */\n\n.rectangle {\n\t@include size(100px, 60px);\n}\n\n.square {\n\t@include size(40px, 40px);\n}\n\n/* Compiles to: */\n.rectangle {\n  width: 100px;\n  height: 60px;\n}\n\n.square {\n  width: 40px;\n  height: 40px;\n}\n\n\n\n/* Functions\n============================== */\n\n\n\n/* Sass provides functions that can be used to accomplish a variety of\n   tasks. Consider the following */\n\n/* Functions can be invoked by using their name and passing in the\n   required arguments */\nbody {\n  width: round(10.25px);\n}\n\n.footer {\n  background-color: fade_out(#000000, 0.25);\n}\n\n/* Compiles to: */\n\nbody {\n  width: 10px;\n}\n\n.footer {\n  background-color: rgba(0, 0, 0, 0.75);\n}\n\n/* You may also define your own functions. Functions are very similar to\n   mixins. When trying to choose between a function or a mixin, remember\n   that mixins are best for generating CSS while functions are better for\n   logic that might be used throughout your Sass code. The examples in\n   the 'Math Operators' section are ideal candidates for becoming a reusable\n   function. */\n\n/* This function will take a target size and the parent size and calculate\n   and return the percentage */\n\n@function calculate-percentage($target-size, $parent-size) {\n  @return $target-size / $parent-size * 100%;\n}\n\n$main-content: calculate-percentage(600px, 960px);\n\n.main-content {\n  width: $main-content;\n}\n\n.sidebar {\n  width: calculate-percentage(300px, 960px);\n}\n\n/* Compiles to: */\n\n.main-content {\n  width: 62.5%;\n}\n\n.sidebar {\n  width: 31.25%;\n}\n\n\n\n/* Extend (Inheritance)\n============================== */\n\n\n\n/* Extend is a way to share the properties of one selector with another. */\n\n.display {\n\t@include size(5em, 5em);\n\tborder: 5px solid $secondary-color;\n}\n\n.display-success {\n\t@extend .display;\n\tborder-color: #22df56;\n}\n\n/* Compiles to: */\n.display, .display-success {\n  width: 5em;\n  height: 5em;\n  border: 5px solid #51527F;\n}\n\n.display-success {\n  border-color: #22df56;\n}\n\n/* Extending a CSS statement is preferable to creating a mixin\n   because of the way Sass groups together the classes that all share\n   the same base styling. If this was done with a mixin, the width,\n   height, and border would be duplicated for each statement that\n   called the mixin. While it won't affect your workflow, it will\n   add unnecessary bloat to the files created by the Sass compiler. */\n\n\n\n/* Nesting\n============================== */\n\n\n\n/* Sass allows you to nest selectors within selectors */\n\nul {\n\tlist-style-type: none;\n\tmargin-top: 2em;\n\n\tli {\n\t\tbackground-color: #FF0000;\n\t}\n}\n\n/* '&' will be replaced by the parent selector. */\n/* You can also nest pseudo-classes. */\n/* Keep in mind that over-nesting will make your code less maintainable.\nBest practices recommend going no more than 3 levels deep when nesting.\nFor example: */\n\nul {\n\tlist-style-type: none;\n\tmargin-top: 2em;\n\n\tli {\n\t\tbackground-color: red;\n\n\t\t&:hover {\n\t\t  background-color: blue;\n\t\t}\n\n\t\ta {\n\t\t  color: white;\n\t\t}\n\t}\n}\n\n/* Compiles to: */\n\nul {\n  list-style-type: none;\n  margin-top: 2em;\n}\n\nul li {\n  background-color: red;\n}\n\nul li:hover {\n  background-color: blue;\n}\n\nul li a {\n  color: white;\n}\n\n\n\n/* Partials and Imports\n============================== */\n\n\n\n/* Sass allows you to create partial files. This can help keep your Sass\n   code modularized. Partial files should begin with an '_', e.g. _reset.css.\n   Partials are not generated into CSS. */\n\n/* Consider the following CSS which we'll put in a file called _reset.css */\n\nhtml,\nbody,\nul,\nol {\n  margin: 0;\n  padding: 0;\n}\n\n/* Sass offers @import which can be used to import partials into a file.\n   This differs from the traditional CSS @import statement which makes\n   another HTTP request to fetch the imported file. Sass takes the\n   imported file and combines it with the compiled code. */\n\n@import 'reset';\n\nbody {\n  font-size: 16px;\n  font-family: Helvetica, Arial, Sans-serif;\n}\n\n/* Compiles to: */\n\nhtml, body, ul, ol {\n  margin: 0;\n  padding: 0;\n}\n\nbody {\n  font-size: 16px;\n  font-family: Helvetica, Arial, Sans-serif;\n}\n\n\n\n/* Placeholder Selectors\n============================== */\n\n\n\n/* Placeholders are useful when creating a CSS statement to extend. If you\n   wanted to create a CSS statement that was exclusively used with @extend,\n   you can do so using a placeholder. Placeholders begin with a '%' instead\n   of '.' or '#'. Placeholders will not appear in the compiled CSS. */\n\n%content-window {\n  font-size: 14px;\n  padding: 10px;\n  color: #000;\n  border-radius: 4px;\n}\n\n.message-window {\n  @extend %content-window;\n  background-color: #0000ff;\n}\n\n/* Compiles to: */\n\n.message-window {\n  font-size: 14px;\n  padding: 10px;\n  color: #000;\n  border-radius: 4px;\n}\n\n.message-window {\n  background-color: #0000ff;\n}\n\n\n\n/* Math Operations\n============================== */\n\n\n\n/* Sass provides the following operators: +, -, *, /, and %. These can\n   be useful for calculating values directly in your Sass files instead\n   of using values that you've already calculated by hand. Below is an example\n   of a setting up a simple two column design. */\n\n$content-area: 960px;\n$main-content: 600px;\n$sidebar-content: 300px;\n\n$main-size: $main-content / $content-area * 100%;\n$sidebar-size: $sidebar-content / $content-area * 100%;\n$gutter: 100% - ($main-size + $sidebar-size);\n\nbody {\n  width: 100%;\n}\n\n.main-content {\n  width: $main-size;\n}\n\n.sidebar {\n  width: $sidebar-size;\n}\n\n.gutter {\n  width: $gutter;\n}\n\n/* Compiles to: */\n\nbody {\n  width: 100%;\n}\n\n.main-content {\n  width: 62.5%;\n}\n\n.sidebar {\n  width: 31.25%;\n}\n\n.gutter {\n  width: 6.25%;\n}\n\n```\n\n## SASS or Sass?\nHave you ever wondered whether Sass is an acronym or not? You probably haven't, but I'll tell you anyway. The name of the language is a word, \"Sass\", and not an acronym.\nBecause people were constantly writing it as \"SASS\", the creator of the language jokingly called it \"Syntactically Awesome StyleSheets\".\n\n\n## Practice Sass\nIf you want to play with Sass in your browser, check out [SassMeister](http://sassmeister.com/).\nYou can use either syntax, just go into the settings and select either Sass or SCSS.\n\n\n## Compatibility\nSass can be used in any project as long as you have a program to compile it into CSS. You'll want to verify that the CSS you're using is compatible with your target browsers.\n\n[QuirksMode CSS](http://www.quirksmode.org/css/) and [CanIUse](http://caniuse.com) are great resources for checking compatibility.\n\n\n## Further reading\n* [Official Documentation](http://sass-lang.com/documentation/file.SASS_REFERENCE.html)\n* [The Sass Way](http://thesassway.com/) provides tutorials (beginner-advanced) and articles."
