package mysql

import (
	"Goi/models"
	"Goi/settings"
	"testing"
)

func init() {
	dbCfg := &settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "",
		DbName:       "Web",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Status:      0,
		Title:       "Test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}

	t.Logf("CreatePost insert record into mysql success.")

}
