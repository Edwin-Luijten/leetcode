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
> Run benchmarks & tests
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
[MEDIAN OF TWO SORTED ARRAYS](#median_of_two_sorted_arrays)  
[TWO SUM](#two_sum)  



<a name="add_two_numbers"></a>  

## [ADD TWO NUMBERS](./add_two_numbers)

Runs: 10245224  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkAddTwoNumbers-2 | 120.800000 ns/op | 48.000000 B/op | 3.000000 allocs/op |  

<a name="longest_substring_without_repeating_characters"></a>  

## [LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS](./longest_substring_without_repeating_characters)

Runs: 30588536  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkLongestSubstringWithoutRepeatingCharacters-2 | 38.940000 ns/op | 0.000000 B/op | 0.000000 allocs/op |  

<a name="median_of_two_sorted_arrays"></a>  

## [MEDIAN OF TWO SORTED ARRAYS](./median_of_two_sorted_arrays)

Runs: 24341931  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkTestFindMedianSortedArrays-2 | 50.380000 ns/op | 32.000000 B/op | 1.000000 allocs/op |  

<a name="two_sum"></a>  

## [TWO SUM](./two_sum)

Runs: 8813082  

| Name | ns/op | B/op | allocs/op |  
| ---- | ----- | ---- | --------- |  
| BenchmarkTwoSum-2 | 130.900000 ns/op | 16.000000 B/op | 1.000000 allocs/op |  
