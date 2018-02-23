# TourofGo-LoopsandFunctionsExercise
An answer to the Loops and Functions exercise in the Tour of Go from golang

### Software Requirements ###

* [Go](http://golang.org/)

### Description ###

This is a simple program demonstrating Newton's method of iteratively computing an unknown value. In this case we use the square root of 2 as the unknown value. Some initial guess as to what the unknown value is must be entered, and with each loop this guess approaches the correct value until the guess is within .001 of the correct value. For the square root of 2, the amount the guess is adjusted by is

(z^2 - x)/(2z),

where z is the current guess, x is 2, and 1/(2z) is the derivative of z^2. Dividing by the derivative makes the next adjustment to z smaller when there was less change to z on the previous loop.
