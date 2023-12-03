package handler
import(
	"github.com/suhasbhagate/GoCode/GoCodeRevision/GRPCDemo/pb"
)

type Server struct{
	pb.UnimplementedServiceServer
}