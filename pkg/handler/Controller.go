package handler

import "github.com/willie-lin/YEVER/pkg/database/ent"

type Controller struct {
	Client *ent.Client
	Err    error
}
