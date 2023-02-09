package main

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/employee"
	"tokoku/item"
	"tokoku/transaction"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var employeeAuth = employee.EmployeeAuth{DB: conn}
	var itemAuth = item.ItemAuth{DB: conn}
	var custAuth = customer.CustAuth{DB: conn}
	var transAuth = transaction.TransAuth{DB: conn}
	//Menu Login
	loginMenu := 0
	for loginMenu != 9 {
		fmt.Print("\n===TOKOKU Project===\n\n")
		fmt.Println("1. Login")
		fmt.Println("2. About")
		fmt.Println("9. Exit")
		fmt.Print("Enter an option : ")
		fmt.Scanln(&loginMenu)

		switch loginMenu {
		case 1:
			username, password := "", ""
			fmt.Print("\n--- Login\n")
			fmt.Print("Username : ")
			fmt.Scanln(&username)
			fmt.Print("Password : ")
			fmt.Scanln(&password)

			if username == "admin" && password == "admin" {
				//Menu Utama Admin
				menuAdm := 0
				for menuAdm != 9 {
					fmt.Print("\n--- Administrator Menu\n\n")
					fmt.Println("1. Add new employee")
					fmt.Println("2. Delete employee")
					fmt.Println("3. Delete item")
					fmt.Println("4. Delete customer")
					fmt.Println("5. Delete transaction")
					fmt.Println("9. Log out")
					fmt.Print("Enter an option : ")
					fmt.Scanln(&menuAdm)

					switch menuAdm {
					case 1:
						var newEmp employee.Employee
						var tmpS string
						fmt.Print("\n---Employee Register\n")
						fmt.Print("Insert username\t: ")
						fmt.Scanln(&tmpS)
						newEmp.SetUsername(tmpS)
						fmt.Print("Insert name\t: ")
						scanner.Scan()
						tmpS = scanner.Text()
						newEmp.SetName(tmpS)
						fmt.Print("Insert password\t: ")
						fmt.Scanln(&tmpS)
						newEmp.SetPassword(tmpS)
						res, err := employeeAuth.RegisterEmp(newEmp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Register Success")
						} else {
							fmt.Println("Register Failed")
						}
					case 2:
						fmt.Print("\n--- Delete Employee\n")
						employeeAuth.ShowEmps()
						var username string
						fmt.Print("\nEnter username that you want to delete : ")
						fmt.Scanln(&username)
						res, err := employeeAuth.DeleteEmp(username)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete Employee success")
						} else {
							fmt.Println("Delete Employee failed")
						}

					case 3:
						fmt.Print("\n--- Delete Item\n")
						itemAuth.ShowItems()
						var idItem int
						fmt.Print("\nEnter id item that you want to delete : ")
						fmt.Scanln(&idItem)
						res, err := itemAuth.DeleteItem(idItem)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete item success")
						} else {
							fmt.Println("Delete item failed")
						}

					case 4:
						fmt.Print("\n--- Delete Customer\n")
						custAuth.ShowCust()
						var noHp string
						fmt.Print("\nEnter phone number of customer you want to delete : ")
						fmt.Scanln(&noHp)
						res, err := custAuth.DeleteCust(noHp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Delete Customer success")
						} else {
							fmt.Println("Delete Customer  failed")
						}
					case 5:
						//Delete Transaction
						fmt.Print("\n--- Delete Transaction\n")
						listTransaction := transAuth.GetTransaction()
						if len(listTransaction) == 0 {
							fmt.Println("Empty history transaction")
							continue
						}
						fmt.Println("\n--- Delete Transaction")
						fmt.Println("Transaction id\t\tEmployee Name\t\tCutomer Name\t\tTransaction Date")
						for _, v := range listTransaction {
							fmt.Println(v.GetIdTrans(), "\t\t", v.GetEmpName(), "\t\t", v.GetCustName(), "\t\t", v.GetCreateDate())
						}
						var idtrans int
						fmt.Print("Enter transaction ID that you want to delete : ")
						fmt.Scanln(&idtrans)
						if _, ok := listTransaction[idtrans]; ok {
							res, err := transAuth.DeleteTrans(idtrans)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Delete Transaction success")
							} else {
								fmt.Println("Delete Transaction  failed")
							}

						} else {
							fmt.Println("Transaction id not exist")
						}

					case 9:
						continue
					default:
						fmt.Println("\nSorry, option doesn't exist")
					}
				}
			} else {
				emp, err := employeeAuth.Login(username, password)
				if err != nil {
					fmt.Println(err.Error())
				}

				//Menu Utama Pegawai
				if emp.GetName() != "" {
					menuEmp := 0
					for menuEmp != 9 {
						fmt.Print("\n--- Employee Menu\n")
						fmt.Print("\nWelcome back, ", emp.GetName(), "\n")
						fmt.Println("1. Insert item")
						fmt.Println("2. Show All items")
						fmt.Println("3. Edit item")
						fmt.Println("4. Update quantity")
						fmt.Println("5. Add new customer")
						fmt.Println("6. Transaction")
						fmt.Println("9. Log out")
						fmt.Print("Enter an option : ")
						fmt.Scanln(&menuEmp)

						switch menuEmp {
						case 1:
							fmt.Print("\n--- Insert Item\n")
							var newItem item.Item
							fmt.Print("Insert item name\t: ")
							scanner.Scan()
							tmps := scanner.Text()
							newItem.SetItemName(tmps)
							fmt.Print("Insert quantity\t\t: ")
							tmpi := 0
							fmt.Scanln(&tmpi)
							newItem.SetQuantity(tmpi)
							newItem.SetIdEmployee(emp.GetId())
							res, err := itemAuth.InsertItem(newItem)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Insert item Success")

							} else {
								fmt.Println("Insert item failed")
							}
						case 2:
							fmt.Print("\n--- List All Items\n")
							itemAuth.ShowItems()
						case 3:
							fmt.Print("\n--- Edit Item\n")
							itemAuth.ShowItems()
							var idItem int
							var newName string
							fmt.Print("Insert Id item\t: ")
							fmt.Scanln(&idItem)
							fmt.Print("Insert New Name\t: ")
							fmt.Scanln(&newName)
							res, err := itemAuth.EditItems(idItem, newName)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Update Name Success")

							} else {
								fmt.Println("Update Name failed")
							}

						case 4:
							fmt.Print("\n--- Edit Quantity\n")
							itemAuth.ShowItems()
							var idItem, qty int
							fmt.Print("Insert Id item\t\t: ")
							fmt.Scanln(&idItem)
							fmt.Print("Insert New Quantity\t: ")
							fmt.Scanln(&qty)
							res, err := itemAuth.UpdateQty(idItem, qty)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Update Quantity Success")

							} else {
								fmt.Println("Update Quantity failed")
							}
						case 5:
							fmt.Print("\n--- Add New Customer\n")
							var newCust customer.Customer
							tmps := ""
							fmt.Print("Insert Phone number\t: ")
							fmt.Scanln(&tmps)
							newCust.SetNohp(tmps)
							fmt.Print("Insert Customer Name\t: ")
							scanner.Scan()
							tmps = scanner.Text()
							newCust.SetName(tmps)
							newCust.SetIdEmployee(emp.GetId())
							res, err := custAuth.InsertCust(newCust)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Insert New Customer Success")

							} else {
								fmt.Println("Insert New Customer failed")
							}
						case 6:
							//Menu Transaksi
							transMenu := 0
							cart := map[int]*item.Item{}
							items := itemAuth.ItemList()

							for transMenu != 9 {
								fmt.Println("\n--- Transaction Menu")
								fmt.Println("1. Show all items")
								fmt.Println("2. Add item to cart")
								fmt.Println("3. Show cart")
								fmt.Println("4. Checkout")
								fmt.Println("5. Print receipt")
								fmt.Println("9. Back")
								fmt.Print("Enter an option : ")
								fmt.Scanln(&transMenu)
								switch transMenu {
								case 1:
									fmt.Print("\n--- All items\n")
									itemAuth.ShowItems()
								case 2:
									//show  item avaiable
									fmt.Print("\n--- Add item to cart\n")
									var itemCart item.Item
									idItem, qty := 0, 0
									fmt.Println("Item Id\t\tItem Name\tQuantity")
									for _, v := range items {
										if v.GetQuantity() > 0 {
											fmt.Println(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t", v.GetQuantity())
										}
									}

									//proses masukin ke keranjang
									fmt.Print("\nInsert ID item\t: ")
									fmt.Scanln(&idItem)

									fmt.Print("Insert quantity\t: ")
									fmt.Scanln(&qty)
									if _, ok := items[idItem]; !ok {
										fmt.Println("\nThe item you choose doesn't exist")
										continue
									}
									if qty > items[idItem].GetQuantity() {
										fmt.Println("\nThe quantity you've entered is higher than the maximum available")
										continue
									}

									// ngecek apakah inputan ada di cart
									//kondisi true
									if thisItem, ok := cart[idItem]; ok {
										thisItem.SetIdItem(idItem)
										thisItem.SetIdEmployee(items[idItem].GetIDEmployee())
										thisItem.SetItemName(items[idItem].GetItemName())
										thisItem.SetQuantity(thisItem.GetQuantity() + qty)
										cart[idItem] = thisItem
									} else { // kondisi false barang belum ada di cart
										itemCart.SetIdItem(idItem)
										itemCart.SetIdEmployee(items[idItem].GetIDEmployee())
										itemCart.SetItemName(items[idItem].GetItemName())
										itemCart.SetQuantity(qty)
										cart[idItem] = &itemCart
									}
									// update di show item ketika mau masukin keranjang
									items[idItem].SetQuantity(items[idItem].GetQuantity() - qty)
								case 3:
									// show isi cart
									fmt.Println("\n---Cart")
									if len(cart) == 0 {
										fmt.Println("\nYour cart empty, please add item to cart")
									} else {
										fmt.Print("Item ID\t\tItem Name\t\tQuantity\n")
										for _, v := range cart {
											fmt.Print(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t\t", v.GetQuantity(), "\n")
										}
									}
								case 4:
									//Checkout
									fmt.Println("\n--- Cart")
									if len(cart) == 0 {
										fmt.Println("\nYour is cart empty, please add item to cart")
										continue
									}
									fmt.Print("Item ID\t\tItem Name\t\tQuantity\n")
									for _, v := range cart {
										fmt.Print(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t", v.GetQuantity(), "\n")
									}
									tmp, noHp := "", ""
									fmt.Print("Are you sure want to finish your cart? (y/n)\t: ")
									fmt.Scanln(&tmp)

									if tmp == "y" {
										fmt.Print("Insert customer phone number\t: ")
										fmt.Scanln(&noHp)
										res, err := transAuth.Checkout(emp.GetId(), noHp, cart)
										if err != nil {
											fmt.Println(err.Error())
										}
										if res {
											fmt.Println("\nCheckout Success")

										} else {
											fmt.Println("\nCheckout failed")
										}
										for _, v := range cart {
											itemAuth.UpdateQty(v.GetIdItem(), items[v.GetIdItem()].GetQuantity())
										}
									} else if tmp == "n" {
										fmt.Println("are you okay ?")
										continue
									} else {
										fmt.Println("Wrong Input")
									}
								case 5:
									//Print Receipt
									fmt.Println("\n--- Transaction List")
									listTrans := transAuth.GetTransaction()
									if len(listTrans) == 0 {
										fmt.Println("No has record transaction")
									} else {
										fmt.Println("ID\t\tEmployee Name\t\tCustomer Name\t\tTransaction Date")
									}
									for _, v := range listTrans {
										fmt.Println(v.GetIdTrans(), "\t", v.GetEmpName(), "\t\t", v.GetCustName(), "\t\t", v.GetCreateDate())
									}
									idTrans := 0
									fmt.Print("\nEnter transaction id that you want to print : ")
									fmt.Scanln(&idTrans)
									fmt.Println("\n--- Transaction Details")
									fmt.Println("Transaction ID\t\t: ", listTrans[idTrans].GetIdTrans())
									fmt.Println("Employee Name\t\t: ", listTrans[idTrans].GetEmpName())
									fmt.Println("Customer Name\t\t: ", listTrans[idTrans].GetCustName())
									fmt.Println("Transaction Date\t: ", listTrans[idTrans].GetCreateDate())
									fmt.Println()

									listItems := transAuth.GetItemsTransaction(idTrans)
									fmt.Println("Item ID\t\t   Item Name\t\t      Quantity")
									for _, v := range listItems {
										fmt.Println(v.GetIdItem(), "\t\t", v.GetItemName(), "\t\t", v.GetQty())
									}
								case 9:
									continue
								default:
									fmt.Println("\nSorry, option doesn't exist")
								}
							}
						case 9:
							continue
						default:
							fmt.Println("\nSorry, option doesn't exist")
						}
					}
				}
			}
		case 2:
			fmt.Println("\nAplikasi Tokoku v.1.0")

		case 9:
			fmt.Println("\nThank you, see you later ...")
			continue
		default:
			fmt.Println("\nSorry, option doesn't exist")
		}
	}
}
