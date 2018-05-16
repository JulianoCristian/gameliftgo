package gameliftgo

type DescribePlayerSessionsRequest struct {
	GameSessionID             string
	Limit                     int
	NextToken                 string
	PlayerID                  string
	PlayerSessionID           string
	PlayerSessionStatusFilter string
}

type DescribePlayerSessionsResponse struct {
	NextToken      string
	PlayerSessions []PlayerSession
}
