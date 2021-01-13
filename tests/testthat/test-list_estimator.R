test_that("changepoints are found correctly", {
  have_PyChest <- reticulate::py_module_available("PyChest")
  if (!have_PyChest)
    skip("PyChest not available for testing")
  
  set.seed(0)
  
  data <- rnorm(750, 0, 1)
  data <- append(data, rnorm(1250, 3, 2))

  expect_lt(list_estimator(data, 0.125)[1] - 750, 2)
})
