package core

func JoinEmbeddings(emb ...[]Vector) []Vector {
	joined := make([]Vector, len(emb[0]))
	for _, vectors := range emb {
		joined = append(joined, vectors...)
	}
	return joined
}
