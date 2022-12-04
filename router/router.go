func route() {
	router := gin.Default()

	v0 := route.Group("/user"){
		v0.GET("/",)
	}
	v1 := router.Group("/ping"){

	}

	
}