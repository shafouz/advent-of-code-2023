BEGIN {
  count = 0
  split("one two three four five six seven eight nine", names)
  arr["one"] = 1
  arr["two"] = 2
  arr["three"] = 3
  arr["four"] = 4
  arr["five"] = 5
  arr["six"] = 6
  arr["seven"] = 7
  arr["eight"] = 8
  arr["nine"] = 9
}
{
  rev = ""
  split($1, tmp, "")
  for (i = length($1); i >= 1; i--) {
    rev = rev tmp[i]
  }

  orig = $1
  orig_lowest = length($1)
  rev_lowest = length($1)
  first = 999
  last = 999
  nom_idx = 999
  ord_idx = 999

  for (i = 1; i <= length(names); i++) {
    ord = i
    nom = names[i]

    ord_idx = index(orig, ord)
    nom_idx = index(orig, nom)

    if (ord_idx != 0 && ord_idx <= orig_lowest) {
      orig_lowest = ord_idx
      first = ord
    }

    if (nom_idx != 0 && nom_idx <= orig_lowest) {
      orig_lowest = nom_idx
      first = ord
    }

    
    cmd = "echo "nom" | rev"
    cmd | getline curr_nom
    close(cmd)

    ord_idx = index(rev, ord)
    nom_idx = index(rev, curr_nom)

    if (ord_idx != 0 && ord_idx <= rev_lowest) {
      rev_lowest = ord_idx
      last = ord
    }

    if (nom_idx != 0 && nom_idx <= rev_lowest) {
      rev_lowest = nom_idx
      last = ord 
    }
  }

  count += first last
}
END {
  print count
}
