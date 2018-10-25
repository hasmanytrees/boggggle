package die

type Die []string

func (d Die) Roll(rand Randomizer) string {
	return d[rand.Intn(len(d))]
}
