# Locating distributional changes in piece-wise stationary time-series with long-range dependencies

Implementations of the change-point estimation algorithms proposed in the following two papers.

 - A. Khaleghi, D. Ryabko, Asymptotically Consistent Estimation of the Number of Change Points in Highly Dependent Time Series In Proceedings of the International Conference on Machine Learning, 2014.

 - A. Khaleghi, D. Ryabko, Locating Changes in Highly-Dependent Data with an Unknown Number of Change-Points, In Proceedings of Neural Information Processing Systems, 2012.


written in Go and distributed in [Python](#python), [R](#R) and [Go](#go).


## Python
### Installing

The Python package is precompiled for Windows (32/64 bit), MacOs (64 bit) and many Linux (32/64 bit) distributions. If you want to use the code on system for which there are no precompiled binaries like Raspberry Pis or Android phones head to the [Compiling](#compiling-the-go-code) section in the Appendix.

If Python is installed on your system simply run the following code

```
pip install PyChest
```

### How to use

To find changepoint positions in a sequence (given as either a python list or a numpy array), with a minimum distance of `0.03 * len(sequence)` between changepoints which have been generated by `2` different processes simply call:

```Python
import PyChest
estimates = PyChest.find_changepoints(sequence, 0.03, 2)
```

`estimates` will now be a python list containing a list of changepoints in increasing order. The list can be emtpy. 

If you do not know the number of generating processes you can call the `list_estimator` instead like following

```Python
import PyChest
estimates = PyChest.list_estimator(sequence, 0.03)
```

Now `estimates` is a list of changepoint estimates at least `0.03 * len(sequence)` apart in decreasing order of a `score` associated to each estimate by the algorithm, which reflects the quality of the estimate; the list can be empty.
 
## R
### Installing
The R implementation uses the Python package which is precompiled for Windows (32/64 bit), MacOs (64 bit) and many Linux (32/64 bit) distributions. If you want to use the code on system for which there are no precompiled binaries like Raspberry Pis or Android phones head to the [Compiling](#compiling-the-go-code) section in the Appendix.

You will need to have [Python](https://www.python.org/) installed on your system in order to run the R-package.

Install the Package by running the following code in any R console:

```R
devtools::install_github("azalk/GoChest")
library("RChest")
init_RChest()
```

`init_RChest` is used to install PyChest onto your local Python environment, call it to update the PyChest version R uses. You should not need to do this everytime you use the R-package. 

### How to use

To find changepoint positions in a sequence (given as either a Matrix/Array or a Multi-element vector), with a minimum distance of `0.03 * length(sequence)` between changepoints which have been generated by `2` different processes simply call:

```R
library("RChest")
estimates <- find_changepoints(sequence, 0.03, 2L)
```

Now `estimates` will now be a Multi-element vector containing a list of changepoints at least `0.03 * length(sequence)` apart in increasing order. The list can be emtpy. 

If you do not know the number of generating processes you can call the `list_estimator` instead like following

```R
library("RChest")
estimates <- list_estimator(sequence, 0.03)
```

Now `estimates` is a list of changepoints at least `0.03 * length(sequence)` apart in decreasing likelihood. The list can be empty. 

## Go
### Installing
Install the Go source code by entering the following line in the console:
```
go get github.com/azalk/GoChest
```

Go will now complain that the directory layout is unexpected. That is because of the duality of it being a Python and Go package and can safely be ignored.

### How to use
 
 To find changepoint positions in a sequence (given as a slice), with a minimum distance of `0.03 * len(sequence)` between changepoints which have been generated by `2` different processes simply call:

```go
import (
    "github.com/azalk/GoChest/GoChest"
)
func main() {
    sequence := make([]float64, 0)

    // fill sequence somehow

    estimates = GoChest.FindChangepoints(sequence, 0.03, 2)
}
```

`estimates` will now be a slice containing a list of changepoints at least `0.03 * len(sequence)` apart in increasing order. The slice can be emtpy. 

If you do not know the number of generating processes you can call the `ListEstimator` instead like following

```Go
import (
    "github.com/azalk/GoChest/GoChest"
)

func main() {
    sequence := make([]float64, 0)

    // fill sequence somehow

    estimates = GoChest.ListEstimator(sequence, 0.03)
}
```

Now `estimates` is a slice of changepoints at least `0.03 * len(sequence)` apart in decreasing likelihood. The slice can be empty. 


## How to cite

If you use this package, please cite the following papers:

 - A. Khaleghi, D. Ryabko, Asymptotically Consistent Estimation of the Number of Change Points in Highly Dependent Time Series In Proceedings of the International Conference on Machine Learning, 2014.

 - A. Khaleghi, D. Ryabko, Locating Changes in Highly-Dependent Data with an Unknown Number of Change-Points, In Proceedings of Neural Information Processing Systems, 2012.



## Appendix
### Compiling the Go code
You should not need to compile the Go code yourself if you just want to use the package on Windows, MacOS or a common Unix distribution, you can find instructions how to install the code here: [Python](#python).

If you want to compile the Go code yourself you need to install [Go](https://golang.org/) first, then clone and enter the repository by typing the following code into a commandline:
```
git clone https://github.com/azalk/GoChest
cd GoChest
```

Next you will need to actually compile the code, for Windows enter the following line in the console to do that
```bat
go build -buildmode=c-shared -o PyChestBuild/GoChest.dll CWrapper.go
```
And for MacOS and Linux
```bat
go build -buildmode=c-shared -o PyChestBuild/GoChest.so CWrapper.go
```

Your compiled c-shared library is now situated in the `PyChestBuild` folder. To complete the installation with simply type:

```
pip install .
```

The package will now prefer your compiled library over the precompiled libraries.

#### Compiling for R
If you want to use the R-package with your own compiled binary, first install the R-Package by typing the following into an R-Commandline:

```R
devtools::install_github("azalk/GoChest")
```

Next, compile the Go code by following the [Compiling the Go code](#Compiling the Go code) chapter.
Finally, install the Python package into the virtual environment used by reticulate.

**Do not** run the `init_RChest` command as that might override your self-compiled binary, if the Python or Go code updates you will need to recompile the code.

#### Cross Compiling
Cross compiling, for example to make the package available on Android phones, is a little more complicated. In order to build the code you will need to enable cgo specifically and provide a c-compiler for your target platform. Your final command should look something like this:
```
CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=ValidCCompilerForThePlatform go build -buildmode=c-shared -o PyChestBuild/GoChest.dll CWrapper.go
```
Which would compile for windows/amd64 architecture given that the `CC=ValidCCompilerForThePlatform` is replaced by an actual c-compiler for windows/amd64.

I recommend compiling on system wherever possible.
