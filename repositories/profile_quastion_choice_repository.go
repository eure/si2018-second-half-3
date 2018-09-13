package repositories

import (
	"github.com/eure/si2018-second-half-3/entities"
)

type ProfileQuastionChoiceRepository struct {
	RootRepository
}

func NewProfileQuastionChoiceRepository(s *Session) *ProfileQuastionChoiceRepository {
	return &ProfileQuastionChoiceRepository{NewRootRepository(s)}
}

func (r *ProfileQuastionChoiceRepository) Create(ent entities.ProfileQuastionChoice) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *ProfileQuastionChoiceRepository) FindByProfileQuastionId(id int64) ([]entities.ProfileQuastionChoice, error) {
	var profileQuastionChoices []entities.ProfileQuastionChoice

	err := engine.
		Where("profile_quastion_element_id = ?", id).
		Find(&profileQuastionChoices)

	if err != nil {
		return profileQuastionChoices, err
	}

	return profileQuastionChoices, nil
}
