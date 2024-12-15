Only works on linux (ubuntu)

Before running, lpsolve must be installed with:

`sudo apt-get install liblpsolve55-dev`

Then environment variables must be set to:

`export CGO_CFLAGS="-I/usr/include/lpsolve"`

`export CGO_LDFLAGS="-llpsolve55 -lm -ldl -lcolamd"`