package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID                uint               `json:"id"`
	UserID            uint               `json:"userId" gorm:"not null" validate:"required"`
	User              User               `json:"user"`
	Name              string             `json:"name" gorm:"size:255;not null" validate:"required,maz=255"`
	StripeId          string             `json:"stripeId" gorm:"size:255;not null" validate:"required,maz=255"`
	StripeStatus      string             `json:"stripeStatus" gorm:"size:255;not null" validate:"required,maz=255"`
	StripePlan        string             `json:"stripePlan" gorm:"size:255"`
	Quantity          int                `json:"quantity"`
	TrialEndsAt       time.Time          `json:"trialEndsAt"`
	EndsAt            time.Time          `json:"endsAt"`
	CreatedAt         time.Time          `json:"createdAt" gorm:"not null"`
	UpdatedAt         time.Time          `json:"updatedAt" gorm:"not null"`
	SubscriptionItems []SubscriptionItem `json:"subscriptionItems"`
}

func (s *Subscription) TableName() string {
	return "subscriptions"
}

func (s *Subscription) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(s)
	if err != nil {
		return err
	}
	if ok := s.isExistsUser(tx); !ok {
		return errors.New("ユーザーが存在しません。")
	}
	return
}

func (s *Subscription) isExistsUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", s.UserID).First(user).Error; err != nil {
		return false
	}
	return true
}
