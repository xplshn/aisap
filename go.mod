module github.com/mgord9518/aisap

go 1.22.5

toolchain go1.23.1

replace github.com/mgord9518/aisap/permissions => ./permissions

replace github.com/mgord9518/aisap/profiles => ./profiles

require (
	github.com/CalebQ42/squashfs v1.0.2
	github.com/adrg/xdg v0.5.0
	github.com/mgord9518/aisap/helpers v0.0.0-20240904034515-85384a4a17b0
	github.com/mgord9518/aisap/permissions v0.0.0-20240904034515-85384a4a17b0
	github.com/mgord9518/aisap/profiles v0.0.0-20240904034515-85384a4a17b0
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/CalebQ42/fuse v0.1.0 // indirect
	github.com/klauspost/compress v1.17.10 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/rasky/go-lzo v0.0.0-20200203143853-96a758eda86e // indirect
	github.com/seaweedfs/fuse v1.2.3 // indirect
	github.com/therootcompany/xz v1.0.1 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	golang.org/x/sys v0.26.0 // indirect
)
