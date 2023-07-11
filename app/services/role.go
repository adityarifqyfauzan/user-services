package services

import (
	"context"

	"github.com/adityarifqyfauzan/user-services/app/domain/model"
	"github.com/adityarifqyfauzan/user-services/app/errors"
	"github.com/adityarifqyfauzan/user-services/app/usecase"
	pb "github.com/adityarifqyfauzan/user-services/proto/role"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleServices struct {
	pb.UnimplementedRoleServicesServer
	uc usecase.RoleUsecase
}

func NewRoleServices(uc usecase.RoleUsecase) RoleServices {
	return RoleServices{
		uc: uc,
	}
}

func (svc RoleServices) Create(ctx context.Context, in *pb.RoleCreateRequest) (*pb.Role, error) {

	role, err := svc.uc.Save(ctx, in.Name)
	if err != nil {
		switch e := err.(type) {
		case *errors.BadRequestError:
			return nil, status.Error(codes.Unimplemented, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	return &pb.Role{
		Id:   role.ID.String(),
		Name: role.Name,
		Slug: role.Slug,
	}, nil

}

func (svc RoleServices) FindAll(ctx context.Context, in *pb.RoleFindAllRequest) (*pb.Roles, error) {
	roles, err := svc.uc.FindAll(ctx, int(in.Page), int(in.Size))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	res := make([]*pb.Role, 0)

	for _, r := range roles.([]model.Role) {
		var role pb.Role

		role.Id = r.ID.String()
		role.Name = r.Name
		role.Slug = r.Slug

		res = append(res, &role)
	}

	return &pb.Roles{
		Roles: res,
	}, nil
}
