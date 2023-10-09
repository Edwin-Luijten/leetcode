# LeetCode 
Solutions and Benchmarks 

--- 

## Getting Started 

> Install gobenchdata to parse benchmark results and create a bench.json
> ```bash 
> go install go.bobheadxi.dev/gobenchdata@latest
> ```
> Build parser
> ```bash 
> make build-parser
> ```
> Run benchmarks
> ```bash 
> make run
> ```
> Create a new challenge
> ```bash
> make new challenge="foo bar" 
> ```










## Challenges

[ADD TWO NUMBERS](#add_two_numbers)  
[LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS](#longest_substring_without_repeating_characters)  
[TWO SUM](#two_sum)  



<a name="add_two_numbers"></a>  

## [ADD TWO NUMBERS](./add_two_numbers)

Runs: 7798587  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkAddTwoNumbers-2 | 151.500000 ns/op | 48.000000 B/op | 3.000000 allocs/op |  

<a name="longest_substring_without_repeating_characters"></a>  

## [LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS](./longest_substring_without_repeating_characters)

Runs: 29090780  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkLongestSubstringWithoutRepeatingCharacters-2 | 42.440000 ns/op | 0.000000 B/op | 0.000000 allocs/op |  

<a name="two_sum"></a>  

## [TWO SUM](./two_sum)

Runs: 7277935  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkTwoSum-2 | 171.300000 ns/op | 16.000000 B/op | 1.000000 allocs/op |  
