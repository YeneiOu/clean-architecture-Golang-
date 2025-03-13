package servers

import (
	_usersHttp "bank/modules/users/controllers"
	_usersRepository "bank/modules/users/respositories"
	_usersUsecase "bank/modules/users/usecase"
	"bank/pkg/utils"
	//"fmt"
)

func (s *Server) MapHandler() error {
	s.App.Use(utils.CORS)

	v1 := s.App.Group("/api/v1")

	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUserRepository(s.Db)
	usersUsecase := _usersUsecase.NewUserUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	//_ = usersRepository
	//usersRepositoryMockDB := _usersMockRepository.NewUserRepositoryMock()
	//userService := _userService.NewUserService(usersRepositoryMockDB)

	//userServiceMock, err := userService.GetUsers()
	//if err != nil {
	//	fmt.Println("Error fetching userServiceMock:", err)
	//	return err
	//}
	//fmt.Println("userServiceMock:", userServiceMock)
	//
	//userMockResult, err := usersUsecase.GetUsers()
	//if err != nil {
	//	fmt.Println("Error fetching userMockResult:", err)
	//	return err
	//}
	//fmt.Println("userMockResult:", userMockResult)

	//userService := _userService.NewUserService(userMockDB)
	//users, err := userService.GetUsers()
	//if err != nil {
	//	fmt.Println("Error fetching users:", err)
	//	return err
	//}
	//fmt.Println("Users:", users)
	//fmt.Println("usersGroup:", usersGroup)

	return nil
}
