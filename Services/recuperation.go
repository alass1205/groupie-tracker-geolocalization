package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetData(url string, data interface{}) error {
	//envoyer le requete au serveur et recuperer  la reponse
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ici", err)
		return err
	}

	//pour s'assurer que la ressource du corps de la reponse HTTP est correctement fermee apres son utilisation
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("la requête a retourné un code d'état %d", resp.StatusCode)
	}

	//Decoder les donnees JSON dans la structure fournie
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}

	return nil
}
