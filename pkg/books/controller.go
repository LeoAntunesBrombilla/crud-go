package books

import (
  "gorm.io/gorm"
)

type handler struct {
  DB *gorm.DB
}
