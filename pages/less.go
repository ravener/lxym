
package pages

const Less = "Less is a CSS pre-processor, that adds features such as variables, nesting, mixins and more.\nLess (and other preprocessors, such as [Sass](http://sass-lang.com/)) help developers to write maintainable and DRY (Don't Repeat Yourself) code.\n\n```css\n\n\n//Single line comments are removed when Less is compiled to CSS.\n\n/*Multi line comments are preserved. */\n\n\n\n/* Variables\n==============================*/\n\n\n/* You can store a CSS value (such as a color) in a variable.\n   Use the '@' symbol to create a variable. */\n\n@primary-color: #a3a4ff;\n@secondary-color: #51527f;\n@body-font: 'Roboto', sans-serif;\n\n/* You can use the variables throughout your stylesheet.\n   Now if you want to change a color, you only have to make the change once.*/\n\nbody {\n\tbackground-color: @primary-color;\n\tcolor: @secondary-color;\n\tfont-family: @body-font;\n}\n\n/* This would compile to: */\n\nbody {\n\tbackground-color: #a3a4ff;\n\tcolor: #51527F;\n\tfont-family: 'Roboto', sans-serif;\n}\n\n\n/* This is much more maintainable than having to change the color\n   each time it appears throughout your stylesheet. */\n\n\n\n/* Mixins\n==============================*/\n\n\n/* If you find you are writing the same code for more than one\n   element, you might want to reuse that easily.*/\n\n.center {\n\tdisplay: block;\n\tmargin-left: auto;\n\tmargin-right: auto;\n\tleft: 0;\n\tright: 0;\n}\n\n/* You can use the mixin by simply adding the selector as a style */\n\ndiv {\n\t.center;\n\tbackground-color: @primary-color;\n}\n\n/* Which would compile to: */\n\n.center {\n  display: block;\n  margin-left: auto;\n  margin-right: auto;\n  left: 0;\n  right: 0;\n}\ndiv {\n\tdisplay: block;\n\tmargin-left: auto;\n\tmargin-right: auto;\n\tleft: 0;\n\tright: 0;\n\tbackground-color: #a3a4ff;\n}\n\n/* You can omit the mixin code from being compiled by adding parenthesis\n   after the selector */\n\n.center() {\n  display: block;\n  margin-left: auto;\n  margin-right: auto;\n  left: 0;\n  right: 0;\n}\n\ndiv {\n  .center;\n  background-color: @primary-color;\n}\n\n/* Which would compile to: */\ndiv {\n  display: block;\n  margin-left: auto;\n  margin-right: auto;\n  left: 0;\n  right: 0;\n  background-color: #a3a4ff;\n}\n\n\n\n/* Nesting\n==============================*/\n\n\n/* Less allows you to nest selectors within selectors */\n\nul {\n\tlist-style-type: none;\n\tmargin-top: 2em;\n\n\tli {\n\t\tbackground-color: #f00;\n\t}\n}\n\n/* '&' will be replaced by the parent selector. */\n/* You can also nest pseudo-classes. */\n/* Keep in mind that over-nesting will make your code less maintainable.\n   Best practices recommend going no more than 3 levels deep when nesting.\n   For example: */\n\nul {\n\tlist-style-type: none;\n\tmargin-top: 2em;\n\n\tli {\n\t\tbackground-color: red;\n\n\t\t&:hover {\n\t\t  background-color: blue;\n\t\t}\n\n\t\ta {\n\t\t  color: white;\n\t\t}\n\t}\n}\n\n/* Compiles to: */\n\nul {\n  list-style-type: none;\n  margin-top: 2em;\n}\n\nul li {\n  background-color: red;\n}\n\nul li:hover {\n  background-color: blue;\n}\n\nul li a {\n  color: white;\n}\n\n\n\n/* Functions\n==============================*/\n\n\n/* Less provides functions that can be used to accomplish a variety of\n   tasks. Consider the following: */\n\n/* Functions can be invoked by using their name and passing in the\n   required arguments. */\n\nbody {\n  width: round(10.25px);\n}\n\n.header {\n\tbackground-color: lighten(#000, 0.5);\n}\n\n.footer {\n  background-color: fadeout(#000, 0.25)\n}\n\n/* Compiles to: */\n\nbody {\n  width: 10px;\n}\n\n.header {\n  background-color: #010101;\n}\n\n.footer {\n  background-color: rgba(0, 0, 0, 0.75);\n}\n\n/* You may also define your own functions. Functions are very similar to\n   mixins. When trying to choose between a function or a mixin, remember\n   that mixins are best for generating CSS while functions are better for\n   logic that might be used throughout your Less code. The examples in\n   the 'Math Operators' section are ideal candidates for becoming a reusable\n   function. */\n\n/* This function calculates the average of two numbers: */\n\n.average(@x, @y) {\n  @average-result: ((@x + @y) / 2);\n}\n\ndiv {\n  .average(16px, 50px); // \"call\" the mixin\n  padding: @average-result;    // use its \"return\" value\n}\n\n/* Compiles to: */\n\ndiv {\n  padding: 33px;\n}\n\n\n\n/*Extend (Inheritance)\n==============================*/\n\n\n/*Extend is a way to share the properties of one selector with another. */\n\n.display {\n  height: 50px;\n}\n\n.display-success {\n  &:extend(.display);\n\tborder-color: #22df56;\n}\n\n/* Compiles to: */\n.display,\n.display-success {\n  height: 50px;\n}\n.display-success {\n  border-color: #22df56;\n}\n\n/* Extending a CSS statement is preferable to creating a mixin\n   because of the way it groups together the classes that all share\n   the same base styling. If this was done with a mixin, the properties\n   would be duplicated for each statement that\n   called the mixin. While it won't affect your workflow, it will\n   add unnecessary bloat to the files created by the Less compiler. */\n\n\n\n/*Partials and Imports\n==============================*/\n\n\n/* Less allows you to create partial files. This can help keep your Less\n   code modularized. Partial files conventionally begin with an '_',\n   e.g. _reset.less. and are imported into a main less file that gets\n   compiled into CSS */\n\n/* Consider the following CSS which we'll put in a file called _reset.less */\n\nhtml,\nbody,\nul,\nol {\n  margin: 0;\n  padding: 0;\n}\n\n/* Less offers @import which can be used to import partials into a file.\n   This differs from the traditional CSS @import statement which makes\n   another HTTP request to fetch the imported file. Less takes the\n   imported file and combines it with the compiled code. */\n\n@import 'reset';\n\nbody {\n  font-size: 16px;\n  font-family: Helvetica, Arial, Sans-serif;\n}\n\n/* Compiles to: */\n\nhtml, body, ul, ol {\n  margin: 0;\n  padding: 0;\n}\n\nbody {\n  font-size: 16px;\n  font-family: Helvetica, Arial, Sans-serif;\n}\n\n\n\n/* Math Operations\n==============================*/\n\n\n/* Less provides the following operators: +, -, *, /, and %. These can\n   be useful for calculating values directly in your Less files instead\n   of using values that you've already calculated by hand. Below is an example\n   of a setting up a simple two column design. */\n\n@content-area: 960px;\n@main-content: 600px;\n@sidebar-content: 300px;\n\n@main-size: @main-content / @content-area * 100%;\n@sidebar-size: @sidebar-content / @content-area * 100%;\n@gutter: 100% - (@main-size + @sidebar-size);\n\nbody {\n  width: 100%;\n}\n\n.main-content {\n  width: @main-size;\n}\n\n.sidebar {\n  width: @sidebar-size;\n}\n\n.gutter {\n  width: @gutter;\n}\n\n/* Compiles to: */\n\nbody {\n  width: 100%;\n}\n\n.main-content {\n  width: 62.5%;\n}\n\n.sidebar {\n  width: 31.25%;\n}\n\n.gutter {\n  width: 6.25%;\n}\n\n\n```\n\n## Practice Less\n\nIf you want to play with Less in your browser, check out:\n* [Codepen](http://codepen.io/)\n* [LESS2CSS](http://lesscss.org/less-preview/)\n\n## Compatibility\n\nLess can be used in any project as long as you have a program to compile it into CSS. You'll want to verify that the CSS you're using is compatible with your target browsers.\n\n[QuirksMode CSS](http://www.quirksmode.org/css/) and [CanIUse](http://caniuse.com) are great resources for checking compatibility.\n\n## Further reading\n* [Official Documentation](http://lesscss.org/features/)\n* [Less CSS - Beginner's Guide](http://www.hongkiat.com/blog/less-basic/)"
