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
// MarginAccountOrderStrategies struct for MarginAccountOrderStrategies
type MarginAccountOrderStrategies struct {
	Session Session `json:"session,omitempty"`
	Duration Duration `json:"duration,omitempty"`
	OrderType OrderType `json:"orderType,omitempty"`
	CancelTime OrderCancelTime `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType ComplexOrderStrategyType `json:"complexOrderStrategyType,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
	FilledQuantity float32 `json:"filledQuantity,omitempty"`
	RemainingQuantity float32 `json:"remainingQuantity,omitempty"`
	RequestedDestination DestinationExchange `json:"requestedDestination,omitempty"`
	DestinationLinkName string `json:"destinationLinkName,omitempty"`
	ReleaseTime string `json:"releaseTime,omitempty"`
	StopPrice float32 `json:"stopPrice,omitempty"`
	StopPriceLinkBasis PriceLinkBasis `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType PriceLinkType `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset float32 `json:"stopPriceOffset,omitempty"`
	StopType StopType `json:"stopType,omitempty"`
	PriceLinkBasis PriceLinkBasis `json:"priceLinkBasis,omitempty"`
	PriceLinkType PriceLinkType `json:"priceLinkType,omitempty"`
	Price float32 `json:"price,omitempty"`
	TaxLotMethod TaxLotMethod `json:"taxLotMethod,omitempty"`
	OrderLegCollection []SavedOrderOrderLegCollection `json:"orderLegCollection,omitempty"`
	ActivationPrice float32 `json:"activationPrice,omitempty"`
	SpecialInstruction SpecialInstruction `json:"specialInstruction,omitempty"`
	OrderStrategyType OrderStrategyType `json:"orderStrategyType,omitempty"`
	OrderId float32 `json:"orderId,omitempty"`
	Cancelable bool `json:"cancelable,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
	EnteredTime string `json:"enteredTime,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
	Tag string `json:"tag,omitempty"`
	AccountId float32 `json:"accountId,omitempty"`
	OrderActivityCollection Execution `json:"orderActivityCollection,omitempty"`
	ReplacingOrderCollection []map[string]interface{} `json:"replacingOrderCollection,omitempty"`
	ChildOrderStrategies []map[string]interface{} `json:"childOrderStrategies,omitempty"`
	StatusDescription string `json:"statusDescription,omitempty"`
}
