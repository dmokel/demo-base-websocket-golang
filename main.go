package main

import chatexample "github.com/dmokel/demo-base-websocket-golang/chat-example"

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type handlerForSome struct{}

// func (h handlerForSome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}

// 	for {

// 	}
// }

func main() {
	chatexample.CoreLogic()
	// http.Handle("/handle", &handlerForSome{})
	// http.HandleFunc("/handle-func", func(w http.ResponseWriter, r *http.Request) {
	// 	conn, err := upgrader.Upgrade(w, r, nil)
	// 	if err != nil {
	// 		fmt.Println("err:", err)
	// 		return
	// 	}

	// 	for {

	// 	}
	// })

	// http.ListenAndServe("0.0.0.0", nil)
}
