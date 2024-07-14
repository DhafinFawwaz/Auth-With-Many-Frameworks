package config

func JWT_KEY() string {
	return GetEnv("JWT_KEY")
}

func DB_PASSWORD() string {
	return GetEnv("DB_PASSWORD")
}

func DB_USERNAME() string {
	return GetEnv("DB_USERNAME")
}

func DB_PORT() string {
	return GetEnv("DB_PORT")
}

func DB_NAME() string {
	return GetEnv("DB_NAME")
}

func DB_HOST() string {
	return GetEnv("DB_HOST")
}

func DB_TYPE() string {
	return GetEnv("DB_TYPE")
}
