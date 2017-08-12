package main

import "fmt"
import s "strings"

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasPrefix:", s.HasPrefix("test", "t"))
	p("hasSuffiex:", s.HasSuffix("test", "t"))
	p("index", s.Index("test", "e"))
	p("join", s.Join([]string{"a", "b"}, "-"))
	p("repeat", s.Repeat("a", 4))
	p("replace", s.Replace("foo", "o", "0", -1))
	p("replace", s.Replace("foo", "o", "0", 1))
	p("split", s.Split("test", "e"))
	p("tolower", s.ToLower("TEsT"))
	p("char", "test"[1])
}
