# Golang Interview

1) Downloaded Go from https://golang.org/dl/
1) Setup VS Code with various extensions to aid in development following this [article](https://levene.me/boost-your-golang-development-with-these-top-vscode-extensions)
3) Initialised the project using: go mod init fizzbuzz
4) Ran the tests and outputted test PASS, inital results:
    - ^TestPrimeNumberReturnsUnchanged$ fizzbuzz ok
    - ^TestMultipleOfThreeReturnsFizz$ fizzbuzz ok
    - Multiple of 5 does not return buzz, got: 25. --- FAIL: TestMultipleOfFiveReturnsBuzz (0.00s)
5) Having looked at the fizzbuzz function, from the start i can see we're not using modulo but constants for comparison. So i changed == to % and checked for remainder of 0
6) After changing line 7 and 10 to use modulo instead of ==, all of the tests now pass.
7) Pseudocode for fizzbuzz should be: remainder of num / 5 = 0 and remainder of num / 3 = 0
8) Like other languages, golang uses && for an 'and'
9) I have to place the fizzbuzz output at the top, as the previous statements will also be true, returning too early to give the correct value
10) In its current form, fizzbuzz will be O(1) notation, as we only ever give it a single number it will always take the same amount of time
11) If we gave it a range of numbers, the notation would be o(n), n being the number of items in the range 
12) This covers the implementation, we now need to test it thoroughly, we already have some tests but we could expand this:
    - Test Invalid Values -> numbers, arrays, objects 
    - Test Edge Cases -> 
        - 0 should return fizzbuzz as its divisible by any number, 
        - int32 limit is -2,147,483,648 to 2,147,483,647 use those
        - negative numbers, -3, -15 
    - Test More an array of numbers which satisfy the requirements -> 
        - multiples of 3: 3,6,9,12,15 
        - multiples of 5: 5,10,15,20,25 
        - multiples of 3 and 5: 15, 30, 60, 120  
        - not multiples of either: 1,2,4,7,8,11
    - Test a range,
       - 1 -> 10 in a loop, to ensure correctness of output
13) Golang is strongly typed, so you cannot enter other types into the fizzbuzz function. That removes the need to test for invalid types.
14) Lets Implement these tests, the best way to arrange these tests, with comments would be to have:
      - Range Tests
      - Edge Cases
      - Basic Functionality Tests
15) I'm updating the existing tests to use an array of numbers which should satisfy the requirement
> When looking up naming conventions, i found a lot of differing opinions. [Effective Golang](https://go.dev/doc/effective_go) states that variable names should be short, especially local variables. I find in my frontend work, that longer variable names, such as 'multiplesOfThree' is easier to read and understand then 'm'. I am going to be using more verbose naming for this as no body else is working and naming convention and style has not been specified. 
16) Tests are:
    - Test Edge Cases
        - **TestZeroReturnsFizzBuzz** -> 0 should return fizzbuzz as its divisible by any number, 
        - **TestHighNumberMultipleOfThreeReturnsFizz** -> int32 limit is -2,147,483,648 to 2,147,483,647 so will use: 2,147,483,649
        - **TestHighNumberMultipleOfFiveReturnsBuzz** -> int32 limit is -2,147,483,648 to 2,147,483,647 so will use: 2,147,483,650
        - **TestHighNumberMultipleOfThreeReturnsFizzBuzz** -> int32 limit is -2,147,483,648 to 2,147,483,647 so will use: 2,147,483,655
        - **TestHighNumberPrimesReturnsUnchanged** -> int32 limit is -2,147,483,648 to 2,147,483,647 so will use: 2,147,483,659
        - **TestNegativeNumberMultipleOfThreeReturnsFizz** -> negative numbers, -3
        - **TestNegativeNumberMultipleOfFiveReturnsBuzz** -> negative numbers, -5
        - **TestNegativeNumberMultipleOfThreeReturnsFizzBuzz** -> negative numbers, -15 
        - **TestNegativeNumberPrimesReturnsUnchanged** -> negative numbers, -1
    - Test More an array of numbers which satisfy the requirements
        - **TestMultipleOfThreeReturnsFizz** -> multiples of 3: 3,6,9,12,15 
        - **TestMultipleOfFiveReturnsBuzz** -> multiples of 5: 5,10,15,20,25 
       -  **TestMultipleOfFiveAndThreeReturnsFizzBuzz** -> multiples of 3 and 5: 15, 30, 60, 120  
        - **TestPrimeNumberReturnsUnchanged** -> not multiples of either: 1,2,4,7,8,11
    - Test a range
        - **TestRangeOfValuesReturnsCorrectly** -> 1 -> 10 in a loop, to ensure correctness of output
17) I am aware that 14 tests for a simple function like Fizzbuzz can be considered overkill. However, I wanted to demonstrate some of the varying factors that you should test for when testing functions. 

