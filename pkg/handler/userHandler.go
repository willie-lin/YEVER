package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/willie-lin/YEVER/pkg/database/ent"
	"github.com/willie-lin/YEVER/pkg/database/ent/user"
	"github.com/willie-lin/YEVER/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"unsafe"
)

type Controller struct {
	Client *ent.Client
}

// CreateUser create user into user tables
func CreateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user *ent.User
		// 接受json数组并解析绑定到user
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return err
		}

		// 加密原始密码
		pwd, err := utils.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(pwd)
		cc := client.User.Create().SetID(utils.UUID())
		if user.Name != "" {
			cc.SetName(user.Name)
		}
		if user.Age != -1 {
			cc.SetAge(user.Age)
		}

		if user.Email != "" {
			cc.SetEmail(user.Email)
		}
		if user.Phone != "" {
			cc.SetPhone(user.Phone)
		}
		if user.Description != "" {
			cc.SetDescription(user.Description)
		}
		if user.Password != "" {
			//cc.SetPassword(user.Password)
			cc.SetPassword(*(*string)(unsafe.Pointer(&pwd)))
		}
		// insert record
		newUser, err := cc.Save(context.Background())
		if err != nil {
			c.Logger().Error("Insert: ", err)
			return c.String(http.StatusBadRequest, "Save: "+err.Error())
		}
		c.Logger().Infof("inserted comment: %v", newUser.ID)
		return c.NoContent(http.StatusCreated)
	}
}

// GetAllUser 获取所有用户
func GetAllUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		//var user ent.User
		log, _ := zap.NewDevelopment()
		users, err := client.User.Query().Order(ent.Desc(user.FieldCreated)).Limit(100).All(context.Background())
		//users, err := client.User.Query().All(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				log.Fatal("GetAll User Error: ", zap.Error(err))
				return c.String(http.StatusBadRequest, "Get: "+err.Error())
			}
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.JSON(http.StatusOK, users)
	}
}

// FindUserByName 根据用户名查找
func FindUserByName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u *ent.User
		// 接受json数组并解析绑定到user
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return err
		}
		us, err := client.User.Query().Where(user.NameEQ(u.Name)).Only(context.Background())
		if err != nil {
			if !ent.IsNotFound(err) {
				c.Logger().Error("Get: ", err)
				return c.String(http.StatusBadRequest, "Get: "+err.Error())
			}
			return c.String(http.StatusNotFound, "Not Found")
		}

		return c.JSON(http.StatusOK, us)
	}
}

// DeleteUser 删除用户
func DeleteUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		var u *ent.User

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			log.Fatal("json decode error: %v", zap.Error(err))
			return err
		}

		// 获取用户基本信息
		fmt.Println(u.Name)
		user, err := client.User.Query().Where(user.NameEQ(u.Name)).Only(context.Background())
		if err != nil {
			log.Fatal("Failed Querying User: %v", zap.Error(err))
			return fmt.Errorf("Failed Querying User: %v", err)
		}
		err = client.User.DeleteOne(user).Exec(context.Background())
		if err != nil {
			log.Fatal("Delete user error:", zap.Error(err))
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}

// UpdateUserByID  UpdateUser 更新用户
func UpdateUserByID(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u *ent.User

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			log.Fatal("json decode error: %v", zap.Error(err))
			return err
		}

		us, err := client.User.Query().Where(user.IDEQ(u.ID)).Only(context.Background())
		if err != nil {
			log.Fatal("Query user error:", zap.Error(err))
		}
		if u.Name != us.Name {
			log.Fatal("用户名不一致")

		}

		ur, err := client.User.Update().Where(user.IDEQ(u.ID)).Where(user.NameEQ(u.Name)).
			SetAge(u.Age).
			SetEmail(u.Email).
			SetPhone(u.Phone).
			SetDescription(u.Description).Save(context.Background())
		if err != nil {
			//panic(err)
			log.Fatal("Update  error:", zap.Error(err))
			//fmt.Println("update  err: ", err)
			return err
		}
		return c.JSON(http.StatusOK, &ur)
	}
}

func UpdateUserByName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u *ent.User

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			log.Fatal("json decode error: %v", zap.Error(err))
			return err
		}

		ur, err := client.User.Update().Where(user.NameEQ(u.Name)).
			SetAge(u.Age).
			SetEmail(u.Email).
			SetPhone(u.Phone).
			SetDescription(u.Description).Save(context.Background())
		if err != nil {
			//panic(err)
			log.Fatal("Update  error:", zap.Error(err))
			//fmt.Println("update  err: ", err)
			return err
		}

		return c.JSON(http.StatusOK, &ur)
	}
}

//func  (controller *Controller) FindUserByUsername(c echo.Context) error {
//		//client, err := database.Client()
//		//client, err := config.NewClient()
//		//if err != nil {
//		//	return err
//
//		u := new(ent.User)
//
//		//// 接收raw数据
//		//result, err := ioutil.ReadAll(c.Request().Body)
//		//if err != nil {
//		//	fmt.Println("ioutil.ReadAll err:", err)
//		//	return err
//		//}
//		//// 解析raw为json
//		//err = json.Unmarshal(result, &u)
//		//if err != nil {
//		//	fmt.Println("json.Unmarshal err:", err)
//		//	return err
//		//}
//
//		//fmt.Println(u.Username)
//		// 直接解析raw数据为json
//		log, _ := zap.NewDevelopment()
//		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
//			log.Fatal("json decode error", zap.Error(err))
//			return err
//		}
//		//// or for DisallowUnknownFields() wrapped in your custom func
//		//decoder := json.NewDecoder(c.Request().Body)
//		//decoder.DisallowUnknownFields()
//		//if err := decoder.Decode(&payload); err != nil {
//		//	return err
//		//}
//
//		us, err := client.User.Query().Where(user.NameEQ(u.Name)).Only(context.Background())
//		if err != nil {
//			return err
//		}
//		return c.JSON(http.StatusOK, &us)
//	}
//}

func (controller *Controller) InsertComment(c echo.Context) error {
	var user ent.User

	// bind request to user struct
	if err := c.Bind(&user); err != nil {
		c.Logger().Error("Bind:", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}

	// insert record

	cc := controller.Client.User.Create().SetDescription(user.Description)

	// 加密原始密码
	pwd, err := utils.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println(pwd)

	user.Password = string(pwd)

	//fmt.Println(string(pwd))

	if user.Name != "" {
		cc.SetName(user.Name)
	}
	//if user.Password != "" {
	//	cc.SetPassword(user.Password)
	//}
	if user.Email != "" {
		cc.SetEmail(user.Email)
	}
	if user.Age >= 0 {
		cc.SetAge(user.Age)

	}
	if user.ID != "" {
		cc.SetID(utils.UUID())
	}
	if user.Password != "" {

		cc.SetPassword(user.Password)
		//cc.SetPassword(*(*string)(unsafe.Pointer(&pwd)))
		//cc.SetPassword(string(pwd))
	}
	if user.Phone != "" {
		cc.SetPhone(user.Phone)

	}

	newUser, err := cc.Save(context.Background())
	//fmt.Println(newUser.Password)
	if err != nil {
		c.Logger().Error("Insert: ", err)
		return c.String(http.StatusBadRequest, "Save: "+err.Error())
	}

	c.Logger().Infof("inserted comment: %v", newUser.ID)
	return c.NoContent(http.StatusCreated)
}
