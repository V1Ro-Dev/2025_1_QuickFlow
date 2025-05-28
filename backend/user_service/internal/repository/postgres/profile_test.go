package postgres

import (
    "context"
    "database/sql"
    "errors"
    "testing"
    "time"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"

    "quickflow/shared/models"
)

func TestPostgresProfileRepository_SaveProfile(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()
    profile := models.Profile{
        UserId: userId,
        BasicInfo: &models.BasicInfo{
            Name:        "Test",
            Surname:     "User",
            Sex:         models.MALE,
            DateOfBirth: time.Now(),
            Bio:         "Test bio",
            AvatarUrl:   "avatar.jpg",
        },
    }

    mock.ExpectExec("insert into profile").
        WithArgs(userId, profile.BasicInfo.Bio, profile.BasicInfo.AvatarUrl, nil,
            profile.BasicInfo.Name, profile.BasicInfo.Surname, profile.BasicInfo.Sex,
            profile.BasicInfo.DateOfBirth).
        WillReturnResult(sqlmock.NewResult(1, 1))

    err = repo.SaveProfile(context.Background(), profile)
}

func TestPostgresProfileRepository_GetProfile(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()
    now := time.Now()

    // Mock main profile query
    rows := sqlmock.NewRows([]string{"id", "bio", "profile_avatar", "profile_background", "firstname", "lastname", "sex", "birth_date", "school_id", "contact_info_id", "last_seen"}).
        AddRow(userId, "bio", "avatar.jpg", nil, "Test", "User", 1, now, nil, nil, now)

    mock.ExpectQuery("select id, bio, profile_avatar, profile_background, firstname, lastname, sex, birth_date, school_id, contact_info_id, last_seen").
        WithArgs(userId).
        WillReturnRows(rows)

    // Mock education query - return no rows
    mock.ExpectQuery("select u.name, u.city, f.name, e.graduation_year").
        WithArgs(userId).
        WillReturnError(sql.ErrNoRows)

    _, err = repo.GetProfile(context.Background(), userId)
    assert.NoError(t, err)
}

func TestPostgresProfileRepository_UpdateProfileTextInfo(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()
    profile := models.Profile{
        UserId: userId,
        BasicInfo: &models.BasicInfo{
            Name:        "Test",
            Surname:     "User",
            Sex:         models.MALE,
            DateOfBirth: time.Now(),
            Bio:         "Test bio",
        },
    }

    mock.ExpectBegin()
    mock.ExpectExec("update profile").
        WithArgs(userId, profile.BasicInfo.Bio, profile.BasicInfo.Name, profile.BasicInfo.Surname, profile.BasicInfo.Sex, profile.BasicInfo.DateOfBirth).
        WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()

    err = repo.UpdateProfileTextInfo(context.Background(), profile)
    assert.NoError(t, err)
}

func TestPostgresProfileRepository_UpdateProfileAvatar(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()
    avatarUrl := "new_avatar.jpg"

    mock.ExpectQuery("update profile set profile_avatar = \\$1 where id = \\$2").
        WithArgs(avatarUrl, userId).
        WillReturnRows(sqlmock.NewRows([]string{}))

    err = repo.UpdateProfileAvatar(context.Background(), userId, avatarUrl)
    assert.NoError(t, err)
}

func TestPostgresProfileRepository_GetPublicUserInfo(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()
    now := time.Now()

    rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "profile_avatar", "username", "last_seen"}).
        AddRow(userId, "Test", "User", "avatar.jpg", "testuser", now)

    mock.ExpectQuery("select u.id, firstname, lastname, profile_avatar, username, last_seen").
        WithArgs(userId).
        WillReturnRows(rows)

    _, err = repo.GetPublicUserInfo(context.Background(), userId)
    assert.NoError(t, err)
}

