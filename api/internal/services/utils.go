package services

import (
	"strings"
)

func getPageNum(subdirectory string) string {
  pageNumWithSlash := strings.Split(subdirectory, "/frases_pensadores/")[1]
  pageNum := strings.Split(pageNumWithSlash, "/")[0]

  return pageNum
}

