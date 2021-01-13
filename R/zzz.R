PyChest <- NULL

.onLoad <- function(libname, pkgname) {
  # use superassignment to update global reference to PyChest
  PyChest <<- reticulate::import("PyChest", delay_load = TRUE)
}
