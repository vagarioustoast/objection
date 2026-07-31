package trial

type Objection string

const (
	Hearsay Objection = "Hearsay"
)

type Exchange struct {
	Question          string
	Answer            string
	PossibleObjection Objection
}

func (exchange Exchange) IsCorrect(objection Objection) bool {
	return objection == exchange.PossibleObjection
}
