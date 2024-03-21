export interface UserDTO {
  id: string;
  email: string;
  password: string;
  role: string;
}

export interface UserCreateDTO {
  email: string;
  password: string;
  role: string;
}

export interface UserLoginDTO {
  email: string;
  password: string;
}

export interface UserTokenDTO {
  token: string;
}
