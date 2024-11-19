package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// Variabel hosts
	hosts := []string{
		"10.41.144.107:4151",
		// "2",
		// "3",

	}

	// Loop tak berujung untuk menjalankan perintah setiap 10 detik
	for {
		fmt.Println("Starting to execute commands...")

		// Loop untuk setiap host
		for _, host := range hosts {
			// Buat URL endpoint berdasarkan host
			url := fmt.Sprintf("http://%s/channel/empty?topic=promocatalog__sync_user_promo_coupon_v2&channel=promocatalog__sync", host)
			// url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", host)

			// Command curl
			command := fmt.Sprintf("curl --location --request POST %s", url)
			fmt.Printf("Executing: %s\n", command)

			// Eksekusi command menggunakan exec.Command
			cmd := exec.Command("sh", "-c", command)
			output, err := cmd.CombinedOutput()

			// Tampilkan output atau error
			if err != nil {
				fmt.Printf("Error executing command for host %s: %s\n", host, err)
			} else {
				fmt.Printf("Output for host %s:\n%s\n", host, string(output))
			}

			fmt.Println("------------------------------") // Pemisah antar command
		}
		fmt.Println("All commands executed. Waiting for 10 seconds...\n")
		time.Sleep(10 * time.Second) // Tunggu 10 detik sebelum iterasi berikutnya
	}
}
