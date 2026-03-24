package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 01_read_binary_file.go - Protokol Tingkat Rendah (Binary)
// Analogi: Membaca Kode Morse dengan urutan yang tepat.

type FileHeader struct {
	MagicNumber uint32
	Version     uint16
	Reserved    uint16
}

func main() {
	// Simulasi data biner (8 byte)
	// Big Endian: 0xCAFEBABE (Magic), 0x0001 (Ver), 0x0000 (Res)
	data := []byte{0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x01, 0x00, 0x00}
	
	reader := bytes.NewReader(data)
	var header FileHeader

	// Membaca data biner langsung ke dalam struct
	err := binary.Read(reader, binary.BigEndian, &header)
	if err != nil {
		fmt.Println("Binary Read Error:", err)
		return
	}

	fmt.Printf("Magic: %X\n", header.MagicNumber)
	fmt.Printf("Version: %d\n", header.Version)
}
