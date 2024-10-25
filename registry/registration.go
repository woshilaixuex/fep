package registry
/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 服务注册 
 * @Date: 2024-10-17 17:54
 */
type ServiceName string
type Registration struct{
    ServiceName ServiceName
    ServiceURL string
}

const (
    LogService = ServiceName("LogService")
)
