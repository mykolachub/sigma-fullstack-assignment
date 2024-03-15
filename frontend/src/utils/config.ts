export default {
  env: {
    apiKey: process.env.REACT_APP_API_URL || 'http://localhost:8080/api',
    jwtSecret: process.env.REACT_APP_JWT_SECRET || 'JWT_SECRET',
  },
};
