package blockone


//-- Imported libraries ---------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

import (        "strconv"
		"bytes"
	        "net/url"
		"net/http"
		"encoding/json"
		"fmt" )

//-- Class-simulation structure -------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

type Bloco1 struct {
	resourceURL 	string;
	resourceIP		string;
	resourceToken 	string;
}


//--------- Function Factory : returns a new Bloco1 Object ----------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func NewClient(rurl string, resIP string) * Bloco1 {

	tvar := new(Bloco1);

	tvar.resourceURL = rurl;
	tvar.resourceIP = resIP;

	return tvar;
}


//--------- Function Login: returns an Object in case of success, or an Error object on failure ----
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func (b *Bloco1) Login(email string, password string, tokens ...string ) (interface{}){ 

	 
 	message := map[string]interface{}{
		"email": email, 
		"password": password,
		"tokens": tokens 	}
	
	
		buf := new(bytes.Buffer);
		json.NewEncoder(buf).Encode(message);

    u, _ := url.ParseRequestURI( b.resourceURL);
    u.Path ="auth/login";
	urlStr := u.String(); 
	
	var response interface{}


    client := &http.Client{}
    r, err := http.NewRequest("POST", urlStr, buf ); 
	if err != nil {
		
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}
	 
    r.Header.Add("Content-Type", "application/json");
   
	resp, err := client.Do(r)
		if err != nil {
		
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close();

	   
	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;
	
}

//----- function refresh: refreshes the token received after login event.
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func (b *Bloco1) Refresh(investor int, email string, token string) (interface{}){


	message := map[string]interface{}{
		"investor": investor,
		"email": email	}
	
	
		buf := new(bytes.Buffer);
		json.NewEncoder(buf).Encode(message);

    u, _ := url.ParseRequestURI( b.resourceURL);
    u.Path ="auth/refreshtoken"
	urlStr := u.String() 

    client := &http.Client{}
    r, err := http.NewRequest("POST", urlStr, buf ) 
	if err != nil {
		fmt.Printf("Ooops");
	}

	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json")

	var response interface{}
	
	resp, err := client.Do(r)
		if err != nil {
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {

		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;

 }


 //----- function sendMarket: sends market actions to the server     ------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func (b *Bloco1) SendMarket(investor int, book int, isSale bool, requiredValue float32, 
												  requiredAmount float32, token string) (interface{}) { 


	message := map[string]interface{}{
		"investor": investor,
		"book": book,
		"is_sale":isSale,
		"required_value": requiredValue,
		"required_amount": requiredAmount	}
	
	
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(message)

    u, _ := url.ParseRequestURI( b.resourceURL)
    u.Path ="trading/sendmarket"
	urlStr := u.String()  

	var response interface{}

    client := &http.Client{}
    r, err := http.NewRequest("POST", urlStr, buf )  
	if err != nil { 

		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json")
     

	resp, err := client.Do(r)
		if err != nil { 

		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 
 

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
 
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;


}



 //----- function sendLimit: sends limiting parameters to the server    ---------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------


func (b *Bloco1) SendLimit(hidden bool, limitVal int, isOco bool, ocoVal int, investor int, 
				book int, isSale bool, reqVal  float32, reqAmount float64, token string)  (interface{}){ 


	ocobj := map[string]interface{}{
		"is_oco": isOco, 
		"oco_value": ocoVal	}

	message := map[string]interface{}{
		"is_hidden": hidden,
		"limit_value":limitVal,
		"oco":ocobj,
		"investor": investor,
		"book": book,
		"is_sale":isSale,
		"required_value": reqVal,
		"required_amount": reqAmount	}


		var response interface{}

		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(message)

		u, _ := url.ParseRequestURI( b.resourceURL)
		u.Path ="trading/sendlimit"
		urlStr := u.String() // "https://api.com/user/"
	
		client := &http.Client{}
		r, err := http.NewRequest("POST", urlStr, buf ) // URL-encoded payload
		if err != nil {
			//log.Fatalln(err)
			//fmt.Printf("Ooops");
			json.Unmarshal([]byte(`{"error":true}`), &response) 
			return response;
		}

r.Header.Add("Authorization", "Bearer "+token)
r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
		if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 
	//fmt.Printf(newStr)

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("ERROR!");
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;
}




 //----- function cancelOrder: cancels peviously put order           ------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------


func (b *Bloco1) CancelOrder(investor int, orderID int, token string) (interface{}){


	message := map[string]interface{}{
		"investor": investor,
		"orderid":  orderID	}


		var response interface{}

		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(message)

		u, _ := url.ParseRequestURI( b.resourceURL)
		u.Path ="trading/cancelorder"
		urlStr := u.String() // "https://api.com/user/"
	
		client := &http.Client{}
		r, err := http.NewRequest("POST", urlStr, buf ) // URL-encoded payload
		if err != nil {
			//log.Fatalln(err)
			//fmt.Printf("Ooops");
			json.Unmarshal([]byte(`{"error":true}`), &response) 
			return response;
		}

r.Header.Add("Authorization", "Bearer "+token)
r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
		if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 
	//fmt.Printf(newStr)

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("ERROR!");
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;




	}




 //----- function myOrders: list orders from a specific investor     ------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func (b *Bloco1) MyOrders(investor int, isSale bool, filtroStatus int, page int, 
															rows int, token string) (interface{}) { 

	message := map[string]interface{}{
		"investor": investor,
		"isSale":  isSale,
		"filtroStatus": filtroStatus,
		"page":	page,
		"rows": rows}


		var response interface{}

		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(message)

		u, _ := url.ParseRequestURI( b.resourceURL)
		u.Path ="trading/myorders"
		urlStr := u.String() // "https://api.com/user/"
	
		client := &http.Client{}
		r, err := http.NewRequest("POST", urlStr, buf ) // URL-encoded payload
		if err != nil {
			//log.Fatalln(err)
			//fmt.Printf("Ooops");
			json.Unmarshal([]byte(`{"error":true}`), &response) 
			return response;
		}

r.Header.Add("Authorization", "Bearer "+token)
r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
		if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 
	//fmt.Printf(newStr)

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("ERROR!");
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;

}



 //----- function balance: show balance from a specific investor     ------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------


func (b *Bloco1) Balance(investor int, token string)(interface{})  {
	

	message := map[string]interface{}{
		"investor": investor }


		var response interface{}

		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(message)

		u, _ := url.ParseRequestURI( b.resourceURL)
		u.Path ="trading/myorders"
		urlStr := u.String() // "https://api.com/user/"
	
		client := &http.Client{}
		r, err := http.NewRequest("POST", urlStr, buf ) // URL-encoded payload
		if err != nil {
			//log.Fatalln(err)
			//fmt.Printf("Ooops");
			json.Unmarshal([]byte(`{"error":true}`), &response) 
			return response;
		}

r.Header.Add("Authorization", "Bearer "+token)
r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
		if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 
	//fmt.Printf(newStr)

	err = json.Unmarshal([]byte(newStr), &response)
	if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("ERROR!");
		json.Unmarshal([]byte(`{"error":true}`), &response) 
	}

	
	return response;

}



 //----- function queryBook: list last n booking events   -----------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------


func (b *Bloco1) QueryBook(count int) (interface{}){	

	var response interface{}

		buf := new(bytes.Buffer);

		u, _ := url.ParseRequestURI( "https://market-api.bloco1.com.br")
		u.Path ="MarketData/query_book/" + strconv.Itoa(count)
		urlStr := u.String() // "https://api.com/user/"
	
		client := &http.Client{}
		r, err := http.NewRequest("GET", urlStr, buf ) // URL-encoded payload
		if err != nil {
			//log.Fatalln(err)
			//fmt.Printf("Ooops");
			json.Unmarshal([]byte(`{"error":true}`), &response) 
			return response;
		}

 
r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
		if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response);
		return response;
	}

	defer resp.Body.Close()

	

	buf2 := new(bytes.Buffer)
    buf2.ReadFrom(resp.Body)
    newStr := buf2.String()
 


	
	return newStr;

}




 //----- function queryBusiness: list last N concretized business       ---------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------

