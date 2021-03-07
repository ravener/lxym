
package pages

const Angularjs = "## AngularJS Tutorial.\n\nAngularJS version 1.0 was released in 2012.\nMiško Hevery, a Google employee, started to work with AngularJS in 2009.\nThe idea turned out very well, and the project is now officially supported by Google.\n\nAngularJS is a JavaScript framework. It can be added to an HTML page with a \"script\" tag.\nAngularJS extends HTML attributes with Directives, and binds data to HTML with Expressions.\n\n##What You Should Already Know\n\nBefore you study AngularJS, you should have a basic understanding of:\n\n- HTML\n- CSS\n- JavaScript\n\n```html\n// AngularJS is a JavaScript framework. It is a library written in JavaScript.\n// AngularJS is distributed as a JavaScript file, and can be added to a web page with a script tag:\n// <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n\n///////////////////////////////////\n// AngularJS Extends HTML\n\n//AngularJS extends HTML with ng-directives.\n//The ng-app directive defines an AngularJS application.\n//The ng-model directive binds the value of HTML controls (input, select, textarea) to application data.\n//The ng-bind directive binds application data to the HTML view.\n<!DOCTYPE html>\n<html>\n  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n  <body>\n    <div ng-app=\"\">\n      <p>Name: <input type=\"text\" ng-model=\"name\"></p>\n      <p ng-bind=\"name\"></p>\n    </div>\n  </body>\n</html>\n\n/*\n  * Example explained:\n  * AngularJS starts automatically when the web page has loaded.\n  * The ng-app directive tells AngularJS that the <div> element is the \"owner\" of an AngularJS application.\n  * The ng-model directive binds the value of the input field to the application variable name.\n  * The ng-bind directive binds the innerHTML of the <p> element to the application variable name.\n*/\n<tag> Here are content to be interpreted </tag>\n\n///////////////////////////////////\n// AngularJS Expressions\n\n// AngularJS expressions are written inside double braces: {{ expression }}.\n// AngularJS expressions binds data to HTML the same way as the ng-bind directive.\n// AngularJS will \"output\" data exactly where the expression is written.\n// AngularJS expressions are much like JavaScript expressions: They can contain literals, operators, and variables.\n// Example {{ 5 + 5 }} or {{ firstName + \" \" + lastName }}\n<!DOCTYPE html>\n<html>\n  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n  <body>\n    <div ng-app=\"\">\n      <p>My first expression: {{ 5 + 5 }}</p>\n    </div>\n  </body>\n</html>\n\n//If you remove the ng-app directive, HTML will display the expression as it is, without solving it:\n<!DOCTYPE html>\n<html>\n  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n  <body>\n    <div>\n      <p>My first expression: {{ 5 + 5 }}</p>\n    </div>\n  </body>\n</html>\n\n// AngularJS expressions bind AngularJS data to HTML the same way as the ng-bind directive.\n<!DOCTYPE html>\n<html>\n<script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n  <body>\n    <div ng-app=\"\">\n      <p>Name: <input type=\"text\" ng-model=\"name\"></p>\n      <p>{{name}}</p>\n    </div>\n  </body>\n</html>\n\n// AngularJS numbers are like JavaScript numbers:\n<div ng-app=\"\" ng-init=\"quantity=1;cost=5\">\n  <p>Total in dollar: {{ quantity * cost }}</p>\n</div>\n\n//AngularJS strings are like JavaScript strings:\n<div ng-app=\"\" ng-init=\"firstName='John';lastName='Doe'\">\n  <p>The name is <span ng-bind=\"firstName + ' ' + lastName\"></span></p>\n</div>\n\n//AngularJS objects are like JavaScript objects:\n<div ng-app=\"\" ng-init=\"person={firstName:'John',lastName:'Doe'}\">\n  <p>The name is {{ person.lastName }}</p>\n</div>\n\n//AngularJS arrays are like JavaScript arrays:\n<div ng-app=\"\" ng-init=\"points=[1,15,19,2,40]\">\n  <p>The third result is {{ points[2] }}</p>\n</div>\n\n// Like JavaScript expressions, AngularJS expressions can contain literals, operators, and variables.\n// Unlike JavaScript expressions, AngularJS expressions can be written inside HTML.\n// AngularJS expressions do not support conditionals, loops, and exceptions, while JavaScript expressions do.\n// AngularJS expressions support filters, while JavaScript expressions do not.\n\n///////////////////////////////////\n// AngularJS Directives\n\n\n//AngularJS directives are extended HTML attributes with the prefix ng-.\n//The ng-app directive initializes an AngularJS application.\n//The ng-init directive initializes application data.\n//The ng-model directive binds the value of HTML controls (input, select, textarea) to application data.\n<div ng-app=\"\" ng-init=\"firstName='John'\">\n  <p>Name: <input type=\"text\" ng-model=\"firstName\"></p>\n  <p>You wrote: {{ firstName }}</p>\n</div>\n\n//Using ng-init is not very common. You will learn how to initialize data in the chapter about controllers.\n\n//The ng-repeat directive repeats an HTML element:\n<div ng-app=\"\" ng-init=\"names=['Jani','Hege','Kai']\">\n  <ul>\n    <li ng-repeat=\"x in names\">\n      {{ x }}\n    </li>\n  </ul>\n</div>\n\n//The ng-repeat directive used on an array of objects:\n<div ng-app=\"\" ng-init=\"names=[\n{name:'Jani',country:'Norway'},\n{name:'Hege',country:'Sweden'},\n{name:'Kai',country:'Denmark'}]\">\n  <ul>\n    <li ng-repeat=\"x  in names\">\n      {{ x.name + ', ' + x.country }}\n    </li>\n  </ul>\n</div>\n\n// AngularJS is perfect for database CRUD (Create Read Update Delete) applications.\n// Just imagine if these objects were records from a database.\n\n// The ng-app directive defines the root element of an AngularJS application.\n// The ng-app directive will auto-bootstrap (automatically initialize) the application when a web page is loaded.\n// Later you will learn how ng-app can have a value (like ng-app=\"myModule\"), to connect code modules.\n\n// The ng-init directive defines initial values for an AngularJS application.\n// Normally, you will not use ng-init. You will use a controller or module instead.\n// You will learn more about controllers and modules later.\n\n//The ng-model directive binds the value of HTML controls (input, select, textarea) to application data.\n//The ng-model directive can also:\n//Provide type validation for application data (number, email, required).\n//Provide status for application data (invalid, dirty, touched, error).\n//Provide CSS classes for HTML elements.\n//Bind HTML elements to HTML forms.\n\n//The ng-repeat directive clones HTML elements once for each item in a collection (in an array).\n\n///////////////////////////////////\n// AngularJS Controllers\n\n// AngularJS controllers control the data of AngularJS applications.\n// AngularJS controllers are regular JavaScript Objects.\n\n// AngularJS applications are controlled by controllers.\n// The ng-controller directive defines the application controller.\n// A controller is a JavaScript Object, created by a standard JavaScript object constructor.\n\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n\nFirst Name: <input type=\"text\" ng-model=\"firstName\"><br>\nLast Name: <input type=\"text\" ng-model=\"lastName\"><br>\n<br>\nFull Name: {{firstName + \" \" + lastName}}\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('myCtrl', function($scope) {\n    $scope.firstName = \"John\";\n    $scope.lastName = \"Doe\";\n});\n</script>\n\n//Application explained:\n\n//The AngularJS application is defined by  ng-app=\"myApp\". The application runs inside the <div>.\n//The ng-controller=\"myCtrl\" attribute is an AngularJS directive. It defines a controller.\n//The myCtrl function is a JavaScript function.\n//AngularJS will invoke the controller with a $scope object.\n//In AngularJS, $scope is the application object (the owner of application variables and functions).\n//The controller creates two properties (variables) in the scope (firstName and lastName).\n//The ng-model directives bind the input fields to the controller properties (firstName and lastName).\n\n//The example above demonstrated a controller object with two properties: lastName and firstName.\n//A controller can also have methods (variables as functions):\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\nFirst Name: <input type=\"text\" ng-model=\"firstName\"><br>\nLast Name: <input type=\"text\" ng-model=\"lastName\"><br>\n<br>\nFull Name: {{fullName()}}\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('personCtrl', function($scope) {\n    $scope.firstName = \"John\";\n    $scope.lastName = \"Doe\";\n    $scope.fullName = function() {\n        return $scope.firstName + \" \" + $scope.lastName;\n    }\n});\n</script>\n\n//In larger applications, it is common to store controllers in external files.\n//Just copy the code between the <script> </script> tags into an external file named personController.js:\n\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\nFirst Name: <input type=\"text\" ng-model=\"firstName\"><br>\nLast Name: <input type=\"text\" ng-model=\"lastName\"><br>\n<br>\nFull Name: {{firstName + \" \" + lastName}}\n\n</div>\n\n<script src=\"personController.js\"></script>\n\n// For the next example we will create a new controller file:\nangular.module('myApp', []).controller('namesCtrl', function($scope) {\n    $scope.names = [\n        {name:'Jani',country:'Norway'},\n        {name:'Hege',country:'Sweden'},\n        {name:'Kai',country:'Denmark'}\n    ];\n});\n\n//Save the file as  namesController.js:\n//And then use the controller file in an application:\n\n<div ng-app=\"myApp\" ng-controller=\"namesCtrl\">\n\n<ul>\n  <li ng-repeat=\"x in names\">\n    {{ x.name + ', ' + x.country }}\n  </li>\n</ul>\n\n</div>\n\n<script src=\"namesController.js\"></script>\n\n///////////////////////////////////\n// AngularJS Filters\n\n// Filters can be added to expressions and directives using a pipe character.\n// AngularJS filters can be used to transform data:\n\n- **currency**:  Format a number to a currency format.\n- **filter**:  Select a subset of items from an array.\n- **lowercase**: Format a string to lower case.\n- **orderBy**: Orders an array by an expression.\n- **uppercase**: Format a string to upper case.\n\n//A filter can be added to an expression with a pipe character (|) and a filter.\n//(For the next two examples we will use the person controller from the previous chapter)\n//The uppercase filter format strings to upper case:\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\n<p>The name is {{ lastName | uppercase }}</p>\n\n</div>\n\n//The lowercase filter format strings to lower case:\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\n<p>The name is {{ lastName | lowercase }}</p>\n\n</div>\n\n//The currency filter formats a number as currency:\n<div ng-app=\"myApp\" ng-controller=\"costCtrl\">\n\n<input type=\"number\" ng-model=\"quantity\">\n<input type=\"number\" ng-model=\"price\">\n\n<p>Total = {{ (quantity * price) | currency }}</p>\n\n</div> \n\n//A filter can be added to a directive with a pipe character (|) and a filter.\n//The orderBy filter orders an array by an expression:\n<div ng-app=\"myApp\" ng-controller=\"namesCtrl\">\n\n<ul>\n  <li ng-repeat=\"x in names | orderBy:'country'\">\n    {{ x.name + ', ' + x.country }}\n  </li>\n</ul>\n\n<div>\n\n//An input filter can be added to a directive with a pipe character (|) \n//and filter followed by a colon and a model name.\n//The filter selects a subset of an array:\n\n<div ng-app=\"myApp\" ng-controller=\"namesCtrl\">\n\n<p><input type=\"text\" ng-model=\"test\"></p>\n\n<ul>\n  <li ng-repeat=\"x in names | filter:test | orderBy:'country'\">\n    {{ (x.name | uppercase) + ', ' + x.country }}\n  </li>\n</ul>\n\n</div>\n\n///////////////////////////////////\n// AngularJS AJAX - $http\n\n//$http is an AngularJS service for reading data from remote servers.\n\n// The following data can be provided by a web server:\n// http://www.w3schools.com/angular/customers.php\n// **Check the URL to see the data format**\n\n// AngularJS $http is a core service for reading data from web servers.\n// $http.get(url) is the function to use for reading server data.\n<div ng-app=\"myApp\" ng-controller=\"customersCtrl\"> \n\n<ul>\n  <li ng-repeat=\"x in names\">\n    {{ x.Name + ', ' + x.Country }}\n  </li>\n</ul>\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('customersCtrl', function($scope, $http) {\n    $http.get(\"http://www.w3schools.com/angular/customers.php\")\n    .success(function(response) {$scope.names = response.records;});\n});\n</script>\n\nApplication explained:\n\n// The AngularJS application is defined by ng-app. The application runs inside a <div>.\n// The ng-controller directive names the controller object.\n// The customersCtrl function is a standard JavaScript object constructor.\n// AngularJS will invoke customersCtrl with a $scope and $http object.\n// $scope is the application object (the owner of application variables and functions).\n// $http is an XMLHttpRequest object for requesting external data.\n// $http.get() reads JSON data from http://www.w3schools.com/angular/customers.php.\n// If success, the controller creates a property (names) in the scope, with JSON data from the server.\n\n\n// Requests for data from a different server (than the requesting page), are called cross-site HTTP requests.\n// Cross-site requests are common on the web. Many pages load CSS, images, and scripts from different servers.\n// In modern browsers, cross-site HTTP requests from scripts are restricted to same site for security reasons.\n// The following line, in our PHP examples, has been added to allow cross-site access.\nheader(\"Access-Control-Allow-Origin: *\");\n\n\n///////////////////////////////////\n// AngularJS Tables\n\n// Displaying tables with angular is very simple:\n<div ng-app=\"myApp\" ng-controller=\"customersCtrl\"> \n\n<table>\n  <tr ng-repeat=\"x in names\">\n    <td>{{ x.Name }}</td>\n    <td>{{ x.Country }}</td>\n  </tr>\n</table>\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('customersCtrl', function($scope, $http) {\n    $http.get(\"http://www.w3schools.com/angular/customers.php\")\n    .success(function (response) {$scope.names = response.records;});\n});\n</script>\n\n// To sort the table, add an orderBy filter: \n<table>\n  <tr ng-repeat=\"x in names | orderBy : 'Country'\">\n    <td>{{ x.Name }}</td>\n    <td>{{ x.Country }}</td>\n  </tr>\n</table>\n\n// To display the table index, add a <td> with $index: \n<table>\n  <tr ng-repeat=\"x in names\">\n    <td>{{ $index + 1 }}</td>\n    <td>{{ x.Name }}</td>\n    <td>{{ x.Country }}</td>\n  </tr>\n</table>\n\n// Using $even and $odd\n<table>\n  <tr ng-repeat=\"x in names\">\n    <td ng-if=\"$odd\" style=\"background-color:#f1f1f1\">{{ x.Name }}</td>\n    <td ng-if=\"$even\">{{ x.Name }}</td>\n    <td ng-if=\"$odd\" style=\"background-color:#f1f1f1\">{{ x.Country }}</td>\n    <td ng-if=\"$even\">{{ x.Country }}</td>\n  </tr>\n</table>\n\n///////////////////////////////////\n// AngularJS HTML DOM\n\n//AngularJS has directives for binding application data to the attributes of HTML DOM elements.\n\n// The ng-disabled directive binds AngularJS application data to the disabled attribute of HTML elements.\n\n<div ng-app=\"\" ng-init=\"mySwitch=true\">\n\n<p>\n<button ng-disabled=\"mySwitch\">Click Me!</button>\n</p>\n\n<p>\n<input type=\"checkbox\" ng-model=\"mySwitch\">Button\n</p>\n\n</div>\n\n//Application explained:\n\n// The ng-disabled directive binds the application data mySwitch to the HTML button's disabled attribute.\n// The ng-model directive binds the value of the HTML checkbox element to the value of mySwitch.\n// If the value of mySwitch evaluates to true, the button will be disabled: \n<p>\n<button disabled>Click Me!</button>\n</p>\n\n// If the value of mySwitch evaluates to false, the button will not be disabled: \n<p>\n  <button>Click Me!</button>\n</p>\n\n// The ng-show directive shows or hides an HTML element.\n\n<div ng-app=\"\">\n\n<p ng-show=\"true\">I am visible.</p>\n\n<p ng-show=\"false\">I am not visible.</p>\n\n</div>\n\n// The ng-show directive shows (or hides) an HTML element based on the value of ng-show.\n// You can use any expression that evaluates to true or false:\n<div ng-app=\"\">\n<p ng-show=\"hour > 12\">I am visible.</p>\n</div>\n\n///////////////////////////////////\n// AngularJS Events\n\n// AngularJS has its own HTML events directives.\n\n// The ng-click directive defines an AngularJS click event.\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n\n<button ng-click=\"count = count + 1\">Click me!</button>\n\n<p>{{ count }}</p>\n\n</div>\n<script>\nvar app = angular.module('myApp', []);\napp.controller('myCtrl', function($scope) {\n    $scope.count = 0;\n});\n</script>\n\n// The ng-hide directive can be used to set the visibility of a part of an application.\n// The value ng-hide=\"true\" makes an HTML element invisible.\n// The value ng-hide=\"false\" makes the element visible.\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\n<button ng-click=\"toggle()\">Toggle</button>\n\n<p ng-hide=\"myVar\">\nFirst Name: <input type=\"text\" ng-model=\"firstName\"><br>\nLast Name: <input type=\"text\" ng-model=\"lastName\"><br>\n<br>\nFull Name: {{firstName + \" \" + lastName}}\n</p>\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('personCtrl', function($scope) {\n    $scope.firstName = \"John\",\n    $scope.lastName = \"Doe\"\n    $scope.myVar = false;\n    $scope.toggle = function() {\n        $scope.myVar = !$scope.myVar;\n    };\n});\n</script>\n\n//Application explained:\n\n// The first part of the personController is the same as in the chapter about controllers.\n// The application has a default property (a variable): $scope.myVar = false;\n// The ng-hide directive sets the visibility, of a <p> element with two input fields, \n// according to the value (true or false) of myVar.\n// The function toggle() toggles myVar between true and false.\n// The value ng-hide=\"true\" makes the element invisible.\n\n\n// The ng-show directive can also be used to set the visibility of a part of an application.\n// The value ng-show=\"false\" makes an HTML element invisible.\n// The value ng-show=\"true\" makes the element visible.\n// Here is the same example as above, using ng-show instead of ng-hide:\n<div ng-app=\"myApp\" ng-controller=\"personCtrl\">\n\n<button ng-click=\"toggle()\">Toggle</button>\n\n<p ng-show=\"myVar\">\nFirst Name: <input type=\"text\" ng-model=\"firstName\"><br>\nLast Name: <input type=\"text\" ng-model=\"lastName\"><br>\n<br>\nFull Name: {{firstName + \" \" + lastName}}\n</p>\n\n</div>\n\n<script>\nvar app = angular.module('myApp', []);\napp.controller('personCtrl', function($scope) {\n    $scope.firstName = \"John\",\n    $scope.lastName = \"Doe\"\n    $scope.myVar = true;\n    $scope.toggle = function() {\n        $scope.myVar = !$scope.myVar;\n    }\n});\n</script>\n\n///////////////////////////////////\n// AngularJS Modules\n\n// An AngularJS module defines an application.\n// The module is a container for the different parts of an application.\n// The module is a container for the application controllers.\n// Controllers always belong to a module.\n\n// This application (\"myApp\") has one controller (\"myCtrl\"):\n\n<!DOCTYPE html>\n<html>\n<script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n<body>\n\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n{{ firstName + \" \" + lastName }}\n</div>\n\n<script>\nvar app = angular.module(\"myApp\", []);\napp.controller(\"myCtrl\", function($scope) {\n    $scope.firstName = \"John\";\n    $scope.lastName = \"Doe\";\n});\n</script>\n\n</body>\n</html>\n\n// It is common in AngularJS applications to put the module and the controllers in JavaScript files.\n// In this example, \"myApp.js\" contains an application module definition, while \"myCtrl.js\" contains the controller:\n\n<!DOCTYPE html>\n<html>\n<script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n<body>\n\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n{{ firstName + \" \" + lastName }}\n</div>\n\n<script src=\"myApp.js\"></script>\n<script src=\"myCtrl.js\"></script>\n\n</body>\n</html>\n\n//myApp.js\nvar app = angular.module(\"myApp\", []); \n\n// The [] parameter in the module definition can be used to define dependent modules.\n\n// myCtrl.js\napp.controller(\"myCtrl\", function($scope) {\n    $scope.firstName  = \"John\";\n    $scope.lastName= \"Doe\";\n});\n\n// Global functions should be avoided in JavaScript. They can easily be overwritten \n// or destroyed by other scripts.\n\n// AngularJS modules reduces this problem, by keeping all functions local to the module.\n\n// While it is common in HTML applications to place scripts at the end of the \n// <body> element, it is recommended that you load the AngularJS library either\n// in the <head> or at the start of the <body>.\n\n// This is because calls to angular.module can only be compiled after the library has been loaded.\n\n<!DOCTYPE html>\n<html>\n<body>\n<script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js\"></script>\n\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n{{ firstName + \" \" + lastName }}\n</div>\n\n<script>\nvar app = angular.module(\"myApp\", []);\napp.controller(\"myCtrl\", function($scope) {\n    $scope.firstName = \"John\";\n    $scope.lastName = \"Doe\";\n});\n</script>\n\n</body>\n</html>\n\n\n///////////////////////////////////\n// AngularJS Applications\n\n// AngularJS modules define AngularJS applications.\n// AngularJS controllers control AngularJS applications.\n// The ng-app directive defines the application, the ng-controller directive defines the controller.\n<div ng-app=\"myApp\" ng-controller=\"myCtrl\">\n  First Name: <input type=\"text\" ng-model=\"firstName\"><br>\n  Last Name: <input type=\"text\" ng-model=\"lastName\"><br>\n  <br>\n  Full Name: {{firstName + \" \" + lastName}}\n</div>\n<script>\n  var app = angular.module('myApp', []);\n  app.controller('myCtrl', function($scope) {\n      $scope.firstName= \"John\";\n      $scope.lastName= \"Doe\";\n  });\n</script>\n\n// AngularJS modules define applications:\nvar app = angular.module('myApp', []);\n\n// AngularJS controllers control applications:\napp.controller('myCtrl', function($scope) {\n    $scope.firstName= \"John\";\n    $scope.lastName= \"Doe\";\n});\n```\n\n## Source & References\n\n**Examples**\n\n- [http://www.w3schools.com/angular/angular_examples.asp](http://www.w3schools.com/angular/angular_examples.asp)\n\n**References**\n\n- [http://www.w3schools.com/angular/angular_ref_directives.asp](http://www.w3schools.com/angular/angular_ref_directives.asp)\n- [http://www.w3schools.com/angular/default.asp](http://www.w3schools.com/angular/default.asp)\n- [https://teamtreehouse.com/library/angular-basics/](https://teamtreehouse.com/library/angular-basics/)"
