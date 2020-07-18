/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 01:55:32
 * @LastEditTime : 2020-07-14 02:29:10
 * @Description  : token
 */ 
 
 package user
 
 import(
	"micro_demo/comm/jwt"
	"micro_demo/comm/logging"
 )

 // GenerateToken 生成token
 func (s *service) GenerateToken(userId uint64, expireDate int) (string, error){
	token := ""
	token ,err := jwt.GenerateToken(userId,30)
	if err !=nil {
		logging.Logger().Error(err)	
		return token, err
	}

	return token,nil
}  


 
// ParseToken 解析token
 func (s *service)ParseToken(token string) (*jwt.Claims, error){
	claims ,err := jwt.ParseToken(token)
	if err !=nil {
		logging.Logger().Error(err)	
		return nil, err
	}
	
	return claims,nil
 } 				 