# iching_display 

This is a program to display *I Ching* castings.  Input is meant to come
either manually or via a helper script. 

The `/scripts/` directory has several powershell scripts to provide input
to the program:

 * **ich_coins.ps1**  uses the 3-Coins method
 * **ich_stalks.ps1** uses the yarrow stalks method
 * **ich_static.ps1** gives a random hexagram with no moving lines

I also include a bash shell version of **ich_coins** as an example.

Usage:

    $ ich_stalks.ps1 
    Casting for <877987>:
    
    50 Ting -- The Cauldron
     --Changing To-->
    18 Ku -- Degeneration
    
      ▄▄▄▄▄▄▄▄     ▄▄▄▄▄▄▄▄
      ▄▄▄  ▄▄▄     ▄▄▄  ▄▄▄
      ▄▄▄▄▄▄▄▄ --> ▄▄▄  ▄▄▄
      ▄▄▄▄▄▄▄▄     ▄▄▄▄▄▄▄▄
      ▄▄▄▄▄▄▄▄     ▄▄▄▄▄▄▄▄
      ▄▄▄  ▄▄▄     ▄▄▄  ▄▄▄

