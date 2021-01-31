#' list_estimator
#'
#' Returns the position of changepoints in the sequence
#'
#' @param sample A vector of floats corresponding to the piecewise stationary sample where the retrospective changes are to be sought
#' @param minimum_distance A real number between 0 and 1 corresponding to a lower-bound on the minimum normalized length of the stationary segments (as percentage of total sample length)
#'
#' @return The list of changepoints in order of score
#'
#' @export
list_estimator <- function(sample, minimum_distance) {
  PyChest$list_estimator(sample, minimum_distance)
}
