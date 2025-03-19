package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	rest_err "github.com/theus-ortiz/api-go/config/restErr"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validade = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
    var jsonErr *json.UnmarshalTypeError
    var jsonValidationError validator.ValidationErrors
    var mysqlErr *mysql.MySQLError

    // Verifica se é um erro de tipo inválido no JSON
    if errors.As(validation_err, &jsonErr) {
        return rest_err.NewBadRequestError("Invalid field type")
    }

    // Verifica se é um erro de validação dos campos
    if errors.As(validation_err, &jsonValidationError) {
        errorsCauses := []rest_err.Causes{}

        for _, e := range validation_err.(validator.ValidationErrors) {
            cause := rest_err.Causes{
                Field:   e.Field(), // Nome do campo com erro
                Message: e.Translate(transl), // Mensagem de erro traduzida
            }

            errorsCauses = append(errorsCauses, cause)
        }

        // Retorna um erro de validação com as causas
        return rest_err.NewBadResquestValidationError("Invalid fields", errorsCauses)
    }

    // Verifica se é um erro de violação de constraint UNIQUE no MySQL
    if errors.As(validation_err, &mysqlErr) && mysqlErr.Number == 1062 {
        // Cria uma lista de causas com o motivo do erro
        causes := []rest_err.Causes{
            {
                Field:   "Email",
                Message: "Email already exists",
            },
        }

        // Retorna um erro personalizado para email duplicado
        return rest_err.NewEmailAlreadyExistsError("Email already exists", causes)
    }

    // Outros tipos de erro
    return rest_err.NewInternalServerError("Failed to execute SQL statement")
}