func (b *Bloco1) QueryBusiness(count int)  (interface{}){ 


	var response interface{}

	buf := new(bytes.Buffer);

	u, _ := url.ParseRequestURI( "https://market-api.bloco1.com.br")
	u.Path ="MarketData/query_business/" + strconv.Itoa(count)
	urlStr := u.String() // "https://api.com/user/"


	fmt.Printf(urlStr);

	client := &http.Client{}
	r, err := http.NewRequest("GET", urlStr, buf ) // URL-encoded payload
	if err != nil {
		//log.Fatalln(err)
		//fmt.Printf("Ooops");
		json.Unmarshal([]byte(`{"error":true}`), &response) 
		return response;
	}


r.Header.Add("Content-Type", "application/json")

resp, err := client.Do(r)
	if err != nil {
	//log.Fatalln(err)
	//fmt.Printf("Ooops");
	json.Unmarshal([]byte(`{"error":true}`), &response);
	return response;
}

defer resp.Body.Close()

buf2 := new(bytes.Buffer)
buf2.ReadFrom(resp.Body)
newStr := buf2.String()

return newStr;


}




 //----- Test implementation: ---------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------
/*
func main() {

    Bloco := newBloco1("https://trade-api.bloco1.com.br", "1.1.1.1");

	fmt.Printf("%+v\n",  Bloco.login("eu", "string", "123456", "123456", "123456")   );
	fmt.Printf("%+v\n",  Bloco.refresh(1, "eu", "123456")   );
	fmt.Printf("%+v\n",  Bloco.sendMarket(1, 1, true, 10, 33.44, "123456")   );
	fmt.Printf("%+v\n",  Bloco.sendLimit(false, 1, true, 12, 1, 1, true, 33, 12, "123456")   );
	fmt.Printf("%+v\n",  Bloco.cancelOrder(1 , 3, "123456") );
	fmt.Printf("%+v\n",  Bloco.myOrders(1, true, 2, 1, 50, "123456") );
	fmt.Printf("%+v\n",  Bloco.balance(1, "123456") );
	fmt.Printf("%+v\n",  Bloco.queryBook(50) );
	fmt.Printf("%+v\n",  Bloco.queryBusiness(100) );
}
*/

