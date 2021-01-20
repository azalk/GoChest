#' install_PyChest
#'
#' Initializes the package and installs/updates PyChest into the local reticulate-Python environment
#'
#' @export
install_PyChest <- function() {
  reticulate::py_install("cython")
  reticulate::py_install("PyChest", pip=TRUE)
}
