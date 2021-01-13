test_that("changepoints are found correctly", {
  have_PyChest <- reticulate::py_module_available("PyChest")
  if (!have_PyChest)
    skip("PyChest not available for testing")
  
  set.seed(0)
  
  data <- rnorm(750, 0, 1)
  data <- append(data, rnorm(1250, 3, 2))

  expect_lt(find_changepoints(data, 0.125, 2L)[1] - 750, 2)
})
