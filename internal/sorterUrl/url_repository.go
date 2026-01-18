package sorterUrl

import "database/sql"

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) save(code string, originalUrl string) error {
	_, err := r.db.Exec(`
	INSERT INTO public.short_urls (code, original_url)
	VALUES ($1, $2)`,
		code, originalUrl)

	return err
}

func (r *URLRepository) FindByCode(code string) (string, error) {
	var url string

	err := r.db.QueryRow(`
		SELECT original_url
		FROM short_urls
		WHERE code = $1
	`, code).Scan(&url)

	return url, err
}

func (r *URLRepository) List(offset int, limit int) ([]ShortURL, error) {
	rows, err := r.db.Query(`
		SELECT code as Code, original_url as OriginalURL
		FROM short_urls
		offset $1 limit $2
	`, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	shorts := []ShortURL{}
	for rows.Next() {
		var short ShortURL
		if err := rows.Scan(&short.Code, &short.OriginalURL); err != nil {
			return nil, err
		}
		shorts = append(shorts, short)

	}

	return shorts, err
}
