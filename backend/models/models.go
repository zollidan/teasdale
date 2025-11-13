package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username   string    `gorm:"uniqueIndex;not null;size:255"`
	Email      string    `gorm:"uniqueIndex;not null;size:255"`
	KeycloakID string    `gorm:"uniqueIndex;not null;size:255"`
	AvatarURL  string    `gorm:"type:text"`
	Bio        string    `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Likes             []Like              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments          []Comment           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Reviews           []Review            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	News              []News              `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL"`
}

type Artist struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"not null;size:255;index"`
	Bio       string    `gorm:"type:text"`
	ImageURL  string    `gorm:"type:text"`
	Country   string    `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Albums []Album `gorm:"foreignKey:ArtistID;constraint:OnDelete:CASCADE"`
}

type Genre struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex;not null;size:100"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time

	Albums []Album `gorm:"many2many:album_genres;constraint:OnDelete:CASCADE"`
}

type Album struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `gorm:"not null;size:255;index"`
	ArtistID    uuid.UUID `gorm:"type:uuid;not null;index"`
	CoverURL    string    `gorm:"type:text"`
	ReleaseDate *time.Time `gorm:"type:date;index"`
	Description string    `gorm:"type:text"`
	AlbumType   string    `gorm:"size:50;default:'album';index"` // album, single, ep
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Artist  Artist   `gorm:"foreignKey:ArtistID;constraint:OnDelete:CASCADE"`
	Tracks  []Track  `gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
	Genres  []Genre  `gorm:"many2many:album_genres;constraint:OnDelete:CASCADE"`
	Reviews []Review `gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
}

type Track struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `gorm:"not null;size:255;index"`
	AlbumID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Duration    int       `gorm:"comment:продолжительность в секундах"`
	TrackNumber int
	FileURL     string `gorm:"type:text;not null"`
	PlaysCount  int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Album            Album              `gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
	Likes            []Like             `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
	Comments         []Comment          `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
	TrendingScores   []TrendingScore    `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
}

type AlbumGenre struct {
	AlbumID uuid.UUID `gorm:"type:uuid;primaryKey"`
	GenreID uuid.UUID `gorm:"type:uuid;primaryKey"`
}

type Like struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_track"`
	TrackID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_track"`
	CreatedAt time.Time

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Track Track `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
}

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	TrackID   uuid.UUID `gorm:"type:uuid;not null;index"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Track Track `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
}

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_album"`
	AlbumID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_album"`
	Rating    int       `gorm:"not null;comment:от 1 до 5"`
	Title     string    `gorm:"size:255"`
	Content   string    `gorm:"type:text;not null"`
	IsEditor  bool      `gorm:"default:false;index;comment:редакторский обзор"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Album Album `gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
}

type News struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string     `gorm:"not null;size:255"`
	Content     string     `gorm:"type:text;not null"`
	AuthorID    uuid.UUID  `gorm:"type:uuid;not null;index;comment:editor user_id"`
	ImageURL    string     `gorm:"type:text"`
	PublishedAt *time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Author User `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL"`
}

type TrendingScore struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	TrackID      uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_track_date"`
	Score        float64   `gorm:"not null;index;comment:рассчитанный score для трендов"`
	Date         time.Time `gorm:"type:date;not null;index;uniqueIndex:idx_track_date"`
	LikesCount   int       `gorm:"default:0"`
	CalculatedAt time.Time `gorm:"default:now()"`

	Track Track `gorm:"foreignKey:TrackID;constraint:OnDelete:CASCADE"`
}