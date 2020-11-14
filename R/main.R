
init_RChest <- function() {
  library(reticulate)
  py_install("PyChest", pip=TRUE)
}

find_changepoints <- function(sequence, minimum_distance, process_count) {
  library(reticulate)
  PyChest <- import("PyChest")
  PyChest$find_changepoints(sequence, minimum_distance, process_count)
}

list_estimator <- function(sequence, minimum_distance) {
  library(reticulate)
  PyChest <- import("PyChest")
  PyChest$list_estimator(sequence, minimum_distance)
}
