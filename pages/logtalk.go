
package pages

const Logtalk = "Logtalk is an object-oriented logic programming language that extends and leverages Prolog with modern code encapsulation and code reuse mechanisms without compromising its declarative programming features. Logtalk is implemented in highly portable code and can use most modern and standards compliant Prolog implementations as a back-end compiler.\n\nTo keep its size reasonable, this tutorial necessarily assumes that the reader have a working knowledge of Prolog and is biased towards describing Logtalk object-oriented features.\n\n# Syntax\n\nLogtalk uses standard Prolog syntax with the addition of a few operators and directives for a smooth learning curve and wide portability. One important consequence is that Prolog code can be easily encapsulated in objects with little or no changes. Moreover, Logtalk can transparently interpret most Prolog modules as Logtalk objects.\n\nThe main operators are:\n\n* `::/2` - sending a message to an object\n* `::/1` - sending a message to _self_ (i.e. to the object that received the message being processed)\n* `^^/1` - _super_ call (of an inherited or imported predicate)\n\nSome of the most important entity and predicate directives will be introduced in the next sections.\n\n# Entities and roles\n\nLogtalk provides _objects_, _protocols_, and _categories_ as first-class entities. Relations between entities define _patterns of code reuse_ and the _roles_ played by the entities. For example, when an object _instantiates_ another object, the first object plays the role of an instance and the second object plays the role of a class. An _extends_ relation between two objects implies that both objects play the role of prototypes, with one of them extending the other, its parent prototype.\n\n# Defining an object\n\nAn object encapsulates predicate declarations and definitions. Objects can be created dynamically but are usually static and defined in source files. A single source file can contain any number of entity definitions. A simple object, defining a list member public predicate:\n\n```logtalk\n:- object(list).\n\n\t:- public(member/2).\n\tmember(Head, [Head| _]).\n\tmember(Head, [_| Tail]) :-\n\t\tmember(Head, Tail).\n\n:- end_object.\n```\n\n# Compiling and loading source files\n\nAssuming that the code above for the `list` object is saved in a `list.lgt` file, it can be compiled and loaded using the `logtalk_load/1` built-in predicate or its abbreviation, `{}/1`, with the file path as argument (the extension can be omitted):\n\n```logtalk\n?- {list}.\nyes\n```\n\nIn general, entities may have dependencies on entities defined in other source files (e.g. library entities). To load a file and all its dependencies, the advised solution is to define a \n_loader_ file that loads all the necessary files for an application. A loader file is simply a source file, typically named `loader.lgt`, that makes calls to the `logtalk_load/1-2`\nbuilt-in predicates, usually from an `initialization/1` directive for portability and\nstandards compliance. Loader files are provided for all libraries, tools, and examples.\n\n# Sending a message to an object\n\nThe `::/2` infix operator is used to send a message to an object. As in Prolog, we can backtrack for alternative solutions:\n\n```logtalk\n?- list::member(X, [1,2,3]).\nX = 1 ;\nX = 2 ;\nX = 3\nyes\n```\n\nEncapsulation is enforced. A predicate can be declared _public_, _protected_, or _private_. It can also be _local_ when there is no scope directive for it. For example:\n\n```logtalk\n:- object(scopes).\n\n\t:- private(bar/0).\n\tbar.\n\n\tlocal.\n\n:- end_object.\n```\n\nAssuming the object is saved in a `scopes.lgt` file:\n\n```logtalk\n?- {scopes}.\nyes\n\n?- catch(scopes::bar, Error, true).\nError = error(\n\tpermission_error(access, private_predicate, bar/0),\n\tlogtalk(scopes::bar, user)\n)\nyes\n\n?- catch(scopes::local, Error, true).\nError = error(\n\texistence_error(predicate_declaration, local/0),\n\tlogtalk(scopes::local, user)\n)\nyes\n```\n\nWhen the predicate in a message is unknown for the object (the role it plays determines the lookup procedures), we also get an error. For example:\n\n```logtalk\n?- catch(scopes::unknown, Error, true).\nError = error(\n\texistence_error(predicate_declaration, unknown/0),\n\tlogtalk(scopes::unknown, user)\n)\nyes\n```\n\nA subtle point is that predicate scope directives specify predicate _calling_ semantics, not _definition_ semantics. For example, if an object playing the role of a class declares a predicate private, the predicate can be defined in subclasses and instances *but* can only be called in its instances _from_ the class.\n\n# Defining and implementing a protocol\n\nProtocols contain predicate declarations that can be implemented by any number of objects and categories:\n\n```logtalk\n:- protocol(listp).\n\n\t:- public(member/2).\n\n:- end_protocol.\n\n:- object(list,\n\timplements(listp)).\n\n\tmember(Head, [Head| _]).\n\tmember(Head, [_| Tail]) :-\n\t\tmember(Head, Tail).\n\n:- end_object.\n```\n\nThe scope of the protocol predicates can be restricted using protected or private implementation. For example:\n\n```logtalk\n:- object(stack,\n\timplements(private::listp)).\n\n:- end_object.\n```\n\nIn fact, all entity relations (in an entity opening directive) can be qualified as public (the default), protected, or private.\n\n# Prototypes\n\nAn object without an _instantiation_ or _specialization_ relation with another object plays the role of a prototype. A prototype can _extend_ another object, its parent prototype.\n\n```logtalk\n% clyde, our prototypical elephant\n:- object(clyde).\n\n\t:- public(color/1).\n\tcolor(grey).\n\n\t:- public(number_of_legs/1).\n\tnumber_of_legs(4).\n\n:- end_object.\n\n% fred, another elephant, is like clyde, except that he's white\n:- object(fred,\n\textends(clyde)).\n\n\tcolor(white).\n\n:- end_object.\n```\n\nWhen answering a message sent to an object playing the role of a prototype, we validate the message and look for an answer first in the prototype itself and, if not found, we delegate to the prototype parents if any:\n\n```logtalk\n?- fred::number_of_legs(N).\nN = 4\nyes\n\n?- fred::color(C).\nC = white\nyes\n```\n\nA message is valid if the corresponding predicate is declared (and the sender is within scope) but it will fail, rather then throwing an error, if the predicate is not defined. This is called the _closed-world assumption_. For example, consider the following object, saved in a `foo.lgt` file:\n\n```logtalk\n:- object(foo).\n\n\t:- public(bar/0).\n\n:- end_object.\n```\n\nLoading the file and trying to call the `bar/0` predicate fails as expected. Note that this is different from calling an _unknown_ predicate, which results in an error:\n\n```logtalk\n?- {foo}.\nyes\n\n?- foo::bar.\nno\n\n?- catch(foo::baz, Error, true).\nError = error(\n\texistence_error(predicate_declaration, baz/0),\n\tlogtalk(foo::baz, user)\n)\nyes\n```\n\n# Classes and instances\n\nIn order to define objects playing the role of classes and/or instances, an object must have at least an instantiation or a specialization relation with another object. Objects playing the role of meta-classes can be used when we need to see a class also as an instance. We use the following example to also illustrate how to dynamically create new objects at runtime:\n\n```logtalk\n% a simple, generic, metaclass defining a new/2 predicate for its instances\n:- object(metaclass,\n\tinstantiates(metaclass)).\n\n\t:- public(new/2).\n\tnew(Instance, Clauses) :-\n\t\tself(Class),\n\t\tcreate_object(Instance, [instantiates(Class)], [], Clauses).\n\n:- end_object.\n\n% a simple class defining age/1 and name/1 predicate for its instances\n:- object(person,\n\tinstantiates(metaclass)).\n\n\t:- public([\n\t\tage/1, name/1\n\t]).\n\n\t% a default value for age/1\n\tage(42).\n\n:- end_object.\n\n% a static instance of the class person\n:- object(john,\n\tinstantiates(person)).\n\n\tname(john).\n\tage(12).\n\n:- end_object.\n```\n\nWhen answering a message sent to an object playing the role of an instance, we validate the message by starting in its class and going up to its class superclasses if necessary. Assuming that the message is valid, then we look for an answer starting in the instance itself:\n\n```logtalk\n?- person::new(Instance, [name(paulo)]).\nInstance = o1\nyes\n\n?- o1::name(Name).\nName = paulo\nyes\n\n?- o1::age(Age).\nAge = 42\nyes\n\n?- john::age(Age).\nAge = 12\nyes\n```\n\n# Categories\n\nA category is a fine grained unit of code reuse, used to encapsulate a _cohesive_ set of predicate declarations and definitions, implementing a _single_ functionality, that can be imported into any object. A category can thus be seen as the dual concept of a protocol. In the following example, we define categories representing car engines and then import them into car objects:\n\n```logtalk\n% a protocol describing engine characteristics\n:- protocol(carenginep).\n\n\t:- public([\n\t\treference/1,\n\t\tcapacity/1,\n\t\tcylinders/1,\n\t\thorsepower_rpm/2,\n\t\tbore_stroke/2,\n\t\tfuel/1\n\t]).\n\n:- end_protocol.\n\n% a typical engine defined as a category\n:- category(classic,\n\timplements(carenginep)).\n\n\treference('M180.940').\n\tcapacity(2195).\n\tcylinders(6).\n\thorsepower_rpm(94, 4800).\n\tbore_stroke(80, 72.8).\n\tfuel(gasoline).\n\n:- end_category.\n\n% a souped up version of the previous engine\n:- category(sport,\n\textends(classic)).\n\n\treference('M180.941').\n\thorsepower_rpm(HP, RPM) :-\n\t\t^^horsepower_rpm(ClassicHP, ClassicRPM),\t% \"super\" call\n\t\tHP is truncate(ClassicHP*1.23),\n\t\tRPM is truncate(ClassicRPM*0.762).\n\n:- end_category.\n\n% with engines (and other components), we may start \"assembling\" some cars\n:- object(sedan,\n\timports(classic)).\n\n:- end_object.\n\n:- object(coupe,\n\timports(sport)).\n\n:- end_object.\n```\n\nCategories are independently compiled and thus allow importing objects to be updated by simple updating the imported categories without requiring object recompilation. Categories also provide _runtime transparency_. I.e. the category protocol adds to the protocol of the objects importing the category:\n\n```logtalk\n?- sedan::current_predicate(Predicate).\nPredicate = reference/1 ;\nPredicate = capacity/1 ;\nPredicate = cylinders/1 ;\nPredicate = horsepower_rpm/2 ;\nPredicate = bore_stroke/2 ;\nPredicate = fuel/1\nyes\n```\n\n# Hot patching\n\nCategories can be also be used for hot-patching objects. A category can add new predicates to an object and/or replace object predicate definitions. For example, consider the following object:\n\n```logtalk\n:- object(buggy).\n\n\t:- public(p/0).\n\tp :- write(foo).\n\n:- end_object.\n```\n\nAssume that the object prints the wrong string when sent the message `p/0`:\n\n```logtalk\n?- {buggy}.\nyes\n\n?- buggy::p.\nfoo\nyes\n```\n\nIf the object source code is not available and we need to fix an application running the object code, we can simply define a category that fixes the buggy predicate:\n\n```logtalk\n:- category(patch,\n\tcomplements(buggy)).\n\n\t% fixed p/0 def\n\tp :- write(bar).\n\n:- end_category.\n```\n\nAfter compiling and loading the category into the running application we will now get:\n\n```logtalk\n?- {patch}.\nyes\n\n?- buggy::p.\nbar\nyes\n```\n\nAs hot-patching forcefully breaks encapsulation, there is a `complements` compiler flag that can be set (globally or on a per-object basis) to allow, restrict, or prevent it.\n\n# Parametric objects and categories\n\nObjects and categories can be parameterized by using as identifier a compound term instead of an atom. Object and category parameters are _logical variables_ shared with all encapsulated predicates. An example with geometric circles:\n\n```logtalk\n:- object(circle(_Radius, _Color)).\n\n\t:- public([\n\t\tarea/1, perimeter/1\n\t]).\n\n\tarea(Area) :-\n\t\tparameter(1, Radius),\n\t\tArea is pi*Radius*Radius.\n\n\tperimeter(Perimeter) :-\n\t\tparameter(1, Radius),\n\t\tPerimeter is 2*pi*Radius.\n\n:- end_object.\n```\n\nParametric objects are used just as any other object, usually providing values for the parameters when sending a message:\n\n```logtalk\n?- circle(1.23, blue)::area(Area).\nArea = 4.75291\nyes\n```\n\nParametric objects also provide a simple way of associating a set of predicates with a plain Prolog predicate. Prolog facts can be interpreted as _parametric object proxies_ when they have the same functor and arity as the identifiers of parametric objects. Handy syntax is provided to for working with proxies. For example, assuming the following clauses for a `circle/2` predicate:\n\n```logtalk\ncircle(1.23, blue).\ncircle(3.71, yellow).\ncircle(0.39, green).\ncircle(5.74, black).\ncircle(8.32, cyan).\n```\n\nWith these clauses loaded, we can easily compute for example a list with the areas of all the circles:\n\n```logtalk\n?- findall(Area, {circle(_, _)}::area(Area), Areas).\nAreas = [4.75291, 43.2412, 0.477836, 103.508, 217.468]\nyes\n```\n\nThe `{Goal}::Message` construct proves `Goal`, possibly instantiating any variables in it, and sends `Message` to the resulting term.\n\n# Events and monitors\n\nLogtalk supports _event-driven programming_ by allowing defining events and monitors for those events. An event is simply the sending of a message to an object. Interpreting message sending as an atomic activity, a _before_ event and an _after_ event are recognized. Event monitors define event handler predicates, `before/3` and `after/3`, and can query, register, and delete a system-wide event registry that associates events with monitors. For example, a simple tracer for any message being sent using the `::/2` control construct can be defined as:\n\n```logtalk\n:- object(tracer,\n\timplements(monitoring)).    % built-in protocol for event handlers\n\n\t:- initialization(define_events(_, _, _, _, tracer)).\n\n\tbefore(Object, Message, Sender) :-\n\t\twrite('call: '), writeq(Object), write(' <-- '), writeq(Message),\n\t\twrite(' from '), writeq(Sender), nl.\n\n\tafter(Object, Message, Sender) :-\n\t\twrite('exit: '), writeq(Object), write(' <-- '), writeq(Message),\n\t\twrite(' from '), writeq(Sender), nl.\n\n:- end_object.\n```\n\nAssuming that the `tracer` object and the `list` object defined earlier are compiled and loaded, we can observe the event handlers in action by sending a message:\n\n```logtalk\n?- list::member(X, [1,2,3]).\n\ncall: list <-- member(X, [1,2,3]) from user\nexit: list <-- member(1, [1,2,3]) from user\nX = 1 ;\nexit: list <-- member(2, [1,2,3]) from user\nX = 2 ;\nexit: list <-- member(3, [1,2,3]) from user\nX = 3\nyes\n```\n\nEvents can be set and deleted dynamically at runtime by calling the `define_events/5` and `abolish_events/5` built-in predicates.\n\nEvent-driven programming can be seen as a form of _computational reflection_. But note that events are only generated when using the `::/2` message-sending control construct.\n\n# Lambda expressions\n\nLogtalk supports lambda expressions. Lambda parameters are represented using a list with the `(>>)/2` infix operator connecting them to the lambda. Some simple examples using library meta-predicates:\n\n```logtalk\n?- {library(metapredicates_loader)}.\nyes\n\n?- meta::map([X,Y]>>(Y is 2*X), [1,2,3], Ys).\nYs = [2,4,6]\nyes\n```\n\nCurrying is also supported:\n\n```logtalk\n?- meta::map([X]>>([Y]>>(Y is 2*X)), [1,2,3], Ys).\nYs = [2,4,6]\nyes\n```\n\nLambda free variables can be expressed using the extended syntax `{Free1, ...}/[Parameter1, ...]>>Lambda`.\n\n# Macros\n\nTerms and goals in source files can be _expanded_ at compile time by specifying a _hook object_ that defines term-expansion and goal-expansion rules. For example, consider the following simple object, saved in a `source.lgt` file:\n\n```logtalk\n:- object(source).\n\n\t:- public(bar/1).\n\tbar(X) :- foo(X).\n\n\tfoo(a). foo(b). foo(c).\n\n:- end_object.\n```\n\nAssume the following hook object, saved in a `my_macros.lgt` file, that expands clauses and calls to the `foo/1` local predicate:\n\n```logtalk\n:- object(my_macros,\n\timplements(expanding)).    % built-in protocol for expanding predicates\n\n\tterm_expansion(foo(Char), baz(Code)) :-\n\t\tchar_code(Char, Code). % standard built-in predicate\n\n\tgoal_expansion(foo(X), baz(X)).\n\n:- end_object.\n```\n\nAfter loading the macros file, we can then expand our source file with it using the `hook` compiler flag:\n\n```logtalk\n?- logtalk_load(my_macros), logtalk_load(source, [hook(my_macros)]).\nyes\n\n?- source::bar(X).\nX = 97 ;\nX = 98 ;\nX = 99\ntrue\n```\n\nThe Logtalk library provides support for combining hook objects using different workflows (for example, defining a pipeline of expansions).\n\n# Further information\n\nVisit the [Logtalk website](http://logtalk.org) for more information."
