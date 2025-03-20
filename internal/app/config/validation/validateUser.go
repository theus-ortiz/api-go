package validation

import (
    "encoding/json"
    "errors"

    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/go-sql-driver/mysql"
    rest_err "github.com/theus-ortiz/api-go/internal/app/config/restErr"

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

// ValidateUserError valida erros de usuário e retorna um RestErr apropriado
func ValidateUserError(validationErr error) *rest_err.RestErr {
    var jsonErr *json.UnmarshalTypeError
    var jsonValidationError validator.ValidationErrors
    var mysqlErr *mysql.MySQLError

    // Verifica se é um erro de tipo inválido no JSON
    if errors.As(validationErr, &jsonErr) {
        return rest_err.NewBadRequestError("Tipo de campo inválido")
    }

    // Verifica se é um erro de validação dos campos
    if errors.As(validationErr, &jsonValidationError) {
        errorsCauses := []rest_err.Causes{}

        for _, e := range validationErr.(validator.ValidationErrors) {
            cause := rest_err.Causes{
                Field:   e.Field(), // Nome do campo com erro
                Message: e.Translate(transl), // Mensagem de erro traduzida
            }

            errorsCauses = append(errorsCauses, cause)
        }

        // Retorna um erro de validação com as causas
        return rest_err.NewBadRequestValidationError("Campos inválidos", errorsCauses)
    }

    // Verifica se é um erro de violação de constraint UNIQUE no MySQL
    if errors.As(validationErr, &mysqlErr) && mysqlErr.Number == 1062 {
        // Cria uma lista de causas com o motivo do erro
        causes := []rest_err.Causes{
            {
                Field:   "Email",
                Message: "Email já existe",
            },
        }

        // Retorna um erro personalizado para email duplicado
        return rest_err.NewEmailAlreadyExistsError("Email já existe", causes)
    }

    // Outros tipos de erro
    return rest_err.NewInternalServerError("Falha ao executar a instrução SQL")
}