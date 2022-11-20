package day3

type Triangle struct {
  s1 int
  s2 int
  s3 int
}

func validateTriangle(tri Triangle) bool {
  return tri.s1 + tri.s2 > tri.s3 &&
    tri.s2 + tri.s3 > tri.s1 &&
    tri.s3 + tri.s1 > tri.s2 
}
