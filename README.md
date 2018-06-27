# file-stats

Simple CLI tool for generating file statistics.

## Installation

`go get github.com/jessie-codes/file-stats`

## CLI Usage

`file-stats -i input.txt -k keyword.txt -o output.txt`

+ **input, i**
	+ CLI flag for which glob pattern of files to generate statistics for. 
+ **keyword, k**
	+ CLI flag for which file to get keywords from.
	+ Keyword file must be formatted to have a single keyword per line.
+ **output, o**
	+ CLI flag for which file to save the results to.

## Docker Usage

This repo also comes with a `docker-compose` file in order to allow easy usage via Docker. To run it via Docker, you'll need to do the following:

1. Clone the repo
2. Replace the files under `test_files` with the files you wish to process.
3. -OR- Update the `docker-compose` file to have a different volume map.

By default, the Docker image expects the following folder format:

```
/top-level-folder
	/files
		+ file1.txt
		+ file2.text
		+ ...
	keywords.txt
```

The resulting statistics will be saved to `stats.txt` in the top level folder.

To change the files the image is looking for, or the output file, edit the `CMD` in the `Dockerfile` to account for these changes.