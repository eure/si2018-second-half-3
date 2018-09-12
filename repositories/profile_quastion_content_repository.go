package repositories

import (
	"github.com/eure/si2018-second-half-3/entities"
)

type ProfileQuastionContentRepository struct {
  RootRepository
}

func NewProfileQuastionContentRepository(s *Session) *ProfileQuastionContentRepository {
  return &ProfileQuastionContentRepository{ NewRootRepository(s) }
}

func (r *ProfileQuastionContentRepository) Create(ent entities.ProfileQuastionContent) error {
  s := r.GetSession()
  if _, err := s.Insert(&ent); err != nil {
    return err
  }

  return nil
}
