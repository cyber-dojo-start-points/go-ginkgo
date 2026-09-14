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

  # A run in which no spec actually ran proves nothing, and ginkgo calls that
  # SUCCESS. The suite may hold no specs at all (Ran 0 of 0) or hold specs that
  # that are all marked pending with PIt (Ran 0 of 1), so it is the first count
  # that keeps either out of green.
  return :amber if /^Ran 0 of \d+ Specs/.match(output)

  # A spec marked with FIt runs alone and skips its siblings, so the suite says
  # SUCCESS while saying nothing about the specs it never ran. ginkgo announces
  # that by adding FOCUSED to the PASS line and exiting 197.
  return :amber if /^PASS \| FOCUSED$/.match(output)

  return :green if /^SUCCESS! --/.match(output)
  return :red if /^FAIL! --/.match(output)
  return :amber
}
