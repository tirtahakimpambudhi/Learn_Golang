package repository
import (
  "ttd/models"
  )
type StudentRepository interface {
  FindAll()([]models.Student)
  FindByNIS(NIS int)(models.Student,error)
  CreateStudent(student models.Student)(error)
  UpdateStudent(NIS int,student models.Student)(error)
  DeleteStudent(NIS int)(error)
}