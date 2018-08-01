package model

import "gopkg.in/mgo.v2/bson"

type PatientRequest struct {
	FirstName string `json:"firstname"bson:"firstname"`
	LastName  string `json:"lastname"bson:"lastname"`
	Age       int    `json:"age"bson:"age"`
}

type PatientResponse struct {
	PatientID string `json:"patientID,omitempty"bson:"patientID"`
	FirstName string `json:"firstname"bson:"firstname"`
	LastName  string `json:"lastname"bson:"lastname"`
	Age       int    `json:"age"bson:"age"`
}

type Patient struct {
	ID        bson.ObjectId `json:"id,omitempty"bson:"_id,omitempty"`
	PatientID string        `json:"patientID,omitempty"bson:"patientID"`
	FirstName string        `json:"firstname"bson:"firstname"`
	LastName  string        `json:"lastname"bson:"lastname"`
	Age       int           `json:"age"bson:"age"`
}

type PatientCount struct {
	ID    bson.ObjectId `json:"id,omitempty"bson:"_id,omitempty"`
	Count int           `json:"count"bson:"count"`
}
