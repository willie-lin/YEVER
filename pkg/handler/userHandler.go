package handler

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/willie-lin/YEVER/pkg/database/ent"
	"github.com/willie-lin/YEVER/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Controller struct {
	Client *ent.Client
}

func CreateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user ent.User

		// bind request to user struct
		//if err := c.Bind(&user); err != nil {
		//	c.Logger().Error("Bind:", err)
		//	return c.String(http.StatusBadRequest, "Bind: "+err.Error())
		//}

		// 接受json数组并解析绑定到user
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return err
		}

		newPsd, err := utils.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			//fmt.Println("加密密码失败", err)
			return err
		}

		cc := client.User.Create().SetID(utils.UUID())

		if user.Name != "" {
			cc.SetName(user.Name)
		} else if user.Age >= 0 {
			cc.SetAge(user.Age)

		} else if user.Password != "" {
			cc.SetName(string(newPsd))
		}

		// insert record

		if user.Name != "" {
			cc.SetName(user.Name)
		}
		if user.Password != "" {
			cc.SetPassword(user.Password)

		}
		if user.Email != "" {
			cc.SetEmail(user.Email)
		}
		if user.Age >= 0 {
			cc.SetAge(user.Age)

		}
		if user.ID != "" {
			cc.SetID(utils.UUID())

		}

		newUser, err := cc.Save(context.Background())
		if err != nil {
			c.Logger().Error("Insert: ", err)
			return c.String(http.StatusBadRequest, "Save: "+err.Error())
		}

		c.Logger().Infof("inserted comment: %v", newUser.ID)
		return c.NoContent(http.StatusCreated)

	}
}

// 根据用户名查找
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

	pwd, err := utils.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

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
		cc.SetPassword(string(pwd))
	}

	newUser, err := cc.Save(context.Background())
	if err != nil {
		c.Logger().Error("Insert: ", err)
		return c.String(http.StatusBadRequest, "Save: "+err.Error())
	}

	c.Logger().Infof("inserted comment: %v", newUser.ID)
	return c.NoContent(http.StatusCreated)
}
