package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/codewithdas/students-api/internal/storage"
	"github.com/codewithdas/students-api/internal/types"
	"github.com/codewithdas/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
	// "github.com/go-playground/validator/v10/non-standard/validators"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return 
		}

		if err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError((err)))
			return 
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest,response.ValidationError(validateErrs))
			return 
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user created succesfully", slog.String("userId",fmt.Sprint(lastId)))

		
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError,err)
			return 
		}

		// w.Write([]byte("welcome to students api"))
		response.WriteJson(w, http.StatusCreated,map[string]int64{"id": lastId})
	}

}

// 8:02:25 tak hua tha tutorial
// {
//     "name":"Rakesh",
//     "email":"rakesh@gmail.com",
//     "age":35
// }
