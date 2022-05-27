package models

import "time"

type Post struct {
	ID            string
	Owner         string
	Title         string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ImgRelatedURL string
	Status        int
}

type New struct {
	ID            string
	Owner         string
	Title         string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ImgRelatedURL string
	Status        int
}

type Departament struct {
	ID             string
	Name           string
	BossID         string
	Address        string
	City           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Workers        int
	LastInspection time.Time
}

type (
	Report struct {
		ID          string
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Priority    int
		TagID       string
	}

	Tag struct {
		ID   string
		Name string
	}
)

type LeakedCertificate struct {
	ID     string
	Secret string
	Range  string
}
