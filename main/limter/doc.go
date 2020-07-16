/*
Package dislimit 基于redis+lua+令牌桶的分布式限流组件.

支持按照接口限流、按照uid+接口限流、按照IP限流
rate 每秒限制并发数
capacity 可以突发的流量-令牌桶支持突发流量
一般配置rate与capacity一样即可

*/
package dislimit
