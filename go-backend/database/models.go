package database

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// --- 1. Table USERS ---
type User struct {
	Base

	Name string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-"`

	Role string `json:"role" gorm:"default:user"`
	Credits int `json:"credits" gorm:"default:3"`
	AestheticVibe string `json:"aesthetic_vibe"`
	IsBanned bool `json:"is_banned" gorm:"default:false"`

	Provider string `json:"provider" gorm:"default:email"`

	Age int `json:"age"`
	Gender string `json:"gender"`
	AvatarURL string `json:"avatar_url"`
	Bio string `json:"bio" gorm:"type:text"`
	InstagramHandle string `json:"instagram_handle"`
	IsVerified bool `json:"is_verified" gorm:"default:false"`

	// Relations
	Scans []ScanHistory `json:"scans" gorm:"foreignKey:UserID"`
	Likes []Like `json:"likes" gorm:"foreignKey:UserID"`
}

// --- 2. Table SCAN HISTORIES ---
type ScanHistory struct {
	Base

	UserID uint `json:"user_id"`
	User User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ImageURL string `json:"image_url"`

	Score int `json:"score"`
	AiResponse string `json:"ai_response" gorm:"type:text"`
	AiRoast string `json:"ai_roast" gorm:"type:text"`
	AestheticDetected string `json:"aesthetic_detected"`

	Caption string `json:"caption" gorm:"type:text"`
	Location string `json:"location"`
	Tags string `json:"tags" gorm:"type:text"`

	// Relations
	Likes []Like `json:"likes" gorm:"foreignKey:ScanHistoryID"`
}

// --- 3. Table LIKES ---
type Like struct {
	Base

	UserID uint `json:"user_id" gorm:"uniqueIndex:idx_user_scan"`
	User User `json:"user"`

	ScanHistoryID uint `json:"scan_history_id" gorm:"uniqueIndex:idx_user_scan"`
	ScanHistory ScanHistory `json:"scan_history"`
}

// --- 4. Table SYSTEM CONFIG ---
type SystemConfig struct {
	Key string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}