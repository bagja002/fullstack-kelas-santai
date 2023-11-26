package route

import (
	"backend1/models"
	"backend1/database"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

//membuat route untuk login
func Login(c *fiber.Ctx)error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//ambil data user di dalam database dan di cek email nya ada atau tidak
	var user models.User
	//pengambilan data dengan gorm
	database.DB.Where("username = ?", data["username"]).Find(&user)
	//pengecekan jika tidak ada userdi database
	if user.Id_user == 0{
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"Pesan":"Username Tidak Ditemuka",
		})
	}
	//Pengecekan jika Password yang di database dan password yang ada di dalam postman

	if  err :=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) ; err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"Pesan":"Password salah",
		})
	}else{
		//membuat data JWT baru
		claim:= jwt.MapClaims{
			"id_user":user.Id_user,
			"nama":user.Nama,
			"username":user.Username,
		}
		//Membuat dataJWT yang baru
		token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		t, err := token.SignedString([]byte("secret"))
		if err!= nil{
			return c.JSON(fiber.Map{
				"pesan":"eror membuat token",
			})
		}
		//Menampilkan isi token
		return c.JSON(fiber.Map{
			"token":t,
		})
	}
}

func Register(c *fiber.Ctx)error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]),bcrypt.DefaultCost)
	user:= models.User{
		Username: data["username"],
		Password: string(password),
		Nama: data["nama"],
		Id_admin: 1,
		Is_active: true,
	}
	database.DB.Create(&user)
	
	return c.JSON(fiber.Map{
		"PESAN":"Registrasi Sudah Berhasil data berhasil di simpan di database",
		"data":user,
	})
}

func GetUSer(c *fiber.Ctx)error {
	id_user, _ := c.Locals("id_user").(int)
	//mengambil data user di database 
	var user models.User
	database.DB.Where("id_user = ?", id_user).Find(&user)

	if user.Id_user == 0{
		return c.JSON(fiber.Map{
			"Pesan":"Data USER TIDAK ADA",
		})
	}

	return c.JSON(fiber.Map{
		"data":user,
	})
}

func GetAlluser(c *fiber.Ctx)error {
	//id_user, _ := c.Locals("id_user").(int)

	id_user:= c.Query("id_user")
	var user []models.User

	if id_user ==""{
		database.DB.Find(&user)
	}else{
		database.DB.Where("id_user = ?", id_user).Find(&user)
	}
	return c.JSON(fiber.Map{
		"data":user,
	})
}

func Update(c *fiber.Ctx)error{

	id_user, _ := c.Locals("id_user").(int)
	//ambil data yang ada
	var user models.User
	database.DB.Where("id_user = ?", id_user).Find(&user)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//
	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]),bcrypt.DefaultCost)
	update := models.User{
		Nama:data["nama"] ,
		Username: data["username"],
		Password: string(password),
	}
	database.DB.Model(&user).Updates(update)

	return c.JSON(fiber.Map{
		"data_paling_baru":user,
	})
}