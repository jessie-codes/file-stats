# file-stats

Simple CLI tool for generating file statistics.

## Usage

`file-stats -i input.txt -k keyword.txt -o output.txt`

+ **input, i**
	+ CLI flag for which glob pattern of files to generate statistics for. 
+ **keyword, k**
	+ CLI flag for which file to get keywords from.
	+ Keyword file must be formatted to have a single keyword per line.
+ **output, o**
	+ CLI flag for which file to save the results to.