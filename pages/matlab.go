
package pages

const Matlab = "MATLAB stands for MATrix LABoratory. It is a powerful numerical computing language commonly used in engineering and mathematics.\n\nIf you have any feedback please feel free to reach me at\n[@the_ozzinator](https://twitter.com/the_ozzinator), or\n[osvaldo.t.mendoza@gmail.com](mailto:osvaldo.t.mendoza@gmail.com).\n\n```matlab\n%% Code sections start with two percent signs. Section titles go on the same line.\n% Comments start with a percent sign.\n\n%{\nMulti line comments look\nsomething\nlike\nthis\n%}\n\n% Two percent signs denote the start of a new code section\n% Individual code sections can be run by moving the cursor to the section followed by\n% either clicking the \"Run Section\" button\n% or     using Ctrl+Shift+Enter (Windows) or Cmd+Shift+Return (OS X)\n\n%% This is the start of a code section\n%  One way of using sections is to separate expensive but unchanging start-up code like loading data\nload myFile.mat y\n\n%% This is another code section\n%  This section can be edited and run repeatedly on its own, and is helpful for exploratory programming and demos\nA = A * 2;\nplot(A);\n\n%% Code sections are also known as code cells or cell mode (not to be confused with cell arrays)\n\n\n% commands can span multiple lines, using '...':\n a = 1 + 2 + ...\n + 4\n\n% commands can be passed to the operating system\n!ping google.com\n\nwho % Displays all variables in memory\nwhos % Displays all variables in memory, with their types\nclear % Erases all your variables from memory\nclear('A') % Erases a particular variable\nopenvar('A') % Open variable in variable editor\n\nclc % Erases the writing on your Command Window\ndiary % Toggle writing Command Window text to file\nctrl-c % Abort current computation\n\nedit('myfunction.m') % Open function/script in editor\ntype('myfunction.m') % Print the source of function/script to Command Window\n\nprofile on \t% turns on the code profiler\nprofile off \t% turns off the code profiler\nprofile viewer \t% Open profiler\n\nhelp command \t% Displays documentation for command in Command Window\ndoc command \t% Displays documentation for command in Help Window\nlookfor command % Searches for command in the first commented line of all functions\nlookfor command -all % searches for command in all functions\n\n\n% Output formatting\nformat short \t% 4 decimals in a floating number\nformat long \t% 15 decimals\nformat bank \t% only two digits after decimal point - for financial calculations\nfprintf('text') % print \"text\" to the screen\ndisp('text') \t% print \"text\" to the screen\n\n% Variables & Expressions\nmyVariable = 4 \t% Notice Workspace pane shows newly created variable\nmyVariable = 4; % Semi colon suppresses output to the Command Window\n4 + 6  \t\t% ans = 10\n8 * myVariable \t% ans = 32\n2 ^ 3 \t\t% ans = 8\na = 2; b = 3;\nc = exp(a)*sin(pi/2) % c = 7.3891\n\n% Calling functions can be done in either of two ways:\n% Standard function syntax:\nload('myFile.mat', 'y') % arguments within parentheses, separated by commas\n% Command syntax:\nload myFile.mat y \t% no parentheses, and spaces instead of commas\n% Note the lack of quote marks in command form: inputs are always passed as\n% literal text - cannot pass variable values. Also, can't receive output:\n[V,D] = eig(A);  % this has no equivalent in command form\n[~,D] = eig(A);  % if you only want D and not V\n\n\n\n% Logicals\n1 > 5 % ans = 0\n10 >= 10 % ans = 1\n3 ~= 4 % Not equal to -> ans = 1\n3 == 3 % equal to -> ans = 1\n3 > 1 && 4 > 1 % AND -> ans = 1\n3 > 1 || 4 > 1 % OR -> ans = 1\n~1 % NOT -> ans = 0\n\n% Logicals can be applied to matrices:\nA > 5\n% for each element, if condition is true, that element is 1 in returned matrix\nA( A > 5 )\n% returns a vector containing the elements in A for which condition is true\n\n% Strings\na = 'MyString'\nlength(a) % ans = 8\na(2) % ans = y\n[a,a] % ans = MyStringMyString\n\n\n% Cells\na = {'one', 'two', 'three'}\na(1) % ans = 'one' - returns a cell\nchar(a(1)) % ans = one - returns a string\n\n% Structures\nA.b = {'one','two'};\nA.c = [1 2];\nA.d.e = false;\n\n% Vectors\nx = [4 32 53 7 1]\nx(2) % ans = 32, indices in Matlab start 1, not 0\nx(2:3) % ans = 32 53\nx(2:end) % ans = 32 53 7 1\n\nx = [4; 32; 53; 7; 1] % Column vector\n\nx = [1:10] % x = 1 2 3 4 5 6 7 8 9 10\nx = [1:2:10] % Increment by 2, i.e. x = 1 3 5 7 9\n\n% Matrices\nA = [1 2 3; 4 5 6; 7 8 9]\n% Rows are separated by a semicolon; elements are separated with space or comma\n% A =\n\n%     1     2     3\n%     4     5     6\n%     7     8     9\n\nA(2,3) % ans = 6, A(row, column)\nA(6) % ans = 8\n% (implicitly concatenates columns into vector, then indexes into that)\n\n\nA(2,3) = 42 % Update row 2 col 3 with 42\n% A =\n\n%     1     2     3\n%     4     5     42\n%     7     8     9\n\nA(2:3,2:3) % Creates a new matrix from the old one\n%ans =\n\n%     5     42\n%     8     9\n\nA(:,1) % All rows in column 1\n%ans =\n\n%     1\n%     4\n%     7\n\nA(1,:) % All columns in row 1\n%ans =\n\n%     1     2     3\n\n[A ; A] % Concatenation of matrices (vertically)\n%ans =\n\n%     1     2     3\n%     4     5    42\n%     7     8     9\n%     1     2     3\n%     4     5    42\n%     7     8     9\n\n% this is the same as\nvertcat(A,A);\n\n\n[A , A] % Concatenation of matrices (horizontally)\n\n%ans =\n\n%     1     2     3     1     2     3\n%     4     5    42     4     5    42\n%     7     8     9     7     8     9\n\n% this is the same as\nhorzcat(A,A);\n\n\nA(:, [3 1 2]) % Rearrange the columns of original matrix\n%ans =\n\n%     3     1     2\n%    42     4     5\n%     9     7     8\n\nsize(A) % ans = 3 3\n\nA(1, :) =[] % Delete the first row of the matrix\nA(:, 1) =[] % Delete the first column of the matrix\n\ntranspose(A) % Transpose the matrix, which is the same as:\nA.' % Concise version of transpose (without taking complex conjugate)\nctranspose(A) % Hermitian transpose the matrix, which is the same as:\nA'  % Concise version of complex transpose\n    % (the transpose, followed by taking complex conjugate of each element)\n\n\n\n\n\n% Element by Element Arithmetic vs. Matrix Arithmetic\n% On their own, the arithmetic operators act on whole matrices. When preceded\n% by a period, they act on each element instead. For example:\nA * B % Matrix multiplication\nA .* B % Multiply each element in A by its corresponding element in B\n\n% There are several pairs of functions, where one acts on each element, and\n% the other (whose name ends in m) acts on the whole matrix.\nexp(A) % exponentiate each element\nexpm(A) % calculate the matrix exponential\nsqrt(A) % take the square root of each element\nsqrtm(A) %  find the matrix whose square is A\n\n\n% Plotting\nx = 0:.10:2*pi; % Creates a vector that starts at 0 and ends at 2*pi with increments of .1\ny = sin(x);\nplot(x,y)\nxlabel('x axis')\nylabel('y axis')\ntitle('Plot of y = sin(x)')\naxis([0 2*pi -1 1]) % x range from 0 to 2*pi, y range from -1 to 1\n\nplot(x,y1,'-',x,y2,'--',x,y3,':') % For multiple functions on one plot\nlegend('Line 1 label', 'Line 2 label') % Label curves with a legend\n\n% Alternative method to plot multiple functions in one plot.\n% while 'hold' is on, commands add to existing graph rather than replacing it\nplot(x, y)\nhold on\nplot(x, z)\nhold off\n\nloglog(x, y) % A log-log plot\nsemilogx(x, y) % A plot with logarithmic x-axis\nsemilogy(x, y) % A plot with logarithmic y-axis\n\nfplot (@(x) x^2, [2,5]) % plot the function x^2 from x=2 to x=5\n\ngrid on % Show grid; turn off with 'grid off'\naxis square % Makes the current axes region square\naxis equal % Set aspect ratio so data units are the same in every direction\n\nscatter(x, y); % Scatter-plot\nhist(x); % Histogram\nstem(x); % Plot values as stems, useful for displaying discrete data\nbar(x); % Plot bar graph\n\nz = sin(x);\nplot3(x,y,z); % 3D line plot\n\npcolor(A) % Heat-map of matrix: plot as grid of rectangles, coloured by value\ncontour(A) % Contour plot of matrix\nmesh(A) % Plot as a mesh surface\n\nh = figure % Create new figure object, with handle h\nfigure(h) % Makes the figure corresponding to handle h the current figure\nclose(h) % close figure with handle h\nclose all % close all open figure windows\nclose % close current figure window\n\nshg % bring an existing graphics window forward, or create new one if needed\nclf clear % clear current figure window, and reset most figure properties\n\n% Properties can be set and changed through a figure handle.\n% You can save a handle to a figure when you create it.\n% The function get returns a handle to the current figure\nh = plot(x, y); % you can save a handle to a figure when you create it\nset(h, 'Color', 'r')\n% 'y' yellow; 'm' magenta, 'c' cyan, 'r' red, 'g' green, 'b' blue, 'w' white, 'k' black\nset(h, 'LineStyle', '--')\n % '--' is solid line, '---' dashed, ':' dotted, '-.' dash-dot, 'none' is no line\nget(h, 'LineStyle')\n\n\n% The function gca returns a handle to the axes for the current figure\nset(gca, 'XDir', 'reverse'); % reverse the direction of the x-axis\n\n% To create a figure that contains several axes in tiled positions, use subplot\nsubplot(2,3,1); % select the first position in a 2-by-3 grid of subplots\nplot(x1); title('First Plot') % plot something in this position\nsubplot(2,3,2); % select second position in the grid\nplot(x2); title('Second Plot') % plot something there\n\n\n% To use functions or scripts, they must be on your path or current directory\npath % display current path\naddpath /path/to/dir % add to path\nrmpath /path/to/dir % remove from path\ncd /path/to/move/into % change directory\n\n\n% Variables can be saved to .mat files\nsave('myFileName.mat') % Save the variables in your Workspace\nload('myFileName.mat') % Load saved variables into Workspace\n\n% M-file Scripts\n% A script file is an external file that contains a sequence of statements.\n% They let you avoid repeatedly typing the same code in the Command Window\n% Have .m extensions\n\n% M-file Functions\n% Like scripts, and have the same .m extension\n% But can accept input arguments and return an output\n% Also, they have their own workspace (ie. different variable scope).\n% Function name should match file name (so save this example as double_input.m).\n% 'help double_input.m' returns the comments under line beginning function\nfunction output = double_input(x)\n\t%double_input(x) returns twice the value of x\n\toutput = 2*x;\nend\ndouble_input(6) % ans = 12\n\n\n% You can also have subfunctions and nested functions.\n% Subfunctions are in the same file as the primary function, and can only be\n% called by functions in the file. Nested functions are defined within another\n% functions, and have access to both its workspace and their own workspace.\n\n% If you want to create a function without creating a new file you can use an\n% anonymous function. Useful when quickly defining a function to pass to\n% another function (eg. plot with fplot, evaluate an indefinite integral\n% with quad, find roots with fzero, or find minimum with fminsearch).\n% Example that returns the square of its input, assigned to the handle sqr:\nsqr = @(x) x.^2;\nsqr(10) % ans = 100\ndoc function_handle % find out more\n\n% User input\na = input('Enter the value: ')\n\n% Stops execution of file and gives control to the keyboard: user can examine\n% or change variables. Type 'return' to continue execution, or 'dbquit' to exit\nkeyboard\n\n% Reading in data (also xlsread/importdata/imread for excel/CSV/image files)\nfopen(filename)\n\n% Output\ndisp(a) % Print out the value of variable a\ndisp('Hello World') % Print out a string\nfprintf % Print to Command Window with more control\n\n% Conditional statements (the parentheses are optional, but good style)\nif (a > 23)\n\tdisp('Greater than 23')\nelseif (a == 23)\n\tdisp('a is 23')\nelse\n\tdisp('neither condition met')\nend\n\n% Looping\n% NB. looping over elements of a vector/matrix is slow!\n% Where possible, use functions that act on whole vector/matrix at once\nfor k = 1:5\n\tdisp(k)\nend\n\nk = 0;\nwhile (k < 5)\n\tk = k + 1;\nend\n\n% Timing code execution: 'toc' prints the time since 'tic' was called\ntic\nA = rand(1000);\nA*A*A*A*A*A*A;\ntoc\n\n% Connecting to a MySQL Database\ndbname = 'database_name';\nusername = 'root';\npassword = 'root';\ndriver = 'com.mysql.jdbc.Driver';\ndburl = ['jdbc:mysql://localhost:8889/' dbname];\njavaclasspath('mysql-connector-java-5.1.xx-bin.jar'); %xx depends on version, download available at http://dev.mysql.com/downloads/connector/j/\nconn = database(dbname, username, password, driver, dburl);\nsql = ['SELECT * from table_name where id = 22'] % Example sql statement\na = fetch(conn, sql) %a will contain your data\n\n\n% Common math functions\nsin(x)\ncos(x)\ntan(x)\nasin(x)\nacos(x)\natan(x)\nexp(x)\nsqrt(x)\nlog(x)\nlog10(x)\nabs(x) %If x is complex, returns magnitude\nmin(x)\nmax(x)\nceil(x)\nfloor(x)\nround(x)\nrem(x)\nrand % Uniformly distributed pseudorandom numbers\nrandi % Uniformly distributed pseudorandom integers\nrandn % Normally distributed pseudorandom numbers\n\n%Complex math operations\nabs(x) \t % Magnitude of complex variable x\nphase(x) % Phase (or angle) of complex variable x\nreal(x)  % Returns the real part of x (i.e returns a if x = a +jb)\nimag(x)  % Returns the imaginary part of x (i.e returns b if x = a+jb)\nconj(x)  % Returns the complex conjugate \n\n\n% Common constants\npi\nNaN\ninf\n\n% Solving matrix equations (if no solution, returns a least squares solution)\n% The \\ and / operators are equivalent to the functions mldivide and mrdivide\nx=A\\b % Solves Ax=b. Faster and more numerically accurate than using inv(A)*b.\nx=b/A % Solves xA=b\n\ninv(A) % calculate the inverse matrix\npinv(A) % calculate the pseudo-inverse\n\n% Common matrix functions\nzeros(m,n) % m x n matrix of 0's\nones(m,n) % m x n matrix of 1's\ndiag(A) % Extracts the diagonal elements of a matrix A\ndiag(x) % Construct a matrix with diagonal elements listed in x, and zeroes elsewhere\neye(m,n) % Identity matrix\nlinspace(x1, x2, n) % Return n equally spaced points, with min x1 and max x2\ninv(A) % Inverse of matrix A\ndet(A) % Determinant of A\neig(A) % Eigenvalues and eigenvectors of A\ntrace(A) % Trace of matrix - equivalent to sum(diag(A))\nisempty(A) % Tests if array is empty\nall(A) % Tests if all elements are nonzero or true\nany(A) % Tests if any elements are nonzero or true\nisequal(A, B) % Tests equality of two arrays\nnumel(A) % Number of elements in matrix\ntriu(x) % Returns the upper triangular part of x\ntril(x) % Returns the lower triangular part of x\ncross(A,B) %  Returns the cross product of the vectors A and B\ndot(A,B) % Returns scalar product of two vectors (must have the same length)\ntranspose(A) % Returns the transpose of A\nfliplr(A) % Flip matrix left to right\nflipud(A) % Flip matrix up to down\n\n% Matrix Factorisations\n[L, U, P] = lu(A) % LU decomposition: PA = LU,L is lower triangular, U is upper triangular, P is permutation matrix\n[P, D] = eig(A) % eigen-decomposition: AP = PD, P's columns are eigenvectors and D's diagonals are eigenvalues\n[U,S,V] = svd(X) % SVD: XV = US, U and V are unitary matrices, S has non-negative diagonal elements in decreasing order\n\n% Common vector functions\nmax     % largest component\nmin     % smallest component\nlength  % length of a vector\nsort    % sort in ascending order\nsum     % sum of elements\nprod    % product of elements\nmode    % modal value\nmedian  % median value\nmean    % mean value\nstd     % standard deviation\nperms(x) % list all permutations of elements of x\nfind(x) % Finds all non-zero elements of x and returns their indexes, can use comparison operators, \n        % i.e. find( x == 3 ) returns indexes of elements that are equal to 3\n        % i.e. find( x >= 3 ) returns indexes of elements greater than or equal to 3\n\n\n% Classes\n% Matlab can support object-oriented programming. \n% Classes must be put in a file of the class name with a .m extension. \n% To begin, we create a simple class to store GPS waypoints.\n% Begin WaypointClass.m\nclassdef WaypointClass % The class name.\n  properties % The properties of the class behave like Structures\n    latitude \n    longitude \n  end\n  methods \n    % This method that has the same name of the class is the constructor. \n    function obj = WaypointClass(lat, lon)\n      obj.latitude = lat;\n      obj.longitude = lon;\n    end\n\n    % Other functions that use the Waypoint object\n    function r = multiplyLatBy(obj, n)\n      r = n*[obj.latitude];\n    end\n\n    % If we want to add two Waypoint objects together without calling\n    % a special function we can overload Matlab's arithmetic like so:\n    function r = plus(o1,o2)\n      r = WaypointClass([o1.latitude] +[o2.latitude], ...\n                        [o1.longitude]+[o2.longitude]);\n    end\n  end\nend\n% End WaypointClass.m\n\n% We can create an object of the class using the constructor\na = WaypointClass(45.0, 45.0)\n\n% Class properties behave exactly like Matlab Structures.\na.latitude = 70.0\na.longitude = 25.0\n\n% Methods can be called in the same way as functions\nans = multiplyLatBy(a,3)\n\n% The method can also be called using dot notation. In this case, the object \n% does not need to be passed to the method.\nans = a.multiplyLatBy(1/3)\n\n% Matlab functions can be overloaded to handle objects. \n% In the method above, we have overloaded how Matlab handles \n% the addition of two Waypoint objects.\nb = WaypointClass(15.0, 32.0)\nc = a + b\n\n```\n\n## More on Matlab\n\n* [The official website](http://www.mathworks.com/products/matlab/)\n* [The official MATLAB Answers forum](http://www.mathworks.com/matlabcentral/answers/)\n* [Loren on the Art of MATLAB](http://blogs.mathworks.com/loren/)\n* [Cleve's Corner](http://blogs.mathworks.com/cleve/)"