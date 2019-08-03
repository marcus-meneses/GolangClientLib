package main


import ( "fmt" 
	  blocone "../bloco1")

 

func main() {

    Bloco := blocone.NewClient("https://trade-api.bloco1.com.br", "1.1.1.1");

	fmt.Printf("%+v\n",  Bloco.Login("eu", "string", "123456", "123456", "123456")   );
	fmt.Printf("%+v\n",  Bloco.Refresh(1, "eu", "123456")   );
	fmt.Printf("%+v\n",  Bloco.SendMarket(1, 1, true, 10, 33.44, "123456")   );
	fmt.Printf("%+v\n",  Bloco.SendLimit(false, 1, true, 12, 1, 1, true, 33, 12, "123456")   );
	fmt.Printf("%+v\n",  Bloco.CancelOrder(1 , 3, "123456") );
	fmt.Printf("%+v\n",  Bloco.MyOrders(1, true, 2, 1, 50, "123456") );
	fmt.Printf("%+v\n",  Bloco.Balance(1, "123456") );
	fmt.Printf("%+v\n",  Bloco.QueryBook(50) );
	fmt.Printf("%+v\n",  Bloco.QueryBusiness(100) );
}

