package routes

import (
	"database/sql"

	"github.com/adityarifqyfauzan/user-services/app/domain/repository"
	"github.com/adityarifqyfauzan/user-services/app/services"
	"github.com/adityarifqyfauzan/user-services/app/usecase"
	"google.golang.org/grpc"

	pb "github.com/adityarifqyfauzan/user-services/proto/role"
)

func InitServices(server *grpc.Server, db *sql.DB) {

	// role repo
	roleRepo := repository.NewRoleRepository()
	roleUsecase := usecase.NewRoleUsecase(db, roleRepo)
	roleServices := services.NewRoleServices(roleUsecase)

	pb.RegisterRoleServicesServer(server, roleServices)
}
