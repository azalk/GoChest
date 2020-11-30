#' list_estimator
#'
#' Returns the position of changepoints in the sequence
#'
#' @param sequence A vector of floats
#' @param minimum_distance The minimum distance between changepoints as percentage of sequence length
#'
#' @return The list of changepoints in order of likelihood
#'
#' @export
list_estimator <- function(sequence, minimum_distance) {
  PyChest <- reticulate::import("PyChest")
  PyChest$list_estimator(sequence, minimum_distance)
}
