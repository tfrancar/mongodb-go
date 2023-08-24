package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Para maiores detalhes assista ao vídeo https://www.youtube.com/watch?v=TtnI6nPhzgQ - Canal HunCoding - Prof. Otavio Celestino

// Definição da estrutura do dado que sera salva no banco de dados.
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int32              `bson:"age"`
}

func main() {
	//Informações importantes sobre  do banco de dados
	username := "seu_usuario" // usuário do seu banco de dados
	password := "sua_senha"   // Senha do seu banco de dados
	host := "seu_local_host"  // nome do host que o banco está hospedado.
	port := 27017             // Porta de conexão com o banco de dados
	authDatabase := "admin"   // Autenticação no banco de dado
	// databaseName := "nome_do_seu_banco"

	//Aqui a conexão esta sendo construída
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		username, password, host, port, authDatabase)

	//Execução da conexão com o banco de dados
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// user := User{
	// 	Name: "Kelma",
	// 	Age:  30,
	// }

	collection := client.Database("users").Collection("userData")

	// result, err := collection.InsertOne(context.Background(), user)
	// if err != nil {
	// 	panic(err)
	// }

	//Fazendo um find com o filter
	// filter := bson.D{{"name", user.Name}}
	// userResult := User{}
	// errFinding := collection.FindOne(context.Background(), filter).Decode(&userResult)
	// if errFinding != nil {
	// 	panic(errFinding)
	// }

	// fmt.Println(result)
	// fmt.Println(userResult)

	//Criando um find sem filter
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {

		userResult := User{}

		err := cur.Decode(&userResult)
		if err != nil {
			panic(err)
		}

		fmt.Println(userResult)
	}

	//Encerrando a conexão com o banco de dados
	err = client.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão encerrada.")
}
