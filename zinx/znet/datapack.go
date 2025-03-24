package znet

/**
封包拆包模块
直接面向TCP连接中的数据流，把TCP的数据流封装成包，在拆包时，
按照 zinx 的格式拆包，再把包在封装到 Message 中，然后传递给 api层
*/
