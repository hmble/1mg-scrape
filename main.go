package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

	type medicinesDetails struct {
		Name string `selector:"p.Card__productName__qM0Yq.bodyMedium"`
		Price string `selector:"span.l3Regular"`
		CompanyName string
		Contents string
	}
func main() {
	fName := "medicineDetails.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Price", "CompanyName", "Contents"})
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("1mg.com", "www.1mg.com"),
		
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./1mg_cache"),
		colly.Async(true),
	)
// 	c.Limit(&colly.LimitRule{
// 	RandomDelay: 2 * time.Second,
// 	Parallelism: 4,
// })

// proxySwitcher, err := proxy.RoundRobinProxySwitcher(
// 	"socks5://185.189.199.75:23500",
// "socks5://185.61.92.207:43947",
// "socks5://78.60.203.75:47385" ,
// "socks5://211.24.95.49:47615" ,
// "socks5://43.248.24.158:51166"	 ,
// "socks5://176.62.178.247:47556",
// "socks5://198.58.11.20:55443",
// "socks5://183.87.153.98:49602"	 ,
// "socks5://117.2.165.159:53281" ,
// "socks5://109.232.106.236:49565",
// "socks5://95.104.54.227:42119",
// "socks5://103.111.225.17:8080",
// "socks5://185.128.136.244:3128",
// )
// if err != nil {
//   log.Fatal(err)
// }
// c.SetProxyFunc(proxySwitcher)
// extensions.RandomUserAgent(c)
// 	q, _ := queue.New(
// 			2, // Number of consumer threads
// 			&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
// 		)
// 		c.OnRequest(func(r *colly.Request) {
// 			log.Println("visiting", r.URL.String())
			
// 		})
// c.OnResponse(func(r *colly.Response) {
// 		log.Println(r.StatusCode)
// 	})
		medicines := make([]*medicinesDetails, 0, 20)
	c.OnHTML("div.Card__container__35Kl4", func(e *colly.HTMLElement) {
			medicine := &medicinesDetails{}
			e.Unmarshal(medicine)	
			fullText := e.ChildText("p.Card__productDescription__2HHjt")
			if strings.Contains(fullText, "Prescription") {
				medicine.CompanyName = e.ChildText("p.Card__productDescription__2HHjt:nth-child(4)")
				medicine.Contents = e.ChildText("p.Card__productDescription__2HHjt:nth-child(5)")
				}else {
					medicine.CompanyName = e.ChildText("p.Card__productDescription__2HHjt:nth-child(3)")
					medicine.Contents = e.ChildText("p.Card__productDescription__2HHjt:nth-child(4)")
					
			}
			medicines = append(medicines, medicine)
			fmt.Printf("%#v", medicine)
			writer.Write([]string{
				medicine.Name,
				medicine.Price,
				medicine.CompanyName,
				medicine.Contents,
			})
	})

	url := "https://www.1mg.com/drugs-all-medicines"

	c.Visit(url)
	// for i := 1; i <= 10; i++ {
	// 	// Add URLs to the queue
	// 	q.AddURL(fmt.Sprintf("%s?page=%d", url, i))
	// }

	// q.Run(c)

	// c.Wait()

	// fmt.Println("Total length of medicines ", len(medicines))

}