package routes

import (
	"github.com/adminboard/adminboard/internal/routes/deposit"
	"github.com/eqto/api-server"
)

func Load(g *api.Group) {

	g.PostAction(deposit.Add).Secure()
	g.PostAction(deposit.Approve).Secure()

}
