test_that("changepoints are found correctly", {

  data <- rnorm(750, 0, 1)
  data <- append(data, rnorm(1250, 3, 2))

  expect_lt(find_changepoints(data, 0.125, 2L)[1] - 750, 2)
})
