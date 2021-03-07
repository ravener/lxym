
package pages

const Groovy = "Groovy - A dynamic language for the Java platform [Read more here.](http://www.groovy-lang.org/)\n\n```groovy\n\n/*\n  Set yourself up:\n\n  1) Install SDKMAN - http://sdkman.io/\n  2) Install Groovy: sdk install groovy\n  3) Start the groovy console by typing: groovyConsole\n\n*/\n\n//  Single line comments start with two forward slashes\n/*\nMulti line comments look like this.\n*/\n\n// Hello World\nprintln \"Hello world!\"\n\n/*\n  Variables:\n\n  You can assign values to variables for later use\n*/\n\ndef x = 1\nprintln x\n\nx = new java.util.Date()\nprintln x\n\nx = -3.1499392\nprintln x\n\nx = false\nprintln x\n\nx = \"Groovy!\"\nprintln x\n\n/*\n  Collections and maps\n*/\n\n//Creating an empty list\ndef technologies = []\n\n/*** Adding a elements to the list ***/\n\n// As with Java\ntechnologies.add(\"Grails\")\n\n// Left shift adds, and returns the list\ntechnologies << \"Groovy\"\n\n// Add multiple elements\ntechnologies.addAll([\"Gradle\",\"Griffon\"])\n\n/*** Removing elements from the list ***/\n\n// As with Java\ntechnologies.remove(\"Griffon\")\n\n// Subtraction works also\ntechnologies = technologies - 'Grails'\n\n/*** Iterating Lists ***/\n\n// Iterate over elements of a list\ntechnologies.each { println \"Technology: $it\"}\ntechnologies.eachWithIndex { it, i -> println \"$i: $it\"}\n\n/*** Checking List contents ***/\n\n//Evaluate if a list contains element(s) (boolean)\ncontained = technologies.contains( 'Groovy' )\n\n// Or\ncontained = 'Groovy' in technologies\n\n// Check for multiple contents\ntechnologies.containsAll(['Groovy','Grails'])\n\n/*** Sorting Lists ***/\n\n// Sort a list (mutates original list)\ntechnologies.sort()\n\n// To sort without mutating original, you can do:\nsortedTechnologies = technologies.sort( false )\n\n/*** Manipulating Lists ***/\n\n//Replace all elements in the list\nCollections.replaceAll(technologies, 'Gradle', 'gradle')\n\n//Shuffle a list\nCollections.shuffle(technologies, new Random())\n\n//Clear a list\ntechnologies.clear()\n\n//Creating an empty map\ndef devMap = [:]\n\n//Add values\ndevMap = ['name':'Roberto', 'framework':'Grails', 'language':'Groovy']\ndevMap.put('lastName','Perez')\n\n//Iterate over elements of a map\ndevMap.each { println \"$it.key: $it.value\" }\ndevMap.eachWithIndex { it, i -> println \"$i: $it\"}\n\n//Evaluate if a map contains a key\nassert devMap.containsKey('name')\n\n//Evaluate if a map contains a value\nassert devMap.containsValue('Roberto')\n\n//Get the keys of a map\nprintln devMap.keySet()\n\n//Get the values of a map\nprintln devMap.values()\n\n/*\n  Groovy Beans\n\n  GroovyBeans are JavaBeans but using a much simpler syntax\n\n  When Groovy is compiled to bytecode, the following rules are used.\n\n    * If the name is declared with an access modifier (public, private or\n      protected) then a field is generated.\n\n    * A name declared with no access modifier generates a private field with\n      public getter and setter (i.e. a property).\n\n    * If a property is declared final the private field is created final and no\n      setter is generated.\n\n    * You can declare a property and also declare your own getter or setter.\n\n    * You can declare a property and a field of the same name, the property will\n      use that field then.\n\n    * If you want a private or protected property you have to provide your own\n      getter and setter which must be declared private or protected.\n\n    * If you access a property from within the class the property is defined in\n      at compile time with implicit or explicit this (for example this.foo, or\n      simply foo), Groovy will access the field directly instead of going though\n      the getter and setter.\n\n    * If you access a property that does not exist using the explicit or\n      implicit foo, then Groovy will access the property through the meta class,\n      which may fail at runtime.\n\n*/\n\nclass Foo {\n    // read only property\n    final String name = \"Roberto\"\n\n    // read only property with public getter and protected setter\n    String language\n    protected void setLanguage(String language) { this.language = language }\n\n    // dynamically typed property\n    def lastName\n}\n\n/*\n  Methods with optional parameters\n*/\n\n// A method can have default values for parameters\ndef say(msg = 'Hello', name = 'world') {\n    \"$msg $name!\"\n}\n\n// It can be called in 3 different ways\nassert 'Hello world!' == say()\n// Right most parameter with default value is eliminated first.\nassert 'Hi world!' == say('Hi')\nassert 'learn groovy' == say('learn', 'groovy')\n\n/*\n  Logical Branching and Looping\n*/\n\n//Groovy supports the usual if - else syntax\ndef x = 3\n\nif(x==1) {\n    println \"One\"\n} else if(x==2) {\n    println \"Two\"\n} else {\n    println \"X greater than Two\"\n}\n\n//Groovy also supports the ternary operator:\ndef y = 10\ndef x = (y > 1) ? \"worked\" : \"failed\"\nassert x == \"worked\"\n\n//Groovy supports 'The Elvis Operator' too!\n//Instead of using the ternary operator:\n\ndisplayName = user.name ? user.name : 'Anonymous'\n\n//We can write it:\ndisplayName = user.name ?: 'Anonymous'\n\n//For loop\n//Iterate over a range\ndef x = 0\nfor (i in 0 .. 30) {\n    x += i\n}\n\n//Iterate over a list\nx = 0\nfor( i in [5,3,2,1] ) {\n    x += i\n}\n\n//Iterate over an array\narray = (0..20).toArray()\nx = 0\nfor (i in array) {\n    x += i\n}\n\n//Iterate over a map\ndef map = ['name':'Roberto', 'framework':'Grails', 'language':'Groovy']\nx = \"\"\nfor ( e in map ) {\n    x += e.value\n    x += \" \"\n}\nassert x.equals(\"Roberto Grails Groovy \")\n\n/*\n  Operators\n\n  Operator Overloading for a list of the common operators that Groovy supports:\n  http://www.groovy-lang.org/operators.html#Operator-Overloading\n\n  Helpful groovy operators\n*/\n//Spread operator:  invoke an action on all items of an aggregate object.\ndef technologies = ['Groovy','Grails','Gradle']\ntechnologies*.toUpperCase() // = to technologies.collect { it?.toUpperCase() }\n\n//Safe navigation operator: used to avoid a NullPointerException.\ndef user = User.get(1)\ndef username = user?.username\n\n\n/*\n  Closures\n  A Groovy Closure is like a \"code block\" or a method pointer. It is a piece of\n  code that is defined and then executed at a later point.\n\n  More info at: http://www.groovy-lang.org/closures.html\n*/\n//Example:\ndef clos = { println \"Hello World!\" }\n\nprintln \"Executing the Closure:\"\nclos()\n\n//Passing parameters to a closure\ndef sum = { a, b -> println a+b }\nsum(2,4)\n\n//Closures may refer to variables not listed in their parameter list.\ndef x = 5\ndef multiplyBy = { num -> num * x }\nprintln multiplyBy(10)\n\n// If you have a Closure that takes a single argument, you may omit the\n// parameter definition of the Closure\ndef clos = { print it }\nclos( \"hi\" )\n\n/*\n  Groovy can memoize closure results [1][2][3]\n*/\ndef cl = {a, b ->\n    sleep(3000) // simulate some time consuming processing\n    a + b\n}\n\nmem = cl.memoize()\n\ndef callClosure(a, b) {\n    def start = System.currentTimeMillis()\n    mem(a, b)\n    println \"Inputs(a = $a, b = $b) - took ${System.currentTimeMillis() - start} msecs.\"\n}\n\ncallClosure(1, 2)\ncallClosure(1, 2)\ncallClosure(2, 3)\ncallClosure(2, 3)\ncallClosure(3, 4)\ncallClosure(3, 4)\ncallClosure(1, 2)\ncallClosure(2, 3)\ncallClosure(3, 4)\n\n/*\n  Expando\n\n  The Expando class is a dynamic bean so we can add properties and we can add\n  closures as methods to an instance of this class\n\n  http://mrhaki.blogspot.mx/2009/10/groovy-goodness-expando-as-dynamic-bean.html\n*/\n  def user = new Expando(name:\"Roberto\")\n  assert 'Roberto' == user.name\n\n  user.lastName = 'Pérez'\n  assert 'Pérez' == user.lastName\n\n  user.showInfo = { out ->\n      out << \"Name: $name\"\n      out << \", Last name: $lastName\"\n  }\n\n  def sw = new StringWriter()\n  println user.showInfo(sw)\n\n\n/*\n  Metaprogramming (MOP)\n*/\n\n//Using ExpandoMetaClass to add behaviour\nString.metaClass.testAdd = {\n    println \"we added this\"\n}\n\nString x = \"test\"\nx?.testAdd()\n\n//Intercepting method calls\nclass Test implements GroovyInterceptable {\n    def sum(Integer x, Integer y) { x + y }\n\n    def invokeMethod(String name, args) {\n        System.out.println \"Invoke method $name with args: $args\"\n    }\n}\n\ndef test = new Test()\ntest?.sum(2,3)\ntest?.multiply(2,3)\n\n//Groovy supports propertyMissing for dealing with property resolution attempts.\nclass Foo {\n   def propertyMissing(String name) { name }\n}\ndef f = new Foo()\n\nassertEquals \"boo\", f.boo\n\n/*\n  TypeChecked and CompileStatic\n  Groovy, by nature, is and will always be a dynamic language but it supports\n  typechecked and compilestatic\n\n  More info: http://www.infoq.com/articles/new-groovy-20\n*/\n//TypeChecked\nimport groovy.transform.TypeChecked\n\nvoid testMethod() {}\n\n@TypeChecked\nvoid test() {\n    testMeethod()\n\n    def name = \"Roberto\"\n\n    println naameee\n\n}\n\n//Another example:\nimport groovy.transform.TypeChecked\n\n@TypeChecked\nInteger test() {\n    Integer num = \"1\"\n\n    Integer[] numbers = [1,2,3,4]\n\n    Date date = numbers[1]\n\n    return \"Test\"\n\n}\n\n//CompileStatic example:\nimport groovy.transform.CompileStatic\n\n@CompileStatic\nint sum(int x, int y) {\n    x + y\n}\n\nassert sum(2,5) == 7\n\n\n```\n\n## Further resources\n\n[Groovy documentation](http://www.groovy-lang.org/documentation.html)\n\n[Groovy web console](http://groovyconsole.appspot.com/)\n\nJoin a [Groovy user group](http://www.groovy-lang.org/usergroups.html)\n\n## Books\n\n* [Groovy Goodness] (https://leanpub.com/groovy-goodness-notebook)\n\n* [Groovy in Action] (http://manning.com/koenig2/)\n\n* [Programming Groovy 2: Dynamic Productivity for the Java Developer] (http://shop.oreilly.com/product/9781937785307.do)\n\n[1] http://roshandawrani.wordpress.com/2010/10/18/groovy-new-feature-closures-can-now-memorize-their-results/\n[2] http://www.solutionsiq.com/resources/agileiq-blog/bid/72880/Programming-with-Groovy-Trampoline-and-Memoize\n[3] http://mrhaki.blogspot.mx/2011/05/groovy-goodness-cache-closure-results.html"
