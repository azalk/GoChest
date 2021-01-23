For a CRAN submission we recommend that you fix all NOTEs, WARNINGs and ERRORs.
## Test environments
- R-hub windows-x86_64-devel (r-devel)
- R-hub ubuntu-gcc-release (r-release)
- R-hub fedora-clang-devel (r-devel)

## R CMD check results
❯ On windows-x86_64-devel (r-devel), ubuntu-gcc-release (r-release), fedora-clang-devel (r-devel)
  checking CRAN incoming feasibility ... NOTE
  Maintainer: 'Lukas Zierahn <lukas@kappa-mm.de>'
  
  New submission

0 errors ✔ | 0 warnings ✔ | 1 note ✖


## Authors comments

* The test coverage of this package is quite small at about 40% which is due to the nature of the package also being quite small
* Additionally the unit-tests have been skipped on the rhub servers as they require the Python Package "PyChest" to be installed on the reticulate environment of the server testing the package. You can run "RChest::install_PyChest()" to install "PyChest" into the environment, the tests should then Pass.