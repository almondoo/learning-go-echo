package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type SubscriptionItem struct {
	ID             uint         `json:"id" `
	SubscriptionID uint         `json:"subscriptionId" gorm:"not null" validate:"requreid"`
	Subscription   Subscription `json:"subscription" `
	StripeId       string       `json:"stripeId" gorm:"size:255;not null" validate:"requreid,max=255"`
	StripePlan     string       `json:"stripePlan" gorm:"size:255;not null" validate:"requreid,max=255"`
	Quantity       uint         `json:"quantity" gorm:"not null" validate:"required,max=999999999"`
	CreatedAt      time.Time    `json:"createdAt" gorm:"not null"`
	UpdatedAt      time.Time    `json:"updatedAt" gorm:"not null"`
}

func (si *SubscriptionItem) TableName() string {
	return "subscription_items"
}

func (si *SubscriptionItem) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(si)
	if err != nil {
		return err
	}
	if ok := si.isExistsSubscription(tx); !ok {
		return errors.New("subscriptionsが存在しません。")
	}
	return
}

func (si *SubscriptionItem) isExistsSubscription(tx *gorm.DB) bool {
	subscription := &Subscription{}
	if err := tx.Where("id = ?", si.SubscriptionID).First(subscription).Error; err != nil {
		return false
	}
	return true
}
