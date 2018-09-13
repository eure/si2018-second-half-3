package repositories

import (
	"github.com/eure/si2018-second-half-3/entities"
)

type ProfileQuastionElementRepository struct {
	RootRepository
}

func NewProfileQuastionElementRepository(s *Session) *ProfileQuastionElementRepository {
	return &ProfileQuastionElementRepository{NewRootRepository(s)}
}

func (r *ProfileQuastionElementRepository) Create(ent entities.ProfileQuastionElement) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *ProfileQuastionElementRepository) Update(ent *entities.ProfileQuastionElement) error {
	s := r.GetSession().Where("id = ?", ent.ID)
	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}
