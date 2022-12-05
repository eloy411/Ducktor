package routes

import (
	// "bytes"
	// "context"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// "io/ioutil"
	"log"
	"net/http"

	// "os"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
	// "github.com/eloy411/project-M12-BACK/config"
	// "github.com/eloy411/project-M12-BACK/models"
)

func IniPainting(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**GET INFO FROM USER*/
	var user models.GetDibujo
	err := json.NewDecoder(r.Body).Decode(&user)
	
	/**LIMITE DE DIBUJOS ===> si llega al total de dibujos se reinicia*/
	
	if (user.IdDibujo == 5) {
		user.IdDibujo = 0
	}
	
	/**GET INFO FROM DIBUJO*/
	var result []models.Dibujos
	
	config.DB.Table("dibujos").Select("*").Where("Id_Dibujo = ?",user.IdDibujo+1).Scan(&result)

	/**CONFIG RESPONSE*/
	jsonResp, err := json.Marshal(&result)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	/**RESPONSE*/
	w.Write(jsonResp)
}





func SavePaint(w http.ResponseWriter, r *http.Request) {

	/**request for multipart form data*/
   err := r.ParseMultipartForm(0<<30)

   if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	return
}

	/** DATOS DEL DIBUJO*/
	var file models.FileData
	var dibujo models.PaintingDaily

	dibujo.IdDibujo = r.PostFormValue("IdDibujo") 
	dibujo.NombreDibujo =  r.PostFormValue("NombreDibujo")
	dibujo.IdUser = r.PostFormValue("IdUser")
	file.MyFile = r.PostFormValue("MyFile")

	res1 := strings.Split(file.MyFile, ",")

	fmt.Println(dibujo.IdDibujo)
	file.MyFile = res1[1]
	
	/** CAPTAR LA IMAGEN*/

	tempFile, err := ioutil.TempFile("temp-images","-*.png")

	defer tempFile.Close()

	dec, _ := base64.StdEncoding.DecodeString(file.MyFile)
	tempFile.Write(dec)

	/*** SUBIR LA IMAGEN A CLOUDINARY*/
	resp,err := config.CLD.Upload.Upload(context.Background(), tempFile.Name(), uploader.UploadParams{})

	// /** guardamos en el struck la respuesta que es el url*/
	dibujo.URL_Dibujo = resp.SecureURL 
	
	// /**Borramos el archivo temporal*/
	nombreDelarchivo := tempFile.Name()
	tempFile.Close()

	eerr := os.Remove(".\\"+nombreDelarchivo) 

	if eerr != nil {
		fmt.Println(eerr)
		return
	}else{
		fmt.Println("Archivo borrado")
	}

	/**REGISTRO EN LA BD LA IMAGEN*/
	config.DB.Create(&dibujo)

	tempFile.Close()
	/**updateamos el numero de dibujos que ha hecho el usuario*/
	var user models.User
	config.DB.Model(&user).Where("Id_User = ?",dibujo.IdUser).Update("NumDibujos",dibujo.IdDibujo)

	
}


func GetPaintDaily(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**GET INFO FROM USER*/
	var user models.InfoUser
	err := json.NewDecoder(r.Body).Decode(&user)


	fmt.Println(user.IdUser)
	/**GET INFO FROM DIBUJO*/
	var result []models.PaintingDaily
	
	config.DB.Table("painting_dailies").Select("*").Where("Id_User = ?",user.IdUser).Scan(&result)

	/**CONFIG RESPONSE*/
	jsonResp, err := json.Marshal(&result)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	/**RESPONSE*/
	w.Write(jsonResp)
}