package square

import (
	"strconv"
	"time"
)

type Payment struct {
	Id                  string        `json:"id"`
	Device              Device        `json:"device"`
	Itemizations        []Itemization `json:"itemizations"`
	Refunds             []Refund      `json:"refunds"`
	Tender              []Tender      `json:"tender"`
	TaxMoney            Money         `json:"tax_money"`
	TipMoney            Money         `json:"tip_money"`
	DiscountMoney       Money         `json:"discount_money"`
	RefundedMoney       Money         `json:"refunded_money"`
	NetSalesMoney       Money         `json:"net_sales_money"`
	NetTotalMoney       Money         `json:"net_total_money"`
	GrossSalesMoney     Money         `json:"gross_sales_money"`
	AdditiveTaxMoney    Money         `json:"additive_tax_money"`
	InclusiveTaxMoney   Money         `json:"inclusive_tax_money"`
	ProcessingFeeMoney  Money         `json:"processing_fee_money"`
	TotalCollectedMoney Money         `json:"total_collected_money"`
}

func (p *Payment) WasFullyRefunded() bool {
	for _, r := range p.Refunds {
		if r.Type == "FULL" {
			return true
		}
	}
	return false
}

type Itemization struct {
	Name                string                `json:"name"`
	QuantityString      string                `json:"quantity"`
	Modifiers           []ItemizationModifier `json:"modifiers"`
	Taxes               []ItemizationTax      `json:"taxes"`
	Discounts           []Discount            `json:"discounts"`
	ItemDetail          ItemDetail            `json:"item_detail"`
	ItemVariationName   string                `json:"item_variation_name"`
	NetSalesMoney       Money                 `json:"net_sales_money"`
	DiscountMoney       Money                 `json:"discount_money"`
	TotalMoney          Money                 `json:"total_money"`
	GrossSalesMoney     Money                 `json:"gross_sales_money"`
	SingleQuantityMoney Money                 `json:"single_quantity_money"`
	SurragateID         string
	RefundSurragateID   string
}

func (i *Itemization) Quantity() float64 {
	if i.QuantityString == "" {
		return 0.0
	}

	f, err := strconv.ParseFloat(i.QuantityString, 32)

	if err != nil {
		panic(err)
	}

	return f
}

type Refund struct {
	Type                       string    `json:"type"`
	Reason                     string    `json:"reason"`
	CreatedAt                  time.Time `json:"created_at"`
	PaymentId                  string    `json:"payment_id"`
	MerchantId                 string    `json:"merchant_id"`
	ProcessedAt                time.Time `json:"processed_at"`
	RefundedMoney              Money     `json:"refunded_money"`
	RefundedTipMoney           Money     `json:"refunded_tip_money"`
	RefundedDiscountMoney      Money     `json:"refunded_discount_money"`
	RefundedAdditiveTaxMoney   Money     `json:"refunded_additive_tax_money"`
	RefundedInclusiveTaxMoney  Money     `json:"refunded_inclusive_tax_money"`
	RefundedProcessingFeeMoney Money     `json:"refunded_processing_fee_money"`
}

type Tender struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	TotalMoney    Money  `json:"total_money"`
	RefundedMoney Money  `json:"refunded_money"`
	TenderedMoney Money  `json:"tendered_money"`
	CardBrand     string `json:"card_brand"`
	PanSuffix     string `json:"pan_suffix"`
	EntryMethod   string `json:"entry_method"`
}

type ItemDetail struct {
	Sku string `json:"sku"`
}

type ItemizationModifier struct {
	Name         string `json:"name"`
	AppliedMoney Money  `json:"applied_money"`
}

type Discount struct {
	Name         string `json:"name"`
	AppliedMoney Money  `json:"applied_money"`
}

type ItemizationTax struct {
	Name          string `json:"name"`
	AppliedMoney  Money  `json:"applied_money"`
	Rate          string `json:"rate"`
	InclusionType string `json:"inclusion_type"`
}

type Device struct {
	Name string `json:"name"`
}

type Money struct {
	Amount       int64  `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}
