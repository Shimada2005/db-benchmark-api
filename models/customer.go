package models

type Customer struct {
    CustomerID       string `json:"customer_id"`
    FirstName        string `json:"first_name"`
    LastName         string `json:"last_name"`
    Company          string `json:"company"`
    City             string `json:"city"`
    Country          string `json:"country"`
    Phone1           string `json:"phone1"`
    Phone2           string `json:"phone2"`
    Email            string `json:"email"`
    SubscriptionDate string `json:"subscription_date"`
    Website          string `json:"website"`
}