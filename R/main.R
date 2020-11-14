library(reticulate)
py_install("PyChest", pip=TRUE)

PyChest <- import("PyChest")

find_changepoints <- function(sequence, minimum_distance, process_count) {
  PyChest$find_changepoints(sequence, minimum_distance, process_count)
}

list_estimator <- function(sequence, minimum_distance) {
  PyChest$list_estimator(sequence, minimum_distance)
}
