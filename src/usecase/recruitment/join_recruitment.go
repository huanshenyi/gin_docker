package recruitment

type JoinInpt struct {
	UserID        int
	RecruitmentID int `json:"recruitment_id" form:"recruitment_id" validate:"required"`
}

func (i *interactor) Join(input JoinInpt) error {
	tx := i.tx.Begin()
	// TODO: SkipDefaultTransactionをtrueにしないとチェンの呼び出しだできない
	// fmt.Println("SkipDefaultTransaction:", tx.DB().Config.SkipDefaultTransaction)
	defer tx.Rollback()

	_, err := i.repository.GetRecruitmentByID(tx, input.RecruitmentID)
	if err != nil {
		return err
	}

	err = i.repository.JoinRecruitment(tx, input.UserID, input.RecruitmentID)
	if err != nil {
		return err
	}
	_, err = tx.Commit()
	return nil
}
