package main

import "log"

func sampleFunc(n int) {
	defer log.Println("I finished")

	if n%2 == 0 {
		return
	}

}

func databaseTx() {
	// tx := db.Conn.Begin()
	// defer tx.Rollback()

	// if err := tx.ProcessOne(); err != nil {
	// 	return err
	// }

	// if err := tx.ProcessTwo(); err != nil {
	// 	return err
	// }

	// if err := tx.ProcessThree(); err != nil {
	// 	return err
	// }

	// if err := tx.ProcessFour(); err != nil {
	// 	return err
	// }

	// tx.Commit() -> we ended the tx process
	// return nil
}

func main() {
	// main purpose of defer: to ensure we called the function
	// we delayed the execution of the function until before the function returns or ends
	defer log.Println("I ended")
	defer log.Println("I ended - 2nd")
	// Last In - First Out, LIFO

	log.Println("Hello World")
	log.Println("Hello World - 2nd")

}
