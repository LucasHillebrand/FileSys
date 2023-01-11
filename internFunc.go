package main

func growFileTree(inp []FileTree, newValue FileTree) []FileTree {
	out := make([]FileTree, len(inp)+1)
	for i := 0; i < len(inp); i++ {
		out[i] = inp[i]
	}
	out[len(inp)] = newValue
	return out
}

func growList(inp []string, newValue string) []string {
	out := make([]string, len(inp)+1)
	for i := 0; i < len(inp); i++ {
		out[i] = inp[i]
	}
	out[len(inp)] = newValue
	return out
}

func shrinkLeft(orig string, items int) string {
	out := ""
	for i := items; i < len(orig); i++ {
		out += string(orig[i])
	}
	return out
}
