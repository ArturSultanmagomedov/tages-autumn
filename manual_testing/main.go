package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"log"
	"tages-autumn/internal/api"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewFileServiceClient(conn)

	ctx := context.Background()

	fileName := "testimg.jpg"

	//проверка на ограничение скорости
	//wg := &sync.WaitGroup{}
	//for i := 0; i < 10000; i++ {
	//	go func() {
	//		wg.Add(1)
	/////// UPLOAD FILE ///////

	buf, err := ioutil.ReadFile("manual_testing/" + fileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if _, err = c.UploadFile(ctx, &api.File{Name: fileName, B: buf}); err != nil {
		log.Fatalf("%v", err)
	}

	/////// DOWNLOAD FILE ///////

	f, err := c.DownloadFile(ctx, &api.FileName{Name: fileName})
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := ioutil.WriteFile(f.Name, f.B, 0406); err != nil {
		log.Fatalf("%v", err)
	}
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()

	/////// GET FILES LIST ///////
	list, err := c.GetFilesList(ctx, &api.Nothing{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	for i := range list.Names {
		fmt.Println(list.Names[i])
	}
}
