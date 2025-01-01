package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Devanshgoel-123/students-api/internal/storage"
	"github.com/Devanshgoel-123/students-api/internal/types"
	"github.com/Devanshgoel-123/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)


func New(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return 
		}

		//request Validations using go-playground validator

		if err:=validator.New().Struct(student); err!=nil{
			validateErrs:=err.(validator.ValidationErrors) //type conversion of error into validator.ValidationErrors type
			response.WriteJson(w,http.StatusBadRequest,response.ValdationError(validateErrs))
			return
		}
		slog.Info("creating a student")

		lastId,err2:=storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user created successfully", slog.String("userId",fmt.Sprint(lastId)))
		if(err2!=nil){
			response.WriteJson(w,http.StatusInternalServerError,err)
			return
		}
		response.WriteJson(w,http.StatusCreated,map[string]int{"int":int(lastId)})

	}
}

func GetStudentById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r * http.Request){
		id:=r.PathValue("id")
		slog.Info("getting a student", slog.String("id",id))
		numberId,_:=strconv.Atoi(id)
		student,err2:=storage.GetStudentById(numberId)
		if err2!=nil{
			slog.Error("error getitng user", slog.String("id",id))
			response.WriteJson(w,http.StatusInternalServerError,response.GeneralError(err2))
			return
		}

		response.WriteJson(w,http.StatusOK,student)
	}
}