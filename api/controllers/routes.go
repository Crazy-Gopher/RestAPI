package controllers

func (s *Server) initializeRoutes() {

	//Users routes
	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users", s.GetUsers).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.GetUser).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
}
