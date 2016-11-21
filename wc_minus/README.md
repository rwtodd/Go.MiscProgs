# WC_MINUS

This small program counts the characters in the input files, except for 
ASCII spaces and tabs.  It's something I wanted for a one-off investigation,
and I thought it would be a fun way to play with Go's concurrency primitives.

It reads up to 8 4k buffers of the input files at a time,  and counts the 
characters in parallel (at least, if the input data can be read fast enough, it can
go in parallel!).  My first inclination was to use a mutex or an atomic add for the overall
sum, but in the spirit of Go I stuck with channels for communication.  Each worker thread
tracks its own sum, and when the input channel closes, they send their sum to another channel
to be part of a grand total computed in the main() function.  For this
small program it worked very well.

This program is over-elaborate for what I used it for, but it made for a good
exercise.
