package daos

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Video struct {
	Slug        string
	Title       string
	Description string
	Status      string
	Uploader    User
}

type VideoDAO struct {
	conn *pgxpool.Pool
}

func NewVideoDAO(conn *pgxpool.Pool) *VideoDAO {
	return &VideoDAO{conn}
}

func (v *VideoDAO) Get(slug string) (*Video, error) {
	var video Video
	row := v.conn.QueryRow(context.Background(), "SELECT slug, username, title, description, status FROM videos LEFT JOIN users ON videos.uploader = users.id WHERE slug=$1;", slug)
	err := row.Scan(&video.Slug, &video.Uploader.Username, &video.Title, &video.Description, &video.Status)
	return &video, err
}

func (v *VideoDAO) GetAll() ([]Video, error) {
	var videos []Video

	rows, err := v.conn.Query(context.Background(), "SELECT slug, username, title, description, status FROM videos LEFT JOIN users ON videos.uploader = users.id;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var video Video
		err = rows.Scan(&video.Slug, &video.Uploader.Username, &video.Title, &video.Description, &video.Status)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (v *VideoDAO) Create(slug string, uploader *User, title string, description string) (*Video, error) {
	_, err := v.conn.Exec(context.Background(), "INSERT INTO videos (slug, uploader, title, description, status) VALUES ($1, $2, $3, $4, $5);", slug, uploader.ID, title, description, "PROCESSING")
	return &Video{Slug: slug, Uploader: *uploader, Title: title, Description: description}, err
}
