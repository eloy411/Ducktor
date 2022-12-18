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
		Gravedad:     1, 
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
		ParamX:       3.0,
		ParamY:       4.0, 
		ParamaZ:      -5.0,
		RotationY:    -1.6,
		Price:        100, 
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358915/ducktor-shop/cuadro_dengri.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1671358833/ducktor-shop/cuadro_rh27eb.png",
	}

	d2 := models.RewardsShop{
		IdRewardShop: 2,
		Name:         "Gato",
		ParamX:       5.0,
		ParamY:       0.3,
		ParamaZ:      0.0,
		Price:        200,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358367/ducktor-shop/catglb_bnw0p6.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/cat_xdy80e.png",
	}

	d3 := models.RewardsShop{
		IdRewardShop: 3,
		Name:         "Reloj",
		ParamX:       -25.0,
		ParamY:       7.0,
		ParamaZ:      -16.0,
		Price:        100,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358341/ducktor-shop/reloj_h7ccas.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/clock_sylvex.png",
	}

	d4 := models.RewardsShop{
		IdRewardShop: 4, 
		Name:         "Cohete",
		ParamX:       19.0,
		ParamY:       0.35,  
		ParamaZ:      -14.0, 
		Price:        1000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1670788091/ducktor-shop/rocket_is23zs.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/rocket_t7veyc.png",
	}

	d5 := models.RewardsShop{
		IdRewardShop: 5,
		Name:         "Pelota",
		ParamX:       8.0,
		ParamY:       0.75,
		ParamaZ:      0.0,
		Price:        300,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358360/ducktor-shop/ball_okz7ub.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788089/ducktor-shop/img_items/ball_z9aobh.png",
	}

	d6 := models.RewardsShop{
		IdRewardShop: 6,
		Name:         "Auriculares",
		ParamX:       -10.0,
		ParamY:       0.0,
		ParamaZ:      0.0,
		Price:        7000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358827/ducktor-shop/auriculares_duj5up.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/headphones_cpn9i1.png",
	}
	d7 := models.RewardsShop{
		IdRewardShop: 7,
		Name:         "book",
		ParamX:       14.0,
		ParamY:       3.0, 
		ParamaZ:      -13.0,
		Price:        700,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358372/ducktor-shop/book_vopdsh.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788088/ducktor-shop/img_items/book_n2ttv7.png",
	}
	d8 := models.RewardsShop{ 
		IdRewardShop: 8,
		Name:         "Regalo", 
		ParamX:       19.5,
		ParamY:       2.7, 
		ParamaZ:      4.0,
		Price:        7000,
		Url:          "https://res.cloudinary.com/eloy411/image/upload/v1671358349/ducktor-shop/regalo_dvh9uo.glb",
		UrlImagen:    "https://res.cloudinary.com/eloy411/image/upload/v1670788090/ducktor-shop/img_items/present_kctnqr.png",
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
