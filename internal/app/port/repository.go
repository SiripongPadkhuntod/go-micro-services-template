package port

type Repository struct {
	Sql           SqlRepository
	Misc          MiscRepository
	ObjectStorage ObjectStorageRepository
}

type SqlRepository interface {
}
type MiscRepository interface {
}

type ObjectStorageRepository interface {
}
