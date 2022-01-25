package main

import (
	"context"
	"fmt"
	"greet/blog/blogpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: ", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	//create Blob
	fmt.Println("Creating a blog")
	blog := blogpb.Blog{
		AuthorId: "Jagdish",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	createBlog_res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: &blog,
	})
	if err != nil {
		log.Fatalf("unexpected error :", err)
	}

	fmt.Println("Blog has been created: ", createBlog_res)
	blogId := createBlog_res.GetBlog().GetId()

	//Read Blog
	fmt.Println("Reading the blog")

	read_res, read_err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "12",
	})
	if read_err != nil {
		fmt.Println("Error Happened while reading :", read_err)
	}
	fmt.Println("Blog was read :", read_res)

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}

	read_resp, read_err2 := c.ReadBlog(context.Background(), readBlogReq)
	if read_err2 != nil {
		fmt.Println("Error Happened while reading :\n", read_err2)
	}
	fmt.Println("Blog was read :", read_resp)

	//update the blog
	fmt.Println("Update the blog")
	newblog := blogpb.Blog{
		Id:       blogId,
		AuthorId: "Jagdish J",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog (edited)",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: &newblog,
	})

	if updateErr != nil {
		fmt.Println("Error happened while updating : ", updateErr)
	}
	fmt.Println("Blog was updated :", updateRes)

	//Delete Blog
	deleteRes, deleteerr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogId,
	})
	if deleteerr != nil {
		fmt.Println("Error happened while deleting : ", deleteerr)
	}
	fmt.Printf("Blog was deleted : ", deleteRes)

	//List Blog
	stream, err1 := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if err1 != nil {
		log.Fatalf("error while calling ListBlog RPC:%v", err)
	}
	for {
		res, err2 := stream.Recv()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatalf("Something happened %v", err2)
		}
		fmt.Println(res.GetBlog())
	}
}
