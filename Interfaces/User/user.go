package User

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name          string             `bson:"name" json:"name,omitempty"`
	Email         string             `bson:"email" json:"email,omitempty"`
	Cpf           string             `bson:"cpf" json:"cpf,omitempty"`
	Rg            string             `bson:"rg" json:"rg,omitempty"`
	Sex           string             `bson:"sex" json:"sex,omitempty"`
	BirthDay      time.Time          `bson:"birthDay" json:"birthDay,omitempty"`
	PhoneNumber   []string           `bson:"phoneNumber" json:"phoneNumber,omitempty"`
	PictureUrl    string             `bson:"pictureUrl" json:"pictureUrl,omitempty"`
	Terms         UserTerms          `bson:"terms" json:"terms,omitempty"`
	MedicalRecord MedicalRecord      `bson:"medicalRecord" json:"medicalRecord,omitempty"`
	Address       UserAddress        `bson:"address" json:"address,omitempty"`
	DentalArch    []DentalArchObj    `bson:"dentalArch" json:"dentalArch,omitempty"`
}

type DentalArchObj struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string
	Procedures []DentalArchProceduresObj
}

type DentalArchProceduresObj struct {
	ScheduleId string
	Services   []DentalArchProceduresServicesObj
}

type UserAddress struct {
	Cep        string
	StreetName string
	Number     int32
	Complement string
	City       string
	State      string
}

type UserTerms struct {
	AcceptRegisterTerms      bool
	AcceptMedicalRecordTerms bool
	AcceptClinicsAccess      bool
}

type MedicalRecord struct {
	Ocupation             string
	EmergencyNumber       string
	EmergencyName         string
	AestheticTreatment    string
	MedicinesAllergy      string
	ConstantMedice        string
	UsesAcidOnSkin        string
	CurrentMedicTreatment string
	HasCancer             string
	BloodPressure         string
	BloodType             string
	HearthProblems        string
	HasBruxism            bool
	Smoke                 bool
	ClenchTeeth           bool
	SunExposition         bool
	LactoseIntolerance    bool
	HasDiabetes           bool
}

type DentalArch struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string
	Procedures []DentalArchProceduresObj
}

type DentalArchProcedures struct {
	ScheduleId string
	Services   []DentalArchProceduresServicesObj
}

type DentalArchProceduresServicesObj struct {
	ServiceId string
}
