package main

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"log"
	"encoding/json"
)

var RedisPort = "8769"
var RedisIP = "140.115.153.185"
var Password = "mwnlmwnl"

func Redis_IDtoMAC(ID string)(CONTENT []string){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
		CheckError(err)
		defer c.Close()


		c.Do("AUTH",Password)
		_, err2 := c.Do("SELECT", "2")
		if err2 != nil{
			CONTENT =append(CONTENT,"you don't have new url1")
			return
		}

		MAC, err2 := redis.String(c.Do("GET", ID))
		if err2 != nil{
			CONTENT =append(CONTENT,"錯誤:ID與MAC未正確輸入")
			return
		}

		_, err2 = c.Do("SELECT", "0")
		if err2 != nil{
			CONTENT =append(CONTENT,"you don't have new url3")
			return
		}

		binary, err2 := redis.Bytes(c.Do("GET", MAC))
		if err2 != nil{
			CONTENT =append(CONTENT,"錯誤:你沒有新資料!")
			return
		}

		user := new(USER_MAC)
		json.Unmarshal(binary,&user)

		var count int
		count = 0
		for i:=len(user.CRAWLER);i>0 && count<3 ; i--{
			for j:=0;j< len(user.CRAWLER[i-1].GOOGLE);j++{
				// CONTENT = CONTENT+user.CRAWLER[i-1].GOOGLE[j].CONTENT+"\n\n"
				CONTENT  = append(CONTENT,user.CRAWLER[i-1].GOOGLE[j].CONTENT)
				count ++
			}
		}
		if CONTENT ==[]{
			CONTENT =append(CONTENT,"錯誤:你沒有新資料")
			return
		}
		return CONTENT

}

func Redis_Get(KEY_NAME string) ([]byte, error){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()

	c.Do("AUTH",Password)
	v, err2 := redis.Bytes(c.Do("GET", KEY_NAME))
	return v, err2
}

func Redis_Set(KEY_NAME string, VAL []byte) error{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()

		c.Do("AUTH",Password)
	_, err2 := c.Do("SET", KEY_NAME, VAL)
	return err2
}

func Redis_Del(KEY_NAME string) error{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	c.Do("AUTH",Password)
	_, err2 := c.Do("DEL", KEY_NAME)
	CheckError(err2)
	return err2
}


// func Redis_DelAllUser()error{
// 	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
// 	c, err := redis.Dial("tcp", RedisIPPORT)
// 	CheckError(err)
// 	defer c.Close()
//
// 	c.Do("AUTH",Password)
// 	user := Redis_AllUser()
// 	var err2 error = nil
// 	for i:=0 ; i<len(user) ; i++{
// 		_, err2 = c.Do("DEL", user[i])
// 		CheckError(err2)
// 		if err2 != nil{
// 			return err2
// 		}
// 	}
// 	return err2
// }


//this will find all user with a strings slice
func Redis_AllUser()([]string){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	c.Do("AUTH",Password)
	user, err2 := redis.Strings(c.Do("KEYS", "*"))
	CheckError(err2)
	return user
}

//if not find, return is fail
func Redis_FindUser(USER string)(bool){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	c.Do("AUTH",Password)
	EXIST, err2 := redis.Bool(c.Do("EXISTS", USER))
	CheckError(err2)
	return EXIST
}



func CheckError(err error) {
    if err  != nil {
        log.Println("Error: " , err)
        // os.Exit(0)
    }
}
