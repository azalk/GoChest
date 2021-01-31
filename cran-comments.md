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
  
    Khaleghi (10:8, 10:93)
  Possibly mis-spelled words in DESCRIPTION:
    Ryabko (10:21, 10:106)

0 errors ✔ | 0 warnings ✔ | 1 note ✖


## Authors comments

### Comments for Version 1.0.2 (Current Version)
* We amended this version following our first submission and hope to have fixed all errors, see the NEWS file for all changes.
* The possibly mis-spelled words in DESCRIPTION are the names of the authors of the packages we are citing.

### Comments for Version 1.0.1
* The test coverage of this package is quite small at about 40% which is due to the nature of the package also being quite small.
* Additionally the unit-tests have been skipped on the rhub servers as they require the Python Package "PyChest" to be installed on the reticulate environment of the server testing the package. You can run "RChest::install_PyChest()" to install "PyChest" into the environment, the tests should then Pass.