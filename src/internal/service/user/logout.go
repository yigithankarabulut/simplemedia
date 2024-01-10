package usersevice

import "context"

func (s *userService) Logout(ctx context.Context, userID uint) error {
	if err := s.userStorage.ChangeActive(ctx, userID, false); err != nil {
		return err
	}
	return nil
}
