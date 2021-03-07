
package pages

const Html = "HTML stands for HyperText Markup Language.\n\nIt is a language which allows us to write pages for the world wide web.\nIt is a markup language, it enables us to write webpages using code to indicate\nhow text and data should be displayed.  In fact, html files are simple text\nfiles.\n\nWhat is this markup? It is a method of organising the page's data by\nsurrounding it with opening tags and closing tags.  This markup serves to give\nsignificance to the text that it encloses.  Like other computer languages, HTML\nhas many versions. Here we will talk about HTML5.\n\n**NOTE :**  You can test the different tags and elements as you progress through\nthe tutorial on a site like [codepen](http://codepen.io/pen/) in order to see\ntheir effects, understand how they work and familiarise yourself with the\nlanguage.  This article is concerned principally with HTML syntax and some\nuseful tips.\n\n\n```html\n<!-- Comments are enclosed like this line! -->\n\n<!--\n\tComments\n\tcan\n\tspan\n\tmultiple\n\tlines!\n-->\n\n<!-- #################### The Tags #################### -->\n\n<!-- Here is an example HTML file that we are going to analyse. -->\n\n\n<!doctype html>\n\t<html>\n\t\t<head>\n\t\t\t<title>My Site</title>\n\t\t</head>\n\t\t<body>\n\t\t\t<h1>Hello, world!</h1>\n\t\t\t<a href=\"http://codepen.io/anon/pen/xwjLbZ\">\n\t\t\t\tCome look at what this shows\n\t\t\t</a>\n\t\t\t<p>This is a paragraph.</p>\n\t\t\t<p>This is another paragraph.</p>\n\t\t\t<ul>\n\t\t\t\t<li>This is an item in a non-enumerated list (bullet list)</li>\n\t\t\t\t<li>This is another item</li>\n\t\t\t\t<li>And this is the last item on the list</li>\n\t\t\t</ul>\n\t\t</body>\n\t</html>\n\n<!--\n\tAn HTML file always starts by indicating to the browser that the page is HTML.\n-->\n<!doctype html>\n\n<!-- After this, it starts by opening an <html> tag. -->\n<html>\n\n<!-- that will be closed at the end of the file with </html>. -->\n</html>\n\n<!-- Nothing should appear after this final tag. -->\n\n<!-- Inside (between the opening and closing tags <html></html>), we find: -->\n\n<!-- A header defined by <head> (it must be closed with </head>). -->\n<!--\n\tThe header contains some description and additional information which are not\n\tdisplayed; this is metadata.\n-->\n\n<head>\n\t<!--\n\t\tThe tag <title> indicates to the browser the title to show in browser\n\t\twindow's title bar and tab name.\n\t-->\n\t<title>My Site</title>\n</head>\n\n<!-- After the <head> section, we find the tag - <body> -->\n<!-- Until this point, nothing described will show up in the browser window. -->\n<!-- We must fill the body with the content to be displayed. -->\n\n<body>\n\t<!-- The h1 tag creates a title. -->\n\t<h1>Hello, world!</h1>\n\t<!--\n\t\tThere are also subtitles to <h1> from the most important (h2) to the most\n\t\tprecise (h6).\n\t-->\n\n\t<!-- a hyperlink to the url given by the attribute href=\"\" -->\n\t<a href=\"http://codepen.io/anon/pen/xwjLbZ\">\n\t\tCome look at what this shows\n\t</a>\n\n\t<!-- The tag <p> lets us include text in the html page. -->\n\t<p>This is a paragraph.</p>\n\t<p>This is another paragraph.</p>\n\n\t<!-- The tag <ul> creates a bullet list. -->\n\t<!--\n\t\tTo have a numbered list instead we would use <ol> giving 1. for the first\n\t\telement, 2. for the second, etc.\n\t-->\n\t<ul>\n\t\t<li>This is an item in a non-enumerated list (bullet list)</li>\n\t\t<li>This is another item</li>\n\t\t<li>And this is the last item on the list</li>\n\t</ul>\n</body>\n\n<!-- And that's it, creating an HTML file can be simple. -->\n\n<!-- But it is possible to add many additional types of HTML tags. -->\n\n<!-- The <img /> tag is used to insert an image. -->\n<!--\n\tThe source of the image is indicated using the attribute src=\"\"\n\tThe source can be an URL or even path to a file on your computer.\n-->\n<img src=\"http://i.imgur.com/XWG0O.gif\"/>\n\n<!-- It is also possible to create a table. -->\n\n<!-- We open a <table> element. -->\n<table>\n\n\t<!-- <tr> allows us to create a row. -->\n\t<tr>\n\n\t\t<!-- <th> allows us to give a title to a table column. -->\n\t\t<th>First Header</th>\n\t\t<th>Second Header</th>\n\t</tr>\n\n\t<tr>\n\n\t\t<!-- <td> allows us to create a table cell. -->\n\t\t<td>first row, first column</td>\n\t\t<td>first row, second column</td>\n\t</tr>\n\n\t<tr>\n\t\t<td>second row, first column</td>\n\t\t<td>second row, second column</td>\n\t</tr>\n</table>\n\n```\n\n## Usage\n\nHTML is written in files ending with `.html` or `.htm`. The mime type is\n`text/html`.\n**HTML is NOT a programming language**\n## To Learn More\n\n* [wikipedia](https://en.wikipedia.org/wiki/HTML)\n* [HTML tutorial](https://developer.mozilla.org/en-US/docs/Web/HTML)\n* [W3School](http://www.w3schools.com/html/html_intro.asp)"
