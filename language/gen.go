//go:generate pigeon -no-recover -o language.go language.peg
//go:generate goimports -w language.go

package language