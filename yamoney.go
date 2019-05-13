package yamoney

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	paymentURL      = "https://payment.yandex.net/api/v3/payments/"
	refundURL       = "https://payment.yandex.net/api/v3/refunds/"
	defaultCurrency = "RUB"
)

var (
	debug bool
)

// SetDebug - Устанавливаем debug
func SetDebug(dbg bool) {
	debug = dbg
}

// CreatePayment - Создание платежа
func (ya *API) CreatePayment(id string, o *InitObj) (ans PaymentInfo, err error) {
	if o.Amount.Currency == "" {
		o.Amount.Currency = defaultCurrency
	}

	if o.Confirmation.Type == "" {
		o.Confirmation.Type = "redirect"
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Формируем запрос
	req, err := http.NewRequest("POST", paymentURL, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Добавляем заголовое о том что это json
	req.Header.Set("Content-Type", "application/json")
	// Добавляем уникальный ключ
	req.Header.Set("Idempotence-Key", createIdempotenceKey(id, "create"))

	return ya.sendRq(req)
}

// Capture - Подтверждаем платеж
func (ya *API) Capture(id string, o *InitObj) (ans PaymentInfo, err error) {
	if o.Amount.Currency == "" {
		o.Amount.Currency = defaultCurrency
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Формируем запрос
	req, err := http.NewRequest("POST", paymentURL+o.InvoiceID+"/capture", bytes.NewBuffer(b))
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Добавляем заголовое о том что это json
	req.Header.Set("Content-Type", "application/json")
	// Добавляем уникальный ключ
	req.Header.Set("Idempotence-Key", createIdempotenceKey(id, "capture"))

	return ya.sendRq(req)
}

// Refund - Возврат средств
func (ya *API) Refund(id string, o *InitObj) (ans PaymentInfo, err error) {
	if o.Amount.Currency == "" {
		o.Amount.Currency = defaultCurrency
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Формируем запрос
	req, err := http.NewRequest("POST", refundURL, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Добавляем заголовое о том что это json
	req.Header.Set("Content-Type", "application/json")
	// Добавляем уникальный ключ
	req.Header.Set("Idempotence-Key", createIdempotenceKey(id, "refund"))

	return ya.sendRq(req)
}

// GetPaymentInfo - получаем информацию о платеже
func (ya *API) GetPaymentInfo(invoiceID string) (ans PaymentInfo, err error) {

	// Формируем запрос
	req, err := http.NewRequest("GET", paymentURL+invoiceID, nil)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return ya.sendRq(req)
}

func (ya *API) sendRq(req *http.Request) (ans PaymentInfo, err error) {

	// Добавляем авторизацию
	req.SetBasicAuth(strconv.Itoa(ya.ShopID), ya.Secret)

	// Делаем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	//if debug {
	//	b, _ := ioutil.ReadAll(req2.Body)
	//	log.Println("[debug]", string(b))
	//}

	// Читаем ответ
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if debug {
		log.Println("[debug]", string(content))
	}

	// Если проблема с ответом
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		err = errors.New(resp.Status)
		log.Println("[error]", resp.Status, resp.StatusCode)
		log.Println("[error]", string(content))
	}

	// Парсим ответ
	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err, string(content))
		return
	}

	return
}

func createIdempotenceKey(id, method string) string {
	return method + "_" + id
}
