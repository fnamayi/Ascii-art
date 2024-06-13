ASCII-ART
```
/\ / | / | | | | | /\ | \ | | / \ | ( | | | | | | / \ | |) | | |
/ /\ \ _ \ | | | | | | || / /\ \ | _ / | |
/ \ ) | | | | | | |_ / \ | | \ \ | |
// _\ |/ _| || || // _\ || _\ |_|
 ## About
This is a program that takes a string and outputs it in an art form out of ASCII characters.This project This Project was developed during Module phase at zone01 Kisumu cohort
## Installation
This is achieved by the folloowing command:
  git clone https://learn.zone01kisumu.ke/git/enungo/ascii-art.git
cd ascii-art
 ## Functions
### 1.InputArgs
This fuction exists in the input.go file and mainly it's concerned with formatting arguments from os.Args### 2.FileChoice
This is a function that uses the os.Args as an input slice then uses it to determine which file it should - <b>-sh:</b> this flag is used to read the shadow.txt file which has its own part of printing characters
- <b>-t </b>: this flag is used to read the thinkertoy.txt file which has its own style of printing the- Absence of a flag basically means the default file to be read is the standard.txt file which has a moreof printing charcters in the terminal
 The flags are actually restricted to be the third argument which is an optional one
check the bottom of Usage to see more examples
### 3.PrintWords
This is actually a function that solely does the printing of the characters it looks up from the file read### 4.Start
This functions just helps to locate where the character is located(position)in the file read and it helps### 5.ValidChar
This is a function that solely limits only characters in the ASCII read from the file to be printed and ## Usage
go run . <input-string> <flag--optional>

student$ go run . “” | cat -e student$ go run . “\n” | cat -e $ student$ go run . “Hello\n” | cat -e _ _ _ _ $ | | | | | | | | $ | || | | | | | $ | | / _ \ | | | |
/ _ \ $ | | | | | / | | | | | () | $ || |_| _| || || _/ $ $ $ $ student$ go run . “hello” | cat -e _ _ _ $ | | | | | | $ | | | | | | $ | _ \ / _ \ | | | | / _ \ $ | | | | | / | | |
| | () | $ || |_| _| || || _/ $ $ $ student$ go run . “HeLlO” | cat -e _ _ _ _ $ | | | | | | | | / \ $ | || | | | | | | | | | $ | | / _ \ | | | | | | | | $ | | | | | / | | | | | ||
| $ || || _| || || _/ $ $ $ student$ go run . “Hello There” | cat -e _ _ _ _ _ $ | | | | | | | | | | | | $ | || | | | | | | | | | _ $ | | / _ \ | | | | / _ \ | | | _ \ / _ \ | ‘| /
_ \ $ | | | | | / | | | | | () | | | | | | | | / | | | / $ || || ___| || || ___/ || || || _| |_| _| $ $ $ student$ go run . “1Hello 2There” | cat -e _ _ _ _ _ $ _ | | | | | |
| | | | | | $ / | | || | | | | | | \ | | | | _ $ | | | | / _ \ | | | | / _ \ ) | | | | _ \ / _ \ | ’| / _ \ $ | | | | | | | / | | | | | (_) | / / | | | | | | | / | | | / $ || || || ___| || || _/ || || ||
|_| _| || _| $ $ $ student$ go run . “{Hello There}” | cat -e _ _ _ _ _ $ / / | | | | | | | | | | | | \ \ $ | | | || | | | | | | | | | _ | | $ / / | | / _ \ | | | | / _ \ | |
| _ \ / _ \ | ‘| / _ \ \ \ $ \ \ | | | | | / | | | | | (_) | | | | | | | | / | | | / / / $ | | || || _| || || _/ || || |_| _| || ___| | | $ _\ // $ $ student$ go run .
“Hello\nThere” | cat -e _ _ _ _ $ | | | | | | | | $ | || | | | | | $ | | / _ \ | | | | / _ \ $ | | | | | / | | | | | () | $ || |_| _| || || _/ $ $ $ _ $ | | | | $ | | | | _ $ | | | _ \ /
_ \ | ’| / _ \ $ | | | | | | | / | | | __/ $ || || || _| |_| _| $ $ $ student$ go run . “Hello\n\nThere” | cat -e _ _ _ _ $ | | | | | | | | $ | || | | | | | $ | | / _ \ | | | |
/ _ \ $ | | | | | / | | | | | () | $ || |_| _| || || _/ $ $ $ $ _ $ | | | | $ | | | | _ $ | | | _ \ / _ \ | ‘| / _ \ $ | | | | | | | / | | | __/ $ || || || _| |_| _| $ $ $ student$
student$ go run . “Hello\n” -t $ o o o–o o o o-o $ | | | | | o o $ O–O O-o | | | | $ | | | | | o o $ o o o–o O—o O—o o-o $ $ $
```
handles an input with numbers, letters, spaces, special characters and \n.This project is developed using the Go language.
as its input thus handling special characters,ensuring the user inputs the correct number of arguments expected by the program and eventually returning a string slice and also a second return value which checks the length of the arguements.The function has its associated error handling in each case of its requirements
read from ,depending on presence of specific flags or absence of the same.There are two flags used in this program and a default one as follows:
same ascii characters in the file
familiar 'font' apppearance style/
already.It prints the 8 rows that form the characters of a word printing each line of individual charcters of the word first before going to the following lines until the 8th one
the PrintWords function locate charcters to be printed
nothing outside the scope
Contributors
Nungo Edwin
Namayi Franklyn
