package dictionary

type Match struct {
	Score float64
	Source interface{}
}

type Corpus struct {
	fragments []fragment
}

func (c *Corpus) AddSample(words []string, source interface{}) {
	f := makeFragment(words, source)
	c.fragments = append(c.fragments, f)
}

func (c *Corpus) FindMatch(words []string) Match {
	f := makeFragment(words, nil)

	winner := Match {
		0,
		nil,
	} 

	for _, fragment := range c.fragments {
		score := fragment.ComputeSimilarity(&f)
		if score > winner.Score {
			winner.Score = score
			winner.Source = fragment.source
		}
	}
	
	return winner
}

func MakeCorpus() Corpus {
	c := Corpus {
		[]fragment{},
	}
	return c
}

