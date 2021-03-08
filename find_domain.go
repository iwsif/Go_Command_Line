
package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"github.com/spf13/cobra"
	"net"

)

func find_the_ip(host string) error {
	result,err := net.LookupHost(host)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		for i:=0;i<len(result);i++ {
			fmt.Println(result)
		}
	}()
	return err
}

func ns_records(domain_name string)  bool {
	var res bool
	result,err := net.LookupNS(domain_name)
	if err!=nil {
		fmt.Println(err)
		res = false
		return res
	}
	defer func() {
		for _,v := range result {
			fmt.Println(v)
		}		
		res = true
	}()
	return res
}

func mx_records(domain_name string) bool {
	var res bool
	result, err := net.LookupMX(domain_name)
	if err!=nil {
		fmt.Println(err)
		res = false 
		return res
	}
	
	defer func() {
		for i:=0;i<len(result);i++ {
			fmt.Println(result)
		}
		res = true
	}()
	return res
}



func main() {
	log.Print()
	fmt.Println("Welcome this is simple program written in golang with cobra-package")
	fmt.Println("Starting...")
	
	time.Sleep(time.Duration(time.Millisecond*1000))
	
	command := &cobra.Command {
		Short:"Find ip",
		Long:"Find ip from a given domain_name",
		Run:func(com *cobra.Command,thing[]string) {

			if  os.Args[1] == "--domain" || os.Args[1] == "-d" {
				if len(os.Args) < 1 {
					fmt.Println("Error..")
					fmt.Println("Exiting..")
					os.Exit(0)
				}else {
					result := find_the_ip(os.Args[2])
					if result!=nil {
						fmt.Println("Exiting...")
						time.Sleep(time.Duration(time.Millisecond*2000))
					}
				}
			}else if os.Args[1] == "--dsrecords" || os.Args[1] != "-r" {
				if len(os.Args) < 1 {
					fmt.Println("Error..")
					time.Sleep(time.Duration(time.Millisecond*2000))
					os.Exit(0)
				}else {
					final_result := ns_records(os.Args[2])
					if final_result == false {
						fmt.Println("Exiting..")
						time.Sleep(time.Duration(time.Millisecond*2000))
						os.Exit(0)
					}
				}				
			}else if os.Args[1] == "--mxrecords" || os.Args[1] == "-m" {
				if len(os.Args) < 1 {
					fmt.Println("Error")
					os.Exit(0)
				}else {
					check_bool := mx_records(os.Args[2])
					if check_bool == false {
						fmt.Println("Error")
						time.Sleep(time.Duration(time.Millisecond*2000))
						os.Exit(0)
					}
				}
			}
		},
	}
	command.Flags().StringP("domain","d","","Type the domain_name you want to find the ip")
	command.Flags().StringP("dsrecords","r","","Check the ds_records for your domain name")
	command.Flags().StringP("mxrecords","m","","Check the mx_records for your domain")
	if len(os.Args) > 1 {
		command.Execute()
	} else {
		fmt.Println("Error...")
		fmt.Println("Bye bye...")
		time.Sleep(time.Duration(time.Millisecond*3000))
	}
}
