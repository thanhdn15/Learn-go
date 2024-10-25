package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/thanhdn15/concrete_lean_go/gRPC/contact/contactpb"
	"log"
)

type ContactInfo struct {
	PhoneNumber string `orm:"size(15);pk"`
	Name        string
	Address     string `orm:"type(text)"`
}

func ConvertPbContact2ContactInfo(pbContact *contactpb.Contact) *ContactInfo {
	return &ContactInfo{
		PhoneNumber: pbContact.PhoneNumber,
		Name:        pbContact.Name,
		Address:     pbContact.Address,
	}
}

func ConvertContactInfo2PbContact(contact *ContactInfo) *contactpb.Contact {
	return &contactpb.Contact{
		PhoneNumber: contact.PhoneNumber,
		Name:        contact.Name,
		Address:     contact.Address,
	}
}

func (c *ContactInfo) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(c)

	if err != nil {
		log.Printf("Insert Contact %+v err %v \n", c, err)
		return err
	}

	log.Printf("Insert %+v successfully \n", c)
	return nil
}

func Read(phoneNumber string) (*ContactInfo, error) {
	o := orm.NewOrm()

	ci := &ContactInfo{
		PhoneNumber: phoneNumber,
	}

	err := o.Read(ci)

	if err != nil {
		log.Printf("Read contact %+v err %v\n", ci, err)

		return nil, err
	}

	return ci, nil
}

func (c *ContactInfo) Update() error {
	o := orm.NewOrm()

	num, err := o.Update(c)

	if err != nil {
		log.Printf("Read contact %+v err %v\n", c, err)

		return err
	}

	log.Printf("update contact %+v, affect %d row\n", c, num)

	return nil
}

func (c *ContactInfo) Delete() error {
	o := orm.NewOrm()

	num, err := o.Delete(c)

	if err != nil {
		log.Printf("delete %+v err %v\n", c, num)
		return err
	}

	log.Printf("delete contact %+v, affect %d row\n", c, num)

	return nil
}

func SearchByName(name string) ([]*ContactInfo, error) {
	result := []*ContactInfo{}

	o := orm.NewOrm()

	num, err := o.QueryTable(new(ContactInfo)).Filter("name__icontains", name).All(&result)

	if err == orm.ErrNoRows {
		log.Printf("search %s found no rows \n", name)

		return result, nil
	}

	if err != nil {
		log.Printf("search %s err %v\n", name, err)
		return nil, err
	}

	log.Printf("search %s found %d rows\n", name, num)
	return result, nil
}
