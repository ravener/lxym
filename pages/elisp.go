
package pages

const Elisp = "```scheme\n;; This gives an introduction to Emacs Lisp in 15 minutes (v0.2d)\n;;\n;; First make sure you read this text by Peter Norvig:\n;; http://norvig.com/21-days.html\n;;\n;; Then install GNU Emacs 24.3:\n;;\n;; Debian: apt-get install emacs (or see your distro instructions)\n;; OSX: http://emacsformacosx.com/emacs-builds/Emacs-24.3-universal-10.6.8.dmg\n;; Windows: http://ftp.gnu.org/gnu/windows/emacs/emacs-24.3-bin-i386.zip\n;;\n;; More general information can be found at:\n;; http://www.gnu.org/software/emacs/#Obtaining\n\n;; Important warning:\n;;\n;; Going through this tutorial won't damage your computer unless\n;; you get so angry that you throw it on the floor.  In that case,\n;; I hereby decline any responsibility.  Have fun!\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;;\n;; Fire up Emacs.\n;;\n;; Hit the `q' key to dismiss the welcome message.\n;;\n;; Now look at the gray line at the bottom of the window:\n;;\n;; \"*scratch*\" is the name of the editing space you are now in.\n;; This editing space is called a \"buffer\".\n;;\n;; The scratch buffer is the default buffer when opening Emacs.\n;; You are never editing files: you are editing buffers that you\n;; can save to a file.\n;;\n;; \"Lisp interaction\" refers to a set of commands available here.\n;;\n;; Emacs has a built-in set of commands available in every buffer,\n;; and several subsets of commands available when you activate a\n;; specific mode.  Here we use the `lisp-interaction-mode', which\n;; comes with commands to evaluate and navigate within Elisp code.\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;;\n;; Semi-colons start comments anywhere on a line.\n;;\n;; Elisp programs are made of symbolic expressions (\"sexps\"):\n(+ 2 2)\n\n;; This symbolic expression reads as \"Add 2 to 2\".\n\n;; Sexps are enclosed into parentheses, possibly nested:\n(+ 2 (+ 1 1))\n\n;; A symbolic expression contains atoms or other symbolic\n;; expressions.  In the above examples, 1 and 2 are atoms,\n;; (+ 2 (+ 1 1)) and (+ 1 1) are symbolic expressions.\n\n;; From `lisp-interaction-mode' you can evaluate sexps.\n;; Put the cursor right after the closing parenthesis then\n;; hold down the control and hit the j keys (\"C-j\" for short).\n\n(+ 3 (+ 1 2))\n;;           ^ cursor here\n;; `C-j' => 6\n\n;; `C-j' inserts the result of the evaluation in the buffer.\n\n;; `C-xC-e' displays the same result in Emacs bottom line,\n;; called the \"minibuffer\".  We will generally use `C-xC-e',\n;; as we don't want to clutter the buffer with useless text.\n\n;; `setq' stores a value into a variable:\n(setq my-name \"Bastien\")\n;; `C-xC-e' => \"Bastien\" (displayed in the mini-buffer)\n\n;; `insert' will insert \"Hello!\" where the cursor is:\n(insert \"Hello!\")\n;; `C-xC-e' => \"Hello!\"\n\n;; We used `insert' with only one argument \"Hello!\", but\n;; we can pass more arguments -- here we use two:\n\n(insert \"Hello\" \" world!\")\n;; `C-xC-e' => \"Hello world!\"\n\n;; You can use variables instead of strings:\n(insert \"Hello, I am \" my-name)\n;; `C-xC-e' => \"Hello, I am Bastien\"\n\n;; You can combine sexps into functions:\n(defun hello () (insert \"Hello, I am \" my-name))\n;; `C-xC-e' => hello\n\n;; You can evaluate functions:\n(hello)\n;; `C-xC-e' => Hello, I am Bastien\n\n;; The empty parentheses in the function's definition means that\n;; it does not accept arguments.  But always using `my-name' is\n;; boring, let's tell the function to accept one argument (here\n;; the argument is called \"name\"):\n\n(defun hello (name) (insert \"Hello \" name))\n;; `C-xC-e' => hello\n\n;; Now let's call the function with the string \"you\" as the value\n;; for its unique argument:\n(hello \"you\")\n;; `C-xC-e' => \"Hello you\"\n\n;; Yeah!\n\n;; Take a breath.\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;;\n;; Now switch to a new buffer named \"*test*\" in another window:\n\n(switch-to-buffer-other-window \"*test*\")\n;; `C-xC-e'\n;; => [screen has two windows and cursor is in the *test* buffer]\n\n;; Mouse over the top window and left-click to go back.  Or you can\n;; use `C-xo' (i.e. hold down control-x and hit o) to go to the other\n;; window interactively.\n\n;; You can combine several sexps with `progn':\n(progn\n  (switch-to-buffer-other-window \"*test*\")\n  (hello \"you\"))\n;; `C-xC-e'\n;; => [The screen has two windows and cursor is in the *test* buffer]\n\n;; Now if you don't mind, I'll stop asking you to hit `C-xC-e': do it\n;; for every sexp that follows.\n\n;; Always go back to the *scratch* buffer with the mouse or `C-xo'.\n\n;; It's often useful to erase the buffer:\n(progn\n  (switch-to-buffer-other-window \"*test*\")\n  (erase-buffer)\n  (hello \"there\"))\n\n;; Or to go back to the other window:\n(progn\n  (switch-to-buffer-other-window \"*test*\")\n  (erase-buffer)\n  (hello \"you\")\n  (other-window 1))\n\n;; You can bind a value to a local variable with `let':\n(let ((local-name \"you\"))\n  (switch-to-buffer-other-window \"*test*\")\n  (erase-buffer)\n  (hello local-name)\n  (other-window 1))\n\n;; No need to use `progn' in that case, since `let' also combines\n;; several sexps.\n\n;; Let's format a string:\n(format \"Hello %s!\\n\" \"visitor\")\n\n;; %s is a place-holder for a string, replaced by \"visitor\".\n;; \\n is the newline character.\n\n;; Let's refine our function by using format:\n(defun hello (name)\n  (insert (format \"Hello %s!\\n\" name)))\n\n(hello \"you\")\n\n;; Let's create another function which uses `let':\n(defun greeting (name)\n  (let ((your-name \"Bastien\"))\n    (insert (format \"Hello %s!\\n\\nI am %s.\"\n                    name       ; the argument of the function\n                    your-name  ; the let-bound variable \"Bastien\"\n                    ))))\n\n;; And evaluate it:\n(greeting \"you\")\n\n;; Some functions are interactive:\n(read-from-minibuffer \"Enter your name: \")\n\n;; Evaluating this function returns what you entered at the prompt.\n\n;; Let's make our `greeting' function prompt for your name:\n(defun greeting (from-name)\n  (let ((your-name (read-from-minibuffer \"Enter your name: \")))\n    (insert (format \"Hello!\\n\\nI am %s and you are %s.\"\n                    from-name ; the argument of the function\n                    your-name ; the let-bound var, entered at prompt\n                    ))))\n\n(greeting \"Bastien\")\n\n;; Let's complete it by displaying the results in the other window:\n(defun greeting (from-name)\n  (let ((your-name (read-from-minibuffer \"Enter your name: \")))\n    (switch-to-buffer-other-window \"*test*\")\n    (erase-buffer)\n    (insert (format \"Hello %s!\\n\\nI am %s.\" your-name from-name))\n    (other-window 1)))\n\n;; Now test it:\n(greeting \"Bastien\")\n\n;; Take a breath.\n\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n;;\n;; Let's store a list of names:\n;; If you want to create a literal list of data, use ' to stop it from\n;; being evaluated - literally, \"quote\" the data.\n(setq list-of-names '(\"Sarah\" \"Chloe\" \"Mathilde\"))\n\n;; Get the first element of this list with `car':\n(car list-of-names)\n\n;; Get a list of all but the first element with `cdr':\n(cdr list-of-names)\n\n;; Add an element to the beginning of a list with `push':\n(push \"Stephanie\" list-of-names)\n\n;; NOTE: `car' and `cdr' don't modify the list, but `push' does.\n;; This is an important difference: some functions don't have any\n;; side-effects (like `car') while others have (like `push').\n\n;; Let's call `hello' for each element in `list-of-names':\n(mapcar 'hello list-of-names)\n\n;; Refine `greeting' to say hello to everyone in `list-of-names':\n(defun greeting ()\n    (switch-to-buffer-other-window \"*test*\")\n    (erase-buffer)\n    (mapcar 'hello list-of-names)\n    (other-window 1))\n\n(greeting)\n\n;; Remember the `hello' function we defined above?  It takes one\n;; argument, a name.  `mapcar' calls `hello', successively using each\n;; element of `list-of-names' as the argument for `hello'.\n\n;; Now let's arrange a bit what we have in the displayed buffer:\n\n(defun replace-hello-by-bonjour ()\n    (switch-to-buffer-other-window \"*test*\")\n    (goto-char (point-min))\n    (while (search-forward \"Hello\")\n      (replace-match \"Bonjour\"))\n    (other-window 1))\n\n;; (goto-char (point-min)) goes to the beginning of the buffer.\n;; (search-forward \"Hello\") searches for the string \"Hello\".\n;; (while x y) evaluates the y sexp(s) while x returns something.\n;; If x returns `nil' (nothing), we exit the while loop.\n\n(replace-hello-by-bonjour)\n\n;; You should see all occurrences of \"Hello\" in the *test* buffer\n;; replaced by \"Bonjour\".\n\n;; You should also get an error: \"Search failed: Hello\".\n;;\n;; To avoid this error, you need to tell `search-forward' whether it\n;; should stop searching at some point in the buffer, and whether it\n;; should silently fail when nothing is found:\n\n;; (search-forward \"Hello\" nil t) does the trick:\n\n;; The `nil' argument says: the search is not bound to a position.\n;; The `'t' argument says: silently fail when nothing is found.\n\n;; We use this sexp in the function below, which doesn't throw an error:\n\n(defun hello-to-bonjour ()\n    (switch-to-buffer-other-window \"*test*\")\n    (erase-buffer)\n    ;; Say hello to names in `list-of-names'\n    (mapcar 'hello list-of-names)\n    (goto-char (point-min))\n    ;; Replace \"Hello\" by \"Bonjour\"\n    (while (search-forward \"Hello\" nil t)\n      (replace-match \"Bonjour\"))\n    (other-window 1))\n\n(hello-to-bonjour)\n\n;; Let's boldify the names:\n\n(defun boldify-names ()\n    (switch-to-buffer-other-window \"*test*\")\n    (goto-char (point-min))\n    (while (re-search-forward \"Bonjour \\\\(.+\\\\)!\" nil t)\n      (add-text-properties (match-beginning 1)\n                           (match-end 1)\n                           (list 'face 'bold)))\n    (other-window 1))\n\n;; This functions introduces `re-search-forward': instead of\n;; searching for the string \"Bonjour\", you search for a pattern,\n;; using a \"regular expression\" (abbreviated in the prefix \"re-\").\n\n;; The regular expression is \"Bonjour \\\\(.+\\\\)!\" and it reads:\n;; the string \"Bonjour \", and\n;; a group of            | this is the \\\\( ... \\\\) construct\n;;   any character       | this is the .\n;;   possibly repeated   | this is the +\n;; and the \"!\" string.\n\n;; Ready?  Test it!\n\n(boldify-names)\n\n;; `add-text-properties' adds... text properties, like a face.\n\n;; OK, we are done.  Happy hacking!\n\n;; If you want to know more about a variable or a function:\n;;\n;; C-h v a-variable RET\n;; C-h f a-function RET\n;;\n;; To read the Emacs Lisp manual with Emacs:\n;;\n;; C-h i m elisp RET\n;;\n;; To read an online introduction to Emacs Lisp:\n;; https://www.gnu.org/software/emacs/manual/html_node/eintr/index.html\n```"