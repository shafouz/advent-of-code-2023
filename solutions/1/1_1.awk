BEGIN {
  count = 0
}
{
  gsub(/[^[:digit:]]/, "")
  first = substr($1, 0, 1)
  last = substr($1, length($1), 1)
  count += first last
}
END {
  print count
}
