declare var process : {
    env: {
        AUTH_SECRET: string
        API_URL: string
    }
}

export const ENV = {
    AUTH_SECRET: (() => process.env.AUTH_SECRET)() || '123456',
    API_URL: (() => process.env.API_URL)() || 'http://localhost:8080/',
}
