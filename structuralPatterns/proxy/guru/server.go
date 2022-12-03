package structuralPatterns

type server interface {
	handleRequest(string, string) (int, string)
}
