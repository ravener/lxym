
package pages

const Typescript = "TypeScript is a language that aims at easing development of large scale\napplications written in JavaScript.  TypeScript adds common concepts such as\nclasses, modules, interfaces, generics and (optional) static typing to\nJavaScript.  It is a superset of JavaScript: all JavaScript code is valid\nTypeScript code so it can be added seamlessly to any project. The TypeScript\ncompiler emits JavaScript.\n\nThis article will focus only on TypeScript extra syntax, as opposed to\n[JavaScript](/docs/javascript).\n\nTo test TypeScript's compiler, head to the\n[Playground](https://www.typescriptlang.org/play) where you will be able\nto type code, have auto completion and directly see the emitted JavaScript.\n\n```ts\n// There are 3 basic types in TypeScript\nlet isDone: boolean = false;\nlet lines: number = 42;\nlet name: string = \"Anders\";\n\n// But you can omit the type annotation if the variables are derived\n// from explicit literals\nlet isDone = false;\nlet lines = 42;\nlet name = \"Anders\";\n\n// When it's impossible to know, there is the \"Any\" type\nlet notSure: any = 4;\nnotSure = \"maybe a string instead\";\nnotSure = false; // okay, definitely a boolean\n\n// Use const keyword for constants\nconst numLivesForCat = 9;\nnumLivesForCat = 1; // Error\n\n// For collections, there are typed arrays and generic arrays\nlet list: number[] = [1, 2, 3];\n// Alternatively, using the generic array type\nlet list: Array<number> = [1, 2, 3];\n\n// For enumerations:\nenum Color { Red, Green, Blue };\nlet c: Color = Color.Green;\n\n// Lastly, \"void\" is used in the special case of a function returning nothing\nfunction bigHorribleAlert(): void {\n  alert(\"I'm a little annoying box!\");\n}\n\n// Functions are first class citizens, support the lambda \"fat arrow\" syntax and\n// use type inference\n\n// The following are equivalent, the same signature will be inferred by the\n// compiler, and same JavaScript will be emitted\nlet f1 = function (i: number): number { return i * i; }\n// Return type inferred\nlet f2 = function (i: number) { return i * i; }\n// \"Fat arrow\" syntax\nlet f3 = (i: number): number => { return i * i; }\n// \"Fat arrow\" syntax with return type inferred\nlet f4 = (i: number) => { return i * i; }\n// \"Fat arrow\" syntax with return type inferred, braceless means no return\n// keyword needed\nlet f5 = (i: number) => i * i;\n\n// Interfaces are structural, anything that has the properties is compliant with\n// the interface\ninterface Person {\n  name: string;\n  // Optional properties, marked with a \"?\"\n  age?: number;\n  // And of course functions\n  move(): void;\n}\n\n// Object that implements the \"Person\" interface\n// Can be treated as a Person since it has the name and move properties\nlet p: Person = { name: \"Bobby\", move: () => { } };\n// Objects that have the optional property:\nlet validPerson: Person = { name: \"Bobby\", age: 42, move: () => { } };\n// Is not a person because age is not a number\nlet invalidPerson: Person = { name: \"Bobby\", age: true };\n\n// Interfaces can also describe a function type\ninterface SearchFunc {\n  (source: string, subString: string): boolean;\n}\n// Only the parameters' types are important, names are not important.\nlet mySearch: SearchFunc;\nmySearch = function (src: string, sub: string) {\n  return src.search(sub) != -1;\n}\n\n// Classes - members are public by default\nclass Point {\n  // Properties\n  x: number;\n\n  // Constructor - the public/private keywords in this context will generate\n  // the boiler plate code for the property and the initialization in the\n  // constructor.\n  // In this example, \"y\" will be defined just like \"x\" is, but with less code\n  // Default values are also supported\n\n  constructor(x: number, public y: number = 0) {\n    this.x = x;\n  }\n\n  // Functions\n  dist(): number { return Math.sqrt(this.x * this.x + this.y * this.y); }\n\n  // Static members\n  static origin = new Point(0, 0);\n}\n\n// Classes can be explicitly marked as implementing an interface.\n// Any missing properties will then cause an error at compile-time.\nclass PointPerson implements Person {\n    name: string\n    move() {}\n}\n\nlet p1 = new Point(10, 20);\nlet p2 = new Point(25); //y will be 0\n\n// Inheritance\nclass Point3D extends Point {\n  constructor(x: number, y: number, public z: number = 0) {\n    super(x, y); // Explicit call to the super class constructor is mandatory\n  }\n\n  // Overwrite\n  dist(): number {\n    let d = super.dist();\n    return Math.sqrt(d * d + this.z * this.z);\n  }\n}\n\n// Modules, \".\" can be used as separator for sub modules\nmodule Geometry {\n  export class Square {\n    constructor(public sideLength: number = 0) {\n    }\n    area() {\n      return Math.pow(this.sideLength, 2);\n    }\n  }\n}\n\nlet s1 = new Geometry.Square(5);\n\n// Local alias for referencing a module\nimport G = Geometry;\n\nlet s2 = new G.Square(10);\n\n// Generics\n// Classes\nclass Tuple<T1, T2> {\n  constructor(public item1: T1, public item2: T2) {\n  }\n}\n\n// Interfaces\ninterface Pair<T> {\n  item1: T;\n  item2: T;\n}\n\n// And functions\nlet pairToTuple = function <T>(p: Pair<T>) {\n  return new Tuple(p.item1, p.item2);\n};\n\nlet tuple = pairToTuple({ item1: \"hello\", item2: \"world\" });\n\n// Including references to a definition file:\n/// <reference path=\"jquery.d.ts\" />\n\n// Template Strings (strings that use backticks)\n// String Interpolation with Template Strings\nlet name = 'Tyrone';\nlet greeting = `Hi ${name}, how are you?`\n// Multiline Strings with Template Strings\nlet multiline = `This is an example\nof a multiline string`;\n\n// READONLY: New Feature in TypeScript 3.1\ninterface Person {\n  readonly name: string;\n  readonly age: number;\n}\n\nvar p1: Person = { name: \"Tyrone\", age: 42 };\np1.age = 25; // Error, p1.age is read-only\n\nvar p2 = { name: \"John\", age: 60 };\nvar p3: Person = p2; // Ok, read-only alias for p2\np3.age = 35; // Error, p3.age is read-only\np2.age = 45; // Ok, but also changes p3.age because of aliasing\n\nclass Car {\n  readonly make: string;\n  readonly model: string;\n  readonly year = 2018;\n\n  constructor() {\n    this.make = \"Unknown Make\"; // Assignment permitted in constructor\n    this.model = \"Unknown Model\"; // Assignment permitted in constructor\n  }\n}\n\nlet numbers: Array<number> = [0, 1, 2, 3, 4];\nlet moreNumbers: ReadonlyArray<number> = numbers;\nmoreNumbers[5] = 5; // Error, elements are read-only\nmoreNumbers.push(5); // Error, no push method (because it mutates array)\nmoreNumbers.length = 3; // Error, length is read-only\nnumbers = moreNumbers; // Error, mutating methods are missing\n\n// Tagged Union Types for modelling state that can be in one of many shapes\ntype State = \n  | { type: \"loading\" }\n  | { type: \"success\", value: number }\n  | { type: \"error\", message: string };\n\ndeclare const state: State;\nif (state.type === \"success\") {\n  console.log(state.value);\n} else if (state.type === \"error\") {\n  console.error(state.message);\n}\n\n// Iterators and Generators\n\n// for..of statement\n// iterate over the list of values on the object being iterated\nlet arrayOfAnyType = [1, \"string\", false];\nfor (const val of arrayOfAnyType) {\n    console.log(val); // 1, \"string\", false\n}\n\nlet list = [4, 5, 6];\nfor (const i of list) {\n   console.log(i); // 4, 5, 6\n}\n\n// for..in statement\n// iterate over the list of keys on the object being iterated\nfor (const i in list) {\n   console.log(i); // 0, 1, 2\n}\n\n// Type Assertion\n\nlet foo = {} // Creating foo as an empty object\nfoo.bar = 123 // Error: property 'bar' does not exist on `{}`\nfoo.baz = 'hello world' // Error: property 'baz' does not exist on `{}`\n\n// Because the inferred type of foo is `{}` (an object with 0 properties), you \n// are not allowed to add bar and baz to it. However with type assertion,\n// the following will pass:\n\ninterface Foo { \n  bar: number;\n  baz: string;\n}\n\nlet foo = {} as Foo; // Type assertion here\nfoo.bar = 123;\nfoo.baz = 'hello world'\n\n```\n\n## Further Reading\n * [TypeScript Official website] (http://www.typescriptlang.org/)\n * [TypeScript language specifications] (https://github.com/Microsoft/TypeScript/blob/master/doc/spec.md)\n * [Anders Hejlsberg - Introducing TypeScript on Channel 9] (http://channel9.msdn.com/posts/Anders-Hejlsberg-Introducing-TypeScript)\n * [Source Code on GitHub] (https://github.com/Microsoft/TypeScript)\n * [Definitely Typed - repository for type definitions] (http://definitelytyped.org/)"
