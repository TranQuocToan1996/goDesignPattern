package creationalPatterns

type Inode interface {
	print(string)
	clone() Inode
}
