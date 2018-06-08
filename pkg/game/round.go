package game

// Round is an entire round of Aggronym.
type Round struct {
  // Number of the round for the game.
	Number int64
	// Acronym for players to submit answers for the round.
  Acronym string
	// SubmissionTimer is the total time allowed for submitting answers per round.
	SubmissionTimer int
	// VotingTimer is the total time allowed for voting to take place per round.
	VotingTimer int
	// ResponseTimer is the total time from the start of the
	ResponseTimer int
  // Participants represents all participating Discord IDs this round.
	Participants []string
	// Responses is the Response associated with a Discord ID.
	Responses map[string]Response
	// MaximumPoints is the maximum possible points for the round.
	MaximumPoints int
}

type Response struct {
	// Answer is the acronym submission from a participant.
	Answer string
	// Duration is the total time taken to for participant to submit answer.
	Duration int64
  // Votes is the total number of votes this response has accrued.
	Votes    int64
}

func NewRound() *Round {
	// start timer for response and
	return &Round{}
}

func generateAcronym() string {
	return "asdf"
}

func (r *Round) TotalPlayers() int {
	return len(r.Participants)
}
