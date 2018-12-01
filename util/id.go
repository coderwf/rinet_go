package util

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
)

/*生成一个随机的种子*/

func Seed() (seed int64,err error){
	err = binary.Read(rand.Reader,binary.BigEndian,&seed)
	return
}

/*产生一个随机的id*/
func RandId(idlen int)(id string){
	b := make([]byte,idlen)
	var i int
	//共8byte
	randVar  := mrand.Uint64()
	for i = 0 ;i <idlen ;i++{
       b[i]  = byte(randVar & 0xFF)
       randVar  >>= 8
       if i % 8 == 0 {
       	//重新生成一个值
			randVar  = mrand.Uint64()
		}//if
	}//for
	id = fmt.Sprintf("%x",b)
	return
}

/*产生一个安全的随机的id*/
func SecureRandId(idlen int)(id string , err error){
	b := make([]byte,idlen)
	n , err := rand.Read(b)
	if n != idlen{
		err = fmt.Errorf("Only generated %d random bytes, %d requested",n,idlen)
		return
	}
	if err != nil{
		return
	}//if
	return
}

