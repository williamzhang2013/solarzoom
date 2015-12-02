package utils

// #include <stdio.h>
// #include <stdlib.h>
// #include <time.h>
// #include <string.h>
/*

unsigned char user_serial11[8] = {0};  // chip SN

unsigned char tx_data_orig[8] = {0};   // original random number
unsigned char tx_data_cipher[16] = {0};// the ciphertext send to device
unsigned char rx_data_cipher[16] = {0};// the ciphertext send by device
unsigned char rx_data_orig[8] = {0};   // decode the ciphertext

extern unsigned char aselp_pitch_check3(unsigned char *,unsigned char *);
extern unsigned char enc_process(unsigned char *,unsigned char *);

unsigned char _alpu_rand(void) {
	srand((unsigned)time(NULL));
	return (unsigned char)(rand());
}

void set_user_serial11(unsigned char data, int i) {
	if (i >= 0 && i <= 7) user_serial11[i] = data;
}

void set_rx_cipher_data(unsigned char data, int i) {
	if (i >= 0 && i <= 15) rx_data_cipher[i] = data;
}

void reset_user_serial11() {
	memset(user_serial11, 0, sizeof(user_serial11));
}

void reset_rx_cipher_data() {
	memset(rx_data_cipher, 0, sizeof(rx_data_cipher));
}

int decode() {
	int error = 0;
	int i = 0;

	// send the rand data to tx_data & encode it
	for (i=0; i<8; i++) tx_data_orig[i] = _alpu_rand();
	enc_process(tx_data_orig, tx_data_cipher);

	// get user_serial11
	// get rx_data_cipher

	aselp_pitch_check3(rx_data_orig, rx_data_cipher);
	if (memcmp(tx_data_orig, rx_data_orig, 8) == 0) {
		error = 0;
	} else {
		error = 50;
	}

	return error;
}

void print(char *str) {
    printf("%s\n", str);
}

void print_user_serial() {
	printf("user_serial11:[0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x]\n",
		user_serial11[0], user_serial11[1], user_serial11[2], user_serial11[3],
		user_serial11[4], user_serial11[5], user_serial11[6], user_serial11[7]);
}

void print_tx_data_orig() {
	printf("tx_data_orig:[0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x]\n",
		tx_data_orig[0], tx_data_orig[1], tx_data_orig[2], tx_data_orig[3],
		tx_data_orig[4], tx_data_orig[5], tx_data_orig[6], tx_data_orig[7]);
}

void print_tx_data_cipher() {
	printf("tx_data_cipher:[0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x]\n",
		tx_data_cipher[0], tx_data_cipher[1], tx_data_cipher[2], tx_data_cipher[3],
		tx_data_cipher[4], tx_data_cipher[5], tx_data_cipher[6], tx_data_cipher[7],
		tx_data_cipher[8], tx_data_cipher[9], tx_data_cipher[10], tx_data_cipher[11],
		tx_data_cipher[12], tx_data_cipher[13], tx_data_cipher[14], tx_data_cipher[15]);
}

void print_rx_data_cipher() {
	printf("rx_data_cipher:[0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x]\n",
		rx_data_cipher[0], rx_data_cipher[1], rx_data_cipher[2], rx_data_cipher[3],
		rx_data_cipher[4], rx_data_cipher[5], rx_data_cipher[6], rx_data_cipher[7],
		rx_data_cipher[8], rx_data_cipher[9], rx_data_cipher[10], rx_data_cipher[11],
		rx_data_cipher[12], rx_data_cipher[13], rx_data_cipher[14], rx_data_cipher[15]);
}

void print_rx_data_orig() {
	printf("rx_data_orig:[0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x, 0x%x]\n",
		rx_data_orig[0], rx_data_orig[1], rx_data_orig[2], rx_data_orig[3],
		rx_data_orig[4], rx_data_orig[5], rx_data_orig[6], rx_data_orig[7]);
}

unsigned char aselp_pitch_check3(unsigned char *dst,unsigned char *src) {
	int i = 0;
	for (i=0; i<8; i++) {
		dst[i] = tx_data_orig[i];
	}
}

unsigned char enc_process(unsigned char *src,unsigned char *dst) {
	int i = 0;
	for (i=0; i<16; i++) {
		dst[i] = 0x11+i;
	}

	return 0;
}

*/
import "C"

import (
	"fmt"
)

func SetChipSNArrayItem(sn [8]uint8) {
	for i, v := range sn {
		C.set_user_serial11(C.uchar(v), C.int(i))
	}
}

func SetChipCipherArrayItem(rx [16]uint8) {
	for i, v := range rx {
		C.set_rx_cipher_data(C.uchar(v), C.int(i))
	}
}

func PrintOrigRandom() {
	fmt.Println("PrintOrigRandom")
	C.print_tx_data_orig()
}

func PrintChipSN() {
	fmt.Println("PrintChipSN")
	C.print_user_serial()
}

func PrintAlpuPlainText() {
	fmt.Println("PrintAlpuPlainText")
	C.print_rx_data_orig()
}

func PrintAlpuCipherText() {
	fmt.Println("PrintAlpuCipherText")
	C.print_rx_data_cipher()
}

func IsPassedAuth() bool {
	err := C.decode()
	PrintOrigRandom()
	PrintAlpuPlainText()
	PrintAlpuCipherText()

	fmt.Println("IsPassedAuth", err)
	if err == 0 {
		return true
	}
	return false
}
