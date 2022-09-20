## Homework 5: Due 2022.09.25 

This week we have a written part and an implementation part. 

### Written part 

Answer the questions found in `main.tex` and turn your answers in by uploading them 
back here before midnight on the due date. You may do this as many times as 
you want. Only your final submission counts.

Remember this is math class so be sure to justify your answers. You will be 
graded on correctness and 

Note: when you submit, your filename must be `solutions.pdf` and it must be a pdf. 

### Implementation part 

This week we start with the first algorithmic attempt at faster solution to 
solving the discrete logarithm problem: Shanks's algorithm. 

Make `hw5.go` pass the tests. The tests only check to see if you correctly 
compute the discrete log. Shanks's algorithm value is the speed up from 
a brute force search to O(N^{1/2} log N). Read the proof of Propositon 2.21 
and take a look at the documentation for [`sort.Search`](https://pkg.go.dev/sort#Search) 
and [`sort.Sort`](https://pkg.go.dev/sort#Sort).
