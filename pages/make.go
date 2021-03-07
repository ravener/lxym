
package pages

const Make = "A Makefile defines a graph of rules for creating a target (or targets).\nIts purpose is to do the minimum amount of work needed to update a\ntarget to the most recent version of the source. Famously written over a\nweekend by Stuart Feldman in 1976, it is still widely used (particularly\non Unix and Linux) despite many competitors and criticisms.\n\nThere are many varieties of make in existence, however this article\nassumes that we are using GNU make which is the standard on Linux.\n\n```make\n\n# Comments can be written like this.\n\n# File should be named Makefile and then can be run as `make <target>`.\n# Otherwise we use `make -f \"filename\" <target>`.\n\n# Warning - only use TABS to indent in Makefiles, never spaces!\n\n#-----------------------------------------------------------------------\n# Basics\n#-----------------------------------------------------------------------\n\n# Rules are of the format\n# target: <prerequisite>\n# where prerequisites are optional.\n\n# A rule - this rule will only run if file0.txt doesn't exist.\nfile0.txt:\n\techo \"foo\" > file0.txt\n\t# Even comments in these 'recipe' sections get passed to the shell.\n\t# Try `make file0.txt` or simply `make` - first rule is the default.\n\n# This rule will only run if file0.txt is newer than file1.txt.\nfile1.txt: file0.txt\n\tcat file0.txt > file1.txt\n\t# use the same quoting rules as in the shell.\n\t@cat file0.txt >> file1.txt\n\t# @ stops the command from being echoed to stdout.\n\t-@echo 'hello'\n\t# - means that make will keep going in the case of an error.\n\t# Try `make file1.txt` on the commandline.\n\n# A rule can have multiple targets and multiple prerequisites\nfile2.txt file3.txt: file0.txt file1.txt\n\ttouch file2.txt\n\ttouch file3.txt\n\n# Make will complain about multiple recipes for the same rule. Empty\n# recipes don't count though and can be used to add new dependencies.\n\n#-----------------------------------------------------------------------\n# Phony Targets\n#-----------------------------------------------------------------------\n\n# A phony target. Any target that isn't a file.\n# It will never be up to date so make will always try to run it.\nall: maker process\n\n# We can declare things out of order.\nmaker:\n\ttouch ex0.txt ex1.txt\n\n# Can avoid phony rules breaking when a real file has the same name by\n.PHONY: all maker process\n# This is a special target. There are several others.\n\n# A rule with a dependency on a phony target will always run\nex0.txt ex1.txt: maker\n\n# Common phony targets are: all make clean install ...\n\n#-----------------------------------------------------------------------\n# Automatic Variables & Wildcards\n#-----------------------------------------------------------------------\n\nprocess: file*.txt\t#using a wildcard to match filenames\n\t@echo $^\t# $^ is a variable containing the list of prerequisites\n\t@echo $@\t# prints the target name\n\t#(for multiple target rules, $@ is whichever caused the rule to run)\n\t@echo $<\t# the first prerequisite listed\n\t@echo $?\t# only the dependencies that are out of date\n\t@echo $+\t# all dependencies including duplicates (unlike normal)\n\t#@echo $|\t# all of the 'order only' prerequisites\n\n# Even if we split up the rule dependency definitions, $^ will find them\nprocess: ex1.txt file0.txt\n# ex1.txt will be found but file0.txt will be deduplicated.\n\n#-----------------------------------------------------------------------\n# Patterns\n#-----------------------------------------------------------------------\n\n# Can teach make how to convert certain files into other files.\n\n%.png: %.svg\n\tinkscape --export-png $^\n\n# Pattern rules will only do anything if make decides to create the\n# target.\n\n# Directory paths are normally ignored when matching pattern rules. But\n# make will try to use the most appropriate rule available.\nsmall/%.png: %.svg\n\tinkscape --export-png --export-dpi 30 $^\n\n# make will use the last version for a pattern rule that it finds.\n%.png: %.svg\n\t@echo this rule is chosen\n\n# However make will use the first pattern rule that can make the target\n%.png: %.ps\n\t@echo this rule is not chosen if *.svg and *.ps are both present\n\n# make already has some pattern rules built-in. For instance, it knows\n# how to turn *.c files into *.o files.\n\n# Older makefiles might use suffix rules instead of pattern rules\n.png.ps:\n\t@echo this rule is similar to a pattern rule.\n\n# Tell make about the suffix rule\n.SUFFIXES: .png\n\n#-----------------------------------------------------------------------\n# Variables\n#-----------------------------------------------------------------------\n# aka. macros\n\n# Variables are basically all string types\n\nname = Ted\nname2=\"Sarah\"\n\necho:\n\t@echo $(name)\n\t@echo ${name2}\n\t@echo $name    # This won't work, treated as $(n)ame.\n\t@echo $(name3) # Unknown variables are treated as empty strings.\n\n# There are 4 places to set variables.\n# In order of priority from highest to lowest:\n# 1: commandline arguments\n# 2: Makefile\n# 3: shell environment variables - make imports these automatically.\n# 4: make has some predefined variables\n\nname4 ?= Jean\n# Only set the variable if environment variable is not already defined.\n\noverride name5 = David\n# Stops commandline arguments from changing this variable.\n\nname4 +=grey\n# Append values to variable (includes a space).\n\n# Pattern-specific variable values (GNU extension).\necho: name2 = Sara # True within the matching rule\n\t# and also within its remade recursive dependencies\n\t# (except it can break when your graph gets too complicated!)\n\n# Some variables defined automatically by make.\necho_inbuilt:\n\techo $(CC)\n\techo ${CXX}\n\techo $(FC)\n\techo ${CFLAGS}\n\techo $(CPPFLAGS)\n\techo ${CXXFLAGS}\n\techo $(LDFLAGS)\n\techo ${LDLIBS}\n\n#-----------------------------------------------------------------------\n# Variables 2\n#-----------------------------------------------------------------------\n\n# The first type of variables are evaluated each time they are used.\n# This can be expensive, so a second type of variable exists which is\n# only evaluated once. (This is a GNU make extension)\n\nvar := hello\nvar2 ::=  $(var) hello\n#:= and ::= are equivalent.\n\n# These variables are evaluated procedurally (in the order that they\n# appear), thus breaking with the rest of the language !\n\n# This doesn't work\nvar3 ::= $(var4) and good luck\nvar4 ::= good night\n\n#-----------------------------------------------------------------------\n# Functions\n#-----------------------------------------------------------------------\n\n# make has lots of functions available.\n\nsourcefiles = $(wildcard *.c */*.c)\nobjectfiles = $(patsubst %.c,%.o,$(sourcefiles))\n\n# Format is $(func arg0,arg1,arg2...)\n\n# Some examples\nls:\t* src/*\n\t@echo $(filter %.txt, $^)\n\t@echo $(notdir $^)\n\t@echo $(join $(dir $^),$(notdir $^))\n\n#-----------------------------------------------------------------------\n# Directives\n#-----------------------------------------------------------------------\n\n# Include other makefiles, useful for platform specific code\ninclude foo.mk\n\nsport = tennis\n# Conditional compilation\nreport:\nifeq ($(sport),tennis)\n\t@echo 'game, set, match'\nelse\n\t@echo \"They think it's all over; it is now\"\nendif\n\n# There are also ifneq, ifdef, ifndef\n\nfoo = true\n\nifdef $(foo)\nbar = 'hello'\nendif\n```\n\n### More Resources\n\n+ [gnu make documentation](https://www.gnu.org/software/make/manual/)\n+ [software carpentry tutorial](http://swcarpentry.github.io/make-novice/)\n+ learn C the hard way [ex2](http://c.learncodethehardway.org/book/ex2.html) [ex28](http://c.learncodethehardway.org/book/ex28.html)"
