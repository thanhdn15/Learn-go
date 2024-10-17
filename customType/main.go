package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	nhieuLaBai := taoNhieuBai()

	nhieuLaBai = append(nhieuLaBai, "J co")
	nhieuLaBai.in()

	bai1, bai2 := chiaBai(nhieuLaBai, 5)

	bai1.in()
	bai2.in()

	nhieuLaBai.luuFile("cobac")

	nhieuLaBai = taoBaiTuFile("cobac")
	nhieuLaBai.in()
}

func chiaBai(n nhieuBai, sl int) (nhieuBai, nhieuBai) {
	return n[:sl], n[sl:]
}

func taoNhieuBai() nhieuBai {
	kq := nhieuBai{}

	nhieuNuoc := []string{"co", "ro", "chuon", "bich"}
	nhieuNut := []string{"1", "2", "3"}

	for _, nuoc := range nhieuNuoc {
		for _, nut := range nhieuNut {
			bai := nut + " " + nuoc
			kq = append(kq, bai)
		}
	}

	return kq
}

func (n nhieuBai) chuyenThanhString() string {
	return strings.Join(n, ",")
}

func (n nhieuBai) luuFile(tenFile string) error {
	return ioutil.WriteFile(tenFile, []byte(n.chuyenThanhString()), 0666)
}

func taoBaiTuFile(tenFile string) nhieuBai {
	data, err := ioutil.ReadFile(tenFile)

	if err != nil {
		log.Printf("doc file khong duoc %v", err)

		os.Exit(1)
	}

	chuoiBai := string(data)

	nhieuLaBai := nhieuBai(strings.Split(chuoiBai, ","))

	return nhieuLaBai
}
