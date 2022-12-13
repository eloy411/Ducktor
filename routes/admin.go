package routes

import (
	"net/http"
	"time"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

func SetUsersExamples(w http.ResponseWriter, r *http.Request) {

	u1 := models.User{
		IdUser:       "1",
		Nombre:       "Eloy",
		Password:     "1111",
		Edad:         12,
		Ciudad:       "Barcelona",
		Hospital:     "Hospital de Sant Pau",
		Enfermedad:   "Cancer",
		Confianza:    0,
		Gravedad:     3,
		Numtest:      0,
		Numdibujos:   0,
		Coins:        1000000,
		Fechaentrada: time.Now(),
	}
	u2 := models.User{
		IdUser:       "2",
		Nombre:       "Pablo",
		Password:     "1234",
		Edad:         8,
		Ciudad:       "Barcelona",
		Hospital:     "Hospital de Sant Pau",
		Enfermedad:   "Rotura de tibia",
		Confianza:    0,
		Gravedad:     0,
		Numtest:      0,
		Numdibujos:   0,
		Coins:        1000, 
		Fechaentrada: time.Now(), 
	}

	config.DB.Create(&u1)
	config.DB.Create(&u2)
}

func SetDibujos(w http.ResponseWriter, r *http.Request) {

	d1 := models.Dibujos{
		IdDibujo:    1,
		Tipo:        "casa",
		Descripcion: "Tienes que dibujar una casa",
	}

	d2 := models.Dibujos{
		IdDibujo:    2,
		Tipo:        "arbol",
		Descripcion: "Tienes que dibujar un arbol",
	}

	d3 := models.Dibujos{
		IdDibujo:    3,
		Tipo:        "amigo",
		Descripcion: "Tienes que dibujar a un amigo",
	}

	d4 := models.Dibujos{
		IdDibujo:    4,
		Tipo:        "familia",
		Descripcion: "Tienes que dibujar a tu familia",
	}

	d5 := models.Dibujos{
		IdDibujo:    5,
		Tipo:        "perro",
		Descripcion: "Tienes que dibujar a un perro",
	}

	config.DB.Create(&d1)
	config.DB.Create(&d2)
	config.DB.Create(&d3)
	config.DB.Create(&d4)
	config.DB.Create(&d5)
}

func SetRewardsShop(w http.ResponseWriter, r *http.Request) {

	d1 := models.RewardsShop{
		IdRewardShop: 1,
		Name:         "Cuadro",
		ParamX:       20.0,
		ParamY:       5.0,
		ParamaZ:      -40.0,
		RotationY:    -1.6,
		Price:        100,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670868376/ducktor-shop/Cuadro_ldmffk.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/desk_fpun8g.png",
	}

	d2 := models.RewardsShop{
		IdRewardShop: 2,
		Name:         "Dinosaurio",
		ParamX:       50.0,
		ParamY:       1.5,
		ParamaZ:      -40.0,
		Price:        200,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788092/ducktor-shop/cat_wrdf4p.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/cat_xdy80e.png",
	}

	d3 := models.RewardsShop{
		IdRewardShop: 3,
		Name:         "Planta",
		ParamX:       10.0,
		ParamY:       1.0,
		ParamaZ:      -40.0,
		Price:        100,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788091/ducktor-shop/present_zaubnh.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/present_kctnqr.png",
	}

	d4 := models.RewardsShop{
		IdRewardShop: 4,
		Name:         "Cohete",
		ParamX:       40.0,
		ParamY:       2.0,
		ParamaZ:      -35.0,
		Price:        1000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788091/ducktor-shop/rocket_is23zs.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/rocket_t7veyc.png",
	}

	d5 := models.RewardsShop{
		IdRewardShop: 5,
		Name:         "Muñeco",
		ParamX:       30.0,
		ParamY:       5.0,
		ParamaZ:      -20.0,
		Price:        300,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/clock_nessi6.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/clock_sylvex.png",
	}

	d6 := models.RewardsShop{
		IdRewardShop: 6,
		Name:         "Moñigo",
		ParamX:       15.0,
		ParamY:       4.0,
		ParamaZ:      -38.0,
		Price:        7000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788091/ducktor-shop/headphones_bqo3gy.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/headphones_cpn9i1.png",
	}
	d7 := models.RewardsShop{
		IdRewardShop: 7,
		Name:         "book",
		ParamX:       25.0,
		ParamY:       0.0,
		ParamaZ:      0.0,
		Price:        7000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/book_wfo1pm.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788088/ducktor-shop/img_items/book_n2ttv7.png",
	}
	d8 := models.RewardsShop{
		IdRewardShop: 8,
		Name:         "Moñigo3",
		ParamX:       0.0,
		ParamY:       0.0,
		ParamaZ:      0.0,
		Price:        7000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/ball_ddikxu.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/ball_z9aobh.png",
	}

	config.DB.Create(&d1)
	config.DB.Create(&d2)
	config.DB.Create(&d3)
	config.DB.Create(&d4)
	config.DB.Create(&d5)
	config.DB.Create(&d6)
	config.DB.Create(&d7)
	config.DB.Create(&d8)

}

func SetRewardsUser(w http.ResponseWriter, r *http.Request) {

	d1 := models.RewardsUsers{
		IdData:       1,
		Id_User:      "1",
		Name:         "Cuadro",
		IdRewardShop: 1,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1669653667/Duckdoc_xzayat.glb",
	}

	config.DB.Create(&d1)
}

func ExempleStateDucktor(w http.ResponseWriter, r *http.Request) {

	d1 := models.StatesDucktorDaily{

		Id:      1,
		IdUser:  "1",
		Date:    time.Now(),
		Good:    0,
		Average: 0,
		Danger:  0,
	}

	config.DB.Create(&d1)
}
