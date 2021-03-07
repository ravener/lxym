
package pages

const Clojure_Macros = "As with all Lisps, Clojure's inherent [homoiconicity](https://en.wikipedia.org/wiki/Homoiconic)\ngives you access to the full extent of the language to write code-generation routines\ncalled \"macros\". Macros provide a powerful way to tailor the language to your needs.\n\nBe careful though. It's considered bad form to write a macro when a function will do.\nUse a macro only when you need control over when or if the arguments to a form will\nbe evaluated.\n\nYou'll want to be familiar with Clojure. Make sure you understand everything in\n[Clojure in Y Minutes](/docs/clojure/).\n\n```clojure\n;; Define a macro using defmacro. Your macro should output a list that can\n;; be evaluated as clojure code.\n;;\n;; This macro is the same as if you wrote (reverse \"Hello World\")\n(defmacro my-first-macro []\n  (list reverse \"Hello World\"))\n\n;; Inspect the result of a macro using macroexpand or macroexpand-1.\n;;\n;; Note that the call must be quoted.\n(macroexpand '(my-first-macro))\n;; -> (#<core$reverse clojure.core$reverse@xxxxxxxx> \"Hello World\")\n\n;; You can eval the result of macroexpand directly:\n(eval (macroexpand '(my-first-macro)))\n; -> (\\d \\l \\o \\r \\W \\space \\o \\l \\l \\e \\H)\n\n;; But you should use this more succinct, function-like syntax:\n(my-first-macro)  ; -> (\\d \\l \\o \\r \\W \\space \\o \\l \\l \\e \\H)\n\n;; You can make things easier on yourself by using the more succinct quote syntax\n;; to create lists in your macros:\n(defmacro my-first-quoted-macro []\n  '(reverse \"Hello World\"))\n\n(macroexpand '(my-first-quoted-macro))\n;; -> (reverse \"Hello World\")\n;; Notice that reverse is no longer function object, but a symbol.\n\n;; Macros can take arguments.\n(defmacro inc2 [arg]\n  (list + 2 arg))\n\n(inc2 2) ; -> 4\n\n;; But, if you try to do this with a quoted list, you'll get an error, because\n;; the argument will be quoted too. To get around this, clojure provides a\n;; way of quoting macros: `. Inside `, you can use ~ to get at the outer scope\n(defmacro inc2-quoted [arg]\n  `(+ 2 ~arg))\n\n(inc2-quoted 2)\n\n;; You can use the usual destructuring args. Expand list variables using ~@\n(defmacro unless [arg & body]\n  `(if (not ~arg)\n     (do ~@body))) ; Remember the do!\n\n(macroexpand '(unless true (reverse \"Hello World\")))\n;; ->\n;; (if (clojure.core/not true) (do (reverse \"Hello World\")))\n\n;; (unless) evaluates and returns its body if the first argument is false.\n;; Otherwise, it returns nil\n\n(unless true \"Hello\") ; -> nil\n(unless false \"Hello\") ; -> \"Hello\"\n\n;; Used without care, macros can do great evil by clobbering your vars\n(defmacro define-x []\n  '(do\n     (def x 2)\n     (list x)))\n\n(def x 4)\n(define-x) ; -> (2)\n(list x) ; -> (2)\n\n;; To avoid this, use gensym to get a unique identifier\n(gensym 'x) ; -> x1281 (or some such thing)\n\n(defmacro define-x-safely []\n  (let [sym (gensym 'x)]\n    `(do\n       (def ~sym 2)\n       (list ~sym))))\n\n(def x 4)\n(define-x-safely) ; -> (2)\n(list x) ; -> (4)\n\n;; You can use # within ` to produce a gensym for each symbol automatically\n(defmacro define-x-hygienically []\n  `(do\n     (def x# 2)\n     (list x#)))\n\n(def x 4)\n(define-x-hygienically) ; -> (2)\n(list x) ; -> (4)\n\n;; It's typical to use helper functions with macros. Let's create a few to\n;; help us support a (dumb) inline arithmetic syntax\n(declare inline-2-helper)\n(defn clean-arg [arg]\n  (if (seq? arg)\n    (inline-2-helper arg)\n    arg))\n\n(defn apply-arg\n  \"Given args [x (+ y)], return (+ x y)\"\n  [val [op arg]]\n  (list op val (clean-arg arg)))\n\n(defn inline-2-helper\n  [[arg1 & ops-and-args]]\n  (let [ops (partition 2 ops-and-args)]\n    (reduce apply-arg (clean-arg arg1) ops)))\n\n;; We can test it immediately, without creating a macro\n(inline-2-helper '(a + (b - 2) - (c * 5))) ; -> (- (+ a (- b 2)) (* c 5))\n\n; However, we'll need to make it a macro if we want it to be run at compile time\n(defmacro inline-2 [form]\n  (inline-2-helper form))\n\n(macroexpand '(inline-2 (1 + (3 / 2) - (1 / 2) + 1)))\n; -> (+ (- (+ 1 (/ 3 2)) (/ 1 2)) 1)\n\n(inline-2 (1 + (3 / 2) - (1 / 2) + 1))\n; -> 3 (actually, 3N, since the number got cast to a rational fraction with /)\n```\n\n### Further Reading\n\n[Writing Macros](http://www.braveclojure.com/writing-macros/)\n\n[Official docs](http://clojure.org/macros)\n\n[When to use macros?](https://lispcast.com/when-to-use-a-macro/)"
