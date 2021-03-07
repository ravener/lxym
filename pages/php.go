
package pages

const Php = "This document describes PHP 5+.\n\n```php\n<?php // PHP code must be enclosed with <?php tags\n\n// If your php file only contains PHP code, it is best practice\n// to omit the php closing tag to prevent accidental output.\n\n// Two forward slashes start a one-line comment.\n\n# So will a hash (aka pound symbol) but // is more common\n\n/*\n     Surrounding text in slash-asterisk and asterisk-slash\n     makes it a multi-line comment.\n*/\n\n// Use \"echo\" or \"print\" to print output\nprint('Hello '); // Prints \"Hello \" with no line break\n\n// () are optional for print and echo\necho \"World\\n\"; // Prints \"World\" with a line break\n// (all statements must end with a semicolon)\n\n// Anything outside <?php tags is echoed automatically\n?>\nHello World Again!\n<?php\n// That is because historically PHP started as a Template engine\n\n\n/************************************\n * Types & Variables\n */\n\n// Variables begin with the $ symbol.\n// A valid variable name starts with a letter or an underscore,\n// followed by any number of letters, numbers, or underscores.\n\n// You don't have to (and cannot) declare variables.\n// Once you assign a value, PHP will create the variable with the right type.\n\n// Boolean values are case-insensitive\n$boolean = true;  // or TRUE or True\n$boolean = FALSE; // or false or False\n\n// Integers\n$int1 = 12;   // => 12\n$int2 = -12;  // => -12\n$int3 = 012;  // => 10 (a leading 0 denotes an octal number)\n$int4 = 0x0F; // => 15 (a leading 0x denotes a hex literal)\n// Binary integer literals are available since PHP 5.4.0.\n$int5 = 0b11111111; // 255 (a leading 0b denotes a binary number)\n\n// Floats (aka doubles)\n$float = 1.234;\n$float = 1.2e3;\n$float = 7E-10;\n\n// Delete variable\nunset($int1);\n\n// Arithmetic\n$sum        = 1 + 1; // 2\n$difference = 2 - 1; // 1\n$product    = 2 * 2; // 4\n$quotient   = 2 / 1; // 2\n\n// Shorthand arithmetic\n$number = 0;\n$number += 1;      // Increment $number by 1\necho $number++;    // Prints 1 (increments after evaluation)\necho ++$number;    // Prints 3 (increments before evaluation)\n$number /= $float; // Divide and assign the quotient to $number\n\n// Strings should be enclosed in single quotes;\n$sgl_quotes = '$String'; // => '$String'\n\n// Avoid using double quotes except to embed other variables\n$dbl_quotes = \"This is a $sgl_quotes.\"; // => 'This is a $String.'\n\n// Special characters are only escaped in double quotes\n$escaped   = \"This contains a \\t tab character.\";\n$unescaped = 'This just contains a slash and a t: \\t';\n\n// Enclose a variable in curly braces if needed\n$apples = \"I have {$number} apples to eat.\";\n$oranges = \"I have ${number} oranges to eat.\";\n$money = \"I have $${number} in the bank.\";\n\n// Since PHP 5.3, nowdocs can be used for uninterpolated multi-liners\n$nowdoc = <<<'END'\nMulti line\nstring\nEND;\n\n// Heredocs will do string interpolation\n$heredoc = <<<END\nMulti line\n$sgl_quotes\nEND;\n\n// String concatenation is done with .\necho 'This string ' . 'is concatenated';\n\n// Strings can be passed in as parameters to echo\necho 'Multiple', 'Parameters', 'Valid';  // Returns 'MultipleParametersValid'\n\n\n/********************************\n * Constants\n */\n\n// A constant is defined by using define()\n// and can never be changed during runtime!\n\n// a valid constant name starts with a letter or underscore,\n// followed by any number of letters, numbers, or underscores.\ndefine(\"FOO\", \"something\");\n\n// access to a constant is possible by calling the chosen name without a $\necho FOO; // Returns 'something'\necho 'This outputs ' . FOO;  // Returns 'This outputs something'\n\n\n\n/********************************\n * Arrays\n */\n\n// All arrays in PHP are associative arrays (hashmaps in some languages)\n\n// Works with all PHP versions\n$associative = array('One' => 1, 'Two' => 2, 'Three' => 3);\n\n// PHP 5.4 introduced a new syntax\n$associative = ['One' => 1, 'Two' => 2, 'Three' => 3];\n\necho $associative['One']; // prints 1\n\n// Add an element to an associative array\n$associative['Four'] = 4;\n\n// List literals implicitly assign integer keys\n$array = ['One', 'Two', 'Three'];\necho $array[0]; // => \"One\"\n\n// Add an element to the end of an array\n$array[] = 'Four';\n// or\narray_push($array, 'Five');\n\n// Remove element from array\nunset($array[3]);\n\n/********************************\n * Output\n */\n\necho('Hello World!');\n// Prints Hello World! to stdout.\n// Stdout is the web page if running in a browser.\n\nprint('Hello World!'); // The same as echo\n\n// echo and print are language constructs too, so you can drop the parentheses\necho 'Hello World!';\nprint 'Hello World!';\n\n$paragraph = 'paragraph';\n\necho 100;        // Echo scalar variables directly\necho $paragraph; // or variables\n\n// If short open tags are configured, or your PHP version is\n// 5.4.0 or greater, you can use the short echo syntax\n?>\n<p><?= $paragraph ?></p>\n<?php\n\n$x = 1;\n$y = 2;\n$x = $y; // $x now contains the same value as $y\n$z = &$y;\n// $z now contains a reference to $y. Changing the value of\n// $z will change the value of $y also, and vice-versa.\n// $x will remain unchanged as the original value of $y\n\necho $x; // => 2\necho $z; // => 2\n$y = 0;\necho $x; // => 2\necho $z; // => 0\n\n// Dumps type and value of variable to stdout\nvar_dump($z); // prints int(0)\n\n// Prints variable to stdout in human-readable format\nprint_r($array); // prints: Array ( [0] => One [1] => Two [2] => Three )\n\n/********************************\n * Logic\n */\n$a = 0;\n$b = '0';\n$c = '1';\n$d = '1';\n\n// assert throws a warning if its argument is not true\n\n// These comparisons will always be true, even if the types aren't the same.\nassert($a == $b); // equality\nassert($c != $a); // inequality\nassert($c <> $a); // alternative inequality\nassert($a < $c);\nassert($c > $b);\nassert($a <= $b);\nassert($c >= $d);\n\n// The following will only be true if the values match and are the same type.\nassert($c === $d);\nassert($a !== $d);\nassert(1 === '1');\nassert(1 !== '1');\n\n// 'Spaceship' operator (since PHP 7)\n// Returns 0 if values on either side are equal\n// Returns 1 if value on the left is greater\n// Returns -1 if the value on the right is greater\n\n$a = 100;\n$b = 1000;\n\necho $a <=> $a; // 0 since they are equal\necho $a <=> $b; // -1 since $a < $b\necho $b <=> $a; // 1 since $b > $a\n\n// Variables can be converted between types, depending on their usage.\n\n$integer = 1;\necho $integer + $integer; // => 2\n\n$string = '1';\necho $string + $string; // => 2 (strings are coerced to integers)\n\n$string = 'one';\necho $string + $string; // => 0\n// Outputs 0 because the + operator cannot cast the string 'one' to a number\n\n// Type casting can be used to treat a variable as another type\n\n$boolean = (boolean) 1; // => true\n\n$zero = 0;\n$boolean = (boolean) $zero; // => false\n\n// There are also dedicated functions for casting most types\n$integer = 5;\n$string = strval($integer);\n\n$var = null; // Null value\n\n\n/********************************\n * Control Structures\n */\n\nif (true) {\n    print 'I get printed';\n}\n\nif (false) {\n    print 'I don\\'t';\n} else {\n    print 'I get printed';\n}\n\nif (false) {\n    print 'Does not get printed';\n} elseif (true) {\n    print 'Does';\n}\n\n// ternary operator\nprint (false ? 'Does not get printed' : 'Does');\n\n// ternary shortcut operator since PHP 5.3\n// equivalent of \"$x ? $x : 'Does'\"\n$x = false;\nprint($x ?: 'Does');\n\n// null coalesce operator since php 7\n$a = null;\n$b = 'Does print';\necho $a ?? 'a is not set'; // prints 'a is not set'\necho $b ?? 'b is not set'; // prints 'Does print'\n\n\n$x = 0;\nif ($x === '0') {\n    print 'Does not print';\n} elseif ($x == '1') {\n    print 'Does not print';\n} else {\n    print 'Does print';\n}\n\n\n\n// This alternative syntax is useful for templates:\n?>\n\n<?php if ($x): ?>\nThis is displayed if the test is truthy.\n<?php else: ?>\nThis is displayed otherwise.\n<?php endif; ?>\n\n<?php\n\n// Use switch to save some logic.\nswitch ($x) {\n    case '0':\n        print 'Switch does type coercion';\n        break; // You must include a break, or you will fall through\n               // to cases 'two' and 'three'\n    case 'two':\n    case 'three':\n        // Do something if $variable is either 'two' or 'three'\n        break;\n    default:\n        // Do something by default\n}\n\n// While, do...while and for loops are probably familiar\n$i = 0;\nwhile ($i < 5) {\n    echo $i++;\n} // Prints \"01234\"\n\necho \"\\n\";\n\n$i = 0;\ndo {\n    echo $i++;\n} while ($i < 5); // Prints \"01234\"\n\necho \"\\n\";\n\nfor ($x = 0; $x < 10; $x++) {\n    echo $x;\n} // Prints \"0123456789\"\n\necho \"\\n\";\n\n$wheels = ['bicycle' => 2, 'car' => 4];\n\n// Foreach loops can iterate over arrays\nforeach ($wheels as $wheel_count) {\n    echo $wheel_count;\n} // Prints \"24\"\n\necho \"\\n\";\n\n// You can iterate over the keys as well as the values\nforeach ($wheels as $vehicle => $wheel_count) {\n    echo \"A $vehicle has $wheel_count wheels\";\n}\n\necho \"\\n\";\n\n$i = 0;\nwhile ($i < 5) {\n    if ($i === 3) {\n        break; // Exit out of the while loop\n    }\n    echo $i++;\n} // Prints \"012\"\n\nfor ($i = 0; $i < 5; $i++) {\n    if ($i === 3) {\n        continue; // Skip this iteration of the loop\n    }\n    echo $i;\n} // Prints \"0124\"\n\n\n/********************************\n * Functions\n */\n\n// Define a function with \"function\":\nfunction my_function () {\n    return 'Hello';\n}\n\necho my_function(); // => \"Hello\"\n\n// A valid function name starts with a letter or underscore, followed by any\n// number of letters, numbers, or underscores.\n\nfunction add ($x, $y = 1) { // $y is optional and defaults to 1\n    $result = $x + $y;\n    return $result;\n}\n\necho add(4); // => 5\necho add(4, 2); // => 6\n\n// $result is not accessible outside the function\n// print $result; // Gives a warning.\n\n// Since PHP 5.3 you can declare anonymous functions;\n$inc = function ($x) {\n    return $x + 1;\n};\n\necho $inc(2); // => 3\n\nfunction foo ($x, $y, $z) {\n    echo \"$x - $y - $z\";\n}\n\n// Functions can return functions\nfunction bar ($x, $y) {\n    // Use 'use' to bring in outside variables\n    return function ($z) use ($x, $y) {\n        foo($x, $y, $z);\n    };\n}\n\n$bar = bar('A', 'B');\n$bar('C'); // Prints \"A - B - C\"\n\n// You can call named functions using strings\n$function_name = 'add';\necho $function_name(1, 2); // => 3\n// Useful for programatically determining which function to run.\n// Or, use call_user_func(callable $callback [, $parameter [, ... ]]);\n\n\n// You can get all the parameters passed to a function\nfunction parameters() {\n    $numargs = func_num_args();\n    if ($numargs > 0) {\n        echo func_get_arg(0) . ' | ';\n    }\n    $args_array = func_get_args();\n    foreach ($args_array as $key => $arg) {\n        echo $key . ' - ' . $arg . ' | ';\n    }\n}\n\nparameters('Hello', 'World'); // Hello | 0 - Hello | 1 - World |\n\n// Since PHP 5.6 you can get a variable number of arguments\nfunction variable($word, ...$list) {\n\techo $word . \" || \";\n\tforeach ($list as $item) {\n\t\techo $item . ' | ';\n\t}\n}\n\nvariable(\"Separate\", \"Hello\", \"World\"); // Separate || Hello | World |\n\n/********************************\n * Includes\n */\n\n<?php\n// PHP within included files must also begin with a PHP open tag.\n\ninclude 'my-file.php';\n// The code in my-file.php is now available in the current scope.\n// If the file cannot be included (e.g. file not found), a warning is emitted.\n\ninclude_once 'my-file.php';\n// If the code in my-file.php has been included elsewhere, it will\n// not be included again. This prevents multiple class declaration errors\n\nrequire 'my-file.php';\nrequire_once 'my-file.php';\n// Same as include(), except require() will cause a fatal error if the\n// file cannot be included.\n\n// Contents of my-include.php:\n<?php\n\nreturn 'Anything you like.';\n// End file\n\n// Includes and requires may also return a value.\n$value = include 'my-include.php';\n\n// Files are included based on the file path given or, if none is given,\n// the include_path configuration directive. If the file isn't found in\n// the include_path, include will finally check in the calling script's\n// own directory and the current working directory before failing.\n/* */\n\n/********************************\n * Classes\n */\n\n// Classes are defined with the class keyword\n\nclass MyClass\n{\n    const MY_CONST      = 'value'; // A constant\n\n    static $staticVar   = 'static';\n\n    // Static variables and their visibility\n    public static $publicStaticVar = 'publicStatic';\n    // Accessible within the class only\n    private static $privateStaticVar = 'privateStatic';\n    // Accessible from the class and subclasses\n    protected static $protectedStaticVar = 'protectedStatic';\n\n    // Properties must declare their visibility\n    public $property    = 'public';\n    public $instanceProp;\n    protected $prot = 'protected'; // Accessible from the class and subclasses\n    private $priv   = 'private';   // Accessible within the class only\n\n    // Create a constructor with __construct\n    public function __construct($instanceProp)\n    {\n        // Access instance variables with $this\n        $this->instanceProp = $instanceProp;\n    }\n\n    // Methods are declared as functions inside a class\n    public function myMethod()\n    {\n        print 'MyClass';\n    }\n\n    // final keyword would make a function unoverridable\n    final function youCannotOverrideMe()\n    {\n    }\n\n    // Magic Methods\n\n    // what to do if Object is treated as a String\n    public function __toString()\n    {\n        return $property;\n    }\n\n    // opposite to __construct()\n    // called when object is no longer referenced\n    public function __destruct()\n    {\n        print \"Destroying\";\n    }\n\n/*\n * Declaring class properties or methods as static makes them accessible without\n * needing an instantiation of the class. A property declared as static can not\n * be accessed with an instantiated class object (though a static method can).\n */\n\n    public static function myStaticMethod()\n    {\n        print 'I am static';\n    }\n}\n\n// Class constants can always be accessed statically\necho MyClass::MY_CONST;    // Outputs 'value';\n\necho MyClass::$staticVar;  // Outputs 'static';\nMyClass::myStaticMethod(); // Outputs 'I am static';\n\n// Instantiate classes using new\n$my_class = new MyClass('An instance property');\n// The parentheses are optional if not passing in an argument.\n\n// Access class members using ->\necho $my_class->property;     // => \"public\"\necho $my_class->instanceProp; // => \"An instance property\"\n$my_class->myMethod();        // => \"MyClass\"\n\n\n// Extend classes using \"extends\"\nclass MyOtherClass extends MyClass\n{\n    function printProtectedProperty()\n    {\n        echo $this->prot;\n    }\n\n    // Override a method\n    function myMethod()\n    {\n        parent::myMethod();\n        print ' > MyOtherClass';\n    }\n}\n\n$my_other_class = new MyOtherClass('Instance prop');\n$my_other_class->printProtectedProperty(); // => Prints \"protected\"\n$my_other_class->myMethod();               // Prints \"MyClass > MyOtherClass\"\n\nfinal class YouCannotExtendMe\n{\n}\n\n// You can use \"magic methods\" to create getters and setters\nclass MyMapClass\n{\n    private $property;\n\n    public function __get($key)\n    {\n        return $this->$key;\n    }\n\n    public function __set($key, $value)\n    {\n        $this->$key = $value;\n    }\n}\n\n$x = new MyMapClass();\necho $x->property; // Will use the __get() method\n$x->property = 'Something'; // Will use the __set() method\n\n// Classes can be abstract (using the abstract keyword) or\n// implement interfaces (using the implements keyword).\n// An interface is declared with the interface keyword.\n\ninterface InterfaceOne\n{\n    public function doSomething();\n}\n\ninterface InterfaceTwo\n{\n    public function doSomethingElse();\n}\n\n// interfaces can be extended\ninterface InterfaceThree extends InterfaceTwo\n{\n    public function doAnotherContract();\n}\n\nabstract class MyAbstractClass implements InterfaceOne\n{\n    public $x = 'doSomething';\n}\n\nclass MyConcreteClass extends MyAbstractClass implements InterfaceTwo\n{\n    public function doSomething()\n    {\n        echo $x;\n    }\n\n    public function doSomethingElse()\n    {\n        echo 'doSomethingElse';\n    }\n}\n\n\n// Classes can implement more than one interface\nclass SomeOtherClass implements InterfaceOne, InterfaceTwo\n{\n    public function doSomething()\n    {\n        echo 'doSomething';\n    }\n\n    public function doSomethingElse()\n    {\n        echo 'doSomethingElse';\n    }\n}\n\n\n/********************************\n * Traits\n */\n\n// Traits are available from PHP 5.4.0 and are declared using \"trait\"\n\ntrait MyTrait\n{\n    public function myTraitMethod()\n    {\n        print 'I have MyTrait';\n    }\n}\n\nclass MyTraitfulClass\n{\n    use MyTrait;\n}\n\n$cls = new MyTraitfulClass();\n$cls->myTraitMethod(); // Prints \"I have MyTrait\"\n\n\n/********************************\n * Namespaces\n */\n\n// This section is separate, because a namespace declaration\n// must be the first statement in a file. Let's pretend that is not the case\n\n<?php\n\n// By default, classes exist in the global namespace, and can\n// be explicitly called with a backslash.\n\n$cls = new \\MyClass();\n\n\n\n// Set the namespace for a file\nnamespace My\\Namespace;\n\nclass MyClass\n{\n}\n\n// (from another file)\n$cls = new My\\Namespace\\MyClass;\n\n//Or from within another namespace.\nnamespace My\\Other\\Namespace;\n\nuse My\\Namespace\\MyClass;\n\n$cls = new MyClass();\n\n// Or you can alias the namespace;\n\nnamespace My\\Other\\Namespace;\n\nuse My\\Namespace as SomeOtherNamespace;\n\n$cls = new SomeOtherNamespace\\MyClass();\n\n\n/**********************\n* Late Static Binding\n*\n*/\n\nclass ParentClass\n{\n    public static function who()\n    {\n        echo \"I'm a \" . __CLASS__ . \"\\n\";\n    }\n\n    public static function test()\n    {\n        // self references the class the method is defined within\n        self::who();\n        // static references the class the method was invoked on\n        static::who();\n    }\n}\n\nParentClass::test();\n/*\nI'm a ParentClass\nI'm a ParentClass\n*/\n\nclass ChildClass extends ParentClass\n{\n    public static function who()\n    {\n        echo \"But I'm \" . __CLASS__ . \"\\n\";\n    }\n}\n\nChildClass::test();\n/*\nI'm a ParentClass\nBut I'm ChildClass\n*/\n\n/**********************\n*  Magic constants\n*\n*/\n\n// Get current class name. Must be used inside a class declaration.\necho \"Current class name is \" . __CLASS__;\n\n// Get full path directory of a file\necho \"Current directory is \" . __DIR__;\n\n    // Typical usage\n    require __DIR__ . '/vendor/autoload.php';\n\n// Get full path of a file\necho \"Current file path is \" . __FILE__;\n\n// Get current function name\necho \"Current function name is \" . __FUNCTION__;\n\n// Get current line number\necho \"Current line number is \" . __LINE__;\n\n// Get the name of the current method. Only returns a value when used inside a trait or object declaration.\necho \"Current method is \" . __METHOD__;\n\n// Get the name of the current namespace\necho \"Current namespace is \" . __NAMESPACE__;\n\n// Get the name of the current trait. Only returns a value when used inside a trait or object declaration.\necho \"Current trait is \" . __TRAIT__;\n\n/**********************\n*  Error Handling\n*\n*/\n\n// Simple error handling can be done with try catch block\n\ntry {\n    // Do something\n} catch (Exception $e) {\n    // Handle exception\n}\n\n// When using try catch blocks in a namespaced environment it is important to\n// escape to the global namespace, because Exceptions are classes, and the\n// Exception class exists in the global namespace. This can be done using a\n// leading backslash to catch the Exception.\n\ntry {\n    // Do something\n} catch (\\Exception $e) {\n    // Handle exception\n}\n\n// Custom exceptions\n\nclass MyException extends Exception {}\n\ntry {\n\n    $condition = true;\n\n    if ($condition) {\n        throw new MyException('Something just happened');\n    }\n\n} catch (MyException $e) {\n    // Handle my exception\n}\n\n```\n\n## More Information\n\nVisit the [official PHP documentation](http://www.php.net/manual/) for reference\nand community input.\n\nIf you're interested in up-to-date best practices, visit\n[PHP The Right Way](http://www.phptherightway.com/).\n\nA tutorial covering basics of language, setting up coding environment and making\nfew practical projects at [Codecourse - PHP Basics](https://www.youtube.com/playlist?list=PLfdtiltiRHWHjTPiFDRdTOPtSyYfz3iLW).\n\nIf you're coming from a language with good package management, check out\n[Composer](http://getcomposer.org/).\n\nFor common standards, visit the PHP Framework Interoperability Group's\n[PSR standards](https://github.com/php-fig/fig-standards)."
