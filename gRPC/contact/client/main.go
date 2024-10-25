package main

import (
	"context"
	"github.com/thanhdn15/concrete_lean_go/gRPC/contact/contactpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50070", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("err while dial %v", err)
	}

	defer cc.Close()

	client := contactpb.NewContactServiceClient(cc)

	insertContact(client, "0907", "Contact", "Address")
	readContact(client, "0905")
	updateContact(client, &contactpb.Contact{
		Address:     "Address Update",
		Name:        "Contact Name Update",
		PhoneNumber: "0905",
	})
	deleteContact(client, "0907")
	searchContact(client, "contact")
}

func searchContact(cli contactpb.ContactServiceClient, name string) {
	resp, err := cli.Search(context.Background(), &contactpb.SearchRequest{
		SearchName: name,
	})

	if err != nil {
		log.Println("search contact err %v\n", err)
		return
	}

	log.Printf("search response %v\n", resp)
}

func deleteContact(cli contactpb.ContactServiceClient, phone string) {
	resp, err := cli.Delete(context.Background(), &contactpb.DeleteRequest{
		PhoneNumber: phone,
	})

	if err != nil {
		log.Println("delete contact err %v\n", err)
		return
	}

	log.Printf("delete response %v\n", resp)
}

func updateContact(cli contactpb.ContactServiceClient, req *contactpb.Contact) {
	resp, err := cli.Update(context.Background(), &contactpb.UpdateRequest{
		NewContact: req,
	})

	if err != nil {
		log.Println("update contact err %v\n", err)
		return
	}

	log.Printf("update response %v\n", resp)
}

func readContact(cli contactpb.ContactServiceClient, phone string) {
	req := contactpb.ReadRequest{
		PhoneNumber: phone,
	}

	resp, err := cli.Read(context.Background(), &req)

	if err != nil {
		log.Println("read contact err %v\n", err)
		return
	}

	log.Printf("Read response %v\n", resp)
}

func insertContact(cli contactpb.ContactServiceClient, phone, name, add string) {
	req := contactpb.InsertRequest{
		Contact: &contactpb.Contact{
			Address:     add,
			PhoneNumber: phone,
			Name:        name,
		},
	}

	resp, err := cli.Insert(context.Background(), &req)

	if err != nil {
		log.Println("call insert err %v\n", err)
		return
	}

	log.Printf("Insert response %v\n", resp)

}
