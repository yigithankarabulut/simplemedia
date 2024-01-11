package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) GetAllRequests(ctx context.Context, userID uint) ([]model.Friends, error) {
	var friends []model.Friends
	err := r.db.Where("receiver_id = ? OR sender_id = ?", userID, userID).Where("accepted = ?", false).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *friendStorage) GetAllFriends(ctx context.Context, userID uint) ([]model.Friends, error) {
	var friends []model.Friends
	err := r.db.Where("receiver_id = ? OR sender_id = ?", userID, userID).Where("accepted = ?", true).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}
