package get_user_by_id

import (
	"context"

	"practices.com/clean_arch_go/domain/user"
)

type GetUserByIdInteractor struct {
	outputPort OutputPort

	repo user.UserRepository
}

func NewGetUserByIdInteractor(repo user.UserRepository, outputPort OutputPort) *GetUserByIdInteractor {
	return &GetUserByIdInteractor{
		outputPort,
		repo,
	}
}

func (uc *GetUserByIdInteractor) Invoke(ctx context.Context, input Input) error {
	ent, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return uc.outputPort.HandleError(err)
	}

	// Convert Entity to Output data.
	output := Output{
		ID:    ent.ID.String(),
		Name:  ent.Name,
		Email: ent.Email,
	}

	// Call the present
	return uc.outputPort.Present(output)
}
