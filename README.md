ASCII-ART
```
             _____    _____   _____   _____                       _____    _______        
    /\      / ____|  / ____| |_   _| |_   _|              /\     |  __ \  |__   __|       
   /  \    | (___   | |        | |     | |    ______     /  \    | |__) |    | |          
  / /\ \    \___ \  | |        | |     | |   |______|   / /\ \   |  _  /     | |          
 / ____ \   ____) | | |____   _| |_   _| |_            / ____ \  | | \ \     | |          
/_/    \_\ |_____/   \_____| |_____| |_____|          /_/    \_\ |_|  \_\    |_|          


```
## About
This Go package provides a command-line interface (CLI) tool for converting input strings into ASCII art. The tool supports different styles of ASCII art based on selected files.
This is a program that takes a string and outputs it in an art form out of ASCII characters.This projectThis Project was developed during Module phase at zone01 Kisumu cohort1. 
## Installation
This is achieved by the folloowing command:
1. Clone the repository:
 git clone https://learn.zone01kisumu.ke/git/enungo/ascii-art.git
 2. Navigate to the project directory:
 cd ascii-art
## Functions

### 1.InputArgs
This fuction exists in the input.go file and mainly it's concerned with formatting arguments from os.Args### 2.FileChoice
This is a function that uses the os.Args as an input slice then uses it to determine which file it should- <b>-sh:</b> this flag is used to read the shadow.txt file which has its own part of printing characters
- <b>-t </b>: this flag is used to read the thinkertoy.txt file which has its own style of printing the- Absence of a flag basically means the default file to be read is the standard.txt file which has a moreof printing charcters in the terminal
The flags are actually restricted to be the third argument which is an optional one
check the bottom of Usage to see more examples
### 3.PrintWords
This is actually a function that solely does the printing of the characters it looks up from the file read
### 4.Start
This functions just helps to locate where the character is located(position)in the file read and it helps
### 5.ValidChar
This is a function that solely limits only characters in the ASCII read from the file to be printed and
## Usage
#### Running the Program
To run the program, use the following command:
go run . <input-string> <flag--optional>
#### Command-Line Arguments
* Input String: A string of text to be converted into ASCII art.
* Optional Flag: A flag to specify the ASCII art style.
   * -t for Thinkertoy style.
   * -sh for Shadow style.
   * If no flag is provided, the Standard style is used by default.
```
student$ go run . “” | cat -e
usage: go run . <input string> <optional flag>

student$ go run . "123456"
                                             $
 _   ____    _____   _  _     ____     __    $
/ | |___ \  |___ /  | || |   | ___|   / /    $
| |   __) |   |_ \  | || |_  |___ \  | '_ \  $
| |  / __/   ___) | |__   _|   __) | | (_) | $
|_| |_____| |____/     |_|   |____/   \___/  $
                                             $
                                             $

student$ go run . “{Hello There}” | cat -e
   __  _    _          _   _           _______   _                           __    $
  / / | |  | |        | | | |         |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___      | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \     | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |    | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/     |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                        /_/   $
                                                                                   $


student$ go run . “Hello\\n\\nThere” | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $

student$ go run . “Hello\n” -t $
$
o        o o     $
|        | |     $
O--o o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
$
student$ go run . “Hello\n” -sh $
                                       $
_|                _| _|                $
_|_|_|     _|_|   _| _|   _|_|         $
_|    _| _|_|_|_| _| _| _|    _|       $
_|    _| _|       _| _| _|    _|       $
_|    _|   _|_|_| _| _|   _|_|         $
                                       $
                                       $

```
## Files
The ASCII art patterns are read from the following files:

   * `standard.txt`
   * `thinkertoy.txt`
   * `shadow.txt`

Make sure these files are present in the same directory as the executable.

## Contributors
* [Nungo Edwin]()
* [Namayi Franklyne](https://github.com/fnamayi)
