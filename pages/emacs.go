
package pages

const Emacs = "Emacs started its life as [\"the extensible, customizable display\neditor\"](https://www.gnu.org/software/emacs/emacs-paper.html) and grew\nover the years into a full-blown ecosystem. Many tasks, usually\nrelegated to a diverse set of tools can be accomplished from within\nEmacs in a consistent, familiar interface. Examples include directory\nmanagement, viewing PDF documents, editing files over SSH, managing git\nrepos,… (the list is quite long). In short, Emacs is yours to make of it\nwhat you will: the spectrum of users varies from those who use it to\nedit text files to extreme purists who use it to virtually replace their\noperating system.\n\nEmacs is extensible via a specialized dialect of Lisp known as Emacs\nLisp (Elisp) which has a lot of macros geared towards editing text and\nmanaging text buffers. Any key (combination) you use in Emacs is bound\nto an Emacs Lisp function and may be remapped to any other function,\nincluding ones you write\nyourself.\n\n# Key Notation\n\n``` text\nThe Emacs manual and the community in general uses a convention to refer to different key combinations used within Emacs. Specifically, Emacs has the notion of a \"modifier key\" that is pressed along with another key to modify its action.\n\nAn example of this notation is \"C-c\". In this key combination \"C\" is the modifier and stands for the \"Ctrl\" key and \"c\" is the key whose action is being modified (the literal character \"c\").\n\nThe modifier shorthand:\n\"C-\" --> The \"CTRL\" key\n\"M-\" --> The \"Meta\" key (usually, the \"Alt\" key)\n\"s-\" --> The \"Super\" key (the \"Cmd\" key on Macs and the \"Windows\" key on PCs)\n\nThere are other, less commonly used modifiers that I will not get into here.\n\nThe key combination \"C-x C-s\" means you press \"Ctrl+x\" followed by \"Ctrl+s\"\n\nIn addition to the above modifiers, the special keys \"Esc\", \"Return (Enter)\" and \"Shift\" are denoted by \"ESC\", \"RET\" and \"S\", respectively.\n```\n\n# Basic Emacs Concepts\n\nHere, I discuss some basic Emacs concepts and terminology that may be\nconfusing to newcomers (especially to people used to Vim terminology)\n\n  - A bunch of text that Emacs is editing is known as a **buffer**\n  - A buffer does not necessarily correspond to an actual file on disk.\n    It may be just a bunch of text in memory.\n  - When a buffer corresponds to a file on disk, we say that the buffer\n    is **visiting** that file.\n  - Emacs typically has many buffers open at once.\n  - The display of Emacs may be split into different **windows** (not to\n    be confused with your operating system's windows: the operating\n    system window for Emacs can have multiple Emacs windows inside it).\n  - An operating system window for Emacs is called an Emacs **frame**.\n    Thus, when the Emacs manual talks about opening a new frame, this\n    essentially means opening a new OS *window* containing an(other)\n    instance of Emacs.\n  - The concepts conventionally known as cutting and pasting are\n    referred to as **killing** and **yanking**, respectively in Emacs\n    parlance.\n  - The current position of the cursor is called the **point** in Emacs.\n    Technically, **point** is defined as the position right before the\n    character where the cursor currently is.\n  - Finally, each buffer may have several **modes** associated with it:\n    a **major mode** and possibly several **minor modes**.\n  - The **major mode** defines the main behavior of Emacs in the\n    currently selected buffer. This can be roughly thought of as the\n    file type. For example, if you're editing a Python file, the major\n    mode is (by default) `python-mode` which causes Emacs to highlight\n    Python syntax and automatically indent and outdent your code blocks\n    as syntactically required by your Python code.\n  - **Minor modes** define subtle changes in behavior and several minor\n    modes may be active at once in the same buffer. An example minor\n    mode is `flyspell-mode` which automatically highlights spelling\n    errors in your\nbuffer.\n\n# Navigation Basics\n\n``` text\nThe GUI version of Emacs can be navigated with the mouse like you would expect from a conventional GUI text editor.\n\nThe aim here is to focus on navigation solely using the keyboard as this enhances productivity immensely.\n\n\n* Line movement\n\nC-n --> Next line\nC-p --> Previous line\n\n* Character movement\n\nC-f --> Go forward one character\nC-b --> Go backward one character\n\n* Word movement\n\nM-f --> Go forward one word\nM-b --> Go backward one word\n\n* Sentence movement\n\nM-a --> Move to the beginning of the sentence\nM-e --> Move to the end of the sentence\n\n* Beginning and end of line\n\nC-a --> Move to the beginning of the line\nC-e --> Move to the end of the line\n\n* Beginning and end of buffer\n\nM-< (\"Meta+Shift+,\") --> Go to the beginning of the buffer \nM-> (\"Meta+Shift+.\") --> Go to the end of the buffer \n\n* Screen movement\n\nC-v --> Scroll down by one screen-full (the last two lines of the previous screen are kept as overlap for a smoother transition)\nM-v --> Scroll up by one screen-full (same as above but with the first two lines)\n\n* Centering the screen\n\nC-l --> Move current line to the screen's center\n\nThe above key combination actually cycles through different states depending on how many times it's been pressed.\n\nC-l --> Move current line to the screen's center\nC-l C-l --> Move current line to the top of the screen\nC-l C-l C-l --> Restore the position of the current line to where it was before the first C-l was pressed\n\nIf you press \"C-l\" a 4th time, it cycles back to centering the current line.\n\n* Repeating movement commands\n\nMost movement commands take a numerical prefix argument that says \"repeat the following command that many times\".\n\nExample:\n\nC-u 3 C-p  --> Go up 3 lines\nC-u 5 C-f  --> Go forward 5 characters\n\nOne notable exception are the screen scrolling commands:\n\nC-u 3 C-v  --> Scroll downward 3 lines (maintaining the position of the cursor)\n```\n\nBonus: many of the above navigation commands are the default navigation\ncommands in Bash (e.g. pressing \"C-b\" while entering a Bash command\ntakes you back one\ncharacter).\n\n# File editing basics\n\n``` text\n* Quitting Emacs [ Now you can't say you don't know how to quit Emacs :-) ]\n\nC-x C-c --> Quit Emacs and get prompted to save any unsaved files (buffers not visiting a file will simply be discarded unless you're running in client-server mode)\n\n* Saving a buffer\n\nC-x C-s --> Save the current buffer. If not visiting a file, it will prompt you for a file name to use to save the buffer.\n\n* Searching within a buffer\n\nC-s --> Search forwards within the buffer. Search is incremental and case-insensitive by default.\n        Press C-s to move to the next match.\n        If you press \"RET\", point is moved to the currently highlighted word and the search ends.\nC-r --> Same as C-s except it searches backward\n\nC-_ or C-/ --> Undo the last action. Keep pressing it to move up the undo tree.\nC-? or M-_ --> Redo the previous change\n\nThe \"undo\" and \"redo\" commands can take prefix numerical arguments to undo or redo that many actions:\n\nC-u 3 C-_ --> Undo the last 3 changes.\n```\n\n# Executing Elisp Functions\n\n``` text\nYou can execute any currently loaded Elisp functions (including ones you have written yourself) via \"M-x\"\n\nM-x RET  --> Prompts you for name of function to execute (Tab completion is available).\n\nExample:\n\nM-x RET search-forward-regexp RET --> Prompts you for a regular expression and searches forward in the buffer for it\n```\n\n# Emacs Configuration\n\nEmacs is configured using Elisp. On startup, it looks for a\nconfiguration file either in `~/.emacs` or `~/.emacs.d/init.el` where\n`~` refers to your home directory. If you're on Windows, consult [this\narticle](https://www.gnu.org/software/emacs/manual/html_node/efaq-w32/Location-of-init-file.html)\nfor the appropriate location of your configuration file.\n\n# Vim inside Emacs\n\nIf you are considering the transition from Vim to Emacs and you're put\noff by the non-modal nature of Emacs editing, there is an Emacs\nextension known as `evil-mode` which lets you have many Vim concepts\ninside Emacs. Here are some things added to Emacs by `evil-mode`:\n\n  - Modal editing: you get normal, insert, visual and block visual modes\n    like Vim. In addition, you get an \"Emacs\" mode where movement and\n    navigation follow the Emacs bindings.\n  - Same movement keys as Vim in normal mode\n  - Leader key combinations\n  - Pressing \":\" in normal mode allows you to execute commands\n    (including system commands)\n\nIn my own experience, `evil-mode` helps make the transition seamless and\nallows you to blend the arguably more intuitive and ergonomic\nkeybindings of Vim with the unbridled power of Emacs for a truly\nsuperior editing experience.\n\n# Discoverable Help\n\nEmacs features a pretty powerful help system that allows you to discover\nnew functionality all the\ntime.\n\n``` text\nObtaining help on specific topics. Tab completion is available for function and variable names.\n\nC-h f RET --> Prompts you for the name of an elisp function and\n              displays help text on it along with a clickable link\n              to its source code.\nC-h v RET --> Same as above with variables  \n\nC-h k RET --> Allows you to enter a key combination and displays the\n              name of the elisp function bound to it.\n\nSearching for help:\n\nC-h a --> Prompts you for a string to search for a command in the\n          help system. Similar to the 'apropos' or 'man -k'\n          commands in Unix systems.\n\nStarting a tutorial:\n\nC-h C-t --> Starts a tutorial designed to familiarize you with\n            basic Emacs functionality.\n```\n\n# Emacs \"Killer Apps\"\n\nAs I hinted above, Emacs functionality goes way beyond being a mere text\neditor. I will list here a couple of Emacs \"apps\" that are fairly\npowerful and popular and may interest you in and of themselves.\n\n## Org\n\nTechnnically, `org-mode`, a major mode for buffer editing that provides\norganizational tools. It is very difficult to succinctly describe what\nOrg can do because it's a behemoth of a tool that has many diverse uses\nto different people. I will attempt to describe the main features I use\nbriefly.\n\n  - Divide your file into sections and sub-sections for easy outlining\n    and organizing of concepts.\n  - Different headings in the outline are foldable/expandable so that\n    you can focus on what you need to focus on and eliminate\n    distractions.\n  - You can maintain a TODO list within Org\n  - You can compile TODO lists from many files into an agenda\n  - Track the time you spend on each TODO task\n  - Manage tables in plain text (including spreadsheet-like\n    capabilities)\n  - Using the extension `org-babel`, write and execute code blocks in\n    your file. The results are captured and are re-usable within the\n    file itself. Think Jupyter notebook for any language.\n  - Display inline images and LaTeX formulas as images within your file\n    (makes for a great note-taking system and/or personal wiki)\n  - Export your file into many different formats (LaTeX, PDF, html,…)\n\nOrg mode is a very powerful tool to add to your productivity arsenal\nand, on a personal note, was the reason that caused me to start using\nEmacs after years of using Vim.\n\n## Magit\n\nThis is a frontend to `git` from within Emacs. It features a very\nintuitive and discoverable interface, yet exposes very powerful\nfunctionality that allows you to manage commits at the chunk level,\ninspect diffs, rebase, cherry-pick, … all from within the comfort of\nyour own editor.\n\n# A Word of Advice\n\nIf you are considering using Emacs, a common trap that beginning users\nfall into is to copy someone else's configuration file and use it as is.\nI highly recommend against doing this for several reasons:\n\n  - It will discourage you from learning and finding things out for\n    yourself\n  - Someone else's configuration will probably contain many things\n    relevant to them that you won't need or ever use.\n  - It defeats the purpose of having a customizable text editor that can\n    fit your own needs.\n\nWhat I encourage you to do is to look at other people's configurations\nand seek to understand them and adapt only what makes sense to you. You\ncan find out about new features of Emacs through many YouTube videos,\nscreencasts or blog posts and then learn for yourself how to add them to\nyour configuration and workflow. This way, you grow your configuration\nincrementally while increasing your knowledge of Emacs along the way.\n\n# Additional Resources\n\n  - [The GNU Emacs Manual](https://www.gnu.org/software/emacs/manual/emacs.html)\n  - [Emacs Stack Exchange](https://emacs.stackexchange.com/)\n  - [Emacs Wiki](https://www.emacswiki.org/emacs/EmacsWiki)"
