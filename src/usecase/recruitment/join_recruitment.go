package recruitment

import (
	"errors"
	"gin_docker/src/utils"
)

type JoinInpt struct {
	UserID        int
	RecruitmentID int `json:"recruitment_id" form:"recruitment_id" validate:"required"`
}

func (i *interactor) Join(input JoinInpt) error {
	tx := i.tx.Begin()
	// TODO: SkipDefaultTransactionをtrueにしないとチェンの呼び出しだできない
	// fmt.Println("SkipDefaultTransaction:", tx.DB().Config.SkipDefaultTransaction)
	defer tx.Rollback()

	recruitment, err := i.repository.GetRecruitmentByID(tx, input.RecruitmentID)
	if err != nil {
		return err
	}

	isFull, err := i.repository.CheckMemberLimit(tx, input.RecruitmentID, recruitment.MemberLimit)
	if err != nil {
		return err
	}

	if isFull {
		return &utils.InvalidParamError{Err: errors.New("the recruitment member is full")}
	}

	err = i.repository.JoinRecruitment(tx, input.UserID, input.RecruitmentID)
	if err != nil {
		return err
	}
	_, err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
