library(reticulate)
py_install("PyChest", pip=TRUE)

find_changepoints <- function(sequence, minimum_distance, process_count) {
  import("PyChest")$find_changepoints(sequence, minimum_distance, process_count)
}

list_estimator <- function(sequence, minimum_distance) {
  import("PyChest")$list_estimator(sequence, minimum_distance)
}
