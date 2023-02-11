[ This program prints "Hello World!" and a newline to the screen, its
  length is 106 active command characters. [It is not the shortest.]

  This loop is an "initial comment loop", a simple way of adding a comment
  to a BF program such that you don't have to worry about any command
  characters. Any ".", ",", "+", "-", "<" and ">" characters are simply
  ignored, the "[" and "]" characters just have to be balanced. This
  loop and the commands it contains are ignored because the current cell
  defaults to a value of 0; the 0 value causes this loop to be skipped.
]
++++++++               Set Cell No0 to 8
[
    >++++               Add 4 to Cell No1; this will always set Cell No1 to 4
    [                   as the cell will be cleared by the loop
        >++             Add 2 to Cell No2
        >+++            Add 3 to Cell No3
        >+++            Add 3 to Cell No4
        >+              Add 1 to Cell No5
        <<<<-           Decrement the loop counter in Cell No1
    ]                   Loop until Cell No1 is zero; number of iterations is 4
    >+                  Add 1 to Cell No2
    >+                  Add 1 to Cell No3
    >-                  Subtract 1 from Cell No4
    >>+                 Add 1 to Cell No6
    [<]                 Move back to the first zero cell you find; this will
                        be Cell No1 which was cleared by the previous loop
    <-                  Decrement the loop Counter in Cell No0
]                       Loop until Cell No0 is zero; number of iterations is 8

The result of this is:
Cell no :   0   1   2   3   4   5   6
Contents:   0   0  72 104  88  32   8
Pointer :   ^

>>.                     Cell No2 has value 72 which is 'H'
>---.                   Subtract 3 from Cell No3 to get 101 which is 'e'
+++++++..+++.           Likewise for 'llo' from Cell No3
>>.                     Cell No5 is 32 for the space
<-.                     Subtract 1 from Cell No4 for 87 to give a 'W'
<.                      Cell No3 was set to 'o' from the end of 'Hello'
+++.------.--------.    Cell No3 for 'rl' and 'd'
>>+.                    Add 1 to Cell No5 gives us an exclamation point
>++.                    And finally a newline from Cell No6
