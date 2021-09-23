GOROOT - > Golang lib (homebrew)
export GOROOT="/opt/homebrew/Cellar/go/1.17.1/libexec"

GOPATH -> third party package
export GOPATH=$HOME/go

GOPATH dir
bin -> binary after build
 bin
├──  dlv
├──  dlv-dap
├──  go-find-references
├──  go-outline
├──  gocode
├──  godef
├──  gopkgs
├──  goplay
├──  gopls
├──  gorename
├──  goreturns
└──  staticcheck
pkg ->
src -> source code 
 src
├──  github.com/
├──  golang.org/
└──  sourcegraph.com/

       

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin