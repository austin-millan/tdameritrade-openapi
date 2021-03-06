/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// FutureOptions struct for FutureOptions
type FutureOptions struct {
	AskPriceInDouble float32 `json:"askPriceInDouble,omitempty"`
	BidPriceInDouble float32 `json:"bidPriceInDouble,omitempty"`
	ClosePriceInDouble float32 `json:"closePriceInDouble,omitempty"`
	ContractType PutOrCall `json:"contractType,omitempty"`
	DeltaInDouble float32 `json:"deltaInDouble,omitempty"`
	Description string `json:"description,omitempty"`
	Digits float32 `json:"digits,omitempty"`
	ExchangeName string `json:"exchangeName,omitempty"`
	ExerciseType string `json:"exerciseType,omitempty"`
	ExpirationType string `json:"expirationType,omitempty"`
	FutureExpirationDate float32 `json:"futureExpirationDate,omitempty"`
	FutureIsActive bool `json:"futureIsActive,omitempty"`
	FutureIsTradable bool `json:"futureIsTradable,omitempty"`
	FuturePercentChange float32 `json:"futurePercentChange,omitempty"`
	FutureTradingHours string `json:"futureTradingHours,omitempty"`
	GammaInDouble float32 `json:"gammaInDouble,omitempty"`
	HighPriceInDouble float32 `json:"highPriceInDouble,omitempty"`
	InTheMoney bool `json:"inTheMoney,omitempty"`
	LastPriceInDouble float32 `json:"lastPriceInDouble,omitempty"`
	LowPriceInDouble float32 `json:"lowPriceInDouble,omitempty"`
	Mark float32 `json:"mark,omitempty"`
	MoneyIntrinsicValueInDouble float32 `json:"moneyIntrinsicValueInDouble,omitempty"`
	MultiplierInDouble float32 `json:"multiplierInDouble,omitempty"`
	NetChangeInDouble float32 `json:"netChangeInDouble,omitempty"`
	OpenInterest float32 `json:"openInterest,omitempty"`
	OpenPriceInDouble float32 `json:"openPriceInDouble,omitempty"`
	RhoInDouble float32 `json:"rhoInDouble,omitempty"`
	SecurityStatus string `json:"securityStatus,omitempty"`
	StrikePriceInDouble float32 `json:"strikePriceInDouble,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	ThetaInDouble float32 `json:"thetaInDouble,omitempty"`
	Tick float32 `json:"tick,omitempty"`
	TickAmount float32 `json:"tickAmount,omitempty"`
	TimeValueInDouble float32 `json:"timeValueInDouble,omitempty"`
	Underlying string `json:"underlying,omitempty"`
	VegaInDouble float32 `json:"vegaInDouble,omitempty"`
	Volatility float32 `json:"volatility,omitempty"`
}
