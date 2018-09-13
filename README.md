# garbage

Garbage generator with options

## Installation

```
        $ go get github.com/zgiles/garbage
```
## Usage

```
    	$ garbage > /dev/null
```
There are a number of options:
* -size X		Specify a size ( human readable even; ex 24GiB )
* -stdout		Output to Stdout even if not in a pipe ( pipe works by default )
* -source [aes,zero]	Source of garbage, either AES garbage or all zeros ( like /dev/zero )
* -output FILE		Output to a file, or either "" or "-" for Stdout
* -threads N		Number of threads to generate data. Default is two, found to be a bit more than 1, non-linear.

## Authors
Zach Giles and Levi Gross co-authors. 
Code released under MIT License (Expat)
