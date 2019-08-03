Métodos implementados para  classe blockone:


NewClient(rurl string, resIP string) * Bloco1

Login(email string, password string, tokens ...string )

Refresh(investor int, email string, token string)

SendMarket(investor int, book int, isSale bool, requiredValue float32,
			          requiredAmount float32, token string) 

SendLimit(hidden bool, limitVal int, isOco bool, ocoVal int, investor int, 
	book int, isSale bool, reqVal  float32, reqAmount float64, token string)  



CancelOrder(investor int, orderID int, token string)


MyOrders(investor int, isSale bool, filtroStatus int, page int, rows int, token string) 


Balance(investor int, token string)


QueryBook(count int)


QueryBusiness(count int)

