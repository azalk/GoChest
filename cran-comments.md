## Test environments
* local R installation, R 4.0.3
* ubuntu 16.04 (on travis-ci), R 4.0.3
* win-builder (devel)

## R CMD check results

0 errors | 0 warnings | 1 note

* This is a new release.


## Authors comments

* The test coverage of this package is quite small at about 40% which is due to the nature of the package also being quite small
* Additionally the unit-tests have been skipped on the rhub servers as they require the Python Package "PyChest" to be installed on the reticulate environment of the server testing the package. You can run "RChest::install_PyChest()" to install "PyChest" into the environment, the tests should then Pass.