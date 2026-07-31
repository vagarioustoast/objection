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
