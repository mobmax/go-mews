package mews_test

import (
	"os"
	"testing"
	"time"

	"github.com/mobmax/go-mews"
	"github.com/mobmax/go-mews/accountingcategories"
	"github.com/mobmax/go-mews/accountingitems"
	"github.com/mobmax/go-mews/bills"
	"github.com/mobmax/go-mews/companies"
	"github.com/mobmax/go-mews/customers"
	"github.com/mobmax/go-mews/reservations"
)

func getClient() *mews.Client {
	// get username & password
	accessToken := os.Getenv("MEWS_ACCESS_TOKEN")
	clientToken := os.Getenv("MEWS_CLIENT_TOKEN")

	// build client
	client := mews.NewClient(nil, accessToken, clientToken)
	client.SetDebug(true)
	client.SetBaseURL(mews.BaseURLDemo)
	client.SetDisallowUnknownFields(true)

	return client
}

func TestBillsAll(t *testing.T) {
	client := getClient()
	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &bills.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	_, err := client.Bills.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountingItems(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &accountingitems.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	_, err := client.AccountingItems.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountingCategories(t *testing.T) {
	client := getClient()

	requestBody := &accountingcategories.AllRequest{}
	_, err := client.AccountingCategories.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestCompanies(t *testing.T) {
	client := getClient()

	requestBody := &companies.AllRequest{}
	_, err := client.Companies.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestCustomers(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &customers.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	requestBody.TimeFilter = customers.CustomerTimeFilterCreated
	_, err := client.Customers.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestReservations(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &reservations.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	requestBody.Extent = reservations.ReservationExtent{
		BusinessSegments:  true,
		Customers:         true,
		Items:             true,
		Products:          true,
		Rates:             true,
		Reservations:      true,
		ReservationGroups: true,
		Services:          true,
		Spaces:            true,
	}
	requestBody.TimeFilter = reservations.ReservationTimeFilterCreated
	_, err := client.Reservations.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestConfig(t *testing.T) {
	client := getClient()

	requestBody := client.Configuration.NewGetRequest()
	_, err := client.Configuration.Get(requestBody)
	if err != nil {
		t.Error(err)
	}
}
