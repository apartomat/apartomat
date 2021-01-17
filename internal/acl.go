package apartomat

import "context"

type Acl struct {
}

func NewAcl() *Acl {
	return &Acl{}
}

func (acl *Acl) CanConfirmLogin(ctx context.Context, subj *UserCtx, obj string) bool {
	return true
}
