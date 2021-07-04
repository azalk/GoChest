#' list_estimator
#'
#' Returns the position of changepoints in the sequence. NOTE: PyChest needs to be installed first by calling `install_PyChest'.
#'
#' @param sample A vector of floats corresponding to the piecewise stationary sample where the retrospective changes are to be sought
#' @param minimum_distance A real number between 0 and 1 corresponding to a lower-bound on the minimum normalized length of the stationary segments (as percentage of total sample length)
#'
#' @return The list of changepoints in order of score
#' @references
#' \insertRef{khaleghi2012locating}{RChest}
#' @export
list_estimator <- function(sample, minimum_distance) {
  PyChest$list_estimator(sample, minimum_distance)
}
