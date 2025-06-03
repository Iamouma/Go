 FUNCTIONS    

DESCRIPTION

Print

Print simply prints whatever input it receives as it is on theoutput console screen, beginning from the current cursor positionwithout appending any space or newlines unless explicitly coded.Apart from printing, it returns two values: the no of bytes writtenand error message in case any write error is encountered.

Printf

Printf formats the input string as the user's choice and then printsthe formatted string onto the output console beginning from thecurrent cursor position without appending any space or newlinesunless explicitly coded. Apart from printing, it returns two values:the number of bytes written and error message in case any erroris encountered.

Println

Println works the same as print function except that it appends newline at the end of the input string and so whatever the output is, in the end, the cursor will move to the next line. Also in case, any variables are added to the input string then this function will ensure that the variables are separated by a space in between. Apart from printing, it returns two values: the no of bytes written and error message in case any write error is encountered.

Sprint

Sprint functions the same way as Print does. The only differenceis that Sprint returns the input string instead of printingon the output console.

Sprintf

Sprintf functions the same way as Printf does. The only differenceis that Sprintf returns the formatted input string insteadof printing on the output console.
Sprintln

Sprintln functions the same way as Println does. the only difference is that Sprintln returns the input stringinstead of printing on the output console.

Fprint

Unlike other print functions, the Fprint does not read or print anything on the console. Fprint formats the input string as per the default formatting and writes the formatted input string into the input file. Whenever two variables are encountered, space is automatically added in between by fprint. Apart from writing into the file, fprint returns two values: no. of bytes written& error message (if any error occurs).

Fprintf

Just similar to fprint but one difference that draws a line between the two is that fprintf formats according to the specified format and does not format according to the default formatting.Fprintf then writes the formatted input string into the file using. Apart from writing into the file, fprintf returns two values: no. of bytes are written and an error message (if any error occurs).

Fprintln

Works exactly the same as fprint. One extra point in Fprintln is thata newline is appended. The formatted input string will then bewritten into the mentioned file with the help of a writer. Apartfrom writing into the file, fprintln returns two values: no. of byteswritten and an error message (if any error occurs).
Scan

Scan collects input from the standard console, and stores this inputin successive arguments. Values that are separated by space ornewlines are treated as multiple values. These are stored inmultiple arguments. Apart from the on-screen scanning work,Scan also returns two values: no. of bytes read from consoleand an error message (if any).
Scanf

Just like the scanf in C language, where a format specifier ismentioned and then the address of the variable where theinput value is to be stored, Scanf in Go scans text read fromstandard input, and stores it in the arguments as per the specifiedformat. Apart from the on-screen scanning work, Scan alsoreturns two values: no. of bytes read from console andan error message (if any).
Scanln

Scanln works similar to Scan, but scanln stops scanning at anewline. Meaning that the last input character should befollowed by a newline for Scanln to stop scanning. Itmay as well be an End of file, no content no scan!
Sscan

Sscan works similar to scan except the difference that sscancollects input as the argument string and not the inputfrom the console screen in default formatting. And just likethe others it returns the two values: no of bytes & error message (if any).
Sscanf

Sscanf works similar to scanf except the difference that sscanfcollects input as the argument string and not the input fromthe console screen in mentioned formatting.
Sscanln

Sscanln is similar to Sscan but sscanln stops scanning at anewline. Meaning that the last input character should befollowed by a newline for Sscanln to stop scanning. Itmay as well be an End of file, no content no scan!
Fscan

Unlike other scan functions which collect input from useror the program console, in default formatting, Fscanreads content or text from the input file and returns theno of items parsed. It reads character wise, progressdirection depends on the pointer location you provide.
Fscanf

Fscanf works similar to fscan. The only difference is that fscanf reads from reader in the specified format rather than following the default formatting. The match-case is sensitive to whole words, spaces and newlines and hence, the argument string and text string must match entirely. Fscanf, apart from reading the characters, also returns the no of items parsed.
Fscanln

Fscanln is similar to Fscan but  Fscanln stops scanning at a newline. Meaning that the last input character should be followed by a newline for Fscanln to stop scanning. Itmay as well be an End of file, no content no scan!
Errorf

Errorf formats according to the mentioned format specifier and assigns this formatted string to a variable. This is the error message. This function allows you to printcustomized error messages, as per the user's will, and print it to console as an error message.
Apart from the functions, Go has some in-built interfaces in its packages. The "fmt" package consists of the following interfaces. The following description has been referred from the official site of Go, you may click here to find out more.

        TYPES        

DESCRIPTION

Formatter

Formatter is an interface that makes custom formattingpossible. Calling format would be something likePrintf(f), Fprintf(f), Sprintf(f) etc.. to generate a customformatted output on the console.
GoStringer

Any code-value that has a GoString method directlyimplements the GoStringer interface implicitly.
ScanState

The ScanState interface holds declaration of valuablefunctions like:ReadRune, UnreadRune, SkipSpace, Token, Width & Read.Basically they tell more about the state of the custom scanner.
Scanner

All the scanning functions that we use - Scan, Scanfand Scanln etc... whenever used make a call to thescanner interface. The scanner interface makes it possiblefor us to collect input through console and/orinput argument strings.
State

It represents the state of the printer that's passedto specific formats. It also holds informationabout flags and available options for the argumentstring's format specifier.
Stringer

Any value that holds a string method (that makes the conversion of string argument into output possible)and uses the default formatting, directly implements the Stringer interface. For instance: Print, Sprint, etc.
