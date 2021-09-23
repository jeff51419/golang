set up break point：
use <package>.<function> or <filename>:<line number>:
# <package>.<function>
(dlv) break main.main
Breakpoint 1 set at 0x10b0958 for main.main() ./delve-demo.go:12
# <filename>:<line number>
(dlv) break delve-demo.go:7
Breakpoint 2 set at 0x10b0758 for main.demo() ./delve-demo.go:7

# use continue
use c or continue to run, you will see dlv stop on break point:
(dlv) b main.main
Breakpoint 1 set at 0x10b0958 for main.main() ./delve-demo.go:12
(dlv) c
> main.main() ./delve-demo.go:12 (hits goroutine(1):1 total:1) (PC: 0x10b0958)
     7: func demo(s string, x int) string {
     8:  ret := fmt.Sprintf("This is a demo, your input is: %s %d", s, x)
     9:  return ret
    10: }
    11:
=>  12: func main() {
    13:  s := "string"
    14:  i := 1111
    15:  fmt.Println(demo(s, i))
    16: }
(dlv)

next line: n 或 next
step in function: s or step
step out function: stepout
args: args

print

goroutine