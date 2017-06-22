# sort_nums

A console application written in Go that performs optimised sorting for a large number of integers in a 'Comma Separated Values' file.

## Methodology used

External merge sorting with Direct K-way merge. A simplified description of it:

* **Splitting input file into sorted chunks in slow memory (hard disk)**

  Given that we want to be efficient and not keep a very large amount of integers in memory, we split the input file into chunks.

  Each chunk will be sequentially read in fast memory (RAM) and will be sorted on the spot. After that it will be written in the slow memory (hard disk).

* **Merging the sorted chunks back into an output file**

  A Direct K-Way merge is performed where k is the number of chunks we created from the previous step. The first entry of each chunk is read, stored in fast memory and evaluated so we can get the minimum of those entries. After that, the minimum entry is written on the output file and the next entry from the chunk that was used now takes the place of the previous entry.

  This process continues until all the stored chunks that were stored in slow memory are read until their end.

* **Finding a minimum entry using a Priority Queue**

  For better performance an implementation of a Priority Queue was used to find the minimum of a series of entries as described in the previous step.

  A Priority Queue associates a priority with each of its elements. Depending on our needs, when we remove an element from the queue it has the highest or the lowest priority. It is usually implemented using heaps and the implementation in this project follows that.

  A reference implementation was used as shown in the official Go documentation:

  https://golang.org/pkg/container/heap/#example__priorityQueue

## Requirements
The following software must be installed on your environment to run this project.

### Go
Your environment must be configured with Go version 1.8.*

https://golang.org/

No third-party Go libraries were used. The implementation uses only the facilities given by the Go standard library.

#### MacOS (with `brew`)

`$ brew install go`

### Apt based GNU/Linux distributions (Debian/Ubuntu)
`$ sudo apt-get install golang-go`

### Windows (with `chocolatey`)

`C:\> choco install golang`

## Get it!

`go get github.com/angelospanag/sort_nums`

## Usage

A Makefile has been provided at the root of the project with some convenient shortcut commands.

### Build
Compiles and builds a binary of the project

`make build`

### Documentation (godoc)
After running the below command, visit http://localhost:6060/pkg/github.com/angelospanag/sort_nums/ on your browser

`make doc`

### Unit testing
`make test`

### Clean generated files
`make clean`

## Running the project

The fastest way to run the project without performing any compilation is to issue a `go run` command from the root of the project. For example:

`go run main.go -file=random_10000.txt -memory=20000`

* Parameters required

  `-file`: specifies the CSV file that will serve as input

  `-memory`: how much memory (in bytes) will be the limit for holding data in RAM

* Sample output
```
  2017/06/22 02:34:20 File random_10000.txt is 48853 bytes, will be split to 4 chunks
  2017/06/22 02:34:20 Trying to merge runs of 4 chunks
  2017/06/22 02:34:20 Sorting took 63.144511ms
```

* Bonus bash command

  A nice one-liner I discovered that counts the number of elements in a CSV file. For example for a file called `sorted_output.txt` do:

  `sed 's/[^,]//g' sorted_output.txt | wc -c`

## Resources Used

This section describes the resources used in the development of this project.

* [Go 1.8.3](https://golang.org/)

* [Atom 1.18](https://atom.io/) with a [series of preferred packages](https://github.com/angelospanag/atom_packages)

* [go-plus for Atom](https://atom.io/packages/go-plus)

* My trusty MacBook Pro 2015
  * OS: 64bit Mac OS X 10.12.5 16F73
  * CPU: Intel Core i5-5257U @ 2.70GHz
  * GPU: Intel Iris Graphics 6100
  * RAM: 10328MiB / 16384MiB


* Lots and lots of coffee
