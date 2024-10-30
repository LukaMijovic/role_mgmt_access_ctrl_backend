package services

import (
	"errors"
	"fmt"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ConfirmCreationByAdmin(u *dto.UserCredentialsDTO, ctx *gin.Context) error {
	fmt.Println("Entered Goroutine!")

	// wsh := &util.WebSocketHandler{
	// 	Upgrader: &websocket.Upgrader{
	// 		ReadBufferSize:  1024,
	// 		WriteBufferSize: 1024,
	// 	},
	// }

	//conn, err := wsh.Connect(ctx)

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err.Error())

	// 	return err
	// }

	//defer wsh.Disconnect(conn)

	// send signal

	//msg := fmt.Sprintf("%v", uint8(u.User_ID))
	//err = util.WebAppConnection.WriteMessage(websocket.TextMessage, []byte(msg))
	err := util.WebAppConnection.WriteJSON(&dto.UserCredentialConfirmationDTO{Email: u.Email, UserId: u.User_ID})

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	var res dto.UserRoleDTO
	err = util.WebAppConnection.ReadJSON(&res)
	//mt, data, err := util.WebAppConnection.ReadMessage()

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	//fmt.Printf("User_id: %v, role_id: %v\n", res.User_id, res.Role_id)

	ur := repository.NewUserRepository()
	err = ur.SetRoleIdOfUser(res.User_id, res.Role_id)

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	//err = ur.SaveUserCredentials(u)
	err = RegisterUserToDatabase(u)

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	//fmt.Printf("MT: %v, msg: %s\n", mt == websocket.BinaryMessage, string(data))
	//fmt.Printf("User %v got role: %v\n", res.User_id, res.Role_id)

	//msg = fmt.Sprintf("User %v has been created with role %v.", uint8(res.User_id), uint8(res.Role_id))
	//conn.WriteMessage(websocket.TextMessage, []byte(data))

	err = util.MobileAppConnection.WriteJSON(&u)

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	err = util.WebAppConnection.WriteMessage(websocket.TextMessage, []byte("Successful"))

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())

		return err
	}

	return nil
}

func ConfirmTempRightByAdmin(tr *model.TempAccessRight) error {
	//send signal
	return nil
}

func SignInAdmin(a *dto.AdminCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	userRepository := repository.NewUserRepository()

	u := dto.UserCredentialsDTO{
		Email:    a.Email,
		Password: a.Password,
	}

	err := userRepository.GetUserIDFromDataBase(&u)

	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository()
	creds, err := adminRepository.GetAdminCredentials(&u)

	if err != nil {
		return nil, err
	}

	if a.Email == creds.Email && util.CheckPassword(u.Password, creds.Password) {
		return creds, nil
	}

	return nil, errors.New("Login failed. Bad credentials.")
}
