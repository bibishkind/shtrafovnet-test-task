package grpc

import (
	"context"

	"github.com/bibishkind/shtrafovnet-test-task/pb"
)

func (srv *Server) GetCompanyByINN(ctx context.Context, req *pb.CompanyByINNRequest) (*pb.CompanyByINNResponse, error) {
	company, err := srv.service.GetCompanyByINN(req.Inn)
	if err != nil {
		return nil, err
	}

	return &pb.CompanyByINNResponse{
		Company: &pb.Company{
			Inn: company.INN,
			Kpp: company.KPP,
			Owner: &pb.Owner{
				Firstname:  company.Owner.Firstname,
				Middlename: company.Owner.Middlename,
				Lastname:   company.Owner.Lastname,
			},
		},
	}, nil
}
