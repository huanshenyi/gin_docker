package recruitment

import (
	"errors"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"

	"gin_docker/src/domain"
	"gin_docker/src/domain/recruitment"
	"gin_docker/src/domain/user"
	"gin_docker/src/infra/model"
	"gin_docker/src/utils"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) ListRecruitmentForUserID(tx domain.Tx, userID int) (recruitment.Recruitments, error) {
	conn := tx.DB()
	var rows []model.Recruitment
	if err := conn.Where("user_id = ?", userID).Preload("Tags").Find(&rows).Error; err != nil {
		return recruitment.Recruitments{}, err
	}
	rs := make([]domain.Recruitment, len(rows))
	for k, v := range rows {
		rs[k] = v.ToDomain()
	}
	return recruitment.Recruitments{
		Recruitments: rs,
	}, nil
}

func (r Repository) CreateRecruitment(tx domain.Tx, input domain.Recruitment) error {
	db := tx.DB()
	var tags []model.Tag
	if len(input.Tags) != 0 {
		for _, i := range input.Tags {
			tags = append(tags, model.Tag{ID: i.ID})
		}
	}
	if err := db.Create(&model.Recruitment{
		Title:       input.Title,
		Place:       input.Place,
		Start:       input.Start,
		End:         input.End,
		Content:     input.Content,
		Paid:        input.Paid,
		Reward:      input.Reward,
		MemberLimit: input.MemberLimit,
		Type:        model.RecruitmentTypeToSQL(input.Type),
		UserID:      input.UserID,
		Tags:        tags,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) JoinListRecruitment(tx domain.Tx, userID int, page int, limit int) (recruitment.JoinListRecruitment, error) {
	db := tx.DB()
	query := db.
		Select(`
	  R.id AS recruitment_id,
	  R.title AS recruitment_title,
	  R.place AS recruitment_place,
	  R.start AS recruitment_start,
      R.end AS recruitment_end,
	  R.content AS recruitment_content,
	  R.paid AS recruitment_paid,
	  R.reward AS recruitment_reward,
	  R.memberLimit AS recruitment_member_limit,
	  R.user_id AS recruitment_user_id,
	  R.type AS recruitment_type,
	  U.id AS user_id,
	  U.username AS user_name,
	  U.Icon AS user_icon
	`).Table(fmt.Sprintf("%s as UR", new(model.UserRecruitment).TableName())).
		Joins(fmt.Sprintf("INNER JOIN %s AS R ON R.id = UR.recruitment_id", new(model.Recruitment).TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s AS U ON R.user_id = U.id", new(model.User).TableName())).
		Where("UR.user_id = ?", userID)

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return recruitment.JoinListRecruitment{}, err
	}

	if totalCount == 0 {
		return recruitment.JoinListRecruitment{}, nil
	}

	var rows []JoinListRecruitmentRow
	if err := query.Limit(limit).Offset((page - 1) * limit).Find(&rows).Error; err != nil {
		return recruitment.JoinListRecruitment{}, err
	}

	jrs := make([]recruitment.JoinRecruitment, len(rows))
	for k, i := range rows {
		jrs[k] = i.ToDomain()
	}

	jrs, err := addTags(tx.ReadDB(), jrs)
	if err != nil {
		return recruitment.JoinListRecruitment{}, err
	}

	return recruitment.JoinListRecruitment{
		Recruitment: jrs,
		TotalPage:   int(math.Ceil(float64(totalCount) / float64(limit))),
		TotalCount:  int(totalCount),
	}, nil
}

type JoinListRecruitmentRow struct {
	RecruitmentID          int       `gorm:"recruitment_id"`
	RecruitmentTitle       string    `gorm:"recruitment_title"`
	RecruitmentPlace       string    `gorm:"recruitment_place"`
	RecruitmentStart       time.Time `gorm:"recruitment_start"`
	RecruitmentEnd         time.Time `gorm:"recruitment_end"`
	RecruitmentContent     string    `gorm:"recruitment_content"`
	RecruitmentPaid        bool      `gorm:"recruitment_paid"`
	RecruitmentReward      string    `gorm:"recruitment_reward"`
	RecruitmentMemberLimit int       `gorm:"recruitment_member_limit"`
	RecruitmentUserID      int       `gorm:"recruitment_user_id"`
	RecruitmentType        string    `gorm:"recruitment_type"`
	UserID                 int       `gorm:"user_id"`
	UserName               string    `gorm:"user_name"`
	UserIcon               string    `gorm:"user_icon"`
}

type recruitmentTag struct {
	ID            int    `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	RecruitmentID int    `gorm:"column:recruitment_id"`
}

func (r JoinListRecruitmentRow) ToDomain() recruitment.JoinRecruitment {
	jr := recruitment.JoinRecruitment{
		Recruitment: domain.Recruitment{
			ID:          r.RecruitmentID,
			Title:       r.RecruitmentTitle,
			Place:       r.RecruitmentPlace,
			Start:       r.RecruitmentStart,
			End:         r.RecruitmentEnd,
			Content:     r.RecruitmentContent,
			Paid:        r.RecruitmentPaid,
			Reward:      r.RecruitmentReward,
			MemberLimit: r.RecruitmentMemberLimit,
			UserID:      r.RecruitmentUserID,
			Type:        domain.RecruitmentType(r.RecruitmentType),
		},
		Owner: user.UserData{
			ID:       r.UserID,
			UserName: r.UserName,
			Icon:     r.UserIcon,
		},
	}
	return jr
}

func (r Repository) GetRecruitmentByID(tx domain.Tx, id int) (domain.Recruitment, error) {
	conn := tx.ReadDB()
	var row model.Recruitment
	if err := conn.Where("id = ?", id).First(&row).Error; err != nil {
		return domain.Recruitment{}, err
	}
	return row.ToDomain(), nil
}

func (r Repository) JoinRecruitment(tx domain.Tx, userID int, recruitmentID int) error {
	conn := tx.DB()
	ur := model.UserRecruitment{UserID: userID, RecruitmentID: recruitmentID}
	res := conn.Table(new(model.UserRecruitment).TableName()).Where("user_id = ? and recruitment_id = ?", userID, recruitmentID).First(&model.UserRecruitment{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if err := conn.Create(&ur).Error; err != nil {
				return err
			}
		} else {
			return res.Error
		}
		return nil
	}
	return &utils.InvalidParamError{Err: errors.New(fmt.Sprintf("Resource Existence recruitmentID:%d userID %d", recruitmentID, userID))}
}

func addTags(conn *gorm.DB, rows []recruitment.JoinRecruitment) (res []recruitment.JoinRecruitment, err error) {
	recruitmentsIDs := extractIDs(rows)
	var tagRows []recruitmentTag
	if err = conn.Debug().Table(fmt.Sprintf("%s AS RT", new(model.RecruitmentTag).TableName())).
		Select(`
	T.id AS id,
	T.name AS name,
	RT.recruitment_id AS recruitment_id
	`).Joins(fmt.Sprintf("INNER JOIN %s T ON T.id = RT.tag_id", new(model.Tag).TableName())).
		Where("RT.recruitment_id IN (?)", recruitmentsIDs).Find(&tagRows).Error; err != nil {
		return nil, err
	}
	tagMap := map[int][]domain.TagData{}
	for _, tagRow := range tagRows {
		tagMap[tagRow.RecruitmentID] = append(tagMap[tagRow.RecruitmentID], domain.TagData{
			ID:   tagRow.ID,
			Name: tagRow.Name,
		})
	}

	for i, row := range rows {
		rows[i].Recruitment.Tags = tagMap[row.Recruitment.ID]
	}
	return rows, err
}

func extractIDs(rows []recruitment.JoinRecruitment) []int {
	var recruitmentsIDs = make([]int, len(rows))
	for i, row := range rows {
		recruitmentsIDs[i] = row.Recruitment.ID
	}
	return recruitmentsIDs
}
