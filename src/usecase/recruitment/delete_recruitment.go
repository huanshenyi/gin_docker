package recruitment

import "gin_docker/src/utils"

type DeleteInput struct {
	UserID        int
	RecruitmentID int `uri:"recruitmentID" binding:"required"`
}

func (i *interactor) Delete(input DeleteInput) error {
	tx := i.tx.Begin()
	defer tx.Rollback()

	recruitment, err := i.repository.GetRecruitmentByID(i.tx, input.RecruitmentID)
	if err != nil {
		return err
	}
	if recruitment.UserID != input.UserID {
		return &utils.ResourceNotPublicError{Resource: "delete recruitment"}
	}

	err = i.repository.DeleteRecruitment(i.tx, recruitment.ID)
	if err != nil {
		return err
	}

	_, err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
