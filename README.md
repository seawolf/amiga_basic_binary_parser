# AmigaBASIC Binary-to-ASCII Converter

### Background

The first couple of Commodore's Amiga products were released with a desktop
environment, Workbench, and these earlier versions included AmigaBASIC on the
"Extras" disk.

### Problem

While AmigaBASIC can read, use and save plain-text files for programs, some
efficiencies were introduced by using a binary format.

* Some special (non-ASCII) characters are used to represent keywords.
* Definitions of and usages of Labels, Subroutine names, etc. are kept towards
  the end of the file, and pointers to them are used in the program.
* ... and many more!

### Solution

The file given to this program by a command-line argument is read as binary.

With the binary data loaded, we are ready to parse the file:
1. Extract the first couple of bytes into the header.
2. Ensure the header denotes an AmigaBASIC file, and one that is not encrypted.
3. Identify the program portion (body).
4. Transform the known special characters into the keywords they represent.
5. Transform the references to variable/label/etc. names back into the program.

The result is then saved to a new file.