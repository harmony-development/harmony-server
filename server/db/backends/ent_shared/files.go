package ent_shared

import (
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/ent/entgen/filehash"
)

func (d *DB) AddFileHash(fileID string, hash []byte) (err error) {
	defer doRecovery(&err)
	d.FileHash.
		Create().
		SetFileid(fileID).
		SetHash(hash).
		SaveX(ctx)
	return
}

func (d *DB) DeleteFileMeta(fileID string) (err error) {
	defer doRecovery(&err)
	d.File.DeleteOneID(fileID).ExecX(ctx)
	return
}

func (d *DB) GetFileIDByHash(hash []byte) (fileID string, err error) {
	defer doRecovery(&err)
	fileID = d.FileHash.Query().Where(filehash.Hash(hash)).OnlyX(ctx).Fileid
	return
}

func (d *DB) GetFileMetadata(fileID string) (file *entgen.File, err error) {
	defer doRecovery(&err)
	file = d.File.GetX(ctx, fileID)
	return
}

func (d *DB) SetFileMetadata(fileID string, contentType, name string, size int) (err error) {
	defer doRecovery(&err)
	d.File.
		UpdateOneID(fileID).
		SetContenttype(contentType).
		SetName(name).
		SetSize(size).
		ExecX(ctx)
	return
}
