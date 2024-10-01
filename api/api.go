package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//Status code (ex: 200)
	Code    int
	Balance int64
}

type Error struct {
	//Error code
	Code    int
	Message string
}
