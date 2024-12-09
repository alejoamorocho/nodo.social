package models

type Store struct {
    ID          string     `json:"id" firestore:"id"`
    ContactInfo Contact    `json:"contactInfo" firestore:"contactInfo"`
    Products    []Product  `json:"products" firestore:"products"`
}

type Contact struct {
    Email     string `json:"email" firestore:"email"`
    Phone     string `json:"phone" firestore:"phone"`
    Address   string `json:"address" firestore:"address"`
    Website   string `json:"website" firestore:"website"`
}

type Product struct {
    ID              string  `json:"id" firestore:"id"`
    Name            string  `json:"name" firestore:"name"`
    Description     string  `json:"description" firestore:"description"`
    Price           float64 `json:"price" firestore:"price"`
    Contact         Contact `json:"contact" firestore:"contact"`
    NodeID          string  `json:"nodeId" firestore:"nodeId"`
    DonationPercent float64 `json:"donationPercent" firestore:"donationPercent"`
    ApprovalStatus  string  `json:"approvalStatus" firestore:"approvalStatus"`
}
