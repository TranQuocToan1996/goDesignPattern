package creationalPatterns

type Inode interface {
	Print(string)
	Clone() Inode
}
