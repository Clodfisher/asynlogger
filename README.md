# asynlogger

### 创建意图    

对于系统程序运行过程中对于日志的产生时比不可少的，通过日志可以知道程序的健康状况。对于公司中日志生产的统一标准很有必要，方便开发人员的查找和运维人员的维护。    

### 达到目标

此项目最终达到的目标，是实现方便开发人员统一调用的日志生成器。    
主要实现的功能如下：      
* 可以选择日志生成级别     
1 .Debug级别：用于调试程序，日志最为详细，对于程序的性能影响比较大。非错误性。          
2 .Trace级别：用于追踪问题。非错误性。          
3 .Info级别：打印程序运行过程中比较重要信息，比如访问日志。非错误性。      
4 .Warn级别：警告日志，说明程序运行出现潜在的问题。错误性。           
5 .Error级别：错误日志，程序运行发生错误，但不影响程序的运行。错误性。      
6 .Fatal级别：严重错误，发生的错误会导致程序退出。错误性。       

* 日志存储的位置     
1 .直接打印到控制台。        
2 .打印到文件里。将错误性和非错误性的日志分别打印到不同的文件内。            

* 日志文件的备份迁移     
1 .根据时间进行备份迁移。    
2 .根据文件的大小进行备份迁移。    

* 异步日志生成     
1 .对于文件写入和控制台输出进行异步处理。        
2 .非核心代码异步化。    

* 统一出错处理    
1 .对于项目开发中出现的错误进行统一处理，提高容错性。    

### Installing    
```
go get github.com/Clodfisher/asynlogger...
```