/**
* @Author: mo tan
* @Description:
* @Date 2021/1/10 9:37
 */
package domain

type PingDo struct {
	Name  string `mapstructure:"name" json:"name"`
	Age   int    `mapstructure:"age" json:"age"`
	Email string `mapstructure:"email" json:"email"`
}
