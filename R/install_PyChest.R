#' install_PyChest
#'
#' Initializes the package and installs/Updates PyChest into the local recticulate-Python evnironment
#'
#' @export
install_PyChest <- function() {
  reticulate::py_install("cython")
  reticulate::py_install("PyChest", pip=TRUE)
}
