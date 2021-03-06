package ent_shared

import (
	"time"

	"github.com/harmony-development/legato/server/db/ent/entgen/session"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

func (d *DB) ExpireSessions() (err error) {
	defer doRecovery(&err)
	d.Session.Delete().Where(session.ExpiresLT(time.Now())).ExecX(ctx)
	return
}

func (d *DB) SessionExpireRoutine() {
	for {
		time.Sleep(15 * time.Minute)
		err := d.ExpireSessions()
		if err != nil {
			d.Logger.Exception(err)
		}
	}
}

func (d *DB) ExtendSession(session string) (err error) {
	defer doRecovery(&err)
	d.Session.UpdateOneID(session).ExecX(ctx)
	return
}

func (d *DB) AddSession(userID uint64, session string) (err error) {
	defer doRecovery(&err)

	d.Session.Create().
		SetUser(
			d.User.Query().
				Where(
					user.ID(userID),
					user.HasLocalUser(),
				).
				OnlyX(ctx),
		).
		SetID(session).
		SaveX(ctx)

	return
}

func (d *DB) SessionToUserID(sid string) (userID uint64, err error) {
	defer doRecovery(&err)

	userID = d.Client.Session.
		Query().
		Where(session.ID(sid)).
		QueryUser().
		OnlyX(ctx).
		ID

	return
}
