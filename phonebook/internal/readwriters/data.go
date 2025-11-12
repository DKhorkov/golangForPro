package readwriters

import (
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
	"time"
)

var (
	now            = time.Now()
	defaultEntries = []models.Entry{
		{
			Name:       "Иван",
			Surname:    "Иванов",
			Phone:      "+7 (999) 123-45-67",
			LastAccess: now,
		},
		{
			Name:       "Мария",
			Surname:    "Петрова",
			Phone:      "+7 (999) 234-56-78",
			LastAccess: now,
		},
		{
			Name:       "Алексей",
			Surname:    "Сидоров",
			Phone:      "+7 (999) 345-67-89",
			LastAccess: now,
		},
		{
			Name:       "Елена",
			Surname:    "Михайлова",
			Phone:      "+7 (999) 456-78-90",
			LastAccess: now,
		},
		{
			Name:       "Дмитрий",
			Surname:    "Фёдоров",
			Phone:      "+7 (999) 567-89-01",
			LastAccess: now,
		},
		{
			Name:       "Ольга",
			Surname:    "Николаева",
			Phone:      "+7 (999) 678-90-12",
			LastAccess: now,
		},
		{
			Name:       "Сергей",
			Surname:    "Александров",
			Phone:      "+7 (999) 789-01-23",
			LastAccess: now,
		},
		{
			Name:       "Анна",
			Surname:    "Васильева",
			Phone:      "+7 (999) 890-12-34",
			LastAccess: now,
		},
		{
			Name:       "Михаил",
			Surname:    "Романов",
			Phone:      "+7 (999) 901-23-45",
			LastAccess: now,
		},
		{
			Name:       "Татьяна",
			Surname:    "Сергеева",
			Phone:      "+7 (999) 012-34-56",
			LastAccess: now,
		},
	}
)
