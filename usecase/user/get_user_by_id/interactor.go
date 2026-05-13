package get_user_by_id

import (
	"context"

	"practices.com/clean_arch_go/interface/repository"
)

type GetUserByIdInteractor struct {
	outputPort OutputPort

	repo repository.UserRepository
}

func NewGetUserByIdInteractor(repo repository.UserRepository, outputPort OutputPort) *GetUserByIdInteractor {
	return &GetUserByIdInteractor{
		outputPort,
		repo,
	}
}

func (uc *GetUserByIdInteractor) Invoke(ctx context.Context, input Input) {
	ent, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		uc.outputPort.HandleError(err)
		return
	}

	// Convert Entity to Output data.
	output := Output{
		ID:    ent.ID.String(),
		Name:  ent.Name,
		Email: ent.Email,
	}

	// Call the present
	uc.outputPort.Present(output)
}
