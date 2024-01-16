package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CompletedAt time.Time `gorm:"type:timestamp" json:"completed_at"` // Tambahkan atribut untuk menandai kapan tugas selesai
}

//memperbarui field UpdatedAt sebelum menyimpan perubahan
func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return nil
}
