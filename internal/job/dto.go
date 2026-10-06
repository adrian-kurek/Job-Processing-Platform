package job

type CreateDTO struct {
	Type    string `validate:"oneof=email"`
	Payload string `validate:"json"`
}
