package sorterUrl

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type URLService struct {
	repo *URLRepository
}

func NewURLService(repo *URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) GetOriginalURL(code string) (string, error) {

	url, err := s.repo.FindByCode(code)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *URLService) CreateShortURL(originalUrl string) (string, error) {
	code := generateCode(6)

	err := s.repo.save(code, originalUrl)
	if err != nil {
		return "", err
	}

	return code, nil
}

func generateCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *URLService) ListShortURLs(offset int, limit int) ([]ShortURL, error) {
	return s.repo.List(offset, limit)
}
