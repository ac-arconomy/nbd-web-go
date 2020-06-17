package model

type Lead struct {
	FirstName string `json:"FIRST_NAME"`
	LastName string `json:"LAST_NAME"`
	Email string `json:"EMAIL"`
	Message string `json:"message__c"`
}