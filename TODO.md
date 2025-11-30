todo: remember to still add the cross platform compiler toolchain

the remaining parts are 
x86-32
parsing, like turning x86 and arm assembly syntax (still -nq like syntax, maybe your choice)
and then assembly into machine code from there
arm support (hopefully easier)

object file creation support (elf64/32, coff, then mach-o)
cross toolchain (go universal assembly), then ts codegen support
linker

start the actual cli

add proper error handling (to prevent bugs that will be unfixably confusing)
build for platforms and release

again make it all super cross platform

use to make an SSA compiler for drift and maybe an acc / ccc / 3c c compiler proof of concept.