# ASCIIPIC

A Go port of the original version I did in scala, [also on github](https://github.com/rwtodd/small_programs/tree/master/ascii_pic).
It is interesting to compare the code.  This time, the Go version isn't really much more verbose than the scala one.

~~~~~~
$ bin/asciipic 
Usage: asciipic fname width

$ bin/asciipic ~/Downloads/tennis-racket.jpg 40
                                        
                  .,**===:.             
               .:$@A%$$$%%%*            
             .=%AA%+:,..,:+@$.          
            :%#@+:..     . :$+          
          .+AA+, .,,,. . .  :%*         
         ,%A%: .,.,,,.. . .  =%         
        :%A=. .,. .:,,  . .. :@,        
       ,$A* .,,.  .:,, . . . ,@*        
      :=@:  ,,,,. .::,....   ,@=        
     ,+$,   .,:,...::,. . .  :@*        
     =%:    ,,,,...:::.      =@,        
    :%+  .  ,,,:...::. . .   $%.        
    +%,     .::,..,,.       :%+         
   .%$  . . .::::,..        =@:         
   ,@=  . . ,,::,          ,%%          
   ,@: . . ..,::.         .=@*          
   ,@: . . ..:,,,         =+%           
   .%* ..    ,,,.        =$%=           
    ++  . .  .,,.      .$@+%.           
    ,%: .   .., .     :%A$%$            
     =%:     .      ,+AA*:@=            
      +%=:.      .:$A#$. *A:            
       =%$+=*:**+%AA$:   =@,            
        .*+$$$+++$+.     +@.            
           .,**++++*.    =@.            
                .,=$+:   =%.            
                   ,=$*  =%.            
                     :$+=@A:            
                      ,%A@A%            
                       =@@@A,           
                       .%@@A$           
                        =@@@A,          
                        .%@@A$          
                         =@@@A,         
                         ,%@@@$         
                          +@@@A,        
                          ,@@%@$        
                           +@@@A,       
                           :@@@@$       
                            $@@@@:      
                            :%@@@@:     
                             $AA%A%     
                             %#%:%@.    
                             +#@=@$     
                             .@#@$.     
                               ,,       

~~~~~~
