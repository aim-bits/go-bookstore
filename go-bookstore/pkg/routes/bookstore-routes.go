package routes



var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}