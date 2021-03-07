
package pages

const Cobol = "COBOL is a business-oriented language revised multiple times since its original design in 1960. It is claimed to still be used in over 80% of \norganizations.\n\n```cobol\n      *COBOL. Coding like it's 1985. \n      *Compiles with GnuCOBOL in OpenCobolIDE 4.7.6.\n       \n      *COBOL has significant differences between legacy (COBOL-85)\n      *and modern (COBOL-2002 and COBOL-2014) versions.\n      *Legacy versions require columns 1-6 to be blank (they are used\n      *to store the index number of the punched card..)\n      *A '*' in column 7 means a comment.\n      *In legacy COBOL, a comment can only be a full line.\n      *Modern COBOL doesn't require fixed columns and uses *> for\n      *a comment, which can appear in the middle of a line.\n      *Legacy COBOL also imposes a limit on maximum line length.\n      *Keywords have to be in capitals in legacy COBOL,\n      *but are case insensitive in modern.\n      *Although modern COBOL allows you to use mixed-case characters\n      *it is still common to use all caps when writing COBOL code.\n      *This is what most professional COBOL developers do.\n      *COBOL statements end with a period.\n      \n      *COBOL code is broken up into 4 divisions.\n      *Those divisions, in order, are:\n      *IDENTIFICATION DIVISION.\n      *ENVIRONMENT DIVISION.\n      *DATA DIVISION.\n      *PROCEDURE DIVISION.\n\n      *First, we must give our program an ID.\n      *Identification division can include other values too,\n      *but they are comments only. Program-id is the only one that is mandatory.\n       IDENTIFICATION DIVISION.\n           PROGRAM-ID.    LEARN.\n           AUTHOR.        JOHN DOE.\n           DATE-WRITTEN.  05/02/2020.\n\n      *Let's declare some variables.\n      *We do this in the WORKING-STORAGE section within the DATA DIVISION.\n      *Each data item (aka variable) starts with a level number, \n      *then the name of the item, followed by a picture clause \n      *describing the type of data that the variable will contain.\n      *Almost every COBOL programmer will abbreviate PICTURE as PIC.\n      *A is for alphabetic, X is for alphanumeric, and 9 is for numeric.\n       \n      *example:\n      01  MYNAME PIC xxxxxxxxxx.    *> A 10 character string.\n       \n      *But counting all those x's can lead to errors, \n      *so the above code can, and should\n      *be re-written as:\n      01 MYNAME PIC X(10).\n       \n      *Here are some more examples:\n      01  AGE             PIC      9(3).   *> A number up to 3 digits.\n      01  LAST_NAME       PIC      X(10).  *> A string up to 10 characters.\n       \n      *In COBOL, multiple spaces are the same as a single space, so it is common\n      *to use multiple spaces to line up your code so that it is easier for other\n      *coders to read.\n      01  inyear picture s9(7). *> S makes number signed.\n                                 *> Brackets indicate 7 repeats of 9,\n                                 *> ie a 6 digit number (not an array).\n\n      *Now let's write some code. Here is a simple, Hello World program.\n      IDENTIFICATION DIVISION.\n      PROGRAM-ID. HELLO.\n      DATA DIVISION.\n      WORKING-STORAGE SECTION.\n      01 THE-MESSAGE      PIC X(20).\n      PROCEDURE DIVISION.\n          DISPLAY \"STARTING PROGRAM\".\n          MOVE \"HELLO WORLD\" TO THE-MESSAGE.\n          DISPLAY THE-MESSAGE.\n          STOP RUN.\n      \n      *The above code will output:\n      *STARTING PROGRAM\n      *HELLO WORLD\n      \n\n      \n      ********COBOL can perform math***************\n      ADD 1 TO AGE GIVING NEW-AGE.\n      SUBTRACT 1 FROM COUNT.\n      DIVIDE VAR-1 INTO VAR-2 GIVING VAR-3.\n      COMPUTE TOTAL-COUNT = COUNT1 PLUS COUNT2.\n      \n      \n      *********PERFORM********************\n      *The PERFORM keyword allows you to jump to another specified section of the code,\n      *and then to return to the next executable\n      *statement once the specified section of code is completed. \n      *You must write the full word, PERFORM, you cannot abbreviate it.\n\n      IDENTIFICATION DIVISION.\n      PROGRAM-ID. HELLOCOBOL.\n\n      PROCEDURE DIVISION.\n         FIRST-PARA.\n             DISPLAY 'THIS IS IN FIRST-PARA'.\n         PERFORM THIRD-PARA THRU FOURTH-PARA. *>skip second-para and perfrom 3rd & 4th\n         *> then after performing third and fourth,\n         *> return here and continue the program until STOP RUN.\n   \n         SECOND-PARA.\n             DISPLAY 'THIS IS IN SECOND-PARA'.\n         STOP RUN.\n   \n         THIRD-PARA.\n             DISPLAY 'THIS IS IN THIRD-PARA'.\n   \n         FOURTH-PARA.\n             DISPLAY 'THIS IS IN FOURTH-PARA'.\n   \n   \n      *When you compile and execute the above program, it produces the following result: \n          THIS IS IN FIRST-PARA\n          THIS IS IN THIRD-PARA\n          THIS IS IN FOURTH-PARA\n          THIS IS IN SECOND-PARA\n          \n          \n      **********Combining variables together using STRING ***********\n      \n      *Now it is time to learn about two related COBOL verbs: string and unstring.\n\n      *The string verb is used to concatenate, or put together, two or more stings.\n      *Unstring is used, not surprisingly, to separate a         \n      *string into two or more smaller strings. \n      *It is important that you remember to use ‘delimited by’ when you\n      *are using string or unstring in your program. \n\n      IDENTIFICATION DIVISION.\n      PROGRAM-ID. LEARNING.\n      ENVIRONMENT DIVISION.\n      DATA DIVISION.\n      WORKING-STORAGE SECTION.\n      01 FULL-NAME PIC X(20).\n      01 FIRST-NAME PIC X(13) VALUE \"BOB GIBBERISH\".\n      01 LAST-NAME PIC X(5) VALUE \"COBB\".\n      PROCEDURE DIVISION.\n          STRING FIRST-NAME DELIMITED BY SPACE\n            \" \"\n            LAST-NAME DELIMITED BY SIZE\n            INTO FULL-NAME\n          END-STRING.\n          DISPLAY \"THE FULL NAME IS: \"FULL-NAME.\n      STOP RUN.\n\n\n      *The above code will output:\n      THE FULL NAME IS: BOB COBB\n\n\n      *Let’s examine it to see why.\n\n      *First, we declared all of our variables, including the one that we are creating\n      *by the string command, in the DATA DIVISION.\n\n      *The action takes place down in the PROCEDURE DIVISION. \n      *We start with the STRING keyword and end with END-STRING. In between we         \n      *list what we want to combine together into the larger, master variable. \n      *Here, we are combining FIRST-NAME, a space, and LAST-NAME. \n\n      *The DELIMITED BY phrase that follows FIRST-NAME and \n      *LAST-NAME tells the program how much of each variable we want to capture. \n      *DELIMITED BY SPACE tells the program to start at the beginning, \n      *and capture the variable until it runs into a space. \n      *DELIMITED BY SIZE tells the program to capture the full size of the variable. \n      *Since we have DELIMITED BY SPACE after FIRST-NAME, the GIBBERISH part is ignored. \n\n      *To make this clearer, change line 10 in the above code to:\n\n      STRING FIRST-NAME DELIMITED BY SIZE\n\n      *and then re-run the program. This time the output is:\n\n      THE FULL NAME IS: BOB GIBBERISH COBB\n\n\n\n\n\n\n```\n\n##Ready For More?\n\n* [GnuCOBOL](https://sourceforge.net/projects/open-cobol/)"