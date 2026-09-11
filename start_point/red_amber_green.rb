lambda { |stdout,stderr,status|
  output = stdout + stderr

  # go says this when a package will not compile, or when the test binary
  # will not link, and in either case no spec ran.
  return :amber if /\[build failed\]/.match(output)
  return :amber if /\[setup failed\]/.match(output)

  # ginkgo marks a spec that panicked, and the house convention makes an
  # error amber where a failed assertion is red. Both are counted as Failed
  # in the summary line, so the summary alone cannot tell them apart.
  return :amber if /\[PANICKED\]/.match(output)

  # A suite holding no specs proves nothing, and ginkgo calls that SUCCESS,
  # so the count is what keeps it out of green.
  return :amber if /^Ran 0 of 0 Specs/.match(output)

  return :green if /^SUCCESS! --/.match(output)
  return :red if /^FAIL! --/.match(output)
  return :amber
}
