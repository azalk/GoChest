#' install_PyChest
#'
#' Initializes the package and installs/updates PyChest into the local reticulate-Python environment
#'
#' @return No return value, called to install the PyChest Package
#'
#' @export
install_PyChest <- function() {
  reticulate::py_install("cython")
  reticulate::py_install("PyChest", pip=TRUE)
}
