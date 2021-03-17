package DAOs

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

func NewVideoDAO(conn *pgxpool.Pool) VideoDAO {
	return VideoDAO{conn}
}

func (v *VideoDAO) getVideo(slug string) (Video, error) {
	var video Video
	row := v.conn.QueryRow(context.Background(), "SELECT slug, username, title, description, status FROM videos LEFT JOIN users ON videos.uploader = users.id WHERE slug=$1;", slug)
	err := row.Scan(&video.Slug, &video.Uploader.Username, &video.Title, &video.Description, &video.Status)
	return video, err
}
