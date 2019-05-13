package yamoney

import "encoding/json"

// API - Объект платежа
type API struct {
	ShopID int
	Secret string
}

// InitObj - объект создания платежа
type InitObj struct {
	InvoiceID    string       `json:"-"`
	Amount       Amount       `json:"amount"`
	Confirmation Confirmation `json:"confirmation"`
	Description  string       `json:"description"`
	PaymentID    string       `json:"payment_id"`
	Receipt      Receipt      `json:"receipt"`
}

// Receipt - объект плательщика
type Receipt struct {
	//	TaxSystemCode int            `json:"tax_system_code"`
	Phone string         `json:"phone"`
	Email string         `json:"email"`
	Items []ReceiptItems `json:"items"`
}

// ReceiptItems - товары
type ReceiptItems struct {
	Description    string `json:"description"`
	Quantity       string `json:"quantity"`
	Amount         Amount `json:"amount"`
	VatCode        int    `json:"vat_code"`
	PaymentSubject string `json:"payment_subject"`
	PaymentMode    string `json:"payment_mode"`
}

// Amount - объект суммы платежа
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// Confirmation - объект редиректа
type Confirmation struct {
	Type            string `json:"type"`
	ReturnURL       string `json:"return_url"`
	ConfirmationURL string `json:"confirmation_url"`
}

// NotifyObj - объект нотификации
type NotifyObj struct {
	Type   string      `json:"type"`
	Event  string      `json:"event"`
	Object PaymentInfo `json:"object"`
}

// PaymentInfo - объект ответа при создании платежа
type PaymentInfo struct {
	ID                  string          `json:"id"`
	Status              string          `json:"status"`
	Description         string          `json:"description"`
	Paid                bool            `json:"paid"`
	Amount              Amount          `json:"amount"`
	RefundedAmount      Amount          `json:"refunded_amount"`
	Confirmation        Confirmation    `json:"confirmation"`
	CreatedAt           string          `json:"created_at"`
	CapturedAt          string          `json:"captured_at"`
	ExpiresAt           string          `json:"expires_at"`
	Metadata            json.RawMessage `json:"metadata"`
	PaymentMethod       PaymentMethod   `json:"payment_method"`
	Recipient           Recipient       `json:"recipient"`
	Test                bool            `json:"test"`
	ReceiptRegistration string          `json:"receipt_registration"`
	PaymentSubject      string          `json:"payment_subject"`
	PaymentMode         string          `json:"payment_mode"`
}

// PaymentMethod - объект информации о методе оплаты
type PaymentMethod struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Saved bool   `json:"saved"`
	Card  Card   `json:"card"`
	Title string `json:"title"`
}

// Card - информация о карте
type Card struct {
	First6      string `json:"first6"`
	Last4       string `json:"last4"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	CardType    string `json:"card_type"`
}

// Recipient - Информация о получателе платежа
type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}