func TestPostgresProfileRepository_GetPublicUsersInfo(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userIds := []uuid.UUID{uuid.New(), uuid.New()}
    now := time.Now()

    rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "profile_avatar", "username", "last_seen"}).
        AddRow(userIds[0], "Test1", "User1", "avatar1.jpg", "testuser1", now).
        AddRow(userIds[1], "Test2", "User2", "avatar2.jpg", "testuser2", now)

    mock.ExpectQuery("select u.id, firstname, lastname, profile_avatar, username, last_seen").
        WithArgs(userIds).
        WillReturnRows(rows)

    _, err = repo.GetPublicUsersInfo(context.Background(), userIds)
}

func TestPostgresProfileRepository_UpdateLastSeen(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewPostgresProfileRepository(db)

    userId := uuid.New()

    mock.ExpectExec("update profile").
        WithArgs(userId, sqlmock.AnyArg()).
        WillReturnResult(sqlmock.NewResult(1, 1))

    err = repo.UpdateLastSeen(context.Background(), userId)
    assert.NoError(t, err)
}

func TestPostgresProfileRepository_ErrorCases(t *testing.T) {
    tests := []struct {
        name     string
        mockFunc func(sqlmock.Sqlmock)
        testFunc func(*PostgresProfileRepository) error
    }{
        {
            name: "GetProfile not found",
            mockFunc: func(m sqlmock.Sqlmock) {
                m.ExpectQuery("select id, bio").WillReturnError(sql.ErrNoRows)
            },
            testFunc: func(r *PostgresProfileRepository) error {
                _, err := r.GetProfile(context.Background(), uuid.New())
                return err
            },
        },
        {
            name: "UpdateProfileTextInfo rollback",
            mockFunc: func(m sqlmock.Sqlmock) {
                m.ExpectBegin()
                m.ExpectExec("update profile").WillReturnError(errors.New("db error"))
                m.ExpectRollback()
            },
            testFunc: func(r *PostgresProfileRepository) error {
                return r.UpdateProfileTextInfo(context.Background(), models.Profile{UserId: uuid.New()})
            },
        },
        {
            name: "GetPublicUserInfo error",
            mockFunc: func(m sqlmock.Sqlmock) {
                m.ExpectQuery("select u.id").WillReturnError(errors.New("db error"))
            },
            testFunc: func(r *PostgresProfileRepository) error {
                _, err := r.GetPublicUserInfo(context.Background(), uuid.New())
                return err
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            db, mock, err := sqlmock.New()
            if err != nil {
                t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
            }
            defer db.Close()

            repo := NewPostgresProfileRepository(db)
            tt.mockFunc(mock)

            err = tt.testFunc(repo)
        })
    }
}

func TestHelperFunctions(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    t.Run("updateContactInfo", func(t *testing.T) {
        mock.ExpectBegin()
        mock.ExpectQuery("SELECT id FROM contact_info").WillReturnError(sql.ErrNoRows)
        mock.ExpectQuery("INSERT INTO contact_info").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
        mock.ExpectCommit()

        tx, _ := db.Begin()
        _, err := updateContactInfo(context.Background(), tx, models.ContactInfo{})
        assert.NoError(t, err)
    })

    t.Run("updateSchoolInfo", func(t *testing.T) {
        db, mock, err := sqlmock.New()
        mock.ExpectBegin()
        mock.ExpectQuery("SELECT id FROM school").WillReturnError(sql.ErrNoRows)
        mock.ExpectQuery("INSERT INTO school").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
        mock.ExpectCommit()

        tx, _ := db.Begin()
        _, err = updateSchoolInfo(context.Background(), tx, models.SchoolEducation{})
        assert.NoError(t, err)
    })

    t.Run("updateUniversityInfo", func(t *testing.T) {
        db, mock, err := sqlmock.New()
        mock.ExpectBegin()
        mock.ExpectQuery("insert into university").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
        mock.ExpectQuery("insert into faculty").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
        mock.ExpectExec("insert into education").WillReturnResult(sqlmock.NewResult(1, 1))
        mock.ExpectCommit()

        tx, _ := db.Begin()
        err = updateUniversityInfo(context.Background(), tx, uuid.New(), models.UniversityEducation{})
        assert.NoError(t, err)
    })
}
