package soteapi

import "time"

// CreatedAndUpdateDate represents created and updated date
type CreatedAndUpdateDate struct {
	CreatedAt time.Time `xml:"created_at"`
	UpdatedAt time.Time `xml:"updated_at"`
}

// Producer represents producer type
type Producer struct {
	ID            int    `xml:"id"`
	Name          string `xml:"name"`
	Image         string `xml:"image"`
	ImageFilename string `xml:"image_filename"`
	URL           string `xml:"url"`
}

// ProducerWithTimestamp represent producer with timestamp
type ProducerWithTimestamp struct {
	CreatedAndUpdateDate
	Producer
}

// Category represents category type
type Category struct {
	ID                   int    `xml:"id"`
	Name                 string `xml:"name"`
	ParentID             int    `xml:"parent_id"`
	Image                string `xml:"image"`
	ImageFilename        string `xml:"image_filename"`
	Description          string `xml:"description"`
	MainPage             int    `xml:"main_page"`
	IsActive             int    `xml:"is_active"`
	IsHidden             int    `xml:"is_hidden"`
	ShowChildrenProducts int    `xml:"show_children_products"`
	URL                  string `xml:"url"`
}

// CategoryWithTimestamp represent category with timestamp
type CategoryWithTimestamp struct {
	CreatedAndUpdateDate
	Category
}

// Order represents order type
type Order struct {
	ID                               int     `xml:"id"`
	Number                           int     `xml:"number"`
	IsConfirmed                      int     `xml:"is_confirmed"`
	SfGuardUserID                    int     `xml:"sf_guard_user_id"`
	OrderStatusID                    int     `xml:"order_status_id"`
	Description                      int     `xml:"description"`
	OrderUserDataDeliveryCountry     string  `xml:"order_user_data_delivery_country"`
	OrderUserDataDeliveryFullName    string  `xml:"order_user_data_delivery_full_name"`
	OrderUserDataDeliveryAddress     string  `xml:"order_user_data_delivery_address"`
	OrderUserDataDeliveryAddressMore string  `xml:"order_user_data_delivery_address_more"`
	OrderUserDataDeliveryName        string  `xml:"order_user_data_delivery_name"`
	OrderUserDataDeliverySurname     string  `xml:"order_user_data_delivery_surname"`
	OrderUserDataDeliveryStreet      string  `xml:"order_user_data_delivery_street"`
	OrderUserDataDeliveryHouse       string  `xml:"order_user_data_delivery_house"`
	OrderUserDataDeliveryFlat        string  `xml:"order_user_data_delivery_flat"`
	OrderUserDataDeliveryCode        string  `xml:"order_user_data_delivery_code"`
	OrderUserDataDeliveryTown        string  `xml:"order_user_data_delivery_town"`
	OrderUserDataDeliveryPhone       string  `xml:"order_user_data_delivery_phone"`
	OrderUserDataDeliveryCompany     string  `xml:"order_user_data_delivery_company"`
	OrderUserDataDeliveryVatNumber   string  `xml:"order_user_data_delivery_vat_number"`
	OrderUserDataBillingCountry      string  `xml:"order_user_data_billing_country"`
	OrderUserDataBillingFullName     string  `xml:"order_user_data_billing_full_name"`
	OrderUserDataBillingAddress      string  `xml:"order_user_data_billing_address"`
	OrderUserDataBillingAddressMore  string  `xml:"order_user_data_billing_address_more"`
	OrderUserDataBillingName         string  `xml:"order_user_data_billing_name"`
	OrderUserDataBillingSurname      string  `xml:"order_user_data_billing_surname"`
	OrderUserDataBillingStreet       string  `xml:"order_user_data_billing_street"`
	OrderUserDataBillingHouse        string  `xml:"order_user_data_billing_house"`
	OrderUserDataBillingFlat         string  `xml:"order_user_data_billing_flat"`
	OrderUserDataBillingCode         string  `xml:"order_user_data_billing_code"`
	OrderUserDataBillingTown         string  `xml:"order_user_data_billing_town"`
	OrderUserDataBillingPhone        string  `xml:"order_user_data_billing_phone"`
	OrderUserDataBillingCompany      string  `xml:"order_user_data_billing_company"`
	OrderUserDataBillingVatNumber    string  `xml:"order_user_data_billing_vat_number"`
	OrderDeliveryName                string  `xml:"order_delivery_name"`
	OrderDeliveryCost                float32 `xml:"order_delivery_cost"`
	OrderDeliveryTax                 float32 `xml:"order_delivery_tax"`
	OrderDeliveryNumber              string  `xml:"order_delivery_number"`
	OrderCurrencyName                string  `xml:"order_currency_name"`
	OrderCurrencyExchange            string  `xml:"order_currency_exchange"`
	OrderCurrencyShortcut            string  `xml:"order_currency_shortcut"`
	MerchantNotes                    string  `xml:"merchant_notes"`
	OrderAllegroAuctionID            string  `xml:"order_allegro_auction_id"`
	OrderDiscountValue               float32 `xml:"order_discount_value"`
	OrderDiscountType                string  `xml:"order_discount_type"`
	TotalAmount                      float32 `xml:"total_amount"`
}

// OrderWithTimestamp represent order with timestamp
type OrderWithTimestamp struct {
	CreatedAndUpdateDate
	Order
}
