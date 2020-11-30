#' init_RChest
#'
#' Initializes the package and installs/Updates PyChest into the local recticulate-Python evnironment
#'
#' @export
init_RChest <- function() {
  reticulate::py_install("PyChest", pip=TRUE)
}
