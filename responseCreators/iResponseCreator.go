package responseCreators

type iResponseCreator interface {
	CreateResponse() Response
}
