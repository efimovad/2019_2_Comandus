package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/store/sqlstore"
	"reflect"
	"testing"
)

func TestJobRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store := sqlstore.New(db)
	rows := sqlmock.
		NewRows([]string{"id"})

	var elemID int64 = 1
	expect := []*model.Job{
		{ ID: elemID },
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID)
	}

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1

	j := testJob(t, m)
	j.BeforeCreate()

	// TODO: uncomment when validation will be implemented
	/*if err := f.Validate(); err != nil {
		t.Fatal()
	}*/

	//ok query
	mock.
		ExpectQuery(`INSERT INTO jobs`).
		WithArgs(j.HireManagerId, j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId,
			j.PaymentAmount, j.Country, j.City, j.JobTypeId, j.Date, j.Status).
		WillReturnRows(rows)

	err = store.Job().Create(j, m)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if j.ID != 1 {
		t.Errorf("bad id: want %v, have %v", u.ID, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// query error
	mock.
		ExpectQuery(`INSERT INTO jobs`).
		WithArgs(j.HireManagerId, j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId,
			j.PaymentAmount, j.Country, j.City, j.JobTypeId, j.Date, j.Status).
		WillReturnError(fmt.Errorf("bad query"))

	err = store.Job().Create(j, m)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestJobRepository_Find(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	var elemID int64 = 1
	// good query
	rows := sqlmock.
		NewRows([]string{"id", "managerId", "title", "description", "files", "specialityId", "experienceLevelId",
			"paymentAmount", "country", "city", "jobTypeId", "date", "status" })

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1

	j := testJob(t, m)
	j.BeforeCreate()

	expect := []*model.Job{
		j,
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.HireManagerId, item.Title, item.Description, item.Files, item.SpecialityId,
			item.ExperienceLevelId, item.PaymentAmount, item.Country, item.City, item.JobTypeId, item.Date, item.Status)
	}

	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
		"country, city, jobTypeId, date, status FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	store := sqlstore.New(db)

	item, err := store.Job().Find(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId, date, status FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = store.Job().Find(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	expect = []*model.Job{
		testJob(t, m),
	}

	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId, date, status FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = store.Job().Find(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestJobRepository_Edit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}

	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store := sqlstore.New(db)

	rows := sqlmock.
		NewRows([]string{"id"})

	var elemID int64 = 1
	expect := []*model.Job{
		{ ID: elemID },
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID)
	}

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1
	j := testJob(t, m)
	j.BeforeCreate()

	//ok query
	j.City = "sim city"
	j.Country = "nnn"
	j.Description = "no description"

	// TODO: uncomment when validation will be implemented
	/*if err := j.Validate(); err != nil {
		t.Fatal()
	}*/

	mock.
		ExpectQuery(`UPDATE jobs SET`).
		WithArgs(j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId, j.PaymentAmount, j.Country,
			j.City, j.JobTypeId, j.Status, j.ID).
		WillReturnRows(rows)

	err = store.Job().Edit(j)
	if err != nil {
		t.Fatal(err)
	}

	if j.ID != 1 {
		t.Errorf("bad id: want %v, have %v", j.ID, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestJobRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "managerId", "title", "description", "files", "specialityId", "experienceLevelId",
			"paymentAmount", "country", "city", "jobTypeId", "date", "status" })

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1

	j1 := testJob(t,m)
	j1.Title = "job1"
	j1.BeforeCreate()

	j2 := testJob(t,m)
	j2.Title = "job2"
	j2.BeforeCreate()

	j3 := testJob(t,m)
	j3.Title = "job3"
	j3.BeforeCreate()

	expect := []*model.Job{
		j1,
		j2,
		j3,
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.HireManagerId, item.Title, item.Description, item.Files, item.SpecialityId,
			item.ExperienceLevelId, item.PaymentAmount, item.Country, item.City, item.JobTypeId, item.Date, item.Status)
	}

	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId, date, status FROM jobs LIMIT 10").
		WillReturnRows(rows)

	store := sqlstore.New(db)

	jobs, err := store.Job().List()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	for i := 0; i < 3; i++ {
		if !jobs[i].IsEqual(*expect[i]) {
			t.Errorf("results not match, want %v, have %v", expect[i], jobs[i])
			return
		}
	}

	// query error
	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId, date, status FROM jobs LIMIT 10").
		WillReturnError(fmt.Errorf("db_error"))

	_, err = store.Job().List()
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}