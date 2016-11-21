# WC_MINUS

This small program counts the characters in the input files, except for 
ASCII spaces and tabs.  It's something I wanted for a one-off investigation,
and I thought it would be a fun way to play with Go's concurrency primitives.

It reads up to 8 4k buffers of the input files at a time,  and counts the 
characters in parallel.  To accomplish this, it took a WaitGroup and a limiter 
channel.  My first inclination was to use a mutex or an atomic add for the overall
sum, but in the spirit of Go I stuck with channels for communication.  For this
small program it worked very well: a 'summer' goroutine sums the results from the
counted buffers, and it puts its final sum in yet another channel for the main
function to read and print.

This program is over-elaborate for what I used it for, but it made for a good
exercise.
