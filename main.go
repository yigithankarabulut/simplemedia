package main

import (
	"github.com/yigithankarabulut/simplemedia/src/apiserver"
	"log"
	"os"
)

func main() {
	if err := apiserver.New(
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
	); err != nil {
		log.Fatal(err)
	}
}

/*
-- comment ve likes handler, servicer
-- paginate gelen getalllar
-- validation değişti usera falan bakılacak image, concurrency incele
-- TODOLAR
-- constantlar eklenecek düzenlenecek
-- post için shorturl eklenecek
-- loglar
-- recover ve cors maybe
-- mainpage gibi bi endpoint oluşturulacak ve burada kullanıcının arkadaşlarının paylaşımları gösterilecek
-- swagger eklenecek
*/
