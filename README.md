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

## Performance
Performance is in the 3-4GB/s range for garbage data:
```
$ ./garbage -size 64GiB | pv > /dev/null
64.0GiB 0:00:18 [3.45GiB/s]
```
Zero data is slighly faster:
```
$ ./garbage -source zero -size 64GiB | pv > /dev/null
64.0GiB 0:00:11 [5.65GiB/s]
```
These tests were performed on a Haswell at 3.5Ghz, the performance is largely clock-speed bound:
```
$ cat /proc/cpuinfo | grep Intel | grep i7 | head -n 1 
model name      : Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
```

## Authors
Zach Giles and Levi Gross co-authors. 
Code released under MIT License (Expat)
