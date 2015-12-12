Jsorm (Javascript Object Rest Message)
---------
This is a very simple library for sending server messages and errors.

#Usage:
```
		const ServerBusy = 13

		func handler(writer http.ResponseWriter, request *http.Request) {
			if !UserIsLoggedIn(request) {
				jsorm.AuthError(writer, "You are not logged in!")
			} 
			
			if ServerIsTooBusyAtTheMoment() {
				jsorm.Code(writer, ServerBusy, "The Server is way too busy right now!")
				return
			}
		
			if output, err := MyFunction(...); err != nil {
				jsorm.Error(writer, err.Error())
			} else {
				jsorm.Send(writer, output)
			}
		}
```
