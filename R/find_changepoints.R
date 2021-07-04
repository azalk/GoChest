#' find_changepoints
#'
#' Returns the position of changepoints in the sequence. NOTE: PyChest needs to be installed first by calling `install_PyChest'.
#'
#' @param sample A vector of floats corresponding to the piecewise stationary sample where the retrospective changes are to be sought
#' @param minimum_distance A real number between 0 and 1 corresponding to a lower-bound on the minimum normalized length of the stationary segments (as percentage of total sample length)
#' @param process_count The different number of distinct stationary processes present. 
#'
#' @return The list of changepoints in increasing size
#'
#' @references
#' \insertRef{khaleghi2014asymptotically}{RChest}
#' 
#' \insertRef{khaleghi2012locating}{RChest}
#' @export
find_changepoints <- function(sample, minimum_distance, process_count) {
  PyChest$find_changepoints(sample, minimum_distance, process_count)
}



