
package pages

const R = "R is a statistical computing language. It has lots of libraries for uploading and cleaning data sets, running statistical procedures, and making graphs. You can also run `R` commands within a LaTeX document.\n\n```r\n\n# Comments start with number symbols.\n\n# You can't make multi-line comments,\n# but you can stack multiple comments like so.\n\n# in Windows you can use CTRL-ENTER to execute a line.\n# on Mac it is COMMAND-ENTER\n\n\n\n#############################################################################\n# Stuff you can do without understanding anything about programming\n#############################################################################\n\n# In this section, we show off some of the cool stuff you can do in\n# R without understanding anything about programming. Do not worry\n# about understanding everything the code does. Just enjoy!\n\ndata()\t        # browse pre-loaded data sets\ndata(rivers)\t# get this one: \"Lengths of Major North American Rivers\"\nls()\t        # notice that \"rivers\" now appears in the workspace\nhead(rivers)\t# peek at the data set\n# 735 320 325 392 524 450\n\nlength(rivers)\t# how many rivers were measured?\n# 141\nsummary(rivers) # what are some summary statistics?\n#   Min. 1st Qu.  Median    Mean 3rd Qu.    Max.\n#  135.0   310.0   425.0   591.2   680.0  3710.0\n\n# make a stem-and-leaf plot (a histogram-like data visualization)\nstem(rivers)\n\n#  The decimal point is 2 digit(s) to the right of the |\n#\n#   0 | 4\n#   2 | 011223334555566667778888899900001111223333344455555666688888999\n#   4 | 111222333445566779001233344567\n#   6 | 000112233578012234468\n#   8 | 045790018\n#  10 | 04507\n#  12 | 1471\n#  14 | 56\n#  16 | 7\n#  18 | 9\n#  20 |\n#  22 | 25\n#  24 | 3\n#  26 |\n#  28 |\n#  30 |\n#  32 |\n#  34 |\n#  36 | 1\n\nstem(log(rivers)) # Notice that the data are neither normal nor log-normal!\n# Take that, Bell curve fundamentalists.\n\n#  The decimal point is 1 digit(s) to the left of the |\n#\n#  48 | 1\n#  50 |\n#  52 | 15578\n#  54 | 44571222466689\n#  56 | 023334677000124455789\n#  58 | 00122366666999933445777\n#  60 | 122445567800133459\n#  62 | 112666799035\n#  64 | 00011334581257889\n#  66 | 003683579\n#  68 | 0019156\n#  70 | 079357\n#  72 | 89\n#  74 | 84\n#  76 | 56\n#  78 | 4\n#  80 |\n#  82 | 2\n\n# make a histogram:\nhist(rivers, col=\"#333333\", border=\"white\", breaks=25) # play around with these parameters\nhist(log(rivers), col=\"#333333\", border=\"white\", breaks=25) # you'll do more plotting later\n\n# Here's another neat data set that comes pre-loaded. R has tons of these.\ndata(discoveries)\nplot(discoveries, col=\"#333333\", lwd=3, xlab=\"Year\",\n     main=\"Number of important discoveries per year\")\nplot(discoveries, col=\"#333333\", lwd=3, type = \"h\", xlab=\"Year\",\n     main=\"Number of important discoveries per year\")\n\n# Rather than leaving the default ordering (by year),\n# we could also sort to see what's typical:\nsort(discoveries)\n#  [1]  0  0  0  0  0  0  0  0  0  1  1  1  1  1  1  1  1  1  1  1  1  2  2  2  2\n# [26]  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  2  3  3  3\n# [51]  3  3  3  3  3  3  3  3  3  3  3  3  3  3  3  3  3  4  4  4  4  4  4  4  4\n# [76]  4  4  4  4  5  5  5  5  5  5  5  6  6  6  6  6  6  7  7  7  7  8  9 10 12\n\nstem(discoveries, scale=2)\n#\n#  The decimal point is at the |\n#\n#   0 | 000000000\n#   1 | 000000000000\n#   2 | 00000000000000000000000000\n#   3 | 00000000000000000000\n#   4 | 000000000000\n#   5 | 0000000\n#   6 | 000000\n#   7 | 0000\n#   8 | 0\n#   9 | 0\n#  10 | 0\n#  11 |\n#  12 | 0\n\nmax(discoveries)\n# 12\nsummary(discoveries)\n#   Min. 1st Qu.  Median    Mean 3rd Qu.    Max.\n#    0.0     2.0     3.0     3.1     4.0    12.0\n\n# Roll a die a few times\nround(runif(7, min=.5, max=6.5))\n# 1 4 6 1 4 6 4\n# Your numbers will differ from mine unless we set the same random.seed(31337)\n\n# Draw from a standard Gaussian 9 times\nrnorm(9)\n# [1]  0.07528471  1.03499859  1.34809556 -0.82356087  0.61638975 -1.88757271\n# [7] -0.59975593  0.57629164  1.08455362\n\n\n\n##################################################\n# Data types and basic arithmetic\n##################################################\n\n# Now for the programming-oriented part of the tutorial.\n# In this section you will meet the important data types of R:\n# integers, numerics, characters, logicals, and factors.\n# There are others, but these are the bare minimum you need to\n# get started.\n\n# INTEGERS\n# Long-storage integers are written with L\n5L # 5\nclass(5L) # \"integer\"\n# (Try ?class for more information on the class() function.)\n# In R, every single value, like 5L, is considered a vector of length 1\nlength(5L) # 1\n# You can have an integer vector with length > 1 too:\nc(4L, 5L, 8L, 3L) # 4 5 8 3\nlength(c(4L, 5L, 8L, 3L)) # 4\nclass(c(4L, 5L, 8L, 3L)) # \"integer\"\n\n# NUMERICS\n# A \"numeric\" is a double-precision floating-point number\n5 # 5\nclass(5) # \"numeric\"\n# Again, everything in R is a vector;\n# you can make a numeric vector with more than one element\nc(3,3,3,2,2,1) # 3 3 3 2 2 1\n# You can use scientific notation too\n5e4 # 50000\n6.02e23 # Avogadro's number\n1.6e-35 # Planck length\n# You can also have infinitely large or small numbers\nclass(Inf)\t# \"numeric\"\nclass(-Inf)\t# \"numeric\"\n# You might use \"Inf\", for example, in integrate(dnorm, 3, Inf);\n# this obviates Z-score tables.\n\n# BASIC ARITHMETIC\n# You can do arithmetic with numbers\n# Doing arithmetic on a mix of integers and numerics gives you another numeric\n10L + 66L # 76      # integer plus integer gives integer\n53.2 - 4  # 49.2    # numeric minus numeric gives numeric\n2.0 * 2L  # 4       # numeric times integer gives numeric\n3L / 4    # 0.75    # integer over numeric gives numeric\n3 %% 2\t  # 1       # the remainder of two numerics is another numeric\n# Illegal arithmetic yields you a \"not-a-number\":\n0 / 0 # NaN\nclass(NaN) # \"numeric\"\n# You can do arithmetic on two vectors with length greater than 1,\n# so long as the larger vector's length is an integer multiple of the smaller\nc(1,2,3) + c(1,2,3) # 2 4 6\n# Since a single number is a vector of length one, scalars are applied \n# elementwise to vectors\n(4 * c(1,2,3) - 2) / 2 # 1 3 5\n# Except for scalars, use caution when performing arithmetic on vectors with \n# different lengths. Although it can be done, \nc(1,2,3,1,2,3) * c(1,2) # 1 4 3 2 2 6\n# Matching lengths is better practice and easier to read\nc(1,2,3,1,2,3) * c(1,2,1,2,1,2) \n\n# CHARACTERS\n# There's no difference between strings and characters in R\n\"Horatio\" # \"Horatio\"\nclass(\"Horatio\") # \"character\"\nclass('H') # \"character\"\n# Those were both character vectors of length 1\n# Here is a longer one:\nc('alef', 'bet', 'gimmel', 'dalet', 'he')\n# =>\n# \"alef\"   \"bet\"    \"gimmel\" \"dalet\"  \"he\"\nlength(c(\"Call\",\"me\",\"Ishmael\")) # 3\n# You can do regex operations on character vectors:\nsubstr(\"Fortuna multis dat nimis, nulli satis.\", 9, 15) # \"multis \"\ngsub('u', 'ø', \"Fortuna multis dat nimis, nulli satis.\") # \"Fortøna møltis dat nimis, nølli satis.\"\n# R has several built-in character vectors:\nletters\n# =>\n#  [1] \"a\" \"b\" \"c\" \"d\" \"e\" \"f\" \"g\" \"h\" \"i\" \"j\" \"k\" \"l\" \"m\" \"n\" \"o\" \"p\" \"q\" \"r\" \"s\"\n# [20] \"t\" \"u\" \"v\" \"w\" \"x\" \"y\" \"z\"\nmonth.abb # \"Jan\" \"Feb\" \"Mar\" \"Apr\" \"May\" \"Jun\" \"Jul\" \"Aug\" \"Sep\" \"Oct\" \"Nov\" \"Dec\"\n\n# LOGICALS\n# In R, a \"logical\" is a boolean\nclass(TRUE)\t# \"logical\"\nclass(FALSE)\t# \"logical\"\n# Their behavior is normal\nTRUE == TRUE\t# TRUE\nTRUE == FALSE\t# FALSE\nFALSE != FALSE\t# FALSE\nFALSE != TRUE\t# TRUE\n# Missing data (NA) is logical, too\nclass(NA)\t# \"logical\"\n# Use | and & for logic operations.\n# OR\nTRUE | FALSE\t# TRUE\n# AND\nTRUE & FALSE\t# FALSE\n# Applying | and & to vectors returns elementwise logic operations\nc(TRUE,FALSE,FALSE) | c(FALSE,TRUE,FALSE) # TRUE TRUE FALSE\nc(TRUE,FALSE,TRUE) & c(FALSE,TRUE,TRUE) # FALSE FALSE TRUE\n# You can test if x is TRUE\nisTRUE(TRUE)\t# TRUE\n# Here we get a logical vector with many elements:\nc('Z', 'o', 'r', 'r', 'o') == \"Zorro\" # FALSE FALSE FALSE FALSE FALSE\nc('Z', 'o', 'r', 'r', 'o') == \"Z\" # TRUE FALSE FALSE FALSE FALSE\n\n# FACTORS\n# The factor class is for categorical data\n# Factors can be ordered (like childrens' grade levels) or unordered (like colors)\nfactor(c(\"blue\", \"blue\", \"green\", NA, \"blue\"))\n#  blue blue green   <NA>   blue\n# Levels: blue green\n# The \"levels\" are the values the categorical data can take\n# Note that missing data does not enter the levels\nlevels(factor(c(\"green\", \"green\", \"blue\", NA, \"blue\"))) # \"blue\" \"green\"\n# If a factor vector has length 1, its levels will have length 1, too\nlength(factor(\"green\")) # 1\nlength(levels(factor(\"green\"))) # 1\n# Factors are commonly seen in data frames, a data structure we will cover later\ndata(infert) # \"Infertility after Spontaneous and Induced Abortion\"\nlevels(infert$education) # \"0-5yrs\"  \"6-11yrs\" \"12+ yrs\"\n\n# NULL\n# \"NULL\" is a weird one; use it to \"blank out\" a vector\nclass(NULL)\t# NULL\nparakeet = c(\"beak\", \"feathers\", \"wings\", \"eyes\")\nparakeet\n# =>\n# [1] \"beak\"     \"feathers\" \"wings\"    \"eyes\"\nparakeet <- NULL\nparakeet\n# =>\n# NULL\n\n# TYPE COERCION\n# Type-coercion is when you force a value to take on a different type\nas.character(c(6, 8)) # \"6\" \"8\"\nas.logical(c(1,0,1,1)) # TRUE FALSE  TRUE  TRUE\n# If you put elements of different types into a vector, weird coercions happen:\nc(TRUE, 4) # 1 4\nc(\"dog\", TRUE, 4) # \"dog\"  \"TRUE\" \"4\"\nas.numeric(\"Bilbo\")\n# =>\n# [1] NA\n# Warning message:\n# NAs introduced by coercion\n\n# Also note: those were just the basic data types\n# There are many more data types, such as for dates, time series, etc.\n\n\n\n##################################################\n# Variables, loops, if/else\n##################################################\n\n# A variable is like a box you store a value in for later use.\n# We call this \"assigning\" the value to the variable.\n# Having variables lets us write loops, functions, and if/else statements\n\n# VARIABLES\n# Lots of way to assign stuff:\nx = 5 # this is possible\ny <- \"1\" # this is preferred\nTRUE -> z # this works but is weird\n\n# LOOPS\n# We've got for loops\nfor (i in 1:4) {\n  print(i)\n}\n# We've got while loops\na <- 10\nwhile (a > 4) {\n\tcat(a, \"...\", sep = \"\")\n\ta <- a - 1\n}\n# Keep in mind that for and while loops run slowly in R\n# Operations on entire vectors (i.e. a whole row, a whole column)\n# or apply()-type functions (we'll discuss later) are preferred\n\n# IF/ELSE\n# Again, pretty standard\nif (4 > 3) {\n\tprint(\"4 is greater than 3\")\n} else {\n\tprint(\"4 is not greater than 3\")\n}\n# =>\n# [1] \"4 is greater than 3\"\n\n# FUNCTIONS\n# Defined like so:\njiggle <- function(x) {\n\tx = x + rnorm(1, sd=.1)\t#add in a bit of (controlled) noise\n\treturn(x)\n}\n# Called like any other R function:\njiggle(5)\t# 5±ε. After set.seed(2716057), jiggle(5)==5.005043\n\n\n\n###########################################################################\n# Data structures: Vectors, matrices, data frames, and arrays\n###########################################################################\n\n# ONE-DIMENSIONAL\n\n# Let's start from the very beginning, and with something you already know: vectors.\nvec <- c(8, 9, 10, 11)\nvec\t#  8  9 10 11\n# We ask for specific elements by subsetting with square brackets\n# (Note that R starts counting from 1)\nvec[1]\t\t# 8\nletters[18]\t# \"r\"\nLETTERS[13]\t# \"M\"\nmonth.name[9]\t# \"September\"\nc(6, 8, 7, 5, 3, 0, 9)[3]\t# 7\n# We can also search for the indices of specific components,\nwhich(vec %% 2 == 0)\t# 1 3\n# grab just the first or last few entries in the vector,\nhead(vec, 1)\t# 8\ntail(vec, 2)\t# 10 11\n# or figure out if a certain value is in the vector\nany(vec == 10) # TRUE\n# If an index \"goes over\" you'll get NA:\nvec[6]\t# NA\n# You can find the length of your vector with length()\nlength(vec)\t# 4\n# You can perform operations on entire vectors or subsets of vectors\nvec * 4\t# 16 20 24 28\nvec[2:3] * 5\t# 25 30\nany(vec[2:3] == 8) # FALSE\n# and R has many built-in functions to summarize vectors\nmean(vec)\t# 9.5\nvar(vec)\t# 1.666667\nsd(vec)\t\t# 1.290994\nmax(vec)\t# 11\nmin(vec)\t# 8\nsum(vec)\t# 38\n# Some more nice built-ins:\n5:15\t# 5  6  7  8  9 10 11 12 13 14 15\nseq(from=0, to=31337, by=1337)\n# =>\n#  [1]     0  1337  2674  4011  5348  6685  8022  9359 10696 12033 13370 14707\n# [13] 16044 17381 18718 20055 21392 22729 24066 25403 26740 28077 29414 30751\n\n# TWO-DIMENSIONAL (ALL ONE CLASS)\n\n# You can make a matrix out of entries all of the same type like so:\nmat <- matrix(nrow = 3, ncol = 2, c(1,2,3,4,5,6))\nmat\n# =>\n#      [,1] [,2]\n# [1,]    1    4\n# [2,]    2    5\n# [3,]    3    6\n# Unlike a vector, the class of a matrix is \"matrix\", no matter what's in it\nclass(mat) # => \"matrix\"\n# Ask for the first row\nmat[1,]\t# 1 4\n# Perform operation on the first column\n3 * mat[,1]\t# 3 6 9\n# Ask for a specific cell\nmat[3,2]\t# 6\n\n# Transpose the whole matrix\nt(mat)\n# =>\n#      [,1] [,2] [,3]\n# [1,]    1    2    3\n# [2,]    4    5    6\n\n# Matrix multiplication\nmat %*% t(mat)\n# =>\n#      [,1] [,2] [,3]\n# [1,]   17   22   27\n# [2,]   22   29   36\n# [3,]   27   36   45\n\n# cbind() sticks vectors together column-wise to make a matrix\nmat2 <- cbind(1:4, c(\"dog\", \"cat\", \"bird\", \"dog\"))\nmat2\n# =>\n#      [,1] [,2]\n# [1,] \"1\"  \"dog\"\n# [2,] \"2\"  \"cat\"\n# [3,] \"3\"  \"bird\"\n# [4,] \"4\"  \"dog\"\nclass(mat2)\t# matrix\n# Again, note what happened!\n# Because matrices must contain entries all of the same class,\n# everything got converted to the character class\nc(class(mat2[,1]), class(mat2[,2]))\n\n# rbind() sticks vectors together row-wise to make a matrix\nmat3 <- rbind(c(1,2,4,5), c(6,7,0,4))\nmat3\n# =>\n#      [,1] [,2] [,3] [,4]\n# [1,]    1    2    4    5\n# [2,]    6    7    0    4\n# Ah, everything of the same class. No coercions. Much better.\n\n# TWO-DIMENSIONAL (DIFFERENT CLASSES)\n\n# For columns of different types, use a data frame\n# This data structure is so useful for statistical programming,\n# a version of it was added to Python in the package \"pandas\".\n\nstudents <- data.frame(c(\"Cedric\",\"Fred\",\"George\",\"Cho\",\"Draco\",\"Ginny\"),\n                       c(3,2,2,1,0,-1),\n                       c(\"H\", \"G\", \"G\", \"R\", \"S\", \"G\"))\nnames(students) <- c(\"name\", \"year\", \"house\") # name the columns\nclass(students)\t# \"data.frame\"\nstudents\n# =>\n#     name year house\n# 1 Cedric    3     H\n# 2   Fred    2     G\n# 3 George    2     G\n# 4    Cho    1     R\n# 5  Draco    0     S\n# 6  Ginny   -1     G\nclass(students$year)\t# \"numeric\"\nclass(students[,3])\t# \"factor\"\n# find the dimensions\nnrow(students)\t# 6\nncol(students)\t# 3\ndim(students)\t# 6 3\n# The data.frame() function converts character vectors to factor vectors\n# by default; turn this off by setting stringsAsFactors = FALSE when\n# you create the data.frame\n?data.frame\n\n# There are many twisty ways to subset data frames, all subtly unalike\nstudents$year\t# 3  2  2  1  0 -1\nstudents[,2]\t# 3  2  2  1  0 -1\nstudents[,\"year\"]\t# 3  2  2  1  0 -1\n\n# An augmented version of the data.frame structure is the data.table\n# If you're working with huge or panel data, or need to merge a few data\n# sets, data.table can be a good choice. Here's a whirlwind tour:\ninstall.packages(\"data.table\") # download the package from CRAN\nrequire(data.table) # load it\nstudents <- as.data.table(students)\nstudents # note the slightly different print-out\n# =>\n#      name year house\n# 1: Cedric    3     H\n# 2:   Fred    2     G\n# 3: George    2     G\n# 4:    Cho    1     R\n# 5:  Draco    0     S\n# 6:  Ginny   -1     G\nstudents[name==\"Ginny\"] # get rows with name == \"Ginny\"\n# =>\n#     name year house\n# 1: Ginny   -1     G\nstudents[year==2] # get rows with year == 2\n# =>\n#      name year house\n# 1:   Fred    2     G\n# 2: George    2     G\n# data.table makes merging two data sets easy\n# let's make another data.table to merge with students\nfounders <- data.table(house=c(\"G\",\"H\",\"R\",\"S\"),\n                       founder=c(\"Godric\",\"Helga\",\"Rowena\",\"Salazar\"))\nfounders\n# =>\n#    house founder\n# 1:     G  Godric\n# 2:     H   Helga\n# 3:     R  Rowena\n# 4:     S Salazar\nsetkey(students, house)\nsetkey(founders, house)\nstudents <- founders[students] # merge the two data sets by matching \"house\"\nsetnames(students, c(\"house\",\"houseFounderName\",\"studentName\",\"year\"))\nstudents[,order(c(\"name\",\"year\",\"house\",\"houseFounderName\")), with=F]\n# =>\n#    studentName year house houseFounderName\n# 1:        Fred    2     G           Godric\n# 2:      George    2     G           Godric\n# 3:       Ginny   -1     G           Godric\n# 4:      Cedric    3     H            Helga\n# 5:         Cho    1     R           Rowena\n# 6:       Draco    0     S          Salazar\n\n# data.table makes summary tables easy\nstudents[,sum(year),by=house]\n# =>\n#    house V1\n# 1:     G  3\n# 2:     H  3\n# 3:     R  1\n# 4:     S  0\n\n# To drop a column from a data.frame or data.table,\n# assign it the NULL value\nstudents$houseFounderName <- NULL\nstudents\n# =>\n#    studentName year house\n# 1:        Fred    2     G\n# 2:      George    2     G\n# 3:       Ginny   -1     G\n# 4:      Cedric    3     H\n# 5:         Cho    1     R\n# 6:       Draco    0     S\n\n# Drop a row by subsetting\n# Using data.table:\nstudents[studentName != \"Draco\"]\n# =>\n#    house studentName year\n# 1:     G        Fred    2\n# 2:     G      George    2\n# 3:     G       Ginny   -1\n# 4:     H      Cedric    3\n# 5:     R         Cho    1\n# Using data.frame:\nstudents <- as.data.frame(students)\nstudents[students$house != \"G\",]\n# =>\n#   house houseFounderName studentName year\n# 4     H            Helga      Cedric    3\n# 5     R           Rowena         Cho    1\n# 6     S          Salazar       Draco    0\n\n# MULTI-DIMENSIONAL (ALL ELEMENTS OF ONE TYPE)\n\n# Arrays creates n-dimensional tables\n# All elements must be of the same type\n# You can make a two-dimensional table (sort of like a matrix)\narray(c(c(1,2,4,5),c(8,9,3,6)), dim=c(2,4))\n# =>\n#      [,1] [,2] [,3] [,4]\n# [1,]    1    4    8    3\n# [2,]    2    5    9    6\n# You can use array to make three-dimensional matrices too\narray(c(c(c(2,300,4),c(8,9,0)),c(c(5,60,0),c(66,7,847))), dim=c(3,2,2))\n# =>\n# , , 1\n#\n#      [,1] [,2]\n# [1,]    2    8\n# [2,]  300    9\n# [3,]    4    0\n#\n# , , 2\n#\n#      [,1] [,2]\n# [1,]    5   66\n# [2,]   60    7\n# [3,]    0  847\n\n# LISTS (MULTI-DIMENSIONAL, POSSIBLY RAGGED, OF DIFFERENT TYPES)\n\n# Finally, R has lists (of vectors)\nlist1 <- list(time = 1:40)\nlist1$price = c(rnorm(40,.5*list1$time,4)) # random\nlist1\n# You can get items in the list like so\nlist1$time # one way\nlist1[[\"time\"]] # another way\nlist1[[1]] # yet another way\n# =>\n#  [1]  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33\n# [34] 34 35 36 37 38 39 40\n# You can subset list items like any other vector\nlist1$price[4]\n\n# Lists are not the most efficient data structure to work with in R;\n# unless you have a very good reason, you should stick to data.frames\n# Lists are often returned by functions that perform linear regressions\n\n##################################################\n# The apply() family of functions\n##################################################\n\n# Remember mat?\nmat\n# =>\n#      [,1] [,2]\n# [1,]    1    4\n# [2,]    2    5\n# [3,]    3    6\n# Use apply(X, MARGIN, FUN) to apply function FUN to a matrix X\n# over rows (MAR = 1) or columns (MAR = 2)\n# That is, R does FUN to each row (or column) of X, much faster than a\n# for or while loop would do\napply(mat, MAR = 2, jiggle)\n# =>\n#      [,1] [,2]\n# [1,]    3   15\n# [2,]    7   19\n# [3,]   11   23\n# Other functions: ?lapply, ?sapply\n\n# Don't feel too intimidated; everyone agrees they are rather confusing\n\n# The plyr package aims to replace (and improve upon!) the *apply() family.\ninstall.packages(\"plyr\")\nrequire(plyr)\n?plyr\n\n\n\n#########################\n# Loading data\n#########################\n\n# \"pets.csv\" is a file on the internet\n# (but it could just as easily be a file on your own computer)\nrequire(RCurl)\npets <- read.csv(textConnection(getURL(\"https://learnxinyminutes.com/docs/pets.csv\")))\npets\nhead(pets, 2) # first two rows\ntail(pets, 1) # last row\n\n# To save a data frame or matrix as a .csv file\nwrite.csv(pets, \"pets2.csv\") # to make a new .csv file\n# set working directory with setwd(), look it up with getwd()\n\n# Try ?read.csv and ?write.csv for more information\n\n\n\n#########################\n# Statistical Analysis\n#########################\n\n# Linear regression!\nlinearModel <- lm(price  ~ time, data = list1)\nlinearModel # outputs result of regression\n# =>\n# Call:\n# lm(formula = price ~ time, data = list1)\n# \n# Coefficients:\n# (Intercept)         time  \n#      0.1453       0.4943  \nsummary(linearModel) # more verbose output from the regression\n# =>\n# Call:\n# lm(formula = price ~ time, data = list1)\n#\n# Residuals:\n#     Min      1Q  Median      3Q     Max \n# -8.3134 -3.0131 -0.3606  2.8016 10.3992 \n#\n# Coefficients:\n#             Estimate Std. Error t value Pr(>|t|)    \n# (Intercept)  0.14527    1.50084   0.097    0.923    \n# time         0.49435    0.06379   7.749 2.44e-09 ***\n# ---\n# Signif. codes:  0 ‘***’ 0.001 ‘**’ 0.01 ‘*’ 0.05 ‘.’ 0.1 ‘ ’ 1\n#\n# Residual standard error: 4.657 on 38 degrees of freedom\n# Multiple R-squared:  0.6124,\tAdjusted R-squared:  0.6022 \n# F-statistic: 60.05 on 1 and 38 DF,  p-value: 2.44e-09\ncoef(linearModel) # extract estimated parameters\n# =>\n# (Intercept)        time \n#   0.1452662   0.4943490 \nsummary(linearModel)$coefficients # another way to extract results\n# =>\n#              Estimate Std. Error    t value     Pr(>|t|)\n# (Intercept) 0.1452662 1.50084246 0.09678975 9.234021e-01\n# time        0.4943490 0.06379348 7.74920901 2.440008e-09\nsummary(linearModel)$coefficients[,4] # the p-values \n# =>\n#  (Intercept)         time \n# 9.234021e-01 2.440008e-09 \n\n# GENERAL LINEAR MODELS\n# Logistic regression\nset.seed(1)\nlist1$success = rbinom(length(list1$time), 1, .5) # random binary\nglModel <- glm(success  ~ time, data = list1, \n\tfamily=binomial(link=\"logit\"))\nglModel # outputs result of logistic regression\n# =>\n# Call:  glm(formula = success ~ time, \n#\tfamily = binomial(link = \"logit\"), data = list1)\n#\n# Coefficients:\n# (Intercept)         time  \n#     0.17018     -0.01321  \n# \n# Degrees of Freedom: 39 Total (i.e. Null);  38 Residual\n# Null Deviance:\t    55.35 \n# Residual Deviance: 55.12 \t AIC: 59.12\nsummary(glModel) # more verbose output from the regression\n# =>\n# Call:\n# glm(formula = success ~ time, \n#\tfamily = binomial(link = \"logit\"), data = list1)\n\n# Deviance Residuals: \n#    Min      1Q  Median      3Q     Max  \n# -1.245  -1.118  -1.035   1.202   1.327  \n# \n# Coefficients:\n#             Estimate Std. Error z value Pr(>|z|)\n# (Intercept)  0.17018    0.64621   0.263    0.792\n# time        -0.01321    0.02757  -0.479    0.632\n# \n# (Dispersion parameter for binomial family taken to be 1)\n#\n#     Null deviance: 55.352  on 39  degrees of freedom\n# Residual deviance: 55.121  on 38  degrees of freedom\n# AIC: 59.121\n# \n# Number of Fisher Scoring iterations: 3\n\n\n#########################\n# Plots\n#########################\n\n# BUILT-IN PLOTTING FUNCTIONS\n# Scatterplots!\nplot(list1$time, list1$price, main = \"fake data\")\n# Plot regression line on existing plot\nabline(linearModel, col = \"red\")\n# Get a variety of nice diagnostics\nplot(linearModel)\n# Histograms!\nhist(rpois(n = 10000, lambda = 5), col = \"thistle\")\n# Barplots!\nbarplot(c(1,4,5,1,2), names.arg = c(\"red\",\"blue\",\"purple\",\"green\",\"yellow\"))\n\n# GGPLOT2\n# But these are not even the prettiest of R's plots\n# Try the ggplot2 package for more and better graphics\ninstall.packages(\"ggplot2\")\nrequire(ggplot2)\n?ggplot2\npp <- ggplot(students, aes(x=house))\npp + geom_bar()\nll <- as.data.table(list1)\npp <- ggplot(ll, aes(x=time,price))\npp + geom_point()\n# ggplot2 has excellent documentation (available http://docs.ggplot2.org/current/)\n\n\n\n```\n\n## How do I get R?\n\n* Get R and the R GUI from [http://www.r-project.org/](http://www.r-project.org/)\n* [RStudio](http://www.rstudio.com/ide/) is another GUI"