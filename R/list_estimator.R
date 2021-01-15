#' list_estimator
#'
#' Returns the position of changepoints in the sequence
#'
#' @param sequence A vector of floats
#' @param minimum_distance A real number between 0 and 1 corresponding to a lower-bound on the minimum normalized length of the stationary segments (as percentage of total sample lengt)
#'
#' @return The list of changepoints in order of likelihood
#'
#' @export
list_estimator <- function(sequence, minimum_distance) {
  PyChest$list_estimator(sequence, minimum_distance)
}
