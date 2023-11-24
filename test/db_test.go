package test

import (
	"testing"
	"ttd/database"
	"ttd/repository"
	"ttd/models"
	// Import your project packages here
)

func TestDBConfig(t *testing.T) {
  db := database.GetConnection()
  studentDB := repository.NewStudent(db)
  students := studentDB.FindAll()
  for _, s := range students {
    t.Log(s.Name)
  }
  t.Log(len(students))
  defer db.Close()
}

func TestFindNIS(t *testing.T) {
  db := database.GetConnection()
  defer db.Close()
  studentDB := repository.NewStudent(db)
  student,err := studentDB.FindByNIS(34672)
  if err != nil {
    t.Fatal(err.Error())
    return
  }
  t.Log(student.Name)
}

func TestInsert(t *testing.T) {
  student := models.Student{
    NIS:9998,
    Name:"Testing 9999",
    Jurusan :"XI SIJA 2",
  }
  db := database.GetConnection()
  defer db.Close()
  studentDB := repository.NewStudent(db)
  err := studentDB.CreateStudent(student)
  
  if err != nil {
    t.Fatal(err.Error())
  }
  t.Log("Succesed Inserted")
}

func TestUpdate(t *testing.T) {
  student := models.Student{
    NIS:34998,
    Name:"Tirta Hakim Pambudhi",
    Jurusan :"XI SIJA 2",
  }
  db := database.GetConnection()
  defer db.Close()
  studentDB := repository.NewStudent(db)
  err := studentDB.UpdateStudent(9999,student)
  
  if err != nil {
    t.Fatal(err.Error())
  }
  t.Log("Succesed Updated")
}

func TestDelete(t *testing.T) {
  db := database.GetConnection()
  defer db.Close()
  studentDB := repository.NewStudent(db)
  err := studentDB.DeleteStudent(9998)
  if err != nil {
    t.Log(err.Error())
  }
  t.Log("Succesed Deleted")
}